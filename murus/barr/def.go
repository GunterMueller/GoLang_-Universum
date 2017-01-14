package barr

// (c) murus.org  v. 120910 - license see murus.go
//
//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 102, 143, 165

type
  Barrier interface { // Points of synchronisation goroutines have to wait for
                      // after having done their part of a common task, until
                      // a certain number M of goroutines have done their part.

// The number of goroutines waiting for the calling barrier is incremented.
// The calling goroutine was evtl. blocked, until the number of waiting goroutines equals
// the number, which was given to the calling barriere as parameter to New.
// Now no goroutines are waiting for the calling barrier.
// The method is atomic, i.e. it cannot be interrupted by other goroutines.
  Wait ()
}
