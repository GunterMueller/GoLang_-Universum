package col

// (c) Christian Maurer   v. 230401 - license see µU.go

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

// Returns the colour Black.
func New() Colour { return new_() }

// Returns the colour defined by (r, g, b).
func New3 (r, g, b byte) Colour { return new3(r,g,b) }

// Returns the colour defined by (r, g, b) with name n.
func New3n (n string, r, g, b byte) Colour { return new3n(n,r,g,b) }

func HeadF() Colour { return flashWhite() }
func HeadB() Colour { return blue() }
func HintF() Colour { return flashWhite() }
func HintB() Colour { return magenta() }
func ErrorF() Colour { return flashYellow() }
func ErrorB() Colour { return red() }
func MenuF() Colour { return flashWhite() }
func MenuB() Colour { return red() }

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

func FlashBrown() Colour       { return flashBrown() }
func BlackBrown() Colour       { return blackBrown() }
func DarkBrown() Colour        { return darkBrown() }
func Brown() Colour            { return brown() }
func LightBrown() Colour       { return lightBrown() }
func WhiteBrown() Colour       { return whiteBrown() }

func DarkOchre() Colour        { return darkOchre() }
func Ochre() Colour            { return ochre() }
func LightOchre() Colour       { return lightOchre() }

func FlashRed() Colour         { return flashRed() }
func BlackRed() Colour         { return blackRed() }
func DarkRed() Colour          { return darkRed() }
func Red() Colour              { return red() }
func LightRed() Colour         { return lightRed() }
func WhiteRed() Colour         { return whiteRed() }
func PompejiRed() Colour       { return pompejiRed() }
func CinnabarRed() Colour      { return cinnabarRed() }
func Carmine() Colour          { return carmine() }
func BrickRed() Colour         { return brickRed() }
func Siena() Colour            { return siena() }
func LightSiena() Colour       { return lightSiena() }

func DarkRose() Colour         { return darkRose() }
func LightRose() Colour        { return lightRose() }
func WhiteRose() Colour        { return whiteRose() }

func FlashOrange() Colour      { return flashOrange() }
func BlackOrange() Colour      { return blackOrange() }
func DarkOrange() Colour       { return darkOrange() }
func Orange() Colour           { return orange() }
func LightOrange() Colour      { return lightOrange() }
func WhiteOrange() Colour      { return whiteOrange() }

func FlashYellow() Colour      { return flashYellow() }
func BlackYellow() Colour      { return blackYellow() }
func DarkYellow() Colour       { return darkYellow() }
func Yellow() Colour           { return yellow() }
func LightYellow() Colour      { return lightYellow() }
func WhiteYellow() Colour      { return whiteYellow() }

func FlashGreen() Colour       { return flashGreen() }
func BlackGreen() Colour       { return blackGreen() }
func DarkGreen() Colour        { return darkGreen() }
func Green() Colour            { return green() }
func LightGreen() Colour       { return lightGreen() }
func WhiteGreen() Colour       { return whiteGreen() }
func GrassGreen() Colour       { return grassGreen() }
func Umber() Colour            { return umber() }
func OliveGreen() Colour       { return oliveGreen() }
func LightOliveGreen() Colour  { return lightOliveGreen() }
func YellowGreen() Colour      { return yellowGreen() }
func MeadowGreen() Colour      { return meadowGreen() }

func FlashCyan() Colour        { return flashCyan() }
func BlackCyan() Colour        { return blackCyan() }
func DarkCyan() Colour         { return darkCyan() }
func Cyan() Colour             { return cyan() }
func LightCyan() Colour        { return lightCyan() }
func WhiteCyan() Colour        { return whiteCyan() }

func FlashBlue() Colour        { return flashBlue() }
func BlackBlue() Colour        { return blackBlue() }
func DarkBlue() Colour         { return darkBlue() }
func Blue() Colour             { return blue() }
func LightBlue() Colour        { return lightBlue() }
func WhiteBlue() Colour        { return whiteBlue() }
func PrussianBlue() Colour     { return prussianBlue() }
func GentianBlue() Colour      { return gentianBlue() }
func SkyBlue() Colour          { return skyBlue() }
func SkyLightBlue() Colour     { return skyLightBlue() }
func Ultramarine() Colour      { return ultramarine() }
func UltramarinBlue() Colour   { return ultramarinBlue() }

