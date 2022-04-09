package pixapp

import SDL "github.com/veandco/go-sdl2/sdl"

var (
	Window  *SDL.Window
	Title   string = "pix-app"
	Width   int    = 360
	Height  int    = 180
	Running bool   = false
	FpsCap  int    = 60
)

func SetTitle(t string) {
	Title = t
	Window.SetTitle(t)
}

func SetSize(w, h int) {
	Width = w
	Height = h
	Window.SetSize(int32(w), int32(h))
}

func SetFpsCap(f int) {
	FpsCap = f
}
