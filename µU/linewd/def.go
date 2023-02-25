package linewd

// (c) Christian Maurer   v. 230112 - license see µU.go

type
  Linewidth byte; const (
  Thin = Linewidth(iota)
  Thick
  Thicker
  VeryThick
  Fat
)
