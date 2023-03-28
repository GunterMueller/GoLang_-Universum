package signal

// (c) Christian Maurer   v. 230107 - license see µU.go

import (
  . "µU/obj"
  . "bahn/kilo"
)
type
  Typ byte; const (
  T0 = Typ(iota)
  T1 // Hp0, Hp1
  T2 // Hp0, Hp1, Hp2
  NT
)
type
  Stellung byte; const (
// Hauptsignale:
  Hp0 = Stellung(iota) // Halt
  Hp1  // Fahrt
  Hp2  // Langsamfahrt
  NS
)
type
  Signal interface {

  Object

// x ist definiert; x hat die Nummer n, die Kilometrierung K, den Type t,
// die Position (z, s) und die stellung Zughalt. 
  Definieren (n uint, t Typ, k Kilometrierung, st Stellung, z, s uint)

// Liefert den Signaltyp von x.
  Signaltyp() Typ

// Vor.: x is defined.
// x hat die Stellung s. x ist an seiner Position auf dem Bildschirm ausgegeben.
  Stellen (s Stellung)

// Wenn x definiert ist, ist es an seiner Position auf dem Bildschirm ausgegeben.
  Ausgeben()
}

func New() Signal { return new_() }
