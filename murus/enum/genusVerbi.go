package enum

// (c) murus.org  v. 140522 - license see murus.go

const (
  UndefGenusVerbi = uint8(iota)
  Akt
  Pass
  NGeneraVerbi
)
var
  lGenusVerbi, sGenusVerbi []string =
  []string { "", "Aktiv", "Passiv" },
  []string { "", "Akt.", "Pass." }


func init() {
  l[GenusVerbi], s[GenusVerbi] = lGenusVerbi, sGenusVerbi
}
