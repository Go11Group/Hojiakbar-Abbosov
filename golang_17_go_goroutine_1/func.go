package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

type Data struct{
	id int
	first_name string
	last_name string
	email string
	age string
}

func main() {

	jsonFile,err := os.Open("data.json")
	
	if err != nil{
		fmt.Println(err)
	}
	
	defer jsonFile.Close()

	var data []Data 

	jsonDecoder := json.NewDecoder(jsonFile)

	err = jsonDecoder.Decode(&data)

	if err != nil{
		fmt.Println(err)
	}


	xlsx := execelize.NewFile()

	sheetName := "Data"

	index := xlsx.NewSheet(sheetName)

	date := []string{"id","first_name","last_name","email","age"}

	for i , j := range date{
		n := execelize.
	}




}