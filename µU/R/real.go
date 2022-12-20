package R

// (c) Christian Maurer   v. 221213 - license see µU.go

import (
  "math"
  "strconv"
  "µU/obj"
  "µU/char"
  "µU/str"
  "µU/col"
  "µU/box"
  "µU/N"
  "µU/errh"
)
const (
  M = 64
  A = 12
  bias = 1023
  w = 12
)
var (
  bx = box.New()
  dp = uint(2)
  errtxt = ""
  errpos = uint(0)
)

func init() {
  bx.Wd (w)
  bx.SetNumerical()
}

func isZero (x float64) bool {
  return eq (x, 0.)
}

func log10 (x float64) float64 {
  return math.Log (x) / math.Ln10
}

func NDigits (x float64) uint {
  return 1 + uint(math.Floor (log10 (math.Abs (x))))
}

func finite (x float64) bool {
  return ! math.IsInf (x, 1) && ! math.IsInf (x, -1) && ! math.IsNaN (x)
}

func eq (x, y float64) bool {
  return math.Abs (x - y) <= Epsilon
}

func pow10 (n uint) uint {
  if n == 0 {
    return 1
  }
  return 10 * pow10 (n - 1)
}

func fail (s string, e uint) {
  errtxt = s + " at position"
  errpos = e
  return
  println (errtxt, errpos) // for testing
}

func float (s string) (float64, uint, bool) {
// float64 = ["-"] digit {digit} [ [","|"."] [ digit {digit} ] ] ["e" ["-"] digit {digit} ]
  str.OffSpc (&s)
  f := NaN()
  x := 0.
  k, t, p, l := N.DigitSequences (s)
  if k == 0 || k > 3 {
    fail ("s invalid", 0)
    return f, 0, false
  }
  j := p[k-1]+l[k-1]
  if uint(len(s)) > j {
    fail ("invalid characters", j)
    return f, 0, false
  }
  for i := uint(0); i < l[0]; i++ {
    x *= 10
    x += float64(t[0][i] - byte('0'))
  }
  if p[0] > 0 {
    if p[0] != 1 {
      fail ("invalid character after the first one", 1)
      return f, 0, false
    }
    if s[0] == '-' {
      x = -x
    } else {
      fail ("invalid first character", 0)
      return f, 0, false
    }
  }
  if k == 1 {
    return x, p[0] + l[0], true
  }
// k >= 2:
  var nodot, noexp bool
  c := s[p[1] - 1]
  if c == '.' || c == ',' {
    z := 0.1
    for i := uint(0); i < l[1]; i++ {
      c := float64(t[1][i] - byte('0'))
      if x < 0 { c = -c }
      x += z * c
      z /= 10
    }
  } else {
    nodot = true
  }
  i := p[k-2]+l[k-2]
  if k == 3 && s[i] != 'e' {
    fail ("invalid character '" + string(s[i]) + "' instead of 'e'", i)
    return f, 0, false
  }
  if s[i] == 'e' {
    c = s[i+1]
    if c != '-' && ! char.IsDigit(c) {
      fail ("invalid character '" + string(c) + "' after 'e'", i + 1)
      return f, 0, false
    }
    ae := uint(0)
    exp := 0
    expneg := false
    i := p[k-1] - 1
    c := s[i]
    if c == '-' {
      expneg = true
    }
    if a, ok := N.Natural (t[k-1]); ! ok {
      fail ("invalid exponent", p[k-1])
      return f, 0, false
    } else {
      ae = a
      exp = int(ae)
    }
    if expneg {
      exp = -exp
      x /= float64(pow10 (ae))
    } else {
      x *= float64(pow10 (ae))
    }
  } else {
    noexp = true
  }
  if nodot && noexp {
    return x, p[0] + l[0], true
  }
  return x, p[k-1] + l[k-1], true
}

func integer (x float64) bool {
  _, b := math.Modf (x)
  return eq (b, 0.)
}

func wd (n uint) {
  bx.Wd (n)
}

func setFormat (n uint) {
  dp = n
}

func string_(x float64) string {
  s := strconv.FormatFloat (x, 'g', 16, 64)
  t := obj.Stream(s)
  b := make(obj.Stream, 0)
  if x < 0 {
    b = append (b, '-')
    t = t[1:]
  }
  n, ss, ps, _ := N.DigitSequences (s)
  s0 := ""; if s[0] == '-' { s0 = "-" }
  if n >= 2 {
    k := uint(len(ss[1]))
    if dp < k {
      ss[1] = ss[1][:dp]
    }
    if dp > k {
      ss[1] += str.Const ('0', dp - k)
    }
  }
  switch n {
  case 1:
    s = s0 + ss[0]
  case 2:
    s = s0 + ss[0] + "." + ss[1]
  case 3:
    s2 := ""; if s[ps[2]-1] == '-' { s2 = "-" }
    s = s0 + ss[0] + "." + ss[1] + "e" + s2 + ss[2]
  }
  return s
}

func colours (f, b col.Colour) {
  bx.Colours (f, b)
}

func write (x float64, l, c uint) {
  bx.Write (string_(x), l, c)
}

func edit (x *float64, l, c uint) {
  s := string_(*x)
  for {
    bx.Edit (&s, l, c)
    if _, _, ok := float (s); ok {
      break
    } else {
      errh.Error (errtxt, errpos)
    }
  }
  write (*x, l, c)
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
    if x, m, ok := float (s[:k]); ok {
      return x, m, true
    }
  }
  return 0., n, false
}

func constStarted (s string) (float64, uint, bool) {
  n := uint(len(s))
  if n >= 1 && s[:1] == "e" {
    return E, n, true
  }
  if n >= 2 && s[:2] == "pi" {
    return Pi, 2, true
  }
  return 0, n, false
}
