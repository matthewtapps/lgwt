package iteration

func Repeat(input string, iterations int) (repeated string) {
	for i := 0; i < iterations; i++ {
		repeated += input
	}
	return repeated
}
