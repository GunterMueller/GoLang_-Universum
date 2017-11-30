package obj

// (c) Christian Maurer   v. 170817 - license see nU.go

type
  Multiplier interface {

// Returns true, iff x is neutral w.r.t. multiplication.
  One () bool

// x is now the product of x before and all y's.
  Mul (y ...Multiplier)

// x is the product of y and z.
  Prod (y, z Multiplier)

// x = x0 * x0, where x0 denotes x before.
  Sqr ()

// x = y / z. // eventually deprecated
  Div (y, z Multiplier)

// x = x0 / y, where x0 denotes x before. // eventually deprecated
  DivBy (y Multiplier)
}
