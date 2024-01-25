package macc

// (c) Christian Maurer   v. 210123 - license see nU.go

type
   MAccount interface { // Ein multitasking-fähiges Konto.
                        // Die exportierten Funktionen können von Aufrufen dieser Funktionen
                        // durch andere Goroutinen nicht unterbrochen werden.

// Das Guthaben von x ist um a erhöht.
// Liefert das neue Guthaben von x.
  Deposit (a uint) uint

// Das Guthaben von x ist um a erniedrigt.
// Liefert das neue Guthaben von x.
// Der aufrufende Prozess ist ggf. solange blockiert,
// bis das Guthaben von x größergleich a ist.
  Draw (a uint) uint
}

// Alle Konstruktoren liefern neue Konten mit dem Guthaben 0.

// Implementierung mit sync Cond's.
func New() MAccount { return new_() }

// Implementierung mit einem universellen Monitor.
func NewM() MAccount { return newM() }

// Implementierung mit einem fernen Monitor.
func NewFM (h string, p uint16, s bool) MAccount { return newFM(h,p,s) }
