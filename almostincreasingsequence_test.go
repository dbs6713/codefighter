package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlmostIncreasingSequence0(t *testing.T) {
	s := []int{1, 3, 2, 1}
	assert.False(t, almostIncreasingSequence(s))
}

func TestAlmostIncreasingSequence1(t *testing.T) {
	s := []int{1, 1, 1}
	assert.False(t, almostIncreasingSequence(s))
}

func TestAlmostIncreasingSequence2(t *testing.T) {
	s := []int{1, 3, 2}
	assert.True(t, almostIncreasingSequence(s))
}

func TestAlmostIncreasingSequence3(t *testing.T) {
	s := []int{1, 2, 3, 1}
	assert.True(t, almostIncreasingSequence(s))
}

func TestAlmostIncreasingSequence4(t *testing.T) {
	s := []int{0, -2, 5, 6}
	assert.True(t, almostIncreasingSequence(s))
}

func TestAlmostIncreasingSequence5(t *testing.T) {
	s := []int{1, 2, 3, 4, 99, 5, 6}
	assert.True(t, almostIncreasingSequence(s))
}
