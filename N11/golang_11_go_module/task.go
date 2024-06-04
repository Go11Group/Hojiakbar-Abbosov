package main

import "fmt"

type Student struct{
	name string
	have bool
}

type Teacher struct{
	class map[string][]Student 
}

func main(){

	teacher := Teacher{
		class: map[string][]Student{
			"06-05":{
				{name: "Ilyos", have: true},
				{name: "Javohir", have: false},
				{name: "Aziz", have: true},
				{name: "Husan", have: true},
			},
		},
	}

	for date, students := range teacher.class{
		fmt.Printf("%v kuni uchun davomat\n ",date)
		count := 0

		for _ , i := range students{
			if i.have {
				count++
				fmt.Printf("%s  bor  \n",i.name)
			}else {
				fmt.Printf("%s  yo'q \n",i.name)
				
			}
		}
		fmt.Printf("Jami %d ta o'quvchi bor\n",count)
	}

}