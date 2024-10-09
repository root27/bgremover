bgremover-cli:
	@go run main.go

server:
	@go run ./http/server.go

example-cli:
	@go run main.go -file=./examples/animal-1.jpg -out=./output/animal-1-out.jpg


##Dockerfile make commands

python-grpc-build:
	@docker build -t python-grpc -f ./Dockerfiles/Dockerfile.python_grpc_server .

python-grpc-server:
	@docker run -p 50051:50051 --name python-grpc-server python-grpc

##Gorelease make commands

release:
	@goreleaser release --rm-dist

