package spc

// (c) Christian Maurer   v. 191019 - license see µU.go

func Set (ex, ey, ez, fx, fy, fz, nx, ny, nz float64) { set (ex,ey,ez,fx,fy,fz,nx,ny,nz) }

func Get() (float64, float64, float64, float64, float64, float64, float64, float64, float64) {
  return get()
}

// func Distance() float64 { return distance() }

func Move (i int, d float64) { move(i,d) }

func Turn (i int, a float64) { turn(i,a) }

func Invert() { invert() }

// func Focus (d float64) { foc(d) }

func TurnAroundFocus (i int, a float64) { turnAroundFocus(i,a) }

func SetLight (n uint) { setLight(n) }
