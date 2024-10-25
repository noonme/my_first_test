package server

import "testing"

// //case 1
// func Test_LeanServiceImpl_LearnGO(t *testing.T) {
// 	mockService := NewLeanService(&mockLearnClient{})
// 	rep, err := mockService.LearnGO("张三")
// 	if err != nil {
// 		t.Error("mockService.LearnGO(\"张三\")")
// 	}
// 	t.Log(rep)

// }

// case 2
func Test_LeanServiceImpl_LearnGO(t *testing.T) {
	mockService := NewLeanService(&mockLearnClient{})
	// rep, err := mockService.LearnGO("张三")
	rep, err := mockService.LearnGo("李四")
	if err != nil {
		t.Error("mockService.LearnGO(\"张三\")")
	}
	t.Log(rep)

}
