package string

func ConcatenateStrings(stringsToConcat ...string) string {
    concatenated := ""
    for _, str := range stringsToConcat {
        concatenated += str
    }
    return concatenated
}