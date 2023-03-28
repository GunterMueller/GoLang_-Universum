package world

// (c) Christian Maurer   v. 230311 - license see µU.go

import (
  "life/species"
  . "µU/obj"
  "µU/mode"
)
const
  Len = 8 // maximal length of the name of the world
type
  World interface {

  Equaler
  Write()
  Edit()
  Stringer
  Persistor
}

// Returns a new empty world.
func New() World { return new_() }

// Returns the mode for life.
func Mode() mode.Mode { return m() }

// s is the actual system.
func Sys (s species.System) { sys(s) }
