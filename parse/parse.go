package parse

import (
	"bufio"
	"encoding/json"
	"os"
)

type ObjReq struct {
	Backend string
	Size int
	Uri string
}

type PageReq struct {
	Objs []ObjReq
}

func parseFile(path string) (chan *PageReq, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	ch := make (chan *PageReq, 1000)
	go parseDaemon(f, ch)

	return ch, nil
}

func parseDaemon (f *os.File, ch chan *PageReq) {
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		req := PageReq{}
		err := json.Unmarshal([]byte(line), &req.Objs)
		if err != nil {
			ch <- &req
		}
	}

	close(ch)
}
