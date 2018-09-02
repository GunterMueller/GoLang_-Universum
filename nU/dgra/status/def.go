package status

// (c) Christian Maurer  v. 180819 - license see µU.go

import
  . "nU/obj"
type
  Status interface { // pair of (phase, id)

  Equaler
  Comparer
  Coder

// Returns the phase of x.
  Phase() uint
// Returns the id of x.
  Id() uint
// phase of x is incremented.
  Inc()
}

func New (p, i uint) Status { return new_(p,i) }
