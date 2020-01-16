package main

import (
	"fmt"
	"sort"
	"time"
)

const (
	size  = 60_000_000
	nLoop = 21
)

func main() {
	perf(false)
	perf(true)
}

func perf(toSort bool) {
	var arr [size]int

	flip := false
	for i := 0; i < size; i++ {
		arr[i] = i + 1
		if flip {
			arr[i] = -arr[i]
		}
		flip = !flip
	}

	if toSort {
		sort.Ints(arr[:size])
	}

	start := time.Now()

	sum := 0
	for ir := 0; ir < nLoop; ir++ {
		for i := 0; i < size; i++ {
			if arr[i] > 100 {
				sum += arr[i]
			}
		}
	}

	dur := time.Since(start)

	fmt.Printf("toSort=%-5v duration=%-15v         array size=%v iterations=%v sum=%v\n", toSort, dur, size, size*nLoop, sum)
}
