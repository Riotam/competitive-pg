package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	number := getNextInt()

	fmt.Println(checker(number))
}

func checker(inputInt int) int {

	if inputInt < 0 {
		inputInt *= -1
	}

	return inputInt
}

func getNextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
