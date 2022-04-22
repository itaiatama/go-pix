# go-pix
Simple golang graphics [`framework`](https://en.wikipedia.org/wiki/Application_framework) ([go-sdl2](https://github.com/veandco/go-sdl2) wrapper).
Contains several modules: 
* `pix-app` - core module that provides `Application` interface and contains main application logic.
* `pix-ext` - extension module:
	- Drawing simple shapes (only rectangels for now).
	- Drawing textures and sub-textures.
	- Drawing text with loaded font.
* `pix-ui` - simple [`imgui`](https://en.wikipedia.org/wiki/Immediate_mode_GUI) module (ugly for now):
	- Simple solid color button.
	- Simple soild color button with text.

## To try
```
$ git clone https://github.com/itaiatama/go-pix.git
$ cd go-pix
$ make run
```

## Example
``` go
package main

import (
	"image/color"

	pixapp "github.com/itaiatama/go-pix/pix-app"
	pixext "github.com/itaiatama/go-pix/pix-ext"
	pixui "github.com/itaiatama/go-pix/pix-ui"
	SDL "github.com/veandco/go-sdl2/sdl"
)

type App struct{}

func (a *App) Init() {}

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
	pixui.End()
}

func main() {
	pixapp.SetTitle("Example")
	pixapp.SetSize(640, 360)
	pixapp.SetFpsCap(120)

	pixapp.Run(&App{})
}
```
