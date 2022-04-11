.PHONY: build clean deploy gomodgen

build:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/payments payments/main.go

clean:
	rm -rf ./bin ./vendor

deploy-dev: clean build
	sls deploy --stage dev --verbose

deploy-prod: clean build
	sls deploy --stage prod --verbose