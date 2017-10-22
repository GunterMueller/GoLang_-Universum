package gl

// (c) Christian Maurer   v. 170908 - license see µU.go

// #include <GL/gl.h>
import
  "C"
import
  "µU/col"
const (
  POINTS = C.GLenum(iota)
  LINES
  LINE_LOOP
  LINE_STRIP
  TRIANGLES
  TRIANGLE_STRIP
  TRIANGLE_FAN
  QUADS
  QUAD_STRIP
  POLYGON
  NClasses
)
const
  MaxL = 16 // <= GL.GL_MAX_LIGHTS

// TODO Spec
func Init (f float64, w, h uint) { init_(f,w,h) } // XXX

func Colour (c col.Colour) { colour(c) }
func Vertex (x, y, z float64) { vertex(x,y,z) }
func VertexC (x, y, z float64, c col.Colour) { vertexC(x,y,z,c) }

func Cls (r, g, b byte) { cls(r,g,b) }

func NewList (n uint) { newList(n) }

func EndList() { endList() }

func CallList (n uint) { callList(n) }

// Pre: n < MaxL, 0 <= h[i] <= 1 für i = 0, 1.
// If Light n is already switched on, nothing has happened.
// Otherwise it is now switched on at position v
// in colour c with ambience h[0] and diffusion h[1].
func InitLight (v, v1, v2, h, h1, h2 float64, r, g, b byte) { initLight(v,v1,v2,h,h1,h2,r,g,b) }

// Pre: Light n is switched on.
// Light n has position v.
func PosLight (n uint, v, v1, v2 float64) { posLight(n,v,v1,v2) }

// TODO Spec
func ActualizeLight (n uint) { actLight(n) } // n < MaxL

// TODO Spec
func ShowLight (on bool) { lightVis = on }

func Clear() { clear() }
func Begin (c C.GLenum) { begin(c) }
func End() { end() }
// func Colour (r, g, b byte) { colour(r,g,b) }

func Point (c col.Colour, x ...float64) { point(c,x) }
func Line (c col.Colour, x ...float64) { line(c,x) }
func Lineloop (c col.Colour, x ...float64) { lineloop(c,x) }
func Linestrip (c col.Colour, x ...float64) { linestrip (c,x) }
func Triangle (c col.Colour, x...float64) { triangle (c,x) }
func Trianglestrip (c col.Colour, x ...float64) { trianglestrip(c,x) }
func Trianglefan (c col.Colour, x ...float64) { trianglefan(c,x) }
func Quad (c col.Colour, x ...float64) { quad(c,x) }
func Quadstrip (c col.Colour, x ...float64) { quadstrip(c,x) }
func Polygon (c col.Colour, x ...float64) { polygon (c,x) }

func HorRectangle (x, y, z, x1, y1 float64, up bool) { horRectangle(x,y,z,x1,y1,up) }
func HorRectangleC (c col.Colour, x, y, z, x1, y1 float64, up bool) { horRectangleC(c,x,y,z,x1,y1,up) }
func VertRectangle (x ...float64) { vertRectangle(x) }
func VertRectangleC (c col.Colour, x ...float64) { vertRectangleC (c,x) }
func Parallelogram (c col.Colour, x ...float64) { parallelogram(c,x) }
func Cube (c col.Colour, x, y, z, a float64) { cube(c,x,y,z,a) }
func CubeC (c []col.Colour, x, y, z, a float64) { cubeC(c,x,y,z,a) }
func Cuboid (c col.Colour, x ...float64) { cuboid (c,x) }
func CuboidC (c []col.Colour, x ...float64) { cuboidC(c,x) }
func Cuboid1 (c col.Colour, x, y, z, b, t, h, a float64) { cuboid1(c,x,y,z,b,t,h,a) }
func Prism (c col.Colour, x ...float64) { prism (c,x) }
func PrismC (c []col.Colour, x ...float64) { prismC (c,x) }
func Parallelepiped (c col.Colour, x ...float64) { parallelepiped (c,x) }
func Pyramid (c col.Colour, x, y, z, a, b, h float64) { pyramid (c,x,y,z,a,b,h) }
func PyramidC (c []col.Colour, x, y, z, a, b, h float64) { pyramidC (c,x,y,z,a,b,h) }
func Octahedron (c col.Colour, x, y, z, r float64) { octahedron (c,x,y,z,r) }
func OctahedronC (c []col.Colour, x, y, z, r float64) { octahedronC (c, x, y, z, r) }
func Multipyramid (c col.Colour, x ...float64) { multipyramid (c,x) }
func MultipyramidC (c []col.Colour, x ...float64) { multipyramidC (c,x) }
func Circle (c col.Colour, x, y, z, r float64) { circle (c,x,y,z,r) }
func CircleSegment (c col.Colour, x, y, z, r, a, b float64) { circleSegment (c,x,y,z,r,a,b) }
func VertCircle (c col.Colour, x, y, z, r, a float64) { vertCircle (c,x,y,z,r,a) }
func Sphere (c col.Colour, x, y, z, r float64) { sphere (c,x,y,z,r) }
func Cone (c col.Colour, x, y, z, r, h float64) { cone (c,x,y,z,r,h) }
// func Frustum (c col.Colour, x, y, z, r, h, h1 float64) { frustum (c,x,y,z,r,h,h1) }
func DoubleCone (c col.Colour, x, y, z, r, h float64) { doubleCone (c,x,y,z,r,h) }
func Cylinder (c col.Colour, x, y, z, r, h float64) { cylinder (c,x,y,z,r,h) }
func CylinderSegment (c col.Colour, x, y, z, r, h, a, b float64) { cylinderSegment (c,x,y,z,r,h,a,b) }
func HorCylinder (c col.Colour, x, y, z, r, l, a float64) { horCylinder (c,x,y,z,r,l,a) }
func Torus (c col.Colour, x, y, z, R, r float64) { torus (c,x,y,z,R,r) }
func HorTorus (c col.Colour, x, y, z, R, r, a float64) { horTorus(c,x,y,z,R,r,a) }
// func Paraboloid (c col.Colour, x, y, z, p float64) { paraboloid(c,x,y,z,p) }
// func HorParaboloid (c col.Colour, x, y, z, p, a float64) { horParaboloid(c,x,y,z,p,a) }
func Curve (c col.Colour, w uint, f1, f2, f3 func (float64) float64, t0, t1 float64) { curve(c,w,f1,f2,f3,t0,t1) }
func Surface (c col.Colour, f func (float64, float64) float64, x, y, z, x1, y1, z1 float64, wd, ht, g uint) { surface (c,f,x,y,z,x1,y1,z1,wd,ht,g) }
// func CoSy (x, y, z float64, mit bool) { coSy(x,y,z,mit) }
