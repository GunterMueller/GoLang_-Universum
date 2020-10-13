package enum

// (c) Christian Maurer   v. 201007 - license see µU.go

const (
  UndefComparatio = uint8(iota)
  Positiv
  Komparativ
  Superlativ
  NComparationes
)
var (
  lComparatio = []string {"", "Positiv", "Komparativ", "Superlativ"}
  sComparatio = []string {"", "Pos.", "Komp.", "Sup."}
)

func init() {
  l[Comparatio] = lComparatio
  s[Comparatio] = sComparatio
}
