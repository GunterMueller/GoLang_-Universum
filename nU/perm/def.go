package perm

// (c) Christian Maurer   v. 210123 - license see nU.go

type
  Permutation interface { // Permutationen der ersten n natürlichen Zahlen.

// x ist zufällig permutiert.
  Permute ()

// Liefert 0, wenn k >= Größe von x,
// andernfalls die k-te Zahl von x.
  F (k uint) uint
}

// Vor.: n > 1.
// x hat die Größe n, d.h.,
// es ist eine Zufallspermutation der natürlichen Zahlen < n.
func New (n uint) Permutation { return new_(n) }
