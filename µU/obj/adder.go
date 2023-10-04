package obj

// (c) Christian Maurer   v. 230924 - license see µU.go

type
  Adder interface {

// Returns true, iff x is neutral w.r.t. addition.
  Zero() bool

// Pre: All y's are of the same type as x.
// x is now the sum of x before and all y's.
  Add (y ...Adder)

// Pre: y and z are of the same type as x.
// x is the sum of y and z.
  Sum (y, z Adder)

// Pre: All y's are of the same type as x.
// x is now the difference of x before and the sum of all y's.
  Sub (y ...Adder)

// Pre: y and z are of the same type as x.
// x is the differenz of y and z.
  Diff (y, z Adder)
}
