package pixui

import (
	"image/color"
	"log"

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

func InRect(x, y, w, h int) bool {
	return (s.MX >= x && s.MX <= x+w && s.MY >= y && s.MY <= y+h)
}

var (
	RED   = color.RGBA{200, 64, 64, 255}
	GREEN = color.RGBA{64, 200, 64, 255}
	BLUE  = color.RGBA{64, 64, 200, 255}
	WHITE = color.RGBA{200, 200, 200, 255}
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

func Button(R *SDL.Renderer, ID, x, y int) bool {
	if InRect(x, y, 64, 48) {
		s.Hot = ID
		if s.Active == 0 && s.ML {
			s.Active = ID
		}
	}

	if s.Hot == ID {
		if s.Active == ID {
			pixext.DrawFillRect(R, x, y, 64, 48, RED)
		} else {
			pixext.DrawFillRect(R, x, y, 64, 48, BLUE)
		}
	} else {
		pixext.DrawFillRect(R, x, y, 64, 48, GREEN)
	}

	return (!s.ML && s.Hot == ID && s.Active == ID)
}

func ButtonText(R *SDL.Renderer, F *TTF.Font, text string, ID, x, y, p int) bool {

	surface, err := F.RenderUTF8Solid(text, SDL.Color(WHITE))

	if err != nil {
		log.Fatal(err)
	}

	texture, err := R.CreateTextureFromSurface(surface)

	if err != nil {
		log.Fatal(err)
	}

	w := int(surface.W)
	h := int(surface.H)

	if InRect(x-p/2, y-p/2, w+p, h+p) {
		s.Hot = ID
		if s.Active == 0 && s.ML {
			s.Active = ID
		}
	}

	if s.Hot == ID {
		if s.Active == ID {
			pixext.DrawFillRect(R, x-p/2, y-p/2, w+p, h+p, RED)
		} else {
			pixext.DrawFillRect(R, x-p/2, y-p/2, w+p, h+p, BLUE)
		}
	} else {
		pixext.DrawFillRect(R, x-p/2, y-p/2, w+p, h+p, GREEN)
	}

	pixext.DrawTexture(R, texture, x, y, w, h)
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
