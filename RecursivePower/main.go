package main
import "fmt"

func RecursivePower(nb int, power int) int {
	if power > 0 {
		return nb * RecursivePower(nb, power - 1)
	} else if power < 0 {
		return 0
	} else {
		return 1
	}		 
}

func main() {
	fmt.Println(RecursivePower(-7, -2))
	fmt.Println(RecursivePower(-8, -7))
	fmt.Println(RecursivePower(4, 8))
	fmt.Println(RecursivePower(1, 3))
	fmt.Println(RecursivePower(-1, 1))
	fmt.Println(RecursivePower(-6, 5))
}
