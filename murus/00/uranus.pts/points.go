package pts

// (c) murus.org  v. 170810 - license see murus.go

import (
  "murus/ker"
  "murus/env"
  "murus/gl"
  "murus/col"
  "murus/scr"
  "murus/errh"
  "murus/vect"
  "murus/files"
  "murus/sel"
  "murus/pseq"
  "murus/pt"
)
const
  suffix = "scn"
type
  points struct {
                pseq.PersistentSequence
                uint "number of points"
     eye, focus,
         vv, nn []vect.Vector
          point []pt.Point
//      started bool
                }

func init() {
  gl.Cls (col.LightWhite)
  files.Cd (env.Home() + ker.Separator + ker.DotMurus + ker.Separator + suffix)
}

func new_() Points {
  x := new (points)
  x.PersistentSequence = pseq.New (pt.New())
  x.eye, x.focus = []vect.Vector { vect.New() }, []vect.Vector { vect.New() }
  return x
}

func (x *points) Clr() {
  x.PersistentSequence.Clr()
  x.uint = 0
//  x.eye[0].Clr()
//  x.focus[0].Clr()
//  x.vv, x.nn = nil, nil
//  x.point = nil
////  x.started = false
}

func (x *points) Empty() bool {
  return x.uint == 0
}

func (x *points) Name (s string) {
  x.PersistentSequence.Name (s + "." + suffix)
  x.eye[0].Set3 (0, -1, 0)
  x.focus[0].Clr()
  x.uint = x.PersistentSequence.Num()
  if x.uint > 0 {
    x.vv, x.nn = make ([]vect.Vector, x.uint), make ([]vect.Vector, x.uint)
    x.point = make ([]pt.Point, x.uint)
    for i := uint(0); i < x.uint; i++ {
      x.vv[i], x.nn[i] = vect.New(), vect.New()
      x.PersistentSequence.Seek (i)
      x.point[i] = x.PersistentSequence.Get().(pt.Point)
    }
    if x.point[x.uint - 1].Class() == pt.Start {
      x.eye[0], x.focus[0] = x.point[x.uint - 1].Read2()
    } else {
      x.eye[0].Set3 (0, 0, 1)
      x.focus[0].Clr()
    }
  }
}

func (x *points) Rename (s string) {
  x.PersistentSequence.Rename (s + "." + suffix)
}

func (x *points) NameCall() {
  x.Name (env.Par (1))
}

func (x *points) Select() {
  name, _ := sel.Names ("Szene:", suffix, 64, 0, 0, scr.ScrColF(), scr.ScrColB())
  if name == "" {
    errh.Error0("nicht vorhanden")
    x.Clr()
  } else {
    x.Name (name)
  }
}

func (x *points) Ins1 (c pt.Class, v []vect.Vector, f col.Colour) {
//  if started { ker.Oops() }
  if c > pt.Polygon { ker.Oops() }
  p := pt.New()
  n := vect.New3 (0, 0, 1)
  a := uint(len (v))
  for i := uint(0); i < a; i++ {
    p.Set (c, a - 1 - i, f, v[i], n)
    x.PersistentSequence.Ins (p)
  }
}

func (x *points) InsLight (l uint, v, n []vect.Vector, f col.Colour) {
//  if started { ker.Oops() }
  p := pt.New()
  p.Set (pt.Light, l, f, v[0], n[0])
  x.PersistentSequence.Ins (p)
}

func (x *points) Ins (c pt.Class, v, n []vect.Vector, f col.Colour) {
//  if started { ker.Oops() }
  if c == pt.Light { ker.Panic ("pts Ins vs. InsLight") }
  a := uint(len (v))
  if uint(len (n)) != a { println ("pts.Ins: len(n) ==", len(n), " != len(v) ==", len(v)) }
  p := pt.New()
  for i := uint(0); i < a; i++ {
    p.Set (c, a - 1 - i, f, v[i], n[i])
    x.PersistentSequence.Ins (p)
  }
}

func (x *points) Start (x0, y0, z0, x1, y1, z1 float64) {
  if x0 == x1 && y0 == y1 && z0 == z1 { ker.Oops() }
  x.eye[0].Set3 (x0, y0, z0)
  x.focus[0].Set3 (x1, y1, z1)
  x.Ins (pt.Start, x.eye, x.focus, col.Red)
// println ("murus/pts/Start done")
//  x.started = true
}

func (x *points) StartCoord() (float64, float64, float64, float64, float64, float64) {
  const unclear = 500.0
  gl.Init (unclear * x.eye[0].Distance (x.focus[0]))
  x0, y0, z0 := x.eye[0].Coord3()
  x1, y1, z1 := x.focus[0].Coord3()
  return x0, y0, z0, x1, y1, z1
}

var öks uint

func (x *points) Write() {
// TODO: point of Class Start at first 
  i := uint (0)
//  ker.Mess0()
  gl.Write0()
//  ker.Mess ("0)")
  j := uint(0)
  var p pt.Point
  for j + 1 < x.uint {
    i = uint(0)
    var a uint
    for {
      k := x.point[j].Number()
      if i == 0 {
        if x.point[j].Class() == pt.Light {
          a = k
          k = 0
        } else {
          a = k + 1 // !
        }
      }
      x.vv[i], x.nn[i] = x.point[j].Read2()
      i ++
      p = x.point[j]
      j ++
      if k == 0 { break }
    }
    if p != x.point[j-1] { errh.Error ("pt.Write: strange stuff", j) }
    var f gl.Figure
    switch p.Class() {
    case pt.None:
//    switch point[j].Class() { case pt.None:
      f = gl.UNDEF
    case pt.Start:
      return
    case pt.Light:
      f = gl.LIGHT
    case pt.Points:
      f = gl.POINTS
    case pt.Lines:
      f = gl.LINES
    case pt.LineStrip:
      f = gl.LINE_STRIP
    case pt.LineLoop:
      f = gl.LINE_LOOP
    case pt.Triangles:
      f = gl.TRIANGLES
    case pt.TriangleStrip:
      f = gl.TRIANGLE_STRIP
    case pt.TriangleFan:
      f = gl.TRIANGLE_FAN
    case pt.Quads:
      f = gl.QUADS
    case pt.QuadStrip:
      f = gl.QUAD_STRIP
    case pt.Polygon:
      f = gl.POLYGON
    }
    gl.Write (f, a, x.vv, x.nn, p.Colour())
//    gl.Write (f, a, x.vv, x.nn, x.point[j].Colour())
  }
// öks++; println ("öks =", öks)
//  ker.Mess ("-)")
  gl.Write1()
//println ("gl.Written")
//  ker.Mess ("1)")
}

func (x *points) Fin() {
  x.PersistentSequence.Fin()
}
