package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	var n int
	fmt.Scanf("%d", &n)

	lines := getLines(n)
	numbers := strings2numbers(lines)
	setSortDesc(numbers)

	count, remainder := 0, 0

	for i, n := range numbers {
		if i == 0 || n < remainder {
			count++
			remainder = n
		}
	}

	fmt.Println(count)
}

func getLines(n int) []string {
	var lines []string
	sc := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		sc.Scan()
		t := sc.Text()
		lines = append(lines, t)
	}
	return lines
}

func setSortDesc(slice []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	return slice
}

func strings2numbers(strings []string) []int {
	var numbers []int
	for _, s := range strings {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}
	return numbers
}
