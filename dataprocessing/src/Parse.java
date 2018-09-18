import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.stream.JsonReader;
import model.PageReq;

import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.UnsupportedEncodingException;

public class Parse {
    public static void readStream(String filepath) {
        try {
            FileInputStream in = new FileInputStream(filepath);
            JsonReader reader = new JsonReader(new InputStreamReader(in, "ASCII"));
            Gson gson = new GsonBuilder().create();

            // Read file in stream mode
            reader.beginArray();
            while (reader.hasNext()) {
                // Read data into object model
                PageReq pr = gson.fromJson(reader, PageReq.class);

                // TODO
            }
            reader.close();
        } catch (UnsupportedEncodingException ex) {
            ex.printStackTrace();
        } catch (IOException ex) {
            ex.printStackTrace();
        }
    }
}
