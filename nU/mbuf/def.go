package mbuf

// (c) Christian Maurer   v. 210123 - license see nU.go

import . "nU/obj"

type MBuffer interface { // Synchronisierte Buffer.
                         // Die exportierten Funktionen können nicht von Aufrufen
                         // dieser Funktionen von anderen Goroutinen unterbrochen werden.

// a ist als letztes Objekt in x eingefügt.
// Der aufrufende Prozess war ggf. solange blockiert, bis x nicht voll war.
  Ins (a Any)

// Liefert das erste Objekt von x und dieses Objekt ist jetzt aus x entfernt.
// Der aufrufende Prozess war ggf. solange blockiert, bis x nicht leer war.
  Get() Any
}

// Vor.: a ist atomar oder der Typ von a implementiert Object.
// Liefert einen leeren Puffer für Objekte des Typs von a,
// der von nebenläufigen Prozessen benutzt werden kann.
func New (a Any) MBuffer { return new_(a) }

// TODO Spec
func NewFarMonitor (a Any, h string, p uint16, s bool) MBuffer { return newfm(a,h,p,s) }
