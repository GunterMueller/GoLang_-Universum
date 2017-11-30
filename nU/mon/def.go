package mon

// (c) Christian Maurer   v. 171125 - license see nU.go

import . "nU/obj"

type Monitor interface {

  Wait (i uint)

  Blocked (i uint) uint

  Awaited (i uint) bool

  Signal (i uint)

  SignalAll (i uint)

  F (a Any, i uint) Any
}

// Returns a new Monitor with Funcspectrum f.
func New (n uint, f FuncSpectrum) Monitor { return new_(n,f) }
