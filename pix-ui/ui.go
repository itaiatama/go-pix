package pixui

import (
	"image/color"

	pixapp "github.com/itaiatama/go-pix/pix-app"
	pixext "github.com/itaiatama/go-pix/pix-ext"
	SDL "github.com/veandco/go-sdl2/sdl"
	TTF "github.com/veandco/go-sdl2/ttf"
)

type State struct {
	MX int
	MY int
	ML bool

	Hot    int
	Active int
}

var s *State = &State{}

type TextObj struct {
	W        int
	H        int
	Texture  *SDL.Texture
	FontName string
}

var texts map[string]*TextObj = make(map[string]*TextObj)

func InRect(x, y, w, h int) bool {
	return (s.MX >= x && s.MX <= x+w && s.MY >= y && s.MY <= y+h)
}

var (
	RED   = color.RGBA{200, 64, 64, 255}
	GREEN = color.RGBA{64, 200, 64, 255}
	BLUE  = color.RGBA{64, 64, 200, 255}
	WHITE = color.RGBA{200, 200, 200, 255}
	BLACK = color.RGBA{0, 0, 0, 255}
)

func Begin() { s.Hot = 0 }

func End() {
	if !s.ML {
		s.Active = 0
	} else {
		if s.Active == 0 {
			s.Active = -1
		}
	}
}

func Button(R *SDL.Renderer, ID, x, y, w, h int) bool {
	bg := GREEN

	if InRect(x, y, w, h) {
		s.Hot = ID
		if s.Active == 0 && s.ML {
			s.Active = ID
		}
	}

	if s.Hot == ID {
		if s.Active == ID {
			bg = RED
		} else {
			bg = BLUE
		}
	}

	pixext.DrawFillRect(R, x, y, w, h, bg)
	return (!s.ML && s.Hot == ID && s.Active == ID)
}

func ButtonText(R *SDL.Renderer, F *TTF.Font, text string, ID, x, y, px, py int) bool {

	t, b := texts[text]

	if !b || t.FontName != F.FaceFamilyName() {
		obj := &TextObj{}
		sur := pixapp.CreateTextSurface(F, text, BLACK)
		obj.FontName = F.FaceFamilyName()
		obj.W = int(sur.W)
		obj.H = int(sur.H)
		obj.Texture = pixapp.CreateImageTexture(sur)
		sur.Free()
		texts[text] = obj
	}

	tex := texts[text].Texture
	w := texts[text].W
	h := texts[text].H
	bg := GREEN

	if InRect(x, y, w+px, h+py) {
		s.Hot = ID
		if s.Active == 0 && s.ML {
			s.Active = ID
		}
	}

	if s.Hot == ID {
		if s.Active == ID {
			bg = RED
		} else {
			bg = BLUE
		}
	}

	pixext.DrawFillRect(R, x, y, w+px, h+py, bg)
	pixext.DrawTexture(R, tex, x+px/2, y+py/2, w, h)
	return (!s.ML && s.Hot == ID && s.Active == ID)
}

func Update(event SDL.Event) {
	switch e := event.(type) {
	case *SDL.MouseMotionEvent:
		s.MX = int(e.X)
		s.MY = int(e.Y)
	case *SDL.MouseButtonEvent:
		s.ML = (e.Button == SDL.BUTTON_LEFT && e.State == SDL.PRESSED)
	}
}

func Destroy() {
	for _, t := range texts {
		t.Texture.Destroy()
	}
}
