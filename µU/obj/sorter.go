package obj

// (c) Christian Maurer   v. 180902 - license see µU.go

type
  Sorter interface {

//
  Ordered () bool

//
  Sort ()

//
  ExGeq (a Any) bool
}

func IsSorter (a Any) bool {
  if a == nil { return false }
  _, ok := a.(Sorter)
  return ok
}
