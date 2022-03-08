package main

import "fmt"

func ArrayChecker(a [10]int) bool {
	for _, element := range a {
		counter := 0
		for _, number := range a {
			if element == number {
				counter += 1
			}
		}
		if counter >= 2 {
			return true
		}
	}
	return false
}

func main() {
	nums := [10]int{1, 42, 3, 4, 5, 4, 7, 8, 9, 10}
	fmt.Println(ArrayChecker(nums))
}
