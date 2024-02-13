package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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
	go func() {
		time.Sleep(2 * time.Second)
		a.SendUp("Greet sending up for " + name)
	}()
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) SendUp(info string) string {
	runtime.EventsEmit(a.ctx, "incoming", "Hello from the server!")

	return "-"
}
