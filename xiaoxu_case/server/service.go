package server

import "fmt"

// // case 1功能使用方--基于面向对象的设计模式需要mockinterface实现所有对应接口过去繁琐
// type LeanService interface{}

// func NewLeanService(LearnClient client.LeanInfClient) *LeanServiceImpl {
// 	return &LeanServiceImpl{
// 		LearnClient: LearnClient,
// 	}
// }

// type LeanServiceImpl struct {
// 	LearnClient client.LeanInfClient
// }

// var _LeanService = (*LeanServiceImpl)(nil)

// func (l *LeanServiceImpl) LearnGO(name string) (string, error) {
// 	// step 1

// 	// step 2

// 	// step 3

// 	//1、 http.Post("", "", nil)

// 	//依然会重复多次调用
// 	// 2、 c := client.NewLean()
// 	// resp, err := c.LeanGo(name)
// 	// if err != nil {
// 	// 	return "", err
// 	// }

// 	//3、只会调用一次,可以实现复用，不会每次创建都调用多次
// 	resp, err := l.LearnClient.LearnGo(name)
// 	if err != nil {
// 		return "", err
// 	}
// 	return resp, nil

// 	// step 4

// 	// step 5

// }

// type mockLearnClient struct {
// }

// func (l *mockLearnClient) LearnGo(req string) (string, error) {
// 	fmt.Println("Go 学习")

// 	return req + " response Go 学习 case1", nil
// }

// func (l *mockLearnClient) LeanC(req string) (string, error) {
// 	http.Post("", "", nil)

// 	return "ClassHome-c", nil
// }

// func (l *mockLearnClient) LeanJava(req string) (string, error) {
// 	http.Post("", "", nil)

// 	return "ClassHome-java", nil
// }

// func (l *mockLearnClient) LeanVioline(req string) (string, error) {
// 	http.Post("", "", nil)

// 	return "ClassHome-violine", nil
// }

// case2 功能使用方
type LeanService interface {
	// LearnGo(req string) (string, error)
}

type LeanProxy interface {
	LearnGo(req string) (string, error)
}

type LeanServiceImpl struct {
	LearnClient LeanProxy
}

func NewLeanService(LearnClient LeanProxy) *LeanServiceImpl {
	return &LeanServiceImpl{
		LearnClient: LearnClient,
	}
}

// var _LeanService = (*leanServiceImpl)(nil)

func (l *LeanServiceImpl) LearnGo(req string) (string, error) {
	// step 1

	// step 2

	// step 3

	//1、 http.Post("", "", nil)

	//依然会重复多次调用
	// 2、 c := client.NewLean()
	// resp, err := c.LeanGo(name)
	// if err != nil {
	// 	return "", err
	// }

	//3、只会调用一次,可以实现复用，不会每次创建都调用多次
	resp, err := l.LearnClient.LearnGo(req)
	if err != nil {
		return "", err
	}
	return resp, nil

	// step 4

	// step 5

}

type mockLearnClient struct {
}

func (l *mockLearnClient) LearnGo(req string) (string, error) {
	fmt.Println("Go 学习")

	return req + " response Go 学习 case1", nil
}
