package obj

// (c) Christian Maurer   v. 201106 - license see µU.go

type
  Adder interface {

// Returns true, iff x is neutral w.r.t. addition.
  Zero() bool

// Pre: All y's are of a numerical type or implement Adder;
//      the result is in the range of the type of x.
// x is now the sum of x before and all y's.
  Add (y ...Adder)

// Pre: All y's are of a numerical type or implement Adder;
//      the result is in the range of the type of x.
// x is now the difference of x before and the sum of all y's.
  Sub (y ...Adder)
}

// Pre: a is of a numerical type or implements Adder.
// Returns true, if a is neutral w.r.t. addition.
func Zero (a Any) bool { return zero(a) }

// Pre: a and all b's are of a numerical type or implement Adder;
//      a and all b's are of the same type.
// Returns the sum of a and all b's.
func Add (a Any, b ...Any) Any { return add(a, b...) }

/*/
// Pre: All a's are of a numerical type or implement Adder;
//      all a's are of the same type. len(as) > 0.
// Returns the sum of all a's.
func Sum (as []Adder) Adder { return sum(as) }
/*/
