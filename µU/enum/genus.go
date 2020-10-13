package enum

// (c) Christian Maurer   v. 201007 - license see µU.go

const (
  UndefGenus = uint8(iota)
  M
  F
  Ne
  NGenera
)
var (
  lGenus = []string {"", "masculinum", "femininum", "neutrum"}
  sGenus = []string {"", "m.", "f.", "n."}
)
func init() {
  l[Genus] = lGenus
  s[Genus] = sGenus
}
