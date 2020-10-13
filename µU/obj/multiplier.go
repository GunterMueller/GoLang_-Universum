package obj

// (c) Christian Maurer   v. 201009 - license see µU.go

type
  Multiplier interface {

// Returns true, iff x is neutral w.r.t. multiplication.
  One() bool

// Pre: All y's are of a numerical type or implement Adder;
//      the result is in the range of the type of x.
// x is now the product of x before and all y's.
  Mul (y ...Any)

// x is now the square of x before.
  Sqr()

// x is now the n-th power of x before.
  Power (n uint)

// Pre: ! Zero(x).
// x is now the quotient of x before and y.
  DivBy (y Any)
}

// Pre: a is of a numerical type or implements Multiplier.
// Returns true, iff a is neutral w.r.t. multiplication.
func One (a Any) bool { return one(a) }

// Pre: a and all b's are of a numerical type or implement Multiplier.
// Returns the product of a and all b's.
func Mul (a Any, b ...Any) Any { return mul(a,b...) }

func Sqr (a Any) Any { return sqr(a) }
