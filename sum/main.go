package sum

func Sum[T int | float64 | float32](a, b T) T { return a + b }
