package macc

// (c) Christian Maurer   v. 201102 - license see µU.go

type
  MAccount interface { // A multitasking capable account.
                       // The functions Deposit and Draw cannot be interrupted
                       // by calls of these functions of other processes.

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
func NewMon() MAccount { return newMon() }

// Implementation with a far monitor.
func NewFMon (h string, p uint16, s bool) MAccount { return newFMon(h,p,s) }
