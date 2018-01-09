package tcp;

import java.io.*;
import java.net.InetAddress;
import java.net.ServerSocket;
import java.net.Socket;

/**
 * Created by Yunga on 2017/12/6.
 */
public class TcpServer {

    public static void main(String[] args) {
        open();
    }

    public static final int TCP_PORT = 9001;

    public static void open() {
        try {
            ServerSocket ss = new ServerSocket(TCP_PORT);
            InetAddress ia = InetAddress.getByName(null);
            System.out.println("TcpServer@" + ia + " start!");

            while (true) {
                Socket socket = ss.accept();
                new TcpClientThread(socket);
            }

        } catch (IOException e) {
            e.printStackTrace();
        }

    }

    private static class TcpClientThread extends Thread {

        private InputStream in;
        private OutputStream out;
        private Socket socket;
        private BufferedReader br;

        public TcpClientThread(Socket s) throws IOException {
            socket = s;
            in = s.getInputStream();
            out = s.getOutputStream();
            start();
        }


        @Override
        public void run() {
            try {
                while (true) {

                    byte[] b = new byte[1024];
                    int len = in.read(b);
                    if (len == -1) break;
                    System.out.println(len);
                    System.out.println(new String(b));

                    out.write(b);
                    out.flush();
                }
            } catch (Exception e) {
                e.printStackTrace();
//                System.out.println(socket.getRemoteSocketAddress() + " 异常");
            } finally {
                try {
                    in.close();
                    out.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }

        }
    }
}
