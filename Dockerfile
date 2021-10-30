FROM golang:alpine AS build
WORKDIR /app
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . /app
RUN go build -o main .

FROM alpine
COPY --from=build /app/main /app/main
ENTRYPOINT ["/app/main"]