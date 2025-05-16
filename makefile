.PHONY: api

api: gen-swagger openapi-client

gen-swagger:
	docker run --rm \
	  -v "$(PWD)":/src \
	  -w /src \
	  golang:1.24 \
	  sh -c 'go install github.com/swaggo/swag/cmd/swag@latest && \
	         $$(go env GOPATH)/bin/swag init \
	             -g cmd/app/main.go \
	             --parseDependency'

openapi-client:
	docker run --rm \
	  -v "$(PWD):/src" \
	  openapitools/openapi-generator-cli generate \
	    -i /src/docs/swagger.json \
	    --skip-validate-spec \
	    -g go \
	    -o /src/pkg/stateclient \
	    --package-name stateclient \
	    --additional-properties=withGoMod=false \

