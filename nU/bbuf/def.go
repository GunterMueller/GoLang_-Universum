package bbuf

// (c) Christian Maurer   v. 171125 - license see nU.go

import (. "nU/obj"; "nU/buf")

type BoundedBuffer interface {

  buf.Buffer

// Liefert genau dann true, wenn x bis zu seiner Kapazität gefüllt ist.
// ! Full() ist Vor. für einen Aufruf von Ins(a).
  Full() bool
}

// Vor.: a ist atomar oder implemtiert Object.
// Liefert einen leeren Puffer der Kapazität n für Objekte vom Typ a.
func New (a Any, n uint) BoundedBuffer { return new_(a,n) }
