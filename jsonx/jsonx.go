package jsonx

import (
	"encoding/json"
)

func MustMarshal(v any) []byte {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return buf
}
