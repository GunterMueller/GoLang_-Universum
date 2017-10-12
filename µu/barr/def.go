package barr

// (c) Christian Maurer   v. 170627 - license see µu.go
//
//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 102, 143, 165

type
  Barrier interface { // Points of synchronisation goroutines have to wait for
                      // after having done their part of a common task, until
                      // a certain number M of goroutines have done their part.

// The number of goroutines waiting for the calling barrier is incremented.
// The calling goroutine was evtl. blocked, until the number of waiting goroutines
// equals the length of x.
// Now no goroutines are waiting for the calling barrier.
// The method is atomic, i.e. it cannot be interrupted by other goroutines.
  Wait ()
}

// Returns a new barrier of length n.
func New (n uint) Barrier { return new_(n) }

func NewMon (n uint) Barrier { return newMon(n) }

func NewGo(n uint) Barrier { return newG(n) }
