package cntry

// (c) Christian Maurer   v. 221213 - license see µU.go

import (
  . "µU/obj"
  "µU/str"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/errh"
  "µU/N"
  "µU/font"
  "µU/pbox"
  "µU/sel"
)
const
  length = 22
type
  id = byte; const (
              undefined = id(iota)
/* Europa */  Albanien; Andorra;
              Belgien; BosnienHerzegowina; Bulgarien
              Dänemark; Deutschland
              Estland
              Finnland; Frankreich
              Griechenland; Großbritannien
              Irland; Island; Italien
              Kroatien
              Lettland; Liechtenstein; Litauen; Luxemburg
              Malta; Mazedonien; Moldau; Monaco; Montenegro
              Niederlande; Norwegen; Österreich
              Polen; Portugal
              Rumänien; Russland
              SanMarino; Schweden; Schweiz; Serbien; Slowakei; Slowenien; Spanien
              Tschechien; Türkei
              Ukraine; Ungarn
              Vatikan
              Weißrussland
              Zypern
/* Afrika */  Ägypten; ÄquatorialGuinea; Äthiopien
              Algerien; Angola; Benin
              Botsuana; BurkinaFaso; Burundi
              Dschibuti
              Elfenbeinküste; Eritrea
              Gabun; Gambia; Ghana; Guinea; GuineaBissau
              Kamerun; KapVerde; Kenia; Komoren; Kongo; KongoDemRep
              Lesotho; Liberia; Libyen
              Madagaskar; Malawi; Mali; Marokko; Mauretanien; Mauritius; Mosambik
              Namibia; Niger; Nigeria
              Ruanda
              Sambia; SanTomePrincipe; Senegal; Seychellen; SierraLeone
                Simbabwe; Somalia; Südafrika; Südsudan; Sudan; Swasiland
              Tansania; Togo; Tschad; Tunesien
              Uganda
              Zentralafrika
/* Amerika */ Antigua; Argentinien
              Bahamas; Barbados; Belize; Bolivien; Brasilien
              Chile; CostaRica
              Dominica; DominikanRep
              Ecuador; ElSalvador
              Grenada; Guatemala; Guyana
              Haiti; Honduras
              Jamaika
              Kanada; StKittsNevis; Kolumbien; Kuba; StLucia
              Mexiko
              Nikaragua
              Panama; Paraguay; Peru
              Suriname
              TrinidadTobago
              Uruguay
              USA
              Venezuela; StVincent
/* Asien */   Afghanistan; Armenien; Aserbaidschan
              Bahrain; Bangladesch; Bhutan; Brunei
              China
              Georgien
              Indien; Indonesien; Irak; Iran; Israel
              Japan; Jemen; Jordanien
              Kambodscha; Kasachstan; Katar; Kirgisistan; Kuwait
              Laos; Libanon
              Malaysia; Malediven; Mongolei; Myanmar
              Nepal; Nordkorea
              Oman; Osttimor
              Pakistan; Palästina; Philippinen
              SaudiArabien; Singapur; SriLanka; Südkorea; Syrien
              Taiwan; Thailand; Tadschikistan; Turkmenistan
              Usbekistan
              VerArabEmirate; Vietnam
/* Australien und Ozeanien */
              Australien
              Cookinseln
              Fidschi
              Kiribati
              Marshallinseln; Mikronesien; Nauru
              Neuseeland; Niue
              Palau; PapuaNeuguinea
              Salomonen; Samoa
              Tonga; Tuvalu
              Vanuatu
              nNations)
type (
  attribut struct {
              iso,       // len 3
              tld,       // len 2
             name string
           prefix uint16
              car,
              ioc,
             fifa string // len 3
                  }
  country struct {
                 id
                 attribut
                 Format
          cF, cB col.Colour
                 font.Font
                 }
  Codes [2]byte // tld without trailing 0
)
var (
  bx = box.New()
  Font uint
  pbx = pbox.New()
  list []attribut
)

func new_() Country {
  x := new(country)
  x.Clr()
  x.Format = Long
  x.cF, x.cB = col.StartCols()
  return x
}

func (x *country) imp (Y any) *country {
  y, ok := Y.(*country)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *country) InContinent (c Continent) bool {
  switch c {
  case Europa:
    return Albanien <= x.id && x.id <= Zypern
  case Afrika:
    return Ägypten <= x.id && x.id <= Zentralafrika
  case Amerika:
    return Antigua <= x.id && x.id <= StVincent
  case Asien:
    return Afghanistan <= x.id && x.id <= Vietnam
  case Ozeanien:
    return Australien <= x.id && x.id <= Vanuatu
  }
  return false
}

