package main

import (
	"fmt"
	"time"
)

var my int=0

func say() {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(my)
		my++
	}
}

func main() {
	go say()
	say()
}
