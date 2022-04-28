package obj

// (c) Christian Maurer   v. 220420 - license see µU.go

type
  TeXer interface {

// Returns an AMSTeX-string representation of x.
  TeX() string
}

func IsTeXer (a any) bool {
  if a == nil { return false }
  _, ok := a.(TeXer)
  return ok
}
