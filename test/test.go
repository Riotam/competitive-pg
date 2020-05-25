package main

import (
	"fmt"
)

func main() {
	fmt.Println(example(6, "a"))
}

// == ロジックがメインの関数に対してのテスト ==

func example(number int, word string) string {

	if number > 5 {

		return "over five"

	} else if number > 0 {

		return "over zero"

	}
	// doSomething(string) は渡された単語が空文字であればエラーを返す
	if err := doSomething(word); err != nil {
		return "wrong word"
	}

	return "too small"
}

func doSomething(word string) error {

	if word == "" {
		return fmt.Errorf("空文字")
	}

	return nil
}
