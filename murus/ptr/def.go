package ptr

// (c) Christian Maurer   v. 140131 - license see murus.go

type
  Pointer byte; const (
  Standard = Pointer(iota)
  Gumby
  Hand
  Gobbler
  Watch
  NPointers
)
