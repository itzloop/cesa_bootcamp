FROM golang:1.22-alpine AS build

WORKDIR /app

COPY . .

RUN GOPROXY='https://goproxy.cn/,direct' go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /uni-app ./cmd/main.go 


FROM alpine:latest

WORKDIR /

COPY --from=build /app/.env . 
COPY --from=build /uni-app /uni-app

ENTRYPOINT [ "/uni-app" ]