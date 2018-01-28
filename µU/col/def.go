package col

// (c) Christian Maurer   v. 171226 - license see µU.go

import
 . "µU/obj"
const
  P6 = 3
type
  Colour interface {

  Object // empty colour is black

// Defined returns true, iff s is a string of 3 values in sedecimal basis
// (with uppercase letters). In that case, c is the colour with
// the corresponding rgb-values; otherwise, nothing has happened.

// String returns "rrggbb", where "rr", "gg" and "bb" are the rgb-values
// in sedecimal basis (with uppercase letters).
  Stringer

// Return the values of red/green/blue intensity of x.
  R() byte; G() byte; B() byte

  SetR (b byte); SetG (b byte); SetB (b byte)

// TODO Spec
  Cc() Stream

// TODO Spec
  Code() uint

// Returns true, if c is, what the name of the func says.
  IsBlack() bool
  IsWhite() bool
  IsLightWhite() bool

// Returns the rgb-values of x scaled to the range from 0 to 1.
  Real() (float32, float32, float32)
  Double() (float64, float64, float64)

// c is changed in a manner suggested by the name of the func.
  Invert()
  Contrast()
}

// Returns the colour White.
func New() Colour { return new_() }

// Returns the colour defined by (r, g, b).
func New3 (r, g, b byte) Colour { return new3(r,g,b) }

// Returns a random colour.
func Rand() Colour { return random() }

func StartCols() (Colour, Colour) { return startCols() }
func StartColsA() (Colour, Colour) { return startColsA() }

// TODO Spec
//  func Change (c *Colour, rgb, d byte, l bool) { change(c,rgb,d,l) }

// Pre: b is one of 4, 8, 15, 16, 24 or 32.
// depth() Colour { return = (b + 4) / 8.
func SetDepth (b uint) { setDepth(b) }

func Depth() uint { return depth } // in bytes - must not be altered after the call to SetDepth !

// Returns the number of available colours, depending on depth.
func NCols() uint { return nCols() }

func P6Encode (a, p Stream) { p6Encode(a,p) }
func P6Colour (a Stream) Colour { return p6Colour(a) }

// func Cd (bs []byte) uint { return cd(bs) }


func HeadF() Colour { return headF }
func HeadB() Colour { return headB }
func HintF() Colour { return hintF }
func HintB() Colour { return hintB }
func ErrorF() Colour { return errorF }
func ErrorB() Colour { return errorB }
func MenuF() Colour { return menuF }
func MenuB() Colour { return menuB }
func MuF() Colour { return muF }
func MuB() Colour { return muB }

func Black() Colour { return black }
func Brown() Colour { return brown }
func BlackBrown() Colour { return blackBrown }
func DarkBrown() Colour { return darkBrown }
func MediumBrown() Colour { return mediumBrown }
func LightBrown() Colour { return lightBrown }
func WhiteBrown() Colour { return whiteBrown }
func BrownWhite() Colour { return brownWhite }
func Siena() Colour { return siena }
func LightSiena() Colour { return lightSiena }
//func RedBrown() Colour { return }
//func Umbrabraun() Colour { return }
func OliveBrown() Colour { return oliveBrown }
func LightOliveBrown() Colour { return lightOliveBrown }
//func OrangeBrown1() Colour { return }
//func Dark Ocker() Colour { return }
//func Ocker() Colour { return }
//func LightOcker() Colour { return }
//func Rosabraun() Colour { return }
//func Hellbeige() Colour { return }
//func Beige2() Colour { return }
//func VeryLightBrown() Colour { return }

func BlackRed() Colour { return blackRed }
func DarkRed() Colour { return darkRed }
func Red() Colour { return red }
func FlashRed() Colour { return flashRed }
func LightRed() Colour { return lightRed }
func WhiteRed() Colour { return whiteRed }
//func Dunkelrosa() Colour { return }
//func Rosa() Colour { return }
//func Hellrosa() Colour { return }
func PompejiRed() Colour { return pompejiRed }
func CinnabarRed() Colour { return cinnabarRed }
func Carmine() Colour { return carmine }
func BrickRed() Colour { return brickRed }

