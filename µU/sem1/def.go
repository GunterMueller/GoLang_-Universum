package sem1

// (c) Christian Maurer   v. 240930 - license see µU.go

// Protocols for access to a ritical sections.
// The functions P and V cannot be interrupted
// by calls of these functions of other processes.
type
  Semaphore1 interface {

// Pre: The calling process is not in the critical section.
// The calling process is the only one in the critical section.
// It might have been delayed, until this was possible.
  P()

// Pre: The calling process is in the critical section.
// The calling process is not in the critical section.
  V()
}

func New() Semaphore1 { return new_() }
