package methods

import (
	"encoding/json"
	"fmt"
	"os"
)

func (recuiter *Recruiter) ReadCompany(id int) {
	recuiters := []Recruiter{}
	companyFile, err := os.OpenFile("manager.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("File is not ")
	}

	data := json.NewDecoder(companyFile)
	err = data.Decode(&recuiters)
	if err != nil {
		fmt.Println()
	}

	for i := 0; i < len(recuiters); i++ {
		if recuiters[i].Id == id {
			for j := 0; j < len(recuiters[i].Companies); j++ {
				fmt.Println(recuiters[i].Companies[j].Name)
			}
		}
	}

}

func (recuiter *Recruiter) UpdateCompany(companyid int, recuiterid int) {

	recuiters := []Recruiter{}
	company := Company{}

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
		if recuiters[i].Id == recuiterid {
			for j := 0; j < len(recuiters[i].Companies); j++ {

				if recuiters[i].Companies[j].Id == companyid {
					fmt.Println("Name:\t")
					fmt.Scan(&company.Name)
					recuiters[i].Companies[j].Name = company.Name
					fmt.Println("Id:\t")
					fmt.Scan(&company.Id)
					recuiters[i].Companies[j].Id = company.Id
					fmt.Println("Phone:\t")
					fmt.Scan(&company.Phone)
					recuiters[i].Companies[j].Phone = company.Phone
					fmt.Println("id:\t")
					recuiters[i].Companies = append(recuiters[i].Companies, company)
					break
				}
			}
		}
	}

	companyFile.Truncate(0)
	companyFile.Seek(0, 0)
	info := json.NewEncoder(companyFile)
	err = info.Encode(recuiters)
	if err != nil {
		fmt.Println("Is not Updated")
		panic(err)
	}

}

// delete company

func (recuiter *Recruiter) DeleteCompany(companyid int, requiredid int) {
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
	son := 0
	for i := 0; i < len(recuiters); i++ {
		if recuiters[i].Id == requiredid {
			for k := 0; k < len(recuiters[i].Companies); k++ {

				if recuiters[i].Companies[k].Id == companyid {
					recuiters[i].Companies = append(recuiters[i].Companies[:k], recuiters[i].Companies[k+1:]...)
					son = i
				}
			}
		}
	}

	companyFile.Truncate(0)
	companyFile.Seek(0, 0)
	data1 := json.NewEncoder(companyFile)
	err = data1.Encode(&recuiters[son])
	if err != nil {
		panic(err)
	}

}

// create Company

func (recuiter *Recruiter) CreateCompany(company Company, index int) {
	recuiters := []Recruiter{}
	f, err := os.OpenFile("manager.json", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	data := json.NewDecoder(f)
	err = data.Decode(&recuiters)
	if err != nil {
		fmt.Println(" ")
	}
	fmt.Println(recuiters)
	for i := 0; i < len(recuiters); i++ {
		if recuiters[i].Id == index {
			recuiters[i].Companies = append(recuiters[i].Companies, company)
		}
	}
	f.Truncate(0)
	f.Seek(0, 0)
	data1 := json.NewEncoder(f)
	err = data1.Encode(recuiters)
	if err != nil {
		fmt.Println()
	}

}
