package parse;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.stream.JsonReader;
import model.PageReq;
import model.Request;

import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.UnsupportedEncodingException;
import java.util.ArrayList;
import java.util.List;

public class Parse {
    public static void readStream(String filepath) {
        try {
            FileInputStream in = new FileInputStream(filepath);
            JsonReader reader = new JsonReader(new InputStreamReader(in, "ASCII"));
            reader.setLenient(true);
            Gson gson = new GsonBuilder().create();

            while (true) {
                // Read file in stream mode
                reader.beginArray();
                List<Request> reqs = new ArrayList<>();
                while (reader.hasNext()) {
                    // Read data into object model
                    Request r = gson.fromJson(reader, Request.class);
                    reqs.add(r);

                }
                reader.endArray();

                PageReq pr = new PageReq(reqs);
                System.out.println(pr.toString());
            }

        } catch (UnsupportedEncodingException ex) {
            ex.printStackTrace();
        } catch (IOException ex) {
            ex.printStackTrace();
        }
    }
}
