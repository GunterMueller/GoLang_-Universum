package col

// (c) Christian Maurer   v. 230326 - license see µU.go

import (
  . "µU/obj"
  "µU/rand"
  "µU/str"
)
type
  colour struct {
        r, g, b byte
                string
                }

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
  c.string = str.Lat1(n)
  return c
}

func (c *colour) imp (Y any) *colour {
  y, ok := Y.(*colour)
  if ! ok {
    TypeNotEqPanic (c, Y)
  }
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
  return c.Eq (black())
}

func (c *colour) IsWhite() bool {
  return c.Eq (white())
}

func (c *colour) IsFlashWhite() bool {
  return c.Eq (flashWhite())
}

func (c *colour) Empty() bool {
  return c.IsBlack()
}

func (c *colour) Clr() {
  c.r, c.g, c.b = 0, 0, 0
}

func (c *colour) Eq (Y any) bool {
  y := c.imp(Y)
  return c.r == y.r && c.g == y.g && c.b == y.b
}

func (c *colour) Less (Y any) bool {
  return false
}

func (c *colour) Leq (Y any) bool {
  return false
}

func (c *colour) Copy (Y any) {
  y := c.imp(Y)
  c.string = y.string
  c.r, c.g, c.b = y.r, y.g, y.b
}

func (c *colour) Clone() any {
  y := new_()
  y.Copy (c)
  return y
}

func (c *colour) Float32() (float32, float32, float32) {
  const f = float32(255)
  return float32(c.r) / f, float32(c.g) / f, float32(c.b) / f
}

func (c *colour) Float64() (float64, float64, float64) {
  const f = 255
  return float64(c.r) / f, float64(c.g) / f, float64(c.b) / f
}

func random() Colour {
  y := new_().(*colour)
  y.r, y.g, y.b = byte(rand.Natural(256)), byte(rand.Natural(256)), byte(rand.Natural(256))
  return y
}

func startCols() (Colour, Colour) {
  return FlashWhite(), Black()
}

func startColF() Colour {
  return FlashWhite()
}

func startColB() Colour {
  return FlashWhite()
}

func startColsA() (Colour, Colour) {
  return Red(), Black()
}

func (c *colour) Invert() {
  c.r, c.g, c.b = 255 - c.r, 255 - c.g, 255 - c.b
}

