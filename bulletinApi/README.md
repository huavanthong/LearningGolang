
# bulletinApi
Bulletin board scalable API ready to ride on Kubernetes

Check out [this video](https://www.youtube.com/watch?v=pkZrgHxJ130&t=1207s) 
to see the whole deployment process of this API on Google Cloud Kubernetes Engine

Disclaimer: this package is created for educational purposes only

### Objectives
    - Implement Go server following microservice architecture.
    - Implement Kubernets to scale out application.
    - Use Docker compose to combine app and PostgreSQL Database together.

### Getting Started
Firstly, you need to start docker compose
```
docker-compose up
```

```bash
$ docker-compose up --force-recreate --build -d
```
