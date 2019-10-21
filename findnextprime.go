package piscine

func FindNextPrime(nb int) int {
	b := true
	if IsPrime(nb) {
		return nb
	} else {
		for b {
			nb = nb + 1
			if IsPrime(nb) {
				b = false
			}
		}
	}
	return nb
}
