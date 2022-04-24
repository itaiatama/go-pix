package main

import (
	"image/color"

	pixapp "github.com/itaiatama/go-pix/pix-app"
	pixext "github.com/itaiatama/go-pix/pix-ext"
	pixui "github.com/itaiatama/go-pix/pix-ui"
	SDL "github.com/veandco/go-sdl2/sdl"
)

type App struct {
}

func (a *App) Init() {
}

func (a *App) Event(event SDL.Event) {
	pixui.Update(event)

	switch event.(type) {
	case *SDL.QuitEvent:
		pixapp.Exit()
	}
}

func (a *App) Update(dt float64) {}

func (a *App) Render(R *SDL.Renderer) {
	pixext.DrawClear(R, color.RGBA{64, 64, 64, 255})
}

func main() {
	pixapp.SetTitle("Example")
	pixapp.SetSize(640, 360)
	pixapp.SetFpsCap(120)

	pixapp.Run(&App{})
}
