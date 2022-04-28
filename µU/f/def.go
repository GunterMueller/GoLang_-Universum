package f

// (c) Christian Maurer   v. 220420 - license see µU.go

const (
  Plus = uint(iota)
  Times
  ToThe
)

// Spec: see below.
func G (a any, i uint) any { return f(a,Times) }

// Pre: a is a coded object of type UintStream,
// containing exactly 2 numbers a and b.
// Returns for i = 0, 1, 2 resp. the value a+b, a*b, a^b.
func F (a any, i uint) any { return f(a,i) }
