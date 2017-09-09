package lock

// (c) Christian Maurer   v. 111111 - license see murus.go

// Tools for the construction of locking algorithms for n goroutines.
// The Ablauf of each of the functions cannot be interrupted
// by other goroutines by calling the same function.

// b == true. Returns the old value of b.
func TestAndSet (b *bool) bool

// n == k. Returns the old value of n.
func ExchangeUint32 (n *uint32, k uint32) uint32
func ExchangeInt32 (n *int32, k int32) int32

// Pre: n + k < math.MaxUint32.
// n is incremented by k. Returns the old value of n.
func FetchAndAddUint32 (n *uint32, k uint32) uint32
func FetchAndAddInt32 (n *int32, k int32) int32

// Pre: n > math.MinInt32.
// n is decremented by 1. Returns true iff now n < 0.
func DecrementInt32 (n *int32) bool
