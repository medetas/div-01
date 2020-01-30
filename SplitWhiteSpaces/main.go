package main
import "fmt"

func GetNT(stringg string, arr []string, n int) {
	str := ""
	for i, l := range stringg {
		if l == '\n' || l == ' ' || l == '\t'{
			if str != "" {
				arr[n] = str
				GetNT(stringg[i+1:], arr, n+1)
				return
			}
			GetNT(stringg[i+1:], arr, n)
			return
		}else{
			str += string(l)
		} 
	}
	if str != "" {
		arr[n] = str
	}
}

func GetNTSize(stringg string, n int, size *int) {
	str := ""
	for i, l := range stringg {
		if l == '\n' || l == ' ' || l == '\t'{
			if str != "" {
				*size++
				GetNTSize(stringg[i+1:], n+1, size)
				return
			}
			GetNTSize(stringg[i+1:], n, size)
			return
		}else{
			str += string(l)
		} 
	}
	if str != "" {
		*size++
	}
}

func SplitWhiteSpaces(str string) []string {
	size := 0
	GetNTSize(str, 0, &size)
	arr := make([]string, size)
	GetNT(str, arr, 0)
	return arr
}

func main() {
	str := ""
	str = "The earliest foundations of what would become computer science predate the invention of the modern digital computer"
	fmt.Println(SplitWhiteSpaces(str))
str = "Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,"
	fmt.Println(SplitWhiteSpaces(str))
str = "aiding in computations such as multiplication and division ."
	fmt.Println(SplitWhiteSpaces(str))
str = "Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."
	fmt.Println(SplitWhiteSpaces(str))
str = "Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"
	fmt.Println(SplitWhiteSpaces(str))
str = "In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"
	fmt.Println(SplitWhiteSpaces(str))
}
