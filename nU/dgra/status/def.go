package status

// (c) Christian Maurer  v. 190409 - license see nU.go

import
  . "nU/obj"
type
  Status interface { // pair of (phase, id)

  Equaler
  Comparer
  Coder

// x has phase p and id i.
  Set (p, i uint)

// Returns the phase of x.
  Phase() uint

// Returns the id of x.
  Id() uint

// phase of x is incremented.
  Inc()
}

// Returns a new status with phase 0 and id = ego.Me.
func New() Status { return new_() }
