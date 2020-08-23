package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lineList := paizaSequenceGetsAsInt()

	resList := checker(lineList)
	for _, res := range resList {
		fmt.Printf("%v\n", res)
	}

}

func checker(lineList [][]int) []string {

	// lineList[0][0]個のスライスを作成
	// 各要素ごとにlineList[0][1]個のスライスを作成 -> [][]string
	// 全て.で埋める
	box := initBox(lineList[0])

	// N個目が落ちる

	for i := 1; i < len(lineList); i++ {
		// for _, line := range lineList {
		// lineList[i][2]〜lineList[i][2]+lineList[i][1] で存在していない(!=#)最下層まで落ちる

		// 最下層(反映する一番後ろのインデックス)を取得
		bottomIndex := getBottomLine(box, lineList[i])

		// 上記最下層 に lineList[i][2]〜lineList[i][2]+lineList[i][1] を #に更新
		// resList[lineList[i][0]] まで上記を反映
		box = updateBox(bottomIndex, box, lineList[i])

	}

	res := make([]string, len(box))
	for i, line := range box {
		resLine := strings.Join(line, "")
		res[(len(box) - i - 1)] = resLine
	}

	return res
}

func updateBox(bottomIndex int, box [][]string, line []int) [][]string {

	for j := bottomIndex; j < bottomIndex+line[0]; j++ {

		// updateLine := []string{}
		// for i := 0; i < len(box[0]); i++ {
		// 	updateLine = append(updateLine, ".")
		// }

		for i := line[2]; i < line[2]+line[1]; i++ {
			box[j][i] = "#"
		}

		// box[j] = updateLine
	}
	return box
}

func getBottomLine(box [][]string, target []int) int {

	bottomIndex := 0
	for i, line := range box {

		ok := true
		for j := target[2]; j < (target[2] + target[1]); j++ {
			if line[j] == "#" {
				ok = false
				break
			}
		}

		if ok {
			bottomIndex = i
			break
		}

	}

	return bottomIndex
}

func initBox(line []int) [][]string {

	resLine := []string{}
	for i := 0; i < line[1]; i++ {
		resLine = append(resLine, ".")
	}

	resBox := [][]string{}
	for i := 0; i < line[0]; i++ {
		resBox = append(resBox, resLine)
	}

	return resBox

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
