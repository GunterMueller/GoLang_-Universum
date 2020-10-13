package enum

// (c) Christian Maurer   v. 201007 - license see µU.go

const (
  UndefWortart = uint8(iota)
  Substantiv
  Adjektiv
  Pronomen
  Numerale
  Verb
  Adverb
  Präposition
  Konjunktion
  Interjektion
  NWortarten
)
var (
  lWortart = []string {"",
                       "Substantiv", "Adjektiv",
                       "Pronomen", "Numerale",
                       "Verb", "Adverb", "Präposition",
                       "Konjunktion", "Interjektion",
                      }
  sWortart = []string {"",
                       "Subst.", "Adj.",
                       "Pron.", "Num.",
                       "Verb", "Adv.", "Präp.",
                       "Konj.", "Interj.",
                      }
)

func init() {
  l[Wortart] = lWortart
  s[Wortart] = sWortart
}
