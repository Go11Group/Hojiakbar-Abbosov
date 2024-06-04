package main

import (
	"fmt"
	"os"
)

func main(){

	f,err := os.Create("file.txt")

	if err != nil{
		panic(err)
	}
	defer f.Close()

	_,err = f.WriteString("fdadweffw")

	if err != nil{
		panic(err)
	}
	
	n,err := f.Seek(0,0)
	if err != nil{
		panic(err)
	}
	fmt.Println(n)

	a := make([]byte,40)
	_,err = f.Read(a)
	if err != nil{
		panic(err)
	}

}