package atomic

// (c) Christian Maurer   v. 190331 - license see µU.go

// Tools for the construction of locking algorithms.
// The execution of any of these cannot be interrupted
// by other goroutines that call the same function.

// *a = true. Returns the former value of *a.
func TestAndSet (a *bool) bool

// *n = k. Returns the former value of n
func Exchange (n *uint, k uint) uint

// Returns true, if *n = k formerly.
// In this case now *n = m, otherwise nothing has changed.
func CompareAndSwap (n *uint, k, m uint) bool

// Pre: *n + 1 < math.MaxUint32 resp math.MaxUint64.
// *n is incremented by 1. Returns the former value of *n.
func FetchAndIncrement (n *uint) uint

// Pre: n + k < math.MaxUint32 resp. < math.MaxUint64.
// *n is incremented by k.
func Add (n *uint, k uint)

// Pre: n + k < math.MaxUint32 resp. < math.MaxUint64.
// *n is incremented by 1.
func Inc (n *uint)

// Pre: n + k < math.MaxUint32 resp. < math.MaxUint64.
// *n is incremented by k. Returns the former value of *n.
func FetchAndAdd (n *uint, k uint) uint

// Pre: n > math.MinInt32 resp. n > math.MinInt64.
// *n is decremented by 1. Returns true iff now *n < 0.
func Decrement (n *int) bool

// Pre: n > 0.
// *n is decremented by 1.
func Dec (n *uint)

// *n = k.
func Store (n *uint, k uint)
