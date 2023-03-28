package col

// (c) Christian Maurer   v. 230326 - license see µU.go
//
// >>> definition of colours under work
// >>> several RAL colours will be incorporated into the colours, the others will be obsolete

import
 . "µU/obj"
type
  Colour interface {

  Object // empty colour is black

// Encodes the colour with red and blue reversed.
  EncodeInv() Stream

// String returns the name of x, defined by the name given with New3.
  Stringer

// String1 returns (rrggbb", where "rr", "gg" and "bb" are the rgb-values
// in sedecimal basis (with uppercase letters).
  String1() string

// Defined returns true, iff s is a string of 3 values in sedecimal basis
// (with uppercase letters). In that case, c is the colour with
// the corresponding rgb-values; otherwise, nothing has happened.
  Defined1 (s string) bool

// Return the values of red/green/blue intensity of x.
  R() byte; G() byte; B() byte

// x is the colour defined by the values of r, g and b.
  Set (r, g, b byte)
  SetR (b byte); SetG (b byte); SetB (b byte)

// Liefert x.R() + 256 * x.G() + x.B().
  Code() uint

// Returns true, if c is, what the name of the func says.
  IsBlack() bool
  IsWhite() bool
  IsFlashWhite() bool

// Returns the rgb-values of x scaled to the range from 0 to 1.
  Float32() (float32, float32, float32)
  Float64() (float64, float64, float64)

// c is changed in a manner suggested by the name of the method.
  Invert()
  Contrast()
}

// Returns the colour White.
func New() Colour { return new_() }

// Returns the colour defined by (r, g, b).
func New3 (r, g, b byte) Colour { return new3(r,g,b) }

// Returns the colour defined by (r, g, b) with name n.
func New3n (n string, r, g, b byte) Colour { return new3n(n,r,g,b) }

func HeadF() Colour { return headF() }
func HeadB() Colour { return headB() }
func HintF() Colour { return hintF() }
func HintB() Colour { return hintB() }
func ErrorF() Colour { return errorF() }
func ErrorB() Colour { return errorB() }
func MenuF() Colour { return menuF() }
func MenuB() Colour { return menuB() }

// Returns a random colour.
func Rand() Colour { return random() }

// Returns the fore- and backgroundcolours at the start of the system
// for unmarked and marked objects.
func StartCols() (Colour, Colour) { return startCols() }
func StartColF() Colour { return startColF() }
func StartColB() Colour { return startColB() }

// Returns (FlashWhite, Black).
func StartColsA() (Colour, Colour) { return startColsA() }

// Returns the slice of all colours defined in this package.
func AllColours() []Colour { return allColours() }
func AllRALColours() []Colour { return allRALColours() }

