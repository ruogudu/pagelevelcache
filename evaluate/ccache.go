package evaluate

import (
	"github.com/Onmysofa/ccache"
	"time"
)

func EvalCcache(size int64, num int, itemsPruning uint32, ttl time.Duration, thread int) {

	var cache = ccache.New(ccache.Configure().MaxSize(size).ItemsToPrune(itemsPruning))

	ins := func (key, val string) {
		cache.Set(key, val, ttl)
	}

	insertUtil(ins, num, thread, "Ccache")
}
