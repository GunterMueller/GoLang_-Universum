package ppm

// (c) Christian Maurer   v. 220120 - license see µU.go

import
  "µU/col"

type
  Image interface {

// Pre: A file with the name n.ppm containing an image
//      in the ppm-format exists in the actual directory.
// x is the image saved in that file.
  Load (n string)

// Pre: x was generated by a call of Get.
// Returns width and height of x.
  Size() (uint, uint)

  Colours() [][]col.Colour

  Store (n string)

  Print (x, y int)
}

func New() Image { return new_() }
