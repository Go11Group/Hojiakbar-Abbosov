package methods

import (
	"encoding/json"
	"fmt"
	"os"
)

func (recuiter *Recruiter) ReadUserInterview(id int) {
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

	

}
