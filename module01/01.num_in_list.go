package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	for _, val := range list {
		if num == val {
			return true
		}
	}
	return false
}
