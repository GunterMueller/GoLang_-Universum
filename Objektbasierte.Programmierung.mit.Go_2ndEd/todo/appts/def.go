package appts

// (c) Christian Maurer   v. 210329 - license see µU.go

import (
  . "µU/obj"
  "µU/day"
)
type
  Appointments interface {

  Object
  Editor
  Printer

  SetFormat (p day.Period)

// Liefert genau dann true, wenn der aktuelle Suchbegriff
// im Suchwort von eine der Termine in x enthalten ist.
  HasWord () bool
}
