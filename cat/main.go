package main

import "os"
import "github.com/01-edu/z01"
import "io/ioutil"
import "io"

func ArgLen(arr []string) int {
	count := 0
	for i := range arr {
		count = i + 1
	}
	return count
}

func PrintText(str string) {
	for _, l := range str {
		z01.PrintRune(rune(l))
	}
}

func main() {
	arr := os.Args[1:]
	lenarg := ArgLen(arr)
	if lenarg < 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			return
		}
	} else {
		for _, word := range arr {
			f, err := ioutil.ReadFile(word)
			if err != nil {
				PrintText("open ")
				PrintText(word)
				PrintText(": no such file or directory")
				z01.PrintRune('\n')
			}
			PrintText(string(f))
		}
	}

}
