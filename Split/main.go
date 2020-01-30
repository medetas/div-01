package main
import "fmt"

func StrLen(str string) int {
	count := 0
	for i := range str {
		count = i + 1
	}
	return count
}

func GetNT(stringg string, arr []string, n int, charset string, setlen int) {
	str := ""
	for i := range stringg {
		if StrLen(stringg[i:]) >= setlen && stringg[i:i+setlen] == charset {
			if str != "" {
				arr[n] = str
				GetNT(stringg[i+setlen:], arr, n+1, charset, setlen)
				return
			}
			arr[n] = ""
			GetNT(stringg[i+setlen:], arr, n+1, charset, setlen)
			//arr[n] = ""
			return
		}else{
			str += string(stringg[i])
		} 
	}
	if str != "" {
		arr[n] = str
	}
}

func GetNTSize(stringg string, n int, size *int, charset string, setlen int) {
	str := ""
	for i := range stringg {
		if StrLen(stringg[i:]) >= setlen && stringg[i:i+setlen] == charset {
			if str != "" {
				*size++
				GetNTSize(stringg[i+setlen:], n+1, size, charset, setlen)
				return
			}
			*size++
			GetNTSize(stringg[i+setlen:], n+1, size, charset, setlen)
			//*size++
			return
		}else{
			str += string(stringg[i])
		} 
	}
	if str != "" {
		*size++
	}
}

func Split(str string, charset string) []string {
	size := 0
	GetNTSize(str, 0, &size, charset, StrLen(charset))
	arr := make([]string, size)
	GetNT(str, arr, 0, charset, StrLen(charset))
	return arr
}

func main() {
	str := ""
	str = "|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,"
	fmt.Println(Split(str, "|=choumi=|"))

str = "!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,"
	fmt.Println(Split(str, "!==!"))

str = "AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,"
	fmt.Println(Split(str, "AFJ"))

str = "<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,"
	fmt.Println(Split(str, "<<==123==>>"))

}
