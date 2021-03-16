package obj

// (c) Christian Maurer   v. 210309 - license see µU.go

type
  Equaler interface {

// Returns true, iff the x has the same type as y
// and coincides with it in all its value[s].
  Eq (y Any) bool

// Pre: y has the same type as x.
// x.Eq (y) (y is unchanged).
  Copy (y Any)

// Returns a clone of x, i.e. x.Eq (x.Clone()).
  Clone() Any
}

// Returns true, iff a implements Equaler.
func IsEqualer (a Any) bool { return isEqualer(a) }

// Pre: a and b have the same type;
//      both are atomic or implement Equaler.
// Returns true, if a and b are equal.
func Eq (a, b Any) bool { return eq(a,b) }

// Pre: a is atomic or implements Equaler.
// Returns a clone of a.
func Clone (a Any) Any { return clone(a) }

// Returns true, iff a is atomic or implements Equaler.
func AtomicOrEqualer (a Any) bool { return Atomic(a) || isEqualer(a) }
