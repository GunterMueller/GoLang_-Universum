package obj

// (c) Christian Maurer   v. 180902 - license see µU.go

type
  Comparer interface {

// Pre: x is of the same type as the calling object.
// Returns true, iff the calling object is smaller than x.
  Less (x Any) bool
}

func IsComparer (a Any) bool { return isComparer(a) }

func Less (a, b Any) bool { return less(a,b) }

func Leq (a, b Any) bool { return leq(a,b) }
