package string

func RemoveVowels(input string) string {
	var result string

	for _, char := range input {
		if char != 'A' && char != 'a' && char != 'E' && char != 'e' && char != 'I' && char != 'i' && char != 'O' && char != 'o' && char != 'U' && char != 'u' {
			result += string(char)
		}
	}

	return result
}