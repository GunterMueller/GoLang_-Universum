package clk

// (c) murus.org  v. 161216 - license see murus.go

import
  . "murus/obj"
const ( // Format
  Hh_mm = iota // e.g. "07.32"
  Hh_mm_ss     // e.g. "13.45:27"
  Mm_ss        // e.g. "04:19"
  NFormats
  NSecondsPerDay = sd
)

func New() Clocktime { return newClk() }

type
  Clocktime interface { // given by a triple of uints h.m:s with h < 24 and m, s < 60.

  Editor
  Stringer
  Formatter
  Printer

// x is equal to the system time.
  Actualize()

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
