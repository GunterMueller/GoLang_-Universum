package enum

// (c) murus.org  v. 140522 - license see murus.go

const (
  UndefModus = uint8(iota)
  Ind
  Konj
  NModi
)
var
  lModus, sModus []string =
  []string { "", "Indikativ", "Konjunktiv", },
  []string { "", "Ind.", "Konj.", }


func init() {
  l[Modus], s[Modus] = lModus, sModus
}
