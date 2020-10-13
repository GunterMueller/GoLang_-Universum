package enum

// (c) Christian Maurer   v. 201007 - license see µU.go

const (
  UndefModus = uint8(iota)
  Ind
  Konj
  NModi
)
var (
  lModus = []string {"", "Indikativ", "Konjunktiv"}
  sModus = []string {"", "Ind.", "Konj."}
)

func init() {
  l[Modus] = lModus
  s[Modus] = sModus
}
