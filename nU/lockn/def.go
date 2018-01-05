package lockn

// (c) Christian Maurer   v. 171231 - license see nU.go

// Die Funktionen Lock und Unlock können durch Aufrufe von
// Lock oder Unlock durch andere Goroutinen nicht unterbrochen werden.

type LockerN interface {

// Vor.: Der aufrufende Prozess ist nicht im kritischen Abschnitt.
// Er ist jetzt als einziger im kritischen Abschnitt.
  Lock (p uint)

// Vor.: Der aufrufende Prozess ist im kritischen Abschnitt.
// Er ist nicht im kritischen Abschnitt.
  Unlock (p uint)
}

// Liefern neue unverschlossene Schlösser für 2 Prozesse
// mit einer Implementierung, die ihr Name verrät.
func NewDekker() LockerN { return newDe() }
func NewDoranThomas() LockerN { return newDT() }
func NewKessels2() LockerN { return newK2() }
func NewPeterson() LockerN { return newP() }

// Liefern neue unverschlossene Schlösser für n Prozesse
// mit einer Implementierung, die ihr Name verrät.
func NewBakery (n uint) LockerN { return newB(n) }
func NewBakery1 (n uint) LockerN { return newB1(n) }
func NewDijkstra (n uint) LockerN { return newD(n) }
func NewHabermann (n uint) LockerN { return newH(n) }
func NewKesselsN (n uint) LockerN { return newKN(n) }
func NewTicket (n uint) LockerN { return newT(n) }
func NewTiebreaker (n uint) LockerN { return newTb(n) }
func NewSzymanski (n uint) LockerN { return newSz(n) }
func NewFast (n uint) LockerN { return newF(n) }
