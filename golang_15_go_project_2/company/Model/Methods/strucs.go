package methods

type User struct {
	Id      int
	Name    string
	Surname string
	Age     int
}

type Recruiter struct {
	Id        int
	Companies []Company
	Name      string
	Surname   string
	Email     string	
	Password  string
}

type Company struct {
	Id        int
	Name      string
	Phone     string
	Address   string
	Vacancies []Vacancy	
}

type Vacancy struct {
	Id         int
	Count      int
	Comments   string
	Interviews []Interview
	QueId      int
}

type Interview struct {
	Id         int
	Time       string
	Place      string
	Candidates []User
}
