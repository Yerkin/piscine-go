package piscine

func Sqrt(nb int) int {
	for i := 1; i <= nb; i++ {
		if IterativePower(i, 2) == nb {
			return i
		}
	}
	return 0
}
