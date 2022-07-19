package toolx

import "strings"

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

func RenderString(src string, m map[string]string) string {
	for k, v := range m {
		var oldStr strings.Builder
		oldStr.WriteString("${")
		oldStr.WriteString(k)
		oldStr.WriteString("}")
		src = strings.Replace(src, oldStr.String(), v, -1)
	}

	if m["prefix"] != "" {
		src = m["prefix"] + src
	}
	if m["suffix"] != "" {
		src = src + m["suffix"]
	}

	return src
}
