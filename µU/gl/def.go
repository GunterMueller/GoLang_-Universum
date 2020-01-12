package gl

// (c) Christian Maurer   v. 191027 - license see µU.go

// #include <GL/gl.h>
import
  "C"
import (
  "µU/col"
  "µU/vect"
)
type
  Figure C.GLenum; const (
  Points = Figure(iota); Lines; LineLoop; LineStrip; Triangles; TriangleStrip;
           TriangleFan; Quads; QuadStrip; Polygons; NFigures; Light // ; Start
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
func ClearColor (c col.Colour) { clearColor(c) }
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

func Point (x ...float64) { point(x) }
func Line (x ...float64) { line(x) }
func Lineloop (x ...float64) { lineloop(x) }
func Linestrip (x ...float64) { linestrip(x) }
func Triangle (x...float64) { triangle(x) }
func Trianglestrip (x ...float64) { trianglestrip(x) }
func Trianglefan (x ...float64) { trianglefan(x) }
func Quad (x ...float64) { quad(x) }
func Quadstrip (x ...float64) { quadstrip(x) }
func Polygon (x ...float64) { polygon (x) }

func HorRectangle (x, y, z, x1, y1 float64, up bool) { horRectangle(x,y,z,x1,y1,up) }
func VertRectangle (x ...float64) { vertRectangle(x) }
func Parallelogram (x ...float64) { parallelogram(x) }
func Cube (x, y, z, a float64) { cube(x,y,z,a) }
func CubeC (c []col.Colour, x, y, z, a float64) { cubeC(c,x,y,z,a) }
func Cuboid (x ...float64) { cuboid (x) }
func CuboidC (c []col.Colour, x ...float64) { cuboidC(c,x) }
func Cuboid1 (x, y, z, w, d, h, a float64) { cuboid1(x,y,z,w,d,h,a) }
func Prism (x ...float64) { prism (x) }
func PrismC (c []col.Colour, x ...float64) { prismC (c,x) }
func Parallelepiped (x ...float64) { parallelepiped (x) }
func Pyramid (x, y, z, a, b, h float64) { pyramid (x,y,z,a,b,h) }
func PyramidC (c []col.Colour, x, y, z, a, b, h float64) { pyramidC (c,x,y,z,a,b,h) }
func Octahedron (x, y, z, r float64) { octahedron (x,y,z,r) }
func OctahedronC (c []col.Colour, x, y, z, r float64) { octahedronC (c, x, y, z, r) }
func Multipyramid (x ...float64) { multipyramid (x) }
func MultipyramidC (c []col.Colour, x ...float64) { multipyramidC (c,x) }
func Circle (x, y, z, r float64) { circle (x,y,z,r) }
func CircleSegment (x, y, z, r, a, b float64) { circleSegment (x,y,z,r,a,b) }
func VertCircle (x, y, z, r, a float64) { vertCircle (x,y,z,r,a) }
func Sphere (x, y, z, r float64) { sphere (x,y,z,r) }
func Cone (x, y, z, r, h float64) { cone (x,y,z,r,h) }
func DoubleCone (x, y, z, r, h float64) { doubleCone (x,y,z,r,h) }
func Cylinder (x, y, z, r, h float64) { cylinder (x,y,z,r,h) }
func CylinderSegment (x, y, z, r, h, a, b float64) { cylinderSegment (x,y,z,r,h,a,b) }
func HorCylinder (x, y, z, r, l, a float64) { horCylinder (x,y,z,r,l,a) }
func Torus (x, y, z, R, r float64) { torus (x,y,z,R,r) }
func HorTorus (x, y, z, R, r, a float64) { horTorus(x,y,z,R,r,a) }
// func Paraboloid (x, y, z, p float64) { paraboloid(x,y,z,p) }
// func HorParaboloid (x, y, z, p, a float64) { horParaboloid(x,y,z,p,a) }
func Curve (w uint, f1, f2, f3 func (float64) float64, t0, t1 float64) { curve(w,f1,f2,f3,t0,t1) }
func Surface (f func (float64, float64) float64, x, y, z, x1, y1, z1 float64, wd, ht, g uint) {
       surface (f,x,y,z,x1,y1,z1,wd,ht,g)
}
// func CoSy (x, y, z float64, mit bool) { coSy(x,y,z,mit) }

// furniture
// (x, y, z) = left front corner, (w, d, h) = width, depth and height

// p, l = thickness of table plate, legs
func Table (x, y, z, w, d, h, p, l, a float64, c col.Colour) {
  table(x,y,z,w,d,h,p,l,a,c)
}

// r, rf, rl = radius of table plate, foot and leg
// hf, hp = height of foot and plate (thickness)
func RoundTable (x, y, z, r, rf, rl, h, hf, hp float64, c col.Colour) {
  roundTable(x,y,z,r,rf,rl,h,hf,hp,c)
}

func OvalTable (x, y, z, w, d, h, a float64, c col.Colour) {
  ovalTable(x,y,z,w,d,h,a,c)
}

// wb, db, hb = width, depth and height of back and arm rests
// hs means the seat height
func ArmChair (x, y, z, w, d, h, wb, db, hs, hb, a float64, c col.Colour) {
  armChair(x,y,z,w,d,h,wb,db,hs,hb,a,c)
}

// h = seat height, p, l = thickness of seat plate, legs
// db, hb = depth, height of back rest
func Bench (x, y, z, w, d, h, p, l, db, hb, a float64, c col.Colour) {
  bench (x,y,z,w,d,h,p,l,db,hb,a,c)
}

func Chair (x, y, z, w, h, p, a float64, c col.Colour) {
  bench (x,y,z,w,w,h/2,p,h/20,h/20,h/2,a,c)
}

// walls, windows and doors

func SetHt (f, c float64) { height0, height = f, f + c }
func SetAng (a float64) { alpha = a }
func SetPos (x, y float64) { xx, yy = x, y }
func Pos() (float64, float64) { return xx, yy }
func Move (x, y float64) { xx += x; yy += y }

func SetColW (c col.Colour) { cWall = c }
func Wall (w float64) { wall(w) }

// w, d, h = width, depth, height
// r = Rand, v = Vorsprung
func Door (w, r, d, v, h float64, c col.Colour) { door(w,r,d,v,h,c) }

// r, rb, rt = Rand left/right, down, top; f = height Fensterbrett
func Window (w, d, h, f, r, rb, rt float64, b bool, c col.Colour) { window(w,d,h,f,r,rb,rt,b,c) }
func Window1 (w, d, h, f float64, c col.Colour) { window1(w,d,h,f,c) }

// light

// Pre: n < MaxL, 0 <= h[i] <= 1 für i = 0, 1.
// If Light n is already switched on, nothing has happened.
// Otherwise it is now switched on at position v
// in colour c with ambience h[0] and diffusion h[1].
func InitLight (n uint, v, h vect.Vector, r, g, b byte) { initLight(n,v,h,r,g,b) }

// Pre: Light n is switched on.
// Light n has position v.
func PosLight (n uint, v vect.Vector) { posLight(n,v) }

func ActualizeLight (n uint) { actLight(n) }
