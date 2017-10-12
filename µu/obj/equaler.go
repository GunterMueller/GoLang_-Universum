package obj

// (c) Christian Maurer   v. 170701 - license see µu.go

type
  Equaler interface {

// Returns true, iff the x has the same type as y
// and coincides with it in all its value[s].
  Eq (y Any) bool

// If y has the same type as x, then x.Eq(y) (y is unchanged).
  Copy (y Any)

// Returns a clone of x, i.e. x.Eq(x.Clone()).
  Clone() Any
}

// Pre: a and b are atomic or implement Equaler.
// Returns true, if a and b are equal.
func Eq (a, b Any) bool { return eq(a,b) }

// Pre: a is atomic or implements Equaler.
// Returns a clone of a.
func Clone (a Any) Any { return clone(a) }
