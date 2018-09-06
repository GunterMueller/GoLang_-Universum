package obj

// (c) Christian Maurer   v. 180902 - license see µU.go

type
  Valuator interface {

// Returns the value of x.
  Val() uint

// Returns true, iff x can be given a value. In that case x.Val() == n.
// Returns otherwise false.
  SetVal (n uint) bool
}

// Returns true, iff a implements Valuator.
func IsValuator (a Any) bool { return isValuator(a) }

// TODO Spec
func Val (a Any) uint { return val(a) }

// TODO Spec
func SetVal (x *Any, n uint) { setVal(x,n) }
