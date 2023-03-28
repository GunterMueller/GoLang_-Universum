package page

// (c) Christian Maurer   v. 210322 - license see µU.go

import (
  . "µU/obj"
  "µU/day"
)
type
  Page interface {

  Object
  Editor
  Printer
  Indexer

// x ist 
  SetFormat (p day.Period)

// d ist das Datum von x.
  Set (d day.Calendarday)

// Liefert das Datum von x.
  Day() day.Calendarday

// Liefert genau dann true, wenn der aktuelle Suchbegriff
// in dem Stichwort eines Termins von x enthalten ist.
  HasWord() bool

// s. dayattr.
  Fin()
}
