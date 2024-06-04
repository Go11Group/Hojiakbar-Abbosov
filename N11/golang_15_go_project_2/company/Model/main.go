package main

import (
	"encoding/json"
	"fmt"
	methods "my_project/Model/Methods"
	"os"
)

type User struct {
	Name string
	Id   int
}

func main() {


	methodInterview := methods.Recruiter{}
	methodVacansi := methods.Recruiter{}
	methodsCompany := methods.Recruiter{}

	var sikl1, sikl2, sikl4, sikl5, sikl6, sikl7 bool
	sikl1, sikl2, sikl4, sikl5, sikl6, sikl7 = true, true, true, true, true, true
	for sikl1 {
		fmt.Println("0->Exit\t1->User\t2->Recuiter")
		var tanlov1 int
		fmt.Scan(&tanlov1)
		switch tanlov1 {
		case 0:
			sikl1 = false
			fmt.Println("E'tiboringiz uchun Raxmat uchun")
		case 1:

			for sikl2 {
				fmt.Println("0->Exit\t1->Register\t2->Login")
				var tanlov2 int
				fmt.Scan(&tanlov2)
				switch tanlov2 {
				case 0:
					sikl1 = false
				case 1:
					methods.RegisterUser()
				case 2:
					// user uchun login qilish kerkak bobur
					var id int
					var name string
					fmt.Printf("Id kiriting: ")
					fmt.Scan(&id)
					fmt.Printf("Ismingizni kiriting: ")
					fmt.Scan(&name)

					userWfile, err := os.OpenFile("users.json", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
					if err != nil {
						panic(err)
					}
					defer userWfile.Close()

					t, err := userWfile.Stat()
					if err != nil {
						panic(err)
					}

					data := make([]byte, t.Size())
					var users []User

					err = json.Unmarshal(data, users)
					if err != nil {
						panic(err)
					}

					for _, k := range users {
						if k.Id == id && k.Name == name {
							fmt.Println("0->Exit\t1-Company\t2-Vakansiya\t3-Interviyu")
							var tanlov3 int
							fmt.Scan(&tanlov3)
							switch tanlov3 {
							case 0:
								return
							case 1:
								c1 := []methods.Company{}
								c1 = methods.CompReadJsonFile()
								// break
								// kompaniylarni ro'yxatini chiqarish kerak
								for i := 0; i < len(c1); i++ {
									fmt.Println(c1[i])
								}
							case 2:
								methods.VacansySeewithFile()
							case 3:
								methods.InterviewSeewithFile()
							}

						}

					}

				}

			}

		case 2:
			for sikl4 {

				fmt.Println("0->Exit\t1->Register\t2->Login")
				var tanlov4 int
				fmt.Scan(&tanlov4)
				switch tanlov4 {
				case 0:
					sikl4 = false
				case 1:
					//  manager uchun registratsiya json faylga yozsin registerni qilsih kerak ++
					methods.RegisterManager()
				case 2:
					// var email string  // login qilib yuboring bobur  managerga
					// var passwors string
					var p, email string
					fmt.Printf("Password kiritng: ")
					fmt.Scan(&p)
					fmt.Printf("Email kiriting: ")
					fmt.Scan(&email)

					r := methods.ManagerReadwithFile()
					
					for _, k := range r {
						if k.Password == p && k.Email == email {
							for sikl5 {
								fmt.Println("0->Exit\t1->Company\t2->Create Company\t3->Delete Company\t4->Update Company")
								var tanlov4 int
								switch tanlov4 {
								case 0:
									sikl5 = false
								case 1:

									for sikl6 {
										fmt.Println("0->Exit\t1->Vacancy\t2->Create Vacancy\t3->Delete Vacancy\t4->Update Vacancy")
										var tanlov5 int
										fmt.Scan(&tanlov5)
										switch tanlov5 {
										case 0:
											sikl6 = false
										case 1:
											// vacancy
											for sikl7 {
												fmt.Println("0->Exit\t1->InterView\t2->Create InterView\t3->Delete InterView\t4->Update InterView")
												var tanlov6 int
												fmt.Scan(&tanlov6)
												switch tanlov6 {
												case 0:
													sikl6 = false
												case 1:
													// interview -----------------------
													var id int
													fmt.Printf("Vakansiya id sini kiriting: ")
													fmt.Scan(&id)
													methodInterview.ReadInterview(id)
												case 2:
													// create interview
													interview1 := methods.Interview{}
													// var id int
													// var t, p string
													fmt.Printf("Interviyu idsini kiriting: ")
													fmt.Scan(&interview1.Id)
													fmt.Printf("Interviyu vaqtini kiriting: ")
													fmt.Scan(&interview1.Time)
													fmt.Printf("Interviyu joyini kiriting: ")
													fmt.Scan(&interview1.Place)

													// fmt.Println(interview1)
													methodInterview.CreateInterview(interview1.Id, interview1)
												case 3:
													// delete interview
													var id int
													fmt.Printf("O`chirmoqchi bo`lga interviyuning idsini kiriting: ")
													fmt.Scan(&id)
													methodInterview.DeleteInterview(id)
												case 4:
													//update interview
													inter := methods.Interview{}
													// var id int
													// var t, p string
													fmt.Printf("Interview idsini kiriting: ")
													fmt.Scan(&inter.Id)
													fmt.Printf("Interviyu vaqtini kiriting: ")
													fmt.Scan(&inter.Time)
													fmt.Printf("Interviyu joyini kiriting: ")
													fmt.Scan(&inter.Place)
													a := methodInterview.UpdateInterview(inter)
													fmt.Println(a)
												}

											}
										case 2:
											// create vacancy
											v := methods.Vacancy{}
											fmt.Printf("vakansiyaga comment yozing: ")
											fmt.Scan(&v.Comments)
											fmt.Printf("Vakansiya sonini kiriting: ")
											fmt.Scan(&v.Count)
											fmt.Printf("Vakansiya qeuid sini kiriting: ")
											fmt.Scan(&v.QueId)
											fmt.Printf("Id kiriting: ")
											fmt.Scan(&v.Id)

											methodVacansi.CreateVacansy(&v)
										case 3:
											// delete vacancies
											var id int
											fmt.Printf("Meniger id sini kiriting: ")
											fmt.Scan(&id)
											a := methodVacansi.DeleteVacansy(id)
											fmt.Println(a)
										case 4:
											// update vacancie
											vac := methods.Vacancy{}
											var qid int
											fmt.Printf("Meniger idsiini kiriting: ")
											fmt.Scan(&qid)

											fmt.Printf("Vakansiya queidsini kiriting: ")
											fmt.Scan(&vac.QueId)
											fmt.Printf("Vakansiya sonini kiriting: ")
											fmt.Scan(&vac.Count)
											fmt.Printf("Vakansiyaga komment yozing: ")
											fmt.Scan(&vac.Comments)
											fmt.Printf("Vakansiya id kiriting: ")
											fmt.Scan(&vac.Id)

											a := methodVacansi.UpdateVancy(qid, vac)
											fmt.Println(a)

										}

									}
								case 2:
									// create company
									comp := methods.Company{}
									fmt.Printf("Kompaniya nomini kiriting: ")
									fmt.Scan(&comp.Name)
									fmt.Printf("Kompaniya id sinikiriting: ")
									fmt.Scan(&comp.Id)
									fmt.Printf("Kompaniya adresini kiriting: ")
									fmt.Scan(&comp.Address)
									fmt.Printf("Kompaniya telefon raqamini kiriting: ")
									fmt.Scan(&comp.Phone)
									methodsCompany.CreateCompany(comp)

								case 3:
									// delete company
									var id int
									fmt.Printf("Kompaniya id sinini kiriting: ")
									fmt.Scan(&id)									
									methodsCompany.DeleteCompany(id)
								case 4:
									// update company
									var id int
									fmt.Printf("Kompaniya id sinini kiriting: ")
									fmt.Scan(&id)
									methodsCompany.UpdateCompany(id)	
								}
							}
						}
					}
					// forni oxiri
				}
			}

		}

	}

}
