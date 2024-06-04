package methods

import (
	"encoding/json"
	"fmt"

	// methods "my_project/Model/Methods"

	"os"
)

func (recuiter *Recruiter) CreateInterview(id int, interview1 Interview) string {
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
	var a, b bool
	a = false
	b = false
	for i := 0; i < len(reuciters); i++ {
		for j := 0; j < len(reuciters[i].Companies); j++ {
			for k := 0; k < len(reuciters[i].Companies[j].Vacancies); k++ {
				if reuciters[i].Id == interview1.Id {
					reuciters[i].Companies[j].Vacancies[k].Interviews = append(reuciters[i].Companies[j].Vacancies[k].Interviews, interview1)
					break
				}
				a = true
				break
			}
			if a == true {
				b = true
				break
			}
		}
		if b == true {
			break
		}
	}
	info3 := json.NewEncoder(companyFile)
	err = info3.Encode(reuciters)
	if err != nil {
		return ("Error is not write file ")
	}
	return "Success "

}

func (recuiter *Recruiter) UpdateInterview(interview1 Interview) string {
	reuciters := []Recruiter{}
	flag := 1
	q := 0
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
	var a, b bool
	a = false
	b = false
	for i := 0; i < len(reuciters); i++ {
		for j := 0; j < len(reuciters[i].Companies); j++ {
			for k := 0; k < len(reuciters[i].Companies[j].Vacancies); k++ {
				q++
				if reuciters[i].Id == interview1.Id {
					reuciters[i].Companies[j].Vacancies[k].Interviews[q].Id = interview1.Id
					reuciters[i].Companies[j].Vacancies[k].Interviews[q].Time = interview1.Time
					reuciters[i].Companies[j].Vacancies[k].Interviews[q].Place = interview1.Place
					a = true
					return "is added "
				} else {
					flag = 0
				}
			}
			if a == true {
				b = true
				break
			}
		}
		if b == true {
			break
		}
	}
	info3 := json.NewEncoder(companyFile)
	err = info3.Encode(reuciters)
	if err != nil {
		return ("Error is not write file ")
	}
	if flag == 0 {
		return "Sizda bunaqa interviyu yoq!"

	}

	return "Success "

}

func (recuiter *Recruiter) ReadInterview(id int) string {
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
	var a, b bool
	a = false
	b = false
	for i := 0; i < len(reuciters); i++ {
		for j := 0; j < len(reuciters[i].Companies); j++ {
			for k := 0; k < len(reuciters[i].Companies[j].Vacancies); k++ {
				if reuciters[i].Id == id {
					fmt.Println(reuciters[i].Companies[j].Vacancies[k].Interviews)
					fmt.Println(reuciters[i].Companies[j].Vacancies[k].Interviews)
					fmt.Println(reuciters[i].Companies[j].Vacancies[k].Interviews) // interview ni yozmasa ham bo`ladi
					fmt.Println(reuciters[i].Companies[j].Vacancies[k].Interviews)
					a = true
					break
				}
			}
			if a == true {
				b = true
				break
			}
		}
		if b == true {
			break
		}
	}
	info3 := json.NewEncoder(companyFile)
	err = info3.Encode(reuciters)
	if err != nil {
		return ("Error is not write file ")
	}
	return "Success "

}

// delete interview part

func (recuiter *Recruiter) DeleteInterview(id int) string {
	reuciters := []Recruiter{}
	// file is create is mode read write and append data in file if file is exits else file error Eof
	companyFile, err := os.OpenFile("companies.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("File is not ")
		panic(err)
	}

	// file is read and parse in struct slice which is companyfile from parse  reuciters
	data := json.NewDecoder(companyFile)
	err = data.Decode(&reuciters)
	if err != nil {
		fmt.Println("Dekodrga err bor")
	}
	var a, b bool
	a = false
	b = false
	for i := 0; i < len(reuciters); i++ {
		for j := 0; j < len(reuciters[i].Companies); j++ {
			for k := 0; k < len(reuciters[i].Companies[j].Vacancies); k++ {
				if reuciters[i].Id == id {
					reuciters[i].Companies[j].Vacancies[k].Interviews = nil
					a = true
					break
				}
			}
			if a == true {
				b = true
				break
			}
		}
		if b == true {
			break
		}
	}
	info3 := json.NewEncoder(companyFile)
	err = info3.Encode(reuciters)
	if err != nil {
		return ("Error is not write file ")
	}
	return "Success "

}
