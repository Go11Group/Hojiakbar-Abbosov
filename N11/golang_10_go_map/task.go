package main

import "fmt"

type Person struct{
	FirstName string
	LastName string
	age int
	gmail string
}
func (a *Person) get(){
	fmt.Println("Enter  First name")
	fmt.Scanln(&a.FirstName)
	fmt.Println("Enter  Last name")
	fmt.Scanln(&a.LastName)
	fmt.Println("Enter  age")
	fmt.Scanln(&a.age)
	fmt.Println("Enter  gmail")
	fmt.Scanln(&a.gmail)
}
func (a *Person) info() {
	fmt.Printf("%s\n", a.FirstName)
	fmt.Printf("%s\n", a.LastName)
	fmt.Printf("%d\n", a.age)
	fmt.Printf("%s\n", a.gmail)
  }



func main(){

	var Person = Person{}

	Person.get()
	Person.info()

}



