package col

// (c) Christian Maurer   v. 230401 - license see µU.go

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
  return c.Eq (black)
}

func (c *colour) IsWhite() bool {
  return c.Eq (white)
}

func (c *colour) IsFlashWhite() bool {
  return c.Eq (flashWhite)
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

func flashBrown() Colour       { return new3n ("FlashBrown",       102,  53,   0) }
func blackBrown() Colour       { return new3n ("BlackBrown",        34,  17,   0) }
func darkBrown() Colour        { return new3n ("DarkBrown",         51,  26,   0) }
func brown() Colour            { return new3n ("Brown",             85,  42,   0) }
func lightBrown() Colour       { return new3n ("LightBrown",       136,  68,  34) }
func whiteBrown() Colour       { return new3n ("WhiteBrown",       221, 119,  68) }
func darkOchre() Colour        { return new3n ("DarkOchre",        102,  61,  34) }
func ochre() Colour            { return new3n ("Ochre",            153,  85,  51) }
func lightOchre() Colour       { return new3n ("LightOchre",       170, 119,  85) }
func flashRed() Colour         { return new3n ("FlashRed",         255,   0,   0) }
func blackRed() Colour         { return new3n ("BlackRed",          85,   0,   0) }
func darkRed() Colour          { return new3n ("DarkRed",          119,   0,   0) }
func red() Colour              { return new3n ("Red",              153,   0,   0) }
func lightRed() Colour         { return new3n ("LightRed",         204,  51,  51) }
func whiteRed() Colour         { return new3n ("WhiteRed",         255, 102, 102) }
func pompejiRed() Colour       { return new3n ("PompejiRed",       136,  53,  53) }
func cinnabarRed() Colour      { return new3n ("CinnabarRed",      238,  64,   0) }
func carmine() Colour          { return new3n ("Carmine",          125,   0,  42) }
func brickRed() Colour         { return new3n ("BrickRed",         205,  63,  51) }
func siena() Colour            { return new3n ("Siena",            153,  51,   0) }
func lightSiena() Colour       { return new3n ("LightSiena",       187,  68,  53) }
func darkRose() Colour         { return new3n ("DarkRose",         170,   0,  85) }
func lightRose() Colour        { return new3n ("LightRose",        187,  54, 102) }
func whiteRose() Colour        { return new3n ("WhiteRose",        204, 119, 136) }
func flashOrange() Colour      { return new3n ("FlashOrange",      255,  51,  51) }
func blackOrange() Colour      { return new3n ("BlackOrange",      136,  34,  17) }
func darkOrange() Colour       { return new3n ("DarkOrange",       185,  34,  17) }
func orange() Colour           { return new3n ("Orange",           255,  51,  34) }
func lightOrange() Colour      { return new3n ("LightOrange",      255,  85,  34) }
func whiteOrange() Colour      { return new3n ("WhiteOrange",      255, 136,  68) }

func flashYellow() Colour      { return new3n ("FlashYellow",      255, 255,   0) }
func blackYellow() Colour      { return new3n ("BlackYellow",      238, 153,  34) }
func darkYellow() Colour       { return new3n ("DarkYellow",       255, 187,   0) }
func yellow() Colour           { return new3n ("Yellow",           255, 255,  34) }
func lightYellow() Colour      { return new3n ("LightYellow",      255, 255, 102) }
func whiteYellow() Colour      { return new3n ("WhiteYellow",      255, 255, 153) }
func flashGreen() Colour       { return new3n ("FlashGreen",         0, 255,   0) }
func blackGreen() Colour       { return new3n ("BlackGreen",         0,  51,   0) }
func darkGreen() Colour        { return new3n ("DarkGreen",          0,  85,   0) }
func green() Colour            { return new3n ("Green",              0, 170,   0) }
func lightGreen() Colour       { return new3n ("LightGreen",        51, 204,  51) }
func whiteGreen() Colour       { return new3n ("WhiteGreen",       102, 255, 102) }
func grassGreen() Colour       { return new3n ("GrassGreen",         0, 144,   0) }
func umber() Colour            { return new3n ("Umber",            149, 135,   0) }
func oliveGreen() Colour       { return new3n ("OliveGreen",        61,  69,  46) }
func lightOliveGreen() Colour  { return new3n ("LightOliveGreen",  170, 196,  85) }
func yellowGreen() Colour      { return new3n ("YellowGreen",      170, 255,  85) }
func meadowGreen() Colour      { return new3n ("MeadowGreen",      106, 212, 106) }
func flashCyan() Colour        { return new3n ("FlashCyan",          0, 255, 255) }
func blackCyan() Colour        { return new3n ("BlackCyan",          0,  51,  51) }
func darkCyan() Colour         { return new3n ("DarkCyan",           0,  85,  85) }
func cyan() Colour             { return new3n ("Cyan",               0, 170, 170) }
func lightCyan() Colour        { return new3n ("LightCyan",         85, 255, 255) }
func whiteCyan() Colour        { return new3n ("WhiteCyan",        153, 255, 255) }
func flashBlue() Colour        { return new3n ("FlashBlue",          0,   0, 255) }
func blackBlue() Colour        { return new3n ("BlackBlue",          0,   0,  51) }
func darkBlue() Colour         { return new3n ("DarkBlue",           0,   0,  85) }
func blue() Colour             { return new3n ("Blue",               0,   0, 170) }
func lightBlue() Colour        { return new3n ("LightBlue",         51,  51, 204) }
func whiteBlue() Colour        { return new3n ("WhiteBlue",        102, 102, 255) }

func prussianBlue() Colour     { return new3n ("PrussianBlue",       0,  85, 136) }
func skyLightBlue() Colour     { return new3n ("SkyLightBlue",      85, 170, 255) }
func skyBlue() Colour          { return new3n ("SkyBlue",            0, 170, 255) }
func gentianBlue() Colour      { return new3n ("GentianBlue",        0,   0, 204) }
func ultramarine() Colour      { return new3n ("Ultramarine",       68,   0, 153) }
func ultramarinBlue() Colour   { return new3n ("UltramarinBlue",     0,  15, 117) }

func flashMagenta() Colour     { return new3n ("FlashMagenta",     255,   0, 255) }
func blackMagenta() Colour     { return new3n ("BlackMagenta",      51,   0,  51) }
func darkMagenta() Colour      { return new3n ("DarkMagenta",       85,   0,  85) }
func magenta() Colour          { return new3n ("Magenta",          170,   0, 170) }
func lightMagenta() Colour     { return new3n ("LightMagenta",     255,  85, 255) }
func whiteMagenta() Colour     { return new3n ("WhiteMagenta",     255, 153, 255) }
func pink() Colour             { return new3n ("Pink",             255,   0, 170) }
func deepPink() Colour         { return new3n ("DeepPink",         255,  17,  51) }
func black() Colour            { return new3n ("Black",              0,   0,   0) }
func flashGray() Colour        { return new3n ("FlashGray",        127, 127, 127) }
func blackGray() Colour        { return new3n ("BlackGray",         34,  34,  34) }
func darkGray() Colour         { return new3n ("DarkGray",          51,  51,  51) }
func gray() Colour             { return new3n ("Gray",              85,  85,  85) }
func lightGray() Colour        { return new3n ("LightGray",        136, 136, 136) }
func whiteGray() Colour        { return new3n ("WhiteGray",        204, 204, 204) }
func silver() Colour           { return new3n ("Silver",           212, 212, 212) }
func flashWhite() Colour       { return new3n ("FlashWhite",       255, 255, 255) }
func white() Colour            { return new3n ("White",            170, 170, 170) }

// RAL-Farben: }
func grünbeige() Colour        { return new3n ("Grünbeige",        214, 199, 148) }
func beige() Colour            { return new3n ("Beige",            217, 186, 140) }
func sandgelb() Colour         { return new3n ("Sandgelb",         214, 176, 117) }
func signalgelb() Colour       { return new3n ("Signalgelb ",      252, 163,  41) }
func goldgelb() Colour         { return new3n ("Goldgelb",         227, 150,  36) }
func honiggelb() Colour        { return new3n ("Honiggelb",        201, 135,  33) }
func maisgelb() Colour         { return new3n ("Maisgelb",         224, 130,  31) }
func narzissengelb() Colour    { return new3n ("Narzissengelb",    227, 122,  31) }
func braunbeige() Colour       { return new3n ("Braunbeige",       173, 122,  79) }
func zitronengelb() Colour     { return new3n ("Zitronengelb",     227, 184,  56) }
func perlweiß() Colour         { return new3n ("Perlweiß",         255, 245, 227) }
func elfenbein() Colour        { return new3n ("Elfenbein",        240, 214, 171) }
func hellelfenbein() Colour    { return new3n ("Hellelfenbein",    252, 235, 204) }
func schwefelgelb() Colour     { return new3n ("Schwefelgelb",     255, 245,  66) }
func safrangelb() Colour       { return new3n ("Safrangelb",       255, 171,  89) }
func zinkgelb() Colour         { return new3n ("Zinkgelb",         255, 214,  77) }
func graubeige() Colour        { return new3n ("Graubeige",        164, 143, 122) }
func olivgelb() Colour         { return new3n ("Olivgelb",         156, 143,  97) }
func rapsgelb() Colour         { return new3n ("Rapsgelb",         252, 189,  31) }
func verkehrsgelb() Colour     { return new3n ("Verkehrsgelb",     252, 184,  33) }
func ockergelb() Colour        { return new3n ("Ockergelb",        181, 140,  79) }
func leuchtgelb() Colour       { return new3n ("Leuchtgelb",       255, 255,  10) }
func currygelb() Colour        { return new3n ("Currygelb",        153, 117,  33) }
func melonengelb() Colour      { return new3n ("Melonengelb",      255, 140,  26) }
func ginstergelb() Colour      { return new3n ("Ginstergelb",      227, 163,  41) }
func dahliengelb() Colour      { return new3n ("Dahliengelb",      255, 148,  54) }
func pastellgelb() Colour      { return new3n ("Pastellgelb",      247, 153,  92) }
func perlbeige() Colour        { return new3n ("Perlbeige",        143, 131, 112) }
func perlgold() Colour         { return new3n ("Perlgold",         128, 100,  64) }
func sonnengelb() Colour       { return new3n ("Sonnengelb",       240, 146,   0) }
func gelborange() Colour       { return new3n ("Gelborange",       224,  94,  31) }
func rotorange() Colour        { return new3n ("Rotorange",        186,  46,  33) }
func blutorange() Colour       { return new3n ("Blutorange",       204,  36,  28) }
func pastellorange() Colour    { return new3n ("Pastellorange",    255,  99,  54) }
func reinorange() Colour       { return new3n ("Reinorange",       242,  59,  28) }
func leuchtorange() Colour     { return new3n ("Leuchtorange",     252,  28,  20) }
func leuchthellorange() Colour { return new3n ("Leuchthellorange", 255, 117,  33) }
func hellrotorange() Colour    { return new3n ("Hellrotorange",    250,  79,  41) }
func verkehrsorange() Colour   { return new3n ("Verkehrsorange",   235,  59,  28) }
func signalorange() Colour     { return new3n ("Signalorange",     212,  69,  41) }
func tieforange() Colour       { return new3n ("Tieforange",       237,  92,  41) }
func lachsorange() Colour      { return new3n ("Lachsorange",      222,  82,  71) }
func perlorange() Colour       { return new3n ("Perlorange",       146,  62,  37) }
func ralorange() Colour        { return new3n ("RALorange",        252,  85,   0) }
func feuerrot() Colour         { return new3n ("Feuerrot",         171,  31,  28) }
func signalrot() Colour        { return new3n ("Signalrot",        163,  23,  26) }
func karminrot() Colour        { return new3n ("Karminrot",        163,  26,  26) }
func rubinrot() Colour         { return new3n ("Rubinrot",         138,  18,  20) }
func purpurrot() Colour        { return new3n ("Purpurrot",        105,  15,  20) }
func weinrot() Colour          { return new3n ("Weinrot",           79,  18,  26) }
func schwarzrot() Colour       { return new3n ("Schwarzrot",        46,  18,  26) }
func oxidrot() Colour          { return new3n ("Oxidrot",           94,  33,  33) }
func braunrot() Colour         { return new3n ("Braunrot",         120,  20,  23) }
func beigerot() Colour         { return new3n ("Beigerot",         204, 130, 115) }
func tomatenrot() Colour       { return new3n ("Tomatenrot",       150,  31,  28) }
func altrosa() Colour          { return new3n ("Altrosa",          217, 102, 117) }
func hellrosa() Colour         { return new3n ("Hellrosa",         232, 156, 181) }
func korallenrot() Colour      { return new3n ("Korallenrot",      166,  36,  38) }
func rose() Colour             { return new3n ("Rose",             202,  85,  93) }
func erdbeerrot() Colour       { return new3n ("Erdbeerrot",       207,  41,  66) }
func verkehrsrot() Colour      { return new3n ("Verkehrsrot",      199,  23,  18) }
func lachsrot() Colour         { return new3n ("Lachsrot",         217,  89,  79) }
func leuchtrot() Colour        { return new3n ("Leuchtrot",        252,  10,  28) }
func leuchthellrot() Colour    { return new3n ("Leuchthellrot",    252,  20,  20) }
func himbeerrot() Colour       { return new3n ("Himbeerrot",       181,  18,  51) }
func reinrot() Colour          { return new3n ("Reinrot",          204,  44,  36) }
func orientrot() Colour        { return new3n ("Orientrot",        166,  28,  46) }
func perlrubinrot() Colour     { return new3n ("Perlrubinrot",     112,  29,  36) }
func perlrosa() Colour         { return new3n ("Perlrosa",         165,  58,  46) }
func rotlila() Colour          { return new3n ("Rotlila",          130,  64, 128) }
func rotviolett() Colour       { return new3n ("Rotviolett",       143,  38,  64) }
func erikaviolett() Colour     { return new3n ("Erikaviolett",     201,  56, 140) }
func bordeauxviolett() Colour  { return new3n ("Bordeauxviolett",   92,   8,  43) }
func blaulila() Colour         { return new3n ("Blaulila",          99,  61, 156) }
func verkehrspurpur() Colour   { return new3n ("Verkehrspurpur",   145,  15, 102) }
func purpurviolett() Colour    { return new3n ("Purpurviolett",     56,  10,  46) }
func signalviolett() Colour    { return new3n ("Signalviolett",    125,  31, 122) }
func pastellviolett() Colour   { return new3n ("Pastellviolett",   158, 115, 148) }
func telemagenta() Colour      { return new3n ("Telemagenta",      187,  64, 119) }
func perlviolett() Colour      { return new3n ("Perlviolett",      110,  63,  87) }
func perlbrombeer() Colour     { return new3n ("Perlbrombeer",     106, 107, 127) }
func violettblau() Colour      { return new3n ("Violettblau",       23,  51, 107) }
func grünblau() Colour         { return new3n ("Grünblau",          10,  51,  84) }
func ultramarinblau() Colour   { return new3n ("Ultramarinblau",     0,  15, 117) }
func saphirblau() Colour       { return new3n ("Saphirblau",         0,  23,  69) }
func schwarzblau() Colour      { return new3n ("Schwarzblau",        3,  13,  31) }
func signalblau() Colour       { return new3n ("Signalblau",         0,  46, 122) }
func brillantblau() Colour     { return new3n ("Brillantblau",      38,  79, 135) }
func graublau() Colour         { return new3n ("Graublau",          26,  41,  56) }
func azurblau() Colour         { return new3n ("Azurblau",          23,  69, 112) }
func enzianblau() Colour       { return new3n ("Enzianblau",         0,  43, 112) }
func stahlblau() Colour        { return new3n ("Stahlblau",          3,  20,  46) }
func lichtblau() Colour        { return new3n ("Lichtblau",         41, 115, 184) }
func kobaltblau() Colour       { return new3n ("Kobaltblau",         0,  18,  69) }
func taubenblau() Colour       { return new3n ("Taubenblau",        77, 105, 153) }
func himmelblau() Colour       { return new3n ("Himmelblau",        23,  97, 171) }
func verkehrsblau() Colour     { return new3n ("Verkehrsblau",       0,  59, 128) }
func türkisblau() Colour       { return new3n ("Türkisblau",        56, 148, 130) }
func capriblau() Colour        { return new3n ("Capriblau",         10,  66, 120) }
func ozeanblau() Colour        { return new3n ("Ozeanblau",          5,  51,  51) }
func wasserblau() Colour       { return new3n ("Wasserblau",        26, 122,  99) }
func nachtblau() Colour        { return new3n ("Nachtblau",          0,   8,  79) }
func fernblau() Colour         { return new3n ("Fernblau",          46,  82, 143) }
func pastellblau() Colour      { return new3n ("Pastellblau",       87, 140, 181) }
func perlenzian() Colour       { return new3n ("Perlenzian",        32, 105, 124) }
func perlnachtblau() Colour    { return new3n ("Perlnachtblau",     15,  48,  82) }
func patinagrün() Colour       { return new3n ("Patinagrün",        51, 120,  84) }
func smaragdgrün() Colour      { return new3n ("Smaragdgrün",       38, 102,  41) }
func laubgrün() Colour         { return new3n ("Laubgrün",          38,  87,  33) }
func olivgrün() Colour         { return new3n ("Olivgrün",          80,  83,  60) }
func blaugrün() Colour         { return new3n ("Blaugrün",          13,  59,  46) }
func moosgrün() Colour         { return new3n ("Moosgrün",          10,  56,  31) }
func grauoliv() Colour         { return new3n ("Grauoliv",          41,  43,  36) }
func flaschengrün() Colour     { return new3n ("Flaschengrün",      44,  50,  34) }
func braungrün() Colour        { return new3n ("Braungrün",         54,  52,  42) }
func tannengrün() Colour       { return new3n ("Tannengrün",        23,  41,  28) }
func grasgrün() Colour         { return new3n ("Grasgrün",          54, 105,  38) }
func resedagrün() Colour       { return new3n ("Resedagrün",        94, 125,  79) }
func schwarzgrün() Colour      { return new3n ("Schwarzgrün",       31,  46,  43) }
func schilfgrün() Colour       { return new3n ("Schilfgrün",       117, 115,  79) }
func gelboliv() Colour         { return new3n ("Gelboliv",          51,  48,  38) }
func schwarzoliv() Colour      { return new3n ("Schwarzoliv",       41,  43,  38) }
func türkisgrün() Colour       { return new3n ("Türkisgrün",         0, 105,  76) }
func maigrün() Colour          { return new3n ("Maigrün",           64, 130,  54) }
func gelbgrün() Colour         { return new3n ("Gelbgrün",          79, 168,  51) }
func weißgrün() Colour         { return new3n ("Weißgrün",         185, 206, 172) }
func chromoxidgrün() Colour    { return new3n ("Chromoxidgrün",     38,  56,  41) }
func blassgrün() Colour        { return new3n ("Blassgrün",        138, 153, 119) }
func braunoliv() Colour        { return new3n ("Braunoliv",         43,  38,  28) }
func verkehrsgrün() Colour     { return new3n ("Verkehrsgrün",      36, 145,  64) }
func farngrün() Colour         { return new3n ("Farngrün",          74, 110,  51) }
func opalgrün() Colour         { return new3n ("Opalgrün",          10,  92,  51) }
func lichtgrün() Colour        { return new3n ("Lichtgrün",        125, 204, 189) }
func kieferngrün() Colour      { return new3n ("Kieferngrün",       38,  74,  51) }
func minzgrün() Colour         { return new3n ("Minzgrün",          18, 120,  38) }
func signalgrün() Colour       { return new3n ("Signalgrün",        41, 138,  64) }
func minttürkis() Colour       { return new3n ("Minttürkis",        66, 140, 120) }
func pastelltürkis() Colour    { return new3n ("Pastelltürkis",    122, 173, 172) }
func perlgrün() Colour         { return new3n ("Perlgrün",          25,  77,  37) }
func perlopalgrün() Colour     { return new3n ("Perlopalgrün",       4,  87,  75) }
func reingrün() Colour         { return new3n ("Reingrün",           0, 139,  41) }
func leuchtgrün() Colour       { return new3n ("Leuchtgrün",         0, 181,  27) }
func fasergrün() Colour        { return new3n ("Fasergrün",        179, 196 , 62) }
func fehgrau() Colour          { return new3n ("Fehgrau",          115, 133, 145) }
func silbergrau() Colour       { return new3n ("Silbergrau",       135, 148, 166) }
func olivgrau() Colour         { return new3n ("Olivgrau",         122, 117,  97) }
func moosgrau() Colour         { return new3n ("Moosgrau",         112, 112,  97) }
func signalgrau() Colour       { return new3n ("Signalgrau",       156, 156, 166) }
func mausgrau() Colour         { return new3n ("Mausgrau",          97, 105, 105) }
func beigegrau() Colour        { return new3n ("Beigegrau",        107,  97,  87) }
func khakigrau() Colour        { return new3n ("Khakigrau",        105,  84,  56) }
func grüngrau() Colour         { return new3n ("Grüngrau",          77,  82,  74) }
func zeltgrau() Colour         { return new3n ("Zeltgrau",          74,  79,  74) }
func eisengrau() Colour        { return new3n ("Eisengrau",         64,  74,  84) }
func basaltgrau() Colour       { return new3n ("Basaltgrau",        74,  84,  89) }
func braungrau() Colour        { return new3n ("Braungrau",         71,  66,  56) }
func schiefergrau() Colour     { return new3n ("Schiefergrau",      61,  66,  82) }
func anthrazitgrau() Colour    { return new3n ("Anthrazitgrau",     38,  46,  56) }
func schwarzgrau() Colour      { return new3n ("Schwarzgrau",       26,  33,  41) }
func umbragrau() Colour        { return new3n ("Umbragrau",         61,  61,  59) }
func betongrau() Colour        { return new3n ("Betongrau",        122, 125, 117) }
func graphitgrau() Colour      { return new3n ("Graphitgrau",       48,  56,  69) }
func granitgrau() Colour       { return new3n ("Granitgrau",        38,  51,  56) }
func steingrau() Colour        { return new3n ("Steingrau",        145, 143, 135) }
func blaugrau() Colour         { return new3n ("Blaugrau",          77,  92, 107) }
func kieselgrau() Colour       { return new3n ("Kieselgrau",       189, 186, 171) }
func zementgrau() Colour       { return new3n ("Zementgrau",       122, 130, 117) }
func gelbgrau() Colour         { return new3n ("Gelbgrau",         143, 135, 112) }
func lichtgrau() Colour        { return new3n ("Lichtgrau",        212, 217, 219) }
func platingrau() Colour       { return new3n ("Platingrau",       158, 150, 156) }
func staubgrau() Colour        { return new3n ("Staubgrau",        122, 125, 128) }
func achatgrau() Colour        { return new3n ("Achatgrau",        186, 189, 186) }
func quarzgrau() Colour        { return new3n ("Quarzgrau",         97,  94,  89) }
func fenstergrau() Colour      { return new3n ("Fenstergrau",      158, 163, 176) }
func verkehrsgrauA() Colour    { return new3n ("VerkehrsgrauA",    143, 150, 153) }
func verkehrsgrauB() Colour    { return new3n ("VerkehrsgrauB",     64,  69,  69) }
func seidengrau() Colour       { return new3n ("Seidengrau",       194, 191, 184) }
func telegrau1() Colour        { return new3n ("Telegrau1",        143, 148, 158) }
func telegrau2() Colour        { return new3n ("Telegrau2",        120, 130, 140) }
func telegrau4() Colour        { return new3n ("Telegrau4",        217, 214, 219) }
func perlmausgrau() Colour     { return new3n ("Perlmausgrau",     129, 123, 115) }
func grünbraun() Colour        { return new3n ("Grünbraun",        125,  92,  56) }
func ockerbraun() Colour       { return new3n ("Ockerbraun",       145,  82,  46) }
func signalbraun() Colour      { return new3n ("Signalbraun",      110,  59,  48) }
func lehmbraun() Colour        { return new3n ("Lehmbraun",        115,  59,  36) }
func kupferbraun() Colour      { return new3n ("Kupferbraun",      133,  56,  43) }
func rehbraun() Colour         { return new3n ("Rehbraun",          94,  51,  31) }
func olivbraun() Colour        { return new3n ("Olivbraun",         99,  61,  36) }
func nussbraun() Colour        { return new3n ("Nussbraun",         71,  38,  28) }
func rotbraun() Colour         { return new3n ("Rotbraun",          84,  31,  31) }
func sepiabraun() Colour       { return new3n ("Sepiabraun",        56,  38,  28) }
func kastanienbraun() Colour   { return new3n ("Kastanienbraun",    77,  31,  28) }
func mahagonibraun() Colour    { return new3n ("Mahagonibraun",     61,  31,  28) }
func schokoladenbraun() Colour { return new3n ("Schokoladenbraun",  46,  28,  28) }
func graubraun() Colour        { return new3n ("Graubraun",         43,  38,  41) }
func schwarzbraun() Colour     { return new3n ("Schwarzbraun",      13,   8,  13) }
func orangebraun() Colour      { return new3n ("Orangebraun",      156,  69,  41) }
func beigebraun() Colour       { return new3n ("Beigebraun",       110,  64,  48) }
func terrabraun() Colour       { return new3n ("Terrabraun",        64,  46,  33) }
func blassbraun() Colour       { return new3n ("Blassbraun",       102,  74,  61) }
func perlkupfer() Colour       { return new3n ("Perlkupfer",       127,  64,  49) }
func cremeweiß() Colour        { return new3n ("Cremeweiß",        255, 252, 240) }
func grauweiß() Colour         { return new3n ("Grauweiß",         240, 237, 230) }
func signalweiß() Colour       { return new3n ("Signalweiß",       255, 255, 255) }
func signalschwarz() Colour    { return new3n ("Signalschwarz",     28,  28,  33) }
func tiefschwarz() Colour      { return new3n ("Tiefschwarz",        3,   5,  10) }
func weißaluminium() Colour    { return new3n ("Weißaluminium",    166, 171, 181) }
func graualuminium() Colour    { return new3n ("Graualuminium",    125, 122, 120) }
func reinweiß() Colour         { return new3n ("Reinweiß",         250, 255, 255) }
func graphitschwarz() Colour   { return new3n ("Graphitschwarz",    13,  18,  26) }
func reinraumweiß() Colour     { return new3n ("Reinraumweiß",     248, 242, 225) }
func verkehrsweiß() Colour     { return new3n ("Verkehrsweiß",     252, 255, 255) }
func verkehrsschwarz() Colour  { return new3n ("Verkehrsschwarz",   20,  23,  28) }
func papyrusweiß() Colour      { return new3n ("Papyrusweiß",      219, 227, 222) }
func perlhellgrau() Colour     { return new3n ("Perlhellgrau",     133, 133, 131) }
func perldunkelgrau() Colour   { return new3n ("Perldunkelgrau",   120, 123, 122) }
