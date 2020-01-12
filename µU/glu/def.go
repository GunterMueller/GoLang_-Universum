package glu

// (c) Christian Maurer   v. 191012 - license see µU.go

// a = angle of field of view in y-z-plane (<= 180.0)
func Perspective (a, p, n, f float64) { perspective(a,p,n,f) }

// (x, y, z) = eye, (x1, y1, z1) = center, (x2, y2, z2) = up
func LookAt (x, y, z, x1, y1, z1, x2, y2, z2 float64) { lookAt(x,y,z,x1,y1,z1,x2,y2,z2) }
