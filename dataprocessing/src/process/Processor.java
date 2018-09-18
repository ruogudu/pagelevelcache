package process;

import model.PageReq;

public interface Processor {
    void take(PageReq pr);
    void store(String filename);
}
