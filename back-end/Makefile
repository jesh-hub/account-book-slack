.PHONY: build clean docs serve

export APP_NAME=abs
export BIN_DIR=bin


build: clean
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o ${BIN_DIR}/${APP_NAME} main.go

clean:
	rm -rf bin/*

docs:
	~/go/bin/swag init --parseDependency --parseInternal

serve: docs
	cd cmd && go run app.go