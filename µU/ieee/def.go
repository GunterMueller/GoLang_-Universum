package ieee

// (c) Christian Maurer   v. 201204 - license see µU.go

// >>> IEEE 754 binary representation of double precision real numbers

type
  IEEE interface { // pair (float64, string)

// f is the value of x.
  SetFloat64 (f float64)

// Pre: The value of x was set by a call of SetFloat64.
// Returns the IEEE-representation as string.
  String() string

// Pre: len(s) == 64.
// s is the string of x.
  SetString (s string)

// Pre: The string of x was set by a call of SetString.
// Returns x has the value f.
  Float64() float64
}

// Returns a new IEEE with value 0. and empty string.
func New() IEEE { return new_() }

