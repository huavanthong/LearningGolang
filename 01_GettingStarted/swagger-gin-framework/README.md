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
### Generate document
Generate document by swagger.

* **Windows:**
```
???????????
```

* **Linux:**
```
swag init -g ginsimple/main.go --output docs/ginsimple
```

### Open SwaggerUI
Open
```
http://localhost:3000/swagger/index.html
```