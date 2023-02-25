package spc

// (c) Christian Maurer   v. 230213 - license see µU.go

import (
  "µU/vect"
  "µU/spc/stack"
  "µU/gl"
)
const
  epsilon = 1e-6
var (
  origin, focus = vect.New(), vect.New()
  right, front, top = vect.New(), vect.New(), vect.New()
// Invariant: Right, front and top build an orthogonal right-handed trihedron.
  temp = vect.New()
)

func init() {
  right = vect.New3 (1, 0, 0)
  front = vect.New3 (0, 1, 0)
  top   = vect.New3 (0, 0, 1)
}

func set (ox, oy, oz, fx, fy, fz, tx, ty, tz float64) {
  origin.Set3 (ox, oy, oz)
  focus.Set3 (fx, fy, fz)
  top.Set3 (tx, ty, tz)
  front.Copy (focus); front.Sub (origin)
  right.Ext (front, top)
  right.Ext (front, top)
  right.Norm(); front.Norm(); top.Norm()
/*/
x, y, z := getFront(); println ("set front =", x, y, z)
x, y, z  = getTop();   println ("set top   =", x, y, z)
x, y, z  = getRight(); println ("set right =", x, y, z); println()
/*/
}

func getOrigin() (float64, float64, float64) {
  return origin.Coord3()
}

func getFocus() (float64, float64, float64) {
  return focus.Coord3()
}

func getRight() (float64, float64, float64) {
  return right.Coord3()
}

func getFront() (float64, float64, float64) {
  return front.Coord3()
}

func getTop() (float64, float64, float64) {
  return top.Coord3()
}

func adjustFocus() {
  d := origin.Distance (focus)
  focus.Scale (d, front)
  focus.Add (origin) // focus = origin + | origin - focus | * front
}

func moveR (d float64) {
  temp.Scale (d, right)
  origin.Add (temp) // origin += d * right
  adjustFocus()
}

func moveF (d float64) {
  temp.Scale (d, front)
  origin.Add (temp) // origin += d * front
  adjustFocus()
}

func moveT (d float64) {
  temp.Scale (d, top)
  origin.Add (temp) // origin += d * top
  adjustFocus()
}

func move1R (d float64) {
  temp.Scale (d, right)
  origin.Add (temp) // origin += d * right
  focus.Add (temp) // focus += d * right
}

func move1F (d float64) {
  temp.Scale (d, front)
  origin.Add (temp) // origin += d * front
  focus.Add (temp) // focus += d * front 
}

func move1T (d float64) {
  temp.Scale (d, top)
  origin.Add (temp) // origin += d * top
  focus.Add (temp) // focus += d * top
}

func rotR (alpha float64) {
  front.Rot (right, alpha)
  front.Norm()
  top.Ext (right, front)
  top.Norm()
}

func rotF (alpha float64) {
  top.Rot (front, alpha)
  top.Norm()
  right.Ext (front, top)
  right.Norm()
}

func rotT (alpha float64) {
  right.Rot (top, alpha)
  right.Norm()
  front.Ext (top, right)
  front.Norm()
}

func tilt (alpha float64) {
  rotR (alpha)
  adjustFocus()
}

func roll (alpha float64) {
  rotF (alpha)
  adjustFocus()
}

func turn (alpha float64) {
  rotT (alpha)
  adjustFocus()
}

func turnAroundFocusR (alpha float64) {
  d := origin.Distance (focus)
  if d < epsilon { return }
  rotR (-alpha)
  temp.Scale (d, front)
  origin.Diff (focus, temp) // origin = focus - | focus - origin | * front
}

func turnAroundFocusT (alpha float64) {
  d := origin.Distance (focus)
  if d < epsilon { return }
  rotT (-alpha)
  temp.Scale (d, front)
  origin.Diff (focus, temp) // origin = focus - | focus - origin | * front
}

func empty() bool {
  return stack.Empty()
}

func push() {
  ox, oy, oz := origin.Coord3()
  fx, fy, fz := focus.Coord3()
  tx, ty, tz := top.Coord3()
  stack.Push (ox, oy, oz, fx, fy, fz, tx, ty, tz)
}

func pop() {
  x := stack.Pop()
  set (x[0], x[1], x[2], x[3], x[4], x[5], x[6], x[7], x[8])
}

func empty1() bool {
  return stack.Empty1()
}

func push1() {
  ox, oy, oz := origin.Coord3()
  fx, fy, fz := focus.Coord3()
  tx, ty, tz := top.Coord3()
  stack.Push1 (ox, oy, oz, fx, fy, fz, tx, ty, tz)
}

func pop1() {
  x := stack.Pop1()
  set (x[0], x[1], x[2], x[3], x[4], x[5], x[6], x[7], x[8])
}

func setLight (n uint) {
  x, y, z := origin.Coord3()
  gl.PosLight (n, x, y, z)
}
