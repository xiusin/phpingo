package str

import (
	"strings"
)

func Len(str string) int {
	return len(str)
}

func Pos(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func Sub(str string, offset int, length ...int) string {
	rs := []rune(str)
	l := len(rs)
	if offset > l {
		return ""
	}
	if l > 0 {
		return string(rs[offset : offset+length[0]])
	} else {
		return string(rs[offset:])
	}
}

func Rev(str string) string {
	rs := []rune(str)
	leftPos, rightPos := 0, len(rs)-1
	for rightPos > leftPos {
		rs[leftPos], rs[rightPos] = rs[rightPos], rs[leftPos]
		leftPos++
		rightPos--
	}
	return string(rs)
}

func Contains(haystack string, needle string) bool {
	return strings.Contains(haystack, needle)
}

func EndWith(haystack string, needle string) bool {
	return strings.HasSuffix(haystack, needle)
}

func StartWith(haystack string, needle string) bool {
	return strings.HasPrefix(haystack, needle)
}

func ToUpper(str string) string {
	return strings.ToUpper(str)
}

func ToLower(str string) string {
	return strings.ToLower(str)
}

func UcFirst(str string) string {
	if len(str) == 0 {
		return str
	}
	firstRune := rune(str[0])
	if firstRune > 'a' && firstRune < 'z' {
		firstRune = firstRune - 32
		rs := []rune(str)
		rs[0] = firstRune
		return string(rs)
	} else {
		return str
	}
}

func UcWords(str string) string {
	return strings.Title(str)
}

func Val[T ~int | ~bool | ~float32](v T) string {
	return ""
}

// (string $string, int $length, string $pad_string = " ", int $pad_type = STR_PAD_RIGHT)
func Pad(str string, length int, padStr string, pad_type int) string {
	return ""
}

func Repeat(str string, times uint) string {
	return strings.Repeat(str, int(times))
}

func Trim(str string, chars string) string {
	return strings.Trim(str, chars)
}

func RTrim(str string, chars string) string {
	return strings.TrimRight(str, chars)
}

func LTrim(str string, chars string) string {
	return strings.TrimLeft(str, chars)
}

func TrimFunc(str string, f func(rune) bool) string {
	return strings.TrimFunc(str, f)
}

// 洗牌
func Shuffle(str string) string {
	return ""
}
