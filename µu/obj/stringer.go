package obj

// (c) Christian Maurer   v. 150418 - license see µu.go

type
  Stringer interface {

// Returns a string representation of x.
  String() string

// Returns true, iff s represents an object.
// In this case, x is that object, otherwise x is undefined.
  Defined (s string) bool
}
