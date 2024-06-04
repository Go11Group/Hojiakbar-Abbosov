package methods

import (
	"encoding/json"
	"fmt"

	// methods "my_project/Model/Methods"

	"os"
)

func (recuiter *Recruiter) CreateInterview(recuiterid5, companyId5, vacancyId5 int, interview1 Interview) bool {
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
		if reuciters[i].Id == recuiterid5 {
			for j := 0; j < len(reuciters[i].Companies); j++ {
				if reuciters[i].Companies[j].Id == companyId5 {
					for k := 0; k < len(reuciters[i].Companies[j].Vacancies); k++ {
						if reuciters[i].Companies[j].Vacancies[j].Id == vacancyId5 {
							reuciters[i].Companies[j].Vacancies[j].Interviews = append(reuciters[i].Companies[j].Vacancies[j].Interviews, interview1)
							return true
						}
					}
				}
			}
		}
	}
	info3 := json.NewEncoder(companyFile)
	err = info3.Encode(reuciters)
	if err != nil {
		return false
	}
	return true

}

func (recuiter *Recruiter) UpdateInterview(recuiterId6, companyId6, vacancyId6, inetrviewId6 int, interview1 Interview) bool {
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
		if reuciters[i].Id == recuiterId6 {
			for j := 0; j < len(reuciters[i].Companies); j++ {
				if reuciters[i].Companies[j].Id == companyId6 {
					for k := 0; k < len(reuciters[i].Companies[j].Vacancies); k++ {
						if reuciters[i].Companies[j].Vacancies[k].Id == vacancyId6 {
							for l := 0; l < len(reuciters[i].Companies[j].Vacancies[k].Interviews); l++ {
								if reuciters[i].Companies[j].Vacancies[k].Interviews[k].Id == inetrviewId6 {
									reuciters[i].Companies[j].Vacancies[k].Interviews[k].Place = interview1.Place
									reuciters[i].Companies[j].Vacancies[k].Interviews[k].Time = interview1.Time
									reuciters[i].Companies[j].Vacancies[k].Interviews[k].Candidates = interview1.Candidates
									return true

								}
							}
						}
					}
				}
			}
		}
	}
	info3 := json.NewEncoder(companyFile)
	err = info3.Encode(reuciters)
	if err != nil {
		return false
	}
	return false

}

func (recuiter *Recruiter) ReadInterview(recuiterid7,companyId7,vacansyid7 int)  {
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
		if reuciters[i].Id==recuiterid7{
			for j := 0; j < len(reuciters[i].Companies); j++ {
				if reuciters[i].Companies[j].Id==companyId7{
					for k := 0; k < len(reuciters[i].Companies[j].Vacancies); k++ {
						if reuciters[i].Companies[j].Vacancies[k].Id==vacansyid7{
							for l := 0; l < len(reuciters[i].Companies[j].Vacancies[k].Interviews); l++ {
								fmt.Println(reuciters[i].Companies[j].Vacancies[k].Interviews[l].Candidates)
								fmt.Println(reuciters[i].Companies[j].Vacancies[k].Interviews[l].Place)
								fmt.Println(reuciters[i].Companies[j].Vacancies[k].Interviews[l].Time)
							}
						}
					}
				}
			}
		}
	}
	info3 := json.NewEncoder(companyFile)
	err = info3.Encode(reuciters)
	if err != nil {
		fmt.Println()
	}
	fmt.Println()

}

// delete interview part

func (recuiter *Recruiter) DeleteInterview(recuiterid8,companyid8,vacansyid8,interviewid8 int) bool {
	reuciters := []Recruiter{}
	// file is create is mode read write and append data in file if file is exits else file error Eof
	companyFile, err := os.OpenFile("manager.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("File is not ")
		panic(err)
	}

	// file is read and parse in struct slice which is companyfile from parse  reuciters
	data := json.NewDecoder(companyFile)
	err = data.Decode(&reuciters)
	if err!=nil{
		fmt.Println()
	}
	for i := 0; i < len(reuciters); i++ {
		if reuciters[i].Id==recuiterid8{
			for j := 0; j < len(reuciters[i].Companies); j++ {
				if reuciters[i].Companies[j].Id==companyid8{
					for k := 0; k < len(reuciters[i].Companies[j].Vacancies); k++ {
						if reuciters[i].Companies[j].Vacancies[k].Id==vacansyid8{
							for l := 0; l < len(reuciters[i].Companies[j].Vacancies[k].Interviews); l++ {
								if reuciters[i].Companies[j].Vacancies[k].Interviews[l].Id==interviewid8{
									reuciters[i].Companies[j].Vacancies[k].Interviews=nil
									return true
								}
							}
						}
					}
				}
			}
		}
	}
	
	info3 := json.NewEncoder(companyFile)
	err = info3.Encode(reuciters)
	if err != nil {
		return false
	}
	return false

}
