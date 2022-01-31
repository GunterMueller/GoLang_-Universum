package obj

// (c) Christian Maurer   v. 220126 - license see µU.go

type
  RValuator interface {

// Pre: IsRealValuator(x).
// Returns the real value of x.
  Val() float64

// Pre: IsRealValuator(x).
// x.Val() == r.
  SetVal (r float64)
}

// Returns true, iff a is of type RValuator.
func IsRValuator (a Any) bool { return isRValuator(a) }

// Pre: IsRealValuator(x).
// Returns the real value of a.
func RVal (a Any) float64 { return rVal(a) }
