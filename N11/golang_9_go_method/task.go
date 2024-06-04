package main

import "fmt"

type book struct{
	name string
	author string
	page int
	bookNo int 
	year int 
	pages []string
}
func (b *book) setBook(){
	b.bookNo = 123
	b.year = 2022
	b.author = "Tohir Malik"
	b.page = 345
	b.name = "Alvido bolalik"
}

func (b *book) setAuthor(){

}



func main(){

	bk := book{}

	bk.setBook()

	fmt.Println(bk)

}