package lock

// (c) Christian Maurer   v. 210123 - license see nU.go

// Funktionen zum Schutz kritischer Abschnitte, die
// durch Aufrufe dieser Funktionen von anderen Prozessen
// nicht unterbrochen werden können.

type
  Locker interface {

// Vor.: Der aufrufende Prozess ist nicht im kritischen Abschnitt.
// Er ist es als einziger Prozess.
  Lock()

// Vor.: Der aufrufende Prozess ist im kritischen Abschnitt.
// Er ist es jetzt nicht mehr.
  Unlock()
}

// Liefern neue offene Schlösser
// mit einer Implementierung, die aus dem Namen hervorgeht.
func NewChannel() Locker { return newChan() }
func NewTAS() Locker { return newTAS() }
func NewXCHG() Locker { return newXCHG() }
func NewCAS() Locker { return newCAS() }
func NewDEC() Locker { return newDEC() }
func NewMutex() Locker { return newMutex() }
func NewUdding() Locker { return newUdding() }
func NewMorris() Locker { return newMorris() }
