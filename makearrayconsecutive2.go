package main

import (
	"fmt"
)

func main() {
	statues := []int{8, 1, 0, 4, 7}
	fmt.Println(makeArrayConstatuesecutive2(statues))
}

func makeArrayConstatuesecutive2(statues []int) int {
	l := len(statues) - 1
	for i := 0; i < l; i++ {
		if statues[i] > statues[i+1] {
			t := statues[i+1]
			statues[i+1] = statues[i]
			statues[i] = t
		}
	}
	for i := 0; i < l; i++ {
		if statues[i] > statues[i+1] {
			t := statues[i+1]
			statues[i+1] = statues[i]
			statues[i] = t
		}
	}
	final := []int{}
	for i := 0; i < l; i++ {
		d := statues[i+1] - statues[i]
		if d > 1 {
			for x := 1; x < d; x++ {
				final = append(final, statues[i]+x)
			}
		}
	}
	return len(final)
}
