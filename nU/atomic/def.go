package atomic

// (c) Christian Maurer   v. 210123 - license see nU.go

// Funktionen für die Konstruktion von Schlossalgorithmen.
// Die Ausführung jeder dieser Funktionen kann nicht von anderen
// Goroutinen unterbrochen werden, die die gleiche Funktion aufrufen.

// *a = true. Liefert den vorherigen Wert von *a.
func TestAndSet (a *bool) bool

// *n = k. Liefert den vorherigen Wert von n.
func Exchange (n *uint, k uint) uint

// Liefert genau dann true, wenn vorher *n = k war.
// In diesem Fall ist jetzt *n = m, andernfalls ist nichts verändert.
func CompareAndSwap (n *uint, k, m uint) bool

// Vor.: *n + 1 < math.MaxUint32 bzw. math.MaxUint64.
// *n ist um 1 erhöht. Liefert den vorherigen Wert von n.
func FetchAndIncrement (n *uint) uint

// Vor.: n + k < math.MaxUint32 bzw. < math.MaxUint64.
// *n ist um k erhöht.
func Add (n *uint, k uint)

// Vor.: n + k < math.MaxUint32 bzw. < math.MaxUint64.
// *n ist um k erhöht. Liefert den vorherigen Wert von *n.
func FetchAndAdd (n *uint, k uint) uint

// Vor.: n > math.MinInt32 bzw. n > math.MinInt64.
// *n ist um 1 erniedrigt. Liefert genau dann true, wenn jetzt *n < 0.
func Decrement (n *int) bool

// *n = k.
func Store (n *uint, k uint)
