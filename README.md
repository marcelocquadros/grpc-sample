# grpc-sample

- This project shows how to implement a service and the client using grpc and golang.

### How to run the project

> From the root folder, run the commands below to start the server and the client.

- Running the server
```cmd
go run server/server.go
```
- Running the client
```cmd
go run client/client.go
```
> Case you want to recompile the .proto file

```cmd
protoc --proto_path=proto proto/service.proto --go_out=plugins=grpc:pb
```