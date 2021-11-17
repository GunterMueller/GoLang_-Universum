package obj

// (c) Christian Maurer   v. 221206 - license see µU.go

type
  Multiplier interface {

// Returns true, iff x is neutral w.r.t. multiplication.
  One() bool

// Pre: All y's implement Multiplier and are of the same type as x.
// x is now the product of x before and all y's.
  Mul (y ...Multiplier)

// Pre: y and z are of the same type as x.
// x is the product of y and z.
  Prod (y, z Multiplier)

// x is now the square of x before.
  Sqr()

// x is now the n-th power of x before.
  Power (n uint)

// Returns true, if x is invertible.
  Invertible() bool

// Pre: x is invertible.
// x is inverted.
  Invert()

// Pre: ! Zero(y).
// x is now the quotient of x before and y.
  DivBy (y Multiplier)

// Pre: y and z are of the same type as x.
// x is the quotient of y and z.
  Quot (y, z Multiplier)
}