func FlashOrange() Colour { return flashOrange }
func DarkOrange() Colour { return darkOrange }
func Orange() Colour { return orange }
func LightOrange() Colour { return lightOrange }
func WhiteOrange() Colour { return whiteOrange }
//func BloodOrange() Colour { return }

func FlashYellow() Colour { return flashYellow }
func DarkYellow() Colour { return darkYellow }
func Yellow() Colour { return yellow }
func LightYellow() Colour { return lightYellow }
func WhiteYellow() Colour { return whiteYellow }
func Sandgelb1() Colour { return sandGelb1 }
func Zitronengelb1() Colour { return zitronenGelb1 }

func FlashGreen() Colour { return flashGreen }
func BlackGreen() Colour { return blackGreen }
func DarkGreen() Colour { return darkGreen }
func Green() Colour { return green }
func LightGreen() Colour { return lightGreen }
func WhiteGreen() Colour { return whiteGreen }
func BirchGreen() Colour { return birchGreen }
func GrassGreen() Colour { return grassGreen }
//func ChromeGreen() Colour { return }
//func LightChromeGreen() Colour { return }
func OliveGreen() Colour { return oliveGreen }
//func LightOliveGreen() Colour { return }
func YellowGreen() Colour { return yellowGreen }

func BlackCyan() Colour { return blackCyan }
func DarkCyan() Colour { return darkCyan }
func Cyan() Colour { return cyan }
func LightCyan() Colour { return lightCyan }
func WhiteCyan() Colour { return whiteCyan }
func FlashCyan() Colour { return flashCyan }

func FlashBlue() Colour { return flashBlue }
func BlackBlue() Colour { return blackBlue }
func PrussianBlue() Colour { return prussianBlue }
func DarkBlue() Colour { return darkBlue }
func Blue() Colour { return blue }
func LightBlue() Colour { return lightBlue }
func WhiteBlue() Colour { return whiteBlue }
func GentianBlue() Colour { return gentianBlue }
func SkyBlue() Colour { return skyBlue }
func Ultramarine() Colour { return ultramarine }

func BlackMagenta() Colour { return blackMagenta }
func DarkMagenta() Colour { return darkMagenta }
func Magenta() Colour { return magenta }
func LightMagenta() Colour { return lightMagenta }
func FlashMagenta() Colour { return flashMagenta }
func WhiteMagenta() Colour { return whiteMagenta }
func Pink() Colour { return pink }
func DeepPink() Colour { return deepPink }

func BlackGray() Colour { return blackGray }
func DarkGray() Colour { return darkGray }
func Gray() Colour { return gray }
func LightGray() Colour { return lightGray }
func WhiteGray() Colour { return whiteGray }
func Silver() Colour { return silver }

func White() Colour { return white }
func LightWhite() Colour { return lightWhite }

