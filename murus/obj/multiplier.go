package obj

// (c) murus.org  v. 140509 - license see murus.go

type
  Multiplier interface {

// Returns true, iff x is neutral w.r.t. all multiplication.
  One () bool

// x is now the product of x before and all y's. Returns x.
  Mul (y ...Multiplier) Multiplier

// x = x0 * x0, where x0 denotes x before.
  Sqr ()

// x = y / z. // eventually deprecated
  Div (y, z Multiplier)

// x = x0 / y, where x0 denotes x before. // eventually deprecated
  DivBy (y Multiplier)
}
