package col

// (c) Christian Maurer   v. 200126 - license see µU.go

func HeadF() Colour            { return LightWhite() }
func HeadB() Colour            { return Blue() }
func HintF() Colour            { return LightWhite() }
func HintB() Colour            { return Magenta() }
func ErrorF() Colour           { return FlashYellow() }
func ErrorB() Colour           { return Red() }
func MenuF() Colour            { return LightWhite() }
func MenuB() Colour            { return Red() }
func MUF() Colour              { return new3 ("µuF",               0,  16,  64) }
func MUB() Colour              { return new3 ("µuB",             231, 238, 255) }

func Black() Colour            { return new3 ("Black",             0,   0,   0) }
func Brown() Colour            { return new3 ("Brown",            95,  53,  34) }
func BlackBrown() Colour       { return new3 ("BlackBrown",       30,  16,  12) }
func DarkBrown() Colour        { return new3 ("DarkBrown",        60,  33,  24) }
func MediumBrown() Colour      { return new3 ("MediumBrown",     149, 106,   0) }
func LightBrown() Colour       { return new3 ("LightBrown",      160,  88,  63) }
func WhiteBrown() Colour       { return new3 ("WhiteBrown",      221, 153, 106) }
func BrownWhite() Colour       { return new3 ("BrownWhite",      249, 202, 160) }
func Siena() Colour            { return new3 ("Siena",           153,  85,  42) }
func LightSiena() Colour       { return new3 ("LightSiena",      191, 127,  42) }
func RedBrown() Colour         { return new3 ("RedBrown",        170,  64,  64) }
func OliveBrown() Colour       { return new3 ("OliveBrown",      127, 127,   0) }
func LightOliveBrown() Colour  { return new3 ("LightOliveBrown", 170, 170,  85) }
func Umber() Colour            { return new3 ("umber",           149, 135,   0) }
func DarkOchre() Colour        { return new3 ("darkochre",       170, 127,  21) }
func Ochre() Colour            { return new3 ("ochre",           255, 170,  64) }
func LightOchre() Colour       { return new3 ("lightochre",      255, 191, 106) }
func RoseBrown() Colour        { return new3 ("rosebrown",       255, 191, 149) }
func LightBeige() Colour       { return new3 ("lightbeige",      234, 212, 170) }
func Beige1() Colour           { return new3 ("beige1",          212, 191, 149) }
func VeryLightBrown() Colour   { return new3 ("verylightbrown",  206, 170, 127) }
func BlackRed() Colour         { return new3 ("blackred",         46,  18,  26) }
func DarkRed() Colour          { return new3 ("darkred",          85,   0,   0) }
func Red() Colour              { return new3 ("red",             170,   0,   0) }
func FlashRed() Colour         { return new3 ("flashred",        255,   0,   0) }
func LightRed() Colour         { return new3 ("lightred",        255,  85,  85) }
func WhiteRed() Colour         { return new3 ("whitered",        255, 170, 170) }
func DarkRose() Colour         { return new3 ("darkrose",        234,   0, 127) }
func Rose1() Colour            { return new3 ("rose1",           255, 170, 170) }
func LightRose() Colour        { return new3 ("lightrose",       255, 191, 191) }
func PompejiRed() Colour       { return new3 ("pompejired",      187,  68,  68) }
func CinnabarRed() Colour      { return new3 ("cinnabarred",     238,  64,   0) }
func Carmine() Colour          { return new3 ("carmine",         125,   0,  42) }
func BrickRed() Colour         { return new3 ("brickred",        205,  63,  51) }
func FlashOrange() Colour      { return new3 ("rlashorange",     255, 127,   0) }
func DarkOrange() Colour       { return new3 ("darkorange",      221, 127,  68) }
func Orange() Colour           { return new3 ("orange",          255, 153,  51) }
func LightOrange() Colour      { return new3 ("lightorange",     255, 164,  31) }
func WhiteOrange() Colour      { return new3 ("whiteorange",     255, 170,   0) }
func BloodOrange() Colour      { return new3 ("bloodorange",     255, 112,  85) }
func FlashYellow() Colour      { return new3 ("flashyellow",     255, 255,   0) }
func DarkYellow() Colour       { return new3 ("darkyellow",      255, 187,   0) }
func Yellow() Colour           { return new3 ("yellow",          255, 255,  34) }
func LightYellow() Colour      { return new3 ("lightyellow",     255, 255, 102) }
func WhiteYellow() Colour      { return new3 ("whiteyellow",     255, 255, 153) }
func SandYellow() Colour       { return new3 ("sandyellow1",     234, 206, 127) }
func LemonYellow() Colour      { return new3 ("lemonyellow",     192, 255,  85) }
func FlashGreen() Colour       { return new3 ("flashgreen",        0, 255,   0) }
func BlackGreen() Colour       { return new3 ("blackgreen",        0,  51,   0) }
func VeryDarkGreen() Colour    { return new3 ("verydarkgreen",     0,  63,   0) }
func DarkGreen() Colour        { return new3 ("darkgreen",         0,  85,   0) }
func Green() Colour            { return new3 ("green",             0, 170,   0) }
func LightGreen() Colour       { return new3 ("lightgreen",       85, 255,  85) }
func WhiteGreen() Colour       { return new3 ("whitegreen",      170, 255, 170) }
func BirchGreen() Colour       { return new3 ("birchgreen",       42, 153,  42) }
func GrassGreen() Colour       { return new3 ("grassgreen",        0, 144,   0) }
func OliveGreen() Colour       { return new3 ("olivegreen",       85, 170,   0) }
func LightOliveGreen() Colour  { return new3 ("lightolivegreen", 170, 196,  85) }
func YellowGreen() Colour      { return new3 ("yellowgreen",     170, 255,  85) }
func MeadowGreen() Colour      { return new3 ("meadowgreen",     106, 212, 106) }
func BlackCyan() Colour        { return new3 ("blackcyan",         0,  51,  51) }
func DarkCyan() Colour         { return new3 ("darkcyan",          0,  85,  85) }
func Cyan() Colour             { return new3 ("cyan",              0, 170, 170) }
func LightCyan() Colour        { return new3 ("lightcyan",        85, 255, 255) }
func WhiteCyan() Colour        { return new3 ("whitecyan",       170, 255, 255) }
func FlashCyan() Colour        { return new3 ("flashcyan",         0, 255, 255) }
func FlashBlue() Colour        { return new3 ("flashblue",         0,   0, 255) }
func BlackBlue() Colour        { return new3 ("blackblue",         0,   0,  51) }
func PrussianBlue() Colour     { return new3 ("prussianblue",      0, 102, 170) }
func DarkBlue() Colour         { return new3 ("darkblue",          0,   0,  85) }
func Blue() Colour             { return new3 ("blue",              0,   0, 170) }
func LightBlue() Colour        { return new3 ("lightblue",        51, 102, 255) }
func WhiteBlue() Colour        { return new3 ("whiteblue",       170, 170, 255) }
func SkyLightBlue() Colour     { return new3 ("skylightblue",     85, 170, 255) }
func SkyBlue() Colour          { return new3 ("skyblue",           0, 170, 255) }
func GentianBlue() Colour      { return new3 ("gentianblue",       0,   0, 212) }
func Ultramarine() Colour      { return new3 ("ultramarine",      68,   0, 153) }
func BlackMagenta() Colour     { return new3 ("blackmagenta",     51,   0,  51) }
func DarkMagenta() Colour      { return new3 ("darkmagenta",      85,   0,  85) }
func Magenta() Colour          { return new3 ("magenta",         170,   0, 170) }
func LightMagenta() Colour     { return new3 ("lightmagenta",    255,  85, 255) }
func FlashMagenta() Colour     { return new3 ("flashmagenta",    255,   0, 255) }
func WhiteMagenta() Colour     { return new3 ("whitemagenta",    255, 170, 255) }
func Pink() Colour             { return new3 ("pink",            255,   0, 170) }
func DeepPink() Colour         { return new3 ("deeppink",        255,  17,  51) }
func BlackGray() Colour        { return new3 ("blackgray",        34,  34,  34) }
func DarkGray() Colour         { return new3 ("darkgray",         51,  51,  51) }
func Gray() Colour             { return new3 ("gray",             85,  85,  85) }
func LightGray() Colour        { return new3 ("lightgray",       136, 136, 136) }
func WhiteGray() Colour        { return new3 ("whitegray",       204, 204, 204) }
func Silver() Colour           { return new3 ("silver",          212, 212, 212) }
func LightSilver() Colour      { return new3 ("lightsilver",     234, 234, 234) }
func White() Colour            { return new3 ("white",           170, 170, 170) }
func LightWhite() Colour       { return new3 ("lightwhite",      255, 255, 255) }
