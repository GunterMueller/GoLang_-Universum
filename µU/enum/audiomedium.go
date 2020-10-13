package enum

// (c) Christian Maurer   v. 201011 - license see µU.go

var (
  laudioMedium = []string {"   ",
                           "SP ",
                           "LP ",
                           "CD ",
                           "DVD",
                           "BR ",
                          }
)

func init() {
  l[AudioMedium] = laudioMedium
  s[AudioMedium] = laudioMedium
}
