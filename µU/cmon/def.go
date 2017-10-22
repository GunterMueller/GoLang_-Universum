package cmon

// (c) Christian Maurer   v. 171019 - license see µU.go

import
  . "µU/obj"
type
  Monitor interface {
// Specs: Buy my book and read the section on conditioned universal monitors.

  Blocked (i uint) uint

  Awaited (i uint) bool

  F (i uint) uint
}

// Returns a new conditioned monitor with NFunc- and CondSpectrum f and c resp.
// with Signal-Urgent-Wait-semantics.
func New (n uint, f NFuncSpectrum, c CondSpectrum) Monitor { return new_(n,f,c) }
