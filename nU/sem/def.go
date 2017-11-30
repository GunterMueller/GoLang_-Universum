package sem

// (c) Christian Maurer   v. 170411 - license see nU.go

type Semaphore interface { // Protocols for critical sections.

// The calling process is inside the critical section among at most n-1 other processes.
// where n is the number of allowed processes to enter.
// It might have been delayed, until this was possible.
  P()

// The calling process is outside the critical section.
  V()

// P and V cannot be interrupted by calls of these functions of other processes.
}

// All constructors return a new Semaphore, that allows
// exactly n processes to enter the critical section:

// Naive incorrect solution
func NewNaive (n uint) Semaphore { return newNaive(n) }

// Corrected naive solution
func NewCorrect (n uint) Semaphore { return newCorrect(n) }

// Elegant solution with asynchronous message passing
func New (n uint) Semaphore { return new_(n) }

// Implementation of the Go authors
func NewGo (n int) Semaphore { return newGo(n) }

// Solution with the algorithm of Barz
func NewBarz (n uint) Semaphore { return newBarz(n) }

// Solution with synchronous message passing
func NewChannel (n uint) Semaphore { return newChannel(n) }

// Solution with guarded select
func NewGSel (n uint) Semaphore { return newGSel(n) }
