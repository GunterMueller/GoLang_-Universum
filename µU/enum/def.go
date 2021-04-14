package enum

// (c) Christian Maurer   v. 201007 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
const ( // Format
  Short = Format(iota)
  Long
  NFormats
)
type
  Type = uint8; const ( // details s. corresponding files
  Title = Type(iota)
  AudioC
  BookC
  Writer
  Composer
  AudioMedium
  Religion
  Subject
  Wortart
  Casus
  Genus
  Persona
  Numerus
  Tempus
  Modus
  GenusVerbi
  Comparatio
  NTypes
)
type
  Enumerator interface { // A set of at most 256 enumerated elements,
                         // represented by strings of len <= 64.
                         // The 0-th element is "empty", represented by spaces.
  Object
  Formatter
  col.Colourer
  Editor
  Stringer
  Printer

// Returns the type of x.
  Typ() uint8 // Type

// Returns the number of elements in the type of x.
  Num() uint8

// Returns the order number of x.
  Ord() uint8

// Returns the width of the string representation of x (common for all elements).
  Wd() uint

// Returns true, iff the type of x has an n-th element.
// In this case x is that element, otherwise x is empty.
  Set (n uint8) bool

// If s equals an entry in sort,
// then the number of that entry is returned,
// otherwise 0.
  Found (s string, sort uint8, f Format) uint8
}
var
  N []Type

// Pre: e < NEnums.
// Returns a new empty enumerator for objects of Type t.
func New (t Type) Enumerator { return new_(t) }
