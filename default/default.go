package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	count, line := paizaGets()

	resList := checker(count, line)
	for _, res := range resList {
		fmt.Printf("%v\n", res)
	}

}

func checker(count int, line []string) []string {

	resList := []string{}
	for _, l := range line {

		lInt, _ := strconv.Atoi(l)

		divisorList := []int{}
		for a := 1; a*a <= lInt; a++ {
			if lInt%a != 0 {
				continue
			}

			b := lInt / a

			divisorList = append(divisorList, a)
			divisorList = append(divisorList, b)
		}

		divisorList = sliceUniqueInt(divisorList)

		divisorSum := 0
		for _, divisor := range divisorList {
			if divisor == lInt {
				continue
			}
			divisorSum += divisor
		}

		if divisorSum == lInt {
			resList = append(resList, "perfect")
		} else if lInt-divisorSum == 1 {
			resList = append(resList, "nearly")
		} else {
			resList = append(resList, "neither")
		}
	}

	return resList
}

func sliceUniqueInt(target []int) (unique []int) {
	m := map[int]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

//  paizaGets は文字列一括取得・最初行数、以下文字列
func paizaGets() (count int, line []string) {
	scanner := bufio.NewScanner(os.Stdin)

	n := -1
	for scanner.Scan() {
		str := scanner.Text()
		if n == -1 {
			count, _ = strconv.Atoi(strings.TrimSpace(str))
		} else {
			line = append(line, strings.TrimSpace(str))
		}
		n++

		if n >= count {
			break
		}
	}

	return
}
