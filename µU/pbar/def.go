package pbar

// (c) Christian Maurer   v. 170919 - license see µU.go

import
  "µU/col"
type
  ProgressBar interface {

  col.Colourer

// Pre: x + w <= scr.Wd, y + h <= scr.Ht.
// x has the top left corner (x, y), width w and height h.
  Locate (x, y, w, h uint)

// X has capacity c and fill degree 0.
  SetCap (c uint)

// Pre: i <= capacity of X.
// x has fill degree i, i.e. X is filled up to i/k % (c = capacity of B).
  Fill (i uint)

// Returns the fill degree of X.
  Filldegree () uint

// x is written to the screen as rectangle with its top left corner
// and its width and height, the fraction of the capacity of B, that
// corresponds to the fill degree of B (with horizontal fill direction
// the left, otherwise the bottom part of the rectangle)
// in its foregroundcolour, the rest in its backgroundcolour.
  Write()

// TODO Spec
  Edit (i *uint)
}

// Returns for h == true/false a new progress bar with horizontal
// resp. vertical fill direction, capacity 100 and fill degree 0
func New(h bool) ProgressBar { return new_(h) }
