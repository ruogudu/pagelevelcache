package main

import (
	"fmt"
	"github.com/Onmysofa/pagelevelcache/evaluate"
	"math"
	"time"
)

func main() {

	for i := 4; i < 7; i++ {
		for j := 4; j < 7; j++ {
			size := math.Pow10(i)
			num := math.Pow10(j)
			qps := evaluate.EvalCcache(int64(size), int(num), 10, time.Minute * 10, 10)
			fmt.Printf("%v ", qps);
		}
		fmt.Println("")
	}

	//for i := 4; i < 7; i++ {
	//	for j := 4; j < 7; j++ {
	//		size := math.Pow10(i)
	//		num := math.Pow10(j)
	//		qps := evaluate.EvalGcache(int(size), int(num), 10)
	//		fmt.Printf("%v ", qps);
	//	}
	//	fmt.Println("")
	//}

	//evaluate.EvalGcache(1000, 1000000, 10)
}
