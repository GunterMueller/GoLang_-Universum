package obj

// (c) Christian Maurer   v. 220126 - license see µU.go

type
  ZValuator interface {

// Pre: IsZValuator(x).
// Returns the integer value of x.
  Val() int

// Pre: IsZValuator(x).
// x.Val() == z.
  SetVal (z int)
}

// Returns true, iff a is of type ZValuatoer.
func IsZValuator (a Any) bool { return isZValuator(a) }

// Pre: IsZValuator(x).
// Returns the natural value of a.
func ZVal (a Any) int { return zVal(a) }
