package main

//go:generate mockgen -destination=./mock/mock_LearnInf.go -package=mock -source=./interface.go
type LeranInf interface {
	LearnGo(req string) (string, error)
	LeanC(req string) (string, error)
}
