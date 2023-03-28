package gl

// (c) Christian Maurer   v. 230322 - license see µU.go
//
// >>> At the moment, nothing has been said about the orientation of the figures. TODO

// #include <GL/gl.h>
import
  "C"
import (
  "µU/obj"
  "µU/col"
)
type (
  F = C.GLfloat
  D = C.GLdouble
)
type
  Figure C.GLenum; const (
  Points = Figure(iota)
  Lines
  LineLoop
  LineStrip
  Triangles
  TriangleStrip
  TriangleFan
  Quads
  QuadStrip
  Polygons
  NFigures
  Light
//  Start
)
const (
  RGB = C.GL_RGB
  Double = C.GL_DOUBLE
  Projection = C.GL_PROJECTION
  Modelview = C.GL_MODELVIEW
  Depthtest = C.GL_DEPTH_TEST
  Flat = C.GL_FLAT
  Smooth = C.GL_SMOOTH
)

// For the specification of the following functions,
// please consult the technical literature on OpenGL.
func Clear() { clear() }
func ClearDepth (d float64) { clearDepth(d) }
func Cls (c col.Colour) { cls(c) }
func Colour (c col.Colour) { colour(c) }
func ClearColour (c col.Colour) { clearColour(c) }
func Begin (f Figure) { begin(f) }
func End() { end() }
func Vertex (x, y, z float64) { vertex(x,y,z) }
func NewList (n uint) { newList(n) }
func EndList() { endList() }
func CallList (n uint) { callList(n) }
func Frustum (l, r, b, t, n, f float64) { frustum (l,r,b,t,n,f) }
func Ortho (l, r, b, t, n, f float64) { ortho(l,r,b,t,n,f) }
func LoadIdentity() { C.glLoadIdentity() }
func Viewport (x, y int, w, h uint) { viewport(x,y,w,h) }
func MatrixMode (m uint) { matrixMode(m) }
func Translate (x, y, z float64) { translate(x,y,z) }
func Rotate (a, x, y, z float64) { rotate(a,x,y,z) }
func Scale (x, y, z float64) { scale(x,y,z) }
func PushMatrix() { pushMatrix() }
func PopMatrix() { popMatrix() }
func Enable (i uint) { enable(i) }
func ShadeModel (m uint) { shadeModel(m) }
func Flush() { flush() }
func Error() string { return error() }

// Figures in the 3-dimensional space ///////////////////////////////////////////////////////////

// Pre: len(x) == 3.
// A point is created at (x[0], x[1], x[2]).
func Point (x ...float64) { point(x...) }

//
func Linewidth (w float64) { linewidth(w) }

// Pre: len(x) == 6.
// A line is created from (x[0], x[1], x[2]) to (x[3], x[4], x[5]).
func Line (x ...float64) { line(x...) }

// Pre: len(x) % 3 == 0, len(x) >= 9.
// A sequence of  connected lines is created, from (x[0], x[1], x[2]) to (x[3], x[4], x[5]),
// from (x[3], x[4], x[5]) to (x[6], x[7], x[8]), and so on, until
// from (x[n-6], x[n-5], x[n-4]) to (x[n-3], x[n-2], x[n-1]).
func Lineloop (x ...float64) { lineloop(x...) }

// Pre: len(x) % 3 == 0, len(x) >= 9.
// Like Lineloop, with an additional line from (x[n-3], x[n-2], x[n-1]) to (x[0], x[1], x[2]).
func Linestrip (x ...float64) { linestrip(x...) }

// Pre: len(x) == 9.
// A triangle is created between (x[0], x[1], x[2]), (x[3], x[4], x[5]) and (x[6], x[7], x[8]).
func Triangle (x ...float64) { triangle(x...) }

// Pre: len(x) % 3 == 0, len(x) >= 12.
// A series of connected triangles is created, starting with the triangle between
// (x[0], x[1], x[2]), (x[3], x[4], x[5]) and (x[6], x[7], x[8]), followed by the triangle
// between (x[3], x[4], x[5]), (x[6], x[7], x[8]) and (x[9], x[10], x[11]), and so on.
func Trianglestrip (x ...float64) { trianglestrip(x...) }

// Pre: len(x) % 3 == 0, len(x) >= 12.
// A series of connected triangles is created, that fan around the central point
// (x[0], x[1], x[2]). The second triangle is given by the central point, (x[6], x[7], x[8])
// and (x[9], x[10], x[11]), and so on. 
func Trianglefan (x ...float64) { trianglefan(x...) }

// Pre: len(x) == 12, all used points lie in a plane.
// A quadrangle is created between (x[0], x[1], x[2]), (x[3], x[4], x[5]), (x[6], x[7], x[8])
// and (x[9], x[10], x[11]).
func Quad (x ...float64) { quad(x...) }

// Pre: len(x) % 3 == 0, len(x) >= 18.
// Specification analogous to Triangle -> Trianglestrip.
func Quadstrip (x ...float64) { quadstrip(x...) }

