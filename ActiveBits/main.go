package main
import "fmt"

func ActiveBits(n int) uint {
	var count uint
	for true {
		if n != 0 {
			if n % 2 == 1 {
				count++
			}
			n = n / 2
		} else {
			break
		}
	}
	return count
}

func main() {
	nbits := ActiveBits(456546)
	fmt.Println(nbits)
}
