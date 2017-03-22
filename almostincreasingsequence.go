package main

import (
	"fmt"
	"math"
	"time"
)

const (
	bitsPerWord = 32 << uint(^uint(0)>>63)
	BitsPerWord = bitsPerWord            // either 32 or 64
	MaxInt      = 1<<(BitsPerWord-1) - 1 // either 1<<31 - 1 or 1<<63 - 1
	MinInt      = -MaxInt - 1            // either -1 << 31 or -1 << 63
)

func main() {
	seq0 := []int{1, 3, 2, 1}  // false
	seq1 := []int{1, 3, 2}     // true
	seq2 := []int{1, 2, 3, 1}  // true
	seq3 := []int{0, -2, 5, 6} // true
	fmt.Println(almostIncreasingSequence(seq0))
	fmt.Println(almostIncreasingSequence(seq1))
	fmt.Println(almostIncreasingSequence(seq2))
	fmt.Println(almostIncreasingSequence(seq3))
}

func almostIncreasingSequence(sequence []int) bool {
	start := time.Now()
	if preConditionsMet(sequence) == false {
		return false
	}
	mV := MaxInt
	pmV := MinInt
	n := 0
	for i := 0; i < len(sequence); i++ {
		//for i, _ := range sequence {
		if sequence[i] <= pmV || sequence[i] == mV {
			n = n + 1
		}
		if i-1 > -1 && (sequence[i] <= sequence[i-1]) {
			mV = sequence[i-1]
			if i-2 > -1 {
				pmV = sequence[i-2]
			}
			n = n + 1
		}
	}
	duration := time.Since(start)
	fmt.Println("total:", duration)
	if n == 1 || n == 0 {
		return true
	}
	return false
}

func preConditionsMet(s []int) bool {
	constraint := math.Pow(10, 5)
	n := len(s)
	if n < 2 || float64(n) > constraint {
		return false
	}
	for i := 0; i < n; i++ {
		if float64(s[i]) < constraint*-1 || float64(s[i]) > constraint {
			return false
		}
	}
	return true
}
