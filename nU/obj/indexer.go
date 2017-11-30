package obj

// (c) Christian Maurer   v. 170919 - license see nU.go

type
  Indexer interface {

  Object
  Editor

  Index() Func
  RotOrder()
}