func Schwarz() Colour { return black }
func Braun() Colour { return brown }
func Dunkelbraun() Colour { return darkBrown }
func Hellbraun() Colour { return lightBrown }
func Weißbraun() Colour { return whiteBrown }
func Braunweiß() Colour { return brownWhite }
func Rot() Colour { return red }
func Hellrot() Colour { return lightRed }
func Dunkelrot() Colour { return darkRed }
func Grellrot() Colour { return flashRed }
func Weißrot() Colour { return whiteRed }
func Grellorange() Colour { return flashOrange }
func Dunkelorange() Colour { return darkOrange }
func Hellorange() Colour { return lightOrange }
func Weißorange() Colour { return whiteOrange }
func Grellgelb() Colour { return flashYellow }
func Dunkelgelb() Colour { return darkYellow }
func Gelb() Colour { return yellow }
func Hellgelb() Colour { return lightYellow }
func Weißgelb() Colour { return whiteYellow }
func Grellgrün() Colour { return flashGreen }
func Dunkelgrün() Colour { return darkGreen }
func Grün() Colour { return green }
func Schwarztürkis() Colour { return blackCyan }
func Dunkeltürkis() Colour { return darkCyan }
func Türkis() Colour { return cyan }
func Helltürkis() Colour { return lightCyan }
func Weißtürkis() Colour { return whiteCyan }
func Grelltürkis() Colour { return flashCyan }
func Grellblau() Colour { return lightBlue }
func Dunkelblau() Colour { return darkBlue }
func Preußischblau() Colour { return prussianBlue }
func Blau() Colour { return blue }
func Hellblau() Colour { return lightBlue }
func Weißblau() Colour { return whiteBlue }
func Grellviolett() Colour { return flashMagenta }
func Schwarzviolett() Colour { return blackMagenta }
func Dunkelviolett() Colour { return darkMagenta }
func Violett() Colour { return magenta }
func Hellviolett() Colour { return lightMagenta }
func Weißviolett() Colour { return whiteMagenta }
func Dunkelgrau() Colour { return darkGray }
func Grau() Colour { return gray }
func Hellgrau() Colour { return lightGray }
func Weißgrau() Colour { return whiteGray }
func Silber() Colour { return silver }
func Anthrazit() Colour { return blackGray }
func Weiß() Colour { return white }
func HellWhite() Colour { return lightWhite }

// RAL-Farben

func Grünbeige() Colour { return grünbeige }
func Beige() Colour { return beige }
func Sandgelb() Colour { return sandgelb }
func Signalgelb() Colour { return signalgelb }
func Goldgelb() Colour { return goldgelb }
func Honiggelb() Colour { return honiggelb }
func Maisgelb() Colour { return maisgelb }
func Narzissengelb() Colour { return narzissengelb }
func Braunbeige() Colour { return braunbeige }
func Zitronengelb() Colour { return zitronengelb }
func Perlweiß() Colour { return perlweiß }
func Elfenbein() Colour { return elfenbein }
func Hellelfenbein() Colour { return hellelfenbein }
func Schwefelgelb() Colour { return schwefelgelb }
func Safrangelb() Colour { return safrangelb }
func Zinkgelb() Colour { return zinkgelb }
func Graubeige() Colour { return graubeige }
func Olivgelb() Colour { return olivgelb }
func Rapsgelb() Colour { return rapsgelb }
func Verkehrsgelb() Colour { return verkehrsgelb }
func Ockergelb() Colour { return ockergelb }
func Leuchtgelb() Colour { return leuchtgelb }
func Currygelb() Colour { return currygelb }
func Melonengelb() Colour { return melonengelb }
func Ginstergelb() Colour { return ginstergelb }
func Dahliengelb() Colour { return dahliengelb }
func Pastellgelb() Colour { return pastellgelb }

func Gelborange() Colour { return gelborange }
func Rotorange() Colour { return rotorange }
func Blutorange() Colour { return blutorange }
func Pastellorange() Colour { return pastellorange }
func Reinorange() Colour { return reinorange }
func Leuchtorange() Colour { return leuchtorange }
func Leuchthellorange() Colour { return leuchthellorange }
func Hellrotorange() Colour { return hellrotorange }
func Verkehrsorange() Colour { return verkehrsorange }
func Signalorange() Colour { return signalorange }
func Tieforange() Colour { return tieforange }
func Lachsorange() Colour { return lachsorange }

