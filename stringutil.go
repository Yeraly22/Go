package stringutil

// Reverse возвращает обратную к строке строку.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Length возвращает количество символов в строке.
func Length(s string) int {
	return len([]rune(s))
}
