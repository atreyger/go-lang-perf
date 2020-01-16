package main

import (
	"fmt"
	"sort"
	"time"
)

const size = 60_000_000
const nLoop = 21

var arr [size]int

func main() {
	sl := make([]int, size)
	process("local  mixed", mixed(sl))
	process("local  1x", multi(mixed(sl), 1))
	process("local  2x", multi(mixed(sl), 2))
	process("local  3x", multi(mixed(sl), 3))
	process("local  ordered", ordered(sl))
	process("local  sorted", sorted(mixed(sl)))

	sl = arr[:]
	process("global mixed", mixed(sl))
	process("global 1x", multi(mixed(sl), 1))
	process("global 2x", multi(mixed(sl), 2))
	process("global 3x", multi(mixed(sl), 3))
	process("global ordered", ordered(sl))
	process("global sorted", sorted(mixed(sl)))
}

func mixed(sl []int) []int {
	for i := 0; i < size; i = i + 2 {
		sl[i] = i + 1
	}
	for i := 1; i < size; i = i + 2 {
		sl[i] = -(i + 1)
	}
	return sl
}

func multi(sl []int, k int) []int {
	for i := 0; i < size; i = i + 3 {
		sl[i] = sl[i] % k
	}
	return sl
}

func sorted(sl []int) []int {
	sort.Ints(sl)
	return sl
}

func ordered(sl []int) []int {
	for i := 0; i < size; i++ {
		sl[i] = i + 1
	}
	return sl
}

func process(name string, sl []int) {
	start := time.Now()

	sum := 0
	for ir := 0; ir < nLoop; ir++ {
		for _, v := range sl {
			if v > 100 {
				sum += v
			}
		}
	}

	dur := time.Since(start)

	fmt.Printf("%-14s duration=%-15v            sum=%v\n", name, dur, sum)
}
