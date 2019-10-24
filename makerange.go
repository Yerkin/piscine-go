package piscine

func MakeRange(min, max int) []int {
	ln := max - min
	var ans []int

	if min >= max {
		return nil
	}

	ans = make([]int, ln)
	for i := 0; i < ln; i++ {
		ans[i] = min
		min++
	}
	return ans

}
