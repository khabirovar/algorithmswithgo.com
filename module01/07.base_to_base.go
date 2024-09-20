package module01

import "strings"

// BaseToBase takes in a number, the base it is currently
// in, and the base you want it to be converted to. It then
// returns a string representing the number in the new base.
//
// Eg:
//
//   BaseToBase("E", 16, 2) => "1110"
//
func BaseToBase(value string, base, newBase int) string {
	//return DecToBase(BaseToDec(value, base), newBase)
	dec := 0
	power := 1
	for _, digit := range Reverse(value) {
		dec += power * unhex(digit)
		power *= base
	}
	var sb strings.Builder
	for dec > 0 {
		rem := dec % newBase
		sb.WriteString(getNumber(rem))
		dec /= newBase
	}
	return Reverse(sb.String())
}
