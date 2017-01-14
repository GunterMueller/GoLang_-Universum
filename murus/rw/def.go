package rw

// (c) murus.org  v. 120910 - license see murus.go

//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 75 ff., 85 ff., 137 ff. u.v.a.

type
  ReaderWriter interface { // protocols for the reader writer problem

// Pre: The calling goroutine is neither reading or writing.
// The calling goroutine is reading; no goroutine is writing.
// If at the time of the call there was a writing goroutine,
// the goroutine has been delayed, until there was no writing goroutine.
  ReaderIn ()

// Pre: The calling goroutine is reading.
// The calling goroutine is neither reading or writing.
  ReaderOut ()

// Pre: The calling goroutine is neither reading or writing.
// The calling goroutine is writing;
// no other goroutine is writing and there are no reading goroutines.
// If at the time of the call there were reading goroutines or a writing one,
// the calling goroutine has been delayed, until there were no reading or writing goroutines.
  WriterIn ()

// Pre: The calling goroutine is writing.
// The calling goroutine is neither reading or writing.
  WriterOut ()
}
