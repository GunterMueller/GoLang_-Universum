package word

// (c) Christian Maurer   v. 210329 - license see µU.go

import
  . "µU/obj"
const
  Wd = 12
type
  Word interface {

  Object
  Editor
  Stringer
  Printer

// Liefert genau dann true, wenn der aktuelle Suchbegriff in x enthalten ist.
  Ok() bool
}

func New() Word { return new_() }
