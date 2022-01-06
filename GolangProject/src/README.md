## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)

## General info
This project is simple Golang MVC model.
	
## Technologies
Project is created with:
* Golang version: 1.17
* Gorilla/mux version: github.com/gorilla/mux (for implement REST API)
* pgx driver: github.com/jackc/pgx/v4/stdlib (for connect PostgreSQL Admin4)

	
## Setup
To run this project, install it locally using Command Line (cmd):

```
$ set GOPATH=${Working_Space}
$ go get github.com/gorilla/mux
$ go get github.com/jackc/pgx/v4/stdlib


```
## 

## Issue knowledge
Here will summary all issue knowledge happened when you build this project.
## 1. Issue related to port a package 
#### Problem:
```
$ cd src\webapp 
$ go run main.go
=> Error: 
package webapp/bootstrap is not in GOROOT
```
#### Solution
```
$ set GOPATH=${working_space}

```
## 2. Issue related to set GOPATH variable env
#### Problem 
```
$ set GOPATH=${working_space}
$ go init mod golangMVC.com/m/v1
$ go run main.go
=> Error: 
$GOPATH/go.mod exists but should not
```
#### Solution

