package species

// (c) Christian Maurer   v. 230311 - license see µU.go

import (
  "µU/col"
  . "µU/obj"
  "µU/scr"
)
const ( // Eco
  plant = byte(iota)
  hare
  fox
  nSpec // > nSpec of Life
)
const ( // Life
  nothing = byte(iota)
  cell
)
const (
  size = 16
  point = "."
)
type (
  image [size]string
  species struct {
                 byte // < nSpec
                 }
  representation struct {
                        image
                        col.Colour
                        }
)
var (
  system  System
  rep     [nSpec]representation
  compare [nSpec]*species
)

func init() {
  sys (Life)
  for n := byte(0); n < nSpec; n++ {
    compare[n] = New().(*species)
    compare[n].byte = n
  }
}

func sys (s System) {
  system = s
  var p image
  switch system {
  case Eco:
    rep[plant].Colour = col.Green()
    for i := 0; i < size; i++ {
      rep[plant].image[i] = "****************"
    }
    rep[hare].Colour = col.DarkYellow()
    p = image {"                ",
               "     *     *    ",
               "     **   **    ",
               "     *** ***    ",
               "      *****     ",
               "      *****     ",
               "      *****     ",
               "      *****     ",
               "     *******    ",
               "    *********   ",
               "   ***********  ",
               "  ** ******* ** ",
               "     *** ***    ",
               "     **   **    ",
               "    ***   ***   ",
               "                "}
    for i := 0; i < size; i++ {
      rep[hare].image[i] = p[i]
    }
    rep[fox].Colour = col.Brown()
    p = image {"                ",
               "     * *        ",
               "     **         ",
               " ******         ",
               "*******         ",
               "   ***          ",
               "   ********     ",
               "  **********    ",
               "  ***********   ",
               "  ********** *  ",
               "  ***    ***  * ",
               "  * *    *  *  *",
               "  * *    *  *   ",
               "  * *    *  *   ",
               "  * *    *  *   ",
               "                "}
    for i := 0; i < size; i++ {
      rep[fox].image[i] = p[i]
    }
    NNeighbours = 4
    Suffix = "eco"
  case Life:
    rep[nothing].Colour = col.Red()
    rep[cell].Colour = rep[nothing].Colour
    NNeighbours = 8
    Suffix = "life"
  }
}

func new_() Species {
  return new(species)
}

func (x *species) imp (Y any) *species {
  y, ok := Y.(*species)
  if ! ok { TypeNotEqPanic(x, Y) }
  return y
}

func (x *species) Eq (Y any) bool {
  return x.byte == x.imp(Y).byte
}

func (x *species) Copy (Y any) {
  x.byte = x.imp(Y).byte
}

func (x *species) Clone() any {
  y := New()
  y.Copy (x)
  return y
}

func (x *species) Set (k uint) {
  if k <= 1 {
    x.byte = byte(k)
  } else {
    if system == Life {
      x.byte = 1
    } else {
      x.byte = 2
    }
  }
}

func (X *species) Write (l, c uint) {
  r := rep[X.byte]
  switch system {
  case Life:
    x, y := 8 * int(c) + 8, 16 * int(l) + 8
    if X.byte == nothing {
      scr.ColourF (col.LightWhite())
      scr.CircleFull (x, y, 6)
      scr.ColourF (r.Colour)
      scr.CircleFull (x, y, 1)
    } else {
      scr.ColourF (r.Colour)
      scr.CircleFull (x, y, 6)
    }
  case Eco:
    for y := 0; y < size; y++ {
      for x := 0; x < size; x++ {
        f := r.Colour
        b := r.Colour
        if r.image[y][x] == ' ' {
          f = b
        }
        scr.Colours (f, b)
        scr.Point (size * int(c) / 2 + x, size * int(l) + y)
      }
    }
  }
}

func (x *species) Modify (numberOf func (Species) uint) {
  switch system {
  case Life:
    c := numberOf (compare[cell])
    switch x.byte {
    case nothing:
      if c == 3 {
        x.byte = cell
      }
    case cell:
      if c < 2 || c > 3 {
        x.byte = nothing
      }
    }
  case Eco:
    h, f := numberOf (compare[hare]), numberOf (compare[fox])
    switch x.byte {
    case plant:
      if h > 0 && h < 4 {
        x.byte = hare
      }
    case hare:
      if h == 4 {
        x.byte = plant
      }
      if f > 0 {
        x.byte = fox
      }
    case fox:
      if h == 0 {
        x.byte = plant
      }
    }
  }
}

func (x *species) String() string {
  switch system {
  case Life:
    if x.byte == cell {
      return "o"
    }
  case Eco:
    switch x.byte {
    case fox:
      return "f"
    case hare:
      return "h"
    }
  }
  return point
}

func (x *species) Defined (s string) bool {
  switch system {
  case Life:
    switch s {
    case "o":
      x.byte = cell
    case point:
      x.byte = nothing
    default:
      return false
    }
  case Eco:
    switch s {
    case "f":
      x.byte = fox
    case "h":
      x.byte = hare
    case point:
      x.byte = plant
    default:
      return false
    }
  }
  return true
}
