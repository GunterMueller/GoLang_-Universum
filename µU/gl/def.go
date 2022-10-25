package gl

// (c) Christian Maurer   v. 221024 - license see µU.go

// #include <GL/gl.h>
import
  "C"
import (
  "µU/obj"
  "µU/col"
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
  MaxL = 16 // <= GL.GL_MAX_LIGHTS
  RGB = C.GL_RGB
  Flat = C.GL_FLAT
  Double = C.GL_DOUBLE
  Projection = C.GL_PROJECTION
  Modelview = C.GL_MODELVIEW
  Depthtest = C.GL_DEPTH_TEST
  ColorMaterial = C.GL_COLOR_MATERIAL
)

func Clear() { clear() }
func Cls (c col.Colour) { cls(c) }
func Colour (c col.Colour) { colour(c) }
func ClearColour (c col.Colour) { clearColour(c) }
func Begin (f Figure) { begin(f) }
func End() { end() }
func Vertex (x, y, z float64) { vertex(x,y,z) }
func NewList (n uint) { newList(n) }
func EndList() { endList() }
func CallList (n uint) { callList(n) }
func Frustum (l, r, b, t, h, f float64) { frustum (l,r,b,t,h,f) }
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
func Error() string { return error() }

// Figures in the 3-dimensional space ///////////////////////////////////////////////////////////

// Pre: len(x) == 3.
func Point (x ...float64) { point(x...) }

// Pre: len(x) == 6.
func Line (x ...float64) { line(x...) }

// Long series of lines must be included by "begin (Lines)" ... "end()"

// Pre: len(x) % 3 == 0, len(x) >= 9.
func Lineloop (x ...float64) { lineloop(x...) }
func Linestrip (x ...float64) { linestrip(x...) }

// Pre: len(x) == 9.
func Triangle (x ...float64) { triangle(x...) }

// Pre: len(x) % 3 == 0, len(x) >= 12.
func Trianglestrip (x ...float64) { trianglestrip(x...) }
func Trianglefan (x ...float64) { trianglefan(x...) }

// Pre: len(x) == 12.
func Quad (x ...float64) { quad(x...) }

// Pre: len(x) % 3 == 0, len(x) >= 18.
func Quadstrip (x ...float64) { quadstrip(x...) }

func HorRectangle (x, y, z, x1, y1 float64, up bool) { horRectangle(x,y,z,x1,y1,up) }

// Pre: len(x) == 6, x[2] != x[5].
func VertRectangle (x ...float64) { vertRectangle(x...) }

// Pre: len(x) == 9.
func Parallelogram (x ...float64) { parallelogram(x...) }

// Pre: len(x) % 3 == 0, len(x) >= 9.
func Polygon (x ...float64) { polygon (x...) }
func Polygon1 (x, y, z []float64) { polygon1 (x,y,z) }

func Curve (w uint, f, f1, f2 func (float64) float64, t, t1 float64) { curve(w,f,f1,f2,t,t1) }

// Pre: a != 0.
func Cube (x, y, z, a float64) { cube(x,y,z,a) }
func CubeC (c []col.Colour, x, y, z, a float64) { cubeC(c,x,y,z,a) }

// Pre: len(x) == 6, x[2] != x[5].
func Cuboid (x ...float64) { cuboid (x...) }
// Pre: len(c) == 6.
func CuboidC (c []col.Colour, x ...float64) { cuboidC(c,x...) }

func Cuboid1 (x, y, z, w, d, h, a float64) { cuboid1(x,y,z,w,d,h,a) }

// Pre: len(x) % 3 == 0, len(x) >= 12.
func Prism (x ...float64) { prism (x...) }
// Pre: len(x) == len(x).
func PrismC (c []col.Colour, x ...float64) { prismC (c,x...) }

// Pre: len(x) == 12.
func Parallelepiped (x ...float64) { parallelepiped (x...) }

func Pyramid (x, y, z, a, b, h float64) { pyramid (x,y,z,a,b,h) }
func PyramidC (c []col.Colour, x, y, z, a, b, h float64) { pyramidC (c,x,y,z,a,b,h) }

// Pre: len(x) % 3 == 0, len(x) >= 12.
func Multipyramid (x ...float64) { multipyramid (x...) }
func MultipyramidC (c []col.Colour, x ...float64) { multipyramidC (c,x...) }

func Octahedron (x, y, z, r float64) { octahedron (x,y,z,r) }
// Pre: len(c) == 8.
func OctahedronC (c []col.Colour, x, y, z, r float64) { octahedronC (c, x, y, z, r) }

func Circle (x, y, z, r float64) { circle (x,y,z,r) }

