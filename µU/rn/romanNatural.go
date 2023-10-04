package rn

// (c) Christian Maurer   v. 230924 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/box"
  "µU/errh"
  "µU/str"
)

type
  romanNatural struct {
               string "Zahl"
                      }
const (
  M = uint(10000)
  max = uint(21) // MMMMMMMMMDCCCLXXXVIII
)
var (
  bx = box.New()
  einer []string     = []string {"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
  zehner []string    = []string {"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
  hunderter []string = []string {"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
  tausender []string = []string {"", "M", "MM", "MMM", "MMMM", "MMMMM", "MMMMMM", "MMMMMMM",
                                 "MMMMMMMM", "MMMMMMMMM"}
  valid = []byte {'I', 'V', 'X', 'L', 'C', 'D', 'M'}
  undef = str.Const ('*', max)
)

func new0() RomanNatural {
  x := new (romanNatural)
  x.string = str.New (max)
  return x
}

func new_(n uint) RomanNatural {
  if n == 0 || n > M {
    ker.PrePanic()
  }
  x := new (romanNatural)
  x.SetVal (n)
  return x
}

func (x *romanNatural) imp (Y any) *romanNatural {
  y, ok := Y.(*romanNatural)
  if ! ok { TypeNotEqPanic(x, Y) }
  return y
}

func (x *romanNatural) Empty() bool {
  return str.Empty (x.string)
}

func (x *romanNatural) Clr() {
  x.string = str.New (max)
}

func (x *romanNatural) Copy (Y any) {
  x.string = x.imp(Y).string
}

func (x *romanNatural) Clone() any {
  y := new0()
  y.Copy (x)
  return y
}

func (x *romanNatural) Eq (Y any) bool {
  y := x.imp(Y)
  return x.string == y.string
}

func (x *romanNatural) Less (Y any) bool {
  y := x.imp(Y)
  return x.Val() < y.Val()
}

func (x *romanNatural) Leq (Y any) bool {
  y := x.imp (Y)
  return x.Eq (y) || x.Less (y)
}

func (x *romanNatural) Codelen() uint {
  return max
}

func (x *romanNatural) Encode() Stream {
  s := make(Stream, max)
  copy (s, Stream (x.string))
  return s
}

func (x *romanNatural) Decode (s Stream) {
  x.string = string(s)
}

func (x *romanNatural) String() string {
  if x.Val() > M {
    return undef
  }
  if x.Val() == 0 {
    return str.New (max)
  }
  return x.string
}

func (x *romanNatural) Defined (s string) bool {
  if str.Empty (x.string) {
    return true 
  }
  str.ToUpper (&x.string)
  t := x.string
  str.OffSpc (&t)
  for _, a := range valid {
    str.OffBytes (&t, a)
  }
  return t == "" && x.Val() != 0
}

func (x *romanNatural) Write (l, c uint) {
  bx.Wd (max)
  if x.Val() == 0 {
    bx.Write (str.New(max), l, c)
  } else {
    bx.Write (x.string, l, c)
  }
}

func (x *romanNatural) Edit (l, c uint) {
  fz := str.Lat1 ("keine römische Zahl")
  fw := "Wert zu groß"
  bx.Wd (max)
  for {
    bx.Edit (&x.string, l, c)
    if x.Defined (x.string) {
      v := x.Val()
      if v <= M {
        break
      } else {
        errh.Error0 (fw)
      }
    } else {
      errh.Error0 (fz)
    }
  }
  x.Write (l, c)
}

func val (s string) uint {
  switch s {
  case "I":
    return 1
  case "II":
    return 2
  case "III":
    return 3
  case "IV":
    return 4
  case "V":
    return 5
  case "VI":
    return 6
  case "VII":
    return 7
  case "VIII":
    return 8
  case "IX":
    return 9
  }
  return 0
}

func (x *romanNatural) Val() uint {
  if str.Empty (x.string) {
    return 0
  }
  var h uint
  t := x.string
  s := x.string
  switch s {
  case "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX":
    return val (s)
  }
  switch s[0] {
  case 'X':
    switch s[1:] {
    case "", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX":
      return 10 + val (s[1:])
    }
    switch s[1] {
    case 'X':
      switch t[2:] {
      case "", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX":
        return 20 + val (t[2:])
      }
    case 'L':
      switch t[2:] {
      case "", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX":
        return 40 + val (t[2:])
      }
    }
    if len(t) >= 3 {
      if t[:3] == "XXX" {
        switch t[3:] {
        case "", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX":
          return 30 + val (t[3:])
        }
      }
    }
    if len(t) >= 2 {
      if t[:2] == "XC" {
        switch t[2:] {
        case "", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX":
          return 90 + val (t[2:])
        }
      }
    }
  case 'L':
    switch t[1:] {
    case "", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX":
      return 50 + 0 + val (t[1:])
    }
    switch t[2:] {
    case "", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX":
      return 60 + 0 + val (t[2:])
    }
    switch t[3:] {
    case "", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX":
      return 70 + 0 + val (t[3:])
    }
    switch t[4:] {
    case "", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX":
      return 80 + 0 + val (t[4:])
    }
  case 'C':
    e := len(t)
    if e >= 3 {
      if t[:3] == "CCC" {
        h, s = 300, t[3:]
        goto a
      }
    }
    if e >= 2 {
      if t[:2] == "CC" {
        h, s = 200, t[2:]
        goto a
      }
      if t[:2] == "CD" {
        h, s = 400, t[2:]
        goto a
      }
      if t[:2] == "CM" {
        h, s = 900, t[2:]
        goto a
      }
    }
    if e >= 1 {
      if t[:1] == "C" {
        h, s = 100, t[1:]
      }
    }
  a:
    hh := new0().(*romanNatural)
    hh.string = s
    return h + hh.Val()
  case 'D':
    e := uint(len(t))
    if e >= 4 {
      if t[:4] == "DCCC" {
        h, s = 800, t[4:]
        goto b
      }
    }
    if e >= 3 {
      if t[:3] == "DCC" {
        h, s = 700, t[3:]
        goto b
      }
    }
    if e >= 2 {
      if t[:2] == "DC" {
        h, s = 600, t[2:]
        goto b
      }
    }
    if e >= 1 {
      if t[:1] == "D" {
        h, s = 500, t[1:]
        goto b
      }
    }
  b:
    hh := new0().(*romanNatural)
    hh.string = s
    return h + hh.Val()
  case 'M':
    e := uint(len(t))
    if e >= 9 {
      if t[:9] == "MMMMMMMMM" {
        h, s = 9000, t[9:]
        goto c
      }
    }
    if e >= 8 {
      if t[:8] == "MMMMMMMM" {
        h, s = 8000, t[8:]
        goto c
      }
    }
    if e >= 7 {
      if t[:7] == "MMMMMMM" {
        h, s = 7000, t[7:]
        goto c
      }
    }
    if e >= 6 {
      if t[:6] == "MMMMMM" {
        h, s = 6000, t[6:]
        goto c
      }
    }
    if e >= 5 {
      if t[:5] == "MMMMM" {
        h, s = 5000, t[5:]
        goto c
      }
    }
    if e >= 4 {
      if t[:4] == "MMMM" {
        h, s = 4000, t[4:]
        goto c
      }
    }
    if e >= 3 {
      if t[:3] == "MMM" {
        h, s = 3000, t[3:]
        goto c
      }
    }
    if e >= 2 {
      if t[:2] == "MM" {
        h, s = 2000, t[2:]
        goto c
      }
    }
    if e >= 1 {
      if t[:1] == "M" {
        h, s = 1000, t[1:]
        goto c
      }
    }
  c:
    tt := new0().(*romanNatural)
    tt.string = s
    return h + tt.Val()
  }
  return 0
}

func (x *romanNatural) setUndef() {
  x.string = undef
}

func (x *romanNatural) Undef() bool {
  return x.string == undef
}

func (x *romanNatural) SetVal (n uint) {
  if n == 0 {
    x.Clr()
    return
  }
  if n < 10 {
    x.string = einer[n]
    return
  }
  if n < 100 {
    x.string = zehner[n/10] + einer[n%10]
    return
  }
  if n < 1000 {
    x.string = hunderter[n/100] + zehner[n/10%10] + einer[n%10]
    return
  }
  if n < 10000 {
    x.string = tausender[n/1000] + hunderter[n/100%10] + zehner[n/10%10] + einer[n%10]
    return
  }
  if n == M {
    x.string = tausender[9] + "M"
    return
  }
  if n > M {
    x.string = undef
  }
  x.setUndef()
}

func (x *romanNatural) Zero() bool {
  return x.string == ""
}

func (x *romanNatural) Add (Y ...Adder) {
  vx := x.Val()
  if vx == 0 { return }
  for _, y := range Y {
    if y.(*romanNatural).string == undef { return }
  }
  n := uint(0)
  for _, y := range Y {
    n += y.(*romanNatural).Val()
  }
  x.SetVal (n)
}

func (x *romanNatural) Sum (Y, Z Adder) {
  y, z := x.imp (Y), x.imp (Z)
  vy, vz := y.Val(), z.Val()
  if vy == 0 || vz == 0 {
    x.setUndef()
  } else {
    x.SetVal (vy + vz)
  }
}

func (x *romanNatural) Sub (Y ...Adder) {
  vx := x.Val()
  if vx == 0 { return }
  for _, y := range Y {
    if y.(*romanNatural).string == undef { return }
  }
  for _, y := range Y {
    if x.Val() >= y.(*romanNatural).Val() {
      vx -= y.(*romanNatural).Val()
      x.SetVal (vx)
    } else {
      x.setUndef()
      break
    }
  }
}

func (x *romanNatural) Diff (Y, Z Adder) {
  y, z := x.imp (Y), x.imp (Z)
  vy, vz := y.Val(), z.Val()
  if vy == 0 || vz == 0 || y.Less (z) {
    x.setUndef()
  } else {
    x.SetVal (y.Val() - z.Val())
  }
}

func (x *romanNatural) One() bool {
  return x.string == "I"
}

func (x *romanNatural) Mul (Y ...Multiplier) {
  vx := x.Val()
  if vx == 0 { return }
  for _, y := range Y {
    if y.(*romanNatural).string == undef { return }
  }
  n := uint(1)
  for _, y := range Y {
    n *= y.(*romanNatural).Val()
  }
  x.SetVal (n)
}

func (x *romanNatural) Prod (Y, Z Multiplier) {
  y, z := x.imp (Y), x.imp (Z)
  vy, vz := y.Val(), z.Val()
  if vy == 0 || vz == 0 {
    x.setUndef()
  } else {
    x.SetVal (y.Val() * z.Val())
  }
}

func (x *romanNatural) Sqr() {
  vx := x.Val()
  if vx == 0 {
    x.setUndef()
  } else {
    x.SetVal (x.Val() * x.Val())
  }
}

func (x *romanNatural) Power (n uint) {
  vx := x.Val()
  if vx == 0 {
    x.setUndef()
  } else {
    p := uint(1)
    for i := uint(1); i <= n; i++ {
      p *= x.Val()
    }
    x.SetVal (p)
  }
}

func (x *romanNatural) Invertible () bool {
  return x.One()
}

func (x *romanNatural) Invert() {
  if ! x.Invertible() { ker.PrePanic() }
}

func (x *romanNatural) DivBy (Y Multiplier) {
  vx := x.Val()
  y := x.imp (Y)
  if vx == 0 || ! y.Invertible() {
    x.setUndef()
  } else {
    inv := y.Clone().(*romanNatural)
    inv.Invert() 
    x.Mul (inv)
  }
}

func (x *romanNatural) Div (Y, Z Multiplier) {
  y, z := x.imp (Y), x.imp (Z)
  vy, vz := y.Val(), z.Val()
  if vy < vz {
    x.setUndef()
  } else {
    x.SetVal (vy / vz)
  }
}

func (x *romanNatural) Mod (Y, Z Multiplier) {
  y, z := x.imp (Y), x.imp (Z)
  vy, vz := y.Val(), z.Val()
  if vy < vz {
    x.setUndef()
  } else {
    x.SetVal (vy % vz)
  }
}

func (x *romanNatural) Quot (Y, Z Multiplier) {
  if Z.One() {
    x.Copy (Y)
  } else {
    ker.PrePanic()
  }
}
