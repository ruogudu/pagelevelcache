package evaluate

import (
	"fmt"
	"strconv"
	"time"
)

func insertUtil(insertFunc func (key, val string), num int, thread int, algo string) {
	endChan := make (chan interface{}, 1000)

	start := time.Now()

	stride := (num + thread - 1) / thread

	if thread > 0 {
		for i := 0; i < thread - 1; i++ {
			go insertDaemon(insertFunc, i * stride, (i + 1) * stride, endChan)
		}

		go insertDaemon(insertFunc, (thread - 1) * stride, num, endChan)
	}

	duration := time.Now().Sub(start)
	qps := float64(num)/ duration.Seconds()
	fmt.Printf("Algorithm:%v spent %v to finish %v insertions using %v threads, QPS: %v\n",
		algo, duration, num, thread, qps)
}

func insertDaemon(insertFunc func (key, val string), start int, end int, endChan chan interface {}) {
	for i := start; i < end; i++ {
		insertFunc(strconv.Itoa(i), strconv.Itoa(i))
	}
	endChan <- nil
}