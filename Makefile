bgremover-cli:
	@go run main.go

server:
	@go run ./http/server.go

example-cli:
	@go run main.go -file=./examples/animal-1.jpg -out=./output/animal-1-out.jpg



