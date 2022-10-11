package buf

// (c) Christian Maurer   v. 220702 - license see nU.go

type Buffer interface { // Fifo-Queues

// Liefert genau dann true, wenn x keine Objekte enthält.
  Empty() bool

// Liefert die Anzahl der Objekte in x.
  Num() int

// a ist als letztes Objekt in x eingefügt.
  Ins (a any)

// Liefert das Musterobjekt von x, wenn x leer ist.
// Liefert andernfalls das erste Objekt von x
// und dieses Objekt ist aus x entfernt.
  Get() any
}

// Vor.: a ist atomar oder implementiert Equaler.
// Liefert eine leere Schlange für Objekte des Typs von a
// mit Musterobjekt a.
func New (a any) Buffer { return new_(a) }
func NewS (a any) Buffer { return newS(a) }
