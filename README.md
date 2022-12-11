# gpb
Google Protocol Buffers client-server application

### Running protoc for Java Server:
```shell
protoc -I=proto \
       --java_out=server/src/main/java \
       Chat.proto
```

### Running protoc for Golang Client:
```shell
protoc -I=proto \
       --go_out=client \
       --go-grpc_out=client \
       Chat.proto
```