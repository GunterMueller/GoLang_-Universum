package obj

// (c) Christian Maurer   v. 170701 - license see µU.go

type
  Valuator interface {

// Returns the value of x.
  Val() uint

// Returns true, iff x can be given a value. In that case x.Val() == n.
// Returns otherwise false.
  SetVal (n uint) bool
}

func Val (a Any) uint { return val(a) }
func SetVal (x *Any, n uint) { setVal(x,n) }
