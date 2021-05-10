package obj

// (c) Christian Maurer   v. 210509 - license see µU.go

type
  TeXer interface {

// Returns a TeX-string representation of x.
  TeX() string
}

func IsTeXer (a Any) bool {
  if a == nil { return false }
  _, ok := a.(TeXer)
  return ok
}
