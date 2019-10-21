package piscine

func IterativePower(nb int, power int) int {
	a := nb
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		for i := 1; i < power; i++ {
			nb = nb * a
		}
		return nb
	}
}
