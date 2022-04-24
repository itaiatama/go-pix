package pixext

import (
	"image/color"

	SDL "github.com/veandco/go-sdl2/sdl"
)

func DrawClear(R *SDL.Renderer, clr color.RGBA) {
	R.SetDrawColor(clr.R, clr.G, clr.B, clr.A)
	R.Clear()
}

func DrawWireRect(R *SDL.Renderer, x, y, w, h int, clr color.RGBA) {
	R.SetDrawColor(clr.R, clr.G, clr.B, clr.A)
	R.DrawRect(&SDL.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)})
}

func DrawFillRect(R *SDL.Renderer, x, y, w, h int, clr color.RGBA) {
	R.SetDrawColor(clr.R, clr.G, clr.B, clr.A)
	R.FillRect(&SDL.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)})
}

func DrawTexture(R *SDL.Renderer, T *SDL.Texture, x, y, w, h int) {
	R.Copy(T, nil, &SDL.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)})
}

func DrawTextureRegion(R *SDL.Renderer, T *SDL.Texture, rx, ry, rw, rh, x, y, w, h int) {
	src := &SDL.Rect{X: int32(rx), Y: int32(ry), W: int32(rw), H: int32(rh)}
	dst := &SDL.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)}
	R.Copy(T, src, dst)
}
