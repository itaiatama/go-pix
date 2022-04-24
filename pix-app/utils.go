package pixapp

import (
	"image/color"
	"log"

	IMG "github.com/veandco/go-sdl2/img"
	SDL "github.com/veandco/go-sdl2/sdl"
	TTF "github.com/veandco/go-sdl2/ttf"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LoadImageSurface(path string) *SDL.Surface {
	sur, err := IMG.Load(path)
	CheckError(err)
	return sur
}

func LoadImageTexture(path string) *SDL.Texture {
	sur := LoadImageSurface(path)
	tex := CreateImageTexture(sur)
	sur.Free()
	return tex
}

func CreateImageTexture(S *SDL.Surface) *SDL.Texture {
	tex, err := Renderer.CreateTextureFromSurface(S)
	CheckError(err)
	return tex
}

func LoadFont(path string, size int) *TTF.Font {
	font, err := TTF.OpenFont(path, size)
	CheckError(err)
	return font
}

func CreateTextSurface(F *TTF.Font, text string, clr color.RGBA) *SDL.Surface {
	sur, err := F.RenderUTF8Blended(text, SDL.Color(clr))
	CheckError(err)
	return sur
}

func CreateTextTexture(F *TTF.Font, text string, clr color.RGBA) (tex *SDL.Texture, w int, h int) {
	sur := CreateTextSurface(F, text, clr)
	w = int(sur.W)
	h = int(sur.H)
	tex = CreateImageTexture(sur)
	sur.Free()
	return tex, w, h
}
