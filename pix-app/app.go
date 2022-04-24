package pixapp

import (
	SDL "github.com/veandco/go-sdl2/sdl"
	TTF "github.com/veandco/go-sdl2/ttf"
)

func Exit() { Running = false }

func Run(a App) {
	CheckError(SDL.Init(SDL.INIT_EVERYTHING))
	defer SDL.Quit()

	CheckError(TTF.Init())
	defer TTF.Quit()

	window, err := SDL.CreateWindow(Title, SDL.WINDOWPOS_CENTERED, SDL.WINDOWPOS_CENTERED, int32(Width), int32(Height), SDL.WINDOW_SHOWN)
	CheckError(err)
	defer window.Destroy()
	Window = window

	renderer, err := SDL.CreateRenderer(window, -1, SDL.RENDERER_ACCELERATED)
	CheckError(err)
	defer renderer.Destroy()
	Renderer = renderer

	Running = true

	a.Init()

	delta := 0.0

	for Running {

		start := SDL.GetTicks()

		for event := SDL.PollEvent(); event != nil; event = SDL.PollEvent() {
			a.Event(event)
		}

		a.Update(delta)

		a.Render(renderer)
		renderer.Present()

		wait := int64(1.0/float64(FpsCap)*1000) - (int64(SDL.GetTicks()) - int64(start))

		if wait >= 0 {
			SDL.Delay(uint32(wait))
		}

		delta = (float64(SDL.GetTicks()) - float64(start)) / 1000.0
	}
}
