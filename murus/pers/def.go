package pers

// (c) Christian Maurer   v. 170218 - license see murus.go

import
  . "murus/obj"
const ( // Format
           // ohne Maske:            1 Zeile,
  VeryShort = iota // Name, Vorname    maximal 35 Spalten
  Short    // Name, Vorname                    42 Spalten
  ShortB   // Name, Vorname (GebDat)           53 Spalten
  ShortT   // Anrede, Name, Vorname
  ShortTB  // Anrede, Name, Vorname (GebDat)
           // mit Maske:
  Long     // Name, Vorname, m/w     1 Zeile,  64 Spalten
  LongB    // lang, GebDat           1 Zeile,  80 Spalten
  LongT    // lang, Anrede           2 Zeilen, 64 Spalten
  LongTB   // lang, GebDat, Anrede   2 Zeilen, 80 Spalten
  NFormats
)
type
  Person interface {

  Editor
  Formatter
  Stringer
  Printer

// Returns true, if x and y coincide in surname, first name and birthday.
  Equiv (y Person) bool

// Returns true, iff surname, first name and birthday of x are not empty.
  Identifiable() bool

// Returns true, if x is at most 18 years old.
  FullAged() bool

  RotOrder()
}

func New() Person { return new_() }
