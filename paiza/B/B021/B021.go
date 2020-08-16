// 正規表現
/*
	あなたは英語の文法チェック・修正システムの作成を担当しています。 このシステムを実現するには、英単語を複数形に変換する機能が必要です。

	単語の複数形への変換は、次のルールに従い行われます。

	・末尾が s, sh, ch, o, x のいずれかである英単語の末尾に es を付ける
	・末尾が f, fe のいずれかである英単語の末尾の f, fe を除き、末尾に ves を付ける
	・末尾の1文字が y で、末尾から2文字目が a, i, u, e, o のいずれでもない英単語の末尾の y を除き、末尾に ies を付ける
	・上のいずれの条件にも当てはまらない英単語の末尾には s を付ける

	入力された英単語を複数形に変換するプログラムを作成してください。
*/

/*
	入力例1
		3
		dog
		cat
		pig
	出力例1
		dogs
		cats
		pigs

	入力例2
		7
		box
		photo
		axis
		dish
		church
		leaf
		knife
	出力例2
		boxes
		photoes
		axises
		dishes
		churches
		leaves
		knives

	入力例3
		2
		study
		play
	出力例3
		studies
		plays
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	count, line := paizaGets()

	resList := checker(count, line)
	for _, res := range resList {
		fmt.Printf("%v\n", res)
	}

}

func checker(count int, line []string) []string {

	// 語尾が s, sh, ch, o, x -> +es
	reES := regexp.MustCompile("(s|sh|ch|o|x)+$")

	// 語尾が f, fe -> ves
	reVes := regexp.MustCompile("(f|fe)+$")

	// 語尾1つ前が (a, i, u, e, o)ではない + 語尾がy
	reIes := regexp.MustCompile("[^aiueo]y+$")

	// 語尾が y (iesに変更する)
	reIesRes := regexp.MustCompile("y+$")

	resList := []string{}

	for _, l := range line {

		res := ""

		if reES.MatchString(l) {
			res = l + "es"
		} else if reVes.MatchString(l) {
			res = reVes.ReplaceAllString(l, "ves")
		} else if reIes.MatchString(l) {
			res = reIesRes.ReplaceAllString(l, "ies")
		} else {
			res = l + "s"
		}

		resList = append(resList, res)

	}

	return resList
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
