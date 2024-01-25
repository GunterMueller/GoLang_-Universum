package qu

// (c) Christian Maurer   v. 220801 - license see nU.go

type
   Queue interface {

// Liefert genau dann true, wenn x leer ist,
// dh. keine Objekte enthält.
  Empty() bool

// Liefert die Anzahl der Objekte in x.
  Num() int

// a ist als letztes Objekt in x eingefügt.
  Enqueue (a any)

// Liefert nil, wenn x leer ist;
// liefert andernfalls das erste Objekt aus x
// und dieses Objekt ist aus x entfernt.
  Dequeue() any
}

// Liefert eine neue leere Schlange für Objekte vom Typ von a.
func New(a any) Queue { return new_(a) }
func NewS(a any) Queue { return news(a) }
