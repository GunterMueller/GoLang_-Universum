package mbuf

// (c) Christian Maurer   v. 220702 - license see nU.go

type
   MBuffer interface { // Synchronisierte Buffer.
                       // Die exportierten Funktionen können nicht von Aufrufen
                       // dieser Funktionen von anderen Goroutinen unterbrochen werden.

// a ist als letztes Objekt in x eingefügt.
// Der aufrufende Prozess war ggf. solange blockiert, bis x nicht voll war.
  Ins (a any)

// Liefert das erste Objekt von x und dieses Objekt ist jetzt aus x entfernt.
// Der aufrufende Prozess war ggf. solange blockiert, bis x nicht leer war.
  Get() any
}

// Vor.: a ist atomar oder der Typ von a implementiert Object.
// Liefert einen leeren Puffer für Objekte des Typs von a,
// der von nebenläufigen Prozessen benutzt werden kann.
func New (a any) MBuffer { return new_(a) }

// TODO Spec
func NewFarMonitor (a any, h string, p uint16, s bool) MBuffer { return newfm(a,h,p,s) }
