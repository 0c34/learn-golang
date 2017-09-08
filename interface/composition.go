package main

import (
	"fmt"
)

type User interface {
	PrintNama()
	PrintDetails()
}
type Person struct{
	Nama string
	Asal string
	Email string
}
func (p *Person) PrintNama(){
	fmt.Printf("Nama :%s", p.Nama)
}
func (p *Person) PrintDetails(){
	fmt.Printf("Asal : %s, Email : %s", p.Asal, p.Email) 
}

type Admin struct{
	Person
	Tugas []string
}
func (a *Admin) PrintDetails(){
	a.Person.PrintDetails()
	fmt.Println("tugas :")
	for _,v := range a.Tugas{
		fmt.Println(v)
	}
}
type Member struct{
	Person
	Tugas []string
}
func (m *Member) PrintDetails(){
	m.Person.PrintDetails()
	fmt.Println("tugas :")
	for _,v := range m.Tugas{
		fmt.Println(v)
	}
}
type Team struct{
	Users []User
}
func (t *Team) GetTeamDetails(){
	fmt.Println("Detail Team")
	for _,v := range t.Users{
		v.PrintNama()
		v.PrintDetails()
	}
}
func main() {
	amir := &Admin{
		Person{
			"Amir",
			"Yogyakarta",
			"amir@gmail.com",
		},
		[]string{"Manajemen User", "Manajemen tugas"},
		}
	udin := &Member{
		Person{
			"Udin",
			"Jakarta",
			"udin@gmail.com",
			},
		[]string{"Lihat data", "update profile"},
	}
	
	team := Team{
		[]User{amir, udin},
		}
	team.GetTeamDetails()
}
