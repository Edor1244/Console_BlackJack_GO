package cmd

type app struct{}

func (a app) Run() {
	panic("implement me")
}

type App interface {
	Run()
}

func NewApp() App {
	return &app{}
}
