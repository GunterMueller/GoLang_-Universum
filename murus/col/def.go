package col

// (c) murus.org  v. 170813 - license see murus.go

import
  "C"
const
  P6 = 3
type
  Colour struct {
        R, G, B byte
                }

func StartCols() (Colour, Colour) { return startCols() }
func StartColsA() (Colour, Colour) { return startColsA() }

// Returns the colour defined by (r, g, b).
func Colour3 (r, g, b uint) Colour { return colour3(r,g,b) }

// Returns the rgb-values of c scaled to the range from 0 to 1.
func Float (c Colour) (float32, float32, float32) { return float(c) }
func LongFloat (c Colour) (float64, float64, float64) { return longFloat(c) }

// Returns a random colour.
func ColourRand () Colour { return colourRand() }

// Returns true, if c and c1 coincide in their rgb-values.
func Eq (c, c1 Colour) bool { return eq(c,c1) }

// Returns true, if c is, what the name of the func says.
func IsBlack (c Colour) bool { return isBlack(c) }
func IsLightwhite (c Colour) bool { return isLightWhite(c) }

// c is changed in a manner suggested by the name of the func.
func Invert (c *Colour) { invert(c) }
func Contrast (c *Colour) { contrast(c) }

// Returns true, iff s is a string of 3 values in sedecimal basis
// (with uppercase letters). In that case, c is the colour with
// the corresponding rgb-values; otherwise, nothing has happened.
func Set (c *Colour, s string) bool { return set(c,s) }

// Returns "rrggbb", where "rr", "gg" and "bb" are the rgb-values
// in sedecimal basis (with uppercase letters).
func String (c Colour) string { return string_(c) }

// TODO Spec
func Change (c *Colour, rgb, d byte, l bool) { change(c,rgb,d,l) }

// see murus/obj/coder.go
func Codelen () uint{ return codelen }
func Encode (c Colour) []byte { return encode(c) }
func Decode (c *Colour, b []byte) { decode(c,b) }

// Pre: b is one of 4, 8, 15, 16, 24 or 32.
// depth == (b + 4) / 8.
func SetDepth (b uint) { setDepth(b) }

// Returns the number of available colours, depending on depth.
func NCols() uint { return nCols() }

// TODO Spec
func Code (c Colour) uint { return code(c) }
func P6Encode (A, P []byte) { p6Encode(A,P) }
func P6Colour (A []byte) Colour { return p6Colour(A) }
func Cc (c Colour) []byte { return cc(c) }
// func Cd (bs []byte) uint { return cd(bs) }

