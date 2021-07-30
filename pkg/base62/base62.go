package base62

import "math"

const (
	chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base  = len(chars)
)

func ToBase62(seed int64) string {
	var base62String string

	if seed == 0 {
		return string(chars[0])
	}

	for seed > 0 {
		rem := seed % int64(base)
		base62String = string(chars[rem]) + base62String
		seed /= int64(base)
	}

	return base62String
}

func FromBase62(base62String string) int64 {
	strLen := len(base62String)
	var result int64
	for i := 0; i < strLen; i++ {
		val := indexOf(chars, rune(base62String[i]))
		basePowExp := int(math.Pow(float64(base), (float64(strLen) - float64(i+1))))
		result += int64(val * basePowExp)
	}
	return result
}

func indexOf(str string, s rune) int {
	for i, c := range str {
		if c == s {
			return i
		}
	}

	return -1
}
