package col

// (c) Christian Maurer   v. 171009 - license see µU.go

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

// XConsortium: rgb.txt,v 10.41 94/02/20 18:39:36 rws Exp
func XBlack() Colour { return new3 (  0,   0,   0) }
func XSnow() Colour { return new3 (255, 250, 250) }
func XSnow2() Colour { return new3 (238, 233, 233) }
func XSnow3() Colour { return new3 (205, 201, 201) }
func XSnow4() Colour { return new3 (139, 137, 137) }
func XGhostWhite() Colour { return new3 (248, 248, 255) }
func XWhiteSmoke() Colour { return new3 (245, 245, 245) }
func XGainsboro() Colour { return new3 (220, 220, 220) }
func XFloralWhite() Colour { return new3 (255, 250, 240) }
func XOldLace() Colour { return new3 (253, 245, 230) }
func XLinen() Colour { return new3 (250, 240, 230) }
func XAntiqueWhite() Colour { return new3 (250, 235, 215) }
func XAntiqueWhite1() Colour { return new3 (255, 239, 219) }
func XAntiqueWhite2() Colour { return new3 (238, 223, 204) }
func XAntiqueWhite3() Colour { return new3 (205, 192, 176) }
func XAntiqueWhite4() Colour { return new3 (139, 131, 120) }
func XPapayaWhip() Colour { return new3 (255, 239, 213) }
func XBlanchedAlmond() Colour { return new3 (255, 235, 205) }
func XBisque() Colour { return new3 (255, 228, 196) }
func XBisque1() Colour { return new3 (255, 228, 196) }
func XBisque2() Colour { return new3 (238, 213, 183) }
func XBisque3() Colour { return new3 (205, 183, 158) }
func XBisque4() Colour { return new3 (139, 125, 107) }
func XPeachPuff() Colour { return new3 (255, 218, 185) }
func XPeachPuff1() Colour { return new3 (255, 218, 185) }
func XPeachPuff2() Colour { return new3 (238, 203, 173) }
func XPeachPuff3() Colour { return new3 (205, 175, 149) }
func XPeachPuff4() Colour { return new3 (139, 119, 101) }
func XNavajoWhite() Colour { return new3 (255, 222, 173) }
func XNavajoWhite1() Colour { return new3 (255, 222, 173) }
func XNavajoWhite2() Colour { return new3 (238, 207, 161) }
func XNavajoWhite3() Colour { return new3 (205, 179, 139) }
func XNavajoWhite4() Colour { return new3 (139, 121,  94) }
func XMoccasin() Colour { return new3 (255, 228, 181) }
func XCornsilk() Colour { return new3 (255, 248, 220) }
func XCornsilk1() Colour { return new3 (255, 248, 220) }
func XCornsilk2() Colour { return new3 (238, 232, 205) }
func XCornsilk3() Colour { return new3 (205, 200, 177) }
func XCornsilk4() Colour { return new3 (139, 136, 120) }
func XIvory() Colour { return new3 (255, 255, 240) }
func XIvory1() Colour { return new3 (255, 255, 240) }
func XIvory2() Colour { return new3 (238, 238, 224) }
func XIvory3() Colour { return new3 (205, 205, 193) }
func XIvory4() Colour { return new3 (139, 139, 131) }
func XLemonChiffon() Colour { return new3 (255, 250, 205) }
func XLemonChiffon1() Colour { return new3 (255, 250, 205) }
func XLemonChiffon2() Colour { return new3 (238, 233, 191) }
func XLemonChiffon3() Colour { return new3 (205, 201, 165) }
func XLemonChiffon4() Colour { return new3 (139, 137, 112) }
func XSeashell() Colour { return new3 (255, 245, 238) }
func XSeashell1() Colour { return new3 (255, 245, 238) }
func XSeashell2() Colour { return new3 (238, 229, 222) }
func XSeashell3() Colour { return new3 (205, 197, 191) }
func XSeashell4() Colour { return new3 (139, 134, 130) }
func XHoneydew() Colour { return new3 (240, 255, 240) }
func XHoneydew1() Colour { return new3 (240, 255, 240) }
func XHoneydew2() Colour { return new3 (224, 238, 224) }
func XHoneydew3() Colour { return new3 (193, 205, 193) }
func XHoneydew4() Colour { return new3 (131, 139, 131) }
func XMintCream() Colour { return new3 (245, 255, 250) }
func XAzure() Colour { return new3 (240, 255, 255) }
func XAzure1() Colour { return new3 (240, 255, 255) }
func XAzure2() Colour { return new3 (224, 238, 238) }
func XAzure3() Colour { return new3 (193, 205, 205) }
func XAzure4() Colour { return new3 (131, 139, 139) }
func XAliceBlue() Colour { return new3 (240, 248, 255) }
func XLavender() Colour { return new3 (230, 230, 250) }
func XLavenderBlush() Colour { return new3 (255, 240, 245) }
func XLavenderBlush1() Colour { return new3 (255, 240, 245) }
func XLavenderBlush2() Colour { return new3 (238, 224, 229) }
func XLavenderBlush3() Colour { return new3 (205, 193, 197) }
func XLavenderBlush4() Colour { return new3 (139, 131, 134) }
func XMistyRose() Colour { return new3 (255, 228, 225) }
func XMistyRose1() Colour { return new3 (255, 228, 225) }
func XMistyRose2() Colour { return new3 (238, 213, 210) }
func XMistyRose3() Colour { return new3 (205, 183, 181) }
func XMistyRose4() Colour { return new3 (139, 125, 123) }
func XDarkSlateGray() Colour { return new3 ( 47,  79,  79) }
func XDarkSlateGray1() Colour { return new3 (151, 255, 255) }
func XDarkSlateGray2() Colour { return new3 (141, 238, 238) }
func XDarkSlateGray3() Colour { return new3 (121, 205, 205) }
func XDarkSlateGray4() Colour { return new3 ( 82, 139, 139) }
func XDimGray() Colour { return new3 (105, 105, 105) }
func XSlateGray() Colour { return new3 (112, 128, 144) }
func XSlateGray1() Colour { return new3 (198, 226, 255) }
func XSlateGray2() Colour { return new3 (185, 211, 238) }
func XSlateGray3() Colour { return new3 (159, 182, 205) }
func XSlateGray4() Colour { return new3 (108, 123, 139) }
func XLightSlateGray() Colour { return new3 (119, 136, 153) }
func XGray1() Colour { return new3 (190, 190, 190) }
func XLightGray() Colour { return new3 (211, 211, 211) }
func XDarkGray() Colour { return new3 (169, 169, 169) }
func XMidnightBlue() Colour { return new3 ( 25,  25, 112) }
func XNavyBlue() Colour { return new3 (  0,   0, 128) }
func XCornflowerBlue() Colour { return new3 (100, 149, 237) }
func XDarkSlateBlue() Colour { return new3 ( 72,  61, 139) }
func XSlateBlue() Colour { return new3 (106,  90, 205) }
func XSlateBlue1() Colour { return new3 (131, 111, 255) }
func XSlateBlue2() Colour { return new3 (122, 103, 238) }
func XSlateBlue3() Colour { return new3 (105,  89, 205) }
func XSlateBlue4() Colour { return new3 ( 71,  60, 139) }
func XMediumSlateBlue() Colour { return new3 (123, 104, 238) }
func XLightSlateBlue() Colour { return new3 (132, 112, 255) }
func XMediumBlue() Colour { return new3 (  0,   0, 205) }
func XRoyalBlue() Colour { return new3 ( 65, 105, 225) }
func XRoyalBlue1() Colour { return new3 ( 72, 118, 255) }
func XRoyalBlue2() Colour { return new3 ( 67, 110, 238) }
func XRoyalBlue3() Colour { return new3 ( 58,  95, 205) }
func XRoyalBlue4() Colour { return new3 ( 39,  64, 139) }
func XBlue() Colour { return new3 (  0,   0, 255) }
func XBlue1() Colour { return new3 (  0,   0, 255) }
func XBlue2() Colour { return new3 (  0,   0, 238) }
func XBlue3() Colour { return new3 (  0,   0, 205) }
func XBlue4() Colour { return new3 (  0,   0, 139) }
func XDarkBlue() Colour { return new3 (  0,   0, 139) }
func XDodgerBlue() Colour { return new3 ( 30, 144, 255) }
func XDodgerBlue2() Colour { return new3 ( 28, 134, 238) }
func XDodgerBlue3() Colour { return new3 ( 24, 116, 205) }
func XDodgerBlue4() Colour { return new3 ( 16,  78, 139) }
func XDeepSkyBlue() Colour { return new3 (  0, 191, 255) }
func XDeepSkyBlue2() Colour { return new3 (  0, 178, 238) }
func XDeepSkyBlue3() Colour { return new3 (  0, 154, 205) }
func XDeepSkyBlue4() Colour { return new3 (  0, 104, 139) }
func XSkyBlue() Colour { return new3 (135, 206, 235) }
func XSkyBlue2() Colour { return new3 (126, 192, 238) }
func XSkyBlue3() Colour { return new3 (108, 166, 205) }
func XSkyBlue4() Colour { return new3 ( 74, 112, 139) }
func XLightSkyBlue() Colour { return new3 (135, 206, 250) }
func XLightSkyBlue1() Colour { return new3 (176, 226, 255) }
func XLightSkyBlue2() Colour { return new3 (164, 211, 238) }
func XLightSkyBlue3() Colour { return new3 (141, 182, 205) }
func XLightSkyBlue4() Colour { return new3 ( 96, 123, 139) }
func XSteelBlue() Colour { return new3 ( 70, 130, 180) }
func XSteelBlue1() Colour { return new3 ( 99, 184, 255) }
func XSteelBlue2() Colour { return new3 ( 92, 172, 238) }
func XSteelBlue3() Colour { return new3 ( 79, 148, 205) }
func XSteelBlue4() Colour { return new3 ( 54, 100, 139) }
func XLightSteelBlue() Colour { return new3 (176, 196, 222) }
func XLightSteelBlue1() Colour { return new3 (202, 225, 255) }
func XLightSteelBlue2() Colour { return new3 (188, 210, 238) }
func XLightSteelBlue3() Colour { return new3 (162, 181, 205) }
func XLightSteelBlue4() Colour { return new3 (110, 123, 139) }
func XLightBlue() Colour { return new3 (173, 216, 230) }
func XLightBlue1() Colour { return new3 (191, 239, 255) }
func XLightBlue2() Colour { return new3 (178, 223, 238) }
func XLightBlue3() Colour { return new3 (154, 192, 205) }
func XLightBlue4() Colour { return new3 (104, 131, 139) }
func XPowderBlue() Colour { return new3 (176, 224, 230) }
func XPaleTurquoise() Colour { return new3 (175, 238, 238) }
func XPaleTurquoise1() Colour { return new3 (187, 255, 255) }
func XPaleTurquoise2() Colour { return new3 (174, 238, 238) }
func XPaleTurquoise3() Colour { return new3 (150, 205, 205) }
func XPaleTurquoise4() Colour { return new3 (102, 139, 139) }
func XDarkTurquoise() Colour { return new3 (  0, 206, 209) }
func XMediumTurquoise() Colour { return new3 ( 72, 209, 204) }
func XTurquoise() Colour { return new3 ( 64, 224, 208) }
func XTurquoise1() Colour { return new3 (  0, 245, 255) }
func XTurquoise2() Colour { return new3 (  0, 229, 238) }
func XTurquoise3() Colour { return new3 (  0, 197, 205) }
func XTurquoise4() Colour { return new3 (  0, 134, 139) }
func XCyan() Colour { return new3 (  0, 255, 255) }
func XCyan1() Colour { return new3 (  0, 255, 255) }
func XCyan2() Colour { return new3 (  0, 238, 238) }
func XCyan3() Colour { return new3 (  0, 205, 205) }
func XCyan4() Colour { return new3 (  0, 139, 139) }
func XDarkCyan() Colour { return new3 (  0, 139, 139) }
func XLightCyan() Colour { return new3 (224, 255, 255) }
func XLightCyan2() Colour { return new3 (209, 238, 238) }
func XLightCyan3() Colour { return new3 (180, 205, 205) }
func XLightCyan4() Colour { return new3 (122, 139, 139) }
func XCadetBlue() Colour { return new3 ( 95, 158, 160) }
func XCadetBlue1() Colour { return new3 (152, 245, 255) }
func XCadetBlue2() Colour { return new3 (142, 229, 238) }
func XCadetBlue3() Colour { return new3 (122, 197, 205) }
func XCadetBlue4() Colour { return new3 ( 83, 134, 139) }
func XMediumAquamarine() Colour { return new3 (102, 205, 170) }
func XAquamarine() Colour { return new3 (127, 255, 212) }
func XAquamarine1() Colour { return new3 (127, 255, 212) }
func XAquamarine2() Colour { return new3 (118, 238, 198) }
func XAquamarine3() Colour { return new3 (102, 205, 170) }
func XAquamarine4() Colour { return new3 ( 69, 139, 116) }
func XDarkGreen() Colour { return new3 (  0, 100,   0) }
func XDarkOliveGreen() Colour { return new3 ( 85, 107,  47) }
func XDarkOliveGreen1() Colour { return new3 (202, 255, 112) }
func XDarkOliveGreen2() Colour { return new3 (188, 238, 104) }
func XDarkOliveGreen3() Colour { return new3 (162, 205,  90) }
func XDarkOliveGreen4() Colour { return new3 (110, 139,  61) }
func XDarkSeaGreen() Colour { return new3 (143, 188, 143) }
func XDarkSeaGreen1() Colour { return new3 (193, 255, 193) }
func XDarkSeaGreen2() Colour { return new3 (180, 238, 180) }
func XDarkSeaGreen3() Colour { return new3 (155, 205, 155) }
func XDarkSeaGreen4() Colour { return new3 (105, 139, 105) }
func XSeaGreen() Colour { return new3 ( 46, 139,  87) }
func XSeaGreen1() Colour { return new3 ( 84, 255, 159) }
func XSeaGreen2() Colour { return new3 ( 78, 238, 148) }
func XSeaGreen3() Colour { return new3 ( 67, 205, 128) }
func XSeaGreen4() Colour { return new3 ( 46, 139,  87) }
func XMediumSeaGreen() Colour { return new3 ( 60, 179, 113) }
func XLightSeaGreen() Colour { return new3 ( 32, 178, 170) }
func XLightGreen() Colour { return new3 (144, 238, 144) }
func XPaleGreen() Colour { return new3 (152, 251, 152) }
func XPaleGreen1() Colour { return new3 (154, 255, 154) }
func XPaleGreen2() Colour { return new3 (144, 238, 144) }
func XPaleGreen3() Colour { return new3 (124, 205, 124) }
func XPaleGreen4() Colour { return new3 ( 84, 139,  84) }
func XSpringGreen() Colour { return new3 (  0, 255, 127) }
func XSpringGreen2() Colour { return new3 (  0, 238, 118) }
func XSpringGreen3() Colour { return new3 (  0, 205, 102) }
func XSpringGreen4() Colour { return new3 (  0, 139,  69) }
func XLawnGreen() Colour { return new3 (124, 252,   0) }
func XGreen1() Colour { return new3 (  0, 255,   0) }
func XGreen2() Colour { return new3 (  0, 238,   0) }
func XGreen3() Colour { return new3 (  0, 205,   0) }
func XGreen4() Colour { return new3 (  0, 139,   0) }
func XChartreuse() Colour { return new3 (127, 255,   0) }
func XChartreuse2() Colour { return new3 (118, 238,   0) }
func XChartreuse3() Colour { return new3 (102, 205,   0) }
func XChartreuse4() Colour { return new3 ( 69, 139,   0) }
func XMediumSpringGreen() Colour { return new3 ( 0, 250, 154) }
func XGreenYellow() Colour { return new3 (173, 255,  47) }
func XLimeGreen() Colour { return new3 ( 50, 205,  50) }
func XYellowGreen() Colour { return new3 (154, 205,  50) }
func XForestGreen() Colour { return new3 ( 34, 139,  34) }
func XOliveDrab() Colour { return new3 (107, 142,  35) }
func XOliveDrab1() Colour { return new3 (192, 255,  62) }
func XOliveDrab2() Colour { return new3 (179, 238,  58) }
func XOliveDrab3() Colour { return new3 (154, 205,  50) }
func XOliveDrab4() Colour { return new3 (105, 139,  34) }
func XDarkKhaki() Colour { return new3 (189, 183, 107) }
func XKhaki() Colour { return new3 (240, 230, 140) }
func XKhaki1() Colour { return new3 (255, 246, 143) }
func XKhaki2() Colour { return new3 (238, 230, 133) }
func XKhaki3() Colour { return new3 (205, 198, 115) }
func XKhaki4() Colour { return new3 (139, 134,  78) }
func XPaleGoldenrod() Colour { return new3 (238, 232, 170) }
func XLightGoldenrodYellow() Colour { return new3 (250, 250, 210) }
func XLightGoldenrod1() Colour { return new3 (255, 236, 139) }
func XLightGoldenrod2() Colour { return new3 (238, 220, 130) }
func XLightGoldenrod3() Colour { return new3 (205, 190, 112) }
func XLightGoldenrod4() Colour { return new3 (139, 129,  76) }
func XLightYellow () Colour { return new3 (255, 255, 224) }
func XLightYellow1() Colour { return new3 (255, 255, 224) }
func XLightYellow2() Colour { return new3 (238, 238, 209) }
func XLightYellow3() Colour { return new3 (205, 205, 180) }
func XLightYellow4() Colour { return new3 (139, 139, 122) }
func XYellow() Colour { return new3 (255, 255,   0) }
func XYellow1() Colour { return new3 (255, 255,   0) }
func XYellow2() Colour { return new3 (238, 238,   0) }
func XYellow3() Colour { return new3 (205, 205,   0) }
func XYellow4() Colour { return new3 (139, 139,   0) }
func XGold() Colour { return new3 (255, 215,   0) }
func XGold1() Colour { return new3 (255, 215,   0) }
func XGold2() Colour { return new3 (238, 201,   0) }
func XGold3() Colour { return new3 (205, 173,   0) }
func XGold4() Colour { return new3 (139, 117,   0) }
func XLightGoldenrod() Colour { return new3 (238, 221, 130) }
func XGoldenrod() Colour { return new3 (218, 165,  32) }
func XGoldenrod1() Colour { return new3 (255, 193,  37) }
func XGoldenrod2() Colour { return new3 (238, 180,  34) }
func XGoldenrod3() Colour { return new3 (205, 155,  29) }
func XGoldenrod4() Colour { return new3 (139, 105,  20) }
func XDarkGoldenrod() Colour { return new3 (184, 134,  11) }
func XDarkGoldenrod1() Colour { return new3 (255, 185,  15) }
func XDarkGoldenrod2() Colour { return new3 (238, 173,  14) }
func XDarkGoldenrod3() Colour { return new3 (205, 149,  12) }
func XDarkGoldenrod4() Colour { return new3 (139, 101,   8) }
func XRosyBrown() Colour { return new3 (188, 143, 143) }
func XRosyBrown1() Colour { return new3 (255, 193, 193) }
func XRosyBrown2() Colour { return new3 (238, 180, 180) }
func XRosyBrown3() Colour { return new3 (205, 155, 155) }
func XRosyBrown4() Colour { return new3 (139, 105, 105) }
func XIndianRed() Colour { return new3 (205,  92,  92) }
func XIndianRed1() Colour { return new3 (255, 106, 106) }
func XIndianRed2() Colour { return new3 (238,  99,  99) }
func XIndianRed3() Colour { return new3 (205,  85,  85) }
func XIndianRed4() Colour { return new3 (139,  58,  58) }
func XSaddleBrown() Colour { return new3 (139,  69,  19) }
func XSienna() Colour { return new3 (160,  82,  45) }
func XSienna1() Colour { return new3 (255, 130,  71) }
func XSienna2() Colour { return new3 (238, 121,  66) }
func XSienna3() Colour { return new3 (205, 104,  57) }
func XSienna4() Colour { return new3 (139,  71,  38) }
func XPeru() Colour { return new3 (205, 133,  63) }
func XBurlywood() Colour { return new3 (222, 184, 135) }
func XBurlywood1() Colour { return new3 (255, 211, 155) }
func XBurlywood2() Colour { return new3 (238, 197, 145) }
func XBurlywood3() Colour { return new3 (205, 170, 125) }
func XBurlywood4() Colour { return new3 (139, 115,  85) }
func XBeige1() Colour { return new3 (245, 245, 220) }
func XWheat() Colour { return new3 (245, 222, 179) }
func XWheat1() Colour { return new3 (255, 231, 186) }
func XWheat2() Colour { return new3 (238, 216, 174) }
func XWheat3() Colour { return new3 (205, 186, 150) }
func XWheat4() Colour { return new3 (139, 126, 102) }
func XSandyBrown() Colour { return new3 (244, 164,  96) }
func XTan() Colour { return new3 (210, 180, 140) }
func XTan1() Colour { return new3 (255, 165,  79) }
func XTan2() Colour { return new3 (238, 154,  73) }
func XTan3() Colour { return new3 (205, 133,  63) }
func XTan4() Colour { return new3 (139,  90,  43) }
func XChocolate() Colour { return new3 (210, 105,  30) }
func XChocolate1() Colour { return new3 (255, 127,  36) }
func XChocolate2() Colour { return new3 (238, 118,  33) }
func XChocolate3() Colour { return new3 (205, 102,  29) }
func XChocolate4() Colour { return new3 (139,  69,  19) }
func XFirebrick() Colour { return new3 (178,  34,  34) }
func XFirebrick1() Colour { return new3 (255,  48,  48) }
func XFirebrick2() Colour { return new3 (238,  44,  44) }
func XFirebrick3() Colour { return new3 (205,  38,  38) }
func XFirebrick4() Colour { return new3 (139,  26,  26) }
func XBrown0() Colour { return new3 (165,  42,  42) }
func XBrown1() Colour { return new3 (255,  64,  64) }
func XBrown2() Colour { return new3 (238,  59,  59) }
func XBrown3() Colour { return new3 (205,  51,  51) }
func XBrown4() Colour { return new3 (139,  35,  35) }
func XDarkSalmon() Colour { return new3 (233, 150, 122) }
func XSalmon() Colour { return new3 (250, 128, 114) }
func XSalmon1() Colour { return new3 (255, 140, 105) }
func XSalmon2() Colour { return new3 (238, 130,  98) }
func XSalmon3() Colour { return new3 (205, 112,  84) }
func XSalmon4() Colour { return new3 (139,  76,  57) }
func XLightSalmon() Colour { return new3 (255, 160, 122) }
func XLightSalmon1() Colour { return new3 (255, 160, 122) }
func XLightSalmon2() Colour { return new3 (238, 149, 114) }
func XLightSalmon3() Colour { return new3 (205, 129,  98) }
func XLightSalmon4() Colour { return new3 (139,  87,  66) }
func XOrange1() Colour { return new3 (255, 165,   0) }
func XOrange2() Colour { return new3 (238, 154,   0) }
func XOrange3() Colour { return new3 (205, 133,   0) }
func XOrange4() Colour { return new3 (139,  90,   0) }
func XDarkOrange() Colour { return new3 (255, 140,   0) }
func XDarkOrange1() Colour { return new3 (255, 127,   0) }
func XDarkOrange2() Colour { return new3 (238, 118,   0) }
func XDarkOrange3() Colour { return new3 (205, 102,   0) }
func XDarkOrange4() Colour { return new3 (139,  69,   0) }
func XCoral() Colour { return new3 (255, 127,  80) }
func XCoral1() Colour { return new3 (255, 114,  86) }
func XCoral2() Colour { return new3 (238, 106,  80) }
func XCoral3() Colour { return new3 (205,  91,  69) }
func XCoral4() Colour { return new3 (139,  62,  47) }
func XLightCoral() Colour { return new3 (240, 128, 128) }
func XTomato() Colour { return new3 (255,  99,  71) }
func XTomato2() Colour { return new3 (238,  92,  66) }
func XTomato3() Colour { return new3 (205,  79,  57) }
func XTomato4() Colour { return new3 (139,  54,  38) }
func XOrangeRed() Colour { return new3 (255,  69,   0) }
func XOrangeRed1() Colour { return new3 (255,  69,   0) }
func XOrangeRed2() Colour { return new3 (238,  64,   0) }
func XOrangeRed3() Colour { return new3 (205,  55,   0) }
func XOrangeRed4() Colour { return new3 (139,  37,   0) }
func XRed() Colour { return new3 (255,   0,   0) }
func XRed1() Colour { return new3 (255,   0,   0) }
func XRed2() Colour { return new3 (238,   0,   0) }
func XRed3() Colour { return new3 (205,   0,   0) }
func XRed4() Colour { return new3 (139,   0,   0) }
func XDarkRed() Colour { return new3 (139,   0,   0) }
func XHotPink() Colour { return new3 (255, 105, 180) }
func XHotPink1() Colour { return new3 (255, 110, 180) }
func XHotPink2() Colour { return new3 (238, 106, 167) }
func XHotPink3() Colour { return new3 (205,  96, 144) }
func XHotPink4() Colour { return new3 (139,  58,  98) }
func XDeepPink() Colour { return new3 (255,  20, 147) }
func XDeepPink2() Colour { return new3 (238,  18, 137) }
func XDeepPink3() Colour { return new3 (205,  16, 118) }
func XDeepPink4() Colour { return new3 (139,  10,  80) }
func XPink0() Colour { return new3 (255, 192, 203) }
func XPink1() Colour { return new3 (255, 181, 197) }
func XPink2() Colour { return new3 (238, 169, 184) }
func XPink3() Colour { return new3 (205, 145, 158) }
func XPink4() Colour { return new3 (139,  99, 108) }
func XLightPink() Colour { return new3 (255, 182, 193) }
func XLightPink1() Colour { return new3 (255, 174, 185) }
func XLightPink2() Colour { return new3 (238, 162, 173) }
func XLightPink3() Colour { return new3 (205, 140, 149) }
func XLightPink4() Colour { return new3 (139,  95, 101) }
func XPaleVioletRed() Colour { return new3 (219, 112, 147) }
func XPaleVioletRed1() Colour { return new3 (255, 130, 171) }
func XPaleVioletRed2() Colour { return new3 (238, 121, 159) }
func XPaleVioletRed3() Colour { return new3 (205, 104, 137) }
func XPaleVioletRed4() Colour { return new3 (139,  71,  93) }
func XMaroon() Colour { return new3 (176,  48,  96) }
func XMaroon1() Colour { return new3 (255,  52, 179) }
func XMaroon2() Colour { return new3 (238,  48, 167) }
func XMaroon3() Colour { return new3 (205,  41, 144) }
func XMaroon4() Colour { return new3 (139,  28,  98) }
func XMediumVioletRed() Colour { return new3 (199,  21, 133) }
func XVioletRed() Colour { return new3 (208,  32, 144) }
func XVioletRed1() Colour { return new3 (255,  62, 150) }
func XVioletRed2() Colour { return new3 (238,  58, 140) }
func XVioletRed3() Colour { return new3 (205,  50, 120) }
func XVioletRed4() Colour { return new3 (139,  34,  82) }
func XMagenta0() Colour { return new3 (255,   0, 255) }
func XMagenta1() Colour { return new3 (255,   0, 255) }
func XMagenta2() Colour { return new3 (238,   0, 238) }
func XMagenta3() Colour { return new3 (205,   0, 205) }
func XMagenta4() Colour { return new3 (139,   0, 139) }
func XDarkMagenta() Colour { return new3 (139,   0, 139) }
func XViolet() Colour { return new3 (238, 130, 238) }
func XPlum() Colour { return new3 (221, 160, 221) }
func XPlum1() Colour { return new3 (255, 187, 255) }
func XPlum2() Colour { return new3 (238, 174, 238) }
func XPlum3() Colour { return new3 (205, 150, 205) }
func XPlum4() Colour { return new3 (139, 102, 139) }
func XOrchid() Colour { return new3 (218, 112, 214) }
func XOrchid1() Colour { return new3 (255, 131, 250) }
func XOrchid2() Colour { return new3 (238, 122, 233) }
func XOrchid3() Colour { return new3 (205, 105, 201) }
func XOrchid4() Colour { return new3 (139,  71, 137) }
func XMediumOrchid() Colour { return new3 (186,  85, 211) }
func XMediumOrchid1() Colour { return new3 (224, 102, 255) }
func XMediumOrchid2() Colour { return new3 (209,  95, 238) }
func XMediumOrchid3() Colour { return new3 (180,  82, 205) }
func XMediumOrchid4() Colour { return new3 (122,  55, 139) }
func XDarkOrchid() Colour { return new3 (153,  50, 204) }
func XDarkOrchid1() Colour { return new3 (191,  62, 255) }
func XDarkOrchid2() Colour { return new3 (178,  58, 238) }
func XDarkOrchid3() Colour { return new3 (154,  50, 205) }
func XDarkOrchid4() Colour { return new3 (104,  34, 139) }
func XDarkViolet() Colour { return new3 (148,   0, 211) }
func XBlueViolet() Colour { return new3 (138,  43, 226) }
func XPurple() Colour { return new3 (160,  32, 240) }
func XPurple1() Colour { return new3 (155,  48, 255) }
func XPurple2() Colour { return new3 (145,  44, 238) }
func XPurple3() Colour { return new3 (125,  38, 205) }
func XPurple4() Colour { return new3 ( 85,  26, 139) }
func XMediumPurple() Colour { return new3 (147, 112, 219) }
func XMediumPurple1() Colour { return new3 (171, 130, 255) }
func XMediumPurple2() Colour { return new3 (159, 121, 238) }
func XMediumPurple3() Colour { return new3 (137, 104, 205) }
func XMediumPurple4() Colour { return new3 ( 93,  71, 139) }
func XThistle() Colour { return new3 (216, 191, 216) }
func XThistle1() Colour { return new3 (255, 225, 255) }
func XThistle2() Colour { return new3 (238, 210, 238) }
func XThistle3() Colour { return new3 (205, 181, 205) }
func XThistle4() Colour { return new3 (139, 123, 139) }
