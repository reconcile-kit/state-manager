### Generate swagger api
```shell
 swag init -g main.go --dir ./cmd/app,./internal/http,./internal/dto,./internal/repositories/resources --parseDependency
```