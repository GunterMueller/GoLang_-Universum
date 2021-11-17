package qmat

// (c) Christian Maurer   v. 211106 - license see µU.go

// >>> matrices with rational fractions as entries

import (
  . "µU/obj"
  "µU/ker"
  "µU/errh"
  "µU/n"
  "µU/q"
)
type
  qmatrix struct {
          nl, nc uint // number of lines and columns
               d uint // maximal number of digits of the fractions
         matrix  [][]q.Rational
                 }

func nofit() {
  ker.Panic ("wrong number of lines or columns")
}

func new_(m, n, d uint) QMatrix {
  x := new (qmatrix)
  x.nl, x.nc = m, n
  x.d = d
  x.matrix = make([][]q.Rational, n)
  for l := uint(0); l < n; l++ {
    x.matrix[l] = make([]q.Rational, n)
    for c := uint(0); c < n; c++ {
      x.matrix[l][c] = q.New()
    }
  }
  return x
}

func unit (m, n, d uint) QMatrix {
  u := new_(m, n, d).(*qmatrix)
  for l := uint(0); l < n; l++ {
    for c := uint(0); c < n; c++ {
      if c == l {
        u.matrix[l][c].Set1 (1)
      } else {
        u.matrix[l][c].Set1 (0)
      }
    }
  }
  return u
}

func (x *qmatrix) imp (Y Any) *qmatrix {
  y, ok := Y.(*qmatrix)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *qmatrix) Det() q.Rational {
  if x.nl != x.nc { nofit() }
  b := q.New()
  b.Set1 (0)
  for c := uint(0); c < x.nc; c++ {
    a := x.matrix[0][c].Clone().(q.Rational)
    for l := uint(1); l < x.nc; l++ {
      a.Mul (x.matrix[l][(c + l) % x.nc])
    }
    b.Add (a)
  }
  b1 := q.New()
  b1.Set1 (0)
  for c := uint(0); c < x.nc; c++ {
    a := x.matrix[0][c].Clone().(q.Rational)
    for l := uint(1); l < x.nc; l++ {
      a.Mul (x.matrix[l][(c + x.nc - l) % x.nc])
    }
    b1.Add (a)
  }
  b.Sub (b1)
  return b
}

func (x *qmatrix) Zero() bool {
  for l := uint(0); l < x.nl; l++ {
    for c := uint(0); c < x.nc; c++ {
      if ! x.matrix[l][c].Zero() {
        return false
      }
    }
  }
  return true
}

func (x *qmatrix) Eq (Y Any) bool {
  y := x.imp (Y)
  if y.nl != x.nl || y.nc != x.nc || y.d != x.d {
    return false
  }
  for l := uint(0); l < x.nl; l++ {
    for c := uint(0); c < x.nc; c++ {
      if ! x.matrix[l][c].Eq (y.matrix[l][c]) {
        return false
      }
    }
  }
  return true
}

func (x *qmatrix) Copy (Y Any) {
  y := x.imp (Y)
  if y.nl != x.nl || y.nc != x.nc { nofit() }
  for l := uint(0); l < x.nl; l++ {
    for c := uint(0); c < x.nc; c++ {
      x.matrix[l][c].Copy (y.matrix[l][c])
    }
  }
}

func (x *qmatrix) Clone () Any {
  y := new_(x.nl, x.nc, x.d)
  y.Copy (x)
  return y
}

func (x *qmatrix) Less (Y Any) bool {
  return false
}

func (x *qmatrix) Empty() bool {
  for l := uint(0); l < x.nl; l++ {
    for c := uint(0); c < x.nc; c++ {
      if x.matrix[l][c].Empty() {
        return true
      }
    }
  }
  return false
}

func (x *qmatrix) Clr() {
  for l := uint(0); l < x.nl; l++ {
    for c := uint(0); c < x.nc; c++ {
      x.matrix[l][c].Clr()
    }
  }
}

func (x *qmatrix) Codelen() uint {
  return 2 * C0 + x.nl * x.nc * x.matrix[0][0].Codelen()
}

func (x *qmatrix) Encode() Stream {
  s := make(Stream, x.Codelen())
  i, a := uint(0), C0
  copy (s[i:i+a], Encode (x.nl))
  i += a
  copy (s[i:i+a], Encode (x.nc))
  i += a
  a = x.matrix[0][0].Codelen()
  for l := uint(0); l < x.nl; l++ {
    for c := uint(0); c < x.nc; c++ {
      copy (s[i:i+a], x.matrix[l][c].Encode())
      i += a
    }
  }
  return s
}

