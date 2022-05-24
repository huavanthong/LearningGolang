# Introduction

### Reference
1. Example about swagger with gin framework. Refer: [here](https://levelup.gitconnected.com/tutorial-generate-swagger-specification-and-swaggerui-for-gin-go-web-framework-9f0c038483b5)


## Getting Started
To run this app
```
go run .\ginsimple\main.go
```

## Swagger
### Install 
Install swagger for documentation
```
go get -v github.com/swaggo/swag/cmd/swag
go get -v github.com/swaggo/gin-swagger
go get -v github.com/swaggo/files
```
### Generate
Generate document by swagger.

* **Windows:**
```
.\swagger.bat init -g ginsimple/main.go --output docs/ginsimple
```

* **Linux:**
```
swag init -g ginsimple/main.go --output docs/ginsimple
```