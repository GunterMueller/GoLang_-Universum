package enum

// (c) Christian Maurer   v. 201007 - license see µU.go

var (
  lbookC = []string {"           ",
                     "Prosa      ",
                     "Klassik    ",
                     "Roman      ",
                     "Histo-Roman",
                     "Rom-Roman  ",
                     "Ital. Roman",
                     "Krimi      ",
                     "Rom-Krimi  ",
                     "Ital. Krimi",
                     "Kunst      ",
                     "Ägypten    ",
                     "Etrurien   ",
                     "Sachbuch   ",
                     "Theater    ",
                     "Kinderbuch ",
                     "Märchen    ",
                    }
  sbookC = []string {"  ",
                     "pr",
                     "kl",
                     "r ",
                     "rh",
                     "rr",
                     "ri",
                     "k ",
                     "kr",
                     "ki",
                     "ku",
                     "ae",
                     "et",
                     "sb",
                     "th",
                     "kb",
                     "m ",
                    }
)

func init() {
  l[BookC] = lbookC
  s[BookC] = sbookC
}
