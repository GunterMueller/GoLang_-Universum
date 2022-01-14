package psp

// (c) Christian Maurer   v. 220112 - license see µU.go

import (
  "µU/col"
  "µU/font"
)
type
  PostscriptPage interface { // all float64-parameters in pt

  X (x int) float64
  Y (y int) float64

  Name (n string)

  Fin()

// Default unit 1 pt is replaced by u (measured in pt).
  SetUnit (u float64)

  Translate (l, b float64)

  SetColour (c col.Colour)

//  SetFont (f font.Font)
  SetFontsize (s font.Size)

  Write (s string, x, y float64)

  SetLinewidth (w float64)

  Point (x, y float64)

  Points (x, y []float64)

  Line (x, y, x1, y1 float64)

  Lines (x, y, x1, y1 []float64)

  Segments (x, y []float64)

  Rectangle (x, y, w, h float64, f bool)

  Polygon (x, y []float64, f bool)

  Circle (x, y, r float64, f bool)

  Ellipse (x, y, a, b float64, f bool)

  Curve (x, y []float64)
}

// Returns a new Postscript page.
func New() PostscriptPage { return new_() }
