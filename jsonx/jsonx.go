package jsonx

import (
	"encoding/json"
	"log"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(buf []byte, v any) error {
	return json.Unmarshal(buf, v)
}

func MustMarshal(v any) []byte {
	buf, err := json.Marshal(v)
	if err != nil {
		log.Print(v)
		panic(err)
	}

	return buf
}

func MustUnmarshal(buf []byte, v any) {
	err := json.Unmarshal(buf, v)
	if err != nil {
		if buf != nil && v != nil {
			log.Print(buf, "-->", v)
		}

		panic(err)
	}
}
