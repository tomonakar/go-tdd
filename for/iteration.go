package iteration

func Repeat(character string, n int) string {
	if n == 0 {
		return ""
	}

	var repeat string
	for i := 0; i < n; i++ {
		repeat += character
	}
	return repeat
}
