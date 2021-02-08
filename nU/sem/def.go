package sem

// (c) Christian Maurer   v. 210123 - license see nU.go

type Semaphore interface {
// Ganzzahlige Werte als Zugangsprotokolle zu kritischen Abschnitten
// zum nebenläufigen Zugriff mehrerer Prozesse auf gemeinsame Daten.
// Die Methoden P und V sind durch Aufrufe von P oder V
// von anderen Prozessen nicht unterbrechbar.

// Der aufrufende Prozess ist neben höchstens n-1 weiteren Prozessen
// im kritischen Abschnitt, wobei n der dem Konstruktor übergebene
// Initialwert des Semaphors ist.
  P()

// Der aufrufende Prozess ist nicht mehr im kritischen Abschnitt.
  V()
}

// Alle Konstruktoren liefern ein neues Semaphor,
// das höchstens n Prozesse nebenläufig in den kritischen
// Abschnitt lässt. Kein Prozess ist im kritischen Abschnitt.

// Falsche naive Lösung
func NewNaive (n uint) Semaphore { return newNaive(n) }

// Korrigiert naive Lösung
func NewCorrect (n uint) Semaphore { return newCorrect(n) }

// Elegante Lösung mit asynchronem Botschaftenaustausch
func New (n uint) Semaphore { return new_(n) }

// Implementierung der Go-Autoren
func NewGo (n int) Semaphore { return newGo(n) }

// Lösung mit dem Algorithmus von Barz
func NewBarz (n uint) Semaphore { return newBarz(n) }

// Lösung mit synchronem Botschaftenaustausch
func NewChannel (n uint) Semaphore { return newCh(n) }

// Lösung mit bewachtem Warten
func NewGSel (n uint) Semaphore { return newGS(n) }
