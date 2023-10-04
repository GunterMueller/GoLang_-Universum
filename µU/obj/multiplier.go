package obj

// (c) Christian Maurer   v. 230924 - license see µU.go

type
  Multiplier interface {

// Returns true, iff the value of x is neutral w.r.t. multiplication.
  One() bool

// Pre: All y's are of the same type as x.
// x is now the product of x before and all y's.
  Mul (y ...Multiplier)

// Pre: y and z are of the same type as x.
// x is the product of y and z.
  Prod (y, z Multiplier)

// x is the square of x before.
  Sqr()

// x is now the n-th power of x before.
  Power (n uint)

// Returns true, if x is invertible.
  Invertible() bool

// Pre: x is invertible.
// x is inverted.
  Invert()

// Pre: y is of the same type as x; y is invertible; ! Zero(z).
// x is the quotient of x before and y.
  DivBy (y Multiplier)

// Pre: y and z are of the same type as x and not invertible; ! Zero(z).
// x is the quotient of y and z without remainder.
  Div (y, z Multiplier)

// Pre: y and z are of the same type as x and not invertible; ! Zero(z).
// x is the remainder of the division of y by z.
  Mod (y, z Multiplier)

// Pre: y and z are of the same type as x; y is invertible; ! Zero(z).
// x is the quotient of y and z without remainder.
  Quot (y, z Multiplier)
}
