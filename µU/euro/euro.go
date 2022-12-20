package euro

// (c) Christian Maurer   v. 221213 - license see µU.go

import (
  "math"
  "µU/ker"
  . "µU/obj"
  "µU/str"
  "µU/col"
  "µU/box"
  "µU/errh"
  "µU/N"
  "µU/font"
  "µU/pbox"
)
const (
  undefined = uint(Limit * 100)
  nDigits = 7 // Limit - 1
  length = nDigits + 1 + 2 // 1 for dot or comma, 2 for cents
)
type (
  euro struct {
         cent uint
         f, b col.Colour
              font.Font
              }
)
var (
  bx = box.New()
  pbx = pbox.New()
)

func init() {
  bx.Wd (length)
//  bx.SetNumerical()
}

func new_() Euro {
  x := new(euro)
  x.Clr()
  x.f, x.b = col.StartCols()
  return x
}

func new2 (e, c uint) Euro {
  if e >= Limit || c >= 100 {
    ker.PrePanic()
  }
  x := new_()
  x.SetVal2 (e, c)
  return x
}

func (x *euro) imp (Y any) *euro {
  y, ok := Y.(*euro)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *euro) Empty() bool {
  return x.cent >= undefined
}

func (x *euro) Clr() {
  x.cent = undefined
}

func (x *euro) Copy (Y any) {
  x.cent = x.imp(Y).cent
}

func (x *euro) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *euro) Eq (Y any) bool {
  return x.cent == x.imp(Y).cent
}

func (x *euro) Less (Y any) bool {
  y := x.imp(Y)
  if y.cent == undefined {
    return false
  }
  if x.cent == undefined { // y.cent != undefined
    return true
  }
  return x.cent < y.cent
}

func (x *euro) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *euro) Val() uint {
  return x.cent
}

func (x *euro) SetVal (c uint) {
  if c >= undefined {
    c = undefined
  }
  x.cent = c
}

func (x *euro) RealVal() float64 {
  return float64 (x.cent) / 100
}

func (x *euro) SetRealVal (r float64) {
  if r < 0 || r >= float64(Limit) {
    ker.PrePanic()
  }
  x.cent = uint(100 * r)
}

func (x *euro) Val2() (uint, uint) {
  return x.cent / 100, x.cent % 100
}

func (x *euro) SetVal2 (e, c uint) {
  if c >= 100 || e >= undefined {
    ker.PrePanic()
  }
  x.cent = 100 * e
  x.cent += c
}

func (x *euro) Zero() bool {
  return x.cent == 0
}

func (x *euro) Add (Y ...Adder) {
  if x.cent == undefined { return }
  for _, y := range Y {
    if y.(*euro).cent == undefined { return }
  }
  for _, y := range Y {
    x.cent += y.(*euro).cent
    if x.cent >= undefined {
      x.cent = undefined
      break
    }
  }
}

func (x *euro) Sum (Y, Z Adder) {
  y, z := x.imp (Y), x.imp (Z)
  if y.cent >= undefined || z.cent >= undefined {
    return
  }
  x.cent += y.cent + z.cent
  if x.cent >= undefined {
    x.cent = undefined
  }
}

func (x *euro) Sub (Y ...Adder) {
  if x.cent == undefined { return }
  for _, y := range Y {
    if y.(*euro).cent == undefined { return }
  }
  for _, y := range Y {
    if x.cent >= y.(*euro).cent {
      x.cent -= y.(*euro).cent
    } else {
      x.cent = undefined
      break
    }
  }
}

func (x *euro) Diff (Y, Z Adder) {
  y, z := x.imp (Y), x.imp (Z)
  if y.cent == undefined || z.cent == undefined {
    return
  }
  x.cent += y.cent - z.cent
}

func (x *euro) Operate (Faktor, Divisor uint) {
  if x.cent == undefined { return }
  if Divisor == 0 { x.cent = undefined; return }
  if Faktor == 0 { x.cent = 0; return }
  if x.cent / Divisor < undefined / Faktor {
    x.cent *= Faktor
    x.cent += Divisor / 2
    x.cent /= Divisor
  } else {
    x.cent = undefined
  }
}

func toThe (q float64, n uint) float64 {
  if n == 0 {
    return 1.
  }
  return q * toThe (q, n - 1)
}

func (x *euro) ChargeInterest (p, n uint) {
  if x.cent == undefined { return }
  f := toThe (1.0 + float64(p) / 10000., n)
  b := float64 (x.cent) * f + 0.5
  if b < float64(undefined) {
    x.cent = uint(math.Trunc (b))
  } else {
    x.cent = undefined
  }
}

func (x *euro) Round (Y Euro) {
  yc := x.imp(Y).cent
  if x.cent >= undefined || yc >= undefined { return }
  x.cent = yc * (x.cent / yc)
}

func (x *euro) String() string {
  if x.Empty() { return str.New (length) }
  return N.StringFmt (x.cent / 100, nDigits, false) + "," +
         N.StringFmt (x.cent % 100, 2, true)
}

func (x *euro) Defined (s string) bool {
  if str.Empty (s) {
    x.cent = undefined
    return true
  }
  a, t, P, L := N.DigitSequences (s)
  if len(t) == 0 { return false }
  k, hatKomma := str.Pos (s, ',')
  if ! hatKomma {
    k, hatKomma = str.Pos (s, '.')
  }
  i, ok := N.Natural (t[0])
  if ! ok { return false }
  switch a {
  case 1:
    if hatKomma && k < P[0] { // Komma vor der Ziffernfolge
      switch L[0] {
      case 1:
        x.cent = 10 * i
      case 2:
        x.cent = i
      default:
        return false
      }
      return true
    }
    if hatKomma && k >= P[0] + L[0] || ! hatKomma {
      if L[0] <= nDigits {
        x.cent = 100 * i
        return true
      }
    }
  case 2:
    if ! hatKomma { return false }
    if k < P[0] + L[0] || P[1] <= k { return false }
    if L[0] > nDigits {
      return false
    } else {
      x.cent = 100 * i
    }
    if i, ok = N.Natural (t[1]); ! ok { return false }
    switch L[1] {
    case 1:
      x.cent += 10 * i
    case 2:
      x.cent += i
    default:
      return false
    }
    return true
  }
  return false
}

func (x *euro) Colours (f, b col.Colour) {
  x.f, x.b = f, b
}

func (x *euro) Cols() (col.Colour, col.Colour) {
  return x.f, x.b
}

func (x *euro) Write (l, c uint) {
  bx.Colours (x.f, x.b)
  bx.Write (x.String(), l, c)
}

func (x *euro) Edit (l, c uint) {
  s := x.String()
  bx.Colours (x.f, x.b)
  for {
    bx.Edit (&s, l, c)
    if x.Defined (s) {
      break
    } else {
      errh.Error0 ("kein Geldbetrag") // l + 1, c)
    }
  }
  x.Write (l, c)
}

func (x *euro) SetFont (f font.Font) {
  x.Font = f
}

func (x *euro) Print (l, c uint) {
  pbx.SetFont (x.Font)
  pbx.Print (x.String(), l, c)
}

func (x *euro) Codelen() uint {
  return 4 // Codelen (uint32(0))
}

func (x *euro) Encode() Stream {
  s := make(Stream, 4)
  s = Encode (uint32(x.cent))
  return s
}

func (x *euro) Decode (s Stream) {
  x.cent = uint(Decode (uint32(0), s).(uint32))
}
