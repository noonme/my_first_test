package svc

import "fmt"

type MyLeanProxy interface {
	LeanGo(string) (string, error)
}
type MyPlan struct {
	leanEvent MyLeanProxy
}

func NewMyPlan(leanEvent MyLeanProxy) *MyPlan {
	return &MyPlan{
		leanEvent: leanEvent,
	}
}

func (m *MyPlan) Lean(req string) {
	requset, err := m.leanEvent.LeanGo(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(requset)
}
