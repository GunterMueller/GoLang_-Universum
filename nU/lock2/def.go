package lock2

// (c) Christian Maurer   v. 210123 - license see nU.go

// Funktionen zum Schutz kritischer Abschnitte, die
// durch Aufrufe dieser Funktionen von anderen Prozessen
// nicht unterbrochen werden können.

type
  Locker2 interface {

// Vor.: p < 2.
//       Der aufrufende Prozess ist nicht im kritischen Abschnitt.
// Er ist es als einziger Prozess.
  Lock (p uint)

// Vor.: p < 2.
//       Der aufrufende Prozess ist im kritischen Abschnitt.
// Er ist es jetzt nicht mehr.
  Unlock (p uint)
}

// Liefern neue offene Schlösser für 2 Prozesse.
// mit einer Implementierung, die aus dem Namen hervorgeht.
func NewDekker() Locker2 { return newDekker() }
func NewPeterson() Locker2 { return newPeterson() }
func NewDoranThomas() Locker2 { return newDoranThomas() }
func NewKessels() Locker2 { return newKessels() }
