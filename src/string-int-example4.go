package main

import (
    "fmt"
    "reflect"
    "strconv"
)

func main() {

    fmt.Println("\nConvert String into Integer Data type")
    str3 := "AppDividend"
    fmt.Println("Before String Value: ", str3)
    fmt.Println("Before converting Datatype:", reflect.TypeOf(str3))
    intStr, _ := strconv.ParseInt(str3, 10, 64)
    fmt.Println("After Integer Value: ", intStr)
    fmt.Println("After converting Datatype:", reflect.TypeOf(intStr))
}
