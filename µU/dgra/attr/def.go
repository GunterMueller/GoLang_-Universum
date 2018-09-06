package attr

// (c) Christian Maurer  v. 180829 - license see µU.go

type
  Attribute byte; const (
  Open = Attribute(iota)
  Closed
  Branch
  In
  Out
  Up
  Down
  NAttributes
)
