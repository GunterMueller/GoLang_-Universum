package macc

// (c) Christian Maurer   v. 241019 - license see µU.go

type
  MAccount interface { // A multitasking capable account.
                       // All functions cannot be interrupted
                       // by calls of these functions of other processes.

// The balance is incremented by a.
// Returns the new balance of x.
  Deposit (a uint) uint

// If the balance is >= a, it is decremented by a.
// In this case the the new balance is returned,
// otherwise 0 is returned.
  Draw (a uint) uint

// Returns the actual balance.
  Show (a uint) uint
}

// Implementation with a far monitor.
func New (h string, p uint16, s bool) MAccount { return new_(h,p,s) }
