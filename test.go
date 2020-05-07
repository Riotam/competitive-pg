package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, s := getNextLine(), getNextLine()

	fmt.Printf("%v\n%v", f, s)
}

func getNextLine() string {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()

}
