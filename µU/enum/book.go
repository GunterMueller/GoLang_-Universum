package enum

// (c) Christian Maurer   v. 210415 - license see µU.go

var (
  lbook = []string {"            ",
                    "Prosa       ",
                    "Klassik     ",
                    "Roman       ",
                    "Histo-Roman ",
                    "Rom-Roman   ",
                    "ItalienRoman",
                    "Krimi       ",
                    "Rom-Krimi   ",
                    "ItalienKrimi",
                    "Kunst       ",
                    "Ägypten     ",
                    "Etrurien    ",
                    "Sachbuch    ",
                    "Theater     ",
                    "Kinderbuch  ",
                    "Märchen     ",
                   }
  sbook = []string {"  ",
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
  l[Book] = lbook
  s[Book] = sbook
}
