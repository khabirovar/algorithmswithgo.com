package module01

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(value string, base int) int {
	dec := 0
	value = Reverse(value)
	power := 1
	for _, num := range value {
		dec += power * unhex(num)
		power *= base
	}
	return dec
}

func unhex(r rune) int {
	if r >= '0' && r <= '9' {
		return int(r - '0')
	}
	return int(r-'A') + 10
}
