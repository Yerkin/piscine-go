package piscine

func IsPrime(nb int) bool {
	c := 0
	for i := 2; i <= nb; i++ {
		if nb%i == 0 {
			c++
		}
	}
	if c == 2 {
		return true
	} else {
		return false
	}
}
