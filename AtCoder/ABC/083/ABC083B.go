package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	inputStr := getNextLine()
	inputNumbers := getNumbersByStr(inputStr)

	res := 0
	for i := 1; i <= inputNumbers[0]; i++ {

		number := sumDigits(i)

		if inputNumbers[1] <= number && number <= inputNumbers[2] {
			res += i
		}

	}

	fmt.Println(res)

}

func getNumbersByStr(s string) []int {

	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}

	return n

}

func getNextLine() string {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()

}

func sumDigits(number int) int {
	remainder := 0
	sumResult := 0
	for number != 0 {
		remainder = number % 10
		sumResult += remainder
		number = number / 10
	}
	return sumResult
}
