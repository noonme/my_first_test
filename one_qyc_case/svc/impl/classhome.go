package impl

import "fmt"

type ClassHomeInf interface {
	ClassHomeGo(string) (string, error)
	ClassHomeC(string) (string, error)
	ClassHomeJava(string) (string, error)
	ClassHomeVioline(string) (string, error)
}
type ClassHome struct {
}

var _ClassHomeInf = (*ClassHome)(nil)

func NewClassHome() *ClassHome {
	return &ClassHome{}
}

func (c *ClassHome) LeanGo(s string) (string, error) {
	fmt.Println(s, "在ClassHome学Go语言")
	return "ClassHome-go", nil
}

func (c *ClassHome) LeanC(s string) (string, error) {
	fmt.Println(s, "在ClassHome学C语言")
	return "ClassHome-c", nil
}

func (c *ClassHome) LeanJava(s string) (string, error) {
	fmt.Println(s, "在ClassHome学Java语言")
	return "ClassHome-java", nil
}

func (c *ClassHome) LeanVioline(s string) (string, error) {
	fmt.Println(s, "在ClassHome学小提琴")
	return "ClassHome-violine", nil
}
