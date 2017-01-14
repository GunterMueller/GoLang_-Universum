package fig

// (c) murus.org  v. 140217 - license see murus.go

// >>> experimental package; specifications under work.

import (
  "murus/col"; "murus/pts"
)
type (
  RealFunc func (float64) float64
  RealFunc2 func (float64, float64) float64
)

// Figures are defined by points; curved surfaces are approximated by polygons.

// TODO Spec
func Start (p pts.Points, x, y, z, xf, yf, zf float64) { start(p,x,y,z,xf,yf,zf) }

// TODO Spec
func Light (p pts.Points, n uint, x, y, z float64, ca, cd col.Colour) { light(p,n,x,y,z,ca,cd) }

// The point (x, y, z) with colour c is in s.
func Point (p pts.Points, x, y, z float64, c col.Colour) { point(p,x,y,z,c) }

// The line segment between (x, y, z) and (x1, y1, z1) with colour c is in s.
func Segment (p pts.Points, x, y, z, x1, y1, z1 float64, c col.Colour) { segment(p,x,y,z,x1,y1,z1,c) }

// Pre: Any two subsequent line segments are not parallel.
// The triangle between the points vi = (xi, yi, zi) (i = 0..2) with colour c
// is in s. Its orientation is [v1 - v0, v2 - v0].
func Triangle (p pts.Points, x0, y0, z0, x1, y1, z1, x2, y2, z2 float64, c col.Colour) { triangle(p,x0,y0,z0,x1,y1,z1,x2,y2,z2,c) }

// func TriangleFan (x, y, z []float64, n uint, c col.Colour)
// func TriangleStrip (x, y, z []float64, n uint, c col.Colour)

// Pre: Any two subsequent line segments are not parallel.
// The quadrangle between the points vi = (xi, yi, zi) (i = 0..2) with colour c
// is in s. Its orientation is [v1 - v0, v3 - v0].
func Quad (p pts.Points, x0, y0, z0, x1, y1, z1, x2, y2, z2, x3, y3, z3 float64, c col.Colour) { quad(p,x0,y0,z0,x1,y1,z1,x2,y2,z2,x3,y3,z3,c) }

// func QuadStrip (p pts.Points, x, y, z, x1, y1, z1 []float64, n uint, c col.Colour)

// Pre: x != x1, y != y1.
// Rectangle (x, y, z, x1, y, z, x1, y1, z, x, y1, z); oriented towards the positive z-axis, iff o == true.
func HorRectangle (p pts.Points, x0, y0, z0, x1, y1 float64, o bool, c col.Colour) { horRectangle(p,x0,y0,z0,x1,y1,o,c) }

// Pre: z != z1.
// Rectangle (x, y, z, x1, y1, z, x1, y1, z1, x, y, z1), oriented in direction [v1 - v0, v3 - v0].
func VertRectangle (p pts.Points, x, y, z, x1, y1, z1 float64, c col.Colour) { vertRectangle(p,x,y,z,x1,y1,z1,c) }

// Quad (x, y, z, x1, y1, z1, x1 + x2 - x, y1 + y2 - y, z1 + z2 - z, x2, y2, z2).
func Parallelogram (p pts.Points, x, y, z, x1, y1, z1, x2, y2, z2 float64, c col.Colour) { parallelogram(p,x,y,z,x1,y1,z1,x2,y2,z2,c) }

func Cuboid (p pts.Points, x0, y0, z0, x1, y1, z1 float64, c col.Colour) { cuboid(p,x0,y0,z0,x1,y1,z1,c) }

func Cuboid1 (p pts.Points, x, y, z, b, t, h, a float64, c col.Colour) { cuboid1(p,x,y,z,b,t,h,a,c) }

// Pre: At the moment: z[i] == z[0] for 0 < i < n, z[n] != z[0]. len(x) == len(y) == len(z).
// Corners = (x[0], y[0], z[0], ..., x[n-1], y[n-1], z[n-1]), top = (x[n], y[n], z[n]).
func Prism (p pts.Points, x, y, z []float64, c col.Colour) { prism(p,x,y,z,c) }

func Parallelepiped (p pts.Points, x0, y0, z0, x1, y1, z1, x2, y2, z2, x3, y3, z3 float64, c col.Colour) { parallelepiped(p,x0,y0,z0,x1,y1,z1,x2,y2,z2,x3,y3,z3,c) }

