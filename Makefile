.PHONY: build clean docker help

all: build

build:
ifeq (, $(wildcard $(shell which swag)))
	# go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
	go install github.com/swaggo/swag/cmd/swag@v1.6.5
	go get -u github.com/swaggo/gin-swagger@v1.2.0
	go get -u github.com/swaggo/files
	go get -u github.com/alecthomas/template
endif
	swag init
	go mod tidy
	@go build .

docker:
	docker build . -t simple-blog-service

clean:
	rm -f simple-blog-service
	rm -rf docs
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make docker: build deploy docker image"
	@echo "make clean: remove object files and cached files"
