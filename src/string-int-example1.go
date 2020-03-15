package main

import (
    "fmt"
    "strconv"
)
func main() {
    data := "11"
    if sv, err := strconv.Atoi(data); err == nil {
        fmt.Printf("%T, %v", sv, sv)
    }
}
