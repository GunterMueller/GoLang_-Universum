package pers

// (c) Christian Maurer   v. 240408 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
const ( // Format
  Name = iota // name, first name    1 line,  64 columns
  NameB       // Namg + birth date   1 line,  80 columns
  NameBT      // NameB + title       2 lines, 80 columns
  NFormats
)
const
  N = 4 // surname, first name, title and birth date
type
  Person interface {

  Editor
  Stringer
  col.Colourer
  Formatter
  TeXer
  Printer
  Rotator

// Returns true, if x and y coincide in surname, first name and birthday.
  Equiv (y Person) bool

  Sub (y any) bool

// Returns true, iff surname, first name and birthday of x are not empty.
  Identifiable() bool

// Returns true, if x is at most 18 years old.
  FullAged() bool

// Returns the age of x (in years) at the time of the call.
  Age() uint

  Rotate()
}

func New() Person { return new_() }
