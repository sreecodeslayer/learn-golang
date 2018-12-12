package main

import "fmt"

var y = 100

var z int
func main() {
    x := 41
    fmt.Println(x)
    
    fmt.Println("From main() => ",y)
    foo()

    y += 300
    
    foo()
    fmt.Println("From main() => ",y)

    bar(z)
    fmt.Println("z From main() => ",z)
}

func foo() {
    fmt.Println("From foo() => ",y)
}

func bar(z int) {
    fmt.Println("z From bar(z) => ",z)
    z+=12
    fmt.Println("Increment z From bar(z) => ",z)
}