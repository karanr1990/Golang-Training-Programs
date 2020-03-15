package main

import (
    "fmt"
    "reflect"
    "strings"
)

func main() {
    str1 := []string{"Trump", "In", "India", "On", "Feb 25"}
    fmt.Println(str1)
    fmt.Println(reflect.TypeOf(str1))

    str2 := strings.Join(str1, " ")
    fmt.Println(str2)
    fmt.Println(reflect.TypeOf(str2))

    str3 := strings.Join(str1, ", ")
    fmt.Println(str3)
    fmt.Println(reflect.TypeOf(str3))
}
