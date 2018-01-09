package http;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.net.HttpURLConnection;
import java.net.URL;

/**
 * Created by Yunga on 2017/12/6.
 */
public class HttpClient {
    public static final String host = "http://www.baidu.com";

    public static void get() {
        try {
            URL url = new URL(host);
            HttpURLConnection httpURLConnection = (HttpURLConnection) url.openConnection();
            httpURLConnection.setRequestMethod("GET");
            InputStream is = httpURLConnection.getInputStream();
            byte[] b = new byte[1024];
            ByteArrayOutputStream baos = new ByteArrayOutputStream();
            int len;
            while ((len = is.read(b)) != -1) {
                baos.write(b, 0, len);
            }
            System.out.println(baos.toString("utf-8"));
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

}
