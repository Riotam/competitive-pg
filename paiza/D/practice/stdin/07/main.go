package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputList := paizaGetWordList(",")

	for _, input := range inputList {
		fmt.Println(input)
	}
}

// paizaGetWordList  は文字列取得・複数単語
func paizaGetWordList(delim string) (stringReturned []string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := scanner.Text()
	stringReturned = splitWithoutEmpty(str, delim)

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
