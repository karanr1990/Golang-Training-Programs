package main

import (
    "fmt"
    "strconv"
)

func main() {
    data := "Golang"
    sv, err := strconv.Atoi(data)
    if err == nil {
        fmt.Printf("%T, %v", sv, sv)
    } else {
        fmt.Println(err)
    }
}
