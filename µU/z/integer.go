package z

// (c) Christian Maurer   v. 201226 - license see µU.go

import (
  . "µU/ker"
  "µU/str"
  "µU/col"
  "µU/box"
  "µU/errh"
  "µU/n"
)
const
  max = 11 // sign plus maximal 10 digits
var (
  m2 = uint(-MinInt) // == uint(MaxInt() + 1)
  bx = box.New()
  width uint
)

func init() {
  bx.SetNumerical()
}

func integer (s string) (int, bool) {
  str.Move (&s, true)
  l := str.ProperLen(s)
  s = s[:l]
  negative := s[0] == '-'
  if negative {
    s = s[1:]
  }
  if str.Empty (s) {
    return 0, false
  }
  if k, ok := n.Natural (s); ok {
    if negative {
      if k < m2 {
        return -int(k), true
      } else if k == m2 {
        return MinInt, true
      }
    } else if k <= uint(MaxInt) {
      return int(k), true
    }
  }
  return 0, false
}

func wd (z int) uint {
  if z < 0 { z = - z }
  return 1 + n.Wd (uint(z))
}

func string_(z int) string {
  s := ""
  if z < 0 {
    s = "-"
    z = -z
  }
  return s + n.String (uint(z))
}

func stringFmt (z int, l uint) string {
  a := " "; if z < 0 { a = "-"; z = -z }
  w := Wd (z)
  if l < w { l = w }
  return a + n.StringFmt (uint(z), l - 1, false)
}

func colours (f, b col.Colour) {
  bx.Colours (f, b)
}

func write (z int, l, c uint) {
  w := Wd (z)
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
  w := Wd (*z)
  if width > w { w = width }
  bx.Wd (w)
  s := StringFmt (*z, w)
  ok := false
  for {
    bx.Edit (&s, l, c + 1 - w)
    if *z, ok = integer (s); ok {
      break
    } else {
      errh.Error0("keine Zahl") // , l + 1, c)
    }
  }
  bx.Write (StringFmt (*z, w), l, c + 1 - w)
}
