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

	fmt.Println(checker(count, line))

}

func checker(count int, line []string) (int, int) {

	codes := [][]string{}
	for _, l := range line {
		code, _ := splitAndConvertToString(l, " ")
		codes = append(codes, code)
	}

	v1, v2 := 0, 0
	for _, code := range codes {

		if code[0] == "SET" && code[1] == "1" {
			v1, _ = strconv.Atoi(code[2])
		} else if code[0] == "SET" && code[1] == "2" {
			v2, _ = strconv.Atoi(code[2])
		} else if code[0] == "ADD" {
			a, _ := strconv.Atoi(code[1])
			v2 = v1 + a
		} else if code[0] == "SUB" {
			a, _ := strconv.Atoi(code[1])
			v2 = v1 - a
		}
	}

	return v1, v2
}

// splitAndConvertToString はデリミタで分割してstringスライスを取得
func splitAndConvertToString(stringTargeted, delim string) (stringSlices []string, err error) {
	// 分割
	stringSplited := splitWithoutEmpty(stringTargeted, delim)

	// stringスライスに保存
	for i := range stringSplited {
		// var iparam string
		iparam := stringSplited[i]
		if err != nil {
			return
		}
		stringSlices = append(stringSlices, iparam)
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
