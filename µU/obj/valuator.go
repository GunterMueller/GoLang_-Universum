package obj

// (c) Christian Maurer   v. 220420 - license see µU.go

type
  Valuator interface {

// Returns the value of x, if IsValuator (x).
// Returns otherwise 1.
  Val() uint

// Pre: IsValuator (x).
// x.Val() == n (% 1 << a, if x has the type uint<a>).
  SetVal (n uint)
}

// Returns true, iff a implements Valuator or has an uint-type.
func IsValuator (a any) bool { return isValuator(a) }

// Pre: IsValuator (a).
// Returns the value of a.
func Val (a any) uint { return val(a) }

// Pre: IsValuator (a).
// x.Val() == n (% 1 << a, if a has the type uint<a>).
func SetVal (a *any, n uint) { setVal(a,n) }
