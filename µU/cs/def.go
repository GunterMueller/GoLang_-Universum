package cs

// (c) Christian Maurer   v. 171019 - license see µU.go

/* Conditioned critical sections to ensure consistency of common ressources
   used by processes of at most two classes s.t. the critical sections
   can be entered concurrently by several processes of the same class,
   but by processes of different classes only under mutual exclusion.
   The classes are identified by natural numbers, starting with 0.

   Mutual exclusion is guaranteed by conditions and by statemantes,
   that control these conditions, depending on the class.
   The are bundled in functions of type CondSpectrum and StmtSpectrum,
   that have to be constructed by clients and passed to the constructor.

   The functions Enter and Leave cannot be interrupted
   by calls of these functions of other processes. */

import
  . "µU/obj"
type
  CriticalSection interface {
// In the following x means always the calling critical section.

// Pre: i < number of classes of x.
//      The function is called within the entry conditions of x (see remark).
// Returns true, iff at least one process of the i-th class of x
// is blocked at the moment of the call.
// Bemark: The result can be different immediately after the call;
//         so it is only usable if the call and the following evaluation
//         are atomic, which is the case, if the precondition holds.
  Blocked (i uint) bool

// Pre: i < number of classes of the x.
//      The calling process is not in x.
// It is now in the i-th class of x, i.e. it was eventually blocked,
// until c(i) was true, and now e(i) is executed (where c is the
// condition of x and e the processing during the entry into x).
// Returns the result of e(i)
  Enter (i uint) uint

// Pre: i < number of classes of the x.
//      The calling process is in the i-th class of x.
// It is now not any more in x, i.e. l(i) has been executed
// (where l is the processing at the exit from x), 
// and i the class of x, in which the calling process was before.)
  Leave (i uint)
}

// Returns a new conditinal critical section with n classes
// and the conditions and enter-/leave-functions given
// by c, e and l to be used by concurrent processes.
func New (n uint, c CondSpectrum, e NFuncSpectrum, l StmtSpectrum) CriticalSection {
  return new_(n,c,e,l) }
