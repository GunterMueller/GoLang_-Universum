package date

// (c) Christian Maurer   v. 161216 - license see murus.go

import (
  . "murus/obj"
  "murus/day"
  "murus/clk"
)
type
  DayTime interface {
// Pairs of Calendarday and Clocktime.
// (M/O, 2) means (last Sunday in March/October, 2.00:00)

  Editor
  Stringer
  Printer

// x is (d, t).
  Set (d day.Calendarday, t clk.Clocktime)

// Returns the day/the time of x.
  Day() day.Calendarday
  Time() clk.Clocktime

// Returns true, iff x is not empty and x < (O, 2) or (M, 2) <= x.
  Normal () bool

// x has the format d for its day and t for its clktime.
  SetFormat (d, t Format)

// If x is Normal, then y equals x, otherwise y equals x + 1 hour.
  Actualize (y DayTime)

// spec TODO
  Normalize ()

// If x is not empty, x is increased by y.
  Inc (y clk.Clocktime)

// If x is not empty, x is decreased by y.
  Dec (y clk.Clocktime)
}

// Returns a new empty pair of day and time
func New() DayTime { return new_() }
