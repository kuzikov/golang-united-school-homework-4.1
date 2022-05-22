package reverse_string

// ReverseString get an utf8 string and returns it's reverse form.
func ReverseString(input string) (output string) {
	runeStr := []rune(input)
	tmp := make([]rune, 0, len(input))
	for i := len(runeStr) - 1; i >= 0; i-- {
		tmp = append(tmp, runeStr[i])
	}

	output = string(tmp)
	return output
}
