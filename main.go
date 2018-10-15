package main

import (
	"github.com/Onmysofa/pagelevelcache/evaluate"
	"time"
)

func main() {
	evaluate.EvalGcache(1000, 1000000, 10)
	evaluate.EvalCcache(1000, 1000000, 500, time.Minute * 10, 10)
}
