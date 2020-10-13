package vect

// (c) Christian Maurer   v. 201009 - license see µU.go

import (
  "math"
  "strconv"
  . "µU/obj"
  "µU/str"
//  "µU/col"
)
const (
  um = math.Pi / 180.0
  null = 0.0
)
type
  vector struct {
              x [3]float64
                }
var (
  temp, temp1 = new_().(*vector), new_().(*vector)
//  bx = box.New()
//  pbx = pbox.New()
)

func new_() Vector {
  return new(vector)
}

func (x *vector) imp (Y Any) *vector {
  y, ok := Y.(*vector)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func new3 (x0, x1, x2 float64) Vector {
  return &vector { [3]float64{ x0, x1, x2 } }
}

func (v *vector) Set3 (x0, x1, x2 float64) {
  v.x = [3]float64{x0, x1, x2}
}

func (v *vector) Set (i int, c float64) {
  v.x[i] = c
}

func (v *vector) Coord3() (float64, float64, float64) {
  return v.x[0], v.x[1], v.x[2]
}

func (v *vector) Coord (d int) float64 {
  return v.x[d]
}

func (v *vector) SetPolar (x, y, z, r, phi, theta float64) {
  v.x[0] = x + r * math.Cos (phi * um) * math.Sin (theta * um)
  v.x[1] = y + r * math.Sin (phi * um) * math.Sin (theta * um)
  v.x[2] = z + r                       * math.Cos (theta * um)
}

func (v *vector) Project (A, B, C Vector) {
  a, b, c := v.imp(A), v.imp(B), v.imp(C)
  for d := 0; d < 3; d++ {
    a.x[d], b.x[d], c.x[d] = null, null, null
  }
  a.x[0], b.x[1], c.x[2] = v.x[0], v.x[1], v.x[2]
}

func (v *vector) Empty() bool {
  a := null
  for d := 0; d < 3; d++ {
    a += math.Abs (v.x[d])
  }
  return a < epsilon
}

func (v *vector) Clr() {
  v.Set3 (null, null, null)
}

func (v *vector) Copy (Y Any) {
  y := v.imp(Y)
  for d := 0; d < 3; d++ {
    v.x[d] = y.x[d]
  }
}

func (v *vector) Clone() Any {
  y := New().(*vector)
  y.Copy (v)
  return y
}

func (x *vector) Eq (Y Any) bool {
  y, a := x.imp (Y), null
  for d := 0; d < 3; d++ {
    a += math.Abs (x.x[d] - y.x[d])
  }
  return a < epsilon
}

func (x *vector) Less (Y Any) bool {
  return false
}

func (x *vector) Int (Y Vector) float64 {
  y, a := x.imp (Y), null
  for d := 0; d < 3; d++ {
    a += x.x[d] * y.x[d]
  }
  return a
}

func next (d uint) uint {
  return (d + 1) % 3
}

func prev (d uint) uint {
  return (d + 2) % 3
}

func (x *vector) Cross (Y Vector) {
  y := x.imp (Y)
  var a [3]float64
  for d := 0; d < 3; d++ {
    a[d] = x.x[d]
  }
  for d := uint(0); d < 3; d++ {
    n, p := next (d), prev (d)
    x.x[d] = a[n] * y.x[p] - a[p] * y.x[n]
  }
}

func (x *vector) Ext (Y, Z Vector) {
  y, z := x.imp (Y), x.imp (Z)
  for d := uint(0); d < 3; d++ {
    n, p := next (d), prev (d)
    x.x[d] = y.x[n] * z.x[p] - y.x[p] * z.x[n]
  }
}

func (x *vector) Collinear (Y Vector) bool {
  y := x.imp (Y)
  if x.Empty() || y.Empty() {
    return true
  }
  temp.Copy (x)
  temp.Cross (y)
  return temp.Empty()
}

func (x *vector) Scale (a float64, Y Vector) {
  y := x.imp(Y)
  for d := 0; d < 3; d++ {
    x.x[d] = a * y.x[d]
  }
}

func (x *vector) Translate (Y Vector) {
  y := x.imp(Y)
  for d := 0; d < 3; d++ {
    x.x[d] += y.x[d]
  }
}

func (x *vector) Dilate (a float64) { // TODO name ?
  for d := 0; d < 3; d++ {
    x.x[d] *= a
  }
}

func (x *vector) Null() bool {
  return x.Empty()
}

func (x *vector) Zero() bool {
  return x.x[0] == 0. &&
         x.x[1] == 0. &&
         x.x[2] == 0.
}

func (x *vector) Add (Y ...Adder) {
  for _, y := range Y {
    for d := 0; d < 3; d++ {
      x.x[d] += x.imp(y).x[d]
    }
  }
}

func (x *vector) Sub (Y ...Adder) {
  for _, y := range Y {
    for d := 0; d < 3; d++ {
      x.x[d] -= x.imp(y).x[d]
    }
  }
}

func (x *vector) Sum (Y, Z Vector) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.Add (z)
}

