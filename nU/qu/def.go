package qu

import . "nU/obj"

type Queue interface {

// Liefert genau dann true, wenn x leer ist,
// dh. keine Objekte enthält.
  Empty() bool

// Liefert die Anzahl der Objekte in x.
  Num() int

// a ist als letztes Objekt in x eingefügt.
  Enqueue (a Any)

// Liefert nil, wenn x leer ist;
// liefert andernfalls das erste Objekt aus x
// und dieses Objekt ist aus x entfernt.
  Dequeue() Any
}

// Liefert eine neue leere Schlange für Objekte vom Typ von a.
func New(a Any) Queue { return new_(a) }
func NewS(a Any) Queue { return news(a) }
