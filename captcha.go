package captcha

import (
	"fmt"
	"strconv"
)

var (
	words = map[int]string{
		0: "ZERO",
		1: "ONE",
		2: "TWO",
		3: "THREE",
		4: "FOUR",
		5: "FIVE",
		6: "SIX",
		7: "SEVEN",
		8: "EIGHT",
		9: "NINE",
	}

	operators = map[int]string{
		1: "+",
		2: "-",
		3: "*",
	}
)

// Create new captcha
func Create(o1, n1, o2, n2 int) string {
	var num1, num2 string

	if o1 == 1 {
		num1 = strconv.Itoa(n1)
		num2, _ = words[n2]
	} else {
		num1, _ = words[n1]
		num2 = strconv.Itoa(n2)
	}

	op, _ := operators[o2]

	return fmt.Sprintf("%s %s %s", num1, op, num2)
}
