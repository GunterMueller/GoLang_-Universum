package enum

// (c) Christian Maurer   v. 201011 - license see µU.go

var (
  laudio = []string {"           ",
                     "Klassik    ",
                     "Beat       ",
                     "Jazz       ",
                     "Folklore   ",
                     "Italien    ",
                     "Kinder     ",
                    }
)

func init() {
  l[Audio] = laudio
  s[Audio] = laudio
}
