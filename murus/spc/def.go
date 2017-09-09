package spc

// (c) Christian Maurer   v. 121204 - license see murus.go

type
  Direction byte
// the unit vectors of an orthonormal rightoriented coordinate-system in 3-space
const (
  Right = Direction(iota)
  Front
  Top
  NDirs
)
const (
  D0 = Right
)
type (
  Coord [NDirs]float64
  GridCoord [NDirs]int16
)
// Pre: The values of Origin and Unit must not be changed.
// Origin = (0, 0, 0), Unit[right] = (1, 0, 0), Unit[front] = (0, 1, 0), Unit[top] = (0, 0, 1).
var (
  Origin Coord
  Unit [NDirs]Coord
)

// Returns Front/Top/Right for d = Right/Front/Top.
func Next (d Direction) Direction { return next(d) }

// Returns Top/Right/Front for d = Right/Front/Top (vectorproduct of d and Next).
func Prev (d Direction) Direction { return prev(d) }
