# gpb
Google Protocol Buffers client-server application

### Running protoc for Java Server:
```shell
protoc -I=proto \
       --plugin=protoc-gen-grpc-java=Users/danorel/Workspace/Education/University/KMA/Labs/Distributed\ Systems/ClientServer/gpb/plugins/protoc-gen-grpc-java-1.51.0-osx-x86_64.exe \
       --java_out=server/src/main/java \
       --java-grpc_out=server/src/main/java \
       Chat.proto
```

### Running protoc for Golang Client:
```shell
protoc -I=proto \
       --go_out=client \
       --go-grpc_out=client \
       Chat.proto
```