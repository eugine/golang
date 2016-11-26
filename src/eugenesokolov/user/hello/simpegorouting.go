package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:2], c)
	go sum(a[2:4], c)
	go sum(a[4:], c)
	x, y, z := <-c, <-c, <-c // receive from c

	fmt.Println(x, y, z, x+y+z)
}
