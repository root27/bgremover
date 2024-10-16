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


python-protogen-build:
	@docker build -t python-proto -f ./Dockerfiles/Dockerfile.pythonprotogen .
	@docker create --name grpc-container python-proto
	@docker cp grpc-container:/output/generated ./python 
	@echo python proto generated

##Gorelease make commands

release:
	@goreleaser release --rm-dist

