package status

// (c) Christian Maurer  v. 180824 - license see µU.go

import
  . "µU/obj"
type
  Status interface { // pair of (phase, id)

  Equaler
  Comparer
  Coder

// x has phase p and id i.
  Set (p int, i uint)

// Returns the phase of x.
  Phase() int

// Returns the id of x.
  Id() uint

// The phase of x is incremented.
  Inc()
}

// Returns a new status with phase -1 and id = ego.Me.
func New() Status { return new_() }
