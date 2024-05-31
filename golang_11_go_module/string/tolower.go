package string

func ConvertToLower(n string) string {
    var result string
    for _, char := range n {
        if char >= 'A' && char <= 'Z' {
            result += string(char + 32)
        } else {
            result += string(char)
        }
    }
    return result
}