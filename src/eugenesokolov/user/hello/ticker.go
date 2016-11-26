package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case t := <-tick:
			fmt.Printf("tick! at %v\n", t)
		case b := <-boom:
			fmt.Printf("BOOM! at %v\n", b)
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