func (x *country) Empty() bool {
  return x.id == undefined
}

func (x *country) Clr() {
  x.id = undefined
  x.attribut.iso = "   "
  x.attribut.tld = str.New (2)
  x.attribut.name = str.New (length)
  x.attribut.prefix = 0
  x.attribut.car = "   "
  x.attribut.ioc = "   "
  x.attribut.fifa = "   "
}

func (x *country) Copy (Y any) {
  y := x.imp (Y)
  x.id = y.id
  x.attribut.iso = y.attribut.iso
  x.attribut.tld = y.attribut.tld
  x.attribut.name = y.attribut.name
  x.attribut.prefix = y.attribut.prefix
  x.attribut.car = y.attribut.car
  x.attribut.ioc = y.attribut.ioc
  x.attribut.fifa = y.attribut.fifa
  x.Format = y.Format
  x.cF = y.cF
  x.cB = y.cB
}

func (x *country) Clone() any {
  y := new_()
  y.Copy(x)
  return y
}

func (x *country) Eq (Y any) bool {
  return x.id == x.imp (Y).id
}

func (x *country) Less (Y any) bool {
  return str.Less (x.attribut.name, x.imp (Y).attribut.name)
}

func (x *country) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *country) String() string {
  return x.attribut.name
}

func (x *country) TLD() string {
  return x.attribut.tld
}

func (x *country) TeX() string {
  s := str.TeX (x.attribut.name)
  return s
}

func (x *country) Defined (t string) bool {
  str.OffSpc (&t)
  if str.Empty (t) {
    x.Clr()
    return true
  }
  for c := 0; c < int(nNations); c++ {
    x.id = byte(c)
    x.attribut = list[c]
    switch x.Format {
    case Tld:
      if t == x.attribut.tld {
        x.name = list[c].name
        return true
      }
    case Long:
      if str.Sub0 (t, x.attribut.name) {
        x.tld = list[c].tld
        return true
      }
    }
  }
  return false
}

func (x *country) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *country) Cols() (col.Colour, col.Colour) {
  return x.cF, x.cB
}

func (x *country) GetFormat() Format {
  return x.Format
}

func (x *country) SetFormat (f Format) {
  if f < NFormats {
    x.Format = f
    switch x.Format {
    case Tld:
      bx.Wd (2)
    case Long:
      bx.Wd (length)
    default:
      bx.Wd (3)
    }
  }
}

func (x *country) Write (l, c uint) {
  bx.Colours (x.cF, x.cB)
  switch x.Format {
  case Tld:
    bx.Write (x.attribut.tld, l, c)
  case Long:
    bx.Clr (l, c)
    bx.Write (x.attribut.name, l, c)
  case Tel:
    N.Colours (x.cF, x.cB)
    N.Write (uint(x.attribut.prefix), l, c)
  case Car:
    bx.Write (x.attribut.car, l, c)
  case Ioc:
    bx.Write (x.attribut.ioc, l, c)
  case Fifa:
    bx.Write (x.attribut.fifa, l, c)
  }
}

func (x *country) Edit (l, c uint) {
  bx.Colours (x.cF, x.cB)
  var k uint
  for {
    switch x.Format {
    case Tld:
      k = 2
      bx.Edit (&x.attribut.tld, l, c)
    case Long:
      k = length
      bx.Edit (&x.attribut.name, l, c)
    default:
      return
    }
    if C, _ := kbd.LastCommand(); C == kbd.Search {
      n := uint(x.id)
      sel.Select (func (P, l, c uint, V, H col.Colour) {
                    x.attribut = list[id(P)]
                    x.Colours (V, H)
                    x.Write (l, c)
                  },
                  uint(nNations), scr.NLines(), k, &n, l, c, x.cB, x.cF)
      if id(n) == nNations {
        n = uint(x.id) // Nation unverändert
      } else {
        x.id = id(n)
      }
      x.attribut = list[x.id]
      break
    } else {
      var ok bool
      switch x.Format {
      case Tld:
        ok = x.Defined (x.attribut.tld)
      case Long:
//        f := x.Format
//        x.Format = Tld
        ok = x.Defined (x.attribut.name)
//        x.Format = f
      }
      if ok {
        break
      } else {
        errh.Error0("kein Land")
      }
    }
  }
  x.Write (l, c)
}

func (x *country) SetFont (f font.Font) {
  x.Font = f
}

func (x *country) Print (l, c uint) {
  pbx.SetFont (x.Font)
  switch x.Format {
  case Tld:
    pbx.Print (x.attribut.tld, l, c)
  case Long:
    pbx.Print (x.attribut.name, l, c)
  }
}

