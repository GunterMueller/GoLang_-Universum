package zelle

// (c) Christian Maurer   v. 230107 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  . "bahn/kilo"
  . "bahn/richtung"
)
type
  Zelle interface {

  Object

// x hat die Nummer n.
  Numerieren (n uint)

// Vor. für alle Methoden mit den Parametern z, s am Schluss, die eine Zelle definieren:
// Die Position (z, s) ist noch nicht mit einer Zelle belegt.

// x ist ein Gleis mit der Schräglage a.
  Gleis (n uint, a Richtung, z, s uint)

// Liefert genau dann true, wenn x ein Gleis ist.
  IstGleis() bool

// x ist ein Knick in Richtung der Kilometrierung nach a.
  Knick (n uint, k Kilometrierung, a Richtung, z, s uint)

// x ist ein Prellbock in Richtung k und Position (z, s).
  Prellbock (k Kilometrierung, z, s uint)

// x ist für r = Rechts eine Rechtsweiche, andernfalls eine Linksweiche
// mit Verzweigungsrichtung k, Schräglage l, Stellung st und Position (z, s).
  Weiche (n uint, k Kilometrierung, l, r, st Richtung, z, s uint)

// Liefert genau dann (k, true), wenn x eine Weiche mit der Verzweigungsrichtung k ist.
  IstWeiche() (Kilometrierung, bool)

// Vor.: l != Gerade.
// x ist eine Doppelkreuzungsweiche mit Schräglage l, Stellung r und Position (z, s).
  DKW (n uint, l, r Richtung, z, s uint)

// Liefert genau dann (k, true), wenn x eine DKW mit der Verzweigungsrichtung k ist.
  IstDKW() (Kilometrierung, bool)

  String() string

// Wenn x eine Weiche oder DkW ist, ist sie in Richtung r gestellt.
  Stellen (r Richtung)

// Liefert die Stellung von x, wenn x eine Weiche oder DkW ist; andernfalls Gerade.
  Stellung() Richtung

// Liefert die Kilmetrierung von x.
  Kilo() Kilometrierung

// Liefert die Schräglage von x in Richtung k (bei in
// Richtung k verzweigtem x die des durchgehenden Astes), 
// und die Position von x auf dem Bildschirm.
  Schräglage (k Kilometrierung) (Richtung, uint, uint)

// Liefert genau dann true, wenn x die Position (z, s) hat.
  HatPosition (z, s uint) bool

// Liefert die Position von x.
  Pos() (uint, uint)

  Nummer() uint

// x ist auf dem Bildschirm in der Farbe ausgegeben.
  Ausgeben (c col.Colour)

// Liefert genau dann true, wenn der Mauszeiger auf x zeigt.
  UnterMaus() bool
}

func New() Zelle { return new_() }
