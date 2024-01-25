package barr

// (c) Christian Maurer   v. 170627 - license see nU.go

type
  Barrier interface {

// Die Anzahl der Prozesse, die auf die aufrufende Barriere warten, ist inkrementiert.
// Der aufrufende Prozess war ggf. blockiert, bis die Anzahl der wartenden Prozesse
// mit der Länge der Barriere übereinstimmt.
// Jetzt warten keine Prozesse mehr auf die aufrufende Barriere.
// Die Methode kann von anderen Prozessen nicht unterbrochen werden.
  Wait ()
}

// Liefert eine neue Barriere der Länge n.
func New (n uint) Barrier { return new_(n) }

func NewM (n uint) Barrier { return newM(n) }

func NewGo(n uint) Barrier { return newG(n) }
