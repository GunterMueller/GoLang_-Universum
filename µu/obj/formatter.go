package obj

// (c) Christian Maurer   v. 150425 - license see µu.go

type
  Format byte
type
  Formatter interface {

// Pre: f < Nformats of the objects of the type of x.
// x has the format f.
  SetFormat (f Format)

// Returns the format of x.
  GetFormat() Format
}
