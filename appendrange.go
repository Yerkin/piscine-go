package piscine

func AppendRange(min, max int) []int {
	var ans []int
	for i := min; i < max; i++ {
		ans = append(ans, i)
	}
	return ans
}
