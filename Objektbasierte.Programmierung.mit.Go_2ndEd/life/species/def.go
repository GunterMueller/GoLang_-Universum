package species

// (c) Christian Maurer   v. 230311 - license see µU.go

import
  . "µU/obj"
type
  System byte
const (
  Eco = System(iota) // Ecosystem with foxes, hares and plants
  Life               // Game of Life (John Conway)
)
var (
  Suffix string
  NNeighbours uint
)
type
  Species interface {

  Equaler
  Stringer

  Write (l, c uint)

// if k == 0 in Eco:  x is a plant
//           in Life: x is nothing
// if k == 2 in Eco:  x is a hare
//           in Life: x is a cell
// if k == 3 in Eco:  x is a fox
//           in Life: x is a cell
  Set (k uint)

// The actual species has changed according to func.
  Modify (func (Species) uint)
}

// Returns a new species.
func New() Species { return new_() }

// The actual system is s.
func Sys (s System) { sys(s) }
