package conf

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"jw.lib/toolx"
)

var csm = sync.Map{}

func init() {
	data, err := os.ReadFile("../config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	buf := strings.Split(string(data), "\n")

	for n := range buf {
		buf[n] = strings.TrimSuffix(buf[n], ":")
	}

	for i := len(buf) - 1; i >= 0; i-- {
		str := toolx.TrimAllSpace(buf[i])

		cols := strings.Split(str, ":")
		if len(cols) > 1 {
			pwd := yamlPwd(buf[:i])
			csm.Store(pwd+"."+cols[0], cols[1])

		}
	}
}

func GetYaml(s string) string {
	v, ok := csm.Load(s)
	if ok {
		return v.(string)
	}
	return ""
}

func yamlPwd(elems []string) string {
	per := 999
	var pwd []string
	for i := len(elems) - 1; i >= 0; i-- {
		n := toolx.FirstNotEmptyIndex(elems[i])
		if n < per {
			var temp []string
			pwd = append(temp, elems[i])
			pwd = append(temp, pwd...)
		}
		if n == 0 {
			break
		}
	}
	return strings.Join(pwd, ".")
}
