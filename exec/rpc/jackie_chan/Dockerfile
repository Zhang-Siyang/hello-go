FROM golang:alpine AS build
WORKDIR /app
RUN apk add git
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . /app
RUN go build -ldflags "-X hello/define.BinaryVersion=$(git rev-parse --short HEAD) -X hello/define.binaryBuildTimeString=$(date +%s)" -o hello-api ./exec/api/main.go

FROM alpine
# alpine 镜像默认使用 UTC 时区
COPY --from=build /app/hello-api /app/hello-api
EXPOSE 80
ENTRYPOINT ["/app/hello-api"]
