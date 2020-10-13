package enum

// (c) Christian Maurer   v. 201011 - license see µU.go

var (
  laudioC = []string {"           ",
                      "Klassik    ",
                      "Beat       ",
                      "Jazz       ",
                      "Folklore   ",
                      "Italien    ",
                      "Kinder     ",
                     }
)

func init() {
  l[AudioC] = laudioC
  s[AudioC] = laudioC
}
