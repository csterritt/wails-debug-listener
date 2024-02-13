package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

type ShowLine struct {
	Line   string `json:"line"`
	Sender string `json:"sender"`
	Type   string `json:"type"`
}

func setupListener(wailsApp *App) {
	listenerApp := fiber.New()

	listenerApp.Post("/", func(c *fiber.Ctx) error {
		line := new(ShowLine)
		if err := c.BodyParser(line); err != nil {
			return err
		}

		runtime.LogPrintf(wailsApp.ctx, "%s: %s %s\n", line.Sender, line.Type, line.Line)

		runtime.EventsEmit(wailsApp.ctx, "incoming", line)

		return nil
	})

	port := 3030
	err := listenerApp.Listen(fmt.Sprintf(`:%d`, port))
	if err != nil {
		panic(err)
	}
}

// NewApp creates a new App application struct
func NewApp() *App {
	app := &App{}

	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	setupListener(a)
}

func (a *App) SendUp(info string) string {
	runtime.EventsEmit(a.ctx, "incoming", "Hello from the server!")

	return "-"
}
