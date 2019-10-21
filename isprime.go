package piscine

func IsPrime(nb int) bool {
	c := 1
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
