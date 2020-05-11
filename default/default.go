package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	number1, number2, number3 := getNextInt(), getNextInt(), getNextInt()
	fmt.Println(number1, number2, number3)

	input := "hoge"
	fmt.Println(checker(input))
}

// use here for unit tests
func checker(inputStr string) string {

	// Your code here

	return "res"
}

func getNextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
