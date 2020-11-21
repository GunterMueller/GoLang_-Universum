package spc

// (c) Christian Maurer   v. 201103 - license see µU.go

// The package maintains the following 5 vectors:
// origin,
// focus and
// an orthogonal right-handed trihedron (right, front, top)
//   with len(right) = len(front) = len(top) = 1,
//   s.t. front = focus - origin normed to len 1.
// Maintains furthermore two stacks of such quintupels.

// origin = (ox, oy, oz), focus = (fx, fy, fz), top = (tx, ty, tz),
// front = focus - origin normed to len 1 and right = cross-product front x top.
func Set (ox, oy, oz, fx, fy, fz, tx, ty, tz float64) { set (ox,oy,oz,fx,fy,fz,tx,ty,tz) }

// Returns (ox, oy, oz, fx, fy, fz, tx, ty, tz).
func Get() (float64, float64, float64, float64, float64, float64, float64, float64, float64) { return get() }

// Returns the coordinates of origin.
func Get3() (float64, float64, float64) { return get3() }

// Pre: i < 3.
// origin is moved in direction i by distance d,
// where i = 0/1/2 means direction of right/front/top.
func Move (i uint, d float64) { move(i,d) }

// Pre: i < 3.
// origin and focus are moved in direction i by distance d,
// where i = 0/1/2 means direction of right/front/top.
func Move1 (i uint, d float64) { move1(i,d) }

// Pre: i < 3.
// trihedron[n] (n = (i + 1) % 3) is rotated around trihedron[i] by angle a,
// trihedron[p] (p = (i + 2) % 3) is adjusted.
func Turn (i uint, a float64) { turn(i,a) }

// Pre: i < 3.
// i = 0, 1, 2 means direction of right, front, top.
// TODO Spec
func TurnAroundFocus (i uint, a float64) { turnAroundFocus(i,a) }

// Returns true, iff the corresponding stack is empty.
func Empty() bool { return empty() }
func Empty1() bool { return empty1() }

// The quintuple is pushed onto the corresponding stack.
func Push() { push() }
func Push1() { push1() }

// The quintuple is popped from the corresponding stack.
func Pop() { pop() }
func Pop1() { pop1() }

// TODO Spec
// func SetLight (n uint) { setLight(n) }
