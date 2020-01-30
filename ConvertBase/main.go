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

func PrintNbrBase(nbr int, base string) string {
	slice := []rune(base)
	num := 0
	bl := true
	str := "VN"
	str2 := ""
	j := 0
	q := 2
	for _, l1 := range slice {
		k := 0
		num++
		for _, l2 := range slice {
			if l1 == l2 || l2 == '-' || l2 == '+' {
				k++
			}
		}
		if k > 1 {
			bl = false
		}
	}
	if bl && num > 1 {
		str = ""
		q = 0
		if nbr < 0 {
			str2 = "-"
			nbr = nbr * -1
		} else if nbr == 0 {
			str2 = string(slice[0])
		}
		for i := nbr; i > 0; i = i / num {
			j = i % num
			q++
			str = str + string(slice[j])
		}
	}
	for z := q - 1; z >= 0; z-- {
		str2 = str2 + string(str[z]) 
	}

	return str2
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	num := AtoiBase(nbr, baseFrom)
	res := PrintNbrBase(num, baseTo)
	return res
	
}

func main() {
	result := ""
	result = ConvertBase("4506C", "0123456789ABCDEF", "choumi")
	fmt.Println(result)
	result = ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF")
	fmt.Println(result)
	result = ConvertBase("256850", "0123456789", "01")
	fmt.Println(result)
	result = ConvertBase("uuhoumo", "choumi", "Zone01")
	fmt.Println(result)
	result = ConvertBase("683241", "0123456789", "0123456789")
	fmt.Println(result)

}
