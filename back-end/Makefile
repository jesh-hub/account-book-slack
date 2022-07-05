.PHONY: build clean serve

export APP_NAME=abs
export BIN_DIR=bin

build: clean
	GOOS=linux GOARCH=amd64 go build -ldflags '-d -s -w' -o ${BIN_DIR}/${APP_NAME} main.go

clean:
	rm -rf bin/*

serve:
	cd cmd && go run app.go