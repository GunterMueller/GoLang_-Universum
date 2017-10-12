package figs

// (c) Christian Maurer   v, 170917 - license see µu.go

import (
  "math"
  "µu/col"
  "µu/fig"
)
const (
  pt = iota; ax; ls; ll; tr; ts; tf; qu; qs; po;
       pa; cu; cb; cc; c1; pr; pc; pp; py; oh; oc; mu; mc;
       ci; cS; vc; co; cy; hc; sp; to; ht; su; ss
)
var
  cs = []col.Colour {col.Red, col.Orange, col.Yellow, col.Green, col.Cyan,
                     col.Blue, col.Magenta, col.Gray, col.White, col.LightWhite}

func draw() {
  var figure int
  figure = cu
  figure = oc
  fig.Clear()
  switch figure {
  case pt:
//    fig.Point (col.Yellow, )
  case ax:
  const r = 3.0
    fig.Line (col.Red,     0, 0, 0, r, 0, 0)
    fig.Line (col.Orange, -r, 0, 0, 0, 0, 0)
    fig.Line (col.Green,   0, 0, 0, 0, r, 0)
    fig.Line (col.Yellow,  0,-r, 0, 0, 0, 0)
    fig.Line (col.Blue,    0, 0, 0, 0, 0, r)
    fig.Line (col.Magenta, 0, 0,-r, 0, 0, 0)
  case ls:
    fig.Linestrip (col.Yellow, 0, 0, 0, 1, 0, 0, 2, 0, 0, 2, 1, 0, 0, 3, 0)
  case ll:
    fig.Lineloop (col.Yellow, 0, 0, 0, 1, 0, 0, 2, 0, 0, 2, 1, 0, 0, 3, 0)
  case tr:
    fig.Triangle (col.Yellow, 0, 0, 0, 2, 0, 0, 0, 2, 0)
  case ts:
    fig.Trianglestrip (col.Cyan, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 2, 0, 0, 2, 1, 0)
  case tf:
    fig.Trianglefan (col.Cyan, 0, 0, 0, 3, 0, 0, 3, 1, 0, 1, 3, 0, 0, 3, 0,-1, 2, 0)
  case qu:
    fig.Quad (col.White, 1, 1, 0, -1, 1, 0, -1, -1, 0, 1, -1, 0)
  case qs:
    fig.Quadstrip (col.Pink, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 2, 0, 0, 2, 1, 0)
  case po:
    fig.Polygon (col.Cyan, 3, 0, 0, 2, 2, 0, 0, 2, 0, -2, 1, 0,-1,-2, 0, 2,-2, 0)
  case pa:
    fig.Parallelogram (col.Blue, 0, 1, 1, 2, 0, 0, 1, 1, 0)
  case cu:
    fig.CubeC (cs, 0, 0, 0, 2)
/*
    a := 6
    for x := -a; x <= a; x += 2 {
      for y := -a; y <= a; y += 2 {
        for z := -a; z <= a; z += 2 {
          fig.CubeC (cs, float64(x), float64(y), float64(z), 1)
        }
      }
    }
*/
  case cb:
    fig.Cuboid (col.Green, 0, 0, 0, 0.5, 1, 1.5)
  case cc:
    fig.CuboidC (cs, 0, 0, 0, 0.5, 1, 1.5)
  case c1:
    fig.Cuboid1 (cs[0], 0, 0, 0, 2, 2, 2, 1)
  case pr:
    fig.Prism (cs[2], 1, 1, 2, 2, 0, 0, 0, 2, 0, -1,-2, 0, 0, 0, 0, 1,-1, 0)
  case pc:
    fig.PrismC (cs, 1, 1, 2, 3, 0, 0, 2, 2, 0, 0, 2, 0,-1, 0, 0, 1,-1, 0)
  case pp:
    fig.Parallelepiped (col.Orange, 0, 0, 0, -2, 1, 1, 1, 1, 1, 1, -2, -1)
  case py:
//    fig.PyramidC (cs, )
  case oh:
    fig.Octahedron (col.White, 0, 0, 0, 2)
  case oc:
    fig.OctahedronC (cs, 0, 0, 0, 2)
  case mu:
    fig.Multipyramid (col.Cyan, 1, 1, 2, 2, 0, 0, 0, 2, 0, -1,-2, 0, 0, 0, 0, 1,-1, 0)
  case mc:
    fig.MultipyramidC (cs, 1, 1, 2, 2, 0, 0, 0, 2, 0, -1,-2, 0, 0, 0, 0, 1,-1, 0)
  case ci:
    fig.Circle (col.Yellow, -1, +1, 0, 2)
  case cS:
    fig.CircleSegment (col.Yellow, 3, 2, 0, 2, 0, 90)
  case vc:
    fig.VertCircle (col.Cyan, 0, 0, 0, 2, 45)
  case co:
    fig.DoubleCone (col.Pink, 0, 0, 0, 1, 2)
  case cy:
    fig.Cylinder (col.Green, -2, 0, 0, 0.5, 2)
  case hc:
    fig.HorCylinder (col.Red, 1, 1, 0, 0.5, 1, 45)
  case sp:
    fig.Sphere (col.Gray,   -4, 0, 0, 1)
    fig.Sphere (col.Red,     4, 0, 4, 1)
    fig.Sphere (col.Green,   4, 0, 0, 1)
    fig.Sphere (col.Blue,   -4, 0, 0, 1)
    fig.Sphere (col.Yellow,  0, 4, 0, 1)
    fig.Sphere (col.Cyan,    0, 0, 4, 1)
    fig.Sphere (col.Magenta, 0,-4,-4, 1)
    fig.Sphere (col.Orange,  0,-4, 4, 1)
  case to:
    fig.Torus (col.Green, 0, 0, 0, 5, 1)
  case ht:
    fig.HorTorus (col.Blue, 0, 0, 0, 5, 1, 2)
  case su:
    r := 10.0
    fig.Surface (col.Yellow, func (x, y float64) float64 { return math.Sin(x) * math.Sin(y) },
             -r, -r, -r, r, r, r, 960, 720, 8)
  case ss:
    const r = 4.0
    w, h := uint(960), uint(720)
    fig.Surface (col.Red, func (x, y float64) float64 { return 0.3 * (x * x - y * y) },
                -r, -r, -r, r, r, r, w, h, 8)
    fig.Surface (col.Green, func (x, y float64) float64 { return 3 * math.Exp (-x * x - y * y) },
                -r, -r, -r, r, r, r, w, h, 8)
    fig.Surface (col.Blue, func (x, y float64) float64 { return 0.15 * (x * x + y * y) - 4 },
                -r, -r, -r, r, r, r, w, h, 8)
  }
}

func start() (float64, float64, float64, float64, float64, float64) {
  return 0, 0, 0, 0, 0, 0
  return -5, -10, 1, 500, 0, 0 // surface
  return 0, 0, -10, 0, 0, 0 // surface
  return 2, 1, -5, 0, 0, 0
  return 0, 0, 0, 0,-5,-3 // surface
  return 0,-3, 3, 0, 0, 0
  return 0,-3, 0, 0, 0, 0
  return 0, 0, 0, 0, 0, 0
}
