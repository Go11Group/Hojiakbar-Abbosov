package methods

import (
	"encoding/json"
	"fmt"
	"os"
)


// Create vakasy , first we read data from json file  and parse  struct of requiters

func (recuiter *Recruiter) CreateVacansy(vakansy1 *Vacancy) string {
	recuiters := []Recruiter{}


	companyFile, err := os.OpenFile("companies.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	// if file is not  panic to error in our programm  exit  else this programm continue
	if err != nil {
		fmt.Println("File is not ")
		panic(err)
	}

	data := json.NewDecoder(companyFile)
	err = data.Decode(&recuiters)
	if err != nil {
		fmt.Println("Dekodrga err bor")
	}
	for i := 0; i < len(recuiters); i++ {
		for j := 0; j < len(recuiters[i].Companies); j++ {
			for k := 0; k < len(recuiters[i].Companies[j].Vacancies); k++ {

				recuiters[i].Companies[j].Vacancies[k].Id = vakansy1.Id
				recuiters[i].Companies[j].Vacancies[k].Count = vakansy1.Count
				recuiters[i].Companies[j].Vacancies[k].Comments = vakansy1.Comments
				return "Success"
			}
		}
	}
	companyFile.Truncate(0)
	companyFile.Seek(0, 0)
	info1 := json.NewEncoder(companyFile)
	err = info1.Encode(recuiters)
	if err != nil {
		return "Is not Success"
	}
	return "false"

}

func (recuiter *Recruiter) DeleteVacansy(Id int) string {
	recuiters := []Recruiter{}
	

	companyFile, err := os.OpenFile("companies.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("File is not ")
		panic(err)
	}

	data := json.NewDecoder(companyFile)
	err = data.Decode(&recuiters)
	if err != nil {
		fmt.Println("Dekodrga err bor")
	}
	for i := 0; i < len(recuiters); i++ {
		for j := 0; j < len(recuiters[i].Companies); j++ {
			if recuiters[i].Id == Id {
				recuiters[i].Companies[j].Vacancies = nil
				return "Success"
			}

		}

	}
	companyFile.Truncate(0)
	companyFile.Seek(0, 0)
	info1 := json.NewEncoder(companyFile)
	err = info1.Encode(recuiters)
	if err != nil {
		return "Not Edit"
	}
	return "true"

}

func (recuiter *Recruiter) UpdateVancy(quId int, vacansy1 Vacancy) string {
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
				if reuciters[i].Id == vacansy1.Id {
					reuciters[i].Companies[j].Vacancies[k].Id = vacansy1.Id
					reuciters[i].Companies[j].Vacancies[k].Count = vacansy1.Count
					reuciters[i].Companies[j].Vacancies[k].Comments = vacansy1.Comments
					reuciters[i].Companies[j].Vacancies[k].QueId = vacansy1.Id
					return "Success "
				}
			}
		}
	}
	companyFile.Truncate(0)
	companyFile.Seek(0, 0)
	info1 := json.NewEncoder(companyFile)
	err = info1.Encode(reuciters)
	if err != nil {
		return "Not Edit"
	}
	return "not found"

}

func (recuiter *Recruiter) ReadVanacy(id int) string {
	recuiters := []Recruiter{}
	

	companyFile, err := os.OpenFile("companies.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("File is not ")
		panic(err)
	}

	data := json.NewDecoder(companyFile)
	err = data.Decode(&recuiters)
	if err != nil {
		fmt.Println("Dekodrga err bor")
	}
	a := true
	b := true
	for i := 0; i < len(recuiters); i++ {
		for j := 0; j < len(recuiters[i].Companies); j++ {
			for k := 0; k < len(recuiters[i].Companies[j].Vacancies); k++ {
				if recuiters[i].Id == id {

					fmt.Println(recuiters[i].Companies[j].Vacancies[j].Id)
					fmt.Println(recuiters[i].Companies[j].Vacancies[j].Count)
					fmt.Println(recuiters[i].Companies[j].Vacancies[j].Comments)
				}
				a = false
				return "Success"
			}
			if a == false {
				b = false
				return "Success"
			}
		}
		if b == false {
			return "Success"
		}
	}
	return "Is not Success"

}
