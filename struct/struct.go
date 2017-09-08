package main

import (
	f "fmt"	
)

type orang struct{
	nama string
	umur int
}

func main(){
	var O orang
	O.nama = "cantik"
	O.umur = 17
	f.Println("Nama :", O.nama)
	f.Println("Umur :", O.umur)
}