package model;

public class PageReq {
    public Request[] requests;

    @Override
    public String toString() {
        if (requests.length == 0) {
            return "[]";
        }
        StringBuilder sb = new StringBuilder();
        sb.append("[");
        sb.append(requests[0].toString());
        for (int i = 1; i < requests.length; i++) {
            sb.append(",");
            sb.append(requests[i].toString());
        }
        return sb.toString();
    }
}