func (c *colour) Contrast() {
  const lightlimit = 352 // 320 352 384 416 448 480 512 <-- difficult problem,
                         // highly dependent of the intensity of green,
                         // and our eyes are particularly sensible for green !
  if c.g > 224 {
    c = Black().(*colour)
  } else if int(c.r) + int(c.g) + int(c.b) < lightlimit {
    c = FlashWhite().(*colour)
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
  c.r, c.g, c.b = s[0], s[1], s[2]
}

func (c *colour) Code() uint {
  return (uint(c.r) << 8 + uint(c.g)) << 8 + uint(c.b)
}

func headF() Colour { return flashWhite() }
func headB() Colour { return blue() }
func hintF() Colour { return flashWhite() }
func hintB() Colour { return magenta() }
func errorF() Colour { return flashYellow() }
func errorB() Colour { return red() }
func menuF() Colour { return flashWhite() }
func menuB() Colour { return red() }

func black() Colour            { return new3n ("Black",             0,   0,   0) }

func flashBrown() Colour       { return new3n ("FlashBrown",      102,  53,   0) }
func darkerBrown() Colour      { return new3n ("DarkerBrown",      34,  17,   0) }
func darkBrown() Colour        { return new3n ("DarkBrown",        51,  26,   0) }
func brown() Colour            { return new3n ("Brown",            85,  42,   0) }
func lightBrown() Colour       { return new3n ("LightBrown",      136,  68,  34) }
func lighterBrown() Colour     { return new3n ("LighterBrown",    238, 119,  68) }

func siena() Colour            { return new3n ("Siena",           153,  51,   0) }
func lightSiena() Colour       { return new3n ("LightSiena",      187,  68,  53) }
func redBrown() Colour         { return new3n ("RedBrown",        170,  64,  64) }
func oliveBrown() Colour       { return new3n ("OliveBrown",      127, 127,   0) }
func lightOliveBrown() Colour  { return new3n ("LightOliveBrown", 170, 170,  85) }
func umber() Colour            { return new3n ("Umber",           149, 135,   0) }

func darkOchre() Colour        { return new3n ("DarkOchre",       170, 127,  21) }
func ochre() Colour            { return new3n ("Ochre",           255, 170,  64) }
func lightOchre() Colour       { return new3n ("LightOchre",      255, 191, 106) }
func roseBrown() Colour        { return new3n ("RoseBrown",       255, 191, 149) }
func lightBeige() Colour       { return new3n ("LightBeige",      234, 212, 170) }
func beige1() Colour           { return new3n ("Beige1",          212, 191, 149) }

func flashRed() Colour         { return new3n ("FlashRed",        255,   0,   0) }
func darkerRed() Colour        { return new3n ("DarkerRed",        85,   0,   0) }
func darkRed() Colour          { return new3n ("DarkRed",         119,   0,   0) }
func red() Colour              { return new3n ("Red",             153,   0,   0) }
func lightRed() Colour         { return new3n ("LightRed",        204,  51,  51) }
func lighterRed() Colour       { return new3n ("LighterRed",      255, 102, 102) }

func darkRose() Colour         { return new3n ("DarkRose",        170,   0,  85) }
func rose1() Colour            { return new3n ("Rose1",           221,   0, 102) } // Rose: RAL
func lightRose() Colour        { return new3n ("LightRose",       255,  85, 136) }

func pompejiRed() Colour       { return new3n ("PompejiRed",      187,  68,  68) }
func cinnabarRed() Colour      { return new3n ("CinnabarRed",     238,  64,   0) }
func carmine() Colour          { return new3n ("Carmine",         125,   0,  42) }
func brickRed() Colour         { return new3n ("BrickRed",        205,  63,  51) }

func flashOrange() Colour      { return new3n ("FlashOrange",     255, 127,   0) }
func darkerOrange() Colour     { return new3n ("DarkerOrange",    221, 127,  68) }
func darkOrange() Colour       { return new3n ("DarkOrange",      221, 127,  68) }
func orange() Colour           { return new3n ("Orange",          255, 153,  51) }
func lightOrange() Colour      { return new3n ("LightOrange",     255, 164,  31) }
func lighterOrange() Colour    { return new3n ("LighterOrange",   255, 170,   0) }
func bloodOrange() Colour      { return new3n ("BloodOrange",     255, 112,  85) }

func flashYellow() Colour      { return new3n ("FlashYellow",     255, 255,   0) }
func darkerYellow() Colour     { return new3n ("DarkerYellow",    255, 187,   0) }
func darkYellow() Colour       { return new3n ("DarkYellow",      255, 187,   0) }
func yellow() Colour           { return new3n ("Yellow",          255, 255,  34) }
func lightYellow() Colour      { return new3n ("LightYellow",     255, 255, 102) }
func lighterYellow() Colour    { return new3n ("LighterYellow",   255, 255, 153) }
func sandYellow() Colour       { return new3n ("SandYellow",      234, 206, 127) }
func lemonYellow() Colour      { return new3n ("LemonYellow",     192, 255,  85) }

func flashGreen() Colour       { return new3n ("FlashGreen",        0, 255,   0) }
func darkerGreen() Colour      { return new3n ("DarkerGreen",       0,  51,   0) }
func darkGreen() Colour        { return new3n ("DarkGreen",         0,  85,   0) }
func green() Colour            { return new3n ("Green",             0, 170,   0) }
func lightGreen() Colour       { return new3n ("LightGreen",       51, 204,  51) }
func lighterGreen() Colour     { return new3n ("LighterGreen",    102, 255, 102) }

func birchGreen() Colour       { return new3n ("BirchGreen",       42, 153,  42) }
func grassGreen() Colour       { return new3n ("GrassGreen",        0, 144,   0) }
func oliveGreen() Colour       { return new3n ("OliveGreen",       61,  69,  46) }
func lightOliveGreen() Colour  { return new3n ("LightOliveGreen", 170, 196,  85) }
func yellowGreen() Colour      { return new3n ("YellowGreen",     170, 255,  85) }
func meadowGreen() Colour      { return new3n ("MeadowGreen",     106, 212, 106) }

func flashCyan() Colour        { return new3n ("FlashCyan",         0, 255, 255) }
func darkerCyan() Colour       { return new3n ("DarkerCyan",        0,  51,  51) }
func darkCyan() Colour         { return new3n ("DarkCyan",          0,  85,  85) }
func cyan() Colour             { return new3n ("Cyan",              0, 170, 170) }
func lightCyan() Colour        { return new3n ("LightCyan",        85, 255, 255) }
func lighterCyan() Colour      { return new3n ("LighterCyan",     153, 255, 255) }

func flashBlue() Colour        { return new3n ("FlashBlue",         0,   0, 255) }
func darkerBlue() Colour       { return new3n ("DarkerBlue",        0,   0,  51) }
func darkBlue() Colour         { return new3n ("DarkBlue",          0,   0,  85) }
func blue() Colour             { return new3n ("Blue",              0,   0, 170) }
func lightBlue() Colour        { return new3n ("LightBlue",        51,  51, 204) }
func lighterBlue() Colour      { return new3n ("LighterBlue",     102, 102, 255) }
func prussianBlue() Colour     { return new3n ("PrussianBlue",      0,  85, 136) }
func skyLightBlue() Colour     { return new3n ("SkyLightBlue",     85, 170, 255) }
func skyBlue() Colour          { return new3n ("SkyBlue",           0, 170, 255) }
func gentianBlue() Colour      { return new3n ("GentianBlue",       0,   0, 204) }
func ultramarine() Colour      { return new3n ("Ultramarine",      68,   0, 153) }

func flashMagenta() Colour     { return new3n ("FlashMagenta",    255,   0, 255) }
func darkerMagenta() Colour    { return new3n ("DarkerMagenta",    51,   0,  51) }
func darkMagenta() Colour      { return new3n ("DarkMagenta",      85,   0,  85) }
func magenta() Colour          { return new3n ("Magenta",         170,   0, 170) }
func lightMagenta() Colour     { return new3n ("LightMagenta",    255,  85, 255) }
func lighterMagenta() Colour   { return new3n ("LighterMagenta",  255, 153, 255) }

func pink() Colour             { return new3n ("Pink",            255,   0, 170) }
func deepPink() Colour         { return new3n ("DeepPink",        255,  17,  51) }

func flashGray() Colour        { return new3n ("FlashGray",       127, 127, 127) }
func darkerGray() Colour       { return new3n ("DarkerGray",       34,  34,  34) }
func darkGray() Colour         { return new3n ("DarkGray",         51,  51,  51) }
func gray() Colour             { return new3n ("Gray",             85,  85,  85) }
func lightGray() Colour        { return new3n ("LightGray",       136, 136, 136) }
func lighterGray() Colour      { return new3n ("LighterGray",     204, 204, 204) }

func silver() Colour           { return new3n ("Silver",          212, 212, 212) }

func flashWhite() Colour       { return new3n ("FlashWhite",      255, 255, 255) }
func white() Colour            { return new3n ("White",           170, 170, 170) }

func redWhite() Colour         { return new3n ("RedWhite",        255, 170, 170) }
func orangeWhite() Colour      { return new3n ("OrangeWhite",     255, 204, 170) }
func yellowWhite() Colour      { return new3n ("YellowWhite",     255, 255, 170) }
func greenWhite() Colour       { return new3n ("GreenWhite",      170, 255, 170) }
func cyanWhite() Colour        { return new3n ("CyanWhite",       170, 255, 255) }
func blueWhite() Colour        { return new3n ("BlueWhite",       170, 170, 255) }
func magentaWhite() Colour     { return new3n ("MagentaWhite",    255, 170, 255) }

// RAL-Farben see ral.go
