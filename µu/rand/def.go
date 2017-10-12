package rand

// (c) Christian Maurer   v. 170327 - license see µu.go

// The random generator is initialized.
func Init() { _init() }

// Returns for n == 0 a random number < 2^32,
// otherwise a random number < n.
func Natural (n uint) uint { return natural(n) }

// Returns for i == 0 || i == -2^31 a random number x with |x| < 2^31,
// otherwise a random number x with |x| < |i|.
func Integer (i int) int { return integer(i) }

// Returns a random number x with 0 <= x < 1.
func Real() float64 { return real() }

// Returns true in the long run every second call.
func Even() bool { return even() }
