package clk

// (c) Christian Maurer   v. 171011 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
const ( // Format
  Hh_mm = iota // e.g. "12.34"
  Hh_mm_ss     // e.g. "12.34:56"
  Mm_ss        // e.g. "34:56"
  NFormats
  NSecondsPerDay = sd
)
type
  Clocktime interface { // given by a triple of uints h.m:s with h < 24 and m, s < 60.

  Object
  col.Colourer
  Editor
  Stringer
  Formatter
  Printer

// x is equal to the system time.
  Update()

// Returns true, iff x lies before the system time.
  Elapsed() bool

// Returns the absolute value of the seconds between y and x.
  Distance(y Clocktime) uint

// Returns the Distance between 00.00:00 and x.
  NSeconds() uint

// Returns the value of the hour, the minutes, the seconds of x.
  Hours() uint
  Minutes() uint
  Seconds() uint

// x is increased by y mod 24 hours.
  Inc(y Clocktime)

// x is decreased by y mod 24 hours.
  Dec(y Clocktime)

// Returns true, iff h.m:s defines a clocktime, and x is that clocktime.
// Otherwise, x is empty.
  Set(h, m, s uint) bool

// Returns true, iff s < NSecondsPerDay 24, and x is set to the clocktime y
// such that y.NSeconds() == s. Otherwise x is empty.
  SetSeconds(s uint) bool
}

// Returns a new empty (i.e. with undefined data) clocktime.
func New() Clocktime { return new_() }
