package enum

// (c) murus.org  v. 140522 - license see murus.go

const (
  UndefLexCat = uint8(iota)
  Substantiv
  Adjektiv
  Pronomen
  Numerale
  Verb
  Adverb
  Präposition
  Konjunktion
  Interjektion
  NLexicalCategories
)
var
  lLexCat, sLexCat []string =
  []string { "", "Substantiv", "Adjektiv", "Pronomen", "Numerale",
             "Verb", "Adverb", "Präposition", "Konjunktion", "Interjektion" },
  []string { "", "Subst.", "Adj.", "Pron.", "Num.",
             "Verb", "Adv.", "Präp.", "Konj.", "Interj." }

func init() {
  l[LexicalCategory], s[LexicalCategory] = lLexCat, sLexCat
}