func Black() Colour            { return black() }
func FlashBrown() Colour       { return flashBrown() }
func DarkerBrown() Colour      { return darkerBrown() }
func DarkBrown() Colour        { return darkBrown() }
func Brown() Colour            { return brown() }
func LightBrown() Colour       { return lightBrown() }
func LighterBrown() Colour     { return lighterBrown() }
func RedBrown() Colour         { return redBrown() }
func RoseBrown() Colour        { return roseBrown() }
func OliveBrown() Colour       { return oliveBrown() }
func LightOliveBrown() Colour  { return lightOliveBrown() }
func Umber() Colour            { return umber() }
func DarkOchre() Colour        { return darkOchre() }
func Ochre() Colour            { return ochre() }
func LightOchre() Colour       { return lightOchre() }
func LightBeige() Colour       { return lightBeige() }
func Beige1() Colour           { return beige1() }
func FlashRed() Colour         { return flashRed() }
func DarkerRed() Colour        { return darkerRed() }
func DarkRed() Colour          { return darkRed() }
func Red() Colour              { return red() }
func LightRed() Colour         { return lightRed() }
func LighterRed() Colour       { return lighterRed() }
func PompejiRed() Colour       { return pompejiRed() }
func CinnabarRed() Colour      { return cinnabarRed() }
func Carmine() Colour          { return carmine() }
func BrickRed() Colour         { return brickRed() }
func Siena() Colour            { return siena() }
func LightSiena() Colour       { return lightSiena() }
func DarkRose() Colour         { return darkRose() }
func Rose1() Colour            { return rose1() } // Rose: RAL
func LightRose() Colour        { return lightRose() }
func FlashOrange() Colour      { return flashOrange() }
func DarkerOrange() Colour     { return darkerOrange() }
func DarkOrange() Colour       { return darkOrange() }
func Orange() Colour           { return orange() }
func LightOrange() Colour      { return lightOrange() }
func LighterOrange() Colour    { return lighterOrange() }
func BloodOrange() Colour      { return bloodOrange() }
func FlashYellow() Colour      { return flashYellow() }
func DarkerYellow() Colour     { return darkerYellow() }
func DarkYellow() Colour       { return darkYellow() }
func Yellow() Colour           { return yellow() }
func LightYellow() Colour      { return lightYellow() }
func LighterYellow() Colour    { return lighterYellow() }
func SandYellow() Colour       { return sandYellow() }
func LemonYellow() Colour      { return lemonYellow() }
func FlashGreen() Colour       { return flashGreen() }
func DarkerGreen() Colour      { return darkerGreen() }
func DarkGreen() Colour        { return darkGreen() }
func Green() Colour            { return green() }
func LightGreen() Colour       { return lightGreen() }
func LighterGreen() Colour     { return lighterGreen() }
func BirchGreen() Colour       { return birchGreen() }
func GrassGreen() Colour       { return grassGreen() }
func OliveGreen() Colour       { return oliveGreen() }
func LightOliveGreen() Colour  { return lightOliveGreen() }
func YellowGreen() Colour      { return yellowGreen() }
func MeadowGreen() Colour      { return meadowGreen() }
func FlashCyan() Colour        { return flashCyan() }
func DarkerCyan() Colour       { return darkerCyan() }
func DarkCyan() Colour         { return darkCyan() }
func Cyan() Colour             { return cyan() }
func LightCyan() Colour        { return lightCyan() }
func LighterCyan() Colour      { return lighterCyan() }
func FlashBlue() Colour        { return flashBlue() }
func DarkerBlue() Colour       { return darkerBlue() }
func DarkBlue() Colour         { return darkBlue() }
func Blue() Colour             { return blue() }
func LightBlue() Colour        { return lightBlue() }
func LighterBlue() Colour      { return lighterBlue() }
func PrussianBlue() Colour     { return prussianBlue() }
func GentianBlue() Colour      { return gentianBlue() }
func SkyBlue() Colour          { return skyBlue() }
func SkyLightBlue() Colour     { return skyLightBlue() }
func Ultramarine() Colour      { return ultramarine() }
func FlashMagenta() Colour     { return flashMagenta() }
func DarkerMagenta() Colour    { return darkerMagenta() }
func DarkMagenta() Colour      { return darkMagenta() }
func Magenta() Colour          { return magenta() }
func LightMagenta() Colour     { return lightMagenta() }
func LighterMagenta() Colour   { return lighterMagenta() }
func Pink() Colour             { return pink() }
func DeepPink() Colour         { return deepPink() }
func FlashGray() Colour        { return flashGray() }
func DarkerGray() Colour       { return darkerGray() }
func DarkGray() Colour         { return darkGray() }
func Gray() Colour             { return gray() }
func LightGray() Colour        { return lightGray() }
func LighterGray() Colour      { return lighterGray() }
func Silver() Colour           { return silver() }
func FlashWhite() Colour       { return flashWhite( ) }
func White() Colour            { return white() }
func RedWhite() Colour         { return redWhite() }
func OrangeWhite() Colour      { return orangeWhite() }
func YellowWhite() Colour      { return yellowWhite() }
func GreenWhite() Colour       { return greenWhite() }
func CyanWhite() Colour        { return cyanWhite() }
func BlueWhite() Colour        { return blueWhite() }
func MagentaWhite() Colour     { return magentaWhite() }
