package spc

// (c) Christian Maurer   v. 201031 - license see µU.go

// The package maintains the following 5 vectors:
// origin,
// focus and
// an orthogonal right-handed trihedron (right, front, top)
//   with len(right) = len(front) = len(top) = 1,
//   s.t. front = focus - origin normed to len 1.

// origin = (ox, oy, oz), focus = (fx, fy, fz), top = (tx, ty, tz),
// front = focus - origin normed to len 1 and right = cross-product front x top.
func Set (ox, oy, oz, fx, fy, fz, tx, ty, tz float64) { set (ox,oy,oz,fx,fy,fz,tx,ty,tz) }
func Set3 (ox, oy, oz float64) { set3(ox,oy,oz) }

// Returns (ox, oy, oz, fx, fy, fz, tx, ty, tz).
func Get() (float64, float64, float64, float64, float64, float64, float64, float64, float64) {
  return get()
}

// Returns (ox, oy, oz).
func Get3() (float64, float64, float64) { x, y, z, _, _, _, _, _, _ := get(); return x, y, z }

// Pre: i < 3.
// origin is moved in direction i by distance d,
// where i = 0/1/2 means direction of right/front/top.
func Move (i uint, d float64) { move(i,d) }

// Pre: i < 3.
// trihedron[n] (n = (i + 1) % 3) is rotated around trihedron[i] by angle a,
// trihedron[p] (p = (i + 2) % 3) is adjusted.
func Turn (i uint, a float64) { turn(i,a) }

// Pre: i < 3.
// i = 0, 1, 2 means direction of right, front, top.
// TODO Spec
func TurnAroundFocus (i uint, a float64) { turnAroundFocus(i,a) }

// TODO Spec
// func SetLight (n uint) { setLight(n) }
