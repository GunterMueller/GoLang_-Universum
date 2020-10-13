package enum

// (c) Christian Maurer   v. 201007 - license see µU.go

const (
  UndefCasus = uint8(iota)
  Nom
  Gen
  Dat
  Akk
  Abl
  NCasus
)
var (
  lCasus = []string {"", "Nominativ", "Genitiv", "Dativ", "Akkusativ", "Ablativ"}
  sCasus = []string {"", "Nom.",      "Gen.",    "Dat.",  "Akk.",      "Abl."}
)

func init() {
  l[Casus] = lCasus
  s[Casus] = sCasus
}
