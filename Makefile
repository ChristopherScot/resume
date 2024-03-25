generate: 
	swagger generate server -f swagger.yaml


deploy-local: build-local cleanup
	sam local start-api -t ./deploy/sam.yml 

build-local:
	./deploy/build.sh

cleanup:
	go mod tidy

deploy-dev: generate build-local  cleanup
	./deploy/create-api-gateway-swagger.sh
	sam deploy -t ./deploy/sam.yml --config-file ./samconfig.toml --resolve-s3
	
