package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// getNumbersByStr はスペース区切りの文字列を、数値のスライスに変換して返す。
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

// getNextLine は標準入力された値を1行単位で文字列にして取得する。複数の実行で2行目、3行目と取得可能。
func getNextLine() string {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()

}

// sumDigits は与えられた数値の各桁の和を求めて返す。
func sumDigits(number int) int {
	remainder := 0
	sum := 0

	for number != 0 {
		remainder = number % 10
		sum += remainder
		number = number / 10
	}

	return sum
}
