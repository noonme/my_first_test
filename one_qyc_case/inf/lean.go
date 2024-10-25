package inf

type Lean interface {
	LeanGo(string) (string, error)
	LeanC(string) (string, error)
	LeanJava(string) (string, error)
	LeanVioline(string) (string, error)
}
