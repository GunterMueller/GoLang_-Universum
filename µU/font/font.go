package font

// (c) Christian Maurer   v. 230112 - license see µU.go

import
  . "µU/fontsize"
var (
  sign = []string {"r", "b", "i"}
  size = []string {"t", "s", "n", "b", "b", "b"}
  width  = []uint {  5,   6,   8,  12,  14,  16}
  height = []uint {  7,  10,  16,  24,  28,  32}
)

func init() {
  Name = []string {"tiny  ", //  6pt
                   "small ", //  8pt
                   "normal", // 10pt
                   "big   ", // 12pt
                   "large ", // 14pt
                   "huge  ", // 17pt
                  }
}

func code (f Font, s Size) string {
  x, y := "?", "?"
  if f < NFonts {
    x = sign[f]
  }
  if s < NSizes {
    y = size[s]
  }
  return x + y
}

func wd (s Size) uint {
  return width[s]
}

func ht (s Size) uint {
  return height[s]
}
