package pixui

import (
	"image/color"

	pixext "github.com/itaiatama/go-pix/pix-ext"
	SDL "github.com/veandco/go-sdl2/sdl"
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

func Update(event SDL.Event) {
	switch e := event.(type) {
	case *SDL.MouseMotionEvent:
		s.MX = int(e.X)
		s.MY = int(e.Y)
	case *SDL.MouseButtonEvent:
		s.ML = (e.Button == SDL.BUTTON_LEFT && e.State == SDL.PRESSED)
	}
}
