package spc

// (c) Christian Maurer   v. 201030 - license see µU.go

import (
  "µU/vect"
  "µU/gl"
)
const
  epsilon = 1.0E-6
var (
  origin, focus = vect.New(), vect.New()
  trihedron [3]vect.Vector // right, front, top = trihedron[0], trihedron[1], trihedron[2]
  temp = vect.New()
)

func init() {
  trihedron = [3]vect.Vector {vect.New(), vect.New(), vect.New()}
}

func set (ox, oy, oz, fx, fy, fz, tx, ty, tz float64) {
  origin.Set3 (ox, oy, oz)
  focus.Set3 (fx, fy, fz)
  trihedron[2].Set3 (tx, ty, tz)
  trihedron[1].Copy (focus); trihedron[1].Sub (origin)
  trihedron[0].Ext (trihedron[1], trihedron[2])
  for i := 0; i < 3; i++ { trihedron[i].Norm() }
}

func set3 (ox, oy, oz float64) {
  origin.Set3 (ox, oy, oz)
}

func get() (float64, float64, float64, float64, float64, float64, float64, float64, float64) {
  ox, oy, oz := origin.Coord3()
  fx, fy, fz := focus.Coord3()
  tx, ty, tz := trihedron[2].Coord3()
  return ox, oy, oz, fx, fy, fz, tx, ty, tz
}

func adjustFocus() {
  delta := origin.Distance (focus)
  focus.Scale (delta, trihedron[1])
  focus.Add (origin)
}

func move (i uint, d float64) {
  temp.Scale (d, trihedron[i])
  origin.Add (temp) // origin = origin + d * trihedron[i]
  adjustFocus()
}

func rotate (i uint, alpha float64) {
  n, p := (i + 1) % 3, (i + 2) % 3
  trihedron[n].Rot (trihedron[i], alpha)
  trihedron[n].Norm()
  trihedron[p].Ext (trihedron[i], trihedron[n])
  trihedron[p].Norm()
}

func turn (i uint, alpha float64) {
  rotate (i, alpha)
  adjustFocus()
}

func turnAroundFocus (i uint, alpha float64) {
  delta := origin.Distance (focus)
  if delta < epsilon { return }
  rotate (i, -alpha)
  temp.Scale (delta, trihedron[1])
  origin.Diff (focus, temp)
}

func setLight (n uint) {
  x, y, z := origin.Coord3()
  gl.PosLight (n, x, y, z)
}
