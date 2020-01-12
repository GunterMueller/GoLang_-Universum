package eye

// (c) Christian Maurer   v. 191018 - license see µU.go

import (
  "math"
//  . "µU/obj"
//  "µU/col"
//  "µU/scr"; "µU/errh"
  "µU/vect"
  "µU/gl" // PosLight, Actualize
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
         vec [3]vect.Vector
       delta float64 // invariant: delta == Distance (Auge, focus)
//      colour col.Colour
//        flat bool
             }
var
  nB, nD, nF uint
//  var e = New() // -> Anwendung

func new_() Eye {
  e := new (eye)
  e.origin = vect.New()
  e.originOld = vect.New()
  e.focus = vect.New()
  e.temp = vect.New()
  e.vec[0] = vect.New3 (1, 0, 0)
  e.vec[1] = vect.New3 (0, 1, 0)
  e.vec[2] = vect.New3 (0, 0, 1)
  e.delta = e.origin.Distance (e.focus)
//  e.colour, _ = scr.StartCols()
//  e.colour = col.White()
  return e
}

func (e *eye) SetLight (n uint) {
  gl.PosLight (n, e.origin)
}

/*
func (e *eye) Actualize() {
  gl.Actualize (e.vec[right], e.vec[front], e.vec[top], e.origin)
}
*/

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
  e.focus.Scale (e.delta, e.vec[front])
  e.focus.Add (e.origin)
//  e.Actualize()
}

func (e *eye) Distance() float64 {
  if math.Abs (e.origin.Distance (e.focus) - e.delta) > epsilon {
    e.adjustFocus()
  }
  return e.delta
}

/*
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
*/

func (e *eye) Move (i int, dist float64) {
  nB++
  e.vec[top].Ext (e.vec[right], e.vec[front])
  e.vec[top].Norm()
  e.originOld.Copy (e.origin)
  e.temp.Scale (dist, e.vec[i])
  e.origin.Add (e.temp)
  e.adjustFocus()
}

func (e *eye) rotate (i int, alpha float64) { // ziemlich abenteuerliche Konstruktion
  V1 := e.vec[(i + 1) % 3]
  V1.Rot (e.vec[i], alpha)
  V1.Norm()
//  V2.Copy (e.vec[i])
  V2 := e.vec[i].Clone().(vect.Vector)
  V2.Cross (V1)
  V2.Ext (e.vec[i], V1)
  V2.Norm()
}

func (e *eye) Turn (i int, alpha float64) {
  nD++
  e.rotate (i, alpha)
  e.adjustFocus()
}

func (e *eye) Invert() {
  nD++
  e.vec[right].Dilate (-1.0)
  e.vec[front].Dilate (-1.0)
  e.adjustFocus()
}

func (e *eye) adjustOrigin() {
  e.temp.Scale (e.delta, e.vec[front])
  e.origin.Sub (e.focus, e.temp)
//  e.Actualize()
}

func (e *eye) Focus (d float64) {
  if d < epsilon { return }
  e.delta = d
  e.adjustOrigin()
}

func (e *eye) TurnAroundFocus (i int, alpha float64) {
  if e.delta < epsilon { return }
  nF++
//  println ("TurnAroundFocus")
  e.rotate (i, -alpha)
//  println ("rotated")
// Dieser Vorzeichenwechsel ist ein Arbeitsdrumrum
// um einen mir bisher nicht erklärbaren Fehler.
// Vermutlich liegt das daran, dass ich irgendeine suboptimal
// dokumentierte Eigenschaft von openGL noch nicht begriffen habe.
  e.adjustOrigin()
}

/*
func (e *eye) Setx (x, y, z, x1, y1, z1 float64) {
  e.origin.Set3 (x, y, z)
  e.focus.Set3 (x1, y1, z1)
  e.delta = e.origin.Distance (e.focus)
  if e.delta < epsilon { return } // error
  if math.Abs (z - z1) < epsilon { // Blick horizontal
    e.vec[top].Set3 (0, 0, 1)
    e.vec[front].Sub (e.focus, e.origin)
    e.vec[front].Norm()
////    e.vec[right].Copy (e.vec[front]) e.vec[right].Cross (e.vec[top])
    e.vec[right].Ext (e.vec[front], e.vec[top])
    e.vec[right].Norm()
  } else { // z != z1
    if math.Abs (x - x1) < epsilon && math.Abs (y - y1) < epsilon { // x == x1 und y == y1
      e.vec[right].Set3 (1, 0, 0)
      e.vec[front].Set3 (0, 0, 1)
      e.vec[top].Set3 (0, 1, 0) // XXX
      if z > z1 { // Blick von top, x -> right, y -> top
        e.vec[front].Dilate (-1)
      } else { // z < z1 *) // Blick von unten, x -> right, y -> unten
        e.vec[top].Dilate (-1)
      }
    } else { // x != x1 oder y != y1
      e.vec[front].Sub (e.focus, e.origin)
      e.vec[front].Norm()
      v2 := e.vec[front].Coord (top)
      e.vec[top].Copy (e.vec[front])
      if z < z1 { v2 = -v2 }
      e.temp.Set3 (0., 0., - 1./v2)
      e.vec[top].Add (e.temp)
      e.vec[top].Norm()
      e.vec[right].Ext (e.vec[front], e.vec[top])
      e.vec[right].Norm()
    }
  }
//  e.Actualize()
}
*/

func Report() {
//  errh.Error2 ("Bewegungen:", nB, "/ Drehungen:", nD)
  println (nB, "Bewegungen,", nD, "Drehungen")
}

/*
var (
  stack[]([]byte) = make ([]([]byte), 100)
//  v vect.Vector = vect.New()
)

// Vielleicht geht das folgende ja noch einfacher ...

func (e *eye) zl() uint {
  return 4 * e.temp.Codelen() + e.colour.Codelen()
}

func (e *eye) Push (c col.Colour) {
  B := make ([]byte, e.zl())
  a := 0
  copy (B[a:a+8], e.origin.Encode())
  a += 8
  for i := 0; i < 3; i++ {
    copy (B[a:a+8], e.vec[i].Encode())
    a += 8
  }
  copy (B[a:a+97], c.Encode())
  stack = append (stack, B)
}

func (e *eye) Colour() col.Colour {
  B := stack[len(stack) - 1]
  a := 0
  e.origin.Decode (B[a:a+8])
  a += 8
  for i := 0; i < 3; i++ {
    e.vec[i].Decode (B[a:a+8])
    a += 8
  }
  c := col.New()
  c.Decode (B[a:a+3])
  stack = stack[0:len(stack) - 2]
  return c
}
*/
