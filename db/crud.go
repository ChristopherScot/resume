package db

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/ChristopherScot/resume/models"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var DBClient *dynamodb.Client
var resumeTable = os.Getenv("RESUME_TABLE")

type ResumeWithID struct {
	ID     string        `json:"id" dynamodbav:"id"`
	Resume models.Resume `json:"resume"`
}

func NewClient(sdkConfig aws.Config) *dynamodb.Client {
	return dynamodb.NewFromConfig(sdkConfig)
}

func GetAllResumes(ctx context.Context) ([]models.Resume, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(resumeTable),
	}

	result, err := DBClient.Scan(ctx, input)
	if err != nil {
		return nil, err
	}

	resumes := make([]models.Resume, 0)
	for _, item := range result.Items {
		var resume models.Resume
		err = attributevalue.UnmarshalMap(item, &resume)
		if err != nil {
			return nil, err
		}

		resumes = append(resumes, resume)
	}

	return resumes, nil
}

func GetAllResumeIDs(ctx context.Context) ([]string, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(resumeTable),
	}

	result, err := DBClient.Scan(ctx, input)
	if err != nil {
		return nil, err
	}

	ids := make([]string, 0)
	for _, item := range result.Items {
		ids = append(ids, item["id"].(*types.AttributeValueMemberS).Value)
	}

	return ids, nil
}

func GetResume(ctx context.Context, id string) (models.Resume, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		TableName: aws.String(resumeTable),
	}

	result, err := DBClient.GetItem(ctx, input)
	if err != nil {
		return models.Resume{}, err
	}

	if result.Item == nil {
		return models.Resume{}, fmt.Errorf("resume not found")
	}
	resumeItem := result.Item["Resume"].(*types.AttributeValueMemberM)

	var resume models.Resume
	err = attributevalue.UnmarshalMap(resumeItem.Value, &resume)
	slog.Info("resume", "resume", resume)

	if err != nil {
		return models.Resume{}, err
	}

	return resume, nil
}

func CreateResume(ctx context.Context, id string, resume models.Resume, metadata models.ResumeMetadata) error {
	item, err := attributevalue.MarshalMap(&ResumeWithID{
		ID:     id,
		Resume: resume,
	})
	if err != nil {
		return err
	}
	metadataMap, err := attributevalue.MarshalMap(metadata)
	if err != nil {
		return err
	}
	for k, v := range metadataMap {
		item[k] = v
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(resumeTable),
	}

	_, err = DBClient.PutItem(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func DeleteResume(ctx context.Context, id string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		TableName: aws.String(resumeTable),
	}

	_, err := DBClient.DeleteItem(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func UpdateResume(ctx context.Context, id string, resume models.Resume, metadata models.ResumeMetadata) error {
	item, err := attributevalue.MarshalMap(&ResumeWithID{
		ID:     id,
		Resume: resume,
	})
	if err != nil {
		return err
	}
	metadataMap, err := attributevalue.MarshalMap(metadata)
	if err != nil {
		return err
	}
	for k, v := range metadataMap {
		item[k] = v
	}
	update := expression.Set(expression.Name("Resume"), expression.Value(resume)).
		Set(expression.Name("Updated"), expression.Value(metadata.Updated))

	exp, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return err
	}
	input := &dynamodb.UpdateItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression:          exp.Update(),
		ExpressionAttributeValues: exp.Values(),
		ExpressionAttributeNames:  exp.Names(),
		TableName:                 aws.String(resumeTable),
	}

	_, err = DBClient.UpdateItem(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
