.PHONY: build deploy deploy-all clean serve

export APP_NAME=abs
export BIN_DIR=bin

build: clean
	GOOS=linux GOARCH=amd64 go build -ldflags '-d -s -w' -o ${BIN_DIR}/${APP_NAME} main.go

deploy: build
	sls deploy -f ${APP_NAME} --verbose

deploy-all: build
	sls deploy --verbose

clean:
	rm -rf bin/*

serve:
	go run cmd/app.go