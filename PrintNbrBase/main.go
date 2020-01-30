package main

//import "fmt"
import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	slice := []rune(base)
	num := 0
	bl := true
	str := "VN"
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
			z01.PrintRune('-')
			//nbr = nbr * -1
			for i := nbr; i < 0; i = i / num {
				j = i % num
				q++
				str = str + string(slice[j*-1])
			}
		} else if nbr == 0 {
			z01.PrintRune(slice[0])
		} else {
			for i := nbr; i > 0; i = i / num {
				j = i % num
				q++
				//fmt.Println(j)
				str = str + string(slice[j])
			}
		}
	}

	for z := q - 1; z >= 0; z-- {
		z01.PrintRune(rune(str[z]))
	}

}

func main() {
	PrintNbrBase(919617, "01")
	z01.PrintRune('\n')
	PrintNbrBase(753639, "CHOUMIisDAcat!")
	z01.PrintRune('\n')
	PrintNbrBase(-174336, "CHOUMIisDAcat!")
	z01.PrintRune('\n')
	PrintNbrBase(-661165, "1")
	z01.PrintRune('\n')
	PrintNbrBase(-861737, "Zone01Zone01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "-ab")
	z01.PrintRune('\n')
	PrintNbrBase(-9223372036854775808, "0123456789")
	z01.PrintRune('\n')
}
