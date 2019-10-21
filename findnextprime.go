package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		for {
			if IsPrime(nb) {
				return nb
			}
			nb = nb + 1
		}
	}
}
