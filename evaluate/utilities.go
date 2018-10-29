package evaluate

import (
	"github.com/Onmysofa/pagelevelcache/parse"
	"strconv"
	"time"
)

func insertUtil(insertFunc func (key, val string), num int, thread int, algo string) float64 {
	endChan := make (chan interface{}, thread + 1)

	start := time.Now()

	stride := (num + thread - 1) / thread

	if thread > 0 {
		for i := 0; i < thread - 1; i++ {
			go insertDaemon(insertFunc, i * stride, (i + 1) * stride, endChan)
		}

		go insertDaemon(insertFunc, (thread - 1) * stride, num, endChan)
	}

	for i := 0; i < thread; i++ {
		<- endChan
	}

	duration := time.Now().Sub(start)
	qps := float64(num)/ duration.Seconds()
	//fmt.Printf("Algorithm:%v spent %v to finish %v insertions using %v threads. Throughput: %v\n",
	//	algo, duration, num, thread, qps)
	return qps
}

func insertUtilTrace(ch chan *parse.PageReq, insertFunc func (req *parse.PageReq), num int, thread int, algo string) float64 {
	endChan := make (chan interface{}, thread + 1)

	start := time.Now()

	for i := 0; i < thread; i++ {
		go insertDaemonTrace(insertFunc, ch, endChan)
	}

	for i := 0; i < thread; i++ {
		<- endChan
	}

	duration := time.Now().Sub(start)
	qps := float64(num)/ duration.Seconds()
	//fmt.Printf("Algorithm:%v spent %v to finish %v insertions using %v threads. Throughput: %v\n",
	//	algo, duration, num, thread, qps)
	return qps
}

func insertDaemon(insertFunc func (key, val string), start int, end int, endChan chan interface {}) {
	for i := start; i < end; i++ {
		insertFunc(strconv.Itoa(i), strconv.Itoa(i))
	}
	endChan <- nil
}

func insertDaemonTrace(insertFunc func (req *parse.PageReq), ch chan *parse.PageReq, endChan chan interface {}) {
	for req, ok := <- ch; ok; req, ok = <- ch {
		insertFunc(req)
	}
	endChan <- nil
}