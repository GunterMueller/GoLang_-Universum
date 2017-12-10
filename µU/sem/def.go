package sem

// (c) Christian Maurer   v. 171127 - license see µU.go

type
  Semaphore interface { // Protocols for critical sections.
                        // Neither P nor V can be interrupted by calls
                        // of these functions of other processes.

// The calling process is inside the critical section among at most n-1 other processes.
// where n is the number of allowed processes to enter.
// It might have been delayed, until this was possible.
  P()

// The calling process is outside the critical section.
  V()
}

// All constructors return a new Semaphore, that allows
// exactly n processes to enter the critical section:

// Naive incorrect solution
func NewNaive (n uint) Semaphore { return new_(n) }

// Corrected naive solution
func NewCorrect (n uint) Semaphore { return newC(n) }

// Elegant implementation with asynchronous message passing
func New (n uint) Semaphore { return newS(n) }

// Implementation of the Go authors
func NewGo (n uint) Semaphore { return newG(n) }

// Implementation with the algorithm of Barz
func NewBarz (n uint) Semaphore { return newB(n) }

// Implementation with a universal critical section
func NewCriticalSection (n uint) Semaphore { return newCS(n) }

// Implementation with a universal monitor
func NewMonitor (n uint) Semaphore { return newM(n) }

// Implementation with a conditioned universal monitor
func NewCondMonitor (n uint) Semaphore { return newCM(n) }

// Implementation with synchronous message passing
func NewChannel (n uint) Semaphore { return newCh(n) }

// Implementation with synchronous messsage passing with guarded select
func NewGSel (n uint) Semaphore { return newGS(n) }

// Implementation for distributed use
func NewFMon (n uint, h string, p uint16, s bool) Semaphore { return newFM(n,h,p,s) }
