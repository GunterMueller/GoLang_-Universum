package ppm

// (c) Christian Maurer   v. 220130 - license see µU.go

import (
 // "os/exec"
  "µU/ker"
  "µU/obj"
  "µU/str"
  "µU/scr" // only for PPMWrite
  "µU/col"
  "µU/pseq"
  "µU/files"
  "µU/char"
//  "µU/prt"
)
const
  suffix = ".ppm"
type
  image struct {
               string
          w, h uint
            cs [][]col.Colour
               }
var (
  ppmheader string
  ppmHeaderLength uint
)

func string_(n uint) string {
  if n == 0 { return "0" }
  var s string
  for s = ""; n > 0; n /= 10 {
    s = string(n % 10 + '0') + s
  }
  return s
}

func number (s obj.Stream) (uint, int) {
  n := uint(0)
  i := 0
  for char.IsDigit (s[i]) { i++ }
  for j := 0; j < i; j++ {
    n = 10 * n + uint(s[j] - '0')
  }
  return n, i
}

func new_() Image {
  im := new(image)
  im.string = ""
  im.w, im.h = 0, 0
  return im
}

func ppmHeaderData (s obj.Stream) (uint, uint, uint, int) {
  p := string(s[:2]); if p != "P6" { ker.Panic ("wrong ppm-header: " + p) }
  i := 3
  if s[i] == '#' {
    for {
      i++
      if s[i] == byte(10) {
        i++
        break
      }
    }
  }
  w, dw := number (s[i:])
  i += dw + 1
  h, dh := number (s[i:])
  i += dh + 1
  m, dm := number (s[i:])
  i += dm
  return w, h, m, i + 1
}

func (im *image) Load (n string) {
  if str.Empty (n) {
    ker.Panic ("Loaded called with empty string as parameter")
  }
  str.OffSpc (&n)
  im.string = n
  filename := n + suffix
  if ! files.IsFile (filename) {
    ker.Panic ("file " + n + ".ppm is not in the actual directory")
  }
  k := pseq.Length (filename)
  s := make(obj.Stream, k)
  file := pseq.New (s)
  file.Name (filename)
  s = file.Get().(obj.Stream)
  file.Fin()
  w, h, _, j := ppmHeaderData (s)
  im.w, im.h = w, h
  i := 4 * uint(2) + 2
  c := col.New()
  im.cs = make([][]col.Colour, h)
  for y := uint(0); y < h; y++ {
    im.cs[y] = make([]col.Colour, w)
    for x := uint(0); x < w; x++ {
      im.cs[y][x] = col.New()
      c.Decode (s[j:j+3])
      im.cs[y][x].Copy (c)
      i += 3
      j += 3
    }
  }
}

func (x *image) Size() (uint, uint) {
  return x.w, x.h
}

func (im *image) Colours() [][]col.Colour {
  return im.cs
}

func ppmHeader (w, h uint) string {
  s := "P6 " + string_(w) + " " + string_(h) + " 255" + string(byte(10))
  ppmheader = s
  ppmHeaderLength = uint(len(s))
  return s
}

// absolutely criminal experimental version
func (x *image) Store (n string) {
  if str.Empty (n) { return }
  str.OffSpc (&n)
  w, h := x.w, x.h
  s := scr.Encode (0, 0, w, h)
  s = append (obj.Stream(ppmHeader (w, h)), s[2*4:]...)
  file := pseq.New (s)
  file.Name (n + suffix)
  file.Clr()
  file.Put (s)
  file.Fin()
}

func (im *image) Print (x, y int) {
/*/
  im.Store (n)
  e := exec.Command (prt.PrintCommand, "-o", "fit-to-page", n + suffix).Run()
  if e != nil { ker.Panic (e.Error()) }
/*/
}