func FlashMagenta() Colour     { return flashMagenta() }
func BlackMagenta() Colour     { return blackMagenta() }
func DarkMagenta() Colour      { return darkMagenta() }
func Magenta() Colour          { return magenta() }
func LightMagenta() Colour     { return lightMagenta() }
func WhiteMagenta() Colour     { return whiteMagenta() }

func Pink() Colour             { return pink() }
func DeepPink() Colour         { return deepPink() }

func Black() Colour            { return black() }

func FlashGray() Colour        { return flashGray() }
func BlackGray() Colour        { return blackGray() }
func DarkGray() Colour         { return darkGray() }
func Gray() Colour             { return gray() }
func LightGray() Colour        { return lightGray() }
func WhiteGray() Colour        { return whiteGray() }

func Silver() Colour           { return silver() }

func FlashWhite() Colour       { return flashWhite() }
func White() Colour            { return white() }

// RAL-Farben:

func Grünbeige() Colour        { return grünbeige() }
func Beige() Colour            { return beige() }
func Sandgelb() Colour         { return sandgelb() }
func Signalgelb() Colour       { return signalgelb() }
func Goldgelb() Colour         { return goldgelb() }
func Honiggelb() Colour        { return honiggelb() }
func Maisgelb() Colour         { return maisgelb() }
func Narzissengelb() Colour    { return narzissengelb() }
func Braunbeige() Colour       { return braunbeige() }
func Zitronengelb() Colour     { return zitronengelb() }
func Perlweiß() Colour         { return perlweiß() }
func Elfenbein() Colour        { return elfenbein() }
func Hellelfenbein() Colour    { return hellelfenbein() }
func Schwefelgelb() Colour     { return schwefelgelb() }
func Safrangelb() Colour       { return safrangelb() }
func Zinkgelb() Colour         { return zinkgelb() }
func Graubeige() Colour        { return graubeige() }
func Olivgelb() Colour         { return olivgelb() }
func Rapsgelb() Colour         { return rapsgelb() }
func Verkehrsgelb() Colour     { return verkehrsgelb() }
func Ockergelb() Colour        { return ockergelb() }
func Leuchtgelb() Colour       { return leuchtgelb() }
func Currygelb() Colour        { return currygelb() }
func Melonengelb() Colour      { return melonengelb() }
func Ginstergelb() Colour      { return ginstergelb() }
func Dahliengelb() Colour      { return dahliengelb() }
func Pastellgelb() Colour      { return pastellgelb() }
func Perlbeige() Colour        { return perlbeige() }
func Perlgold() Colour         { return perlgold() }
func Sonnengelb() Colour       { return sonnengelb() }

func Gelborange() Colour       { return gelborange() }
func Rotorange() Colour        { return rotorange() }
func Blutorange() Colour       { return blutorange() }
func Pastellorange() Colour    { return pastellorange() }
func Reinorange() Colour       { return reinorange() }
func Leuchtorange() Colour     { return leuchtorange() }
func Leuchthellorange() Colour { return leuchthellorange() }
func Hellrotorange() Colour    { return hellrotorange() }
func Verkehrsorange() Colour   { return verkehrsorange() }
func Signalorange() Colour     { return signalorange() }
func Tieforange() Colour       { return tieforange() }
func Lachsorange() Colour      { return lachsorange() }
func Perlorange() Colour       { return perlorange() }
func RALorange() Colour        { return ralorange() }

func Feuerrot() Colour         { return feuerrot() }
func Signalrot() Colour        { return signalrot() }
func Karminrot() Colour        { return karminrot() }
func Rubinrot() Colour         { return rubinrot() }
func Purpurrot() Colour        { return purpurrot() }
func Weinrot() Colour          { return weinrot() }
func Schwarzrot() Colour       { return schwarzrot() }
func Oxidrot() Colour          { return oxidrot() }
func Braunrot() Colour         { return braunrot() }
func Beigerot() Colour         { return beigerot() }
func Tomatenrot() Colour       { return tomatenrot() }
func Altrosa() Colour          { return altrosa() }
func Hellrosa() Colour         { return hellrosa() }
func Korallenrot() Colour      { return korallenrot() }
func Rose() Colour             { return rose() }
func Erdbeerrot() Colour       { return erdbeerrot() }
func Verkehrsrot() Colour      { return verkehrsrot() }
func Lachsrot() Colour         { return lachsrot () }
func Leuchtrot() Colour        { return leuchtrot() }
func Leuchthellrot() Colour    { return leuchthellrot() }
func Himbeerrot() Colour       { return himbeerrot() }
func Reinrot() Colour          { return reinrot() }
func Orientrot() Colour        { return orientrot() }
func Perlrubinrot() Colour     { return perlrubinrot() }
func Perlrosa() Colour         { return perlrosa() }

