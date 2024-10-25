package impl

import "fmt"

type GoogleInf interface {
	GoogleGo(string) (string, error)
	GoogleC(string) (string, error)
	GoogleJava(string) (string, error)
	GoogleVioline(string) (string, error)
}
type Google struct {
}

var _GoogleInf = (*Google)(nil)

func NewGoogle() *Google {
	return &Google{}
}

func (g *Google) LeanGo(s string) (string, error) {
	fmt.Println(s, "在Google学Go语言")
	return "Google-go", nil
}

func (g *Google) LeanC(s string) (string, error) {
	fmt.Println(s, "在Google学C语言")
	return "Google-c", nil
}

func (g *Google) LeanJava(s string) (string, error) {
	fmt.Println(s, "在Google学Java语言")
	return "Google-java", nil
}

func (g *Google) LeanVioline(s string) (string, error) {
	fmt.Println(s, "在Google学小提琴")
	return "Google-violine", nil
}
