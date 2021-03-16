package shape

// (c) Christian Maurer   v. 210311 - license see µU.go

type
  Shape byte; const (
  Off = Shape(iota)
  Understroke
  Block
  NShapes
)

// Returns the coordinates of the upper left corner and the height
// of the rectangle for all combinations (c, s) with c != s.
func Cursor (x, y, h uint, c, s Shape) (uint, uint) { return cursor(x,y,h,c,s) }
