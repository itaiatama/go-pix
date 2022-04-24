package main

import (
	"image/color"

	pixapp "github.com/itaiatama/go-pix/pix-app"
	pixext "github.com/itaiatama/go-pix/pix-ext"
	pixui "github.com/itaiatama/go-pix/pix-ui"
	SDL "github.com/veandco/go-sdl2/sdl"
	TTF "github.com/veandco/go-sdl2/ttf"
)

type App struct {
}

var font *TTF.Font = nil
var tex *SDL.Texture = nil
var w, h int = 0, 0

func (a *App) Init() {
	font = pixapp.LoadFont("assets/fonts/m6x11.ttf", 24)
	tex, w, h = pixapp.CreateTextTexture(font, "Hello, World!", color.RGBA{255, 255, 255, 255})
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
	pixext.DrawTexture(R, tex, 640/2-w/2, 10, w, h)

	pixui.Begin()
	pixui.ButtonText(R, font, "Button", 3, 0, 10, 10, 20)
	pixui.ButtonText(R, font, "Button", 4, 100, 10, 10, 20)
	pixui.ButtonText(R, font, "Button", 5, 200, 10, 10, 20)
	pixui.End()
}

func (a *App) Destroy() {
	tex.Destroy()
	font.Close()
	pixui.Destroy()
}

func main() {
	pixapp.SetTitle("Example")
	pixapp.SetSize(640, 360)
	pixapp.SetFpsCap(120)

	pixapp.Run(&App{})
}
