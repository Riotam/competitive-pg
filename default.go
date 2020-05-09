package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputStr := getNextLine()
	fmt.Println(checker(inputStr))
}

// use here for unit tests
func checker(inputStr string) string {

	// Your code here
	res := "hoge"

	return fmt.Sprintf("%d", res)
}

func getNextLine() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}
