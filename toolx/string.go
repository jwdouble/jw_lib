package toolx

func StrTrimAllRune(s string, r rune) string {
	var buf []rune
	for _, v := range s {
		if v == r {
			continue
		}
		buf = append(buf, v)
	}
	return string(buf)
}
