package eye

// (c) Christian Maurer   v. 191117 - license see µU.go

import (
  "math"
  . "µU/obj"
  "µU/col"
  "µU/vect"
  "µU/gl"
)
const (
  epsilon = 1.0E-6
  right = 0
  front = 1
  top = 2
)
type
  eye struct {
      origin,
   originOld,
       focus,
        temp vect.Vector
      vector [3]vect.Vector
       delta float64 // invariant: delta == Distance (Auge, focus)
      colour col.Colour
        flat bool
             }

func next (i int) int {
  return (i + 1) % 3
}

func prev (i int) int {
  return (i + 2) % 3
}

func new_() Eye {
  e := new (eye)
  e.origin = vect.New()
  e.originOld = vect.New()
  e.focus = vect.New()
  e.temp = vect.New()
  e.vector[0] = vect.New3 (1, 0, 0)
  e.vector[1] = vect.New3 (0, 1, 0)
  e.vector[2] = vect.New3 (0, 0, 1)
  e.delta = e.origin.Distance (e.focus)
  e.colour = col.White()
  return e
}

func (e *eye) SetLight (n uint) {
  gl.PosLight (n, e.origin)
}

func (e *eye) actualize() {
//  gl.Actualize (e.vec[0], e.vec[1], e.vec[2], e.origin)
}

func (e *eye) Set (ex, ey, ez, fx, fy, fz float64) {
  e.origin.Set3 (ex, ey, ez)
  e.focus.Set3 (fx, fy, fz)
}

func (e *eye) Get() (float64, float64, float64, float64, float64, float64) {
  ex, ey, ez := e.origin.Coord3()
  fx, fy, fz := e.focus.Coord3()
  return ex, ey, ez, fx, fy, fz
}

func (e *eye) DistanceFrom (aim vect.Vector) float64 {
  return e.origin.Distance (aim)
}

func (e *eye) adjustFocus() {
  e.focus.Scale (e.delta, e.vector[front])
  e.focus.Add (e.origin)
  e.actualize()
}

func (e *eye) Distance() float64 {
  if math.Abs (e.origin.Distance (e.focus) - e.delta) > epsilon {
    e.adjustFocus()
  }
  return e.delta
}

func (e *eye) Read (v []vect.Vector) bool {
  v[0].Copy (e.originOld)
  if len (v) > 1 {
    v[1].Copy (e.origin)
  }
  return e.flat
}

func (e *eye) Flatten (f bool) {
  e.flat = f
}

func (e *eye) Move (i int, dist float64) {
  e.vector[top].Ext (e.vector[right], e.vector[front])
  e.vector[top].Norm()
  e.originOld.Copy (e.origin)
  e.temp.Scale (dist, e.vector[i])
  e.origin.Add (e.temp)
  e.adjustFocus()
}

func (e *eye) rotate (i int, alpha float64) { // ziemlich abenteuerliche Konstruktion
  V1 := e.vector[next(i)]
  V1.Rot (e.vector[i], alpha)
  V1.Norm()
//  V2.Copy (e.vector[i])
  V2 := e.vector[i].Clone().(vect.Vector)
  V2.Cross (V1)
  V2.Ext (e.vector[i], V1)
  V2.Norm()
}

func (e *eye) Turn (i int, alpha float64) {
  e.rotate (i, alpha)
  e.adjustFocus()
}

func (e *eye) Invert() {
  e.vector[right].Dilate (-1)
  e.vector[front].Dilate (-1)
  e.adjustFocus()
}

func (e *eye) adjustOrigin() {
  e.temp.Scale (e.delta, e.vector[front])
  e.origin.Sub (e.focus, e.temp)
  e.actualize()
}

func (e *eye) Focus (d float64) {
  if d < epsilon { return }
  e.delta = d
  e.adjustOrigin()
}

func (e *eye) TurnAroundFocus (i int, alpha float64) {
  if e.delta < epsilon { return }
//  println ("TurnAroundFocus")
  e.rotate (i, -alpha)
//  println ("rotated")
// Dieser Vorzeichenwechsel ist ein Arbeitsdrumrum
// um einen mir bisher nicht erklärbaren Fehler.
// Vermutlich liegt das daran, dass ich irgendeine suboptimal
// dokumentierte Eigenschaft von openGL noch nicht begriffen habe.
  e.adjustOrigin()
}

func (e *eye) Setx (ex, ey, ez, fx, fy, fz float64) {
  e.origin.Set3 (ex, ey, ez)
  e.focus.Set3 (fx, fy, fz)
  e.delta = e.origin.Distance (e.focus)
  if e.delta < epsilon { return } // error
  if math.Abs (ez - fz) < epsilon { // ez == fz: Blick horizontal
    e.vector[top].Set3 (0, 0, 1)
    e.vector[front].Sub (e.focus, e.origin)
    e.vector[front].Norm()
////    e.vector[right].Copy (e.vector[front]) e.vector[right].Cross (e.vector[top])
    e.vector[right].Ext (e.vector[front], e.vector[top])
    e.vector[right].Norm()
  } else { // ez != fz
    if math.Abs (ex - fx) < epsilon && math.Abs (ey - fy) < epsilon { // ex == fx und ey == fy
      e.vector[right].Set3 (1, 0, 0)
      e.vector[front].Set3 (0, 0, 1)
      e.vector[top].Set3 (0, 1, 0) // XXX
      if ez > fz { // Blick von top, x -> right, y -> top
        e.vector[front].Dilate (-1)
      } else { // ez < fz // Blick von unten, x -> right, y -> unten
        e.vector[top].Dilate (-1)
      }
    } else { // ex != fx oder ey != fy
      e.vector[front].Sub (e.focus, e.origin)
      e.vector[front].Norm()
      v2 := e.vector[front].Coord (top)
      e.vector[top].Copy (e.vector[front])
      if ez < fz { v2 = -v2 }
      e.temp.Set3 (0, 0, -1/v2)
      e.vector[top].Add (e.temp)
      e.vector[top].Norm()
      e.vector[right].Ext (e.vector[front], e.vector[top])
      e.vector[right].Norm()
    }
  }
  e.actualize()
}

var
  stack = make ([]Stream, 100)

func (e *eye) codelen() uint {
  return 4 * e.temp.Codelen() + e.colour.Codelen()
}

func (e *eye) Push (c col.Colour) {
  bs := make (Stream, e.codelen())
  i, a := uint(0), e.origin.Codelen()
  copy (bs[i:i+a], e.origin.Encode())
  for j := 0; j < 3; j++ {
    copy (bs[i:i+a], e.vector[j].Encode())
    i += a
  }
  a = 3
  copy (bs[i:i+a], c.Encode())
  stack = append (stack, bs)
}

func (e *eye) Top() col.Colour {
  n := len(stack)
  bs := stack[n - 1]
  i, a := uint(0), e.origin.Codelen()
  e.origin.Decode (bs[i:i+a])
  for j := 0; j < 3; j++ {
    e.vector[j].Decode (bs[i:i+a])
    i += a
  }
  c := col.New()
  a = c.Codelen()
  c.Decode (bs[i:i+a])
  stack = stack[:n - 1]
  return c
}
