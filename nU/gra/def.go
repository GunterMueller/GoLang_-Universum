package gra

// (c) Christian Maurer   v. 220702 - license see nU.go

import (. "nU/obj"; "nU/adj")

type Graph interface {

  Object

// Liefert genau dann true, wenn x gerichtet ist.
  Directed() bool

// Liefert die Anzahl der Ecken von x.
  Num() uint

// Liefert die Anzahl der Kanten von x.
  Num1() uint

// Wenn e nicht vom Type der Ecken von x ist oder schon
// in x enthalten ist, ist nichts verändert. Andernfalls
// ist e als Ecke in x eingefügt. War x vorher leer,
// ist e jetzt die kolokale und die lokale Ecke von x;
// sonst ist e jetzt die lokale Ecke von x und die
// vorherige lokale Ecke ist jetzt kolokale Ecke von x.
  Ins (e any)

// Wenn x leer war oder einen Kantentyp hat oder wenn
// die kolokale Ecke von x mit der lokalen Ecke von x
// übereinstimmt, ist nichts verändert.
// Andernfalls ist e als Kante von der kolokalen zur
// lokalen Ecke von x eingefügt (wenn diese beiden 
// Ecken vorher schon durch eine Kante verbunden
// waren, kann sich ihre Richtung verändert haben.)
  Edge (e any)

// Liefert genau dann true, wenn die kolokale Ecke nicht
// mit der lokalen Ecke von x übereinstimmt und es in x
// keine Kante von der kolokalen zur lokalen Ecke von x gibt.
  Edged() bool

// Liefert genau dann true, wenn v als Ecke in x enthalten
// ist. In diesem Fall ist v jetzt die lokale Ecke von x.
// Die kolokale Ecke von x ist die gleiche wie vorher.
  Ex (v any) bool

// Liefert genau dann true, wenn v und v1 als Ecken in x
// enthalten sind und nicht übereinstimmen. In diesem Fall
// ist v jetzt die kolokale Ecke von x und v1 die lokale.
  Ex2 (v, v1 any) bool

// Vor.: p ist auf Ecken definiert.
// Liefert true, wenn es eine Ecke in x gibt, für die
// p true liefert. In diesem Fall ist jetzt irgendeine
// solche Ecke die lokale Ecke von x.
// Die kolokale Ecke von x ist die gleiche wie vorher.
  ExPred (p Pred) bool

// Liefert nil, wenn x leer ist.
// Liefert andernfalls eine Kopie der lokalen Ecke von x.
  Get() any

// Liefert eine Kopie der Musterkante von x, wenn x
// leer ist oder es keine Kante von der kolokalen
// zur lokalen Ecke von x gibt oder diese beiden Ecken
// übereinstimmen. Liefert sonst eine Kopie der Kante
// von der kolokalen Ecke von x zur lokalen Ecke von x.
  Get1() any

// Liefert (nil, nil), wenn x leer ist.
// Liefert andernfalls ein Paar, bestehend aus einer
// Kopie der kolokalen und einer der lokalen Ecke von x.
  Get2() (any, any)

// Wenn x leer oder v nicht vom Eckentyp von x oder
// v nicht in x enthalten ist, ist nicht verändert.
// Andernfalls ist v jetzt die lokale Ecke von x
// und ist markiert.
// Die kolokale Ecke von x ist die gleiche wie vorher.
  Mark (v any)

// Wenn x leer ist oder wenn v oder v1 nicht vom
// Eckentyp von x sind oder wenn v oder v1 nicht in x
// enthalten sind oder wenn v und v1 zusammenfallen,
// ist nichts verändert.
// Andernfalls ist v jetzt die kolokale und v1
// die lokale Ecke von x und diese beiden Ecken
// und die Kante zwischen ihnen sind jetzt markiert.
  Mark2 (v, v1 any)

// Liefert genau dann true, wenn alle Ecken und Kanten
// von x markiert sind.
  AllMarked() bool

// Wenn x leer ist, ist nichts verändert.
// Andernfalls stimmt jetzt die kolokale Ecke von x
// mit der lokalen Ecke von x überein, wobei das
// für f == true die vorherige lokale Ecke ist und
// für f == false die vorherige kolokale Ecke von x.
// Das einzige markierte Element in x ist diese Ecke.
  Locate (f bool)

// Liefert 0, wenn x leer ist; liefert andernfalls
// die Anzahl der in die lokalen Ecke von x eingehenden Kanten.
  NumNeighboursIn() uint

// Liefert 0, wenn x leer ist; liefert andernfalls
// die Anzahl der von der lokalen Ecke von x ausgehenden Kanten.
  NumNeighboursOut() uint

// Liefert 0, wenn x leer ist. Liefert andernfalls die Anzahl
// aller ein- und ausgehenden Kanten der lokalen Ecke von x.
  NumNeighbours() uint

// Liefert false, wenn x leer oder i >= NumNeighbours() ist;
// liefert andernfalls genau dann true, wenn die Kante zum
// i-ten Nachbarn der lokalen Ecke eine ausgehende Kante ist.
  Outgoing (i uint) bool

// Liefert false, wenn x leer oder i >= NumNeighbours() ist;
// liefert andernfalls genau dann true, wenn die Kante zum
// i-ten Nachbarn der lokalen Ecke eine eingehende Kante ist.
  Incoming (i uint) bool

// Liefert nil, wenn x leer ist oder i >= NumNeighbours()
// gilt; liefert andernfalls eine Kopie des i-ten Nachbarn
// der lokalen Ecke von x.
  Neighbour (i uint) any

// Vor.: o ist auf Ecken definiert.
// o ist auf alle Ecken von x angewendet. Die kolokale und
// die lokale Ecke von x ist jeweils die gleiche wie vorher.
  Trav (o Op)

// Liefert einen leeren Graphen, wenn x leer ist.
// Liefert andernfalls einen Graphen mit der lokalen Ecke
// von x als einziger Ecke allen von ihr aus- und bei ihr
// eingehenden Kanten. Markiert sind nur diese Ecke
// und alle Kanten zwischen ihr und ihren Nachbarn.
  Star() Graph

// Vor.: x ist genau dann gerichtet,
//       wenn es alle Graphen y sind.
// x besteht aus allen Ecken und Kanten
// von x vorher und aus allen Graphen y.
// Dabei sind alle Markierungen von y übernommen.
  Add (y ...Graph)

// Liefert die Repräsentation von x als Adjazenzmatrix.
  Matrix() adj.AdjacencyMatrix

// Vor.: a ist genau dann symmetrisch, wenn x geordnet ist.
// x ist der Graph mit den Ecken a.Vertex(i) und Kanten
// von a.Vertex(i) nach a.Vertex(k), falls a.Val(i,k) > 0 ist
// (i, k < a.Num()).
  SetMatrix (a adj.AdjacencyMatrix)

// w und w2 sind die Operationen zur Ausgabe von x.
  SetWrite (w CondOp, w2 CondOp2)

// Die Werte der Ecken von x sind an ihren Positionen
// und die Kanten von x als einfache Linien ausgegeben.
// durch ":" getrennt dahinter alle ihre Nachbarecken.
  Write()
}

// Vor.: v is atomar oder implementiert Object.
//       e == nil oder e ist von einem uint-Typ
//       oder implementiert Valuator.
// Liefert einen leeren Graphen. der genau dann gerichtet ist,
// wenn d den Wert true hat.
// v ist Musterecke von x, die den Typ der Ecken definiert.
// Für e == nil hat x keinen Kantentyp und alle Kanten
// haben den Wert 1; andernfalls ist e die Musterkante von x,
// die den Typ der Kanten von x definiert.
func New (d bool, v, e any) Graph { return new_(d,v,e) }
