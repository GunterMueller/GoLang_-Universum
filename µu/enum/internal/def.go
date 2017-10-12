package internal

// (c) Christian Maurer   v. 170919 - license see µu.go

import (
  . "µu/obj"
  "µu/col"
)
const (
  Short = Format(iota)
  Long
  NFormats
)
type
  Base interface {
// New (b) returns a new Object of type t,
// where b == byte(e) for e == one of enum/Enum.

  Formatter
  Object
  col.Colourer
  Editor
  Stringer
  Printer

// Returns the type of x.
  Typ () uint8

// Returns the number of elements of Enum (common for all elements).
  Num () uint8

// Returns the order number of x.
  Ord () uint8

// Returns the width of the string representation of x (common for all elements).
  Wd () uint

// Returns true, iff there is an n-th element in Enum.
// In this case x is that element, otherwise x is empty.
  Set (n uint8) bool
}

func New (t uint8, s [NFormats][]string) Base { return new_(t,s) }
