package obj

// (c) Christian Maurer   v. 140425 - license see murus.go

type
  Indexer interface {

  Editor

  Index() Func
  RotOrder()
}