// A horizontal rectangle with edges parallel to the x- and y-axis is created
// between (x, y, z), (x1, y, z), (x1, y1, z) and (x, y, z).
func HorRectangle (x, y, z, x1, y1 float64, up bool) { horRectangle(x,y,z,x1,y1,up) }

// Pre: len(x) == 6, x[2] != x[5].
// A rectangle with horizontal edges parallel to the x-y-plane and vertical edges
// parallel to the z-axis is created between (x[0], x[1], x[2]) and (x[3], x[4], x[5]).
func VertRectangle (x ...float64) { vertRectangle(x...) }

// Pre: len(x) == 9.
// A parallelogram is created between the points at
// (x[0], x[1], x[2]), (x[0]+x[3], x[1]+x[4], x[2]+x[5]),
// (x[0]+x[3]+x[6], x[1]+x[4]+x[7], x[2]+x[5]+x[8]) and (x[0]+x[6], x[1]+x[7], x[2]+x[8]).
func Parallelogram (x ...float64) { parallelogram(x...) }

// Pre: len(x) % 3 == 0, len(x) >= 9, all used points lie in a plane.
// A polygon is created between the points at (x[0], x[1], x[2]), (x[3], x[4], x[5]), ...,
// (x[n-3], x[n-2], x[n-1]), where n == len(x).
func Polygon (x ...float64) { polygon (x...) }

// Pre: t0 < t1.
// The curve given by f is drawn from t == t0 until t == t1.
func Curve (f obj.Ft2xyz, t0, t1 float64) { curve(f,t0,t1) }

// Pre: wx > 0, wy > 0.
// The bounded plane within the area -wx <= x <= wx and -wy <= y <= wy,
// defined by f(x,y) = a * x + b * y + c, is created.
func Plane (a, b, c, wx, wy float64) { plane(a,b,c,wx,wy) }

// Pre: a != 0.
// A cube with edges parallel to the coordinate axes is created
// with the center at (x, y, z) and the edge length a.
func Cube (x, y, z, a float64) { cube(x,y,z,a) }

// Pre: a != 0. len(c) == 6.
// See Cube. It has the front/right/back/left/top/bottom colours c[0]..c[5].
func CubeC (c []col.Colour, x, y, z, a float64) { cubeC(c,x,y,z,a) }

// Pre: len(x) == 6, x[0] != x[3], x[1] != x[4] and x[2] != x[5].
// A cuboid with edges parallel to the coordinate axes is created
// between the points at (x[0], x[1], x[2]) and (x[3], x[4], x[5]).
func Cuboid (x ...float64) { cuboid (x...) }

// Pre: see Couboid and len(c) == 6.
// See Cuboid. It has the front/right/back/left/top/bottom colours c[0]..c[5].
func CuboidC (c []col.Colour, x ...float64) { cuboidC(c,x...) }

// Pre: w != 0, d !=0 and h != 0.
// A cuboid with a vertical edge parallel to the z-axis is created
// between the points at (x, y, z) and (x + dx, y + dy, z + dz) and it is rotated
// by a degrees around the line parallel to the z-axis through the point (x, y, z).
func Cuboid1 (x, y, z, dx, dy, dz, a float64) { cuboid1(x,y,z,dx,dy,dz,a) }

// Pre: len(x) % 3 == 0, len(x) >= 12.
// A prism without bottom and top is created. Its bottom corners are (x[3], x[4], x[5]),
// (x[6], x[7], x[8]) and so on, its top corners are the bottom corners plus (x[0], x[1], x[2]).
func Prism (x ...float64) { prism (x...) }

// Pre: len(x) % 3 == 0, len(c) == len(x) >= 12.
// See Prism. Its colours are given by c.
func PrismC (c []col.Colour, x ...float64) { prismC (c,x...) }

// Pre: len(x) == 12.
// A parallelepiped is created. One of its corners is c = (x[0], x[1], x[2]), the others
// are c + (x[3], x[4], x[5]), c + (x[6], x[7], x[8]) and c + (x[9], x[19], x[11]).
func Parallelepiped (x ...float64) { parallelepiped (x...) }

// Pre: len(x) == 12, len(c) == 6.
// See Parallelepiped. Its colours are given by c.
func ParallelepipedC (c []col.Colour, x ...float64) { parallelepipedC (c,x...) }

// Pre: a > 0, h != 0.
// A pyramid of height h with the center (x, y, z) of its horizonal bottom is created,
// its bottom edges have the length a.
func Pyramid (x, y, z, a, h float64) { pyramid (x,y,z,a,h) }

// Pre: a > 0, h != 0, len(c) == 5.
// See Pyramid. Its colours front/right/back/left/bottom are given by c.
func PyramidC (c []col.Colour, x, y, z, a, h float64) { pyramidC (c,x,y,z,a,h) }

// Pre: len(x) % 2 == 0, len(c) >= 6; the line loop through the c[i] is convex.
// A multipyramid of height h with top (x, y, z + h) and corners (c[0], c[1], 0),
// (c[2], c[3], 0) and so on is created.
func Multipyramid (x, y, z, h float64, c ...float64) { multipyramid (x,y,z,h,c...) }

