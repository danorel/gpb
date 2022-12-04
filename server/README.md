### How to run

Two options:
1. Run grpc server using Spring Boot directly in the code:
   - `mvn spring-boot:run`
2. Build grpc server using Spring Boot and run jar file:
   - `mvn compile`
   - `mvn package`
   - `java -jar target/server-1.0.jar com.danorel.ServerApplication`

Making a request to a grpc server looks like the following:
```shell
grpcurl --plaintext -d '{"id": 1,"title": "Hello from client!","text": "Boo!"}' localhost:9898 com.danorel.chat.Chat/SendMessage
```