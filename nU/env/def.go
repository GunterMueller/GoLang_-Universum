package env

// (c) Christian Maurer   v. 171126 - license see nU.go

// Liefert das 1. Byte des 1. Parameters des Programmaufrufs,
// wenn der gegeben war, andernfalls 0.
func Par1() byte { return par1() }

// Liefert die Anzahl der Parameter des Programmaufrufs.
func NPars() uint { return nPars() }

// Liefert den Namen des Rechners,
// auf dem der aufrufende Prozess gestartet wurde.
func Localhost() string { return localhost() }

// Returns the name of the call.
func Call() string { return call() }
