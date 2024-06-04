package string

func ConvertToUpperCase(n string) string {
    var result string
    for _, char := range n {
        if char >= 'a' && char <= 'z' {
            result += string(char - 32)
        } else {
            result += string(char)
        }
    }
    return result
}