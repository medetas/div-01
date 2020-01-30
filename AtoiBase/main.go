package main
import "fmt"


func StrLen(stringg string) int {
	count := 0	
	for i := range stringg {
		count = i + 1
	}
	return count
}

func StrCheck(stringg string) bool {
	slice := []rune(stringg)
	bl := true
	for _, l1 := range slice {
		k := 0
		for _, l2 := range slice {
			if l1 == l2 || l2 == '-' || l2 == '+' {
				k++
			}
		}
		if k > 1 {
			bl = false
		}
	}
	return bl
}

func LetCheck(str1 rune, str2 string) int {
	count := 0
	res := 0	
	for _,l2 := range str2 {
		if str1 == l2 {
			res = count
			break
		} 
		count++
	}
	return res
}
func AtoiBase(s string, base string) int {
	str1 := StrCheck(base)
	num1 := StrLen(base)
	num2 := StrLen(s)
	res := 0
	if str1 && num1 > 1 && num2 > 1 {
		for _, l := range s {
			res = res * num1 + LetCheck(l, base)
		}
	}
	return res


}

func main() {
	fmt.Println(AtoiBase("bcbbbbaab", "abc"))
	fmt.Println(AtoiBase("0001", "01"))
	fmt.Println(AtoiBase("00", "01"))
	fmt.Println(AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"))
	fmt.Println(AtoiBase("AAho?Ao", "WhoAmI?"))
	fmt.Println(AtoiBase("thisinputshouldnotmatter", "abca"))
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
