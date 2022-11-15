package fig3

// (c) Christian Maurer  v. 221113 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/gl"
)

func init() {
  gl.Clear()
}

func point (c col.Colour, x ...float64) {
  gl.Colour (c)
  gl.Point (x...)
}

func line (c col.Colour, x ...float64) {
  gl.Colour (c)
  gl.Line (x...)
}

func lineloop (c col.Colour, x ...float64) {
  gl.Colour (c)
  gl.Lineloop (x...)
}

func linestrip (c col.Colour, x ...float64) {
  gl.Colour (c)
  gl.Linestrip (x...)
}

func triangle (c col.Colour, x ...float64) {
  gl.Colour (c)
  gl.Triangle (x...)
}

func trianglestrip (c col.Colour, x ...float64) {
  gl.Colour (c)
  gl.Trianglestrip (x...)
}

func trianglefan (c col.Colour, x ...float64) {
  gl.Colour (c)
  gl.Trianglefan (x...)
}

func quad (c col.Colour, x ...float64) {
  gl.Colour (c)
  gl.Quad (x...)
}

func quadstrip (c col.Colour, x ...float64) {
  gl.Colour (c)
  gl.Quadstrip (x...)
}

func polygon (c col.Colour, x ...float64) {
  gl.Colour (c)
  gl.Polygon (x...)
}

func horRectangle (c col.Colour, x, y, z, x1, y1 float64, up bool) {
  gl.Colour (c)
  gl.HorRectangle (x, y, z, x1, y1, up)
}

func vertRectangle (c col.Colour, x ...float64) {
  gl.Colour (c)
  gl.VertRectangle (x...)
}

func parallelogram (c col.Colour, x ...float64) {
  gl.Colour (c)
  gl.Parallelogram (x...)
}

func curve (c col.Colour, f Ft2xyz, t0, t1 float64) {
  gl.Colour (c)
  gl.Curve (f, t0, t1)
}

func plane (f col.Colour, a, b, c, wx, wy float64) {
  gl.Colour (f)
  gl.Plane (a, b, c, wx, wy)
}

func cube (c col.Colour, x, y, z, a float64) {
  gl.Colour (c)
  gl.Cube (x, y, z, a)
}

func cubeC (c []col.Colour, x, y, z, a float64) {
  gl.CubeC (c, x, y, z, a)
}

func cuboid (c col.Colour, x, y, z, x1, y1, z1 float64) {
  gl.Colour (c)
  gl.Cuboid (x, y, z, x1, y1, z1)
}

func cuboid1 (c col.Colour, x, y, z, dx, dy, dz, a float64) {
  gl.Colour (c)
  gl.Cuboid1 (x, y, z, dx, dy, dz, a)
}

func cuboidC (c []col.Colour, x, y, z, dx, dy, dz float64) {
  gl.CuboidC (c, x, y, z, dx, dy, dz)
}

func prism (c col.Colour, x ...float64) {
  gl.Colour (c)
  gl.Prism (x...)
}

func prismC (c []col.Colour, x ...float64) {
  gl.PrismC (c, x...)
}

func parallelepiped (c col.Colour, x ...float64) {
  gl.Colour (c)
  gl.Parallelepiped (x...)
}

func parallelepipedC (c []col.Colour, x ...float64) {
  gl.ParallelepipedC (c, x...)
}

func pyramid (c col.Colour, x, y, z, a, h float64) {
  gl.Colour (c)
  gl.Pyramid (x, y, z, a, h)
}

func pyramidC (c []col.Colour, x, y, z, a, h float64) {
  gl.PyramidC (c, x, y, z, a, h)
}

func multipyramid (f col.Colour, x, y, z, h float64, c ...float64) {
  gl.Colour (f)
  gl.Multipyramid (x, y, z, h, c...)
}

func multipyramidC (f []col.Colour, x, y, z, h float64, c ...float64) {
  gl.MultipyramidC (f, x, y, z, h, c...)
}

func octopus (c col.Colour, x ...float64) {
  gl.Colour (c)
  gl.Octopus (x...)
}

func octopusC (c []col.Colour, x ...float64) {
  gl.OctopusC (c, x...)
}

func octahedron (c col.Colour, x, y, z, r float64) {
  gl.Colour (c)
  gl.Octahedron (x, y, z, r)
}

func octahedronC (c []col.Colour, x, y, z, r float64) {
  gl.OctahedronC (c, x, y, z, r)
}

func circle (c col.Colour, x, y, z, r float64) {
  gl.Colour (c)
  gl.Circle (x, y, z, r)
}

func circleSegment (c col.Colour, x, y, z, r, a, b float64) {
  gl.Colour (c)
  gl.CircleSegment (x, y, z, r, a, b)
}

func vertCircle (c col.Colour, x, y, z, r, a float64) {
  gl.Colour (c)
  gl.VertCircle (x, y, z, r, a)
}

func sphere (c col.Colour, x, y, z, r float64) {
  gl.Colour (c)
  gl.Sphere (x, y, z, r)
}

func cone (c col.Colour, x, y, z, r, h float64) {
  gl.Colour (c)
  gl.Cone (x, y, z, r, h)
}

func doubleCone (c col.Colour, x, y, z, r, h float64) {
  gl.Colour (c)
  gl.DoubleCone (x, y, z, r, h)
}

func cylinder (c col.Colour, x, y, z, r, h float64) {
  gl.Colour (c)
  gl.Cylinder (x, y, z, r, h)
}

func cylinderC (c []col.Colour, x, y, z, r, h float64) {
  gl.CylinderC (c, x, y, z, r, h)
}

func cylinderSegment (c col.Colour, x, y, z, r, h, a, b float64) {
  gl.Colour (c)
  gl.CylinderSegment (x, y, z, r, h, a, b)
}

func horCylinder (c col.Colour, x, y, z, r, l, a float64) {
  gl.Colour (c)
  gl.HorCylinder (x, y, z, r, l, a)
}

func horCylinderC (c []col.Colour, x, y, z, r, l, a float64) {
  gl.HorCylinderC (c, x, y, z, r, l, a)
}

func torus (c col.Colour, x, y, z, R, r float64) {
  gl.Colour (c)
  gl.Torus (x, y, z, R, r)
}

func verTorus (c col.Colour, x, y, z, R, r, a float64) {
  gl.Colour (c)
  gl.VerTorus (x, y, z, R, r, a)
}

func paraboloid (c col.Colour, x, y, z, a, wx, wy float64) {
  gl.Colour (c)
  gl.Paraboloid (x, y, z, a, wx, wy)
}

func surface (c col.Colour, f Fxy2z, wx, wy float64) {
  gl.Colour (c)
  gl.Surface (f, wx, wy)
}
