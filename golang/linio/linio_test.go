package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinioTest(t *testing.T) {
	testCases := []struct {
		input  uint
		result string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Linio"},
		{4, "4"},
		{5, "IT"},
		{6, "Linio"},
		{7, "7"},
		{8, "8"},
		{9, "Linio"},
		{10, "IT"},
		{11, "11"},
		{12, "Linio"},
		{13, "13"},
		{14, "14"},
		{15, "Linianos"},
		{16, "16"},
		{17, "17"},
		{18, "Linio"},
		{19, "19"},
		{20, "IT"},
		{21, "Linio"},
		{22, "22"},
		{23, "23"},
		{24, "Linio"},
		{25, "IT"},
		{26, "26"},
		{27, "Linio"},
		{28, "28"},
		{29, "29"},
		{30, "Linianos"},
	}

	for _, test := range testCases {
		output := LinioTest(test.input)
		if !assert.Equal(t, test.result, output) {
			t.Fail()
		}
	}
}

func BenchmarkLinioTest(b *testing.B) {
	bN := make([]uint, b.N)
	for i, _ := range bN {
		bN[i] = uint(rand.Uint64())
	}
	b.ResetTimer()
	for _, num := range bN {
		LinioTest(num)
	}
}
