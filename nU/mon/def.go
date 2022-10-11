package mon

// (c) Christian Maurer   v. 220702 - license see nU.go

import . "nU/obj"

type Monitor interface {

// Vor.: i < Funktionszahl von x.
// Der aufrufende Prozess war ggf. auf i blockiert, bis
// genügend viele andere Prozesse Signal(i) aufgerufen hatten.
  Wait (i uint)

// Vor.: i < Funktionszahl von x.
// Liefert die Anzahl der Prozesse, die
// im Moment des Aufrufs in x auf i blockiert sind.
// Bem.: Siehe Bemerkung bei Awaited.
  Blocked (i uint) uint

// Vor.: i < Funktionszahl von x.
// Liefert genau dann true, wenn es im Moment des Aufrufs
// Prozesse gibt, die in x auf i blockiert sind.
// Bem.: Siehe Bemerkung bei der Funktion Awaited
//       in der Spezifikation von Monitoren.
  Awaited (i uint) bool

// Vor.: i < Funktionszahl von x.
// Wenn es Prozesse gibt, die im Moment des Aufrufs in x auf
// i blockiert sind, ist genau einer von ihnen deblockiert.
  Signal (i uint)

// Vor.: i < Funktionszahl von x.
// Alle in x auf i blockierten Prozesse sind deblockiert.
  SignalAll (i uint)

// Vor.: i < Funktionszahl von x.
//       a == nil oder a ist das zu bearbeitende Objekt.
// Liefert den Wert der i-ten Funktion für das Argument a
// nachdem der aufrufende Prozess ggf. entsprechend den
// Aufrufen von Wait(i) und Signal(i) oder SignalAll(i)
// in den Funktionen von x blockiert war (wobei f die
// Monitorfunktion von x sind und a ein Objekt, ggf. nil).
// Die Funktion kann durch Monitorfunktionen anderer Prozesse
// nicht unterbrochen werden.
  F (a any, i uint) any
}

// Vor.: n > 0. f (und ggf. p) ist für alle i < n definiert.
// Liefert einen neuen Monitor mit der Funktionszahl n und
// den Funktionen f(-, i) für alle i < n.
// Klienten sind für die Bedingungssynchronisation mit
// passenden Aufrufen von Wait, Signal und SignalAll
// selber verantwortlich.
  func New (n uint, f FuncSpectrum) Monitor { return new_(n,f) }
