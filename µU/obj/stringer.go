package obj

// (c) Christian Maurer   v. 180902 - license see µU.go

type
  Stringer interface {

// Returns a string representation of x.
  String() string

// Returns true, iff s represents an object.
// In this case, x is that object, otherwise x is undefined.
  Defined (s string) bool
}

func IsStringer (a Any) bool {
  if a == nil { return false }
  _, ok := a.(Stringer)
  return ok
}