func Feuerrot() Colour { return feuerrot }
func Signalrot() Colour { return signalrot }
func Karminrot() Colour { return karminrot }
func Rubinrot() Colour { return rubinrot }
func Purpurrot() Colour { return purpurrot }
func Weinrot() Colour { return weinrot }
func Schwarzrot() Colour { return schwarzrot }
func Oxidrot() Colour { return oxidrot }
func Braunrot() Colour { return braunrot }
func Beigerot() Colour { return beigerot }
func Tomatenrot() Colour { return tomatenrot }
func Altrosa() Colour { return altrosa }
func Hellrosa() Colour { return hellrosa }
func Korallenrot() Colour { return korallenrot }
func Rose() Colour { return rose }
func Erdbeerrot() Colour { return erdbeerrot }
func Verkehrsrot() Colour { return verkehrsrot }
func Lachsrot() Colour { return lachsrot }
func Leuchtrot() Colour { return leuchtrot }
func Leuchthellrot() Colour { return leuchthellrot }
func Himbeerrot() Colour { return himbeerrot }
func Orientrot() Colour { return orientrot }

func Rotlila() Colour { return rotlila }
func Rotmagenta() Colour { return rotmagenta }
func Erikamagenta() Colour { return erikamagenta }
func Bordeauxmagenta() Colour { return bordeauxmagenta }
func Blaulila() Colour { return blaulila }
func Verkehrspurpur() Colour { return verkehrspurpur }
func Purpurmagenta() Colour { return purpurmagenta }
func Signalviolett() Colour { return signalviolett }
func Pastellviolett() Colour { return pastellviolett }
func Telemagenta() Colour { return telemagenta }

func Violettblau() Colour { return violettblau }
func Grünblau() Colour { return grünblau }
func Ultramarinblau() Colour { return ultramarinblau }
func Saphirblau() Colour { return saphirblau }
func Schwarzblau() Colour { return schwarzblau }
func Signalblau() Colour { return signalblau }
func Brillantblau() Colour { return brillantblau }
func Graublau() Colour { return graublau }
func Azurblau() Colour { return azurblau }
func Enzianblau() Colour { return enzianblau }
func Stahlblau() Colour { return stahlblau }
func Lichtblau() Colour { return lichtblau }
func Kobaltblau() Colour { return kobaltblau }
func Taubenblau() Colour { return taubenblau }
func Himmelblau() Colour { return himmelblau }
func Verkehrsblau() Colour { return verkehrsblau }
func Türkisblau() Colour { return türkisblau }
func Capriblau() Colour { return capriblau }
func Ozeanblau() Colour { return ozeanblau }
func Wasserblau() Colour { return wasserblau }
func Nachtblau() Colour { return nachtblau }
func Fernblau() Colour { return fernblau }
func Pastellblau() Colour { return pastellblau }

func Patinagrün() Colour { return patinagrün }
func Smaragdgrün() Colour { return smaragdgrün }
func Laubgrün() Colour { return laubgrün }
func Olivgrün() Colour { return olivgrün }
func Blaugrün() Colour { return blaugrün }
func Moosgrün() Colour { return moosgrün }
func Grauoliv() Colour { return grauoliv }
func Flaschengrün() Colour { return flaschengrün }
func Braungrün() Colour { return braungrün }
func Tannengrün() Colour { return tannengrün }
func Grasgrün() Colour { return grasgrün }
func Resedagrün() Colour { return resedagrün }
func Schwarzgrün() Colour { return schwarzgrün }
func Schilfgrün() Colour { return schilfgrün }
func Gelboliv() Colour { return gelboliv }
func Schwarzoliv() Colour { return schwarzoliv }
func Cyangrün() Colour { return cyangrün }
func Maigrün() Colour { return maigrün }
func Gelbgrün() Colour { return gelbgrün }
func Weißgrün() Colour { return weißgrün }
func Chromoxidgrün() Colour { return chromoxidgrün }
func Blassgrün() Colour { return blassgrün }
func Braunoliv() Colour { return braunoliv }
func Verkehrsgrün() Colour { return verkehrsgrün }
func Farngrün() Colour { return farngrün }
func Opalgrün() Colour { return opalgrün }
func Lichtgrün() Colour { return lichtgrün }
func Kieferngrün() Colour { return kieferngrün }
func Minzgrün() Colour { return minzgrün }
func Signalgrün() Colour { return signalgrün }
func Minttürkis() Colour { return minttürkis }
func Pastelltürkis() Colour { return pastelltürkis }

