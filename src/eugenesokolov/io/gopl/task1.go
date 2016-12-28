package main

import "fmt"
import "strings"
import "os"
import "time"

func main() {
    start:=time.Now()
	for i:=0;i<1000;i++ {
        printArgsWithManulStringBuild()
    }
    fmt.Printf("%vms elapsed\n", time.Since(start).Seconds())
}

func printArgsWithStringJoin() {
    fmt.Println(strings.Join(os.Args[1:], " "))
}

func printArgsWithManulStringBuild() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }	
    fmt.Println(s)
}
