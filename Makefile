.PHONY: protos
protos:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/hello.proto

clean:
	rm protos/*.go

api: clean protos
	docker build -t siyangzhang/hello-go-img -f exec/api/Dockerfile .