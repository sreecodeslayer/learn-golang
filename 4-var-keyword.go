package main

import "fmt"

var y = 100

func main() {
    x := 41
    fmt.Println(x)
    
    fmt.Println("From main() => ",y)
    foo()

    y += 300
    
    foo()
    fmt.Println("From main() => ",y)
}

func foo() {
    fmt.Println("From foo() => ",y)
}