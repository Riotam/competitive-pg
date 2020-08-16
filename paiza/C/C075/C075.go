package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {

	x, y, line := paizaGetsXYAndNumbers()

	resList := checker(x, y, line)
	for _, res := range resList {
		fmt.Printf("%v %v\n", res[0], res[1])
	}

}

func checker(x, y int, line []int) [][]int {

	charge := x
	point := 0

	resList := make([][]int, y)

	for i, l := range line {

		if l >= point {
			charge -= l
			point += l / 10
		} else {
			point -= l
		}

		resList[i] = append(resList[i], charge)
		resList[i] = append(resList[i], point)
	}

	return resList
}

// paizaGetsXYAndNumbers は文字列一括取得・最初X,Y、以下Y行数値
func paizaGetsXYAndNumbers() (x, y int, line []int) {
	scanner := bufio.NewScanner(os.Stdin)

	first := true
	for scanner.Scan() {
		str := scanner.Text()
		if first {
			xy, _ := splitAndConvertToInt(str, " ")
			x = xy[0]
			y = xy[1]
			first = false
		} else {
			v, _ := strconv.Atoi(str)
			line = append(line, v)
		}
	}

	return
}

// splitAndConvertToInt はデリミタで分割して整数値スライスを取得
func splitAndConvertToInt(stringTargeted, delim string) (intSlices []int, err error) {
	// 分割
	stringSplited := splitWithoutEmpty(stringTargeted, delim)

	// 整数スライスに保存
	for i := range stringSplited {
		var iparam int
		iparam, err = strconv.Atoi(stringSplited[i])
		if err != nil {
			return
		}
		intSlices = append(intSlices, iparam)
	}
	return
}

// splitWithoutEmpty は空白や空文字だけの値を除去したSplit()
func splitWithoutEmpty(stringTargeted string, delim string) (stringReturned []string) {
	stringSplited := strings.Split(stringTargeted, delim)

	for _, str := range stringSplited {
		if str != "" {
			stringReturned = append(stringReturned, str)
		}
	}

	return
}
