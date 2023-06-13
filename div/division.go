package div

import "fmt"

var res float64

func Details(num1, num2 float64) float64 {
	if num2 == 0 {
		fmt.Println("division error")
	} else {
		res = num1 / num2
		return res

	}
	return 0
}
