package shape

// (c) Christian Maurer   v. 140131 - license see murus.go

type
  Shape byte; const (
  Off = Shape(iota)
  Understroke
  Block
  NShapes
)

func Cursor (x, y, h uint, c, s Shape) (uint, uint) { return cursor(x,y,h,c,s) }
