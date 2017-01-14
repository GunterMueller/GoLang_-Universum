package enum

// (c) murus.org  v. 140522 - license see murus.go

var
  lSubject, sSubject =
  []string { "keinFach",
             "Deutsch", "Englisch", "Französisch", "Italienisch", "Spanisch",
             "Polnisch", "Russisch", "Türkisch", "Japanisch", "Chinesisch",
             "Latein", "Griechisch",
             "Musik", "BildendeKunst", "DarstellendesSpiel",
             "Politikwissenschaft", "Geschichte", "Geografie",
             "Sozialwissenschaften", "Psychologie", "Philosophie", "Recht",
             "Wirtschaftswissenschaft", "Pädagogik", "RechnungswesenUndControlling", "Wirtschaft",
             "Mathematik", "Physik", "Chemie", "Biologie", "Informatik",
             "Physiktechnik", "Physiklabortechnik", "Elektrotechnik", "RegenerativeEnergietechnik",
             "Bautechnik", "Mechatronik", "Metalltechnik_Maschinenbau",
             "Chemietechnik", "Chemielabortechnik",
             "Biologietechnik", "Biologielabortechnik", "Biotechnologie", "AgrartechnikMitBiologie",
             "Gesundheit", "Ernährung", "Medizintechnik", "Umwelttechnik",
             "Wirtschaftsinformatik", "TechnischeInformatik", "Medizininformatik", "Informationstechnik",
             "Medientechnik", "GestaltungsMedientechnik", "Gestaltung",
             "Sport" },
  []string { "  ",
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
             "sp" }
const
  NSubjects = 57


func init() {
  l[Subject], s[Subject] = lSubject, sSubject
  if NSubjects != uint8(len(lSubject)) { panic ("enum.NSubjects wrong") }
}
