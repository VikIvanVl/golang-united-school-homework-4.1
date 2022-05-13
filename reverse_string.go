package reverse_string

func ReverseString(input string) (output string) {
	r := []rune(input)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	output = string(r)
	return output
}
