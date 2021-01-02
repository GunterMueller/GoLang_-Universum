package ppm

// (c) Christian Maurer   v. 201230 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  PPM interface { // portable pixmap files

  Persistor

  Set (w, h uint)

  Get (n string)

  Put()

// x is written to the screen.
  Write()

// c is the actual colour.
  ColourF (c col.Colour)

// Pre: 0 <= x < width, 0 <= y < height of the calling ppm-file.
// The point at screen position (x, y) is set in the actual colour in the calling ppm-file.
  Point (x, y int)
}

// Returns a new ppm-file with width w and height h.
// All its points are black.
func New() PPM { return new_() }
