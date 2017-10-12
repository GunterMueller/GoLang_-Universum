package col

// (c) Christian Maurer   v. 170919 - license see µu.go

import (
  . "µu/obj"
  "strconv"
  "µu/ker"
  "µu/rand"
)
const (
  m = 1 << 8
  m1 = m - 1
  light = byte(m1)
  codelen = 3
)
type
  colour struct {
        r, g, b byte
                }
var (
  bitDepth uint
  depth uint
)

func new_() Colour {
  x := new(colour)
  x.r, x.g, x.b = 0, 0, 0
  return x
}

func new3 (r, g, b byte) Colour {
  x := new(colour)
  x.r, x.g, x.b = r, g, b
  return x
}

//func colour3 (r, g, b uint) Colour {
//  var c Colour
//  c.R, c.G, c.B = byte(r % m), byte(g % m), byte(b % m)
//  return c
//}

func (c *colour) imp (Y Any) *colour {
  y, ok := Y.(*colour)
  if ! ok { TypeNotEqPanic(c, Y) }
  return y
}

func (c *colour) R() byte {
  return c.r
}

func (c *colour) G() byte {
  return c.g
}

func (c *colour) B() byte {
  return c.b
}

func (c *colour) SetR (b byte) {
  c.r = b
}

func (c *colour) SetG (b byte) {
  c.g = b
}

func (c *colour) SetB (b byte) {
  c.b = b
}

func (c *colour) Cc() Stream {
  n, bs := c.Code(), make(Stream, Depth())
  for i := uint(0); i < depth; i++ {
    bs[i] = byte(n)
    n >>= 8
  }
  return bs
}

func (c *colour) IsBlack() bool {
  return c.r == 0 &&
         c.g == 0 &&
         c.b == 0
}

func (c *colour) IsWhite() bool {
  return c.r == 170 &&
         c.g == 170 &&
         c.b == 170
}

func (c *colour) IsLightWhite() bool {
  return c.r == light &&
         c.g == light &&
         c.b == light
}

func (c *colour) Empty() bool {
  return c.r == 0 &&
         c.g == 0 &&
         c.b == 0
}

func (c *colour) Clr() {
  c.r, c.g, c.b = 0, 0, 0
}

func (c *colour) Eq (Y Any) bool {
  y := c.imp(Y)
  return c.r == y.r &&
         c.g == y.g &&
         c.b == y.b
}

func (c *colour) Less (Y Any) bool {
  return false
}

func (c *colour) Copy (Y Any) {
 y := c.imp(Y)
 c.r, c.g, c.b = y.r, y.g, y.b
}

func (c *colour) Clone() Any {
  y := new_()
  y.Copy (c)
  return y
}

func (c *colour) Real() (float32, float32, float32) {
  const f = float32(m1)
  return float32(c.r) / f, float32(c.g) / f, float32(c.b) / f
}

func (c *colour) Double() (float64, float64, float64) {
  const f = m1
  return float64(c.r) / f, float64(c.g) / f, float64(c.b) / f
}

func random() Colour {
  y := new_().(*colour)
  y.r, y.g, y.b = byte (rand.Natural (m)), byte (rand.Natural (m)), byte (rand.Natural (m))
  return y
}

func startCols() (Colour, Colour) {
  return white, black
}

func startColsA() (Colour, Colour) {
  return red, black
}

func (c *colour) ansiEncode() uint { // 0..15 // doch vielleicht Mist
  const (black = 1 << iota / 2; red; green; blue; light)
  const m3 = m1 / 3
  const m2 = 2 * m3
  r, g, b := c.r, c.g, c.b
  n := uint(black)
  if r >= m2 { r -= m2; n += red }
  if g >= m2 { g -= m2; n += green }
  if b >= m2 { b -= m2; n += blue }
  if r >= m3 && g >= m3 && b >= m3 { n += light }
  return n
}

/*
func changeRand (c *Colour) {
  const (N = 32; N2 = N / 2)
  n := byte(rand.Natural (N))
  if n < N2 {
    if c.R >= n {
      c.R -= n
    } else {
      c.R = 0
    }
  } else { // n >= N2
    n -= N2
    if c.R + n <= m1 {
      c.R += n
    } else {
      c.R = light
    }
  }
  n = byte(rand.Natural (N))
  if n < N2 {
    if c.G >= n {
      c.G -= n
    } else {
      c.G = 0
    }
  } else {
    n -= N2
    if c.G + n <= m1 {
      c.G += n
    } else {
      c.G = light
    }
  }
  n = byte(rand.Natural (N))
  if n < N2 {
    if c.B >= n {
      c.B -= n
    } else {
      c.B = 0
    }
  } else {
    n -= N2
    if c.B + n <= m1 {
      c.B += n
    } else {
      c.B = light
    }
  }
}
*/

func (c *colour) Invert() {
  c.r, c.g, c.b = m1 - c.r, m1 - c.g, m1 - c.b
}

