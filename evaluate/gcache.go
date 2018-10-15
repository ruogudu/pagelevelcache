package evaluate

import (
	"github.com/Onmysofa/gcache"
)

func EvalGcache(size int, num int, thread int) float64 {

	gc := gcache.New(size).
		LFU().
		Build()

	ins := func (key, val string) {
		gc.Set(key, val)
	}

	return insertUtil(ins, num, thread, "Gcache")
}

