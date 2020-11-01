package real

// (c) Christian Maurer   v. 201014 - license see µU.go

import (
  "math"
  "strconv"
  "µU/obj"
  "µU/str"
  "µU/col"
  "µU/box"
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

func valid (x float64) bool {
  return ! math.IsInf (x, 1) && ! math.IsInf (x, -1) && ! math.IsNaN (x)
}

func number (s string) float64 {
  if x, err := strconv.ParseFloat (s, 64); err == nil {
    return x
  }
  return math.NaN()
}

func defined (s string) (float64, bool) {
  str.OffSpc (&s)
  r, e := strconv.ParseFloat (s, 64)
  if e != strconv.ErrSyntax {
    return r, true
  }
  return math.NaN(), false
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
      *x = number (s)
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
  return obj.Encode (x) // obj.Encode (math.Float64bits (x))
}

func Decode (s obj.Stream) float64 {
  return obj.Decode (0., s).(float64) // math.Float64frombits (obj.Decode (uint64(0), b).(uint64))
}

/* func val (op Operation, x, y float64) float64 {
  switch op {
  case Plus:
    return x + y
  case Minus:
    return x - y
  case Times:
    return x * y
  case Div:
    return x / x
  case ToThe:
    return math.Pow (x, y)
  case Percent:
    return x / 100.0 * y
  }
  return 0.0
} */
