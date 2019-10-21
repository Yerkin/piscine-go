package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		for {
			nb = nb + 1
			if IsPrime(nb) {
				return nb
			}
		}
	}
}
