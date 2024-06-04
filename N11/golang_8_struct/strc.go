package main

import "fmt"

type building struct{
	name string
	isGov bool
	location location
}
type hospital struct{
	building
	employee int
	rating float32
}


type location struct{
	lat , lot float32
}


func main(){

	hospitals:= []hospital{
		{
	 	building : building{
		name: "shohmed",
		isGov: false,
		location: location{lat: 23.24,lot: 43.44},

		},
		employee:140,
		rating:4.5,
		},
		{},
	}


	for _,i :=  range hospitals{
		
	}

	fmt.Println(hospitals[0].employee,hospitals[0].location.lat,hospitals[0].isGov)



}