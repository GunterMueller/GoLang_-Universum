package mouse

// (c) Christian Maurer   v. 210311 - license see µU.go

// >>> This package only serves the implementations of µU/kbd 
//     and µU/cons; it must not be used elsewhere.

type
  Command byte

// Returns true, iff a mouse exists.
func Ex() bool { return ex() }

// TODO Spec
func Channel() chan Command { return channel() }

// TODO Spec
func Def (x, y, w, h uint) { def(x,y,w,h) }

// Pre: TODO
// The pixel-position of the mouse is (x, y).
func Warp (x, y uint) { warp(x,y) }

// Returns the pixel-position (x, y) of the mouse,
// where (x, y) = (0, 0) is the top left corner.
func Pos() (int, int) { return pos() }
