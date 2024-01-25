package col

// (c) Christian Maurer   v. 230326 - license see nU.go

import
  . "nU/obj"
type
  Colour interface {

  Object

// Liefert den rot/grün/blau-Anteil von x.
  R() byte; G() byte; B() byte

// x hat den Anteil b für rot/grün/blau.
  SetR (b byte); SetG (b byte); SetB (b byte)
}

func New() Colour { return new_() } // Black
func New3 (r, g, b byte) Colour { return new3(r,g,b) }

func Black() Colour          { return new3 (  0,   0,   0) }
func DarkRed() Colour        { return new3 ( 85,   0,   0) }
func Red() Colour            { return new3 (170,   0,   0) }
func LightRed() Colour       { return new3 (255,  85,  85) }
func Brown() Colour          { return new3 ( 95,  53,  34) }
func LightBrown() Colour     { return new3 (160,  88,  63) }
func Orange() Colour         { return new3 (255, 153,  51) }
func DarkYellow() Colour     { return new3 (255, 187,   0) }
func Yellow() Colour         { return new3 (255, 255,  85) }
func SandYellow() Colour     { return new3 (234, 206, 127) }
func DarkGreen() Colour      { return new3 (  0,  85,   0) }
func Green() Colour          { return new3 (  0, 170,   0) }
func LightGreen() Colour     { return new3 ( 85, 255,  85) }
func Cyan() Colour           { return new3 (  0, 170, 170) }
func LightCyan() Colour      { return new3 ( 85, 255, 255) }
func DarkBlue() Colour       { return new3 (  0,   0,  85) }
func Blue() Colour           { return new3 (  0,   0, 170) }
func LightBlue() Colour      { return new3 ( 85,  85, 255) }
func DarkMagenta() Colour    { return new3 ( 85,   0,  85) }
func Magenta() Colour        { return new3 (170,   0, 170) }
func LightMagenta() Colour   { return new3 (255,  85, 255) }
func Gray() Colour           { return new3 ( 85,  85,  85) }
func White() Colour          { return new3 (170, 170, 170) }
func FlashWhite() Colour     { return new3 (255, 255, 255) }
