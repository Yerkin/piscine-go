package piscine

func IterativeFactorial(nb int) int {
	a := 1
	for i := 1; i <= nb; i++ {
		a = a * i
	}

	return a
}
