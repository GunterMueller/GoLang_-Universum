package obj

// (c) murus.org  v. 140509 - license see murus.go

type
  Adder interface {

// Returns true, iff x is neutral w.r.t. addtion.
  Null() bool

// x is now the sum of x before and all y's. Returns x.
  Add (y ...Adder) Adder

// x is now the difference of x before and the sum of all y's. Returns x.
  Sub (y ...Adder) Adder
}

/* TODO
Six = Three.Sum (Four, Two)
Six = Sum (Four, Two)

Three.Mult (Four, Two) // Eight ?  or Twentyfour ?
Three.Mult (Four) // Twelve
Three.Times (Four, Two) // s.o.
Three.Times (Four) // Twelve // in Go: Mult
Three.Prod (Four, Two) Eight
Eight = Three.Prod (Four, Two)
Eight = Prod (Four, Two)
*/
