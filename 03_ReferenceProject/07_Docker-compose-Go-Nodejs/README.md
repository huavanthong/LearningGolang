This is the example reference project for learning Docker-compose with Golang [Refer blog post](https://viblo.asia/p/docker-compose-la-gi-kien-thuc-co-ban-ve-docker-compose-1VgZv8d75Aw)
## Overview
It include: 
* Server is implemented by Golang
* Web is implemented by ReadJS (NodeJS).
* Dockerfile to build image for each part.
* Docker compose will combine all parts together.
## Dependencies 
Please make sure you install all the necessary software below:  
* NodeJS on Windows
* Docker on Windows
* Golang compiler

## How to build this project
### Step 1: Initialize environment for NodeJS
Set up your web application:
```
$ nmp init react-app app
```
### Step 2: Use docker compose to start all service
Build this project by .yml docker-compose.
```
$ docker-compose up
```
### Step 3: Run our application.
```

```
