package pos

// (c) murus.org  v. 161216 - license see murus.go

import (
  . "murus/obj"
    "murus/col"
)
type
// Positions on the screen with surrounding areas.
  Position interface {

  Object // Empty means content is empty

// x has the center (x, y).
  Set (x, y int)

// Returns the coordinates of x.
  Pos() (int, int)

// Returns the width and height of x.
  Contour() (uint, uint)

// The center of x has the position of the mouse.
  Mouse()

// Returns true, iff the mouse has the position of the center of x.
  UnderMouse() bool

// f and b are the colours of x.
  Colours (f, b col.Colour)

// Returns the colours of x.
  Cols() (col.Colour, col.Colour)

// x is written to the screen at its position (evtl. transparent),
// If x.w * x.h > 0, then an ellipse is drawn, that surrounds the
// rectangle around x defined by the size of x.
  Write()

// x has the content, that was edited by the user.
//  Edit() // TODO

  Save()
  Restore()
}
// Returns a new position of size (width, height) = (w, h) (in pixels).
func New (w, h uint) Position { return newPos(w,h) }
