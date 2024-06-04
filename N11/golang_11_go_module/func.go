package main

import (
	"fmt"
	"my_module/string"
)

func main(){

	fmt.Println(string.ConcatenateTwo("hello","world"))
	fmt.Println(string.ConcatenateStrings("hello","world","goodbye"))
	fmt.Println(string.ConvertToLower("Hello"))
	fmt.Println(string.ConvertToUpperCase("hello"))
	fmt.Println(string.SplitLetters("hello"))
	fmt.Println(string.RemoveVowels("hello"))


}