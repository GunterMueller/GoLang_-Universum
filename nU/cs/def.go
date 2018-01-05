package cs

// (c) Christian Maurer   v. 171230 - license see nU.go

import . "nU/obj"

type CriticalSection interface {

// Vor.: i < Anzahl der Klassen von x.
//       Die Funktion ist innerhalb der Eintrittsbedingungen
//       von x aufgerufen (s. Bem.).
// Liefert genau dann true, wenn im Moment des Aufrufs
// mindestens ein Prozess der k-ten Klasse von x blockiert ist.
// Bem.: Der Wert kann sich sofort nach dem Aufruf geändert
//       haben; er ist folglich nur verwertbar,
//       wenn die Unteilbarkeit des Aufrufs von
//       späteren Bearbeitungen sichergestellt ist.
//       Das ist der Fall bei Einhaltung der Voraussetzung.
  Blocked (i uint) bool

// Vor.: i < Anzahl der Klassen von x.
//       Der aufrufende Prozess befindet sich nicht in x.
// Er befindet sich jetzt in der i-ten Klasse von x, d.h.,
// er war ggf. solange blockiert, bis c(i) galt, und jetzt ist
// e(i) ausgeführt (wobei c die Eintrittsbedingung von x
// und e die Funktion beim Eintritt in x sind).
// Liefert den Wert von e(i).
  Enter (i uint) uint

// Vor.: i < Anzahl der Klassen von x.
//       Der aufrufende Prozess befindet sich
//       in der i-ten Klasse von x.
// Er befindet sich nicht mehr in x, d.h., o(a, i) ist
// ausgeführt (wobei o die Bearbeitung beim Austritt aus x
// und i die Klasse von x, in der sich der aufrufende
// Prozess befand, sind). c(i) ist nicht mehr sichergestellt.
  Leave (i uint)
}

// Vor.: n > 1. c, e und l sind für alle i < n definiert.
// Liefert einen neuen kritischen Abschnitt mit n Klassen,
// x hat die Eintrittsbedingung c(i) und die Bearbeitungen
// e(_, k) und l(i) zum Eintritt in die i-te Klasse von x
// bzw. zum Austritt aus ihr zum Gebrauch von nebenläufigen
// Prozessen. Alle Eintrittsbedingungen für x sind erfüllt;
// kein Prozess befindet sich in x.
  func New (n uint, c CondSpectrum, e NFuncSpectrum, l StmtSpectrum) CriticalSection {
    return new_(n,c,e,l)
  }
