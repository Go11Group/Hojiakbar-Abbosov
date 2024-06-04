package string

func SplitLetters(input string) []string {
	var letters []string
	for _, char := range input {
		letters = append(letters, string(char))
	}
	return letters
}