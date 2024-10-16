package barr

// (c) Christian Maurer   v. 240930 - license see µU.go

// Points of synchronisation processes have to wait for
// after having done their part of a common task, until
// a certain number M of processes have done their part.

type
  Barrier interface {

// The number of processes waiting for the calling barrier is
// incremented. The calling process might have been delayed,
// until the number of waiting processes equals the length of x.
// Now no processes are waiting for the calling barrier.
// The function cannot be interrupted by other processes.
  Wait()
}

// Pre: n > 1.
// All constructors return a new barrier of length n.
func New (n uint) Barrier { return new_(n) }

func NewMon (n uint) Barrier { return newMon(n) }

func NewGo(n uint) Barrier { return newG(n) }
