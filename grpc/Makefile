build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/grpc/ \
	  myproto/*.proto