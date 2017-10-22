package ptr

// (c) Christian Maurer   v. 140131 - license see µU.go

type
  Pointer byte; const (
  Standard = Pointer(iota)
  Gumby
  Hand
  Gobbler
  Watch
  NPointers
)
