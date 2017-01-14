package obj

// (c) murus.org  v. 150215 - license see murus.go

type
  Stringer interface {

// Returns a string representation of x.
  String() string

// Returns true, iff s represents an object.
// In this case, x is that object, otherwise x is unchanged.
  Defined (s string) bool // TODO return error instead of bool ?
}
