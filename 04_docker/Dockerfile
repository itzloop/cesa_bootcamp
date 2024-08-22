# Use the official Golang image as the base image
FROM golang:1.23.0-alpine

# set the current working directory inside the container
WORKDIR /app

# copy the source code from the host machine to the container
COPY . .

# compile the go application
RUN go build -o hello-docker .

CMD ["./hello-docker"]