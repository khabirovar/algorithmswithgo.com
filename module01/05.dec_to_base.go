package module01

import (
	"fmt"
	"strings"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//	DecToBase(14, 16) => "E"
//	DecToBase(14, 2) => "1110"
func DecToBase(dec, base int) string {
	var number strings.Builder
	for dec >= base {
		rem := dec % base
		number.WriteString(getNumber(rem))
		dec /= base
	}
	number.WriteString(getNumber(dec))

	return Reverse(number.String())
}

func getNumber(num int) string {
	if num >= 10 {
		return string('A' + byte(num - 10))
	} 
	return fmt.Sprint(num)
	
}
