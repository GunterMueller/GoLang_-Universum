package col

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  . "µU/obj"
  "µU/rand"
)
const (
  m = 1 << 8
  m1 = m - 1
  light = byte(m1)
  codelen = 3
//  dp = uint(4) // XXX
)
type
  colour struct {
        r, g, b byte
                string
                }
var (
  all = make([]uint, 0)
  name = make([]string, 0)
)

func new_() Colour {
  c := new(colour)
  c.r, c.g, c.b = 0, 0, 0
  return c
}

func new3 (r, g, b byte) Colour {
  c := new(colour)
  c.r, c.g, c.b = r, g, b
  c.string = c.String1()
  return c
}

func new3n (n string, r, g, b byte) Colour {
  c := new(colour)
  c.r, c.g, c.b = r, g, b
  if n == "" {
    n = c.String1()
  }
  c.string = n
  return c
}

func (c *colour) imp (Y any) *colour {
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

func (c *colour) Set (r, g, b byte) {
  c.r, c.g, c.b = r, g, b
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
  return c.IsBlack()
}

func (c *colour) Clr() {
  c.r, c.g, c.b = 0, 0, 0
}

func (c *colour) Eq (Y any) bool {
  y := c.imp(Y)
  return c.r == y.r &&
         c.g == y.g &&
         c.b == y.b
}

func (c *colour) Less (Y any) bool {
  return false
}

func (c *colour) Copy (Y any) {
  y := c.imp(Y)
  c.r, c.g, c.b = y.r, y.g, y.b
}

func (c *colour) Clone() any {
  y := new_()
  y.Copy (c)
  return y
}

func (c *colour) Float32() (float32, float32, float32) {
  const f = float32(m1)
  return float32(c.r) / f, float32(c.g) / f, float32(c.b) / f
}

func (c *colour) Float64() (float64, float64, float64) {
  const f = m1
  return float64(c.r) / f, float64(c.g) / f, float64(c.b) / f
}

func random() Colour {
  y := new_().(*colour)
  y.r, y.g, y.b = byte (rand.Natural (m)), byte (rand.Natural (m)), byte (rand.Natural (m))
  return y
}

func startCols() (Colour, Colour) {
  return White(), Black()
}

func startColsA() (Colour, Colour) {
  return Red(), Black()
}

func (c *colour) Invert() {
  c.r, c.g, c.b = m1 - c.r, m1 - c.g, m1 - c.b
}

func (c *colour) Contrast() {
  const lightlimit = 352 // 320 352 384 416 448 480 512 <-- difficult problem,
                         // highly dependent of the intensity of green,
                         // and our eyes are particularly sensible for green !
  if c.g > 224 {
    c = Black().(*colour)
  } else if int(c.r) + int(c.g) + int(c.b) < lightlimit {
    c = LightWhite().(*colour)
  } else {
    c = Black().(*colour)
  }
}

func ok (b byte) bool {
  if b < '9' {
    return true
  }
  if 'A' <= b && b <= 'F' {
    return true
  }
  return false
}

func value (b byte) uint{
  if b < '9' {
    return uint(b - '0')
  }
  if 'A' <= b && b <= 'F' {
    return uint(b - 'A' + 10)
  }
  return 0
}

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
  return c.string
}

func (c *colour) Defined (s string) bool {
  c.string = s
  return true
}

func (c *colour) String1() string {
  s := char (uint(c.r) / 16) + char (uint(c.r) % 16)
  s += char (uint(c.g) / 16) + char (uint(c.g) % 16)
  s += char (uint(c.b) / 16) + char (uint(c.b) % 16)
  return s
}

func (c *colour) Defined1 (s string) bool {
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
  return Stream {c.r, c.g, c.b}
}

func (c *colour) EncodeInv() Stream {
  return Stream {c.b, c.g, c.r}
}

func (c *colour) Decode (s Stream) {
  if len(s) == 3 {
    c.r, c.g, c.b = s[0], s[1], s[2]
  } else {
    c = LightWhite().(*colour)
  }
}

/*/
func depth() uint {
  return dp
}
/*/

func (c *colour) Code() uint {
  return (uint(c.r) << 8 + uint(c.g)) << 8 + uint(c.b)
}
