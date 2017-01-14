package enum

// (c) murus.org  v. 161216 - license see murus.go

import
  . "murus/obj"
const ( // Format
  Short = Format(iota)
  Long
  NFormats
)
const ( // details s. corresponding files
  Title = uint8(iota)
  Writer
  Composer
  RecordLabel
  AudioMedium
  SparsCode
  Religion
  Subject
  LexicalCategory
  Casus
  Genus
  Persona
  Numerus
  Tempus
  Modus
  GenusVerbi
  Comparatio
  NEnums
)
type
  Enumerator interface { // A set of at most 256 enumerated elements,
                         // represented by strings of len <= 64.
                         // The 0-th element is "empty", represented by spaces.
  Formatter
  Editor
  Stringer
  Printer

// Returns the type of x.
  Typ() uint8

// Returns the number of elements in the type of x.
  Num() uint8

// Returns the order number of x.
  Ord() uint8

// Returns the width of the string representation of x (common for all elements).
  Wd() uint

// Returns true, iff the type of x has an n-th element.
// In this case x is that element, otherwise x is empty.
  Set (n uint8) bool
}
// Pre: e < NEnums.
// Returns a new empty enumerator for objects of Title e.
func New (e uint8) Enumerator { return newEnum(e) }
