package mon

// (c) Christian Maurer   v. 220420 - license see µU.go

import
  . "µU/obj"
type
  Monitor interface {
// Specs: Buy my book and read the chapter on universal monitors.

  Wait (i uint)

  Signal (i uint)

  SignalAll (i uint)

  Blocked (i uint) uint

  Awaited (i uint) bool

  F (a any, i uint) any
}

// Returns a new Monitor with FuncSpectrum f
// with Signal-Urgent-Wait-semantics.
func New (n uint, f FuncSpectrum) Monitor { return new_(n,f) }
