package main
import "github.com/01-edu/z01"
import "os"

func StrLen(arr string) int {
	count := 0
	for j := range arr {
		count = j + 1
	}
	return count
}

func CheckVowel(r byte) bool {
	if r == 'e' || r == 'u' || r == 'i' || r == 'o' || r == 'a' || r == 'E' || r == 'U' || r == 'I' || r == 'O' || r == 'A'{
		return true
	} else {
		return false
	}	
}

func main() {
	slice := os.Args[1:]
	str := ""
	for _, l := range slice {
		str += l
		str += " "
	}
	lenstr := StrLen(str)
	numvow := 0
	for i := 0; i < lenstr; i++ {
		if CheckVowel(str[i]) {
			numvow++
		}
	} 
	arr := make([]byte, numvow)
	numvow = 0
	for i := 0; i < lenstr; i++ {
		if CheckVowel(str[i]) {
			arr[numvow] = str[i]
			numvow++
		}
	}
	for i := 0; i < lenstr; i++ {
		if CheckVowel(str[i]) {
			z01.PrintRune(rune(arr[numvow-1]))
			numvow--
		} else {
			z01.PrintRune(rune(str[i]))
		}
	}
	z01.PrintRune('\n')
}