func Rotlila() Colour          { return rotlila() }
func Rotviolett() Colour       { return rotviolett() }
func Erikaviolett() Colour     { return erikaviolett() }
func Bordeauxviolett() Colour  { return bordeauxviolett() }
func Blaulila() Colour         { return blaulila() }
func Verkehrspurpur() Colour   { return verkehrspurpur() }
func Purpurviolett() Colour    { return purpurviolett() }
func Signalviolett() Colour    { return signalviolett() }
func Pastellviolett() Colour   { return pastellviolett() }
func Telemagenta() Colour      { return telemagenta() }
func Perlviolett() Colour      { return perlviolett() }
func Perlbrombeer() Colour     { return perlbrombeer() }

func Violettblau() Colour      { return violettblau() }
func Grünblau() Colour         { return grünblau() }
func Ultramarinblau() Colour   { return ultramarinblau() }
func Saphirblau() Colour       { return saphirblau() }
func Schwarzblau() Colour      { return schwarzblau() }
func Signalblau() Colour       { return signalblau() }
func Brillantblau() Colour     { return brillantblau() }
func Graublau() Colour         { return graublau() }
func Azurblau() Colour         { return azurblau() }
func Enzianblau() Colour       { return enzianblau() }
func Stahlblau() Colour        { return stahlblau() }
func Lichtblau() Colour        { return lichtblau() }
func Kobaltblau() Colour       { return kobaltblau() }
func Taubenblau() Colour       { return taubenblau() }
func Himmelblau() Colour       { return himmelblau() }
func Verkehrsblau() Colour     { return verkehrsblau() }
func Türkisblau() Colour       { return türkisblau() }
func Capriblau() Colour        { return capriblau() }
func Ozeanblau() Colour        { return ozeanblau() }
func Wasserblau() Colour       { return wasserblau() }
func Nachtblau() Colour        { return nachtblau() }
func Fernblau() Colour         { return fernblau() }
func Pastellblau() Colour      { return pastellblau() }
func Perlenzian() Colour       { return perlenzian() }
func Perlnachtblau() Colour    { return perlnachtblau() }

func Patinagrün() Colour       { return patinagrün() }
func Smaragdgrün() Colour      { return smaragdgrün() }
func Laubgrün() Colour         { return laubgrün() }
func Olivgrün() Colour         { return olivgrün() }
func Blaugrün() Colour         { return blaugrün() }
func Moosgrün() Colour         { return moosgrün() }
func Grauoliv() Colour         { return grauoliv() }
func Flaschengrün() Colour     { return flaschengrün() }
func Braungrün() Colour        { return braungrün() }
func Tannengrün() Colour       { return tannengrün() }
func Grasgrün() Colour         { return grasgrün() }
func Resedagrün() Colour       { return resedagrün() }
func Schwarzgrün() Colour      { return schwarzgrün() }
func Schilfgrün() Colour       { return schilfgrün() }
func Gelboliv() Colour         { return gelboliv() }
func Schwarzoliv() Colour      { return schwarzoliv() }
func Türkisgrün() Colour       { return türkisgrün() }
func Maigrün() Colour          { return maigrün() }
func Gelbgrün() Colour         { return gelbgrün() }
func Weißgrün() Colour         { return weißgrün() }
func Chromoxidgrün() Colour    { return chromoxidgrün() }
func Blassgrün() Colour        { return blassgrün() }
func Braunoliv() Colour        { return braunoliv() }
func Verkehrsgrün() Colour     { return verkehrsgrün() }
func Farngrün() Colour         { return farngrün() }
func Opalgrün() Colour         { return opalgrün() }
func Lichtgrün() Colour        { return lichtgrün() }
func Kieferngrün() Colour      { return kieferngrün() }
func Minzgrün() Colour         { return minzgrün() }
func Signalgrün() Colour       { return signalgrün() }
func Minttürkis() Colour       { return minttürkis() }
func Pastelltürkis() Colour    { return pastelltürkis() }
func Perlgrün() Colour         { return perlgrün() }
func Perlopalgrün() Colour     { return perlopalgrün() }
func Reingrün() Colour         { return reingrün() }
func Leuchtgrün() Colour       { return leuchtgrün() }
func Fasergrün() Colour        { return fasergrün() }

