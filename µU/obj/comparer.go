package obj

// (c) Christian Maurer   v. 221021 - license see µU.go

type
  Comparer interface {

// Pre: x is of the same type as the calling object.
// Returns true, iff the calling object is smaller than x.
  Less (x any) bool

// Pre: x is of the same type as the calling object.
// Returns true, iff the calling object is smaller than x
// or equals x.
  Leq (x any) bool
}

// Returns true, iff a implements Comparer.
func IsComparer (a any) bool { return isComparer(a) }

// Pre: a and b have the same type;
//      both are atomic or implement Comparer.
// Returns true, iff a is smaller than b.
func Less (a, b any) bool { return less(a,b) }

// Pre: a and b have the same type; both
//      both are atomic or implement Comparer and Equaler.
// Returns true, if a is smaller than b or a equals b.
func Leq (a, b any) bool { return leq(a,b) }
