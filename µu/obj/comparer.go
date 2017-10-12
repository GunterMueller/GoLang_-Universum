package obj

// (c) Christian Maurer   v. 170701 - license see µu.go

type
  Comparer interface {

// Pre: x is of the same type as the calling object.
// Returns true, iff the calling object is smaller than x.
  Less (x Any) bool
}

func Less (a, b Any) bool { return less(a,b) }
// func Leq (a, b Any) bool { return leq(q,b) }
