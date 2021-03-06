package strings

import "unicode"

// ToLowerInitial converts a string's first character to lower case.
func ToLowerInitial(s string) string {
	if s == "" {
		return ""
	}
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

// ToUpperInitial converts a string's first character to upper case.
func ToUpperInitial(s string) string {
	if s == "" {
		return ""
	}
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}
