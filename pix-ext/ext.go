package pixext

import (
	"image/color"
	"log"

	SDL "github.com/veandco/go-sdl2/sdl"
	TTF "github.com/veandco/go-sdl2/ttf"
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

func DrawSubTexture(R *SDL.Renderer, T *SDL.Texture, sx, sy, sw, sh, x, y, w, h int) {
	R.Copy(T, &SDL.Rect{X: int32(sx), Y: int32(sy), W: int32(sw), H: int32(sh)}, &SDL.Rect{X: int32(x), Y: int32(y), W: int32(w), H: int32(h)})
}

func DrawText(R *SDL.Renderer, F *TTF.Font, text string, clr color.RGBA, x, y int) {
	surface, err := F.RenderUTF8Solid(text, SDL.Color(clr))

	if err != nil {
		log.Fatal(err)
	}

	texture, err := R.CreateTextureFromSurface(surface)

	if err != nil {
		log.Fatal(err)
	}

	DrawTexture(R, texture, x, y, int(surface.W), int(surface.H))
}
