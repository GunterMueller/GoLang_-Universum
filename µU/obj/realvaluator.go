package obj

// (c) Christian Maurer   v. 210309 - license see µU.go

type
  RealValuator interface {

// Pre: IsRealValuator(x).
// Returns the real value of x.
  RealVal() float64

// Pre: IsRealValuator(x).
// x.Val() == r.
  RealSet (r float64)
}

// Returns true, iff a has a number type.
func IsRealValuator (a Any) bool { return isRealValuator(a) }

// Pre: IsRealValuator(x).
// Returns the real value of a.
func RealVal (a Any) float64 { return realVal(a) }
