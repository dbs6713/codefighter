package main

import "fmt"

func main() {
	s1 := "aabcc"
	s2 := "adcaa"
	fmt.Println(commonCharacterCount(s1, s2))
}

func commonCharacterCount(s1 string, s2 string) int {
	b1 := []byte(s1)
	b2 := []byte(s2)
	n := 0
	l := len(b1)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if b1[i] == b2[j] {
				n++
				b2[j] = '%'
				break
			}
		}
	}
	return n
}
