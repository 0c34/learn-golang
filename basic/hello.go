package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

func main(){
	fmt.Println("hello, World")
	fmt.Println("waktu sekarang adalah", time.Now())
	fmt.Println("Bilangan kesukaan anda", rand.Intn(10))
	fmt.Println("Now you have %g problems.", math.Sqrt(7))
	
}


