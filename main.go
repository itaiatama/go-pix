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
	tex *SDL.Texture
	fnt *TTF.Font
}

func (a *App) Init() {
	a.tex = pixapp.LoadTexture("assets/textures/tex.png")
	a.fnt = pixapp.LoadFont("assets/fonts/m6x11.ttf", 16)
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

	pixui.Begin()
	pixui.Button(R, 3, 10, 10)
	pixui.ButtonText(R, a.fnt, "Button", 6, 100, 10, 8)
	pixui.End()

	pixext.DrawText(R, a.fnt, "Hello, World!", color.RGBA{255, 255, 255, 255}, 100, 100)
	pixext.DrawSubTexture(R, a.tex, 0, 0, 16, 16, 200, 200, 64, 64)

}

func main() {
	pixapp.SetTitle("Example")
	pixapp.SetSize(640, 360)
	pixapp.SetFpsCap(120)

	pixapp.Run(&App{})
}
