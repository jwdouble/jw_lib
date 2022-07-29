package jsonx

import (
	"testing"
)

func Test_unmarshal(t *testing.T) {
	type info struct {
		App string `json:"app,omitempty"`
	}
	i := info{}
	err := Unmarshal([]byte(`{"app": ""}`), &i)
	if err != nil {
		t.Error(err)
	}
	t.Log(i)
}
