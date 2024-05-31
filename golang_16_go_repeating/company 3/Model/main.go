package main

import (
	"fmt"
	methods "my_project/Model/Methods"
)

type User struct {
	Name string
	Id   int
}

func main() {

	methodInterview := methods.Recruiter{}
	methodVacansi := methods.Recruiter{}
	methodsCompany := methods.Recruiter{}
	var RecuiterId int
	var CompanyId int
	var VacansyId int

	var sikl1 bool
	sikl1 = true
	for sikl1 {
		fmt.Println("0->Exit\t1->User\t2->Recuiter")
		var tanlov1 int
		fmt.Scan(&tanlov1)
		switch tanlov1 {
		case 0:
			sikl1 = false
			fmt.Println("E'tiboringiz uchun Raxmat uchun")
		case 1:
			var sikl2 = true
			for sikl2 {
			B:
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

					for true {
						var id int
						fmt.Printf("Id kiriting: ")
						fmt.Scan(&id)

						e1 := methods.LoginUser(id)
						if e1 == true {
							break
						}
						var tanlov8 int
						fmt.Println("0->Exit\t1->Continue")
						fmt.Scan(&tanlov8)
						if tanlov8 == 0 {
							goto B
							break
						}
					}

					fmt.Println("0->Exit\t1-Company\t2-Vakansiya\t3-Interviyu")
					var tanlov3 int
					fmt.Scan(&tanlov3)
					switch tanlov3 {
					case 0:
						return
					case 1:
						// c1 := []methods.Company{}
						// c1 = methods.CompReadJsonFile()
						// break
						// kompaniylarni ro'yxatini chiqarish kerak
						// for i := 0; i < len(c1); i++ {
						// 	fmt.Println(c1[i])
						// }
						methods.CompanyViewUser()
					case 2:
						methods.VacansySeewithFile()
					case 3:
						methods.InterviewSeewithFile()
					}

					// }

					// }

				}

			}

		case 2:
			var sikl4 = true
			for sikl4 {
			A:
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
					for true {
						var email, password string
						fmt.Printf("email kiriting: ")
						fmt.Scan(&email)
						fmt.Printf("passowrd kiriting: ")
						fmt.Scan(&password)

						RecuiterId = methods.LoginManager(email, password)
						fmt.Println(RecuiterId)
						if RecuiterId != -1 {
							break
						}
						var tanlov7 int
						fmt.Println("0->Exit\t1->Continue")
						fmt.Scan(&tanlov7)
						if tanlov7 == 0 {
							sikl4 = false
							goto A
							break
						}
					}

					var sikl10 = true
					for sikl10 {
						fmt.Println("0->Exit\t1->Company\t2->Create Company\t3->Delete Company\t4->Update Company")
						var tanlov4 int
						fmt.Scan(&tanlov4)
						switch tanlov4 {
						case 0:
							sikl10 = false
						case 1:
							methodsCompany.ReadCompany(RecuiterId)
							var sikl6 = true

							for sikl6 {
								fmt.Println("0->Exit\t1->Vacancy\t2->Create Vacancy\t3->Delete Vacancy\t4->Update Vacancy")
								var tanlov5 int
								fmt.Scan(&tanlov5)
								switch tanlov5 {
								case 0:
									sikl6 = false
								case 1:
									// vacancy
									fmt.Println("Company ID : ")
									fmt.Scan(&CompanyId)
									methodVacansi.ReadVanacy(RecuiterId, CompanyId)
									var sikl7 = true
									for sikl7 {
										fmt.Println("0->Exit\t1->InterView\t2->Create InterView\t3->Delete InterView\t4->Update InterView")
										var tanlov6 int
										fmt.Scan(&tanlov6)

										switch tanlov6 {
										case 0:
											sikl7 = false
										case 1:
											// interview -----------------------
											fmt.Println("Vackansy id : ")
											fmt.Scan(&VacansyId)

											methodInterview.ReadInterview(RecuiterId, CompanyId, VacansyId)
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
											methodInterview.CreateInterview(RecuiterId, CompanyId, VacansyId, interview1)
										case 3:
											// delete interview
											var ineterviewid int
											fmt.Printf("O`chirmoqchi bo`lga interviyuning idsini kiriting: ")
											fmt.Scan(&ineterviewid)
											methodInterview.DeleteInterview(RecuiterId, CompanyId, VacansyId, ineterviewid)
											
										case 4:
											//update interview
											inter := methods.Interview{}
											var interviewId2 int
											fmt.Println("Vakasiya Id: ")
											fmt.Scan(&interviewId2)
											fmt.Printf("Interview idsini kiriting: ")
											fmt.Scan(&inter.Id)
											fmt.Printf("Interviyu vaqtini kiriting: ")
											fmt.Scan(&inter.Time)
											fmt.Printf("Interviyu joyini kiriting: ")
											fmt.Scan(&inter.Place)

											methodInterview.UpdateInterview(RecuiterId, CompanyId, VacansyId, interviewId2, inter)
											
										}

									}
								case 2:
									// create vacancy
									v := methods.Vacancy{}
									fmt.Printf("vakansiyaga comment yozing: ")
									fmt.Scan(&v.Comments)
									fmt.Printf("Vakansiya sonini kiriting: ")
									fmt.Scan(&v.Count)
									fmt.Printf("Id kiriting: ")
									fmt.Scan(&v.Id)

									methodVacansi.CreateVacansy(RecuiterId, CompanyId, v)
									// if a == true {
									// 	fmt.Println("Success")
									// } else {
									// 	fmt.Println("Failed")
									// }
								case 3:
									// delete vacancies
									var vakansy int
									fmt.Printf("Vakansiya id sini kiriting: ")
									fmt.Scan(&vakansy)
									methodVacansi.DeleteVacansy(RecuiterId, CompanyId, vakansy)

								case 4:
									// update vacancie
									vac := methods.Vacancy{}
									var vakansyId int
									fmt.Printf("Meniger idsiini kiriting: ")
									fmt.Scan(&vakansyId)
									fmt.Printf("Vakansiya sonini kiriting: ")
									fmt.Scan(&vac.Count)
									fmt.Printf("Vakansiyaga komment yozing: ")
									fmt.Scan(&vac.Comments)
									fmt.Printf("Vakansiya id kiriting: ")
									fmt.Scan(&vac.Id)

									methodVacansi.UpdateVancy(RecuiterId, CompanyId, vakansyId, vac)

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
							methodsCompany.CreateCompany(comp, RecuiterId)
							// e3:=methodsCompany.CreateCompany(comp,index)
							// if e3==true{
							// 	fmt.Println("Success")
							// }else{
							// 	fmt.Println("Failed")
							// }

						case 3:
							// delete company
							var companyid int
							fmt.Printf("Kompaniya id sinini kiriting: ")
							fmt.Scan(&companyid)
							methodsCompany.DeleteCompany(companyid, RecuiterId)
						case 4:
							// update company
							var companyid int
							fmt.Printf("Kompaniya id sinini kiriting: ")
							fmt.Scan(&companyid)
							methodsCompany.UpdateCompany(companyid, RecuiterId)
						}
					}
					break
				}
			}
			// forni oxiri
		}
	}

}

// }

// }
