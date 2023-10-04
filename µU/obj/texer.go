package obj

// (c) Christian Maurer   v. 230812 - license see µU.go

import (
  "µU/char"
  "µU/str"
)
type
  TeXer interface {

// Returns an AMSTeX-string representation of x.
  TeX() string
}

func String2TeX (s string) string {
  var b byte
  n := uint(len(s))
  t := ""
  for i := uint(0); i < n; i++ {
    b = s[i]
    switch b {
    case char.Ae:
      t += "\\\"a"
    case char.Oe:
      t += "\\\"o"
    case char.Ue:
      t += "\\\"u"
    case char.Sz:
      t += "\\ss "
    case char.Ä:
      t += "\\\"A"
    case char.Ö:
      t += "\\\"O"
    case char.Ü:
      t += "\\\"U"
    default:
      str.Append (&t, b)
    }
  }
  return t
}

func IsTeXer (a any) bool {
  if a == nil { return false }
  _, ok := a.(TeXer)
  return ok
}
