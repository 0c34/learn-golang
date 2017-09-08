package main 

import (
	f "fmt"
)

const(
	PUTIH = iota
	HITAM
	BIRU
	MERAH
	KUNING
)
type Warna byte

type Kotak struct{
	lebar, tinggi, panjang float64
	warna Warna
}

type DaftarKotak []Kotak

func (k Kotak) volume() float64{
	return k.lebar * k.panjang * k.tinggi
}

func (k *Kotak) SetWarna(w Warna) {
	k.warna = w
}

func (dk DaftarKotak) WarnaTertinggi() Warna {
	v := 0.0
	k := Warna(PUTIH)
	for _, ko := range dk{
		if ko.volume() > v{
			v = ko.volume()
			k = ko.warna
		}

	}
	return k
}
func (dk DaftarKotak) WarnaiKeHitam() {
	for i, _ := range dk {
		dk[i].SetWarna(HITAM)
	}
}
func (w Warna) String() string {
	strings := []string {"PUTIH","HITAM", "BIRU", "MERAH", "KUNING"}
	return strings[w]
}

func main() {
	kotaks := DaftarKotak{
		Kotak{4,4,4,MERAH},
		Kotak{10,10,1, KUNING},
		Kotak{1,1,20, HITAM},
		Kotak{10,10,1, BIRU},
		Kotak{10,30,1,PUTIH},
		Kotak{20,20,20, KUNING},
	}
	for k,v := range kotaks{
		f.Println(k,v)
	}
	f.Printf("jumlah kotak yang ada = %d \n", len(kotaks))
	f.Println("Volume kotak pertama", kotaks[0].volume(), "cm3")
	f.Println("Warna kotak terkahir", kotaks[len(kotaks)-1].warna.String())
	f.Println("Yang paling besar", kotaks.WarnaTertinggi().String())

	f.Println("Warnai semua kotak dengan warna HITAM")
	kotaks.WarnaiKeHitam()
	f.Println("Warna kotak kedua : ", kotaks[1].warna.String())
	f.Println("Sekarang warna terbesar adalah ", kotaks.WarnaTertinggi().String())
}