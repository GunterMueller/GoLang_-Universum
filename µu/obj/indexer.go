package obj

// (c) Christian Maurer   v. 170919 - license see µu.go

type
  Indexer interface {

  Object
  Editor

  Index() Func
  RotOrder()
}
