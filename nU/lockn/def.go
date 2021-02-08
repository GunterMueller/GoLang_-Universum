package lockn

// (c) Christian Maurer   v. 190325 - license see nU.go

// Funktionen zum Schutz kritischer Abschnitte, die
// durch Aufrufe dieser Funktionen von anderen Prozessen
// nicht unterbrochen werden können.

type
  LockerN interface {

// Vor.: p < Anzahl der durch den Konstruktor definierten Prozesse.
//       Der aufrufende Prozess ist nicht im kritischen Abschnitt.
// Er ist es als einziger Prozess.
  Lock (p uint)

// Vor.: p < Anzahl der durch den Konstruktor definierten Prozesse.
//       Der aufrufende Prozess ist im kritischen Abschnitt.
// Er ist es jetzt nicht mehr.
  Unlock (p uint)
}

// Liefern neue offene Schlösser für n Prozesse.
// mit einer Implementierung, die aus dem Namen hervorgeht.
func NewDijkstra (n uint) LockerN { return newDijkstra(n) }
func NewHabermann (n uint) LockerN { return newHabermann(n) }
func NewBakery (n uint) LockerN { return newBakery(n) }
func NewBakery1 (n uint) LockerN { return newBakery1(n) }
func NewTicket (n uint) LockerN { return newTicket(n) }
func NewTiebreaker (n uint) LockerN { return newTiebreaker(n) }
func NewKessels (n uint) LockerN { return newKessels(n) } // Pre: n is a power of 2.
func NewSzymanski (n uint) LockerN { return newSzymanski(n) }
func NewKnuth (n uint) LockerN { return newKnuth(n) }
func NewDeBruijn (n uint) LockerN { return newDeBruijn(n) }
func NewChannel (n uint) LockerN { return newChannel(n) }
