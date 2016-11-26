package main

import (
	"fmt"
)

func main() {
    a := []int{1, 1, 2, 4}
    сума := 0
    добуток := 1

    for _, v := range a {
        сума += v
        добуток *= v
    }
    fmt.Printf("Cума = %v; Добуток = %v\n", сума, добуток)
}
