package ego

// (c) Christian Maurer   v. 170424 - license see nU.go

// Returns the value of the first argument of the programm call,
// that uses this function, if that is a natural number < n.
// Panics otherwise.
//func Ego (n uint) uint { return ego(n) }

// Returns the value of the first argument of the programm call,
// that uses this function, if that is a natural number.
// Returns otherwise uint(1<<16).
func Me() uint { return me() }
