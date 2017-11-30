package cmon

// (c) Christian Maurer   v. 171125 - license see nU.go

import . "nU/obj"

type Monitor interface {

  Blocked (i uint) uint

  Awaited (i uint) bool

  F (i uint) uint
}

// Returns a new conditioned monitor with NFunc- and CondSpectrum f and c resp.
func New (n uint, f NFuncSpectrum, c CondSpectrum) Monitor { return new_(n,f,c) }
