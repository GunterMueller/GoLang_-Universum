package mouse

// (c) Christian Maurer   v. 220119 - license see µU.go

// >>> This package only serves the implementation of µU/kbd/console.go;
//     it must not be used elsewhere.
//     Pre: The mouse device is /dev/input/mice.
//          This device is readable for the world, and a mouse is working on it.

type
  Command = byte; const (
  None = byte(iota)
  Go                // mouse move without any button pressed
  Here; This; That  // left, right, middle button pressed
  Drag; Drop; Move  // move with left, right, middle button
  To; There; Hither // left, right, middle button released
  nCommands
)
var
  Pipe chan Command

// The range of the mouse is the rectangle with
// the upper left corner (x, y) and the lower right corner (x + w - 1, y + h - 1).
// Initially, (x, y) = (0, 0) and (w, h) = (1600, 1200).
func Def (x, y, w, h uint) { def(x,y,w,h) }

// Pre: x, y is inside the range of the mouse.
// The pixel-position of the mouse is (x, y).
func Warp (x, y uint) { warp(x,y) }

// Returns the pixel-position (x, y) of the mouse,
// where (x, y) = (0, 0) is the top left corner.
func Pos() (int, int) { return pos() }
