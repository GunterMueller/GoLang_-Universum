package fahrweg

// (c) Christian Maurer   v. 230305 - license see µU.go

type
  Weg interface {

// Liefert die Nummer des Startblocks von x.
  Start() uint

// Liefert die Nummer des Zielblocks von x.
  Ziel() uint

// x enthält keine Blöcke.
  Clr()

// Der Block mit der Nummer n ist in x eingeordnet.
  Insert (n uint)

// Liefert die Nummer des i-ten Blocks von x.
  Nr (i uint) uint

// Liefert die Anzahl der Blöcke in x.
  Num() uint

// Liefert genau dann true, wenn der i-te Block in x
// kleiner als der j-te Block in x ist.
  Less (i, j int) bool

// Liefert genau dann true, wenn der i-te Block in x
// kleiner als der j-te Block in x ist oder mit ihm übereinstimmt.
  Leq (i, j int) bool

// Liefert genau dann true, wenn in x Weichen oder Doppelkreuzngsweichen
// mit einer ablenkenden Stellung (links oder rechts) vorkommen.
  Ablenkend() bool
}

// Liefert einen neuen leeren Fahrweg.
func New() Weg { return new_() }
