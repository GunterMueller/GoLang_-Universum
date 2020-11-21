package real

// (c) Christian Maurer   v. 201112 - license see µU.go

import (
  "math"
  "strconv"
  "µU/obj"
  "µU/str"
  "µU/col"
  "µU/box"
  "µU/nat"
  "µU/errh"
)
var
  bx = box.New()

func init() {
//  bx.SetNumerical()
//  setFormat()
}

func NDigits (x float64) uint {
  return 1 + uint(math.Floor (math.Log (math.Abs (x)) / math.Ln10))
}

func finite (x float64) bool {
  return ! math.IsInf (x, 1) && ! math.IsInf (x, -1) && ! math.IsNaN (x)
}

func val (s string) float64 {
  if x, e := strconv.ParseFloat (s, 64); e == nil {
    return x
  }
  return math.NaN()
}

func eq (x, y float64) bool {
  return x == y // TODO
}

func defined (s string) (float64, bool) {
// number = ["-"]digit{digit}[ "."digit{digit}]["e"|"E"["-"]digit{digit}]
  str.OffBytes (&s, ' ')
  x := 0.
  n, t, p, l := nat.DigitSequences (s)
  if uint(len(s)) > p[n-1]+l[n-1] {
    return 0, false
  }
  if n == 0 || n > 3 {
    return 0., false
  }
  if n == 2 {
    c := s[p[1] - 1]
    if c != '.' && c != ',' {
      return 0, false
    }
  }
  n0, _ := nat.Natural (t[0])
  x = float64(n0)
  if n >= 2 {
    n1, _ := nat.Natural (t[1])
    z := 1.
    for i := uint(0); i < l[1]; i++ {
      z *= 10
    }
    x += float64(n1) / z
  }
  if n == 3 {
    n2, _ := nat.Natural (t[2])
    c := s[p[2] - 1]
    if c != 'e' && c != 'E' {
      return 0, false
    }
    z := 1.
    for i := uint(0); i < n2; i++ {
      z *= 10
    }
    x *= z
  }
  if s[0] == '-' { x = -x }
  return x, true
}

func string_ (x float64) string {
  s := strconv.FormatFloat (x, 'f', 2, 64)
  str.OffSpc (&s)
  return s
}

func colours (f, b col.Colour) {
  bx.Colours (f, b)
}

func write (x float64, l, c uint) {
  bx.Write (String (x), l, c)
}

func edit (x *float64, l, c uint) {
  s := String (*x)
  for {
    bx.Edit (&s, l, c)
    _, ok := str.Pos (s, 'e')
    _, ok1 := str.Pos (s, 'E')
    if ! ok && ! ok1 {
      *x = val (s)
      if ! math.IsNaN (*x) {
        break
      }
    } else {
    }
    errh.Error0Pos ("keine Zahl", l + 1, c)
  }
  Write (*x, l, c)
}

func Codelen() uint {
  return 8
}

func Encode (x float64) obj.Stream {
  return obj.Encode (x)
}

func Decode (s obj.Stream) float64 {
  return obj.Decode (0., s).(float64)
}

func realStarted (s string) (float64, uint, bool) {
  n := uint(len(s))
  for k := n; k > 0; k-- {
    if x, ok := defined (s[:k]); ok {
      return x, k, true
    }
  }
  return 0., n, false
}
