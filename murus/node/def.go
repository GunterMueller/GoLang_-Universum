package node

// (c) murus.org  v. 161216 - license see murus.go

import (
  . "murus/obj"
  "murus/col"
)

// Returns a new empty node for content of the type of e
// of size (width, height) = (w, h) (in colums, lines).
func New (e EditorGr, w, h uint) Node { return new_(e,w,h) }

type
// Nodes with objects as content and positions on the screen.
  Node interface {

  Object // Empty means content is empty

// Returns the content of x.
  Content() EditorGr

// Returns the width and height of x.
  Size() (uint, uint)

// x has the the center (x, y).
  Set (x, y int)

// Returns the coordinates of x.
  Pos() (int, int)

// TODO Spec
  Contour() (uint, uint)

// The center of x has the position of the mouse.
  Mouse()

// Returns true, iff the mouse has the position of the center of x.
  UnderMouse() bool

// f and b are the normal colours of x.
  Colours (f, b col.Colour)

// f and b are the actual colours of x.
  ColoursA (f, b col.Colour)

// x is written to the screen at its position in its normal colour.
  Write()

// x is written to the screen at its position;
// if a, in the actual colour, otherwise in its normal colour.
  Write1 (a bool)

// x and N are connected by a line in the colour corresponding to a.
  Write3 (e Any, N Node, a bool)
  Write3dir (e Any, d bool, N Node, a bool)

// x has the content, that was edited by the user.
  Edit()
}

func O (n Any, a bool) {
  n.(Node).Write1(a)
}

func O3 (n, e, n1 Any, a bool) {
  n.(Node).Write3 (e, n1.(Node), a)
}

func O3dir (n, e Any, d bool, n1 Any, a bool) {
  n.(Node).Write3dir (e, d, n1.(Node), a)
}