func Fehgrau() Colour          { return fehgrau() }
func Silbergrau() Colour       { return silbergrau() }
func Olivgrau() Colour         { return olivgrau() }
func Moosgrau() Colour         { return moosgrau() }
func Signalgrau() Colour       { return signalgrau() }
func Mausgrau() Colour         { return mausgrau() }
func Beigegrau() Colour        { return beigegrau() }
func Khakigrau() Colour        { return khakigrau() }
func Grüngrau() Colour         { return grüngrau() }
func Zeltgrau() Colour         { return zeltgrau() }
func Eisengrau() Colour        { return eisengrau() }
func Basaltgrau() Colour       { return basaltgrau() }
func Braungrau() Colour        { return braungrau() }
func Schiefergrau() Colour     { return schiefergrau() }
func Anthrazitgrau() Colour    { return anthrazitgrau() }
func Schwarzgrau() Colour      { return schwarzgrau() }
func Umbragrau() Colour        { return umbragrau() }
func Betongrau() Colour        { return betongrau() }
func Graphitgrau() Colour      { return graphitgrau() }
func Granitgrau() Colour       { return granitgrau() }
func Steingrau() Colour        { return steingrau() }
func Blaugrau() Colour         { return blaugrau() }
func Kieselgrau() Colour       { return kieselgrau() }
func Zementgrau() Colour       { return zementgrau() }
func Gelbgrau() Colour         { return gelbgrau() }
func Lichtgrau() Colour        { return lichtgrau() }
func Platingrau() Colour       { return platingrau() }
func Staubgrau() Colour        { return staubgrau() }
func Achatgrau() Colour        { return achatgrau() }
func Quarzgrau() Colour        { return quarzgrau() }
func Fenstergrau() Colour      { return fenstergrau() }
func VerkehrsgrauA() Colour    { return verkehrsgrauA() }
func VerkehrsgrauB() Colour    { return verkehrsgrauB() }
func Seidengrau() Colour       { return seidengrau() }
func Telegrau1() Colour        { return telegrau1() }
func Telegrau2() Colour        { return telegrau2() }
func Telegrau4() Colour        { return telegrau4() }
func Perlmausgrau() Colour     { return perlmausgrau() }

func Grünbraun() Colour        { return grünbraun() }
func Ockerbraun() Colour       { return ockerbraun() }
func Signalbraun() Colour      { return signalbraun() }
func Lehmbraun() Colour        { return lehmbraun() }
func Kupferbraun() Colour      { return kupferbraun() }
func Rehbraun() Colour         { return rehbraun() }
func Olivbraun() Colour        { return olivbraun() }
func Nussbraun() Colour        { return nussbraun() }
func Rotbraun() Colour         { return rotbraun() }
func Sepiabraun() Colour       { return sepiabraun() }
func Kastanienbraun() Colour   { return kastanienbraun() }
func Mahagonibraun() Colour    { return mahagonibraun() }
func Schokoladenbraun() Colour { return schokoladenbraun() }
func Graubraun() Colour        { return graubraun() }
func Schwarzbraun() Colour     { return schwarzbraun() }
func Orangebraun() Colour      { return orangebraun() }
func Beigebraun() Colour       { return beigebraun() }
func Blassbraun() Colour       { return blassbraun() }
func Terrabraun() Colour       { return terrabraun() }
func Perlkupfer() Colour       { return perlkupfer() }

func Cremeweiß() Colour        { return cremeweiß() }
func Grauweiß() Colour         { return grauweiß() }
func Signalweiß() Colour       { return signalweiß() }
func Signalschwarz() Colour    { return signalschwarz() }
func Tiefschwarz() Colour      { return tiefschwarz() }
func Weißaluminium() Colour    { return weißaluminium() }
func Graualuminium() Colour    { return graualuminium() }
func Reinweiß() Colour         { return reinweiß() }
func Graphitschwarz() Colour   { return graphitschwarz() }
func Reinraumweiß() Colour     { return reinraumweiß() }
func Verkehrsweiß() Colour     { return verkehrsweiß() }
func Verkehrsschwarz() Colour  { return verkehrsschwarz() }
func Papyrusweiß() Colour      { return papyrusweiß() }
func Perlhellgrau() Colour     { return perlhellgrau() }
func Perldunkelgrau() Colour   { return perldunkelgrau() }
