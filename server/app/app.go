package app

import "fmt"

type App struct {
	configFile string
}

type Option func(a *App)

func ConfigFile(file string) Option {
	return func(a *App) {
		a.configFile = file
	}
}

func (a *App) InitServer() {
	fmt.Println("<<< INICIANDO >>>")

	x := a.ReadConfig()

	fmt.Println("CONFIG >>>", x)
}
