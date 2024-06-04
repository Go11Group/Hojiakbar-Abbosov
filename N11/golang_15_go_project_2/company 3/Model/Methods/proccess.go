package methods

import (
	"encoding/json"
	"fmt"
	"os"
)

func (recuiter *Recruiter) Proccess(recuiterIndex, companyIndex, vacancyIndex, userId int) string {
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
	// reuciters[recuiterIndex-1].Companies[companyIndex-1].Vacancies[vacancyIndex-1].Interviews[] = append(reuciters[recuiterIndex-1].Companies[companyIndex-1].Vacancies[vacancyIndex].Interviews, )
	return "Success"

}
