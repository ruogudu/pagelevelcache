package parse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type ObjReq struct {
	Backend uint64
	Size int
	Uri uint64
	Obj interface{}
}

type PageReq struct {
	Objs []ObjReq
}

func ParseFile(path string, num int) ([] chan*PageReq, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	chs := make([]chan *PageReq, num, num)
	for i := range chs {
		chs[i] = make(chan *PageReq, 1000000)
	}
	go parseDaemon(f, chs)

	return chs, nil
}

func parseDaemon (f *os.File, chs []chan *PageReq) {
	defer f.Close()

	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	cnt := 0
	num := len(chs)
	for scanner.Scan() {
		line := scanner.Text()
		req := PageReq{}
		err := json.Unmarshal([]byte(line), &req.Objs)
		if err == nil {
			for i, o := range req.Objs {
				req.Objs[i].Obj = NewObject(o.Size)
			}
			chs[cnt % num] <- &req
			cnt++

			//if cnt % 10000 == 0 {
			//	fmt.Println("Current:", cnt)
			//}

		} else {
			fmt.Println(err)
		}
	}
	fmt.Println("Parse finished, num: ", cnt)
	for _, ch := range chs {
		close(ch)
	}
}

func ParseFileWithoutValue(path string, num int) ([] chan*PageReq, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	chs := make([]chan *PageReq, num, num)
	for i := range chs {
		chs[i] = make(chan *PageReq, 1000000)
	}
	go parseDaemonWithoutValue(f, chs)

	return chs, nil
}

func parseDaemonWithoutValue(f *os.File, chs []chan *PageReq) {
	defer f.Close()

	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	cnt := 0
	num := len(chs)
	for scanner.Scan() {
		line := scanner.Text()
		req := PageReq{}
		err := json.Unmarshal([]byte(line), &req.Objs)
		if err == nil {
			for i := range req.Objs {
				req.Objs[i].Obj = NewNilOBject()
			}
			chs[cnt % num] <- &req
			cnt++

		} else {
			fmt.Println(err)
		}
	}
	fmt.Println("Parse finished, num: ", cnt)
	for _, ch := range chs {
		close(ch)
	}
}