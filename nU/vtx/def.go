package vtx

// (c) Christian Maurer   v. 171229 - license see nU.go

import . "nU/obj"

type Vertex interface {

  Object // leer, wenn Wert = 0.
  Valuator

// x has the the center (x, y).
  Set (x, y uint)

// Returns the coordinates of x.
  Pos() (uint, uint)

// x ist an seiner Position in Weiß ausgegeben.
  Write()

// x ist an seiner Position ausgegeben,
// für a == true in Rot, sonst in Weiß.
  Write1 (a bool)
}

// Liefert eine neue leere Ecke.
func New (n uint) Vertex { return new_(n) }

// Vor.: v implementiert Vertex.
// v ist an seiner Position ausgegeben,
// für a == true in Rot, sonst in Weiß.
func W (v Any, a bool) { w(v,a) }

// Vor.: v und v0 implementieren Vertex.
// Die Positionen von v und v0 sind durch eine Linie
// verbunden, für a == true in Rot, sonst in Weiß.
func W2 (v, v1 Any, a bool) { w2(v,v1,a) }
