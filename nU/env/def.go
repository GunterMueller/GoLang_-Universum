package env

// (c) Christian Maurer   v. 171227 - license see nU.go

// Liefert das 1. Byte des 1. Arguments des Programmaufrufs,
// wenn das gegeben war, andernfalls 0.
func Arg1() byte { return arg1() }

// Liefert das i-te Argument des Programmaufrufs,
// wenn das gegeben war, andernfalls "".
func Arg (i uint) string { return arg(i) }

// Liefert den Namen des Rechners,
// auf dem der aufrufende Prozess gestartet wurde.
func Localhost() string { return localhost() }
