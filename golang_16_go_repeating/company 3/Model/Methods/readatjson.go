package methods

import (
	"encoding/json"
	"fmt"
	"os"
)

func VacansySeewithFile() {
	reuciters := []Recruiter{}

	companyFile, err := os.OpenFile("manager.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("File is not ")
		panic(err)
	}

	data := json.NewDecoder(companyFile)
	err = data.Decode(&reuciters)
	if err != nil {
		fmt.Println("Dekodrga err bor")
	}
	for i := 0; i < len(reuciters); i++ {
		for j := 0; j < len(reuciters[i].Companies); j++ {
			for k := 0; k < len(reuciters[i].Companies[j].Vacancies); k++ {
				fmt.Println(reuciters[i].Companies[j].Vacancies[k].Id)
				fmt.Println(reuciters[i].Companies[j].Vacancies[k].Count)
				fmt.Println(reuciters[i].Companies[j].Vacancies[k].Comments)
			}
		}
	}
}

func InterviewSeewithFile() {
	reuciters := []Recruiter{}

	companyFile, err := os.OpenFile("manager.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("File is not ")
		panic(err)
	}

	data := json.NewDecoder(companyFile)
	err = data.Decode(&reuciters)
	if err != nil {
		fmt.Println("Dekodrga err bor")
	}
	for i := 0; i < len(reuciters); i++ {
		for j := 0; j < len(reuciters[i].Companies); j++ {
			for k := 0; k < len(reuciters[i].Companies[j].Vacancies); k++ {
				for l := 0; l < len(reuciters[i].Companies[j].Vacancies[k].Interviews); l++ {
					fmt.Println(reuciters[i].Companies[j].Vacancies[k].Interviews[l].Place)
					fmt.Println(reuciters[i].Companies[j].Vacancies[k].Interviews[l].Time)
				}
			}
		}
	}
}

func CompanyViewUser() {
	reuciters := []Recruiter{}

	companyFile, err := os.OpenFile("manager.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("File is not ")
		panic(err)
	}

	data := json.NewDecoder(companyFile)
	err = data.Decode(&reuciters)
	if err != nil {
		fmt.Println("Dekodrga err bor")
	}
	for i := 0; i < len(reuciters); i++ {
		for j := 0; j < len(reuciters[i].Companies); j++ {
			fmt.Println(reuciters[i].Companies[j].Name)
		}
	}
}
