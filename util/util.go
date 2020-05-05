package util

import (
	"strconv"
	"strings"
)

// getNumbersByStr はスペース区切りの文字列を、数値のスライスに変換して返す
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
