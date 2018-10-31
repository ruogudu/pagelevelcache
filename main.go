package main

import (
	"fmt"
	"github.com/Onmysofa/pagelevelcache/evaluate"
	"github.com/Onmysofa/pagelevelcache/parse"
	"os"
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

	//ch, err := parse.ParseFile("/home/ruogu/Desktop/capstone/data/first1000.json")
	//if err != nil {
	//	return
	//}
	//for i := 7; i < 10; i++ {
	//	size := math.Pow10(i)
	//	qps := evaluate.EvalCcacheTrace(ch, int64(size), 1000, 100, time.Minute * 10, 10)
	//	fmt.Printf("%v ", qps);
	//	fmt.Println("")
	//}

	//parse.ParititionFile("/home/ruogu/Desktop/capstone/data/trace_2018_03_06_24h.json", 8)

	argsWithoutProg := os.Args[1:]

	funBenchTrace(argsWithoutProg[0])

}

func funCalcSize(filename string) {
	ch, err := parse.ParseFile(filename)
	if err != nil {
		return
	}

	res := calcSizeSum(ch)
	fmt.Print("Size sum:")
	fmt.Println(res)
}

func funCalcNum(filename string) {
	ch, err := parse.ParseFile(filename)
	if err != nil {
		return
	}

	res := calcNum(ch)
	fmt.Print("Number:")
	fmt.Println(res)
}

func funBenchTrace(filename string) {
	ch, err := parse.ParseFile(filename)
	if err != nil {
		return
	}

	size := 50000000000
	qps := evaluate.EvalCcacheTrace(ch, int64(size), 100000, 100, time.Minute * 10, 8)
	fmt.Printf("%v ", qps);
	fmt.Println("")

}

func calcSizeSum(ch chan *parse.PageReq) int64 {
	var sum int64 = 0
	for r := range ch {
		for _, o := range r.Objs {
			sum += int64(o.Size)
		}
	}

	return sum
}

func calcNum(ch chan *parse.PageReq) int {
	sum := 0
	for range ch {
		sum++
	}

	return sum
}

