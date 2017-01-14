package sem

// (c) murus.org  v. 140216 - license see murus.go

//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 53ff., 70 ff., 149, 170, 185

type
  Semaphore interface { // Protocols for critical sections.
                        // Constructors are "I(n uint)";
                        // they allow n processes to enter the critical section.

// The calling process is inside the critical section among at most n-1 other processes.
// where n is the number of allowed processes to enter.
// It might have been delayed, until this was possible.
  P()

// The calling process is outside the critical section.
  V()

// P and V cannot be interrupted by calls of these functions of other processes.
}
