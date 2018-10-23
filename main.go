package main

import (
	"fmt"
	"github.com/Onmysofa/pagelevelcache/evaluate"
	"math"
	"time"
)

func main() {

	for i := 4; i < 7; i++ {
		for j := 4; j <= 7; j++ {
			size := math.Pow10(i)
			num := math.Pow10(j)
			qps := evaluate.EvalCcache(int64(size), int(num), 10, time.Minute * 10, 10)
			fmt.Printf("%v ", qps);
		}
		fmt.Println("")
	}

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

	//ch, err := parse.ParseFile("/home/ruogu/Desktop/capstone/data/first1000.json")
	//if err != nil {
	//	return
	//}
	//
	//for req := range ch {
	//	fmt.Println(req)
	//}
}
