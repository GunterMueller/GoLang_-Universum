package ego

// (c) Christian Maurer   v. 171202 - license see nU.go

// Liefert den Wert des ersten Arguments des Programmaufrufs,
// der diese Funktion benutzt, wenn das eine natürliche Zahl < n ist.
// Panict andernfalls.
func Ego (n uint) uint { return ego(n) }

// Liefert den Wert des ersten Arguments des Programmaufrufs,
// der diese Funktion benutzt, wenn das eine natürliche Zahl ist.
// Liefert andernfalls uint(1<<16).
func Me() uint { return me() }
