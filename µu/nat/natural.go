package nat

// (c) Christian Maurer   v. 150215 - license see µu.go

import (
  . "µu/ker"
  "µu/str"
  "µu/col"
//  "µu/scr"
  "µu/box"
)
const (
  max = 10 // maximal number of digits of uint
)
var (
  bx = box.New()
  width uint
)

func init() {
//  bx.SetNumerical() // TODO
}

func isDigit (b byte) bool {
  return '0' <= b && b <= '9'
}

func digSeqs (s string) (uint, []string, []uint, []uint) {
  var t []string
  var p, length []uint
  l := uint(len (s))
  noDigitBefore := true
  n := uint(0)
  for i := uint(0); i < l; i++ {
    if isDigit (s[i]) {
      if noDigitBefore {
        t = append (t, string(s[i]))
        p = append (p, i)
        length = append (length, 1)
        n ++
        noDigitBefore = false
      } else {
        t[n - 1] += string (s[i])
        length[n - 1] ++
      }
    } else {
      noDigitBefore = true
    }
  }
  return n, t, p, length
}

func defined (n *uint, s string) bool {
  if s == "" { return false }
  str.Move (&s, true)
  l := str.ProperLen (s)
  *n = uint(0)
  var b byte
  for i := 0; i < int(l); i++ {
    if isDigit (s[i]) {
      b = s[i] - '0'
      if *n <= (MaxNat() - uint(b)) / 10 {
        *n = 10 * *n + uint(b)
      } else {
        return false
      }
    } else {
      return false
    }
  }
  return true
}

func natural (s string) (uint, bool) {
  var n uint
  if defined (&n, s) {
    return n, true
  }
  return 0, false
}

func wdrec (n uint) uint {
  if n > 0 {
    return 1 + wdrec (n / 10)
  }
  return 0
}

func wd (n uint) uint {
  if n == 0 {
    return 1
  }
  return wdrec (n)
}

func string_ (n uint) string {
  if n == 0 { return "0" }
  var s string
  for s = ""; n > 0; n /= 10 {
    s = string(n % 10 + '0') + s
  }
  return s
}

func stringFmt (n uint, l uint, withZeros bool) string {
  s := string_(n)
  a := " "; if withZeros { a = "0" }
  w := wd(n)
  if l < w { l = w }
  for ; w < l; w++ {
    s = a + s
  }
  return s
}

func colours (f, b col.Colour) {
  bx.Colours (f, b)
}

func write (n uint, l, c uint) {
  w := wd (n)
  if width > w { w = width }
  if w > c + 1 { return }
  bx.Wd (w)
//  scr.SwitchFontsize (scr.Normal)
  bx.Write (StringFmt (n, w, false), l, c + 1 - w)
}

func setWd (w uint) {
  if w == 0 {
    width = 1
  } else if w > max {
    width = max
  } else {
    width = w
  }
}

func edit (n *uint, l, c uint) {
  w := wd (*n)
  if width > w { w = width }
  bx.Wd (w)
  s := String (*n)
  for {
    bx.Edit (&s, l, c + 1 - w)
    if defined (n, s) {
      break
    } else {
//      errh.Report2Error ("keine Zahl", 0) // impossible, dependency cycle
//      scr.Write (" keine Zahl ", scr.NLines() - 1, 0) // provisorial
    }
  }
//  scr.Write ("            ", scr.NLines() - 1, 0) // provisorial
//  bx.Write (String (*n), l, c + 1 - w)
}
