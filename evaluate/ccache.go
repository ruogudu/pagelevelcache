package evaluate

import (
	ccache_page "github.com/Onmysofa/ccache"
	"github.com/karlseguin/ccache"
	"time"
)

func EvalCcachePage(size int64, num int, itemsPruning uint32, ttl time.Duration, thread int) float64 {

	var cache = ccache_page.New(ccache_page.Configure().MaxSize(size).ItemsToPrune(itemsPruning).Buckets(128))

	ins := func (key, val string) {
		cache.Set(key, val, ttl)
	}

	return insertUtil(ins, num, thread, "Ccache")
}

func EvalCcache(size int64, num int, itemsPruning uint32, ttl time.Duration, thread int) float64 {

	var cache = ccache.New(ccache.Configure().MaxSize(size).ItemsToPrune(itemsPruning).Buckets(128))

	ins := func (key, val string) {
		cache.Set(key, val, ttl)
	}

	return insertUtil(ins, num, thread, "Ccache")
}
