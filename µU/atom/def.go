package atom

// (c) Christian Maurer   v. 220831 - license see µU.go
//
// >>> This package is only needed for the implementation of µU/mol;
//     it must not be used elsewhere.

import (
  . "µU/obj"
  "µU/col"
)
const (
  String = iota
  Natural
  Real
  Calendarday
  Clocktime
  Euro
  PhoneNumber
  Country
  Enum
  Ntypes
)
type
  Atom interface {

  Editor
  col.Colourer
  EditIndex()
  Print (l, c uint)

  Place (l, c uint)
  Pos() (uint, uint)
  Width() uint

  PosLess (Y any) bool

  String() string

  Index (b bool)
  IsIndex() bool

// Pre: If x has type Enum, x.EnumSet must have been called before.
// x is the atom interactively selected by the user.
  Select()

// Pre: t < NTypes
// x has the type t and width n.
  Define (t int, n uint)

// Returns the type of x.
  Typ() int

// If x has the type String, true is returned, iff x is a substring of Y.
// Returns otherwise true, iff x.Eq (Y).
  Sub (Y any) bool

  SelectColF()
  SelectColB()

  EnumName (n string)
  EnumSet (l, c uint)
  EnumGet()
}

// Returns a new atom of type Char.
func New() Atom { return new_() }
