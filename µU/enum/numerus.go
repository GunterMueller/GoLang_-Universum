package enum

// (c) Christian Maurer   v. 201010 - license see µU.go

const (
  UndefNumerus = uint8(iota)
  Sing
  Plur
  NNumeri
)
var (
  lNumerus = []string {"", "Singular", "Plural"}
  sNumerus = []string {"", "Sing.", "Plur."}
)

func init() {
  l[Numerus] = lNumerus
  s[Numerus] = sNumerus
}
