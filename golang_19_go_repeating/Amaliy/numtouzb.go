package main

import (
 "fmt"
)

func numberToWordsUZ(n int) string {
 	if n < 1 || n > 9999 {
  		return "Raqam 1 dan 9999 gacha bo'lishi kerak"
 	}

 	var (
	yuzlar  = []string{"", "bir yuz", "ikki yuz", "uch yuz", "to'rt yuz", "besh yuz", "olti yuz", "yetti yuz", "sakkiz yuz", "to'qqiz yuz"}
	onlik   = []string{"", "o'n", "yigirma", "o'ttiz", "qirq", "ellik", "oltmish", "yetmish", "sakson", "to'qson"}
	birlik  = []string{"", "bir", "ikki", "uch", "to'rt", "besh", "olti", "yetti", "sakkiz", "to'qqiz"}
	minglik = []string{"", "bir ming", "ikki ming", "uch ming", "to'rt ming", "besh ming", "olti ming", "yetti ming", "sakkiz ming", "to'qqiz ming"}
 	)

	ming := n / 1000
	yuz := (n % 1000) / 100
	on := (n % 100) / 10
	bir := n % 10

 	result := minglik[ming]
 	if ming > 0 && yuz > 0 {
  		result += " "
 	}
 	result += yuzlar[yuz]

 	if yuz > 0 && (on > 0 || bir > 0) {
  		result += " "
 	}

 	if on > 0 {
  		result += onlik[on]
  	if bir > 0 {
   		result += " "
  	}
 }

 	result += birlik[bir]

 	return result
}

func main() {

	fmt.Println(numberToWordsUZ(100))  
	fmt.Println(numberToWordsUZ(0))    
	fmt.Println(numberToWordsUZ(7000)) 
}