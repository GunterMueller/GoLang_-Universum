package appt

// (c) Christian Maurer   v. 211215 - license see µU.go

import (
  . "µU/obj"
  "µU/kbd"
  "todo/attr"
)
const ( // Format
  Long = Format(iota) // one complete screen line
  Short               // one line with 9 columns
)
type
  Appointment interface {

  Object
  Formatter
  Stringer
  Editor
  Printer

// Liefert genau dann true, wenn der aktuelle Suchbegriff
// im Suchwort von x enthalten ist.
  HasWord() bool

// Liefert das Terminattribut von x.
  Attrib() attr.Attribute
}

func PostEdit() (kbd.Comm, uint) { return CC, DD }
