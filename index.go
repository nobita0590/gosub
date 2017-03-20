package gosub

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)
func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func ReplaceUTF8Character(input string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, input)
	result = strings.Replace(strings.Trim(strings.ToLower(result)," "),"đ","d",-1)
	return  result
}