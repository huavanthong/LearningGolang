# Introduction
In this chapter, you’ll learn how to deploy to
* A server that’s either fully owned by you, physical, or virtual (on an Infrastructure-as-a-Service provider, Digital Ocean)
* A cloud Platform-as-a-Service (PaaS) provider, Heroku
* Another cloud PaaS provider, Google App Engine
* A dockerized container, deployed to a local Docker server, and also to a virtual machine on Digital Ocean
**Notes:** 
- It’s important to remember that there are many different ways of deploying a web app and in each of the methods you’ll be learning in this chapter, there are many variations.  
- The methods described in this chapter are based on a single person deploying the web app. The processes are usually a lot more involved in a production environment, which includes additiona tasks like running test suites, continual integration, staging servers, and so on.
- This chapter also introduces many concepts and tools, each of them worthy of an entire book. Therefore, it’s impossible to cover all of these technologies and services. This chapter aims to cover only a small portion, and if you want to learn more, take this as a starting point.

# Deploying to servers
* [What steps to deploy a Go web application to server?](#Steps)
* [How to set up driver PostgreSQL for web-app?](#Set-up-PostgreSQL)
# Deploying to Google App Engine

# Deploy to Dokcer
* [What is Docker?](#What-is-Docker)
* [What is Docker machine?](#Docker-machine)
* [What is Digital Ocean?](#Digital-Ocean)
* [Why once we push local Docker host to cloud-provider, we can't any images](#Develop-images)

# Comparison of deployment methods
* [Compare the advantages in distinct servers methods? Part 10.5](#https://edu.anarcho-copy.org/Programming%20Languages/Go/Go%20Web%20Programming.pdf)

## Deploying to servers
### Steps
Step to deploy a web app to server:
* Create an executable binary and the running it off a server that's on the internet.
* 

### Getting Started
#### Set-up-PostgreSQL
Create a user named gwp.
```
createuser -P -d gwp
```
Create a database named gwp.
```
createdb gwp
```
Create schema and table data
```
psql -U gwp -f setup.sql -d gwp
```
#### Insert data 
Add records to Database
```
psql -U gwp -f data.sql -d gwp
```
#### Building web app
Set up web-app to account Database
```
From 
    Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
To  
    Db, err = sql.Open("postgres", "user=postgres dbname=gwp password=1234 sslmode=disable")
```

To buidl a web server
```
go build
```
#### Run
To run it on localhosts
**GET:**
```
http://localhost:8080/post/1
```

**POST:**
```
http://localhost:8080/post/1
```

**PUT:**
```
http://localhost:8080/post/1
```

**DELETE:**
```
http://localhost:8080/post/1
```
#### Deploy to server
To run the web service, log into the server and run it from the console:
```
./ws-s
```

## Deploying to docker
### What-is-Docker?
To understand about Docker, refer: [here](#https://www.meisternote.com/app/note/loBzTQ2yKbv4/docker-concepts-and-components)

### Dockerizing a GO web application
#### Build-images 
To build image for this web-app
```
docker build –t ws-d .
```
#### Run image
To run
```
docker run --publish 80:8080 --name simple_web_service --rm ws-d
```
#### Containers
To see weather your container has been created.
```
>docker ps
CONTAINER ID IMAGE ... PORTS NAMES
eeb674e289a4 ws-d ... 0.0.0.0:80->8080/tcp simple_web_service
```
#### Run app on Docker
**POST:**
```
curl -i -X POST -H "Content-Type: application/json" -d '{"content":"My first post","author":"Sau Sheong"}' http://127.0.0.1/post/
```

**GET:**
```
curl -i -X GET http://127.0.0.1/post/1
```
### Pushing your Docker container to the internet
Dockerizing the simple web service sounds great, but it’s still running locally. What you want is to have it running on the internet. There are a number of ways of doing this, but using the Docker machine is probably the most simple (at the moment, because Docker is still evolving).
#### Docker-machine
* **Docker machine** is a command-line interface that allows you to create Docker hosts, either locally or on cloud providers, both public and private. 
* As of this writing, the list of public cloud providers include:
    - AWS  
    - Digital Ocean
    - Google Compute Engine
    - IBM Softlayer
    - Microsoft Azure 
    - Rackspace
    - Exoscale
    -  VMWare vCloud Air.
It can also create hosts on private clouds, including clouds running on OpenStack, VMWare, and Microsoft Hyper-V (which covers most of the private cloud infrastructure to date).
#### Download
Docker Machine isn’t installed along with the main Docker installation; you need to install it separately.
**Method 1:** To install it on Windows, clone this repository
```
https://github.com/docker/machine
```

**Method 2:** Downloading the binary for your platform from
```
https://docs.docker.com/machine/install-machine
``` 

**Method 3:** For Linux, we can use command
```
curl -L https://github.com/docker/machine/releases/download/v0.3.0/docker-machine_linux-amd64 /usr/local/bin/docker-machine
```

Then make it executable:
```
chmod +x /usr/local/bin/docker-machine
```
### Cloud-providers
#### Digital-Ocean
Once you’ve downloaded Docker Machine and made it executable, you can use it to create a Docker host in any of the cloud providers. One of the easiest is probably **Digital Ocean**, a virtual private server (VPS) provider known for its ease-of-use and low cost. (A VPS is a VM sold as a service by a provider.) 
In May 2015, Digital Ocean became the second-largest hosting company in the world, in terms of web-facing servers, after AWS.
##### How-to-use
To create a Docker host on Digital Ocean, you’ll first need to sign up for a **Digital Ocean** account. Once you have an account, go to the Applications & API page at
https://cloud.digitalocean.com/settings/applications.


##### Push-Docker-to-Digital-Ocean
To create a Docker host on Digital Ocean through Docker Machine, execute this command on the console:
```
> docker-machine create --driver digitalocean --digitalocean-access-token<tokenwsd>
Creating CA: /home/sausheong/.docker/machine/certs/ca.pem
Creating client certificate: /home/sausheong/.docker/machine/certs/cert.pem
Creating SSH key...
Creating Digital Ocean droplet...
To see how to connect Docker to this machine, run: docker-machine env wsd
```
##### Connect local Docker host on Digital Ocean
Once the remote Docker host is created, the next step is to connect to it. 
Remember, your Docker client is currently connected to the local Docker host. You need to connect it to our Docker host on Digital Ocean, called wsd. The response from Docker Machine gives you a hint how to do it. You should run:
```
> docker-machine env wsd
export DOCKER_TLS_VERIFY="1"
export DOCKER_HOST="tcp://104.236.0.57:2376"
export DOCKER_CERT_PATH="/home/sausheong/.docker/machine/machines/wsd"
export DOCKER_MACHINE_NAME="wsd"
# Run this command to configure your shell:
# eval "$(docker-machine env wsd)"
```

The command tells you the environment settings for our Docker host on the cloud. To configure the client to point to this Docker host, you need to change your environment settings to match it. Here’s a quick way to do this:
```
eval "$(docker-machine env wsd)"
```

As simple as that, you’re connected to the Docker host on Digital Ocean! How do you know that? Just run
```
docker images
```
##### Develop-images
You’ll see that there are no images listed. Remember that when you were creating the container earlier in this section you created an image locally, so if you’re still connected to the local Docker host, you should see at least one image. No image means you’re no longer connected to the local Docker host.
Because you have no image in this new Docker host, you have to create one again. Issue the same **docker build** command from earlier:
```
docker build –t ws-d .
```
After the command completes, you should see at least two images when you run docker images: the golang base image, and the new ws-d image. The final step is to run the container as you’ve run it before:
```
docker run --publish 80:8080 --name simple_web_service --rm ws-d
```
This command will create and start up a container on the remote Docker host. To prove that you’ve done it, use curl to get the post record. But wait—where’s the
server? It’s at the same IP address that was returned when you ran:
```
docker-machine env wsd
```
#### Run it
To run web-app
```
> curl -i -X GET http://104.236.0.57/post/1
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 03 Aug 2015 11:35:46 GMT
Content-Length: 69
{
    "id": 2,
    "content": "My first post",
    "author": "Sau Sheong"
}
```

## Comparison of developement methods
Pleser refer to: part 10.5 [here](https://edu.anarcho-copy.org/Programming%20Languages/Go/Go%20Web%20Programming.pdf)