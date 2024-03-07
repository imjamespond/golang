package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Curl(val string) string {
	return Curl(val)
}

func (a *App) Save(infs string, funs string, defs string) {
	Save("out.inf.ts", infs)
	Save("out.fun.ts", funs)
	Save("out.def.ts", defs)
}

func (a *App) GetTpl(file string) string {
	data := Read(file)
	if data == nil {
		return ""
	}
	return *data
}
