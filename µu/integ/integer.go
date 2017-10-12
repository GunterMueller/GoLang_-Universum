package integ

// (c) Christian Maurer   v. 170107 - license see µu.go

import (
  . "µu/ker"
  "µu/str"
  "µu/col"
  "µu/box"
  "µu/errh"
  "µu/nat"
)
const (
  max = 11 // sign plus maximal 10 digits
)
var (
  m2 = uint(-MinInt()) // == uint(MaxInt() + 1)
  bx = box.New()
  width uint
)

func init() {
//  bx.SetNumerical() // TODO
}

func defined (z *int, s string) bool {
  str.Move (&s, true)
  negative:= s[0] == '-'
  if negative {
    n:= uint(len (s))
    s = str.Part (s, 1, n - 1)
  }
  if str.Empty (s) {
    return false
  }
  if n, ok:= nat.Natural (s); ok {
    if negative {
      if n < m2 {
        *z = - int(n)
        return true
      } else if n == m2 {
        *z = MinInt()
        return true
      }
    } else if n <= uint(MaxInt()) {
      *z = int(n)
      return true
    }
  }
  return false
}

func integer (s string) (int, bool) {
  var n int
  if defined (&n, s) {
    return n, true
  }
  return 0, false
}

func wd (z int) uint {
  if z < 0 { z = - z }
  return 1 + nat.Wd (uint(z))
}

func string_(z int) string {
  s:= ""
  if z < 0 {
    s = "-"
    z = -z
  }
  return s + nat.String (uint(z))
}

func stringFmt (z int, l uint) string {
  a:= " "; if z < 0 { a = "-"; z = -z }
  w:= Wd (z)
  if l < w { l = w }
  return a + nat.StringFmt (uint(z), l - 1, false)
}

func colours (f, b col.Colour) {
  bx.Colours (f, b)
}

func write (z int, l, c uint) {
  w:= Wd (z)
  if w > c + 1 { return }
  bx.Wd (w)
//  scr.SwitchFontsize (scr.Normal)
  bx.Write (StringFmt (z, w), l, c + 1 - w)
}

func setWd (w uint) {
  if w == 0 {
    width = 2
  } else if w > max {
    width = max
  } else {
    width = w
  }
}

func edit (z *int, l, c uint) {
  w:= Wd (*z)
  if width > w { w = width }
  bx.Wd (w)
  s:= StringFmt (*z, w)
  for {
    bx.Edit (&s, l, c + 1 - w)
    if defined (z, s) {
      break
    } else {
      errh.Error0("keine Zahl") // , l + 1, c)
    }
  }
  bx.Write (StringFmt (*z, w), l, c + 1 - w)
}
