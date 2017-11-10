package macc

// (c) Christian Maurer   v. 171107 - license see µU.go

import
  "µU/host"
type
  MAccount interface { // A multitasking capable account.
                       // The exported functions cannot be interrupted
                       // by calls of these functions of other goroutines.

// The balance of x is incremented by a.
// Returns the new balance of x.
  Deposit (a uint) uint

// The balance of x is decremented by a.
// Returns the new balance of x.
// The calling process was blocked, until the balance of x was greater or Equal than a.
  Draw (a uint) uint
}

// All constructors return new accounts with balance 0.

// Implementation with sync Cond's.
func New() MAccount { return new_() }

// Implementation with a universal monitor.
func NewM() MAccount { return newM() }

// Implementation with message passing
func NewCh() MAccount { return newCh() }

// Implementation with a far monitor.
func NewFM (h host.Host, p uint16, s bool) MAccount { return newFM(h,p,s) }
