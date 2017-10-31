package acc

// (c) Christian Maurer   v. 171020 - license see µU.go

import
  "µU/host"
type
  Account interface { // A multitasking capable account.
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
func New() Account { return new_() }

// Implementation with a universal monitor.
func NewM() Account { return newM() }

// Implementation with message passing
func NewCh() Account { return newCh() }

// Implementation with a far monitor.
func NewFMon (h host.Host, p uint16, s bool) Account { return newFM(h,p,s) }
