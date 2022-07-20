package test

import (
	"fmt"
	"test-generic/sum"
	"testing"
)

func Test1_SUM(t *testing.T) {
	var sum1 = sum.Sum(1, 2)
	fmt.Println("Sum of 1 and 2:", sum1)

	var sum2 = sum.Sum[float32](.1, .2)
	fmt.Println("Sum of .1 and .2:", sum2)

	var sum3 = sum.Sum(.1, .2)
	fmt.Println("Sum of .2 and .3:", sum3, fmt.Sprintf("%.30f", sum3))
}