func (c *colour) Contrast() {
  const lightlimit = 352 // 320 352 384 416 448 480 512 <-- difficult problem,
                         // highly dependent of the intensity of green,
                         // and our eyes are particularly sensible for green !
  if c.g > 224 {
    c = black.(*colour)
  } else if int(c.r) + int(c.g) + int(c.b) < lightlimit {
    c = lightWhite.(*colour)
  } else {
    c = black.(*colour)
  }
}

func ok (b byte) bool {
  if b < '9' {
    return true
  } else if 'A' <= b && b <= 'F' {
    return true
  }
  return false
}

func value (b byte) uint{
  if b < '9' {
    return uint(b - '0')
  } else if 'A' <= b && b <= 'F' {
    return uint(b - 'A' + 10)
  }
  return 0
}

/*
func change (c *Colour, rgb, d byte, lighter bool) {
  if rgb > 2 || d > 127 { return }
  switch rgb { case 0:
    if lighter {
      if c.R <= m1 - d {
        c.R += d
      }
    } else if c.R >= d {
      c.R -= d
    }
  case 1:
    if lighter {
      if c.G <= m1 - d {
        c.G += d
      }
    } else if c.G >= d {

    }
  case 2:
    if lighter {
      if c.B <= m1 - d {
        c.B += d
      }
    } else if c.B >= d {
      c.B -= d
    }
  }
}
*/

func char (n uint) string {
  if n < 10 {
    return string (n + uint('0'))
  }
  if n < 16 {
    return string (n - 10 + uint('A'))
  }
  return string (0)
}

func (c *colour) String() string {
  return char (uint(c.r) / 16) + char (uint(c.r) % 16) +
         char (uint(c.g) / 16) + char (uint(c.g) % 16) +
         char (uint(c.b) / 16) + char (uint(c.b) % 16)
}

func (c *colour) Defined (s string) bool {
  if len(s) != 6 { return false }
  for i := 0; i < 6; i++ {
    if ! ok (s[i]) { return false }
  }
  c.r = byte(16 * value (s[0]) + value (s[1]))
  c.g = byte(16 * value (s[2]) + value (s[3]))
  c.b = byte(16 * value (s[4]) + value (s[5]))
  return true
}

func (c *colour) Codelen() uint {
  return 3
}

func (c *colour) Encode() Stream {
  b := make (Stream, 3)
  b[0], b[1], b[2] = c.r, c.g, c.b
  return b
}

func (c *colour) Decode (b Stream) {
  if len (b) == 3 {
    c.r, c.g, c.b = b[0], b[1], b[2]
  } else {
    c = lightWhite.(*colour)
  }
}

func setDepth (bits uint) {
  switch bits {
  case 4, 8, 15, 16, 24, 32:
    bitDepth = bits
    depth = (bitDepth + 4) / 8
  default:
    ker.Panic ("strange colourdepth: " + strconv.Itoa(int(bits)) + " bits")
  }
}

func nCols() uint{
  switch bitDepth {
  case 4:
    return 16
  case 8:
    return m
  case 15:
    return 128 * m
  case 16:
    return m * m
  case 24, 32:
    return m * m * m
  }
  return 0
}

func (c *colour) Code() uint {
  switch bitDepth {
  case 4:
    return c.ansiEncode()
  case 8:
    return ((uint(c.r) >> 5) << 2 + uint(c.g) >> 5) << 3 + uint(c.b) >> 6 // direct colour
  case 15:
    return ((uint(c.r) >> 3) << 5 + uint(c.g) >> 3) << 5 + uint(c.b) >> 3
  case 16:
    return ((uint(c.r) >> 3) << 5 + uint(c.g) >> 2) << 6 + uint(c.b) >> 3
  case 24:
    return (uint(c.r) << 8 + uint(c.g)) << 8 + uint(c.b)
  case 32:
    return (uint(c.r) << 8 + uint(c.g)) << 8 + uint(c.b)
//    return m * (m * (uint(c.alpha) << 8 + (uint(c.r) << 8 + uint(c.g)) << 8 + uint(c.b)
  }
  return 0
}

// Pre: len(bs) == int(depth)
func cd (bs Stream) uint { // inverse of cc
  n := uint(0)
  if len(bs) == int(depth) {
    for i := int(depth) - 1; i >= 0; i-- {
      n = n * 1<<8 + uint(bs[i])
    }
  }
  return n
}

