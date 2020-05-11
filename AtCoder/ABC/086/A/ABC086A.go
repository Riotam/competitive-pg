package main

import "fmt"

func main() {

	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(checkEven(a, b))

}

func checkEven(a int, b int) string {

	if a%2 == 0 || b%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}

}
