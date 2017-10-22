package obj

// (c) Christian Maurer   v. 170701 - license see µU.go

type
  RealValuator interface {

// Returns the real value of x.
  RealVal() float64

//// Returns true, iff x is defined with x.Val() == r.
  RealSet (r float64) bool
}

func RealVal (a Any) float64 { return realVal(a) }
