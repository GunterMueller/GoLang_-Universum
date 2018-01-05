package bbuf

// (c) Christian Maurer   v. 171125 - license see nU.go

import (. "nU/obj"; "nU/buf")

type BoundedBuffer interface {

  buf.Buffer

// Liefert genau dann true, wenn x bis zu seiner Kapazität gefüllt ist.
// ! Full() ist Vor. für einen Aufruf von Ins(a).
  Full() bool
}

// Vor.: a ist atomar oder implementiert Equaler.
// Liefert einen leeren Puffer der Kapazität n für Objekte vom Typ a.
func New (a Any, n uint) BoundedBuffer { return new_(a,n) }
func New1 (a Any, n uint) BoundedBuffer { return new1(a,n) }
