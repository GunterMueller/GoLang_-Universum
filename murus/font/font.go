package font

// (c) Christian Maurer   v. 170328 - license see murus.go

var (
  sign = []string { "r", "b", "s", "i" }
  size = []string { "t", "s", "n", "b", "h" }
  width = []uint { 5, 6, 8, 12, 16 }
  height = []uint { 7, 10, 16, 24, 32 }
)

func init() {
  Name = []string { "tiny  ", "small ", "normal", "big   ", "huge  " }
}

func code (f Font, s Size) string {
  x, y:= "?", "?"
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