func CircleSegment (x, y, z, r, a, b float64) { circleSegment (x,y,z,r,a,b) }

func VertCircle (x, y, z, r, a float64) { vertCircle (x,y,z,r,a) }

func Sphere (x, y, z, r float64) { sphere (x,y,z,r) }

func Cone (x, y, z, r, h float64) { cone (x,y,z,r,h) }

func DoubleCone (x, y, z, r, h float64) { doubleCone (x,y,z,r,h) }

func Cylinder (x, y, z, r, h float64) { cylinder (x,y,z,r,h) }

func CylinderSegment (x, y, z, r, h, a, b float64) { cylinderSegment (x,y,z,r,h,a,b) }

func HorCylinder (x, y, z, r, l, a float64) { horCylinder (x,y,z,r,l,a) }

// Pre: R > 0, r > 0.
func Torus (x, y, z, R, r float64) { torus (x,y,z,R,r) }
func VerTorus (x, y, z, R, r, a float64) { verTorus(x,y,z,R,r,a) }

func Paraboloid (x, y, z, a, w, h float64) { paraboloid(x,y,z,a,w,h) }

func Surface (f obj.XY2Z, w, h float64) { surface (f,w,h) }

// Furniture ////////////////////////////////////////////////////////////////////////////////////
// (x, y, z) = left front corner, (w, d, h) = width, depth and height, a = angle

// p, l = thickness of table plate, legs
func Table (x, y, z, w, d, h, p, l, a float64, c col.Colour) { table(x,y,z,w,d,h,p,l,a,c) }

// r, rf, rl = radius of table plate, foot and leg, hf, hp = height (thickness) of foot and plate
func RoundTable (x, y, z, r, rf, rl, h, hf, hp float64, c col.Colour) {
     roundTable(x,y,z,r,rf,rl,h,hf,hp,c) }

// Pre: d <= w.
// d = length of straight part
func OvalTable (x, y, z, w, d, h, a float64, c col.Colour) { ovalTable(x,y,z,w,d,h,a,c) }

// wb, db, hb = width, depth and height of back and arm rests
// p = thickness of seat plate
func Chair (x, y, z, w, h, p, a float64, c col.Colour) { chair (x,y,z,w,h,p,a,c) }

// wb, db, hb = width, depth and height of back and arm rests, hs = seat height
func ArmChair (x, y, z, w, d, h, wb, db, hs, hb, a float64, c col.Colour) {
     armChair(x,y,z,w,d,h,wb,db,hs,hb,a,c) }

// h = seat height, p, l = thickness of seat plate, legs, db, hb = depth, height of back rest
func Bench (x, y, z, w, d, h, p, l, db, hb, a float64, c col.Colour) {
     bench (x,y,z,w,d,h,p,l,db,hb,a,c) }

// Walls, Windows and Doors /////////////////////////////////////////////////////////////////////

func SetHt (f, c float64) { height0, height = f, f + c }
func SetAng (a float64) { alpha = a }
func SetPos (x, y float64) { xx, yy = x, y }
func Pos() (float64, float64) { return xx, yy }
func Move (x, y float64) { xx += x; yy += y }

func SetColW (c col.Colour) { cWall = c }
func Wall (w float64) { wall(w) }

// w, d, h = width, depth, height, f = door framg, p = door protrusion
func Door (w, f, d, p, h float64, c col.Colour) { door(w,f,d,p,h,c) }

// f, fb, ft = window frame left/right, bottom, top; wc = height of window cill
func Window (w, d, h, wc, f, fb, ft float64, b bool, c col.Colour) {
     window(w,d,h,wc,f,fb,ft,b,c) }
func Window1 (w, d, h, f float64, c col.Colour) { window1(w,d,h,f,c) }

// Light ////////////////////////////////////////////////////////////////////////////////////////

// Pre: n < MaxL, 0 <= h[i] <= 1 für i = 0, 1.
// If Light n is already switched on, nothing has happened.
// Otherwise it is now switched on at position v0, v1, v2 in colour c
// with ambience h[0] and diffusion h[1]. // XXX
func InitLight (n uint, x, y, z, h0, h1, h2 float64, r, g, b byte) {
  initLight(n,x,y,z,h0,h1,h2,r,g,b)
}

// Pre: n < MaxL; Light n is switched on.
// Light n has position x, y, z.
func PosLight (n uint, x, y, z float64) { posLight(n,x,y,z) }

// Pre: n < MaxL.
func ActualizeLight (n uint) { actLight(n) }

// Pre: n < MaxL.
func Lamp (n uint) { lamp(n) }
