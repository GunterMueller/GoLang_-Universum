package mon

// (c) Christian Maurer   v. 170411 - license see murus.go

//     Nichtsequentielle Programmierung mit Go 1 kompakt,
//     Kap. 6, insbesondere Abschnitt 6.15

import
  . "murus/obj"
type
  Monitor interface {
// Specs: Buy my book and read chapter 6.

  Wait (i uint)

  Awaited (i uint) bool

  Signal (i uint)

  SignalAll (i uint)

  F (a Any, i uint) Any
}

// Returns a new Monitor with Func- and Pred-Spectrum f and p resp.
func New (n uint, f FuncSpectrum, p PredSpectrum) Monitor { return new_(n,f,p) }
