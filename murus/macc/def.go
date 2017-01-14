package macc

// (c) murus.org  v. 140328 - license see murus.go

import
  "murus/euro"
type
  MAccount interface { // A multitasking capable account.
                       // The exported functions cannot be interrupted
                       // by calls of these functions of other goroutines.

// The balance of x is incremented by e.
// Returns the new balance of x.
  Deposit (e euro.Euro) euro.Euro

// The balance of x is decremented by e.
// Returns the new balance of x.
// The calling process was blocked, until the balance of x was greater or Equal than e.
  Draw (e euro.Euro) euro.Euro
}
