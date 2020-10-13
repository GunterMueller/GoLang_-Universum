package obj

// (c) Christian Maurer   v. 201011 - license see µU.go

type
  Sorter interface {

  Iterator

// Returns true, iff x is ordered.
  Ordered() bool

// x is ordered.
  Sort()

// Pre: x is ordered.
// Returns true, iff x contains objects b with Leq (a, b).
// In this case, the next Get operation returns the smallest such object.
  ExGeq (a Any) bool
}

func IsSorter (a Any) bool {
  if a == nil { return false }
  _, ok := a.(Sorter)
  return ok
}