func Pyramid (p pts.Points, x0, y0, z0, x1, y1, z1, x2, y2, z2 float64, c col.Colour) { pyramid(p,x0,y0,z0,x1,y1,z1,x2,y2,z2,c) }

func Octahedron (p pts.Points, x, y, z, r float64, c col.Colour) { octahedron(p,x,y,z,r,c) }

// Pre: At the moment: z[i] == z[0] for 0 < i < n, z[n] != z[0]. len(x) == len(y) == len(z).
// Corners = (x[0], y[0], z[0], ..., x[n-1], y[n-1], z[n-1]), top = (x[n], y[n], z[n]).
func Multipyramid (p pts.Points, x, y, z []float64, c col.Colour) { multipyramid(p,x,y,z,c) }

// For r == 0 a point (x, y, z); otherwise a horizontal circle around (x, y, z) with radius abs(r);
// oriented nach oben, iff r > 0.
func Circle (p pts.Points, x, y, z, r float64, c col.Colour) { circle(p,x,y,z,r,c) }

// Pre: 0 <= a < b <= 360.
// Circle segment between a and b.
func CircleSegment (p pts.Points, x, y, z, r, a, b float64, c col.Colour) { circleSegment(p,x,y,z,r,a,b,c) }

func VertCircle (p pts.Points, x, y, z, r, a float64, c col.Colour) { vertCircle(p,x,y,z,r,a,c) }

// Pre: r >= 0. // TODO
// For r == 0 a point, otherwise a sphere around (x, y, z) with radius r;
// oriented to the outside, iff r > 0.
func Sphere (p pts.Points, x, y, z, r float64, c col.Colour) { sphere(p,x,y,z,r,c) }

// Standing cone (i.e. with rotation axis parallel to the z-axis)
// with (x, y, z) as middlepoint of its bottom circle, radius r and height h.
func Cone (p pts.Points, x, y, z, r, h float64, c col.Colour) { cone(p,x,y,z,r,h,c) }

func Frustum (p pts.Points, x, y, z, r, h, h1 float64, c col.Colour) { frustum(p,x,y,z,r,h,h1,c) }

func DoubleCone (p pts.Points, x, y, z, r, h float64, c col.Colour) { doubleCone(p,x,y,z,r,h,c) }

func Cylinder (p pts.Points, x, y, z, r, h float64, c col.Colour) { cylinder(p,x,y,z,r,h,c) }

func HorCylinder (p pts.Points, x, y, z, r, l, a float64, c col.Colour) { horCylinder(p,x,y,z,r,l,a,c) }

func CylinderSegment (p pts.Points, x, y, z, r, h, a, b float64, c col.Colour) { cylinderSegment(p,x,y,z,r,h,a,b,c) }

// Pre: 0 < r, 0 < R.
// In die Punktfolge ist ein Torus mit einer vertikalen (zur z-Achse parallelen) Rotationsachse eingefügt.
// (x, y, z) ist sein Mittelpunkt, R der Radius des Mittelkreises und r der Radius seines Ringes.
func Torus (p pts.Points, x, y, z, R, r float64, c col.Colour) { torus(p,x,y,z,R,r,c) }

// Pre: 0 < r, 0 < R, -180 < a < 180.
// In die Punktfolge ist ein Torus mit einer horizontalen (zur x-y-Ebene parallelen) Rotationsachse eingefügt.
// (x, y, z) ist sein Mittelpunkt, R der Radius des Mittelkreises, r der Radius seines Ringes
// und a der Winkel (in °) zwischen der x-Achse und seiner Rotationsachse.
func HorTorus (p pts.Points, x, y, z, R, r, a float64, c col.Colour) { horTorus(p,x,y,z,R,r,a,c) }

//  func Paraboloid (p pts.Points, x, y, z, p float64, c col.Colour)

//  func HorParaboloid (p pts.Points, x, y, z, p, a float64, c col.Colour)

func Curve (p pts.Points, f1, f2, f3 RealFunc, t0, t1 float64, c col.Colour) { curve(p,f1,f2,f3,t0,t1,c) }

func Surface (p pts.Points, f RealFunc2, x, y, z, x1, y1, z1 float64, c col.Colour) { surface (p,f,x,y,z,x1,y1,z1,c) }

// Pre: x, y, z > 0.
//  func CoSy (x, y, z float64, with bool)

//  func Tree (p pts.Points, x, y, z, r float64, c col.Colour)