func (x *country) Codelen() uint {
  return 2
}

func (x *country) Encode() Stream {
  s := make (Stream, 2)
  s[0] = x.attribut.tld[0]
  s[1] = x.attribut.tld[1]
  return s
}

func (x *country) Decode (s Stream) {
  t := string(s)
  f := x.Format
  x.Format = Tld
  if ! x.Defined (t) {
    x.Clr()
  }
  x.SetFormat (f)
}

func (x *country) Name (t string) {
  for i := uint(0); i < uint(len(list)); i++ {
    if t == list[i].tld {
      x.name = list[i].name
      return
    }
  }
  x.name = str.New (length)
}

func def (n id, name, iso, tld string, prefix uint16, car, loc, fifa string) {
  list[n].tld = tld
  list[n].name = str.Lat1 (name)
  list[n].iso = iso
  list[n].prefix = prefix
  list[n].car = car
  list[n].ioc = loc
  list[n].fifa = fifa
}

func init() {
  bx.Wd (length)
  list = make ([]attribut, nNations)
                                                   // iso   tld   Tel    car    ioc   fifa
  def (undefined,          "                      ", "   ", "  ", 0   , "   ", "   ", "   ")

  def (Afghanistan,        "Afghanistan",            "AFG", "af", 93  , "AFG", "AFG", "AFG")
  def (Ägypten,            "Ägypten",                "EGY", "eg", 20  , "ET ", "EGY", "EGY")
  def (Albanien,           "Albanien",               "ALB", "al", 355 , "AL ", "ALB", "ALB")
  def (Algerien,           "Algerien",               "DZA", "dz", 213 , "DZ ", "ALG", "ALG")
  def (Andorra,            "Andorra",                "AND", "ad", 376 , "AND", "AND", "AND")
  def (Angola,             "Angola",                 "AGO", "ao", 244 , "ANG", "ANG", "ANG")
  def (Antigua,            "Antigua und Barbuda",    "ATG", "ag", 1268, "AG ", "ANT", "ATG")
  def (ÄquatorialGuinea,   "Äquatorial-Guinea",      "GNQ", "gq", 240 , "   ", "GEQ", "EGQ")
  def (Argentinien,        "Argentinien",            "ARG", "ar", 54  , "RA ", "ARG", "ARG")
  def (Armenien,           "Armenien",               "ARM", "am", 374 , "AR ", "ARM", "ARM")
  def (Aserbaidschan,      "Aserbaidschan",          "AZE", "az", 994 , "AZ ", "AZE", "AZE")
  def (Äthiopien,          "Äthiopien",              "ETH", "et", 251 , "ETH", "ETH", "ETH")
  def (Australien,         "Australien",             "AUS", "au", 61  , "AUS", "AUS", "AUS")

  def (Bahamas,            "Bahamas",                "BHS", "bs", 1242, "BS ", "BAH", "")
  def (Bahrain,            "Bahrain",                "BHR", "bh", 973 , "BRN", "BRN", "BHR")
  def (Bangladesch,        "Bangladesch",            "BGD", "bd", 880 , "BD ", "BAN", "BAN")
  def (Barbados,           "Barbados",               "BRB", "bb", 1246, "BDS", "BAR", "BRB")
  def (Belgien,            "Belgien",                "BEL", "be", 32  , "B  ", "BEL", "BEL")
  def (Belize,             "Belize",                 "BLZ", "bz", 501 , "BZ ", "BIZ", "")
  def (Benin,              "Benin",                  "BEN", "bj", 229 , "DY ", "BEN", "BEN")
  def (Bhutan,             "Bhutan",                 "BTN", "bt", 975 , "BHT", "BHU", "")
  def (Bolivien,           "Bolivien",               "BOL", "bo", 591 , "BOL", "BOL", "BOL")
  def (BosnienHerzegowina, "Bosnien u. Herzegowina", "BIH", "ba", 387 , "BIH", "BIH", "BIH")
  def (Botsuana,           "Botsuana",               "BWA", "bw", 267 , "RB ", "BOT", "BOT")
  def (Brasilien,          "Brasilien",              "BRA", "br", 55  , "BR ", "BRA", "BRA")
  def (Brunei,             "Brunei Darussalam",      "BRN", "bn", 673 , "BRU", "BRU", "BRU")
  def (Bulgarien,          "Bulgarien",              "BGR", "bg", 359 , "BG ", "BUL", "BUL")
  def (BurkinaFaso,        "Burkina Faso",           "BFA", "bf", 226 , "BF ", "BUR", "BFA")
  def (Burundi,            "Burundi",                "BDI", "bi", 257 , "RU ", "BDI", "BDI")

  def (Chile,              "Chile",                  "CHL", "cl", 56  , "RCH", "CHI", "CHI")
  def (China,              "China",                  "CHN", "cn", 86  , "RC ", "CHN", "CHN")
  def (Cookinseln,         "Cookinseln",             "COK", "ck", 682 , "NZ ", "COK", "COK")
  def (CostaRica,          "Costa Rica",             "CRI", "cr", 506 , "CR ", "CRC", "CRC")

  def (Dänemark,           "Dänemark",               "DNK", "dk", 45  , "DK ", "DEN", "DEN")
  def (Deutschland,        "Deutschland",            "DEU", "de", 49  , "D  ", "GER", "GER")
  def (Dominica,           "Dominica",               "DMA", "dm", 1767, "WD ", "DMA", "DMA")
  def (DominikanRep,       "Dominikan. Republik",    "DOM", "do", 1809, "DOM", "DOM", "DOM")
  def (Dschibuti,          "Dschibuti",              "DJI", "dj", 253 , "DSC", "DJI", "DJI")

  def (Ecuador,            "Ecuador",                "ECU", "ec", 593 , "EC ", "ECU", "ECU")
  def (ElSalvador,         "El Salvador",            "SLV", "sv", 503 , "ES ", "ESA", "SLV")
  def (Elfenbeinküste,     "Elfenbeinküste",         "CIV", "ci", 225 , "CI ", "CIV", "CIV")
  def (Eritrea,            "Eritrea",                "ERI", "er", 291 , "ER ", "ERI", "ERI")
  def (Estland,            "Estland",                "EST", "ee", 372 , "EST", "EST", "EST")

  def (Fidschi,            "Fidschi",                "FIJ", "fj", 679 , "FJI", "FIJ", "FIJ")
  def (Finnland,           "Finnland",               "FIN", "fi", 358 , "FIN", "FIN", "FIN")
  def (Frankreich,         "Frankreich",             "FRA", "fr", 33  , "F  ", "FRA", "FRA")

  def (Gabun,              "Gabun",                  "GAB", "ga", 241 , "G  ", "GAB", "GAB")
  def (Gambia,             "Gambia",                 "GMB", "gm", 220 , "WAG", "GAM", "GAM")
  def (Georgien,           "Georgien",               "GEO", "ge", 995 , "GE ", "GEO", "GEO")
  def (Ghana,              "Ghana",                  "GHA", "gh", 233 , "GH ", "GHA", "GHA")
  def (Grenada,            "Grenada",                "GRD", "gd", 1473, "WG ", "GRN", "GRN")
  def (Griechenland,       "Griechenland",           "GRC", "gr", 30  , "GR ", "GRE", "GRE")
  def (Großbritannien,     "Großbritannien",         "GBR", "uk", 44  , "GB ", "GBR", "ENG") // auch gb
//                                                                                    "NIR", "SCO", "WAL"
  def (Guatemala,          "Guatemala",              "GTM", "gt", 502 , "GCA", "GUA", "GUA")
  def (Guinea,             "Guinea",                 "GIN", "gn", 224 , "RG ", "GUI", "GUI")
  def (GuineaBissau,       "Guinea-Bissau",          "GNB", "gw", 245 , "GNB", "GBS", "GNB")
  def (Guyana,             "Guyana",                 "GUY", "gy", 592 , "GUY", "GUY", "GUY")

  def (Haiti,              "Haiti",                  "HTI", "ht", 509 , "RH ", "HAI", "HAI")
  def (Honduras,           "Honduras",               "HND", "hn", 504 , "HN ", "HON", "HON")

  def (Indien,             "Indien",                 "IND", "in", 91  , "IND", "IND", "IND")
  def (Indonesien,         "Indonesien",             "IDN", "id", 62  , "RI ", "INA", "IDN")
  def (Irak,               "Irak",                   "IRQ", "iq", 964 , "IRQ", "IRQ", "IRQ")
  def (Iran,               "Iran",                   "IRN", "ir", 98  , "IR ", "IRI", "IRN")
  def (Irland,             "Irland",                 "IRL", "ie", 353 , "IRL", "IRL", "IRL")
  def (Island,             "Island",                 "ISL", "is", 354 , "IS ", "ISL", "ISL")
  def (Israel,             "Israel",                 "ISR", "il", 972 , "IL ", "ISR", "ISR")
  def (Italien,            "Italien",                "ITA", "it", 39  , "I  ", "ITA", "ITA")

  def (Jamaika,            "Jamaika",                "JAM", "jm", 1876, "JA ", "JAM", "JAM")
  def (Japan,              "Japan",                  "JPN", "jp", 81  , "J  ", "JPN", "JPN")
  def (Jemen,              "Jemen",                  "YEM", "ye", 967 , "YAR", "YEM", "YEM")
  def (Jordanien,          "Jordanien",              "JOR", "jo", 962 , "JOR", "JOR", "JOR")

  def (Kambodscha,         "Kambodscha",             "KHM", "kh", 855 , "K  ", "CAM", "CAM")
  def (Kamerun,            "Kamerun",                "CMR", "cm", 237 , "TC ", "CMR", "CMR")
  def (Kanada,             "Kanada",                 "CAN", "ca", 1   , "CDN", "CAN", "CAN")
  def (KapVerde,           "Kap Verde",              "CPV", "cv", 238 , "CV ", "CPV", "CPV")
  def (Kasachstan,         "Kasachstan",             "KAZ", "kz", 7   , "KZ ", "KAZ", "KAZ")
  def (Katar,              "Katar",                  "QAT", "qa", 974 , "Q  ", "QAT", "QAT")
  def (Kenia,              "Kenia",                  "KEN", "ke", 254 , "EAK", "KEN", "KEN")
  def (Kirgisistan,        "Kirgisistan",            "KGZ", "kg", 996 , "KS ", "KGZ", "KGZ")
  def (Kiribati,           "Kiribati",               "KIR", "ki", 686 , "KI ", "KIR", "")
  def (Kolumbien,          "Kolumbien",              "COL", "co", 57  , "CO ", "COL", "COL")
  def (Komoren,            "Komoren",                "COM", "km", 269 , "COM", "COM", "COM")
  def (Kongo,              "Kongo",                  "COG", "cg", 242 , "RCB", "CGO", "CGO")
  def (KongoDemRep,        "Kongo, Dem.Rep.",        "COD", "cd", 243 , "CD ", "COD", "COD")
//def (Kosovo,             "Kosovo",                 "XXK", "  ", 381 , "   ", "   ", "")
  def (Kroatien,           "Kroatien",               "HRV", "hr", 385 , "HR ", "CRO", "CRO")
  def (Kuba,               "Kuba",                   "CUB", "cu", 53  , "C  ", "CUB", "CUB")
  def (Kuwait,             "Kuwait",                 "KWT", "kw", 965 , "KWT", "KUW", "KUW")

  def (Laos,               "Laos",                   "LAO", "la", 856 , "LAO", "LAO", "LAO")
  def (Lesotho,            "Lesotho",                "LSO", "ls", 266 , "LS ", "LES", "LES")
  def (Lettland,           "Lettland",               "LVA", "lv", 371 , "LV ", "LAT", "LVA")
  def (Libanon,            "Libanon",                "LBN", "lb", 961 , "RL ", "LIB", "LIB")
  def (Liberia,            "Liberia",                "LBR", "lr", 231 , "LB ", "LBR", "LBR")
  def (Libyen,             "Libyen",                 "LBY", "ly", 218 , "LAR", "LBA", "LBY")
  def (Liechtenstein,      "Liechtenstein",          "LIE", "li", 423 , "FL ", "LIE", "LIE")
  def (Litauen,            "Litauen",                "LTU", "lt", 370 , "LT ", "LTU", "LTU")
  def (Luxemburg,          "Luxemburg",              "LUX", "lu", 352 , "L  ", "LUX", "LUX")

  def (Madagaskar,         "Madagaskar",             "MDG", "mg", 261 , "RM ", "MAD", "MAD")
  def (Malawi,             "Malawi",                 "MWI", "mw", 265 , "MW ", "MAW", "MWI")
  def (Malaysia,           "Malaysia",               "MYS", "my", 60  , "MAL", "MAS", "MAS")
  def (Malediven,          "Malediven",              "MDV", "mv", 960 , "MV ", "MDV", "MDV")
  def (Mali,               "Mali",                   "MLI", "ml", 223 , "RMM", "MLI", "MLI")
  def (Malta,              "Malta",                  "MLT", "mt", 356 , "M  ", "MLT", "MLT")
  def (Marokko,            "Marokko",                "MAR", "ma", 212 , "MA ", "MAR", "MAR")
  def (Marshallinseln,     "Marshallinseln",         "MHL", "mh", 692 , "MH ", "MHL", "---")
  def (Mauretanien,        "Mauretanien",            "MRT", "mr", 222 , "RIM", "MTN", "MTN")
  def (Mauritius,          "Mauritius",              "MUS", "mu", 230 , "MS ", "MRI", "MRI")
  def (Mazedonien,         "Mazedonien",             "MKD", "mk", 389 , "MK ", "MKD", "MKD")
  def (Mexiko,             "Mexiko",                 "MEX", "mx", 52  , "MEX", "MEX", "MEX")
  def (Mikronesien,        "Mikronesien",            "FSM", "fm", 691 , "FSM", "FSM", "---")
  def (Moldau,             "Moldau",                 "MDA", "md", 373 , "MD ", "MDA", "MDA")
  def (Monaco,             "Monaco",                 "MCO", "mc", 377 , "MC ", "MON", "---")
  def (Mongolei,           "Mongolei",               "MNG", "mn", 976 , "MGL", "MGL", "MNG")
  def (Montenegro,         "Montenegro",             "MNE", "me", 382 , "MNE", "MNE", "MNE") // auch yu
  def (Mosambik,           "Mosambik",               "MOZ", "mz", 258 , "MOC", "MOZ", "MOZ")
  def (Myanmar,            "Myanmar",                "MMR", "mm", 95  , "MYA", "MYA", "MYA")

  def (Namibia,            "Namibia",                "NAM", "na", 264 , "NAM", "NAM", "NAM")
  def (Nauru,              "Nauru",                  "NRU", "nr", 674 , "NAU", "NRU", "---")
  def (Nepal,              "Nepal",                  "NPL", "np", 977 , "NEP", "NEP", "NEP")
  def (Neuseeland,         "Neuseeland",             "NZL", "nz", 64  , "NZ ", "NZL", "NZL")
  def (Nikaragua,          "Nicaragua",              "NIC", "ni", 505 , "NIC", "NCA", "NAC")
  def (Niederlande,        "Niederlande",            "NLD", "nl", 31  , "NL ", "NED", "NED")
  def (Niger,              "Niger",                  "NER", "ne", 227 , "NIG", "NIG", "NIG")
  def (Nigeria,            "Nigeria",                "NGA", "ng", 234 , "WAN", "NGR", "NGA")
  def (Niue,               "Niue",                   "NIU", "nu", 683 , "NZ ", "   ", "---")
  def (Nordkorea,          "Nordkorea",              "PRK", "kp", 850 , "KP ", "PRK", "PRK")
  def (Norwegen,           "Norwegen",               "NOR", "no", 47  , "N  ", "NOR", "NOR")

  def (Oman,               "Oman",                   "OMN", "om", 968 , "OM ", "OMA", "OMA")
  def (Österreich,         "Österreich",             "AUT", "at", 43  , "A  ", "AUT", "AUT")
  def (Osttimor,           "Timor-Leste",            "TLS", "tl", 670 , "TL ", "TLS", "TLS") // auch tp

  def (Pakistan,           "Pakistan",               "PAK", "pk", 92  , "PK ", "PAK", "PAK")
  def (Palästina,          "Palästin.Autonomiegeb.", "PSE", "ps", 970 , "   ", "PLE", "PLE")
  def (Palau,              "Palau",                  "PLW", "pw", 680 , "PAL", "PLW", "---")
  def (Panama,             "Panama",                 "PAN", "pa", 507 , "PA ", "PAN", "PAN")
  def (PapuaNeuguinea,     "Papua Neuguinea",        "PNG", "pg", 675 , "PNG", "PNG", "PNG")
  def (Paraguay,           "Paraguay",               "PRY", "py", 595 , "PY ", "PAR", "PAR")
  def (Peru,               "Peru",                   "PER", "pe", 51  , "PE ", "PER", "PER")
  def (Philippinen,        "Philippinen",            "PHL", "ph", 63  , "RP ", "PHI", "PHI")
  def (Polen,              "Polen",                  "POL", "pl", 48  , "PL ", "POL", "POL")
  def (Portugal,           "Portugal",               "PRT", "pt", 351 , "P  ", "POR", "POR")

  def (Ruanda,             "Ruanda",                 "RWA", "rw", 250 , "RWA", "RWA", "RWA")
  def (Rumänien,           "Rumänien",               "ROU", "ro", 40  , "R  ", "ROU", "ROU")
  def (Russland,           "Russische Föderation",   "RUS", "ru", 7   , "RUS", "RUS", "RUS") // auch su

  def (Salomonen,          "Salomonen",              "SLB", "sb", 677 , "SOL", "SOL", "SOL")
  def (Sambia,             "Sambia",                 "ZMB", "zm", 260 , "Z  ", "ZAM", "ZAM")
  def (Samoa,              "Samoa",                  "WSM", "ws", 685 , "WS ", "SAM", "SAM")
  def (SanMarino,          "San Marino",             "SMR", "sm", 378 , "RSM", "SMR", "SMR")
  def (SanTomePrincipe,    "Sao Tome und Principe",  "STP", "st", 239 , "STP", "STP", "STP")
  def (SaudiArabien,       "Saudi-Arabien",          "SAU", "sa", 966 , "KSA", "KSA", "KSA")
  def (Schweden,           "Schweden",               "SWE", "se", 46  , "S  ", "SWE", "SWE")
  def (Schweiz,            "Schweiz",                "CHE", "ch", 41  , "CH ", "SUI", "SUI")
  def (Senegal,            "Senegal",                "SEN", "sn", 221 , "SN ", "SEN", "SEN")
  def (Serbien,            "Serbien",                "SRB", "rs", 381 , "SRB", "SRB", "SRB") // auch yu
  def (Seychellen,         "Seychellen",             "SYC", "sc", 248 , "SY ", "SEY", "SEY")
  def (SierraLeone,        "Sierra Leone",           "SLE", "sl", 232 , "WAL", "SLE", "SLE")
  def (Simbabwe,           "Simbabwe",               "ZWE", "zw", 263 , "ZW ", "ZIM", "ZIM")
  def (Singapur,           "Singapur",               "SGP", "sg", 65  , "SGP", "SIN", "SIN")
  def (Slowakei,           "Slowakei",               "SVK", "sk", 421 , "SK ", "SVK", "SVK")
  def (Slowenien,          "Slowenien",              "SVN", "si", 386 , "SLO", "SLO", "SVN")
  def (Somalia,            "Somalia",                "SOM", "so", 252 , "SP ", "SOM", "SOM")
  def (Spanien,            "Spanien",                "ESP", "es", 34  , "E  ", "ESP", "ESP")
  def (SriLanka,           "Sri Lanka",              "LKA", "lk", 93  , "CL ", "SRI", "SRI")
  def (StKittsNevis,       "St. Kitts und Nevis",    "KNA", "kn", 1869, "KAN", "SKN", "SKN")
  def (StLucia,            "St. Lucia",              "LCA", "lc", 1758, "WL ", "LCA", "LCA")
  def (StVincent,          "St. Vincent Grenadinen", "VCT", "vc", 1784, "WV ", "VIN", "VIN")
  def (Südafrika,          "Südafrika",              "ZAF", "za", 27  , "ZA ", "RSA", "RSA")
  def (Sudan,              "Sudan",                  "SDN", "sd", 249 , "SUD", "SUD", "SDN")
  def (Südsudan,           "Südsudan",               "   ", "ss", 292 , "SSD", "   ", "SSD")
  def (Südkorea,           "Südkorea",               "KOR", "kr", 82  , "ROK", "KOR", "KOR")
  def (Suriname,           "Suriname",               "SUR", "sr", 597 , "SME", "SUR", "SUR")
  def (Swasiland,          "Swasiland",              "SWZ", "sz", 268 , "SD ", "SWZ", "SWZ")
  def (Syrien,             "Syrien",                 "SYR", "sy", 963 , "SYR", "SYR", "SYR")

  def (Tadschikistan,      "Tadschikistan",          "TJK", "tj", 992 , "TJ ", "TJK", "TJK")
  def (Taiwan,             "Taiwan",                 "TWN", "tw", 886 , "RC ", "TPE", "TPE")
  def (Tansania,           "Tansania",               "TZA", "tz", 255 , "EAT", "TAN", "TAN")
  def (Thailand,           "Thailand",               "THA", "th", 66  , "THA", "THA", "THA")
  def (Togo,               "Togo",                   "TGO", "tg", 228 , "RT ", "TOG", "TOG")
  def (Tonga,              "Tonga",                  "TON", "to", 676 , "TON", "TGA", "TGA")
  def (TrinidadTobago,     "Trinidad und Tobago",    "TTO", "tt", 1868, "TT ", "TRI", "TRI")
  def (Tschad,             "Tschad",                 "TCD", "td", 235 , "TCD", "CHA", "CHA")
  def (Tschechien,         "Tschechische Republik",  "CZE", "cz", 420 , "CZ ", "CZE", "CZE")
  def (Tunesien,           "Tunesien",               "TUN", "tn", 216 , "TN ", "TUN", "TUN")
  def (Türkei,             "Türkei",                 "TUR", "tr", 90  , "TR ", "TUR", "TUR")
  def (Turkmenistan,       "Turkmenistan",           "TKM", "tm", 993 , "TM ", "TKM", "TKM")
  def (Tuvalu,             "Tuvalu",                 "TUV", "tv", 688 , "TUV", "TUV", "---")

  def (Uganda,             "Uganda",                 "UGA", "ug", 256 , "EAU", "UGA", "UGA")
  def (Ukraine,            "Ukraine",                "UKR", "ua", 380 , "UA ", "UKR", "UKR")
  def (Ungarn,             "Ungarn",                 "HUN", "hu", 36  , "H  ", "HUN", "HUN")
  def (Uruguay,            "Uruguay",                "URY", "uy", 598 , "ROU", "URU", "URU")
  def (USA,                "Ver. Staaten v.Amerika", "USA", "us", 1   , "USA", "USA", "USA")
  def (Usbekistan,         "Usbekistan",             "UZB", "uz", 998 , "UZB", "UZB", "UZB")

  def (Vanuatu,            "Vanuatu",                "VUT", "vu", 678 , "VU ", "VAN", "VAN")
  def (Vatikan,            "Vatikanstadt",           "VAT", "va", 379 , "V  ", "   ", "---")
  def (Venezuela,          "Venezuela",              "VEN", "ve", 58  , "YV ", "VEN", "VEN")
  def (VerArabEmirate,     "Ver. Arabische Emirate", "ARE", "ae", 971 , "UAE", "UAE", "UAE")
  def (Vietnam,            "Vietnam",                "VNM", "vn", 84  , "VN ", "VIE", "VIE")

  def (Weißrussland,       "Weißrussland",           "BLR", "by", 375 , "BY ", "BLR", "BLR")

  def (Zentralafrika,      "Zentralafrikan. Rep.",   "CAF", "cf", 236 , "RZA", "CAF", "CTA")
  def (Zypern,             "Zypern",                 "CYP", "cy", 357 , "CY ", "CYP", "CYP")
/*
                           "Anguilla",               "   ", "ai"                    , "AIA")
                           "Niederl. Antillen",      "   ", "an"
                           "Antarktis",              "   ", "aq"
                           "Amerikan. Samoa",        "   ", "as"                    , "ASA")
                           "Aruba",                  "   ", "aw"                    , "ARU")
                           "Åland",                  "   ", "ax"
                           "Bermuda",                "   ", "bm"                    , "BER")
                           "Bouvet Island no",       "   ", "bv"
                           "Cocos (Keeling) Insel",  "   ", "cc"
                           "Christmas Insel",        "   ", "cx"
                           "Curaçao",                "   ",                         , "CUW")
                           "Europäische Union",      "   ", "eu"
                           "Falkland Insel",         "   ", "fk"
                           "Faröer",                 "   ", "fo"                    . "FRO")
                           "Franz. Guiana",          "   ", "gf"
                           "Guernse",                "   ", "gg"
                           "Gibraltar",              "   ", "gi"
                           "Grönland",               "   ", "gl"
                           "Guadeloupe etc.",        "   ", "gp"
                      "SouthGeorgia+Sandwich Insel", "   ", "gs"
                           "Guam",                   "   ", "gu"                    , "GUM")
                           "HongKong",               "   ", "hk"                    , "HKG")
                           "Heard u.McDonald Insel", "   ", "hm"
                           "Isle of Man",            "   ", "im"
                           "Brit.Terr.im Ind.Ozean", "   ", "io"
                           "Jersey",                 "   ", "je"
                           "Cayman Islands",         "   ", "ky"                    , "CAY")
                           "Macau",                  "   ", "mo"                    , "MAC")
                           "Nord Mariana Insel",     "   ", "mp"
                           "Martinique",             "   ", "mq" // fr
                           "Montserrat",             "   ", "ms"                    , "MSR")
                           "Neu Kaledonien",         "   ", "nc" // fr              , "NCL")
                           "Norfolk Island",         "   ", "nf"
                           "Franz. Polynesien",      "   ", "pf" // fr
                           "Saint-Pierre u.Miquelon","   ", "pm"
                           "Pitcairn Islands",       "   ", "pn"
                           "Puerto Rico",            "   ", "pr"                    , "PUR")
                           "Réunion",                "   ", "re" // fr
                           "Saint Helena",           "   ", "sh"
                           "Svalbard+JanMayen Insel","   ", "sj" // no
                           "Tahiti",                                                , "TAH")
                           "Turks and Caicos Insel", "   ", "tc"                    , "TCA")
                           "Franz.südl.u.antarkt.L." "   ", "tf"
                           "Tokelau",                "   ", "tk"
                           "British Virgin Islands", "   ", "vg"                    , "VGB")
                           "Virgin Islands",         "   ", "vi"                    , "VIR")
                           "Wallis and Futuna",      "   ", "wf" // fr
                           "Mayotte",                "   ", "yt" // fr
*/
}
