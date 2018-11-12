package evaluate

import (
	ccache_page "github.com/Onmysofa/ccache"
	"github.com/Onmysofa/pagelevelcache/parse"
	"github.com/karlseguin/ccache"
	"time"
)

func EvalCcachePage(size int64, num int, itemsPruning uint32, ttl time.Duration, thread int) float64 {

	var cache = ccache_page.New(ccache_page.Configure().MaxSize(size).ItemsToPrune(itemsPruning).Buckets(128).Candidates(32))

	ins := func (key, val string) {
		cache.Set(key, val, ttl)
	}

	return insertUtil(ins, num, thread, "Ccache")
}

func EvalCcacheTrace(chs []chan *parse.PageReq, size int64, num int, itemsPruning uint32, ttl time.Duration, thread int) float64 {

	var cache = ccache_page.New(ccache_page.Configure().MaxSize(size).ItemsToPrune(itemsPruning).Buckets(128).Candidates(32))

	ins := func (req *parse.PageReq) {

		cReqs := make([]*ccache_page.Request, len(req.Objs), len(req.Objs))

		for i, o := range req.Objs {
			cReqs[i] = &ccache_page.Request{o.Backend, o.Uri, o.Obj}
		}

		cache.GetPage(cReqs)

		objsToSet := make([]*ccache_page.Request, 0, len(req.Objs))
		for _, o := range cReqs {
			if parse.IsNilObject(o.Obj) {
				objsToSet = append(objsToSet, o)
			}
		}

		cache.SetPage(objsToSet, ttl)
	}

	return insertUtilTrace(chs, ins, num, thread, "CcacheTrace")
}

func EvalCcacheRatio(ch chan *parse.PageReq, granularity int, size int64, itemsPruning uint32, ttl time.Duration) float64 {

	var cache = ccache_page.New(ccache_page.Configure().MaxSize(size).ItemsToPrune(itemsPruning).Buckets(128).Candidates(32))

	ins := func (req *parse.PageReq) (all int, hit int) {

		cReqs := make([]*ccache_page.Request, len(req.Objs), len(req.Objs))

		for i, o := range req.Objs {
			cReqs[i] = &ccache_page.Request{o.Backend, o.Uri, o.Obj}
		}

		cache.GetPage(cReqs)

		objsToSet := make([]*ccache_page.Request, 0, len(req.Objs))
		for i, o := range cReqs {
			if parse.IsNilObject(o.Obj) {
				o.Obj = parse.NewObject(req.Objs[i].Size)
				objsToSet = append(objsToSet, o)
			}
		}

		cache.SetPage(objsToSet, ttl)

		//return len(cReqs), len(cReqs)-len(objsToSet)
		if len(objsToSet) == 0 {
			return 1, 1
		}
		return 1, 0
	}

	return hitRatioUtilTrace(ch, granularity, ins,"CcacheTrace")
}

func EvalCcache(size int64, num int, itemsPruning uint32, ttl time.Duration, thread int) float64 {

	var cache = ccache.New(ccache.Configure().MaxSize(size).ItemsToPrune(itemsPruning).Buckets(32))

	ins := func (key, val string) {
		cache.Set(key, val, ttl)
	}

	return insertUtil(ins, num, thread, "Ccache")
}
