package client

import (
	"fmt"
	"net/http"
)

// // case1 功能提供方
// type LeanInfClient interface {
// 	LearnGo(string) (string, error)
// 	LeanC(string) (string, error)
// 	LeanJava(string) (string, error)
// 	LeanVioline(string) (string, error)
// }

// func NewLeanClient() LeanInfClient {
// 	return &LeanClientImpl{}
// }

// type LeanClientImpl struct {
// }

// // var _LeanInfClient = (*LeanClientImpl)(nil)

// func (l *LeanClientImpl) LearnGo(req string) (string, error) {
// 	fmt.Println("http.Post(\"\", \"\", nil)")

// 	return req + " ClassHome-go", nil
// }
// func (l *LeanClientImpl) LeanC(req string) (string, error) {
// 	http.Post("", "", nil)

// 	return "ClassHome-c", nil
// }

// func (l *LeanClientImpl) LeanJava(req string) (string, error) {
// 	http.Post("", "", nil)

// 	return "ClassHome-java", nil
// }

// func (l *LeanClientImpl) LeanVioline(req string) (string, error) {
// 	http.Post("", "", nil)

// 	return "ClassHome-violine", nil
// }

//case 2 功能提供方--核心理念interface应该使用者根据需求自定义，而非程序提供方定义

func NewLeanClient() *LeanInfClient {
	return &LeanInfClient{}
}

type LeanInfClient struct {
}

// var _LeanInfClient = (*LeanClientImpl)(nil)

func (l *LeanInfClient) LearnGo(req string) (string, error) {
	fmt.Println("http.Post(\"\", \"\", nil)")

	return req + " ClassHome-go", nil
}
func (l *LeanInfClient) LeanC(req string) (string, error) {
	http.Post("", "", nil)

	return "ClassHome-c", nil
}

func (l *LeanInfClient) GLeanJava(req string) (string, error) {
	http.Post("", "", nil)

	return "ClassHome-java", nil
}

func (l *LeanInfClient) LeanVioline(req string) (string, error) {
	http.Post("", "", nil)

	return "ClassHome-violine", nil
}
