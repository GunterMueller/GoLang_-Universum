package obj

// (c) Christian Maurer   v. 180902 - license see µU.go

type
  Clearer interface {

// Returns true, iff the calling object is empty. What "empty" actually means,
// depends on the very semantics of the type of the objects considered.
// If that type is e.g. a collector, empty means "containing no objects"; otherwise it is
// normally an object with undefined value, represented by strings consisting only of spaces. 
  Empty() bool

// The calling object is empty.
  Clr()
}

// Returns true, iff a implements Clearer.
func IsClearer (a Any) bool { return isClearer(a) }

// TODO Spec
func Clear (a Any) Any { return clear(a) }
