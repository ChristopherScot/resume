PROJECT = resume

generate: 
	rm -rf ./restapi
	swagger generate server -f swagger.yaml


run:
	go run ./cmd/${PROJECT}-server/main.go

build-local: generate cleanup
	sam build -t ./deploy/sam.yml --config-file ./samconfig.toml 
	./deploy/build.sh

deploy-local: build-local
	sam local start-api

build-dev: generate cleanup
	sam build -t ./deploy/sam.yml --config-file ./samconfig.toml 
	./deploy/build.sh arm64

cleanup:
	go mod tidy

deploy-dev: build-local
	./deploy/create-api-gateway-swagger.sh
	sam deploy -t ./deploy/sam.yml --config-file ./samconfig.toml --resolve-s3
	
