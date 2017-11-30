package obj

// (c) Christian Maurer   v. 170817 - license see nU.go

type
  Adder interface {

// Returns true, iff x is neutral w.r.t. addition.
  Null() bool

// x is now the sum of x before and all y's.
  Add (y ...Adder)

// x is the sum of y and z.
  Sum (y, z Adder)

// x is now the difference of x before and the sum of all y's.
  Sub (y ...Adder)

// x is the difference of y and z.
  Diff (y, z Adder)
}