var (
  Depth uint // in bytes - must not be altered after the call to SetDepth !
  HeadF, HeadB, // for headings
  HintF, HintB, ErrorF, ErrorB, // for hints and error reports
  MenuF, MenuB, // for menues
  MurusF, MurusB Colour

  Black =            colour3 (  0,   0,   0)

  Brown =            colour3 ( 95,  53,  34)
  BlackBrown =       colour3 ( 30,  16,  12)
  DarkBrown =        colour3 ( 60,  33,  24)
  MediumBrown =      colour3 (149, 106,   0)
  LightBrown =       colour3 (160,  88,  63)
  WhiteBrown =       colour3 (221, 153, 106)
//  WhiteBrown =       colour3 (255, 212, 149)
  BrownWhite =       colour3 (249, 202, 160)

  Siena =            colour3 (153,  85,  42)
  LightSiena =       colour3 (191, 127,  42)
//  RedBrown =        colour3 (170,  64,  64)
//  Umbrabraun =       colour3 (149, 135,   0)
  OliveBrown =       colour3 (127, 127,   0)
  LightOliveBrown =  colour3 (170, 170,  85)
//  OrangeBrown1 =     colour3 (127, 106,  42)
//  Dark Ocker =      colour3 (170, 127,  21)
//  Ocker =            colour3 (255, 170,  64)
//  Light Ocker =        colour3 (255, 191, 106)
//  Rosabraun =        colour3 (255, 191, 149)
//  Hellbeige =        colour3 (234, 212, 170)
//  Beige2 =           colour3 (212, 191, 149)
//  VeryLightBrown =    colour3 (206, 170, 127)

  BlackRed =         colour3 ( 46,  18,  26)
  DarkRed =          colour3 ( 85,   0,   0)
  Red =              colour3 (170,   0,   0)
  FlashRed =         colour3 (255,   0,   0)
  LightRed =         colour3 (255,  85,  85)
  WhiteRed =         colour3 (255, 187, 170)
//  Dunkelrosa =       colour3 (234,  0,  127)
//  Rosa =             colour3 (255, 170, 170)
//  Hellrosa =         colour3 (255, 191, 191)
  PompejiRed =       colour3 (187,  68,  68)
  CinnabarRed =      colour3 (238,  64,   0)
  Carmine =          colour3 (125,   0,  42)
  BrickRed =         colour3 (205,  63,  51)

  FlashOrange =      colour3 (255, 127,   0)
  DarkOrange =       colour3 (221, 127,  68)
  Orange =           colour3 (255, 153,  51)
  LightOrange =      colour3 (255, 164,  31)
  WhiteOrange =      colour3 (255, 170,   0)
//  BlutOrange1 =      colour3 (255, 112,  85)

  FlashYellow =      colour3 (255, 255,   0)
  DarkYellow =       colour3 (255, 187,   0)
  Yellow =           colour3 (255, 255,  34)
  LightYellow =      colour3 (255, 255, 102)
  WhiteYellow =      colour3 (255, 255, 153)
  Sandgelb1 =        colour3 (234, 206, 127)
  Zitronengelb1 =    colour3 (191, 255,  85)

  FlashGreen =       colour3 (  0, 255,   0)
  BlackGreen =       colour3 (  0,  51,   0)
  DarkGreen =        colour3 (  0,  85,   0)
  Green =            colour3 (  0, 170,   0)
  LightGreen =       colour3 ( 85, 255,  85)
  WhiteGreen =       colour3 (170, 255, 170)
  BirchGreen =       colour3 ( 42, 153,  42)
  GrassGreen =       colour3 (  0, 144,   0)
//  ChromeGreen =    colour3 ( 85, 170,   0)
//  LightChromeGreen =    colour3 ( 85, 170,   0)
  OliveGreen =       colour3 ( 85, 170,   0)
//  LightOliveGreen =     colour3 (170, 196,  85)
  YellowGreen =         colour3 (170, 255,  85)
//  WiesenGreen =       colour3 (106, 212, 106)

  BlackCyan =        colour3 (  0,  51,  51)
  DarkCyan =         colour3 (  0,  85,  85)
  Cyan =             colour3 (  0, 170, 170)
  LightCyan =        colour3 ( 85, 255, 255)
  WhiteCyan =        colour3 (170, 255, 255)
  FlashCyan =        colour3 (  0, 255, 255)

  FlashBlue =        colour3 (  0,   0, 255)
  BlackBlue =        colour3 (  0,   0,  51)
  PrussianBlue =     colour3 (  0, 102, 170)
  DarkBlue =         colour3 (  0,   0,  85)
  Blue =             colour3 (  0,   0, 170)
  LightBlue =        colour3 ( 51, 102, 255)
  WhiteBlue =        colour3 (153, 221, 255)
//  WhiteBlue =        colour3 (170, 170, 255)
  GentianBlue =      colour3 (  0,   0, 212)
  SkyBlue =          colour3 (  0, 170, 255)
  Ultramarine =      colour3 ( 68,   0, 153)

  BlackMagenta =     colour3 ( 51,  0,  51)
  DarkMagenta =      colour3 ( 85,  0,  85)
  Magenta =          colour3 (170,  0, 170)
  LightMagenta =     colour3 (255, 85, 255)
  FlashMagenta =     colour3 (255,  0, 255)
  WhiteMagenta =     colour3 (255,187, 255)
  Pink =             colour3 (255,   0, 170)
  DeepPink =         colour3 (255,  17,  51)

  BlackGray =        colour3 ( 34,  34,  34)
  DarkGray =         colour3 ( 51,  51,  51)
  Gray =             colour3 ( 85,  85,  85)
  LightGray =        colour3 (136, 136, 136)
  WhiteGray =        colour3 (204, 204, 204)
  Silver =           colour3 (212, 212, 212)

  White =            colour3 (170, 170, 170)
  LightWhite =       colour3 (255, 255, 255)
)
