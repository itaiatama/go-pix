package pixapp

import SDL "github.com/veandco/go-sdl2/sdl"

type App interface {
	Init()
	Event(SDL.Event)
	Update(float64)
	Render(*SDL.Renderer)
	Destroy()
}
