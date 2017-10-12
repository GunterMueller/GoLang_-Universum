package cs

// (c) Christian Maurer   v. 130808 - license see µu.go

//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 85 ff.

import
  . "µu/obj"
/* Conditioned critical sections to ensure consistency of common ressources
   used by processes of at most two classes s.t. the critical sections
   can be entered concurrently by several processes of the same class,
   but by processes of different classes only under mutual exclusion.
   The classes are identified by natural numbers, starting with 0.

   Mutual exclusion is guaranteed by conditions and by operations,
   that control these conditions, depending on the class.
   The are bundled in functions of type CondSpectrum and OpSpectrum,
   that have to be constructed by clients and passed to the constructor
     New (n uint, c CondSpectrum, e, l OpSpectrum).

   The functions Enter and Leave cannot be interrupted by calls of these functions
   of other processes. */

type
  CriticalSection interface {
// x means in the following always the calling critical section.

// Pre: k < number of classes of x.
//      The function is called within the entry conditions of x (see remark).
// Returns true, iff at least one process of the k-th class of x
// is blocked at the moment of the call.
// Bemark: The result can be different immediately after the call;
//         so it is only usable if the call and the following evaluation
//         are atomic, which is the case, if the precondition holds.
  Blocked (k uint) bool

// Pre: k < number of classes of the x.
//      The calling process is not in x.
// It is now in the k-th class of x, i.e. it was eventually blocked, until c(k) was true,
// and now e(a, k) is executed (where c is the condition of x
// and e the processing during the entry into x).
  Enter (k uint, a Any)

// Pre: k < number of classes of the x.
//      The calling process is in the k-th class of x.
// It is now not any more in x, i.e. c(a, k) has been executed
// (where l is the processing at the exit from x), 
// and k the class of x, in which the calling process was before.)
// c(k) of x is no longer ensured.
  Leave (k uint, a Any)
}

// Returns a new conditinal critical section with n classes and
// the conditions and enter-/leave-functions given by c, e and l
// to be used by concurrent processes.
func New (n uint, c CondSpectrum, e, l OpSpectrum) CriticalSection { return new_(n,c,e,l) }
