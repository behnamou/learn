package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func Test1(t *testing.T) {
// 	cases := []struct {
// 		value    []int
// 		expected float64
// 	}{
// 		{
// 			value:    []int{1, 2, 3, 4, 5},
// 			expected: 3,
// 		},
// 		{
// 			value:    []int{0, 0, 0, 0, 0},
// 			expected: 0,
// 		},
// 		{
// 			value:    []int{},
// 			expected: 0,
// 		},
// 		{
// 			value:    []int{5, 5, 5, 5, 5},
// 			expected: 5,
// 		},
// 		{
// 			value:    []int{5, 5, 5, 5, 5},
// 			expected: 5,
// 		},
// 	}

// 	for _, item := range cases {
// 		a := Mid(item.value)

// 		if a != item.expected {
// 			t.Error(item.value)
// 		}
// 	}

// }

func TestMid(t *testing.T) {
	assert.Equal(t, float64(1), Mid([]int{1, 1, 1}))
	assert.Equal(t, float64(1.6), Mid([]int{1, 1, 1, 2, 3}))
	assert.Equal(t, float64(3), Mid([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, float64(30), Mid([]int{10, 20, 30, 40, 50}))
	assert.Equal(t, float64(0), Mid([]int{}))
	assert.Equal(t, float64(0), Mid([]int{0}))
	assert.Equal(t, float64(0), Mid([]int{0, 0, 0, 0, 0, 0}))

}

func TestPopulationVariance1(t *testing.T) {
	assert.Equal(t, float64(106.56), "1 4 6 10 30")
	assert.Equal(t, float64(2), "1 2 3 4 5")
	assert.Equal(t, float64(0), "0 0 0")

}
