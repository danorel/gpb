import java.io.FileOutputStream;
import java.io.IOException;

public class Server {
    public static void main(String[] args) {
        if (args.length != 1) {
            System.err.println("Usage: output file is absent in command line call");
            System.exit(-1);
        }
        UserMessage message = UserMessage.newBuilder()
                .setId(1)
                .setTitle("Greetings from a friend")
                .setText("Hello, it's me! Glad to see you!")
                .build();
        FileOutputStream output;
        try {
            output = new FileOutputStream(args[0]);
            message.writeTo(output);
            output.close();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}