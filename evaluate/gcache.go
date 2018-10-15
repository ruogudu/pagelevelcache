package evaluate

import (
	"github.com/Onmysofa/gcache"
)

func EvalGcache(size int, num int, thread int) {

	gc := gcache.New(size).
		LFU().
		Build()

	ins := func (key, val string) {
		gc.Set(key, val)
	}

	insertUtil(ins, num, thread, "Gcache")
}

