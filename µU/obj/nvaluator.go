package obj

// (c) Christian Maurer   v. 220126 - license see µU.go

type
  NValuator interface {

// Pre: IsNValuator(x).
// Neturns the natural value of x.
  Val() uint

// Pre: IsNValuator(x).
// x.Val() == n.
  SetVal (n uint)
}

// Returns true, iff a is of type NValuatoer.
func IsNValuator (a Any) bool { return isNValuator(a) }

// Pre: IsNValuator(x).
// Returns the natural value of a.
func NVal (a Any) uint { return nVal(a) }
