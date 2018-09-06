package mul

// (c) Christian Maurer   v. 180902 - license see µU.go

import
  . "µU/obj"
type
  Multiplier interface {

// Returns true, iff x is neutral w.r.t. multiplication.
  One() bool

// x is now the product of x before and all y's.
  Mul (y ...Multiplier)

// x = x0 * x0, where x0 denotes x before.
  Sqr()

//  Power (n uint)

// x = y / z. // eventually deprecated
  Div (y, z Multiplier)

// x = x0 / y, where x0 denotes x before. // eventually deprecated
  DivBy (y Multiplier)
}

// Pre: a is of a numerical type or implements Multiplier.
// Returns true, iff a is neutral w.r.t. multiplication.
func One (a Any) bool { return one(a) }

// Pre: a and all b's are of a numerical type or implement Multiplier.
// Returns the product of a and all b's.
func Mul (a Any, b ...Any) Any { return mul(a,b...) }

func Sqr (a Any) Any { return sqr(a) }
