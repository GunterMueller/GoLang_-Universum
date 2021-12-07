package col

// (c) Christian Maurer   v. 211127 - license see µU.go

func HeadF() Colour            { return LightWhite() }
func HeadB() Colour            { return Blue() }
func HintF() Colour            { return LightWhite() }
func HintB() Colour            { return Magenta() }
func ErrorF() Colour           { return FlashYellow() }
func ErrorB() Colour           { return Red() }
func MenuF() Colour            { return LightWhite() }
func MenuB() Colour            { return Red() }

func Black() Colour            { return new3n ("Black",             0,   0,   0) }
func Brown() Colour            { return new3n ("Brown",            95,  53,  34) }
func BlackBrown() Colour       { return new3n ("BlackBrown",       30,  16,  12) }
func DarkBrown() Colour        { return new3n ("DarkBrown",        60,  33,  24) }
func MediumBrown() Colour      { return new3n ("MediumBrown",     149, 106,   0) }
func LightBrown() Colour       { return new3n ("LightBrown",      160,  88,  63) }
func WhiteBrown() Colour       { return new3n ("WhiteBrown",      221, 153, 106) }
func BrownWhite() Colour       { return new3n ("BrownWhite",      249, 202, 160) }
func Siena() Colour            { return new3n ("Siena",           153,  85,  42) }
func LightSiena() Colour       { return new3n ("LightSiena",      191, 127,  42) }
func RedBrown() Colour         { return new3n ("RedBrown",        170,  64,  64) }
func OliveBrown() Colour       { return new3n ("OliveBrown",      127, 127,   0) }
func LightOliveBrown() Colour  { return new3n ("LightOliveBrown", 170, 170,  85) }
func Umber() Colour            { return new3n ("Umber",           149, 135,   0) }
func DarkOchre() Colour        { return new3n ("DarkOchre",       170, 127,  21) }
func Ochre() Colour            { return new3n ("Ochre",           255, 170,  64) }
func LightOchre() Colour       { return new3n ("LightOchre",      255, 191, 106) }
func RoseBrown() Colour        { return new3n ("RoseBrown",       255, 191, 149) }
func LightBeige() Colour       { return new3n ("LightBeige",      234, 212, 170) }
func Beige1() Colour           { return new3n ("Beige1",          212, 191, 149) }
func VeryLightBrown() Colour   { return new3n ("VeryLightBrown",  206, 170, 127) }
func BlackRed() Colour         { return new3n ("BlackRed",         46,  18,  26) }
func DarkRed() Colour          { return new3n ("DarkRed",          85,   0,   0) }
func Red() Colour              { return new3n ("Red",             170,   0,   0) }
func FlashRed() Colour         { return new3n ("FlashRed",        255,   0,   0) }
func LightRed() Colour         { return new3n ("LightRed",        255,  85,  85) }
func WhiteRed() Colour         { return new3n ("WhiteRed",        255, 170, 170) }
func DarkRose() Colour         { return new3n ("DarkRose",        234,   0, 127) }
func Rose1() Colour            { return new3n ("Rose1",           255, 170, 170) }
func LightRose() Colour        { return new3n ("LightRose",       255, 191, 191) }
func PompejiRed() Colour       { return new3n ("PompejiRed",      187,  68,  68) }
func CinnabarRed() Colour      { return new3n ("CinnabarRed",     238,  64,   0) }
func Carmine() Colour          { return new3n ("Carmine",         125,   0,  42) }
func BrickRed() Colour         { return new3n ("BrickRed",        205,  63,  51) }
func FlashOrange() Colour      { return new3n ("FlashOrange",     255, 127,   0) }
func DarkOrange() Colour       { return new3n ("DarkOrange",      221, 127,  68) }
func Orange() Colour           { return new3n ("Orange",          255, 153,  51) }
func LightOrange() Colour      { return new3n ("LightOrange",     255, 164,  31) }
func WhiteOrange() Colour      { return new3n ("WhiteOrange",     255, 170,   0) }
func BloodOrange() Colour      { return new3n ("BloodOrange",     255, 112,  85) }
func FlashYellow() Colour      { return new3n ("FlashYellow",     255, 255,   0) }
func DarkYellow() Colour       { return new3n ("DarkYellow",      255, 187,   0) }
func Yellow() Colour           { return new3n ("Yellow",          255, 255,  34) }
func LightYellow() Colour      { return new3n ("LightYellow",     255, 255, 102) }
func WhiteYellow() Colour      { return new3n ("WhiteYellow",     255, 255, 153) }
func SandYellow() Colour       { return new3n ("SandYellow1",     234, 206, 127) }
func LemonYellow() Colour      { return new3n ("LemonYellow",     192, 255,  85) }
func FlashGreen() Colour       { return new3n ("FlashGreen",        0, 255,   0) }
func BlackGreen() Colour       { return new3n ("BlackGreen",        0,  51,   0) }
func VeryDarkGreen() Colour    { return new3n ("VeryDarkGreen",     0,  63,   0) }
func DarkGreen() Colour        { return new3n ("DarkGreen",         0,  85,   0) }
func Green() Colour            { return new3n ("Green",             0, 170,   0) }
func LightGreen() Colour       { return new3n ("LightGreen",       85, 255,  85) }
func WhiteGreen() Colour       { return new3n ("WhiteGreen",      170, 255, 170) }
func BirchGreen() Colour       { return new3n ("BirchGreen",       42, 153,  42) }
func GrassGreen() Colour       { return new3n ("GrassGreen",        0, 144,   0) }
func OliveGreen() Colour       { return new3n ("OliveGreen",       85, 170,   0) }
func LightOliveGreen() Colour  { return new3n ("LightOliveGreen", 170, 196,  85) }
func YellowGreen() Colour      { return new3n ("YellowGreen",     170, 255,  85) }
func MeadowGreen() Colour      { return new3n ("MeadowGreen",     106, 212, 106) }
func BlackCyan() Colour        { return new3n ("BlackCyan",         0,  51,  51) }
func DarkCyan() Colour         { return new3n ("DarkCyan",          0,  85,  85) }
func Cyan() Colour             { return new3n ("Cyan",              0, 170, 170) }
func LightCyan() Colour        { return new3n ("LightCyan",        85, 255, 255) }
func WhiteCyan() Colour        { return new3n ("WhiteCyan",       170, 255, 255) }
func FlashCyan() Colour        { return new3n ("FlashCyan",         0, 255, 255) }
func FlashBlue() Colour        { return new3n ("FlashBlue",         0,   0, 255) }
func BlackBlue() Colour        { return new3n ("BlackBlue",         0,   0,  51) }
func PrussianBlue() Colour     { return new3n ("PrussianBlue",      0, 102, 170) }
func DarkBlue() Colour         { return new3n ("DarkBlue",          0,   0,  85) }
func Blue() Colour             { return new3n ("Blue",              0,   0, 170) }
func LightBlue() Colour        { return new3n ("LightBlue",        51, 102, 255) }
func WhiteBlue() Colour        { return new3n ("WhiteBlue",       170, 170, 255) }
func SkyLightBlue() Colour     { return new3n ("SkyLightBlue",     85, 170, 255) }
func SkyBlue() Colour          { return new3n ("SkyBlue",           0, 170, 255) }
func GentianBlue() Colour      { return new3n ("GentianBlue",       0,   0, 212) }
func Ultramarine() Colour      { return new3n ("Ultramarine",      68,   0, 153) }
func BlackMagenta() Colour     { return new3n ("BlackMagenta",     51,   0,  51) }
func DarkMagenta() Colour      { return new3n ("DarkMagenta",      85,   0,  85) }
func Magenta() Colour          { return new3n ("Magenta",         170,   0, 170) }
func LightMagenta() Colour     { return new3n ("LightMagenta",    255,  85, 255) }
func FlashMagenta() Colour     { return new3n ("FlashMagenta",    255,   0, 255) }
func WhiteMagenta() Colour     { return new3n ("WhiteMagenta",    255, 170, 255) }
func Pink() Colour             { return new3n ("Pink",            255,   0, 170) }
func DeepPink() Colour         { return new3n ("DeepPink",        255,  17,  51) }
func BlackGray() Colour        { return new3n ("BlackGray",        34,  34,  34) }
func DarkGray() Colour         { return new3n ("DarkGray",         51,  51,  51) }
func Gray() Colour             { return new3n ("Gray",             85,  85,  85) }
func LightGray() Colour        { return new3n ("LightGray",       136, 136, 136) }
func WhiteGray() Colour        { return new3n ("WhiteGray",       204, 204, 204) }
func Silver() Colour           { return new3n ("Silver",          212, 212, 212) }
func LightSilver() Colour      { return new3n ("LightSilver",     234, 234, 234) }
func White() Colour            { return new3n ("White",           170, 170, 170) }
func LightWhite() Colour       { return new3n ("LightWhite",      255, 255, 255) }
