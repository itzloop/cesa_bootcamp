# Docker Workshop

## Build the docker image

```sh
docker build -t docker-hello:latest .
```

## List of Images

```sh
docker images
```

Remove an image:

```sh
docker rmi <image_id>
```

## Run the image

```sh
docker run docker-hello:latest
```

Run in detached mode:

```sh
docker run -d docker-hello:latest
```

Run in a specific port:

```sh
docker run -p <host_port>:<container_port> docker-hello:latest
```

Give a name to your container:

```sh
docker run --name dh docker-hello:latest
```

Run the container with an env:

```sh
docker run --env <YOUR_ENV>=<VALUE> docker-hello:latest
```

## Display the list of containers

```sh
docker ps
```

Show all the containers (stopped and running):

```sh
docker ps -a
```

## Stop and Remove the container

Stop a running container:

```sh
docker stop <container_name>
```

Remove a container:

```sh
docker rm <container_name>
```

## DockerHub Registry

Login to DockerHub:

```sh
docker login -u <username>
```

Change the image tag:

```sh
docker tag hello-docker:latest <username>/hello-docker:latest
```

Push the docker image to DockerHub:

```sh
docker push <username>/hello-docker:latest
```

Pull the docker image from DockerHub:

```sh
docker pull <username>/hello-docker:latest
```
