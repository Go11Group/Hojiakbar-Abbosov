package methods

import (
	"encoding/json"
	"fmt"
	"os"
)

func (recuiter *Recruiter) ReadCompany() {
	recuiters := []Recruiter{}
	companyFile, err := os.OpenFile("companies.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("File is not ")
		panic(err)
	}

	data := json.NewDecoder(companyFile)
	err = data.Decode(&recuiters)
	if err != nil {
		fmt.Println()
	}

	if len(recuiters) == 0 {
		fmt.Println("Fayl bo`sh!")
	} else {
		for i := 0; i < len(recuiters); i++ {
			fmt.Println(recuiters[i])
		}
	}

}

func (recuiter *Recruiter) UpdateCompany(quId int) {

	recuiters := []Recruiter{}
	company := Company{}

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
	var queid int
	for i := 0; i < len(recuiters); i++ {
		for j := 0; j < len(recuiters[i].Companies); j++ {

			if recuiters[i].Companies[j].Id == queid {
				fmt.Println("Name:\t")
				fmt.Scan(&company.Name)
				recuiters[i].Companies[j].Name = company.Name
				fmt.Println("Id:\t")
				fmt.Scan(&company.Id)
				recuiters[i].Companies[j].Id = company.Id
				fmt.Println("Phone:\t")
				fmt.Scan(&company.Phone)
				recuiters[i].Companies[j].Phone = company.Phone
				fmt.Println("QueId:\t")
				recuiters[i].Companies = append(recuiters[i].Companies, company)
				break
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

func (recuiter *Recruiter) DeleteCompany(queId int) {
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
	// companies1 := []Company{}
	son := 0
	for i := 0; i < len(recuiters); i++ {
		for k := 0; k < len(recuiters[i].Companies); k++ {
			if recuiters[i].Companies[k].Id == queId {
				// companies1 = append(companies1[:k], companies1[k+1:]...)
				recuiters[i].Companies = append(recuiters[i].Companies[:k], recuiters[i].Companies[k+1:]...)
				son = i
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

func (recuiter *Recruiter) CreateCompany(company Company) {
	recuiters := []Recruiter{}
	recuiters[len(recuiters)].Companies = append(recuiters[len(recuiters)].Companies, company)
	
	f, err := os.OpenFile("companyINregister.json", os.O_APPEND | os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	data := json.NewDecoder(f)
	err = data.Decode(&recuiters)
	if err != nil {
		panic(err)
	}
	



}
