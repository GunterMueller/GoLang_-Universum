package enum

// (c) Christian Maurer   v. 181221 - license see µU.go

import
  "µU/str"
var (
  lEng, sEng[]string =
  []string { "",
    "003",
    "078",
    "089",
    "103",
    "110",
    "111",
    "117",
    "120",
    "144",
    "151",
    "187",
    "194",
    "402",
    "216",
    "260",
    "795",
  },
  lComp
  NEngines = uint8(len(lComp))
)
func init() {
  for i:= 1; i < len(lEng); i++ {
    p, _:= str.Pos (lEng[i], ',')
    sEng[i] = str.Part (lEng[i], 0, p)
  }
  l[Composer], s[Composer] = lEng, sEng
}
