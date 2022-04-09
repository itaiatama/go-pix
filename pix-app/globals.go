package pixapp

var (
	Title     string = "pix-app"
	Width     int    = 360
	Height    int    = 180
	Running   bool   = false
	TargetFPS int    = 60
)

func SetWindowTitle(t string) { Title = t }
func SetWindowWidth(w int)    { Width = w }
func SetWindowHeight(h int)   { Height = h }
func SetTargetFPS(t int)      { TargetFPS = t }
