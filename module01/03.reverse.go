package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	reversed := ""
	for _, letter := range word {
		reversed = string(letter) + reversed
	}
	return reversed
}
