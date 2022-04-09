package main

import (
	pixapp "github.com/itaiatama/go-pix/pix-app"
	SDL "github.com/veandco/go-sdl2/sdl"
)

type App struct{}

func (a *App) Init() {}

func (a *App) Event(event SDL.Event) {
	switch event.(type) {
	case *SDL.QuitEvent:
		pixapp.Exit()
	}
}

func (a *App) Update(dt float64) {}

func (a *App) Render(R *SDL.Renderer) {}

func main() {
	pixapp.SetWindowTitle("GO PIX EXAMPLE")
	pixapp.SetWindowWidth(640)
	pixapp.SetWindowHeight(360)
	pixapp.SetTargetFPS(120)

	pixapp.Run(&App{})
}
