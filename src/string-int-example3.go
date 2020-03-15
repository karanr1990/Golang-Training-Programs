package main

import (
    "fmt"
)

func main() {
    s := "id:00211"
    
    var i int
    
    if _, err := fmt.Sscanf(s, "id:%5d", &i); err == nil {
        fmt.Println(i)
    }
}
