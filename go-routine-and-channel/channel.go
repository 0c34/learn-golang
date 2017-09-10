package main

import (
	"fmt"
	"runtime"
)

func main() {
	
	runtime.GOMAXPROCS(2)
	
	var message = make(chan string)

	var sayHai = func(msg string){
		var data = msg
		message <- data
		}
	go sayHai("soykrucil")
	go sayHai("itthi")
	go sayHai("SoyItthi")
	
	msg1 := <-message
	fmt.Println(msg1)
	
	msg2 := <-message
	fmt.Println(msg2)
	
	msg3 := <-message
	fmt.Println(msg3)
}
