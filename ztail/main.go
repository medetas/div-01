package main
import "os"
import "fmt"

func ArgLen(arr []string) int {
	count := 0
	for i := range arr {
		count = i + 1
	}
	return count
}

func StrLen(str string) int {
	count := 0
	for i := range str {
		count = i + 1
	}
	return count
}

func ToRune(r rune) int {
	a := 0
	for true {
		if r <= 48 {
			break
		}
		r--
		a++

	}
	return a
}

func AtoiZTail(s string, sign bool) int {
	if s == "" {
		fmt.Printf("tail: invalid number of bytes: ''\n")
		os.Exit(1)
	}	
	ans := 0
	q := ""
	k := 1
	slice := []rune(s)
	for _, l := range slice {
		if (l >= 48 && l <= 57) {
			q += string(l)
		} else if (l < 48 || l > 57) && sign {
			fmt.Printf("tail: invalid number of bytes: '" + s + "'\n")
			os.Exit(1)
		} else if (l < 48 || l > 57) && !sign {
			fmt.Printf("tail: invalid number of bytes: '+" + s + "'\n")
			os.Exit(1) 
		}
	}
	for _, l := range q {
		ans = ans * 10 + ToRune(l)
	}
	return ans * k
}

func main() {
	args := os.Args[1:]
	var deleteCarr []string
	minus := true
	c := false
	var bytes int
	if ArgLen(args) == 0 {
		os.Exit(1)
	}
	for i := 0; i < ArgLen(args); i++ {
		word := args[i]
		if word == "-c" {
			c = true
			if ArgLen(args) > i + 1 {
				nextstr := args[i+1]
				if nextstr[0] == '-' {
					nextstr = nextstr[1:]
				} else if nextstr[0] == '+' && StrLen(nextstr) == 1 {
					fmt.Printf("tail: invalid number of bytes: ‘+’\n")
					os.Exit(1)
				} else if nextstr[0] == '+' {
					nextstr = nextstr[1:]
					minus = false
				}
				bytes = AtoiZTail(nextstr, minus)
				i++
			} else {
				fmt.Printf("tail: option requires an argument -- 'c'\nTry 'tail --help' for more information.\n")
				os.Exit(1)
			}
		} else if StrLen(word) > 2 && word[:2] == "-c" {
			c = true
			nextstr := word[2:]
			if nextstr[0] == '-' {
				nextstr = nextstr[1:]
			} else if nextstr[0] == '+' && StrLen(nextstr) == 1 {
				fmt.Printf("tail: invalid number of bytes: ‘+’\n")
				os.Exit(1)
			} 
			if nextstr[0] == '+' {
				nextstr = nextstr[1:]
				minus = false
			}
			bytes = AtoiZTail(nextstr, minus)
		} else {
			deleteCarr = append(deleteCarr, word)
		}
	}
	if !c {
		os.Exit(1)
	}

	wroteSmth := false
	for _, filename := range deleteCarr {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("tail: cannot open '" + filename + "' for reading: No such file or directory\n")
			deleteCarr = deleteCarr[1:]
			wroteSmth = true
			continue
		}
		if ArgLen(deleteCarr) > 0 && deleteCarr[0] != filename {
			fmt.Printf("\n")
		}
		if ArgLen(deleteCarr) > 1 || wroteSmth {
			fmt.Printf("==> " + filename + " <==\n") 
		}
		var data []byte
		if minus {
			var i int64
			stat,_ := f.Stat()
			if stat.Size() < int64(bytes) {
				for i = 0; i < stat.Size(); i++ {
					data = append(data, 0)
				}
				f.ReadAt(data, 0)
			} else {
				for i = 0; i < int64(bytes); i++ {
					data = append(data, 0)
				}
				f.ReadAt(data, stat.Size() - int64(bytes))
			}
		} else {
			stat,_ := f.Stat()
			for i := 0; i < int(stat.Size()); i++ {
				data = append(data, 0)
			}
			if bytes == 0 {
				bytes = 1
			}
			f.ReadAt(data, int64(bytes-1))
		}
		fmt.Printf(string(data))
	}
}
