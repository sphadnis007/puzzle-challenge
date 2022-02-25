package utils

func GetRotatedElem(c byte, n int) byte {
	n %= 26
	//fmt.Println("initial", c, string(c), "n", byte(n))
	c += byte(n)
	//fmt.Println("now", c, string(c))
	if c < 'a' {
		diff := 'a' - c
		c = 'z' - (diff - 1)
	}

	if c > 'z' {
		diff := c - 'z'
		c = 'a' + (diff - 1)
	}

	//fmt.Println("final", string(c))
	return c
}

func Rotate1DArray(arr []byte, num int) []byte {
	res := make([]byte, len(arr))
	for index, elem := range arr {
		res[index] = GetRotatedElem(elem, num)
	}
	return res
}
