.PHONY: build all test clean

all: build

docker-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" ./src/cmd/backend/main.go

docker-run:
	docker-compose -f ./docker-compose.yml up

docker-database:
	docker-compose -f ./docker-compose-database.yml up 

run :
	go run ./src/cmd/backend/main.go

test:
#	测试项
	PROJECT_PWD=$(shell pwd) go test -race -v ./src/storage

test-cover:
	go test -cover --count=1 ./...

cover-total:
	go test -cover --count=1 ./... -coverprofile cover.out
	go tool cover -func cover.out | grep total 

cov:
	PROJECT_PWD=$(shell pwd) go test -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

fmt:
	@gofmt -w $(shell find . -type f -name '*.go' -not -path './*_test.go')

fmt-check:
	@gofmt -l $(shell find . -type f -name '*.go' -not -path './*_test.go')

clean:
	@ro -fR bin
