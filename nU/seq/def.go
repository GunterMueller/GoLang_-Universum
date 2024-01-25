package seq

// (c) Christian Maurer   v. 220702 - license see nU.go

type
  Sequence interface {

// Liefert genau dann true, wenn x keine Objekte enthält.
  Empty() bool

// Liefert die Anzahl der Objekte in x.
  Num() int

// a ist als letztes Objekt in x eingefügt.
  InsLast (a any)

// Wenn x leer ist, ist nichts verändert.
// Andernfalls ist das erste Objekt aus x entfernt.
  DelFirst()

// Liefert nil, falls x leer ist,
// andernfalls das erste Objekt aus x.
  GetFirst() any
}

// Vor.: a ist atomar oder implementiert Equaler.
// Liefert eine leere Folge mit Musterobjekt a.
func New (a any) Sequence { return new_(a) }
