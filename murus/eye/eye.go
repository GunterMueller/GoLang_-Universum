package eye

// (c) murus.org  v. 161216 - license see murus.go

import (
  "math"
  . "murus/spc"; "murus/col"
  "murus/errh"
  "murus/vect"
  "murus/gl" // PosLight, Actualize
)
const
  epsilon = 1.0E-6
type
  eye struct {
      origin,
   originOld,
       focus,
        temp vect.Vector
         vec [NDirs]vect.Vector
       delta float64 // Invariant: delta == Distance (Auge, focus)
      colour col.Colour
    flaechig bool
             }
var
  nB, nD, nF uint
//  var e = New() // -> Anwendung

func newEye() Eye {
  e := new (eye)
  e.origin = vect.New()
  e.originOld = vect.New()
  e.focus = vect.New()
  e.temp = vect.New()
  for d := D0; d < NDirs; d++ {
    e.vec[d] = vect.New()
    e.vec[d].Set (Unit[d])
  }
  e.delta = e.origin.Distance (e.focus)
  e.colour, _ = col.StartCols()
  e.flaechig = false
  return e
}

func (e *eye) SetLight (n uint) {
  gl.PosLight (n, e.origin)
}

func (e *eye) Actualize() {
  gl.Actualize (e.vec[Right], e.vec[Front], e.vec[Top], e.origin)
}

func (e *eye) DistanceFrom (aim vect.Vector) float64 {
  return e.origin.Distance (aim)
}

func (e *eye) focusAnpassen() {
  e.focus.Scale (e.delta, e.vec[Front])
  e.focus.Add (e.origin)
  e.Actualize()
}

func (e *eye) Distance() float64 {
  if math.Abs (e.origin.Distance (e.focus) - e.delta) > epsilon {
    e.focusAnpassen()
  }
  return e.delta
}

func (e *eye) Read (v []vect.Vector) bool {
  v[0].Copy (e.originOld)
  if len (v) > 1 {
    v[1].Copy (e.origin)
  }
  return e.flaechig
}

func (e *eye) Flatten (f bool) {
  e.flaechig = f
}

func (e *eye) Move (d Direction, dist float64) {
  nB ++
  e.vec[Top].Copy (e.vec[Right])
  e.vec[Top].Cross (e.vec[Front])
//  e.vec[Top].Ext (e.vec[Right], e.vec[Front])
  e.vec[Top].Norm()
  e.originOld.Copy (e.origin)
  e.temp.Scale (dist, e.vec[d])
  e.origin.Add (e.temp)
  e.focusAnpassen()
}

func (e *eye) rotate (d Direction, alpha float64) { // ziemlich abenteuerliche Konstruktion
  V1 := e.vec[Next (d)]
  V1.Rot (e.vec[d], alpha)
  V1.Norm()
  V2 := e.vec[Prev (d)]
  V2.Copy (e.vec[d])
  V2.Cross (V1)
//  V2.Ext (e.vec[d], V1)
  V2.Norm()
}

func (e *eye) Turn (d Direction, alpha float64) {
  nD++
  e.rotate (d, alpha)
  e.focusAnpassen()
}

func (e *eye) Invert() {
  nD++
  e.vec[Right].Dilate (-1.0)
  e.vec[Front].Dilate (-1.0)
  e.focusAnpassen()
}

func (e *eye) originAnpassen() {
  e.temp.Scale (e.delta, e.vec[Front])
  e.origin.Copy (e.focus)
  e.origin.Sub (e.temp)
//  e.origin.Sub (e.focus, e.temp)
  e.Actualize()
}

func (e *eye) Focus (d float64) {
  if d < epsilon { return }
  e.delta = d
  e.originAnpassen()
}

func (e *eye) TurnAroundFocus (D Direction, alpha float64) {
  if e.delta < epsilon { return }
  nF++
//  println ("TurnAroundFocus")
  e.rotate (D, - alpha)
//  println ("rotated")
// Dieser Vorzeichenwechsel ist ein Arbeitsdrumrum
// um einen mir bisher nicht erklärbaren Fehler.
// Vermutlich liegt das daran, dass ich irgendeine suboptimal
// dokumentierte Eigenschaft von openGL noch nicht begriffen habe.
  e.originAnpassen()
}

func (e *eye) Set (x, y, z, xf, yf, zf float64) {
  e.origin.Set3 (x, y, z)
  e.focus.Set3 (xf, yf, zf)
  e.delta = e.origin.Distance (e.focus)
  if e.delta < epsilon { return } // error
  if math.Abs (z - zf) < epsilon { // Blick horizontal
    e.vec[Top].Set (Unit[Top])
    e.vec[Front].Copy (e.focus)
    e.vec[Front].Sub (e.origin)
//    e.vec[Front].Sub (e.focus, e.origin)
    e.vec[Front].Norm()
    e.vec[Right].Copy (e.vec[Front])
    e.vec[Right].Cross (e.vec[Top])
//    e.vec[Right].Ext (e.vec[Front], e.vec[Top])
    e.vec[Right].Norm()
  } else { // z != zf
    if math.Abs (x - xf) < epsilon && math.Abs (y - yf) < epsilon { // x == xf und y == yf
      e.vec[Right].Set (Unit[Right])
      e.vec[Front].Set (Unit[Top])
      e.vec[Top].Set (Unit[Right])
      if z > zf { // Blick von Top, x -> Right, y -> Top
        e.vec[Front].Dilate (-1.0)
      } else { // z < zf *) // Blick von unten, x -> Right, y -> unten
        e.vec[Top].Dilate (-1.0)
      }
    } else { // x != xf oder y != yf
      e.vec[Front].Copy (e.focus)
      e.vec[Front].Sub (e.origin)
//      e.vec[Front].Sub (e.focus, e.origin)
      e.vec[Front].Norm()
      v2 := e.vec[Front].Coord (Top)
      e.vec[Top].Copy (e.vec[Front])
      if z < zf { v2 = -v2 }
      e.temp.Set3 (0., 0., - 1. / v2)
      e.vec[Top].Add (e.temp)
      e.vec[Top].Norm()
      e.vec[Right].Copy (e.vec[Front])
      e.vec[Right].Cross (e.vec[Top])
//      e.vec[Right].Ext (e.vec[Front], e.vec[Top])
      e.vec[Right].Norm()
    }
  }
  e.Actualize()
}

func Report() {
  errh.Error2 ("Bewegungen:", nB, "/ Drehungen:", nD)
}

var (
  stack[]([]byte) = make ([]([]byte), 100)
  v vect.Vector = vect.New()
)

// Vielleicht geht das folgende ja noch einfacher ...

func zl() uint {
  return 4 * v.Codelen() + col.Codelen()
}

func (e *eye) Push (c col.Colour) {
  B := make ([]byte, zl())
  a := 0
  copy (B[a:a+8], e.origin.Encode())
  a += 8
  for d := D0; d < NDirs; d++ {
    copy (B[a:a+8], e.vec[d].Encode())
    a += 8
  }
  copy (B[a:a+97], col.Encode (c))
  stack = append (stack, B)
}

func (e *eye) Colour() col.Colour {
  B := stack[len(stack) - 1]
  a := 0
  e.origin.Decode (B[a:a+8])
  a += 8
  for d := D0; d < NDirs; d++ {
    e.vec[d].Decode (B[a:a+8])
    a += 8
  }
  var c col.Colour
  col.Decode (&c, B[a:a+3])
  stack = stack[0:len(stack) - 2]
  return c
}
