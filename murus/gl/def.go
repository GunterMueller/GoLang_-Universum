package gl

// (c) murus.org  v. 170820 - license see murus.go

import (
  "murus/col"
  "murus/vect"
)
type
  Class byte; const (
//  UNDEF = iota; // -> Chaos
  POINTS = iota;
  LINES
  LINE_LOOP
  LINE_STRIP
  TRIANGLES
  TRIANGLE_STRIP
  TRIANGLE_FAN
  QUADS
  QUAD_STRIP
  POLYGON
  UNDEF
  LIGHT
)
const
  MaxL = 16 // <= GL.GL_MAX_LIGHTS

// TODO Spec
func Cls (c col.Colour) { cls(c) }

// TODO Spec // called by pts.Start()
func Init (f float64) { init_(f) }

// Pre: n < MaxL, 0 <= h[i] <= 1 für i = 0, 1.
// If Light n is already switched on, nothing has happened.
// Otherwise it is now switched on at position v
// in colour c with ambience h[0] and diffusion h[1].
func InitLight (n uint, v, h vect.Vector, c col.Colour) { initLight(n,v,h,c) }

// Pre: Light n is switched on.
// Light n has position v.
func PosLight (n uint, v vect.Vector) { posLight(n,v) }

// TODO Spec
func ActualizeLight (n uint) { actLight(n) } // n < MaxL

// TODO Spec
func ShowLight (on bool) { lightVis = on }

// TODO Spec
func Write0() { write0() }

// TODO Spec
func Write (class Class, a uint, vs, ns []vect.Vector, c col.Colour) { write(class,a,vs,ns,c) }

// TODO Spec
func Write1() { write1() }

// TODO Spec
func Actualize (r, v, o, a vect.Vector) { act(r,v,o,a) }
