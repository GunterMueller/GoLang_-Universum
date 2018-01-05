package atomic

// (c) Christian Maurer   v. 171231 - license see nU.go

// Die Ausführung jeder Funktion kann nicht durch den Aufruf der
// gleichen Funktion von einem anderen Prozess unterbrochen werden.

import "sync/atomic"

// *n = k. Liefert den vorigen Wert von n.
func Exchange (n *uint32, k uint32) uint32 {
  return atomic.SwapUint32 (n, k)
}

//// Liefert true, wenn vorher *n = k war.
//// In diesem Fall gilt jetzt *n = m, andernfalls ist nichts verändert.
//func CompareAndSwap (n *uint32, k, m uint32) bool {
//  return atomic.CompareAndSwapUint32 (n, k, m)
//}

// Vor.: n + 1 < math.MaxUint32.
// n ist um 1 erhöht. Liefert den vorigen Wert von n.
func FetchAndIncrement (n *uint32) uint32 { return atomic.AddUint32 (n, 1) - 1 }

// a = true. b ist der vorherige Wert von a.
func TestAndSet (a *bool) (b bool)

// Vor.: n + k < math.MaxUint32.
// n ist um k erhöht. m ist der vorige Wert von n.
func FetchAndAdd (n *uint32, k uint32) (m uint32)

// Vor.: n > math.MinInt32.
// n ist um 1 erniedrigt. b ist genau dann true, wenn jetzt n < 0 ist.
func Decrement (n *int32) (b bool)
