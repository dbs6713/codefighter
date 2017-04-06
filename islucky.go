package main

import (
	"fmt"
)

func main() {
	fmt.Println(isLucky(1234))
	fmt.Println(isLucky(1230))
	fmt.Println(isLucky(440017))
}

func isLucky(n int) bool {
	if n < 0 {
		return false
	}
	s1 := 0
	s2 := 0
	l := nlen(n) // uint64(math.Log10(float64(n)) + 1)
	h := l / 2
	for i := 0; i < h; i++ {
		s1 += n % 10
		n /= 10
	}
	for ; n != 0; n /= 10 {
		s2 += n % 10
	}
	return s1 == s2
}

func nlen(n int) int {
	switch {
	case n < 10:
		return 1
	case n < 100:
		return 2
	case n < 1000:
		return 3
	case n < 10000:
		return 4
	case n < 100000:
		return 5
	case n < 1000000:
		return 6
	case n < 10000000:
		return 7
	case n < 100000000:
		return 8
	case n < 1000000000:
		return 9
	}
	return 10
}
