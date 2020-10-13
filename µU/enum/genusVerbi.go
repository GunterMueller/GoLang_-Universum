package enum

// (c) Christian Maurer   v. 201007 - license see µU.go

const (
  UndefGenusVerbi = uint8(iota)
  Akt
  Pass
  NGeneraVerbi
)
var (
  lGenusVerbi = []string {"", "Aktiv", "Passiv"}
  sGenusVerbi = []string {"", "Akt.", "Pass."}
)

func init() {
  l[GenusVerbi] = lGenusVerbi
  s[GenusVerbi] = sGenusVerbi
}
