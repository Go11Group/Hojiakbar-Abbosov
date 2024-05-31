package methods

import (
	"encoding/json"
	"fmt"
	"os"
)

func CompReadJsonFile() []Company {

	comp := []Company{}

	ucompanyFile, err := os.OpenFile("companies.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("File is not ")
		panic(err)
	}
	t, err := ucompanyFile.Stat()
	if err != nil {
		panic(err)
	}
	data := make([]byte, t.Size())

	err = json.Unmarshal(data, comp)
	if err != nil {
		panic(err)
	}
	return comp
}

func ManagerReadwithFile() []Recruiter {
	manager := []Recruiter{}
	managerfile, err := os.OpenFile("manager.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("File is not ")
		panic(err)
	}
	t, err := managerfile.Stat()
	if err != nil {
		panic(err)
	}
	data := make([]byte, t.Size())

	err = json.Unmarshal(data, manager)
	if err != nil {
		panic(err)
	}
	return manager

}

func VacansySeewithFile() {
	reuciters := []Recruiter{}

	companyFile, err := os.OpenFile("companies.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
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
				fmt.Println(reuciters[i].Companies[j].Vacancies[k].QueId)
			}
		}
	}
}

func InterviewSeewithFile() {
	reuciters := []Recruiter{}

	companyFile, err := os.OpenFile("companies.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
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
				fmt.Println(reuciters[i].Companies[j].Vacancies[k].Interviews)
				fmt.Println(reuciters[i].Companies[j].Vacancies[k].Interviews)
				fmt.Println(reuciters[i].Companies[j].Vacancies[k].Interviews) // interview ni yozmasa ham bo`ladi
				fmt.Println(reuciters[i].Companies[j].Vacancies[k].Interviews)
			}
		}
	}
}