// Pre: len(x) % 2 == 0, len(c) >= 6, len(c) == len(x) / 2 + 1.
// See Multipyramid. Its colours are given by f.
func MultipyramidC (f []col.Colour,x,y,z,h float64,c ...float64) {multipyramidC (f,x,y,z,h,c...) }

// Pre: len(x) % 3 == 0, len(x) >= 12.
// An octopus with top (x[0], x[1], x[2]) and corners (x[3], x[4], x[5]),
// (x[6], x[7], x[8]) and so on is created.
func Octopus (x ...float64) { octopus (x...) }

// Pre: len(x) % 3 == 0, len(x) >= 12, len(c) == len(x) / 3.
// See Octopus. Its colours are given by c.
func OctopusC (c []col.Colour, x ...float64) { octopusC (c,x...) }

// Pre: r != 0.
// An octahedron with the center (x, y, z) and length e of its edges is created.
func Octahedron (x, y, z, e float64) { octahedron (x,y,z,e) }

// Pre: r != 0, len(c) == 8.
// See Octahedron. Its colours are given by c.
func OctahedronC (c []col.Colour, x, y, z, e float64) { octahedronC (c, x, y, z, e) }

// Pre: r != 0.
// A horizontal circle with the center (x, y, z) and the radius r is created.
func Circle (x, y, z, r float64) { circle (x,y,z,r) }

// Pre: r != 0.
// The segment of the horizontal circle given by (x, y, z, r) 
// between the angles of a and b degrees is created.
func CircleSegment (x, y, z, r, a, b float64) { circleSegment (x,y,z,r,a,b) }

// Pre: r != 0.
// A vertical circle with the center (x, y, z) and the radius r is created,
// that lies in the y-z-plane rotated by a degrees around the y-axis.
func VertCircle (x, y, z, r, a float64) { vertCircle (x,y,z,r,a) }

// Pre: r != 0.
// A sphere is created with the center (x, y, z) and the radius r.
func Sphere (x, y, z, r float64) { sphere (x,y,z,r) }

// Pre: r != 0, h != 0.
// A cone of height h is created with the horizontal circle around (x, y, z) with radius r
// as its bottom.
func Cone (x, y, z, r, h float64) { cone (x,y,z,r,h) }

// Pre: r != 0, h != 0.
// Two cones of height h are created, one with the horizontal circle around (x, y, z - h)
// as bottom and the other with the horizontal circle around (x, y, z + h) as top.
func DoubleCone (x, y, z, r, h float64) { doubleCone (x,y,z,r,h) }

// Pre: r != 0, h != 0.
// A cylinder of radius r and height h is created with the horizontal circle around (x, y, z)
// with radius r as bottom and the horizontal circle around (x, y, z + h) as top.
func Cylinder (x, y, z, r, h float64) { cylinder (x,y,z,r,h) }

// Pre: r != 0, h != 0, len(c) == 2.
// See Cylinder. Its colour is c[0] and the colour of its bottom and top is c[1].
func CylinderC (c[]col.Colour, x, y, z, r, h float64) { cylinderC (c,x,y,z,r,h) }

// Pre: r != 0, h != 0.
// The segment of the cylinder given by (x, y, z, r, h)
// between the angles of a and b degrees is created.
func CylinderSegment (x, y, z, r, h, a, b float64) { cylinderSegment (x,y,z,r,h,a,b) }

// Pre: r != 0.
// A horizontal cylinder of radius r and length l is created,
// ended by vertical circles around (x, y, z) and (x + l, y, z).
func HorCylinder (x, y, z, r, l, a float64) { horCylinder (x,y,z,r,l,a) }

// Pre: r != 0, len(c) == 2.
// See HorCylinder. Its colour is c[0] and the colour of the ends is c[1].
// ended by vertical circles around (x, y, z) and (x + l, y, z).
func HorCylinderC (c []col.Colour, x, y, z, r, l, a float64) { horCylinderC (c,x,y,z,r,l,a) }

// Pre: R > 0, r > 0.
// A horizontal torus with the center at (x, y, z),
// the inner radius R-r and the outer radius R+r is created.
func Torus (x, y, z, R, r float64) { torus (x,y,z,R,r) }

// Pre: R > 0, r > 0.
// A vertical torus with the center at (x, y, z), 
// the inner radius R-r and the outer radius R+r is created and rotated
// by a degrees around the line parallel to the z-axis through the point (x, y, z).
func VerTorus (x, y, z, R, r, a float64) { verTorus(x,y,z,R,r,a) }

// Pre: a != 0, wx > 0, wy > 0.
// A paraboloid within the area -wx <= x <= wx and -wy <= y <= wy is created with
// base point (x0, y0, z0), defined by f(x, y) = a^2 * ((x - x0)^2 + (y - y0)^2).
func Paraboloid (x0, y0, z0, a, wx, wy float64) { paraboloid(x0,y0,z0,a,wx,wy) }

// Pre: wx > 0, wy > 0.
// The bounded surface within the area -wx <= x <= wx and -wy <= y <= wy,
// given by the function f is created.
func Surface (f obj.Fxy2z, wx, wy float64) { surface (f,wx,wy) }
