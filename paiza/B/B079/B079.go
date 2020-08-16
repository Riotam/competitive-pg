/*
	あなたは恋愛相談会社 PAIZA のプログラマーです。今回、あなたは二人の相性占いをするプログラムを作ることにしました。

	あなたが作成するプログラムは、まず、以下の流れで「二人の相性」を求めます。

	1. 相性をチェックする二人の名前を並べた英小文字からなる文字列を入力します。
	2. "a" を 1、"b" を 2、...、"z" を 26 として、文字列を数列に変換します。この数列を A とします。
	3. 数列 A の隣り合う 2 つの数を足して前から順番に並べた新しい数列 A' を作り、これを新たに A とします。このとき、A の要素の大きさが 101 を超えていた場合、その要素から 101 を引きます。
	4. 数列 A の要素数が 1 になるまで 3. の手順を繰り返します。A の要素数が 1 となったとき、残った要素の値を「二人の相性」とします。

	名前の並べ方は 2 通りあります。そこで、あなたは相性占いの結果として、 2 通りの方法で計算した「二人の相性」のうち大きい方を出力するようにプログラムを組むことにしました。

	たとえば、pa さんと iza さんの名前を "paiza" として並べたとき、「二人の相性」は 78 になります。

	16, 1, 9, 26, 1
	↓
	17,10,35,27
	↓
	27, 45, 62
	↓
	72, 6
	↓
	78
*/

/*
	検索キーワード
		再帰処理風
		map
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	input := paizaGetWordListWithDelim(" ")

	fmt.Println(checker(input[0], input[1]))

}

func checker(s, t string) int {

	res1 := getResDivination(s + t)

	res2 := getResDivination(t + s)

	res := 0
	if res1 > res2 {
		res = res1
	} else {
		res = res2
	}

	return res

}

func getResDivination(input string) int {

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	}

	inputSplittedList := splitWithoutEmpty(input, "")

	numList := []int{}
	newNumList := []int{}

	for _, inputSplitted := range inputSplittedList {
		numList = append(numList, m[inputSplitted])
	}

	for {

		for i := 0; i < len(numList)-1; i++ {

			newNum := numList[i] + numList[i+1]

			if newNum >= 101 {
				newNum -= 101
			}

			newNumList = append(newNumList, newNum)

		}

		if len(newNumList) == 1 {
			break
		}

		numList = newNumList
		newNumList = []int{}

	}

	return newNumList[0]

}

// paizaGetWordListWithDelim はデリミタで区切られた複数単語の文字列を取得しスライス化する
func paizaGetWordListWithDelim(delim string) (stringReturned []string) {
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
