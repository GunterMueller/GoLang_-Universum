package lock

// (c) Christian Maurer   v. 171231 - license see nU.go

// Die Funktionen Lock und Unlock können durch Aufrufe von
// Lock oder Unlock durch andere Goroutinen nicht unterbrochen werden.

type Locker interface {

// Vor.: Der aufrufende Prozess ist nicht im kritischen Abschnitt.
// Er ist jetzt als einziger im kritischen Abschnitt.
  Lock()

// Vor.: Der aufrufende Prozess ist im kritischen Abschnitt.
// Er ist nicht im kritischen Abschnitt.
  Unlock()
}

// Liefern neue unverschlossene Schlösser für n Prozesse
// mit einer Implementierung, die ihr Name verrät.
func NewCAS() Locker { return newCAS() }
func NewChannel() Locker { return newChan() }
func NewDEC() Locker { return newDEC() }
func NewFA() Locker { return newFA() }
func NewMorris() Locker { return newMorris() }
func NewMutex() Locker { return newMutex() }
func NewTAS() Locker { return newTAS() }
func NewUdding() Locker { return newUdding() }
func NewXCHG() Locker { return newXCHG() }
