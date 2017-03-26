package main

const (
	bitsPerWord = 32 << uint(^uint(0)>>63)
	BitsPerWord = bitsPerWord            // either 32 or 64
	MaxInt      = 1<<(BitsPerWord-1) - 1 // either 1<<31 - 1 or 1<<63 - 1
	MinInt      = -MaxInt - 1            // either -1 << 31 or -1 << 63
)

func almostIncreasingSequence(sequence []int) bool {
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
	if n == 1 || n == 0 {
		return true
	}
	return false
}
