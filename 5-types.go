package main

import  "fmt"

func main() {
    x := 23
    var y string = "this is something unusual for a pythonist"
    raw_string := `I just said, 
    "STATIC typing is alien to me! Duh!"
    `
    // code below is not allowed
    // raw_string := 'I just said, 
    // "STATIC typing is alien to me! Duh!"
    // '
    fmt.Println("Int => ",x)
    fmt.Printf("%T\n",x)

    fmt.Println("String => ",y)
    fmt.Printf("%T\n",y)
    fmt.Println("Raw String => ",raw_string)
    fmt.Printf("%T\n",raw_string)
}