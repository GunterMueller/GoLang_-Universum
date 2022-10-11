package obj

// (c) Christian Maurer   v. 220808 - license see µU.go

type
  Rotator interface {

  Indexer

// The actual index is another one, if there is not only _one_ actual index.
  Rotate()
}

func IsRotator (a any) bool {
  if a == nil { return false }
  _, ok := a.(Rotator)
  return ok
}
