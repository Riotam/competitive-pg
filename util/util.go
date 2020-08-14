package util

import (
	"bufio"
	"fmt"
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

// getSortDesc は与えられたintのスライスを降順にソートして返す
func getSortDesc(slice []int) []int {
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

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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

// splitWithoutEmptyN は空白や空文字だけの値を除去したSplitN()
func splitWithoutEmptyN(stringTargeted string, delim string, n int) (stringReturned []string) {
	stringSplited := strings.SplitN(stringTargeted, delim, n)

	for _, str := range stringSplited {
		if str != "" {
			stringReturned = append(stringReturned, str)
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

// sliceUniqueString はstringスライスの重複処理
func sliceUniqueString(target []string) (unique []string) {
	m := map[string]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

// sliceUniqueInt はintスライスの重複処理
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

/*
	以下、paiza専用
*/

// paizaSequenceGetsAsInt は行指定なしデータ列読み込み
func paizaSequenceGetsAsInt() (line [][]int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()
		data, err := splitAndConvertToInt(str, " ")
		if err != nil || data == nil {
			break
		}
		line = append(line, data)
	}

	return
}

// paizaSequenceGets は行指定なし文字列読み込み
func paizaSequenceGets() (lines [][]string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()
		data := splitWithoutEmpty(str, " ")
		if data == nil {
			break
		}
		lines = append(lines, data)
	}

	return
}

// paizaStrFirstIntAfter は「文字 数値 数値 数値...」の取得
func paizaStrFirstIntAfter(data string, delim string) (head string, body []int) {
	list := splitWithoutEmptyN(data, delim, 2)
	head = list[0]
	body, _ = splitAndConvertToInt(list[1], delim)
	return
}

// paizaGetsXY は文字列一括取得・最初X,Y、以下Y行文字列
func paizaGetsXY() (x, y int, line []string) {
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
			line = append(line, strings.TrimSpace(str))
		}
	}

	return
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
		n += 1

		if n >= count {
			break
		}
	}

	return
}

//  paizaGetNums は数字列1行取得・1整数以上対応
func paizaGetNums() (intReturned []int) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := scanner.Text()
	intReturned, _ = splitAndConvertToInt(str, " ")

	return
}

// paizaGetWordList  は文字列取得・複数単語
func paizaGetWordList() (stringReturned []string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := scanner.Text()
	stringReturned = splitWithoutEmpty(str, " ")

	return
}

//  paizaGetWord 文字列取得・1単語
func paizaGetWord() (stringReturned string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringReturned = scanner.Text()

	return
}

///////////////////////////////////
// 使い方サンプル
///////////////////////////////////

// TestPaizaGetWord 1行・文字・1単語
func TestPaizaGetWord() {
	line := paizaGetWord()
	fmt.Println(line)
}

// TestPaizaGetWordList 1行・文字・複数単語
func TestPaizaGetWordList() {
	list := paizaGetWordList()
	for _, str := range list {
		fmt.Println(str)
	}
}

// TestPaizaGetNum 1行・数値・空白区切り対応
func TestPaizaGetNum() {
	nums := paizaGetNums()
	fmt.Println(nums)
	// 数値であることを示すため
	if len(nums) >= 2 {
		fmt.Printf("%d+%d=%d\n", nums[0], nums[1], nums[0]+nums[1])
	} else {
		fmt.Printf("%d*2=%d\n", nums[0], nums[0]*2)
	}
}

// TestPaizaGets 1行目.......継続行数
// TestPaizaGets 2行目以降...文字列(空白区切り)
func TestPaizaGets() {
	count, line := paizaGets()
	fmt.Printf("count=%d\n", count)
	for _, d := range line {
		strs := splitWithoutEmpty(d, " ")
		fmt.Println(strs)
	}
}

// TestPaizaStrFirstIntAfter 1行目.......継続行数
// TestPaizaStrFirstIntAfter 2行目以降...文字、整数値列(空白区切り)
func TestPaizaStrFirstIntAfter() {
	count, line := paizaGets()
	fmt.Printf("count=%d\n", count)
	for _, d := range line {
		head, body := paizaStrFirstIntAfter(d, " ")
		fmt.Println(head)
		fmt.Println(body)
		fmt.Printf("%d+%d=%d\n", body[0], body[1], body[0]+body[1]) // 数値であることを示すため
	}
}

// TestpaizaGetsXY 1行目.......2整数(x yなど)
// TestpaizaGetsXY 2行目以降...文字列(空白区切り)
func TestpaizaGetsXY() {
	x, y, line := paizaGetsXY()
	fmt.Printf("(x,y)=(%d,%d)\n", x, y)
	for _, d := range line {
		strs := splitWithoutEmpty(d, " ")
		fmt.Println(strs)
	}
}

// TestPaizaSequenceGetsAsInt 改行のみ続くまで数値列入力
func TestPaizaSequenceGetsAsInt() {
	line := paizaSequenceGetsAsInt()
	for _, d := range line {
		fmt.Println(d)
		fmt.Printf("%d+%d+%d=%d\n", d[0], d[1], d[2], d[0]+d[1]+d[2])
	}
}

// TestPaizaSequenceGets 改行のみ続くまで文字列入力
func TestPaizaSequenceGets() {
	line := paizaSequenceGets()
	for i, d := range line {
		fmt.Printf("-- %d --\n", i)
		for _, s := range d {
			fmt.Printf("%s\n", s)
		}
	}
}
