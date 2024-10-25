package main

import (
	"my_first_test/svc"
	"my_first_test/svc/impl"
)

func main() {
	svc.NewMyPlan(impl.NewClassHome()).Lean("张三")
	svc.NewMyPlan(impl.NewGoogle()).Lean("李四")
}
