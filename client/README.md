### How to run

1. Install all required dependencies
```shell
go get -u -v -f all
```

2. Make commands:
```shell
make build # build the code
make test # test the code
make vet # check the vetting
make lint # check the linting
make fmt # check the formatting
make # ensure everything passes and builds
```

Making a request to a grpc server looks like the following:
```shell
grpcurl --plaintext -d '{"id": 1,"title": "Hello from client!","text": "Boo!"}' localhost:9898 com.danorel.chat.Chat/SendMessage
```