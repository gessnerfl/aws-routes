    
export CGO_ENABLED:=0
export GO111MODULE=on
#export GOFLAGS=-mod=vendor

VERSION=$(shell git describe --tags --match "v*" --always --dirty)

all: build test vet lint fmt

build: clean bin/awsroutes

bin/awsroutes:
	@echo "+++++++++++  Run GO Build +++++++++++ "
	@go build -o $@ github.com/gessnerfl/awsroutes

test:
	@echo "+++++++++++  Run GO Test +++++++++++ "
	@go test ./... -cover

vet:
	@echo "+++++++++++  Run GO VET +++++++++++ "
	@go vet -all ./...

lint:
	@echo "+++++++++++  Run GO Lint +++++++++++ "
	@golint -set_exit_status `go list ./...`

fmt:
	@echo "+++++++++++  Run GO FMT +++++++++++ "
	@test -z $$(go fmt ./...) 

update:
	@GOFLAGS="" go get -u
	@go mod tidy

vendor:
	@go mod vendor

clean:
	@echo "+++++++++++  Clean up project +++++++++++ "
	@rm -rf bin