package mon

// (c) murus.org  v. 130726 - license see murus.go

//     Nichtsequentielle Programmierung mit Go 1 kompakt,
//     Kap. 6, insbesondere Abschnitt 6.15

import
  . "murus/obj"
type
  Monitor interface {
// Specifications: Buy my book and read chapter 6.

  Wait (i uint)

  Awaited (i uint) bool

  Signal (i uint)

  SignalAll (i uint)

//  Func (i uint, a Any) Any // deprecated, replaced by F
//                           // with reversed parameter order
  F (a Any, i uint) Any
  S (a Any, i uint, c chan Any) // experimental
}
