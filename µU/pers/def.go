package pers

// (c) Christian Maurer   v. 210511 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
const ( // Format
// ohne Maske
  Short = iota // Name, Vorname         1 Zeile,  43 Spalten
  ShortB       // Short + GebDat        1 Zeile,  53 Spalten
// mit Maske
  Long         // Name, Vorname         1 Zeile,  64 Spalten
  LongB        // Long + GebDat         1 Zeile,  80 Spalten
  LongTB       // Anrede + LongB + m/w  2 Zeilen, 80 Spalten
  NFormats
)
type
  Person interface {

  Object
  col.Colourer
  Editor
  Formatter
  TeXer
  Printer
  Rotator

// Returns true, if x and y coincide in surname, first name and birthday.
  Equiv (y Person) bool

// Returns true, iff surname, first name and birthday of x are not empty.
  Identifiable() bool

// Returns true, if x is at most 18 years old.
  FullAged() bool

// Returns the age of x (in years) at the time of the call.
  Age() uint

  Rotate()
}

func New() Person { return new_() }
