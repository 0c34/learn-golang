package main

import "fmt"

func split(sum int)(x, y int){
	x = sum * 4 / 9
	y = sum - x
	return 
	}

func main(){
	x,y := split(10);
	fmt.Println(x)
	fmt.Println(y)
	}