func Fehgrau() Colour { return fehgrau }
func Silbergrau() Colour { return silbergrau }
func Olivgrau() Colour { return olivgrau }
func Moosgrau() Colour { return moosgrau }
func Signalgrau() Colour { return signalgrau }
func Mausgrau() Colour { return mausgrau }
func Beigegrau() Colour { return beigegrau }
func Khakigrau() Colour { return khakigrau }
func Grüngrau() Colour { return grüngrau }
func Zeltgrau() Colour { return zeltgrau }
func Eisengrau() Colour { return eisengrau }
func Basaltgrau() Colour { return basaltgrau }
func Braungrau() Colour { return braungrau }
func Schiefergrau() Colour { return schiefergrau }
func Anthrazitgrau() Colour { return anthrazitgrau }
func Schwarzgrau() Colour { return schwarzgrau }
func Umbragrau() Colour { return umbragrau }
func Betongrau() Colour { return betongrau }
func Graphitgrau() Colour { return graphitgrau }
func Granitgrau() Colour { return granitgrau }
func Steingrau() Colour { return steingrau }
func Blaugrau() Colour { return blaugrau }
func Kieselgrau() Colour { return kieselgrau }
func Zementgrau() Colour { return zementgrau }
func Gelbgrau() Colour { return gelbgrau }
func Lichtgrau() Colour { return lichtgrau }
func Platingrau() Colour { return platingrau }
func Staubgrau() Colour { return staubgrau }
func Achatgrau() Colour { return achatgrau }
func Quarzgrau() Colour { return quarzgrau }
func Fenstergrau() Colour { return fenstergrau }
func VerkehrsgrauA() Colour { return verkehrsgrauA }
func VerkehrsgrauB() Colour { return verkehrsgrauB }
func Seidengrau() Colour { return seidengrau }
func Telegrau1() Colour { return telegrau1 }
func Telegrau2() Colour { return telegrau2 }
func Telegrau4() Colour { return telegrau4 }

func Grünbraun() Colour { return grünbraun }
func Ockerbraun() Colour { return ockerbraun }
func Signalbraun() Colour { return signalbraun }
func Lehmbraun() Colour { return lehmbraun }
func Kupferbraun() Colour { return kupferbraun }
func Rehbraun() Colour { return rehbraun }
func Olivbraun() Colour { return olivbraun }
func Nussbraun() Colour { return nussbraun }
func Rotbraun() Colour { return rotbraun }
func Sepiabraun() Colour { return sepiabraun }
func Kastanienbraun() Colour { return kastanienbraun }
func Mahagonibraun() Colour { return mahagonibraun }
func Schokoladenbraun() Colour { return schokoladenbraun }
func Graubraun() Colour { return graubraun }
func Schwarzbraun() Colour { return schwarzbraun }
func Orangebraun() Colour { return orangebraun }
func Beigebraun() Colour { return beigebraun }
func Blassbraun() Colour { return blassbraun }
func Terrabraun() Colour { return terrabraun }

func Cremeweiß() Colour { return cremeweiß }
func Grauweiß() Colour { return grauweiß }
func Signalweiß() Colour { return signalweiß }
func Signalschwarz() Colour { return signalschwarz }
func Tiefschwarz() Colour { return tiefschwarz }
func Aluminiumweiß() Colour { return aluminiumweiß }
func Aluminiumgrau() Colour { return aluminiumgrau }
func Reinweiß() Colour { return reinweiß }
func Graphitschwarz() Colour { return graphitschwarz }
func Verkehrweiß() Colour { return verkehrsweiß }
func Verkehrschwarz() Colour { return verkehrsschwarz }
func Papyrusweiß() Colour { return papyrusweiß }
