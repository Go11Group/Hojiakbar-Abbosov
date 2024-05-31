package main

import "fmt"

func main() {
	s := "Hello world saloM dunya"
	a := reverseWords(s)
	fmt.Println(a) 
   }

func reverseWords(s string) string {
 	reversed := ""
 	word := ""
 	for _, char := range s {
  		if char == ' ' {
   			reversed = (word) + " " + reversed
   			word = ""
  		} else {
   			word += string(char)
  		}
 	}
 		reversed = (word) + " " + reversed
 		return reversed[:len(reversed)-1]
}

