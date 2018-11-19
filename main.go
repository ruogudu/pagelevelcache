package main

import (
	"flag"
	"fmt"
	"github.com/Onmysofa/pagelevelcache/evaluate"
	"github.com/Onmysofa/pagelevelcache/parse"
	"time"
)

func main() {

	tracePtr := flag.String("t", "", "The trace to test against")
	sizePtr := flag.Int64("s", 5000000000, "The size of cache by byte")
	granularityPtr := flag.Int("g", 30000, "The granularity to report")
	reportThresPtr := flag.Int("r", 300000, "Threshold when start to report")
	algoPtr := flag.String("a", "LFU", "Algorithm to test")
	bucketPtr := flag.Int("b", 128, "Bucket number")
	samplePtr := flag.Int("m", 32, "Sample number")
	threadPtr := flag.Int("n", 1, "Number of threads")
	itemsPtr := flag.Int("i", 10, "Number of items evicted each time")

	ohrPtr := flag.Bool("o", false, "Calculate OHR")
	phrPtr := flag.Bool("p", false, "Calculate PHR")
	uPtr := flag.Bool("u", false, "Calculate working set")
	cPtr := flag.Bool("c", false, "Count request number")
	evalPtr := flag.Bool("e", false, "Evaluate throughput")

	flag.Parse()

	if *ohrPtr {
		funBenchTraceOHR(*tracePtr, *granularityPtr, *reportThresPtr, *algoPtr, *sizePtr, *bucketPtr, *samplePtr, *itemsPtr)
	}

	if *phrPtr {
		funBenchTracePHR(*tracePtr, *granularityPtr, *reportThresPtr, *algoPtr, *sizePtr, *bucketPtr, *samplePtr, *itemsPtr)
	}
	if *uPtr {
		funCalcUniqueSize(*tracePtr)
	}
	if *cPtr {
		funCalcNum(*tracePtr)
	}
	if *evalPtr {
		funBenchTraceThroughtput(*tracePtr, *algoPtr, *sizePtr, *threadPtr, *itemsPtr)
	}
}

func funCalcSize(filename string) {
	chs, err := parse.ParseFile(filename, 1)
	if err != nil {
		return
	}

	res := calcSizeSum(chs[0])
	fmt.Print("Size sum: ", res)
}

func funCalcUniqueSize(filename string) {
	chs, err := parse.ParseFile(filename, 1)
	if err != nil {
		return
	}

	res := calcUniqueSize(chs[0])
	fmt.Print("Unique size sum: ", res)
}

func funCalcNum(filename string) {
	chs, err := parse.ParseFile(filename, 1)
	if err != nil {
		return
	}

	start := time.Now()

	res := calcNum(chs[0])

	duration := time.Now().Sub(start)
	qps := float64(res)/ duration.Seconds()

	fmt.Println("Number: ", res)
	fmt.Println("QPS: ", qps)
}

func funBenchTraceThroughtput(filename string, algorithm string, size int64, threads int, pruningItems int) {
	chs, err := parse.ParseFileWithoutValue(filename, 1)
	if err != nil {
		return
	}

	fmt.Println("Size: ", size, " Threads: ", threads)
	num := calcNum(chs[0])
	fmt.Println("Num: ", num)

	chs, err = parse.ParseFileWithoutValue(filename, threads)
	if err != nil {
		return
	}

	fmt.Print("Wait 60s for parsing...")
	time.Sleep(60 * time.Second)
	fmt.Println("")

	qps := evaluate.EvalCcacheTrace(chs, algorithm, size, num, uint32(pruningItems), time.Minute * 10, threads)
	fmt.Printf("%v ", qps);
	fmt.Println("")

}

func funBenchTracePHR(filename string, granularity int, reportThreshold int, algorithm string, size int64, buckets int, samplenum int, pruningItems int) {

	fmt.Println("Trace:", filename)
	fmt.Println("Algorithm:", algorithm)
	fmt.Println("Granularity:", granularity)
	fmt.Println("Report threshold:", reportThreshold)
	fmt.Println("Cache size:", size)

	chs, err := parse.ParseFileWithoutValue(filename, 1)
	if err != nil {
		return
	}

	ratio := evaluate.EvalCcachePHR(chs[0], granularity, reportThreshold, algorithm, size, buckets, samplenum, uint32(pruningItems), time.Minute * 10)
	fmt.Printf("Ratio: %v\n ", ratio);
}

func funBenchTraceOHR(filename string, granularity int, reportThreshold int, algorithm string, size int64, buckets int, samplenum int, pruningItems int) {

	fmt.Println("Trace:", filename)
	fmt.Println("Algorithm:", algorithm)
	fmt.Println("Granularity:", granularity)
	fmt.Println("Report threshold:", reportThreshold)
	fmt.Println("Cache size:", size)

	chs, err := parse.ParseFileWithoutValue(filename, 1)
	if err != nil {
		return
	}

	ratio := evaluate.EvalCcacheOHR(chs[0], granularity, reportThreshold, algorithm, size, buckets, samplenum, uint32(pruningItems), time.Minute * 10)
	fmt.Printf("Ratio: %v\n ", ratio);
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

func calcUniqueSize(ch chan *parse.PageReq) int64 {
	var sum int64 = 0
	m := make(map[string]bool)

	for r := range ch {
		for _, o := range r.Objs {
			k := fmt.Sprintf("%v:%v", o.Backend, o.Uri)
			_,ok := m[k]
			if !ok {
				sum += int64(o.Size)
				m[k] = true
			}
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

