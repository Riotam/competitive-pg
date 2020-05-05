package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {

	var secondLine []int
	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i <= 2 && scanner.Scan(); i++ {
		switch i {
		case 2:
			secondLine = getNumbersByStr(scanner.Text())
		}
	}

	minCount := 0
	count := 0

	for i, num := range secondLine {

		for {
			if num%2 != 0 {
				break
			}

			num = num / 2
			count++
		}

		if i == 0 {
			minCount = count
		} else if minCount > count {
			minCount = count
		}

		count = 0

	}

	fmt.Println(minCount)

}
