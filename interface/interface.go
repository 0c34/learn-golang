
package main 

import f "fmt"

type Manusia struct{
	name	string
	age		int
	phone 	string
}

type Pelajar struct{
	Manusia
	sekolah		string
	loan		float32
}

type Karyawan struct{
	Manusia
	company 	string
	uang		float32
}

func (m Manusia) SayHay(){
	f.Printf("Hy nama saya %s telpon saya di nomer ini %s\n", m.name, m.phone)
}
func (m Manusia) Sing(lirik string) {
	f.Println("la la la la", lirik)
}

func (e Karyawan) SayHay(){
	f.Printf("Halo nama saya %s umur saya %d saya bekerja di %s nomer telpon saya %s", e.name, e.age, e.company, e.phone)
}

type Men interface{
	SayHay()
	Sing(lirik string)
}

func main() {

	mike 	:= Pelajar{Manusia{"Mike",21,"082123321123"},"MIT", 0.00}
	paul 	:= Pelajar{Manusia{"Paul",18,"082123321123"},"Harvard", 100}
	sam 	:= Pelajar{Manusia{"Sam",36,"082123321123"},"Golang Inc", 1000}
	tom 	:= Pelajar{Manusia{"Tom",36,"082123321123"},"Tokopedia", 5000}

	var i Men

	i = mike
	f.Println("This is Mike, a student")
	i.SayHay()
	i.Sing("Korea")

	i = tom
	f.Println("This is Tom an Employe")
	i.SayHay()
	i.Sing("Dia")

	f.Println("slice")
	x := make([]Men,3)
	x[0],x[1],x[2] = paul,sam,mike
	for _, value := range x{
		value.SayHay()
	}
}