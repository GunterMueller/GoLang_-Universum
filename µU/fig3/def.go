package fig3

// (c) Christian Maurer  v. 221113 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)

// The spezifications of all functions are found in the file µU/gl/def.go.

func Point (c col.Colour, x, y, z float64) { point(c,x,y,z) }
func Line (c col.Colour, x ...float64) { line(c,x...) }
func Lineloop (c col.Colour, x ...float64) { lineloop(c,x...) }
func Linestrip (c col.Colour, x ...float64) { linestrip(c,x...) }
func Triangle (c col.Colour, x, y, z, x1, y1, z1, x2, y2, z2 float64) {
               triangle(c,x,y,z,x1,y1,z1,x2,y2,z2) }
func Trianglestrip (c col.Colour, x ...float64) { trianglestrip(c,x...) }
func Trianglefan (c col.Colour, x ...float64) { trianglefan(c,x...) }
func Quad (c col.Colour, x ...float64) { quad(c,x...) }
func Quadstrip (c col.Colour, x ...float64) { quadstrip(c,x...) }
func HorRectangle (c col.Colour, x, y, z, x1, y1 float64, up bool) {
                   horRectangle (c,x,y,z,x1,y1,up) }
func VertRectangle (c col.Colour, x ...float64) { vertRectangle(c,x...) }
func Parallelogram (c col.Colour, x ...float64) { parallelogram(c,x...) }
func Polygon (c col.Colour, x ...float64) { polygon(c,x...) }
func Curve (c col.Colour, f Ft2xyz, t0, t1 float64) { curve(c,f,t0,t1) }
func Plane (f col.Colour, a, b, c, wx, wy float64) { plane(f,a,b,c,wx,wy) }
func Cube (c col.Colour, x, y, z, a float64) { cube(c,x,y,z,a) }
func CubeC (c []col.Colour, x, y, z, a float64) { cubeC(c,x,y,z,a) }
func Cuboid (c col.Colour, x, y, z, x1, y1, z1 float64) { cuboid(c,x,y,z,x1,y1,z1) }
func CuboidC (c []col.Colour, x, y, z, dx, dy, dz float64) { cuboidC(c,x,y,z,dx,dy,dz) }
func Cuboid1 (c col.Colour, x, y, z, dx, dy, dz, a float64) { cuboid1(c,x,y,z,dx,dy,dz,a) }
func Prism (c col.Colour, x ...float64) { prism (c,x...) }
func PrismC (c []col.Colour, x ...float64) { prismC (c,x...) }
func Parallelepiped (c col.Colour, x ...float64) { parallelepiped(c,x...) }
func ParallelepipedC (c []col.Colour, x ...float64) { parallelepipedC(c,x...) }
func Pyramid (c col.Colour, x, y, z, a, h float64) { pyramid(c,x,y,z,a,h) }
func PyramidC (c []col.Colour, x, y, z, a, h float64) { pyramidC(c,x,y,z,a,h) }
func Multipyramid (f col.Colour, x, y, z, h float64, c ...float64) { multipyramid(f,x,y,z,h,c...)}
func MultipyramidC (f []col.Colour, x,y,z,h float64,c ...float64) { multipyramidC(f,x,y,z,h,c...)}
func Octopus (c col.Colour, x ...float64) { octopus(c,x...) }
func OctopusC (c []col.Colour, x ...float64) { octopusC(c,x...) }
func Octahedron (c col.Colour, x, y, z, r float64) { octahedron(c,x,y,z,r) }
func OctahedronC (c []col.Colour, x, y, z, r float64) { octahedronC(c,x,y,z,r) }
func Circle (c col.Colour, x, y, z, r float64) { circle(c,x,y,z,r) }
func CircleSegment (c col.Colour, x, y, z, r, a, b float64) { circleSegment(c,x,y,z,r,a,b) }
func VertCircle (c col.Colour, x, y, z, r, a float64) { vertCircle(c,x,y,z,r,a) }
func Sphere (c col.Colour, x, y, z, r float64) { sphere(c,x,y,z,r) }
func Cone (c col.Colour, x, y, z, r, h float64) { cone(c,x,y,z,r,h) }
func DoubleCone (c col.Colour, x, y, z, r, h float64) { doubleCone(c,x,y,z,r,h) }
func CylinderC (c []col.Colour, x, y, z, r, h float64) { cylinderC(c,x,y,z,r,h) }
func Cylinder (c col.Colour, x, y, z, r, h float64) { cylinder(c,x,y,z,r,h) }
func CylinderSegment (c col.Colour, x, y, z, r, h, a, b float64) {
                      cylinderSegment (c,x,y,z,r,h,a,b) }
func HorCylinder (c col.Colour, x, y, z, r, l, a float64) { horCylinder(c,x,y,z,r,l,a) }
func HorCylinderC (c []col.Colour, x, y, z, r, l, a float64) { horCylinderC(c,x,y,z,r,l,a) }
func Torus (c col.Colour, x, y, z, R, r float64) { torus(c,x,y,z,R,r) }
func VerTorus (c col.Colour, x, y, z, R, r, a float64) { verTorus(c,x,y,z,R,r,a) }
func Paraboloid (c col.Colour, x, y, z, a, wx, wy float64) { paraboloid(c,x,y,z,a,wx,wy) }
func Surface (c col.Colour, f Fxy2z, wx, wy float64) { surface(c,f,wx,wy) }
