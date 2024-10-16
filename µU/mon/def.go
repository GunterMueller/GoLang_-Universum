package mon

// (c) Christian Maurer   v. 241003 - license see µU.go

import
  . "µU/obj"
type
  Monitor interface {

// Pre: i < number of monitor functions of x.
// The calling process may have been blocked on i
// until enough other processes have called Signal(i).
  Wait (i uint)

// Pre: i < number of monitor functions of x.
// If there are processes that in the moment of the call
// were blocked in x on i, exactly one of them is deblocked.
  Signal (i uint)

// Pre: i < number of monitor functions of x.
// All in x on i blocked processe are unblocked.
  SignalAll (i uint)

// Pre: i < number of monitor functions of x.
// Returns the number of processes that
// are blocked at the moment of the call in x on i.
// Remark: See remark on the function Awaited.
  Blocked (i uint) uint

// Pre: i < number of monitor functions of x.
// Returns true, iff at the moment of the call
// there are processes that are blocked in x on i.
// Remark: See remark on the function Awaited
//         in the specification of monitors.
  Awaited (i uint) bool

// Pre: i < number of monitor functions of x.
//      a == nil or a is the object to be processed.
// Returns the value of the i-th function for the argument a
// after the calling  process may have been blocked according to the
// calls of Wait(i) and Signal(i) or SignalAll(i) in the functions of x
// (where f are the monitor function of x and a an object, ggf. nil).
// The function cannot be interrupted by monitor functions of other processes.
  F (a any, i uint) any
}

// Pre: n > 0. f is defined for all i < m.
// Returns a new Monitor with Signal-Urgent-Wait-semantics
// with n functions and the functions f(-, i) for all i < n.
// Clients are responsible for synchronizing the conditions
// with suitable calls of Wait, Signal and SignalAll themselves.
func New (n uint, f FuncSpectrum) Monitor { return new_(n,f) }