var (
  headF = lightWhite
  headB = blue
  hintF = lightWhite
  hintB = magenta
  errorF = flashYellow
  errorB = red
  menuF = lightWhite
  menuB = red
  muF =           new3 (  0,  16,  64)
  muB =           new3 (231, 238, 255)

  black =            new3 (  0,   0,   0)

  brown =            new3 ( 95,  53,  34)
  blackBrown =       new3 ( 30,  16,  12)
  darkBrown =        new3 ( 60,  33,  24)
  mediumBrown =      new3 (149, 106,   0)
  lightBrown =       new3 (160,  88,  63)
  whiteBrown =       new3 (221, 153, 106)
//  whiteBrown =       new3 (255, 212, 149)
  brownWhite =       new3 (249, 202, 160)

  siena =            new3 (153,  85,  42)
  lightSiena =       new3 (191, 127,  42)
//  redBrown =        new3 (170,  64,  64)
//  umbrabraun =       new3 (149, 135,   0)
  oliveBrown =       new3 (127, 127,   0)
  lightOliveBrown =  new3 (170, 170,  85)
//  orangeBrown1 =     new3 (127, 106,  42)
//  Dark Ocker =      new3 (170, 127,  21)
//  Ocker =            new3 (255, 170,  64)
//  Light Ocker =        new3 (255, 191, 106)
//  Rosabraun =        new3 (255, 191, 149)
//  Hellbeige =        new3 (234, 212, 170)
//  Beige2 =           new3 (212, 191, 149)
//  VeryLightBrown =    new3 (206, 170, 127)

  blackRed =         new3 ( 46,  18,  26)
  darkRed =          new3 ( 85,   0,   0)
  red =              new3 (170,   0,   0)
  flashRed =         new3 (255,   0,   0)
  lightRed =         new3 (255,  85,  85)
  whiteRed =         new3 (255, 187, 170)
//  dunkelrosa =       new3 (234,  0,  127)
//  rosa =             new3 (255, 170, 170)
//  hellrosa =         new3 (255, 191, 191)
  pompejiRed =       new3 (187,  68,  68)
  cinnabarRed =      new3 (238,  64,   0)
  carmine =          new3 (125,   0,  42)
  brickRed =         new3 (205,  63,  51)

  flashOrange =      new3 (255, 127,   0)
  darkOrange =       new3 (221, 127,  68)
  orange =           new3 (255, 153,  51)
  lightOrange =      new3 (255, 164,  31)
  whiteOrange =      new3 (255, 170,   0)
//  blutOrange1 =      new3 (255, 112,  85)

  flashYellow =      new3 (255, 255,   0)
  darkYellow =       new3 (255, 187,   0)
  yellow =           new3 (255, 255,  34)
  lightYellow =      new3 (255, 255, 102)
  whiteYellow =      new3 (255, 255, 153)
  sandGelb1 =        new3 (234, 206, 127)
  zitronenGelb1 =    new3 (191, 255,  85)

  flashGreen =       new3 (  0, 255,   0)
  blackGreen =       new3 (  0,  51,   0)
  darkGreen =        new3 (  0,  85,   0)
  green =            new3 (  0, 170,   0)
  lightGreen =       new3 ( 85, 255,  85)
  whiteGreen =       new3 (170, 255, 170)
  birchGreen =       new3 ( 42, 153,  42)
  grassGreen =       new3 (  0, 144,   0)
//  chromeGreen =    new3 ( 85, 170,   0)
//  lightChromeGreen =    new3 ( 85, 170,   0)
  oliveGreen =       new3 ( 85, 170,   0)
//  lightOliveGreen =     new3 (170, 196,  85)
  yellowGreen =         new3 (170, 255,  85)
//  wiesenGreen =       new3 (106, 212, 106)

  blackCyan =        new3 (  0,  51,  51)
  darkCyan =         new3 (  0,  85,  85)
  cyan =             new3 (  0, 170, 170)
  lightCyan =        new3 ( 85, 255, 255)
  whiteCyan =        new3 (170, 255, 255)
  flashCyan =        new3 (  0, 255, 255)

  flashBlue =        new3 (  0,   0, 255)
  blackBlue =        new3 (  0,   0,  51)
  prussianBlue =     new3 (  0, 102, 170)
  darkBlue =         new3 (  0,   0,  85)
  blue =             new3 (  0,   0, 170)
  lightBlue =        new3 ( 51, 102, 255)
  whiteBlue =        new3 (153, 221, 255)
//  whiteBlue =        new3 (170, 170, 255)
  gentianBlue =      new3 (  0,   0, 212)
  skyBlue =          new3 (  0, 170, 255)
  ultramarine =      new3 ( 68,   0, 153)

  blackMagenta =     new3 ( 51,  0,  51)
  darkMagenta =      new3 ( 85,  0,  85)
  magenta =          new3 (170,  0, 170)
  lightMagenta =     new3 (255, 85, 255)
  flashMagenta =     new3 (255,  0, 255)
  whiteMagenta =     new3 (255,187, 255)
  pink =             new3 (255,   0, 170)
  deepPink =         new3 (255,  17,  51)

  blackGray =        new3 ( 34,  34,  34)
  darkGray =         new3 ( 51,  51,  51)
  gray =             new3 ( 85,  85,  85)
  lightGray =        new3 (136, 136, 136)
  whiteGray =        new3 (204, 204, 204)
  silver =           new3 (212, 212, 212)

  white =            new3 (170, 170, 170)
  lightWhite =       new3 (255, 255, 255)
)
