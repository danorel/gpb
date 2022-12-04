package com.danorel.chat;

import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;

@GrpcService
public class GrpcChatService extends ChatGrpc.ChatImplBase {
    @Override
    public void sendMessage(ChatProto.Message message, StreamObserver<ChatProto.Reply> responseObserver) {
        ChatProto.Reply reply = ChatProto.Reply.newBuilder()
                .setId(message.getId())
                .setTitle("Greetings from a server")
                .setEmoji("\uD83D\uDE02")
                .build();
        responseObserver.onNext(reply);
        responseObserver.onCompleted();
    }
}
