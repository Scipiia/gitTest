package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello git")
	fmt.Println("buy buy git")
	for i := 0; i < 12; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
