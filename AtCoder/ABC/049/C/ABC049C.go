package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO: 未回答

func main() {
	inputStr := getNextLine()
	fmt.Println(checker(inputStr))
}

func checker(inputStr string) string {

	if len(inputStr) < 5 {
		return "NO"
	}

	reverseS := reverseString(inputStr)
	reverseDreamer := reverseString("dreamer")
	reverseDream := reverseString("dream")
	reverseEraser := reverseString("eraser")
	reverseErase := reverseString("erase")

	res := ""
	if reverseS[:6] == reverseDreamer {
		res = strings.Replace(reverseS, reverseDreamer, "", 1)
	} else if reverseS[:5] == reverseDream {
		res = strings.Replace(reverseS, reverseDream, "", 1)
	} else if reverseS[:6] == reverseEraser {
		res = strings.Replace(reverseS, reverseEraser, "", 1)
	} else if reverseS[:5] == reverseErase {
		res = strings.Replace(reverseS, reverseErase, "", 1)
	}

	if res != reverseDreamer && res != reverseDream && res != reverseEraser && res != reverseErase {
		return "NO"
	}

	return "YES"
}

func getNextLine() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

func reverseString(s string) string {

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
