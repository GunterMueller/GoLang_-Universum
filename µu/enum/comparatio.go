package enum

// (c) Christian Maurer   v. 140525 - license see µu.go

const (
  UndefComparatio = uint8(iota)
  Positiv
  Komparativ
  Superlativ
  NComparationes
)
var
  lComparatio, sComparatio []string =
  []string { "", "Positiv", "Komparativ", "Superlativ" },
  []string { "", "Pos.", "Komp.", "Sup." }

func init() {
  l[Comparatio], s[Comparatio] = lComparatio, sComparatio
}
