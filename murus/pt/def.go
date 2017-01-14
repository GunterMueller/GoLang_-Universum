package pt

// (c) murus.org  v. 140615 - license see murus.go

import (
  . "murus/obj"
  "murus/col"
  "murus/vect"
)
type
  Class byte; const (
  None = Class(iota)
  Start; Light
  Points; Lines; LineStrip; LineLoop
  Triangles; TriangleStrip; TriangleFan
  Quads; QuadStrip
  Polygon
  NClasses
)
type // Coloured points in 3-space with a class, a current number and a normal vector.
  Point interface {

// Returns an new point of class None with number 1, coordinates (0, 0, 0),
// normal (0, 0, 0) in the colour scr.StdColourF.
// New ()

  Object

//  Fin()

// x is the endpoint of v with class c, number a, colour f and normal n.
  Set (c Class, a uint, f col.Colour, v, n vect.Vector)

// Returns the class of x.
  Class() Class

// Returns the current number of x.
  Number () uint

// Returns the colour of x.
  Colour () col.Colour

// Returns the vector with the endpoint x.
  Read () vect.Vector

// Returns the vector with the endpoint x and the normal of x.
  Read2 () (vect.Vector, vect.Vector)
}
