package obj

// (c) Christian Maurer   v. 220420 - license see µU.go

type
  Indexer interface {

  Object
  Editor

  Index() Func
}

func IsIndexer (a any) bool {
  if a == nil { return false }
  _, ok := a.(Indexer)
  return ok
}
