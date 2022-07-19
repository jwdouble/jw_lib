package templatex

import (
	"fmt"
)

type templatex interface {
	SetName(s string) templatex
	SetDescription(s string) templatex
	SetContent(s string) templatex
	Cat() string
	Create()
}

type templateInfo struct {
	Name        string
	Description string
	Content     string
}

func New() *templateInfo {
	return &templateInfo{}
}

func (t *templateInfo) SetName(s string) *templateInfo {
	t.Name = s
	return t
}

func (t *templateInfo) SetDescription(s string) *templateInfo {
	t.Description = s
	return t
}

func (t *templateInfo) SetContent(s string) *templateInfo {
	t.Content = s
	return t
}

func (t *templateInfo) Cat() string {
	return fmt.Sprintf("name: %s, description: %s, content: %s", t.Name, t.Description, t.Content)
}

func (t *templateInfo) Create() {

}
