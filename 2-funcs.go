package main

import "fmt"

func add(x int, y int) int {
    return x+y
}


func subtract(x int, y int) int {
    return x-y
}

func main() {
    fmt.Printf('Add 42 and 13: ' + string(add(42,13)))
    fmt.Printf('Subtract 13 from 42: '+string(subtract(42,13)))
}