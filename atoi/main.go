package main

import "fmt"

//import "github.com/01-edu/z01"
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

func Atoi(s string) int {
	ans := 0
	q := ""
	k := 1
	slice := []rune(s)
	for i, l := range slice {
		if i == 0 && l == '-' {
			k = -1
		} else if i == 0 && l == '+' {
			k = 1
		}else if (l >= 48 && l <= 57) {
			q += string(l)
		} else if (l < 48 || l > 57) {
			q = "0"
			break
		} 
	}
	for _, l := range q {
		ans = ans * 10 + ToRune(l)
	}

	return ans * k
}

func main() {
	s := ""
	s2 := "-"
	s3 := "--123"
	s4 := "1"
	s5 := "-3"
	s6 := "8292"
	s7 := "9223372036854775807"
	s8 := "-9223372036854775808"

	n := Atoi(s)
	n2 := Atoi(s2)
	n3 := Atoi(s3)
	n4 := Atoi(s4)
	n5 := Atoi(s5)
	n6 := Atoi(s6)
	n7 := Atoi(s7)
	n8 := Atoi(s8)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)
}
