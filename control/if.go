package main

import "fmt"

func perbandingan(a, b int) string {
	if a < b {
		return "a lebih kecil dari b"
		}
		return "b lebih kecil dari a"
	}

func main(){
	fmt.Println(perbandingan(10,5))
	}