func (x *vector) Diff (Y, Z Vector) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.Sub (z)
}

func (x *vector) Parametrize (Y, Z Vector, t float64) {
  y, z := x.imp (Y), x.imp (Z)
  for d := 0; d < 3; d++ {
    x.x[d] = y.x[d] + t * (z.x[d] - y.x[d])
  }
}

func (x *vector) Len() float64 {
  return math.Sqrt (x.Int (x))
}

func (x *vector) Distance (Y Vector) float64 {
  y := Y.(*vector)
  a, s := null, null
  for d := 0; d < 3; d++ {
    s = x.x[d] - y.x[d]
    a += s * s
  }
  return math.Sqrt (a)
}

func (x *vector) Centre (Y, Z Vector) float64 {
  y, z := x.imp (Y), x.imp (Z)
  a, s := null, null
  for d := 0; d < 3; d++ {
    x.x[d] = (y.x[d] + z.x[d]) / 2.0
    s = y.x[d] - z.x[d]
    a += s * s
  }
  return math.Sqrt (a) / 2.0
}

func (x *vector) Flat (Y Vector) bool {
  y := Y.(*vector)
  return math.Abs (x.x[2] - y.x[2]) < epsilon
}

func (x *vector) Norm() {
  a := math.Sqrt (x.Int (x))
  for d := 0; d < 3; d++ {
    x.x[d] /= a
  }
}

func (x *vector) Normed() bool {
  return math.Abs (x.Len() - 1.0) < epsilon
}

func (x *vector) Rot (Y Vector, a float64) {
  y := Y.(*vector)
  for a <= -180. { a += 360. }
  for a > 180. { a -= 360. }
  if x.Collinear (y) { return } // error
//  y.Norm() // avoid rounding errors
  c := math.Cos (a * um)
// x |-> cos(a) * x + <x, y> * (1 - cos(a)) * y + sin(a) * [y, x]
  temp.Scale ((1. - c) * x.Int (y), y)
//  temp.Copy (y)
//  temp.Dilate ((1. - c) * x.Int (y))
//  temp1.Copy (y)
//  temp1.Cross (x)
  temp1.Ext (y, x)
  temp1.Dilate (math.Sin (a * um))
  x.Dilate (c)
  x.Add (temp)
  x.Add (temp1)
}

func (x *vector) Defined (s string) bool {
  x.Clr()
  n := uint(len (s))
  if n < 7 { return false }
  if s[0] != '(' || s[n - 1] != ')' { return false }
  t := str.Part (s, 1, n - 2) + ","
  for d := 0; d < 3; d++ {
    p, ok := str.Pos (t, ',')
    if ! ok { return false }
    r, err := strconv.ParseFloat (t[:p], 64)
    if err == nil {
      x.x[d] = r
    } else {
      return false
    }
    str.Rem (&t, 0, p+1)
  }
  return true
}

func (x *vector) String() string {
  s := "("
  for d := 0; d < 3; d++ {
    s += strconv.FormatFloat (x.x[d], 'f', 2, 64)
    if d == 2 {
      s += ") "
    } else {
      s += ", "
    }
  }
  return s
}

var
  clfloat = Codelen(null)

func (x *vector) Codelen() uint {
  return 3 * clfloat
}

func (x *vector) Encode() []byte {
  b := make ([]byte, x.Codelen())
  i, a := uint(0), clfloat
  for d := 0; d < 3; d++ {
    copy (b[i:i+a], Encode (x.x[d]))
    i += a
  }
  return b
}

func (x *vector) Decode (b []byte) {
  i, a := uint(0), clfloat
  for d := 0; d < 3; d++ {
    x.x[d] = Decode (null, b[i:i+a]).(float64)
    i += a
  }
}

func (V *vector) Minimax (N, X Vector) {
  Min, n := N.(*vector)
  Max, x := X.(*vector)
  if ! n || ! x { return }
  for d := 0; d < 3; d++ {
    if V.x[d] < Min.x[d] {
      Min.x[d] = V.x[d]
    }
    if V.x[d] > Max.x[d] {
      Max.x[d] = V.x[d]
    }
  }
}

/*
func (x *vector) Parallelogram (y, z Vector) []Vector {
  x1, x2, x3 := New(), New(), New()
  x1.Add (x, y)
  x2.Add (x1, z)
  x3.Add (x, z)
  return []Vector {x1, x2, x3}
}
*/
