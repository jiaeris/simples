package tcp;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.Socket;

/**
 * Created by Yunga on 2017/12/6.
 */
public class TcpClient {

    public static void main(String[] args) {
        open();
    }

    public static void open() {
        try {
            Socket s = new Socket("localhost", 9001);
            InputStream in = s.getInputStream();
            OutputStream out = s.getOutputStream();

            while (true) {
                Thread.sleep(5000);
                out.write("ping".getBytes());
                byte[] b = new byte[1024];
                int len = in.read(b);
                if (len == -1) break;
                System.out.println(len);
                System.out.println(new String(b));
            }

        } catch (IOException e) {
            e.printStackTrace();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

    }

}
