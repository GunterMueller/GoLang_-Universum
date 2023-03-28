package robi

// (c) Christian Maurer   v. 230309 - license see µU.go
//
// Spezikationen siehe robi/rob/def.go

func M() uint { return m() }

func LinksDrehen() { linksDrehen() }
func RechtsDrehen() { rechtsDrehen() }
func InLinkerObererEcke() bool { return inLinkerObererEcke() }
func VorRand() bool { return vorRand() }
func Laufen() { laufen() }
func Zurücklaufen() { zurücklaufen() }
func Leer() bool { return leer() }
func NachbarLeer() bool { return nachbarLeer() }
func AnzahlKlötze() uint { return anzahlKlötze() }
func HatKlötze() bool { return hatKlötze() }
func Ablegen() { ablegen() }
func Aufnehmen() { aufnehmen() }
func Geschoben() bool { return geschoben() }
func Schießen() { schießen() }
func Markieren() { markieren() }
func Entmarkieren() { entmarkieren() }
func Markiert() bool { return markiert() }
func NachbarMarkiert() bool { return nachbarMarkiert() }
func VorMauer() bool { return vorMauer() }
func Zumauern() { zumauern() }
func Entmauern() { entmauern() }

func Laden (s ...string) { laden(s...) }
func Editieren() { editieren() }
func ProtokollSchalten (ein bool) { protokollSchalten (ein) }
func SokobanSchalten (ein bool) { sokobanSchalten(ein) }
func Ausgeben (n uint) { ausgeben(n) }
func Eingabe() uint { return eingabe() }
func FehlerMelden (s string, n uint) { fehlerMelden(s,n) }
func HinweisAusgeben (s string, n uint) { hinweisAusgeben(s,n) }
func Fertig() { fertig() }

func Pos() (uint, uint) { return pos() }
func Set (x, y uint) { set(x,y) }
