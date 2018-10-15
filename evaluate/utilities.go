package evaluate

import (
	"strconv"
)

func insertUtil(insertFunc func (key, val string), num int, thread int) {
	endChan := make (chan interface{}, 1000)

	
}

func insertDaemon(insertFunc func (key, val string), start int, end int, endChan chan interface {}) {
	for i := start; i < end; i++ {
		insertFunc(strconv.Itoa(i), strconv.Itoa(i))
	}
	endChan <- nil
}