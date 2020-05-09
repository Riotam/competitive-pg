package util

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

// getSliceIntByString はスペース区切りの文字列を、数値のスライスに変換して返す。
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

// getNextLine は標準入力された値を1行単位で文字列にして取得する。複数の実行で2行目、3行目と取得可能。
// TODO: for文だと1.14では使えない？
func getNextLine() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

// getLines は与えられた数値の回数だけ、標準入力された値を取得する。
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

// getSumDigits は与えられた数値の各桁の和を求めて返す。
func getSumDigits(number int) int {
	remainder := 0
	sum := 0
	for number != 0 {
		remainder = number % 10
		sum += remainder
		number = number / 10
	}
	return sum
}

// setSortDesc は与えられたintのスライスを降順にソートして返す
func setSortDesc(slice []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	return slice
}

func getSliceIntBySliceString(strings []string) []int {
	var numbers []int
	for _, s := range strings {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}
	return numbers
}
