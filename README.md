
# About
Simple Gin API/Web app

### Commands
## Run project
```bash
go run main.go
# or
go build main.go && ./main
```

## Run tests
```bash
go test
# or
go test -run test-name
```
## Others
```Bash
go env
go mod init path/project-name
go mod tiny
go mod graph
go mod vendor
go mod verify
go mod why
```


### Generate DOCS
- I is importante to export \$GOPATH in terminal before running Swagger 
 ``$ export PATH=$(go env GOPATH)/bin:$PATH* ``
```bash
  swag init --parseDependency --parseInternal --parseDepth 1
```


