package spc

// (c) Christian Maurer   v. 191019 - license see µU.go

import (
//  "fmt"
//  "math"
//  "µU/ker"
  "µU/vect"
  "µU/gl" // PosLight
)
const
  epsilon = 1.0E-6
var (
  eye, focus, temp vect.Vector
  system [3]vect.Vector
  delta float64 // invariant: delta == eye.Distance (focus)
)

func init() {
  eye = vect.New()
  focus = vect.New()
  system = [3]vect.Vector { vect.New3 (1,0,0), vect.New3 (0,1,0), vect.New3 (0,0,1) }
  temp = vect.New()
  delta = eye.Distance (focus)
}

func set (ex, ey, ez, fx, fy, fz, nx, ny, nz float64) {
  eye.Set3 (ex, ey, ez)
  focus.Set3 (fx, fy, fz)
  system[2].Set3 (nx, ny, nz); system[2].Norm()
  system[1].Copy (focus); system[1].Sub (eye); system[1].Norm()
  system[0].Ext (system[1], system[2]); system[0].Norm()
// print ("right: "); fmt.Println (system[0].Coord3())
// print ("front: "); fmt.Println (system[1].Coord3())
// print ("normal: "); fmt.Println (system[1].Coord3())
}

func get() (float64, float64, float64, float64, float64, float64, float64, float64, float64) {
  ex, ey, ez := eye.Coord3()
  fx, fy, fz := focus.Coord3()
  nx, ny, nz := system[2].Coord3()
  return ex, ey, ez, fx, fy, fz, nx, ny, nz
}

func adjustFocus() {
  delta = eye.Distance (focus)
// print (" vor adj focus: "); fmt.Println (focus.Coord3())
  focus.Scale (delta, system[1])
  focus.Add (eye)
// print ("nach adj focus: "); fmt.Println (focus.Coord3())
}

/*
func distance() float64 {
  delta = eye.Distance (focus)
  if math.Abs (eye.Distance (focus) - delta) > epsilon {
    adjustFocus()
  }
  return delta
}
*/

func move (i int, dist float64) {
//  system[2].Ext (system[0], system[1])
//  system[2].Norm()
  temp.Scale (dist, system[i])
  eye.Add (temp)
  adjustFocus()
}

func rotate (i int, alpha float64) {
// print ("vor rot right:  "); fmt.Println (system[0].Coord3())
// print ("vor rot front:  "); fmt.Println (system[1].Coord3())
// print ("vor rot normal: "); fmt.Println (system[2].Coord3())
  n , p := (i + 1) % 3, (i + 2) % 3
  system[n].Rot (system[i], alpha)
  system[n].Norm()
  system[p].Ext (system[i], system[n])
  system[p].Norm()
// print ("nach rot right:  "); fmt.Println (system[0].Coord3())
// print ("nach rot front:  "); fmt.Println (system[1].Coord3())
// print ("nach rot normal: "); fmt.Println (system[2].Coord3())
}

func turn (i int, alpha float64) {
  rotate (i, alpha)
// print ("turn: "); fmt.Println (focus.Coord3())
  adjustFocus()
// print ("turn1: "); fmt.Println (focus.Coord3())
}

func invert() { // TODO
  system[0].Dilate (-1.)
  system[1].Dilate (-1.)
  adjustFocus()
}

func adjustEye() {
//  if delta != eye.Distance (focus) { ker.Oops() }
  delta := eye.Distance (focus)
  temp.Scale (delta, system[1])
  eye.Diff (focus, temp)
}

/*
func foc (d float64) {
  if d < epsilon { return }
  delta = d
  adjustEye()
}
*/

func turnAroundFocus (i int, alpha float64) {
  delta = eye.Distance (focus)
  if delta < epsilon { return }
//  println ("TurnAroundFocus")
  rotate (i, -alpha)
//  println ("rotated")
// Dieser Vorzeichenwechsel ist ein Arbeitsdrumrum
// um einen mir bisher nicht erklärbaren Fehler.
// Vermutlich liegt das daran, dass ich irgendeine suboptimal
// dokumentierte Eigenschaft von openGL noch nicht begriffen habe.
  adjustEye()
}

func setLight (n uint) {
  gl.PosLight (n, eye)
}
