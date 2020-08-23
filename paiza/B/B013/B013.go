// time.Time型での時刻計算
// 日時に分を加減
// 時刻比較

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	inputList := paizaSequenceGetsAsInt()

	fmt.Println(checker(inputList))

}

func checker(inputList [][]int) string {

	limit := time.Date(2043, 8, 21, 8, 59, 0, 0, time.Local)
	commutingTime := 0

	for _, input1 := range inputList[0] {
		commutingTime += input1
	}

	res := ""

	for i := 2; i < len(inputList); i++ {

		boardingTime := time.Date(2043, 8, 21, inputList[i][0], inputList[i][1], 0, 0, time.Local)
		ImOffTime := boardingTime.Add(-1 * time.Duration(inputList[0][0]) * time.Minute)

		ComeToWork := ImOffTime.Add(time.Duration(commutingTime) * time.Minute)

		if ComeToWork.Before(limit) {
			res = ImOffTime.Format("03:04")
		} else {
			break
		}

	}

	return res
}

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
