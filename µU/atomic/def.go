package atomic

// (c) Christian Maurer   v. 171023 - license see µU.go

// Tools for the construction of locking algorithms.
// The execution of any of these cannot be interrupted
// by other goroutines that call the same function.

import
  "sync/atomic"

// *n = k. Returns the former value of n
func Exchange (n *uint32, k uint32) uint32 {
  return atomic.SwapUint32 (n, k)
}

// Returns true, if *n = k formerly.
// In this case now *n = m, otherwise nothing has changed.
func CompareAndSwap (n *uint32, k, m uint32) bool {
  return atomic.CompareAndSwapUint32 (n, k, m)
}

// Pre: n + 1 < math.MaxUint32.
// n is incremented by 1. Returns the former value of n.
func FetchAndIncrement (n *uint32) uint32 { return atomic.AddUint32 (n, 1) - 1 }

// *n = k.
// func Store (n *uint32, k uint32) { atomic.StoreUint32 (n, k) }

// a = true. b is the former value of a.
func TestAndSet (a *bool) (b bool)

// Pre: n + k < math.MaxUint32.
// n is incremented by k. m is the former value of n.
func FetchAndAdd (n *uint32, k uint32) (m uint32)

// Pre: n > math.MinInt32.
// n is decremented by 1. b is true iff now n < 0.
func Decrement (n *int32) (b bool)
