package evaluate

import (
	ccache_page "github.com/Onmysofa/ccache"
	"github.com/Onmysofa/pagelevelcache/parse"
	"github.com/karlseguin/ccache"
	"time"
)

func EvalCcachePage(size int64, num int, itemsPruning uint32, ttl time.Duration, thread int) float64 {

	var cache = ccache_page.New(ccache_page.Configure().MaxSize(size).ItemsToPrune(itemsPruning).Buckets(64).Candidates(32))

	ins := func (key, val string) {
		cache.Set(key, val, ttl)
	}

	return insertUtil(ins, num, thread, "Ccache")
}

func EvalCcacheTrace(chs []chan *parse.PageReq, size int64, num int, itemsPruning uint32, ttl time.Duration, thread int) float64 {

	var cache = ccache_page.New(ccache_page.Configure().MaxSize(size).ItemsToPrune(itemsPruning).Buckets(64).Candidates(32))

	ins := func (req *parse.PageReq) {

		cReqs := make([]*ccache_page.Request, len(req.Objs), len(req.Objs))

		for i, o := range req.Objs {
			cReqs[i] = &ccache_page.Request{o.Backend, o.Uri, o.Obj}
		}

		res := cache.GetPage(cReqs)

		if res == nil {
			cache.SetPage(cReqs, ttl)
		}
	}

	return insertUtilTrace(chs, ins, num, thread, "CcacheTrace")
}

func EvalCcache(size int64, num int, itemsPruning uint32, ttl time.Duration, thread int) float64 {

	var cache = ccache.New(ccache.Configure().MaxSize(size).ItemsToPrune(itemsPruning).Buckets(32))

	ins := func (key, val string) {
		cache.Set(key, val, ttl)
	}

	return insertUtil(ins, num, thread, "Ccache")
}
