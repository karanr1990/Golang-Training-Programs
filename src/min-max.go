package main

import "fmt"

func main() {

	arr := []int{2, 5, 7, 4, 9}

	max := arr[0]
	min := arr[0]

	for _, value := range arr {

		if value > max {
			max = value
		}

		if value < min {
			min = value
		}

	}
	fmt.Println(max, min)
}
