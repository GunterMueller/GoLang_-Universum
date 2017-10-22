package obj

// (c) Christian Maurer   v. 170919 - license see µU.go

type
  Indexer interface {

  Object
  Editor

  Index() Func
  RotOrder()
}
