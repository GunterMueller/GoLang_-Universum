package pers

// (c) Christian Maurer   v. 250407 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/day"
)
const ( // Format
  Name = iota // name, first name        1 line,  64 columns
  NameB       // Name + birth date       1 line,  80 columns
  NameBT      // NameB + title + remark  2 lines, 80 columns
  NFormats
)
const
  N = 5 // surname, first name, title, birth date and remarks
type
  Person interface {

  Editor
  Stringer
  col.Colourer
  Formatter
  TeXer
  Printer
  Rotator

// Returns the date of birth of x, if it is defined, otherwise the empty day.
  Birthdate() day.Calendarday

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
