package methods

import (
	"encoding/json"
	"fmt"
	"os"
)

// Create vakasy , first we read data from json file  and parse  struct of requiters

func (recuiter *Recruiter) CreateVacansy(recuredId, companyId int, vakansy1 Vacancy) bool {
	recuiters := []Recruiter{}

	companyFile, err := os.OpenFile("manager.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
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
		if recuiters[i].Id == recuredId {
			for j := 0; j < len(recuiters[i].Companies); j++ {
				if recuiters[i].Companies[j].Id == companyId {
					recuiters[i].Companies[j].Vacancies = append(recuiters[i].Companies[j].Vacancies, vakansy1)
					return true
				}
			}
		}
	}
	companyFile.Truncate(0)
	companyFile.Seek(0, 0)
	info1 := json.NewEncoder(companyFile)
	err = info1.Encode(recuiters)
	if err != nil {
		return false
	}
	return false

}

func (recuiter *Recruiter) DeleteVacansy(recuiredId, companyid1, vacansyId int) bool {
	recuiters := []Recruiter{}

	companyFile, err := os.OpenFile("manager.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
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
		if recuiters[i].Id == recuiredId {
			for j := 0; j < len(recuiters[i].Companies); j++ {
				if recuiters[i].Companies[j].Id == companyid1 {
					for k := 0; k < len(recuiters[i].Companies[j].Vacancies); k++ {
						if recuiters[i].Companies[j].Vacancies[k].Id == vacansyId {
							recuiters[i].Companies[j].Vacancies = nil
							return true
						}
					}
				}
			}
		}
	}
	companyFile.Truncate(0)
	companyFile.Seek(0, 0)
	info1 := json.NewEncoder(companyFile)
	err = info1.Encode(recuiters)
	if err != nil {
		return false
	}
	return false

}

func (recuiter *Recruiter) UpdateVancy(recuireId3, companyid3, vacancyId3 int, vacansy1 Vacancy) bool {
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
		if reuciters[i].Id == recuireId3 {
			for j := 0; j < len(reuciters[i].Companies); j++ {
				if reuciters[i].Companies[j].Id == companyid3 {
					for k := 0; k < len(reuciters[i].Companies[j].Vacancies); k++ {
						if reuciters[i].Companies[j].Vacancies[k].Id == vacancyId3 {
							reuciters[i].Companies[j].Vacancies[k].Id = vacansy1.Id
							reuciters[i].Companies[j].Vacancies[k].Count = vacansy1.Count
							reuciters[i].Companies[j].Vacancies[k].Comments = vacansy1.Comments
							return true
						}
					}
				}
			}
		}
	}
	companyFile.Truncate(0)
	companyFile.Seek(0, 0)
	info1 := json.NewEncoder(companyFile)
	err = info1.Encode(reuciters)
	if err != nil {
		return false
	}
	return false

}

func (recuiter *Recruiter) ReadVanacy(recuiterId4, companyId4 int) {
	recuiters := []Recruiter{}

	companyFile, err := os.OpenFile("manager.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
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
		if recuiters[i].Id == recuiterId4 {
			for j := 0; j < len(recuiters[i].Companies); j++ {
				if recuiters[i].Companies[j].Id == companyId4 {
					for k := 0; k < len(recuiters[i].Companies[j].Vacancies); k++ {
						fmt.Println(recuiters[i].Companies[j].Vacancies[k].Count, recuiters[i].Companies[j].Vacancies[k].Count)
					}
				}
			}
		}
	}

}
