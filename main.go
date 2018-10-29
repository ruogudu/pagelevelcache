package main

import (
	"fmt"
	"github.com/Onmysofa/pagelevelcache/evaluate"
	"github.com/Onmysofa/pagelevelcache/parse"
	"math"
	"time"
)

func main() {

	//for i := 4; i < 7; i++ {
	//	for j := 4; j <= 7; j++ {
	//		size := math.Pow10(i)
	//		num := math.Pow10(j)
	//		qps := evaluate.EvalCcachePage(int64(size), int(num), 100, time.Minute * 10, 10)
	//		fmt.Printf("%v ", qps);
	//	}
	//	fmt.Println("")
	//}

	//for t := 1; t < 6; t++ {
	//	qps :=evaluate.EvalCcache(1000000, 10000000, 10, time.Minute * 10, t)
	//	fmt.Printf("%v ", qps);
	//}


	//for i := 4; i < 7; i++ {
	//	for j := 4; j <= 7; j++ {
	//		size := math.Pow10(i)
	//		num := math.Pow10(j)
	//		qps := evaluate.EvalGcache(int(size), int(num), 10)
	//		fmt.Printf("%v ", qps);
	//	}
	//	fmt.Println("")
	//}

	//for t := 1; t < 6; t++ {
	//	qps := evaluate.EvalGcache(1000000, 10000000, 10)
	//	fmt.Printf("%v ", qps);
	//}

	//evaluate.EvalGcache(1000, 1000000, 10)

	ch, err := parse.ParseFile("/home/ruogu/Desktop/capstone/data/first100000.json")
	if err != nil {
		return
	}
	for i := 4; i < 7; i++ {
		size := math.Pow10(i)
		qps := evaluate.EvalCcacheTrace(ch, int64(size), 100000, 100, time.Minute * 10, 10)
		fmt.Printf("%v ", qps);
		fmt.Println("")
	}

	//parse.ParititionFile("/home/ruogu/Desktop/capstone/data/trace_2018_03_06_24h.json", 8)


}
