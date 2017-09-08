package main 
import (
	f "fmt"
	"math"
)

type PersegiPanjang struct{
	panjang, lebar float64
}

type Lingkaran struct{
	jariJari float64
}

func (p PersegiPanjang) area() float64{
	return p.panjang * p.lebar
}

func (l Lingkaran) area() float64 {
	return l.jariJari * l.jariJari * math.Pi 
}

func main() {
	p1 := PersegiPanjang{2, 12}
	p2 := PersegiPanjang{4, 9}
	l1 := Lingkaran{10}
	l2 := Lingkaran{25}

	f.Println("Luas persegi panjang 1",p1.area())
	f.Println("Luas persegi panjang 2",p2.area())
	f.Println("Luas persegi lingkaran 1",l1.area())
	f.Println("Luas persegi lingkaran 2",l2.area())
}