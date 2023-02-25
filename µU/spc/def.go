package spc

// (c) Christian Maurer   v. 230213 - license see µU.go

// The package maintains the following 5 vectors:
// origin, focus and an orthogonal right-handed trihedron (right, front, top)
// with len(right) = len(front) = len(top) = 1, s.t. front = focus - origin normed to len 1.
// Maintains furthermore two stacks of those quintupels.

// origin = (ox, oy, oz), focus = (fx, fy, fz), top = (tx, ty, tz),
// front = focus - origin normed to len 1 and right = cross-product front x top.
func Set (ox, oy, oz, fx, fy, fz, tx, ty, tz float64) { set (ox,oy,oz,fx,fy,fz,tx,ty,tz) }

// Returns the coordinates of origin, focus and top.
func GetOrigin() (float64, float64, float64) { return getOrigin() }
func GetFocus()  (float64, float64, float64) { return getFocus() }
func GetRight()  (float64, float64, float64) { return getRight() }
func GetFront()  (float64, float64, float64) { return getFront() }
func GetTop()    (float64, float64, float64) { return getTop() }

// origin is moved in direction Right/Front/Top by distance d,
func MoveRight (d float64) { moveR(d) }
func MoveFront (d float64) { moveF(d) }
func MoveTop (d float64) { moveT(d) }

// origin and focus are moved in direction Right/Front/Top by distance d,
func Move1Right (d float64) { move1R(d) }
func Move1Front (d float64) { move1F(d) }
func Move1Top (d float64) { move1T(d) }

func Tilt (a float64) { tilt(a) }
// front is rotated around right by angle a, top is adjusted.

func Roll (a float64) { roll(a) }
// top is rotated around front by angle a, right is adjusted.

func Turn (a float64) { turn(a) }
// right is rotated around top by angle a, front is adjusted.

// TODO Spec
func TurnAroundFocusR (a float64) { turnAroundFocusR(a) }

// TODO Spec
func TurnAroundFocusT (a float64) { turnAroundFocusT(a) }

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
func SetLight (n uint) { setLight(n) }
