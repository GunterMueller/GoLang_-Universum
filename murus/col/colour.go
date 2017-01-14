package col

// (c) murus.org  v. 150425 - license see murus.go

import (
  "strconv"
  "murus/ker"; "murus/rand"
)
const (
  m = 1 << 8
  m1 = m - 1
  light = byte(m1)
  codelen = 3
)
var
  bitDepth uint

func colour3 (r, g, b uint) Colour {
  var c Colour
  c.R, c.G, c.B = byte(r % m), byte(g % m), byte(b % m)
  return c
}

func ansiEncode (c Colour) uint { // 0..15 // doch vielleicht Mist
  const (black = 1 << iota / 2; red; green; blue; light)
  const m3 = m1 / 3
  const m2 = 2 * m3
  r, g, b:= c.R, c.G, c.B
  n:= uint(black)
  if r >= m2 { r -= m2; n += red }
  if g >= m2 { g -= m2; n += green }
  if b >= m2 { b -= m2; n += blue }
  if r >= m3 && g >= m3 && b >= m3 { n += light }
  return n
}

func float (c Colour) (float32, float32, float32) {
  const f = float32(m1)
  return float32(c.R) / f, float32(c.G) / f, float32(c.B) / f
}

func longFloat (c Colour) (float64, float64, float64) {
  const f = float64(m1)
  return float64(c.R) / f, float64(c.G) / f, float64(c.B) / f
}

func colourRand() Colour {
  var c Colour
  c.R, c.G, c.B = byte (rand.Natural (m)), byte (rand.Natural (m)), byte (rand.Natural (m))
  return c
}

func changeRand (c *Colour) {
  const (N = 32; N2 = N / 2)
  n:= byte(rand.Natural (N))
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

func eq (c, c1 Colour) bool {
  return c.R == c1.R &&
         c.G == c1.G &&
         c.B == c1.B
}

func isBlack (c Colour) bool {
  return c.R == 0 &&
         c.G == 0 &&
         c.B == 0
}

func isLightWhite (c Colour) bool {
  return c.R == light &&
         c.G == light &&
         c.B == light
}

func invert (c *Colour) {
  c.R, c.G, c.B = m1 - c.R, m1 - c.G, m1 - c.B
}

func contrast (c *Colour) {
  const lightlimit = 352 // 320 352 384 416 448 480 512 <-- difficult problem,
                         // highly dependent of the intensity of green,
                         // and our eyes are particularly sensible for green !
  if c.G > 224 {
    *c = Black
  } else if int(c.R) + int(c.G) + int(c.B) < lightlimit {
    *c = LightWhite
  } else {
    *c = Black
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

func set (c *Colour, s string) bool {
  if len (s) != 6 { return false }
  for i:= 0; i < 6; i++ {
    if ! ok (s[i]) { return false }
  }
  c.R = byte(16 * value (s[0]) + value (s[1]))
  c.G = byte(16 * value (s[2]) + value (s[3]))
  c.B = byte(16 * value (s[4]) + value (s[5]))
  return true
}

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
      c.G -= d
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

func char (n uint) string {
  if n < 10 {
    return string (n + uint('0'))
  }
  if n < 16 {
    return string (n - 10 + uint('A'))
  }
  return string (0)
}

func string_ (c Colour) string {
  return char (uint(c.R) / 16) + char (uint(c.R) % 16) +
         char (uint(c.G) / 16) + char (uint(c.G) % 16) +
         char (uint(c.B) / 16) + char (uint(c.B) % 16)
}

func encode (c Colour) []byte {
  b:= make ([]byte, 3)
  b[0], b[1], b[2] = c.R, c.G, c.B
  return b
}

func decode (c *Colour, b []byte) {
  if len (b) == 3 {
    c.R, c.G, c.B = b[0], b[1], b[2]
  } else {
    *c = LightWhite
  }
}

func setDepth (bits uint) {
  switch bits { case 4, 8, 15, 16, 24, 32:
    bitDepth = bits
    Depth = (bitDepth + 4) / 8
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

func code (c Colour) uint {
  switch bitDepth {
  case 4:
    return ansiEncode (c)
  case 8:
    return ((uint(c.R) >> 5) << 2 + uint(c.G) >> 5) << 3 + uint(c.B) >> 6 // direct colour
  case 15:
    return ((uint(c.R) >> 3) << 5 + uint(c.G) >> 3) << 5 + uint(c.B) >> 3
  case 16:
    return ((uint(c.R) >> 3) << 5 + uint(c.G) >> 2) << 6 + uint(c.B) >> 3
  case 24:
    return (uint(c.R) << 8 + uint(c.G)) << 8 + uint(c.B)
  case 32:
    return /* m * (m * (uint(c.alpha) << 8 + */ (uint(c.R) << 8 + uint(c.G)) << 8 + uint(c.B)
  }
  return 0
}

func cc (c Colour) []byte {
  n, bs:= Code(c), make([]byte, Depth)
  for i:= uint(0); i < Depth; i++ {
    bs[i] = byte(n)
    n >>= 8
  }
  return bs
}

/* // Pre: len(bs) == int(Depth)
func cd (bs[]byte) uint { // inverse of cc
  n:= uint(0)
  if len(bs) == int(Depth) {
    for i:= int(Depth) - 1; i >= 0; i-- {
      n = n * 1<<8 + uint(bs[i])
    }
  }
  return n
} */

func init() {
  HeadF, HeadB = LightWhite, Blue
  HintF, HintB = LightWhite, Magenta
  ErrorF, ErrorB = FlashYellow, Red
  MenuF, MenuB = LightWhite, Red
  MurusF, MurusB = colour3 (0, 16, 64), colour3 (231, 238, 255)
}
