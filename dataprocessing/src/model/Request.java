package model;

import com.google.gson.Gson;

public class Request {
    public String uri;
    public String backend;
    public int size;

    @Override
    public String toString() {
        return new Gson().toJson(this);
    }
}
