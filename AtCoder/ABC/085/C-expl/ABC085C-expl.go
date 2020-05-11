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
	fmt.Println(checker(inputStr))
}

func checker(inputStr string) string {
	inputSliceInt := getSliceIntByString(inputStr)

	for x := 0; x <= inputSliceInt[0]; x++ {
		for y := 0; y <= inputSliceInt[0]-x; y++ {

			z := inputSliceInt[0] - x - y
			if x+y+z == inputSliceInt[0] && 10000*x+5000*y+1000*z == inputSliceInt[1] {
				return fmt.Sprintf("%d %d %d", x, y, z)
			}
		}
	}

	return "-1 -1 -1"
}

func getNextLine() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

func getSliceIntByString(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}
