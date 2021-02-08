package adj

// (c) Christian Maurer   v. 210123 - license see nU.go

import . "nU/obj"

type AdjacencyMatrix interface {
// Quadratische Matrizen mit Paaren (v, e) als Einträgen,
// wobei v atomar ist oder Object implementiert und
// e einen uint-Typ hat oder Valuator implementiert.
// Der Eintrag einer Matrix x in Reihe i und Spalte k ist x(i,k).
// Jede solche Matrix definiert auf folgende Weise einen Graohen:
// x(i,k) = (e, v) bedeutet
// für i == k:
//     v ist eine Ecke in dem Graphen. In diesem Fall ist e die Musterkante von x.
// für i != k:
//     Es gibt eine Kante mit dem Wert e, die von der i-ten Ecke des Graphen ausgeht und
//     bei seiner k-ten Ecke ankommt, wenn v nicht mit der Musterecke von x übereinstimmt.
//     In diesem Fall ist v die Musterecke von x.
// Die Muster sind diejenigen Objekte, die dem Konstruktor als Parameter übergeben werden;
// sie dürfen nicht als Ecken oder Kanten benutzt werden.

  Object

// Liefert die Anzahl der Zeilen/Spalten von x.
  Num() uint

// Liefert genau dann true, wenn x und y die gleiche Zeilenzahl,
// gleiche Musterecken und gleiche Musterkanten haben.
  Equiv (y AdjacencyMatrix) bool

// Vor.: e hat den Typ der Musterecke von x.
// Wenn i oder k >= x.Num(), ist nichts verändert.
// Andernfalls gilt:
// x(i,k) ist das Paar (v, e) mit v = Musterecke von x,
// d.h., in dem ensprechenden Graph gibt es genau dann
// eine Kante mit dem Wert von e von seiner i-ten Ecke
// zu seiner k-ten Ecke, wenn x.Val (i,k) > 0 ist.
  Edge (i, k uint, e Any)

// Liefert das erste Element des Paares x(i,i), also eine Ecke.
  Vertex (i uint) Any

// Vor.: i, k < x.Num().
// Liefert 0, wenn in x(i,k) = (v, e) e die Musterkante
// von x ist; liefert andernfalls den Wert von e.
  Val (i, k uint) uint

// Vor.: v hat den Typ der Musterecke von x und
//       e hat den Typ der Musterkante von x.
// Wenn i oder k >= x.Num(), ist nichts verändert.
// Andernfalls ist jetzt x(i,k) == (v, e).
  Set (i, k uint, v, e Any)

// Liefert genau dann true, wenn x(i,k) == x(k,i) für alle i, k < x.Num(),
// d.h. der entsprechende Graph ist ungerichtet.
  Symmetric() bool

// Vor.: x und y sind äquivalent.
// x enthält alle Einträge von x und dazu alle Einträge
// von y, bei Kanten aber nur diejenigen mit einem Wert > 0.
// Einträge von x, die an den gleichen Stellen in y vorkommen,
// sind dabei von den Einträgen in x überschrieben.
  Add (y AdjacencyMatrix)

// Liefert genau dann true, wenn jede Zeile von x
// mindestens einen Eintrag (v, e) mit x.Val(e) > 0 enthält,
// d.h., wenn jede Ecke im entsprechenden Graphen
// mindestens eine ausgehende Kante hat.
  Full () bool

// Vor.: Die Einträge von x sind vom Typ uint
//       oder implementieren Valuator.
// x ist auf dem Bildschirm ausgegeben.
  Write()
}

// Vor.: n > 0. v ist atomar oder implementiert Object und
//       e hat einen uint-Typ oder implementiert Valuator.
// v ist die Musterecke und e die Musterkante von x.
// Liefert eine n*n-Matrix nur mit Einträgen (v, e).
func New (n uint, v, e Any) AdjacencyMatrix { return new_(n,v,e) }
