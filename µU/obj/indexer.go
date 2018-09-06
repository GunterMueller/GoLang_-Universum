package obj

// (c) Christian Maurer   v. 180902 - license see µU.go

type
  Indexer interface {

  Object
  Editor

  Index() Func
  RotOrder()
}

func IsIndexer (a Any) bool {
  if a == nil { return false }
  _, ok := a.(Indexer)
  return ok
}
