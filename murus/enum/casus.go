package enum

// (c) murus.org  v. 140522 - license see murus.go

var
  lCasus, sCasus []string =
  []string { "", "Nominativ", "Genitiv", "Dativ", "Akkusativ", "Ablativ" },
  []string { "", "Nom.", "Gen.", "Dat.", "Akk.", "Abl." }
const (
  UndefCasus = uint8(iota)
  Nom
  Gen
  Dat
  Akk
  Abl
  NCasus
)

func init() {
  l[Casus], s[Casus] = lCasus, sCasus
}
