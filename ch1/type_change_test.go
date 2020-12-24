package ch1

import (
	"testing"
)

func TestIntToFloat(t *testing.T) {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	t.Logf("the value of mean is: %f\n", mean)
}
