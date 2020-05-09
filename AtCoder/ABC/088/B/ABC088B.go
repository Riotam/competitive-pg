package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var firstLine int
	fmt.Scanf("%d", &firstLine)

	secondLineStr := getNextLine()
	secondLine := getNumbersByStr(secondLineStr)

	sort.Sort(sort.Reverse(sort.IntSlice(secondLine)))

	var alicePoint, bobPoint int

	for i, s := range secondLine {
		if i%2 == 0 {
			alicePoint += s
		} else {
			bobPoint += s
		}
	}

	res := alicePoint - bobPoint
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
