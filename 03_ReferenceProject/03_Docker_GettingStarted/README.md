# Docker Getting Started Tutorial

This tutorial has been written with the intent of helping folks get up and running
with containers and is designed to work with Docker Desktop. While not going too much 
into depth, it covers the following topics:

- Running your first container
- Building containers
- Learning what containers are running and removing them
- Using volumes to persist data
- Using bind mounts to support development
- Using container networking to support multi-container applications
- Using Docker Compose to simplify the definition and sharing of applications
- Using image layer caching to speed up builds and reduce push/pull size
- Using multi-stage builds to separate build-time and runtime dependencies

## Getting Started

If you wish to run the tutorial, you can use the following command after installing Docker Desktop:

```bash
docker run -d -p 80:80 docker/getting-started
```

Once it has started, you can open your browser to [http://localhost](http://localhost).


## Issue knowledge  
1. Port 8080 is already allocated
```
If you have already run application before. And you want yo run it again without remove port 8080.
$ docker run -d -p 80:80 docker/getting-started
..... forgot to remove this container 
$ docker run -d -p 80:80 docker/getting-started
=> error: 
docker: Error response from daemon: driver failed programming external connectivity on endpoint docker-tutorial (cf3d117a14250f7bcd050454f08a9d65aeac10e247279a3adea5d3042b2a7732): Bind for 0.0.0.0:80 failed: port is already allocated.
```

## Development

This project has a `docker-compose.yml` file, which will start the mkdocs application on your
local machine and help you see changes instantly.

```bash
docker-compose up
```

## Contributing

If you find typos or other issues with the tutorial, feel free to create a PR and suggest fixes!

If you have ideas on how to make the tutorial better or new content, please open an issue first before working on your idea. While we love input, we want to keep the tutorial  scoped to newcomers.
As such, we may reject ideas for more advanced requests and don't want you to lose any work you might
have done. So, ask first and we'll gladly hear your thoughts!
