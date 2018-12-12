package main

import  "fmt"

// make our own type
type monkey int
var m monkey
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

    m = 456
    fmt.Println("monkey => ",m)
    fmt.Printf("%T\n",m)
    // cannot use m (type monkey) as type int in assignment
    // x = m
    // but convert it to underlying type
    x = int(m)
    fmt.Println("monkey > int => ",x)
    fmt.Printf("%T\n",x)
}