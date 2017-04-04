package main

import "log"

func main() {
	a0 := []string{"aba", "aa", "ad", "vcd", "aba"}
	log.Println(allLongestStrings(a0))

	a1 := []string{"abc", "eeee", "abcd", "dcd"}
	log.Println(allLongestStrings(a1))
}

func allLongestStrings(inputArray []string) []string {
	l := 0
	r := []string{}
	for _, v := range inputArray {
		c := len(v)
		if c > l {
			l = c
		}
	}
	for _, v := range inputArray {
		if len(v) == l {
			r = append(r, v)
		}
	}
	return r
}
