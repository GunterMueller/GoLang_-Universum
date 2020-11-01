package euro

// (c) Christian Maurer   v. 201014 - license see µU.go

import (
  "math"
  . "µU/obj"
  "µU/str"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/errh"
  "µU/nat"
  "µU/font"
  "µU/pbox"
)
const (
  hundred = uint(100)
  tenMillions = uint(1e7)
  undefined = uint(tenMillions * hundred)
  nDigits = 7 // höchstens 9.999.999 Euro
  length = nDigits + 1 /* Komma */ + 2
)
type (
  euro struct {
         cent uint
       cF, cB col.Colour
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
  x.cF, x.cB = scr.StartCols()
  return x
}

func new2 (e, c uint) Euro {
  x := new(euro)
  x.Set2 (e,c)
  return x
}

func (x *euro) imp (Y Any) *euro {
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

func (x *euro) Copy (Y Any) {
  x.cent = x.imp(Y).cent
}

func (x *euro) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *euro) Eq (Y Any) bool {
  return x.cent == x.imp(Y).cent
}

func (x *euro) Less (Y Any) bool {
  y := x.imp(Y)
  if y.cent == undefined {
    return false
  }
  if x.cent == undefined { // y.cent != undefined
    return true
  }
  return x.cent < y.cent
}

func (x *euro) Val() uint {
  return x.cent
}

func (x *euro) SetVal (c uint) bool {
  x.cent = c
  return x.cent < undefined
}

func (x *euro) Val2() (uint, uint) {
  return x.cent / hundred, x.cent % hundred
}

func (x *euro) Set2 (e, c uint) bool {
  if e >= tenMillions || c >= hundred {
    x.cent = undefined
  } else {
    x.cent = hundred * e
    x.cent += c
  }
  return x.cent < undefined && c < hundred
}

func (x *euro) RealVal() float64 {
  return float64 (x.cent) / float64 (hundred)
}

func (x *euro) SetReal (r float64) bool {
  if r >= 0. && r < float64(tenMillions) {
    x.cent = uint(math.Trunc (float64 (hundred) * r + 0.5))
  } else {
    x.cent = undefined
  }
  return x.cent < undefined
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

func (x *euro) Operate (Faktor, Divisor uint) {
  if x.cent == undefined { return }
  if Divisor == 0 { x.cent = undefined; return }
  if Faktor == 0 { x.cent = 0; return }
  if x.cent / Divisor < (tenMillions * hundred) / Faktor {
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
  f := toThe (1.0 + float64 (p) / 10000.0, n)
  b := float64 (x.cent) * f + 0.5
  if b < float64 (tenMillions * hundred) {
    x.cent = uint(math.Trunc (b))
  } else {
    x.cent = undefined
  }
}

func (x *euro) Round (Y Euro) {
  yc := x.imp(Y).cent
  if x.cent == undefined || yc == undefined { return }
  x.cent = yc * (x.cent / yc)
}

func (x *euro) String() string {
  if x.Empty() { return str.New (length) }
  return nat.StringFmt (x.cent / hundred, nDigits, false) + "," +
         nat.StringFmt (x.cent % hundred, 2, true)
}

func (x *euro) Defined (s string) bool {
  if str.Empty (s) {
    x.cent = undefined
    return true
  }
  a, t, P, L := nat.DigitSequences (s)
  if len(t) == 0 { return false }
  k, hatKomma := str.Pos (s, ',')
  if ! hatKomma {
    k, hatKomma = str.Pos (s, '.')
  }
  n, ok := nat.Natural (t[0])
  if ! ok { return false }
  switch a {
  case 1:
    if hatKomma && k < P[0] { // Komma vor der Ziffernfolge
      switch L[0] {
      case 1:
        x.cent = 10 * n
      case 2:
        x.cent = n
      default:
        return false
      }
      return true
    }
    if hatKomma && k >= P[0] + L[0] || ! hatKomma {
      if L[0] <= nDigits {
        x.cent = hundred * n
        return true
      }
    }
  case 2:
    if ! hatKomma { return false }
    if k < P[0] + L[0] || P[1] <= k { return false }
    if L[0] > nDigits {
      return false
    } else {
      x.cent = hundred * n
    }
    if n, ok = nat.Natural (t[1]); ! ok { return false }
    switch L[1] {
    case 1:
      x.cent += 10 * n
    case 2:
      x.cent += n
    default:
      return false
    }
    return true
  }
  return false
}

func (x *euro) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *euro) Write (l, c uint) {
  bx.Colours (x.cF, x.cB)
  bx.Write (x.String(), l, c)
}

func (x *euro) Edit (l, c uint) {
  s := x.String()
  bx.Colours (x.cF, x.cB)
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
  bs := make (Stream, 4)
  bs = Encode (uint32(x.cent))
  return bs
}

func (x *euro) Decode (bs Stream) {
  x.cent = uint(Decode (uint32(0), bs).(uint32))
}