func (x *qmatrix) Decode (s Stream) {
  i, a := uint(0), C0
  x.nl = Decode (uint(0), s[i:i+a]).(uint)
  i += a
  x.nc = Decode (uint(0), s[i:i+a]).(uint)
  i += a
  a = x.matrix[0][0].Codelen()
  for l := uint(0); l < x.nl; l++ {
    for c := uint(0); c < x.nc; c++ {
      x.matrix[l][c].Decode (s[i:i+a])
      i += a
    }
  }
}

func (x *qmatrix) add (Y Adder) {
  y := x.imp(Y)
  if y.nl != x.nl || y.nc != x.nc || y.d != x.d { nofit() }
  for l := uint(0); l < x.nl; l++ {
    for c := uint(0); c < x.nc; c++ {
      x.matrix[l][c].Add (y.matrix[l][c])
    }
  }
}

func (x *qmatrix) Add (Y ...Adder) {
  for i, _ := range Y {
    x.add (Y[i])
  }
}

func (x *qmatrix) Sum (Y, Z Adder) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.add (z)
}

func (x *qmatrix) Diff (Y, Z Adder) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.sub (z)
}

func (x *qmatrix) sub (Y Adder) {
  y := x.imp(Y)
  if y.nl != x.nl || y.nc != x.nc || y.d != x.d { nofit() }
  for l := uint(0); l < x.nl; l++ {
    for c := uint(0); c < x.nc; c++ {
      x.matrix[l][c].Sub (y.matrix[l][c])
    }
  }
}

func (x *qmatrix) Sub (Y ...Adder) {
  for i, _ := range Y {
    x.sub (Y[i])
  }
}

func (x *qmatrix) One() bool {
  if x.nl != x.nc {
    return false
  }
  for l := uint(0); l < x.nl; l++ {
    for c := uint(0); c < x.nc; c++ {
      if l == c {
        if ! x.matrix[l][c].One() {
          return false
        }
      } else {
        if ! x.matrix[l][c].Zero() {
          return false
        }
      }
    }
  }
  return true
}

func (x *qmatrix) mul (Y Multiplier) {
  y := x.imp(Y)
  if y.nl != x.nc || y.d != x.d { nofit() }
  a := q.New()
  b := q.New()
  xy := New (x.nl, y.nc, x.d).(*qmatrix)
  for l := uint(0); l < xy.nl; l++ {
    for c := uint(0); c < xy.nc; c++ {
      a.Set1 (0)
      for i := uint(0); i < x.nc; i++ {
        b.Copy (x.matrix[l][i])
        b.Mul (y.matrix[i][c])
        a.Add (b)
      }
      xy.matrix[l][c].Copy (a)
    }
  }
  if xy.nl == xy.nc {
    x.Copy (xy)
  } else {
    ker.Panic ("not yet implemented")
  }
}

func (x *qmatrix) Mul (Y ...Multiplier) {
  for i, _ := range Y {
    x.mul (Y[i])
  }
}

func (x *qmatrix) Prod (Y, Z Multiplier) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.mul (z)
}

// x and y are changed in a way s.t.
// x.matrix[m][m] is set to 1.
func (x *qmatrix) norm (y *qmatrix, m uint) {
  a := x.matrix[m][m].Clone().(q.Rational)
  for c := uint(0); c < x.nc; c++ {
    x.matrix[m][c].DivBy (a)
    y.matrix[m][c].DivBy (a)
  }
}

// x and y are changed in a way s.t.
// x.matrix[l][0] is set to 0 for all lines > m.
func (x *qmatrix) do1 (y *qmatrix, m uint) {
  for l := m + 1; l < x.nc; l++ {
    z := x.matrix[l][m].Clone().(q.Rational)
    for c := uint(0); c < x.nc; c++ {
      a := z.Clone().(q.Rational)
      a1 := x.matrix[m][c].Clone().(q.Rational)
      a1.Mul (a)
      x.matrix[l][c].Sub (a1)
      b1 := y.matrix[m][c].Clone().(q.Rational)
      b1.Mul (a)
      y.matrix[l][c].Sub (b1)
    }
  }
}

// x and y are changed in a way s.t.
// x.matrix[l][m] is set to 0 for all lines < m;
func (x *qmatrix) do2 (y *qmatrix, m uint) {
  for l := uint(0); l < m; l++ {
    a := x.matrix[l][m].Clone().(q.Rational)
    for c := uint(0); c < x.nc; c++ {
      a1 := x.matrix[m][c].Clone().(q.Rational)
      a1.Mul (a)
      x.matrix[l][c].Sub (a1)
      b1 := y.matrix[m][c].Clone().(q.Rational)
      b1.Mul (a)
      y.matrix[l][c].Sub (b1)
    }
  }
}

