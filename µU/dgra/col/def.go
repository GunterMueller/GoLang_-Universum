package col

// (c) Christian Maurer  v. 180829 - license see µU.go

import
  . "µU/obj"
type
  Colour interface { // pair of (level, id)

  Equaler
  Comparer
  Coder

// x has level l and id i.
  Set (l uint, i uint)

// Returns the level of x.
  Level() uint

// Returns the id of x.
  Id() uint

// The level of x is incremented.
  Inc()
}

// Returns a new status with level 0 and id = ego.Me.
func New() Colour { return new_() }
