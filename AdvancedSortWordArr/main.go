package main
import "fmt"
import "strings"

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}
func Compare2(b, a string) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	for i := 0; i < len(array) - 1; i++ {
		temp := ""
		for j := len(array) - 1; j > i; j-- {
			if Compare(array[i],array[j]) > 0 {
				temp = array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
	}
}

func AdvancedSortWordArr2(array []string, f func(a, b string) int) {
	for i := 0; i < len(array) - 1; i++ {
		temp := ""
		for j := len(array) - 1; j > i; j-- {
			if Compare2(array[i],array[j]) > 0 {
				temp = array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
	}
}

func main() {

    result := []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", "five", "fingers", "of", "each", "hand"}
    AdvancedSortWordArr(result, strings.Compare)
    fmt.Println(result)
    fmt.Println()
    result = []string{"The", "word", "digital", "comesfrom", "\"digits\"", "or", "fingers"}
    AdvancedSortWordArr(result, strings.Compare)
    fmt.Println(result)
    fmt.Println()
    result = []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
    AdvancedSortWordArr(result, strings.Compare)
    fmt.Println(result)
    fmt.Println()
    result = []string{"The", "computing", "consisted", "device", "each", "earliest", "fingers", "five", "hand", "of", "of", "the", "undoubtedly"}
    AdvancedSortWordArr2(result, strings.Compare)
    fmt.Println(result)
    fmt.Println()
    result = []string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"}
    AdvancedSortWordArr2(result, strings.Compare)
    fmt.Println(result)
    fmt.Println()
    result = []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}
    AdvancedSortWordArr2(result, strings.Compare)
    fmt.Println(result)
}
