package scale

// (c) Christian Maurer   v. 150425 - license see µU.go

// "Virtual points" denote those points, whose coordinates (X, Y)
// can be mapped to the screen (w.r.t. the actual screen mode).
// At the beginning the range of their coordinates
// ist given by 0 <= X < 1 und 0 <= Y < 1/scr.Proportion,
// where (0,0) is the left bottom corner of the screen.

// Pre: The screen is initialized.
// scale is initialized. This is the precondition for all other functions !
func Init() { init_() }

// Pre: w > 0.
// The range of the coordinates (X, Y) of the virtual points is given
// by x <= X <= x + w and y <= Y <= y + h mit h = w / scr.Proportion;
// particularly (x, y) is the left bottom corner.
func SetRange (x, y, w float64) { setRange(x,y,w) }

// Pre: (x, y) is in the range of the coordinates of the virtual points.
// Returns the screen pixel position of (x, y).
// If the calling process runs under X, the result is in the range
// MinInt16 .. MaxInt16, if it runs in a console, in (0 .. Wd, 0 .. Ht).
func Scale (x, y float64) (int, int) { return scale(x,y) }

// Up to rounding effects the inverse function of Scale.
func Rescale (x, y int) (float64, float64) { return rescale(x,y) }

// TODO Spec
func Lim (x, y, w, h, v float64) { lim(x,y,w,h,v) }

// The transformation magnification is manipulated by the user.
func Edit() { edit() }
