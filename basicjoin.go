package piscine

func BasicJoin(strs []string) string {
	my_str := ""
	for _, w := range strs {
		my_str = my_str + w
	}
	return my_str
}
