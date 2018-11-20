package evaluate

import (
	"fmt"
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

func insertUtilTrace(chs []chan *parse.PageReq, insertFunc func (req *parse.PageReq), num int, thread int, algo string) float64 {
	endChan := make (chan interface{}, thread + 1)

	start := time.Now()

	for i := 0; i < thread; i++ {
		go insertDaemonTrace(insertFunc, chs[i], endChan)
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

func hitRatioUtilTrace(ch chan *parse.PageReq, granularity int, reportThresold int, insertFunc func (req *parse.PageReq) (int, int), algo string) float64 {

	all := 0
	hit := 0
	cnt := 0
	next := reportThresold + granularity
	for req, ok := <- ch; ok; req, ok = <- ch {
		curAll, curHit := insertFunc(req)

		cnt++

		if cnt > reportThresold {
			all += curAll
			hit += curHit

			if cnt >= next {

				//fmt.Println("Report:", "All", all, "Hit", hit, "Ratio", float64(hit) / float64(all))
				fmt.Printf("%v\t%v\t%v\t%v\n", cnt, all, hit, float64(hit) / float64(all))

				next += granularity
			}
		}
	}

	return float64(hit) / float64(all)
}