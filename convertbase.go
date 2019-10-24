package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	ln1 := 0
	ln2 := 0
	ln3 := 0
	ans := ""

	for i := range nbr {
		ln1 = i
	}

	arr1 := []rune(nbr)

	for i := range baseFrom {
		ln2 = i
	}
	ln2++

	for i := range baseTo {
		ln3 = i + 1
	}

	myNum := 0
	pow := 1

	for i := ln1; i >= 0; i-- {
		myNum = myNum + int(arr1[i]-'0')*pow
		pow *= ln2
	}

	for myNum != 0 {
		ans = string(baseTo[myNum%ln3]) + ans
		myNum /= ln3
	}

	return ans
}
