gen-task:
	protoc proto/task/task.proto --go_out=paths=source_relative:./ --go-grpc_out=paths=source_relative:./ --grpc-gateway_out=paths=source_relative:./