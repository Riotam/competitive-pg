package main

import (
	"fmt"
	"strings"
)

func main() {

	var s string
	fmt.Scanf("%s", &s)

	slice := strings.Split(s, "")

	count := 0
	for _, v := range slice {
		if "1" == v {
			count++
		}
	}

	fmt.Println(count)

}
