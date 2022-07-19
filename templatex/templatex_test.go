package templatex

import (
	"log"
	"testing"
)

func TestTemplateInfo_templatex(t *testing.T) {
	temp := New().SetName("name").SetDescription("description").SetContent("content")

	log.Print(temp.Cat())
}
