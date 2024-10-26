package main

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
)

type Child struct{}

func (c *Child) GetChild() string {
	return "boy"
}

func Test_Child_GetChild(t *testing.T) {
	// 创建 Child 类型的指针
	cPtr := &Child{}

	// 使用 gomonkey 来 mock GetChild 方法
	re := gomonkey.ApplyMethod(reflect.TypeOf(cPtr), "GetChild", func(_ *Child) string {
		return "girl"
	})
	defer re.Reset()

	// re := gomonkey.ApplyFunc(reflect.TypeOf(cPtr).Elem().Name()+".GetChild", func(_ *Child) string {
	// 	return "girl"
	// })
	// defer re.Reset()

	// 调用 GetChild 方法并检查返回值
	if got := cPtr.GetChild(); got != "girl" {
		t.Errorf("Expected 'girl', but got '%s'", got)
	}
	if got := cPtr.GetChild(); got != "girl" {
		t.Errorf("Expected 'girl', but got '%s'", got)
	}

	// 再次调用以确保 mock 效果被重置
	if got := cPtr.GetChild(); got != "boy" {
		t.Errorf("Expected 'boy', but got '%s'", got)
	}
	// 再次调用以确保 mock 效果被重置
	if got := cPtr.GetChild(); got != "boy" {
		t.Errorf("Expected 'boy', but got '%s'", got)
	}
}