func (x *qmatrix) Invertible() bool {
  d := x.Det()
  _, n, _ := d.Vals()
  return n != 0
}

func (x *qmatrix) repair() {
  zero := make([]bool, x.nl)
  for l := uint(0); l < x.nl; l++ {
    zero[l] = x.matrix[l][0].Zero()
  }
  if zero[0] {
    loop:
    for l := uint(1); l < x.nl; l++ {
      if ! zero[l] {
        for c := uint(0); c < x.nc; c++ {
          x.matrix[0][c], x.matrix[l][c] = x.matrix[l][c], x.matrix[0][c]
        }
        break loop
      }
    }
  }
}

func wait() {
  errh.Error0 ("")
}

func (x *qmatrix) DivBy (Y Multiplier) {
  y := x.imp(Y)
  if ! y.Invertible() {
    DivBy0Panic()
  }
  if y.matrix[0][0].Zero() {
    y.repair()
  }
  y.norm (x, 0)
  withProtocol := false
  c := (x.nc + 1) * (2 * (x.d + 1) + 1)
  if withProtocol { y.Write (0, 0); x.Write (0, c); wait() }
  for i := uint(0); i < x.nc - 1; i++ {
    y.do1 (x, i)
    if withProtocol { y.Write (0, 0); x.Write (0, c); wait() }
    y.norm (x, i + 1)
    if withProtocol { y.Write (0, 0); x.Write (0, c); wait() }
  }
  for i := x.nc - 1; i > 0; i-- {
    y.do2 (x, i)
    if withProtocol { y.Write (0, 0); x.Write (0, c); wait() }
  }
}

func (x *qmatrix) Invert() {
  e := unit (x.nl, x.nc, x.d).(*qmatrix)
  e.DivBy (x)
  x.Copy (e)
}

func (x *qmatrix) Quot (Y, Z Multiplier) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.DivBy (z)
}

func (x *qmatrix) Sqr() {
  if x.nl != x.nc { nofit() }
  y := x.Clone().(*qmatrix)
  x.Mul (y)
}

func (x *qmatrix) Power (n uint) {
  if x.nl != x.nc { nofit() }
  y := x.Clone().(*qmatrix)
  for i := uint(1); i < n; i++ {
    x.Mul (y)
  }
}

func (x *qmatrix) Set1 (i ...int) {
  if uint(len(i)) != x.nl * x.nc { nofit() }
  for l := uint(0); l < x.nl; l++ {
    for c := uint(0); c < x.nc; c++ {
      x.matrix[l][c].Set (i[x.nl * l + c], 1)
    }
  }
  if x.matrix[0][0].Zero() {
    x.repair()
  }
}

func (x *qmatrix) Set (i ...int) {
  if uint(len(i)) != 2 * x.nl * x.nc { nofit() }
  k := uint(0)
  for l := uint(0); l < x.nl; l++ {
    for c := uint(0); c < x.nc; c++ {
      x.matrix[l][c].Set (i[k], i[k + 1])
      k += 2
    }
  }
  if x.matrix[0][0].Zero() {
    x.repair()
  }
}

func (x *qmatrix) Vals (l, c uint) (bool, uint, uint) {
  if l >= x.nl || c >= x.nc { nofit() }
  return x.matrix[l][c].Vals()
}

func (x *qmatrix) Write (z, s uint) {
  for l := uint(0); l < x.nl; l++ {
    for c := uint(0); c < x.nc; c++ {
      x.matrix[l][c].Write (z + l, c + s + (1 + x.d + 1 + x.d) * c)
    }
  }
}

func (x *qmatrix) Edit (z, s uint) {
  for l := uint(0); l < x.nl; l++ {
    for c := uint(0); c < x.nc; c++ {
      x.matrix[l][c].Edit (z + l, c + s + (1 + x.d + 1 + x.d) * c)
    }
  }
}

func (x *qmatrix) TeX() string { // AmSTeX
  const r = "\\"
  s := "$" + r + "pmatrix\n"
  for l := uint(0); l < x.nl; l++ {
    v, a, b := x.matrix[l][0].Vals()
    if ! v {
      s += "-"
    }
    s += n.String(a) + "/" + n.String(b)
    for c := uint(1); c < x.nc; c++ {
      v, a, b := x.matrix[l][c].Vals()
      s += "&"
      if ! v {
        s += "-"
      }
      s +=  n.String(a) + "/" + n.String(b)
    }
    if l + 1 < x.nl {
      s += r + r
    }
    s += "\n"
  }
  return s + r + "endpmatrix$\n"
}
