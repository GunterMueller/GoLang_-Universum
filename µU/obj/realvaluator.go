package obj

// (c) Christian Maurer   v. 220420 - license see µU.go

type
  RealValuator interface {

// Pre: IsRealValuator(x).
// Returns the value of x.
  RealVal() float64

// Pre: IsRealValuator(x).
// x.RealVal() == r.
  SetRealVal (r float64)
}

// Returns true, iff a has a number type.
func IsRealValuator (a any) bool { return isRealValuator(a) }

/*/
// Pre: IsRealValuator(x).
// Returns the real value of a.
func RealVal (a any) float64 { return realVal(a) }
/*/
