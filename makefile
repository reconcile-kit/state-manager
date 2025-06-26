.PHONY: api gen-swagger build

api: gen-swagger

gen-swagger:
	docker run --rm \
	  -v "$(PWD)":/src \
	  -w /src \
	  golang:1.24 \
	  sh -c 'go install github.com/swaggo/swag/cmd/swag@latest && \
	         $$(go env GOPATH)/bin/swag init \
	             -g cmd/app/main.go \
	             --parseDependency'

build:
	docker build -t state-manager .