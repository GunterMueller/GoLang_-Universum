package enum

// (c) Christian Maurer   v. 201007 - license see µU.go

var (
  lSubject = []string {"keinFach",
                       "Deutsch", "Englisch", "Französisch", "Italienisch", "Spanisch",
                       "Polnisch", "Russisch", "Türkisch", "Japanisch", "Chinesisch",
                       "Latein", "Griechisch",
                       "Musik", "BildendeKunst", "DarstellendesSpiel",
                       "Politikwissenschaft", "Geschichte", "Geografie",
                       "Sozialwissenschaften", "Psychologie", "Philosophie", "Recht",
                       "Wirtschaftswissenschaft", "Pädagogik",
                         "RechnungswesenUndControlling", "Wirtschaft",
                       "Mathematik", "Physik", "Chemie", "Biologie", "Informatik",
                       "Physiktechnik", "Physiklabortechnik",
                         "Elektrotechnik", "RegenerativeEnergietechnik",
                       "Bautechnik", "Mechatronik", "Metalltechnik_Maschinenbau",
                       "Chemietechnik", "Chemielabortechnik",
                       "Biologietechnik", "Biologielabortechnik",
                         "Biotechnologie", "AgrartechnikMitBiologie",
                       "Gesundheit", "Ernährung", "Medizintechnik", "Umwelttechnik",
                       "Wirtschaftsinformatik", "TechnischeInformatik",
                         "Medizininformatik", "Informationstechnik",
                       "Medientechnik", "GestaltungsMedientechnik", "Gestaltung",
                       "Sport",
                      }
  sSubject = []string {"  ",
                       "de", "e ", "f ", "i ", "s ",
                       "p ", "r ", "t ", "j ", "c ",
                       "l ", "g ",
                       "mu", "ku", "ds",
                       "pw", "ge", "gg",
                       "sw", "ps", "ph", "re",
                       "ww", "pa", "rc", "wi",
                       "ma", "ph", "ch", "bi", "in",
                       "pt", "pt", "et", "rt",
                       "bt", "me", "mm",
                       "ct", "c",
                       "bt", "bt", "bt", "ab",
                       "gs", "er", "me", "ut",
                       "wi", "ti", "mi", "it",
                       "mt", "gm", "gt",
                       "sp",
                      }
)

func init() {
  l[Subject] = lSubject
  s[Subject] = sSubject
}
