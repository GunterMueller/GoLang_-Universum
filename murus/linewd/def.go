package linewd

// (c) Christian Maurer   v. 140131 - license see murus.go

type
  Linewidth byte; const (
  Thin = Linewidth(iota)
  Thick
  Thicker
)
