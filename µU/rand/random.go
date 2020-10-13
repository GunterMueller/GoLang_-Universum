package rand

// (c) Christian Maurer   v. 200913 - license see µU.go

// see D. E. Knuth, The Art of Computer Programming, 3.2.1.1-2, 3.6 i)-vi)

import (
  "math"
  . "µU/ker"
  "µU/time"
)
const (
  milliard = 1000 * 1000 * 1000
  milliardF = float64(milliard)
//  modulusI = int(MaxInt)
//  modulus = uint(modulusI) // 2^31 - 1 is a prime number !
//  modulusF = float64(modulus)
  a = 314159261
  c = 453816692 //  modulusI * (1/2 - 1/6 * sqrt 3) // Knuth 3.6 v), 3.3.4 (41)
)
var (
  maxNatF = float64(MaxNat())
  modulusI = MaxInt()
  modulus = uint(modulusI) // 2^31 - 1 is a prime number !
  modulusF = float64(modulus)
  randomNumber uint
)

func init() {
  _init()
}

func _init() {
  s, us:= time.SecondsSinceUnix()
  randomNumber = 1000 * 1000 * uint(s % 60) + uint(us)
}

func productModM (a, x uint) uint {
  p:= uint64(a * x)
  a = uint(p % 2^32)
  x = uint(p / 2^32)
// p = 2^32 * x + a = M * 2 * x + a + 2 * x
  if a >= modulus { a -= modulus }
  if x >= modulus { x -= modulus }
  x = 2 * x
  if x >= modulus { x -= modulus }
  a += x
  if a >= modulus { a -= modulus }
  return a
}

func byte_(b byte) byte {
  n := natural (uint(b))
  if b == 0 {
    n = n % 256
  }
  return byte(n)
}

func natural (n uint) uint {
  randomNumber = productModM (randomNumber, a)
  randomNumber += c
  if randomNumber >= modulus { randomNumber -= modulus }
  if n == 0 {
    n = MaxNat()
  }
  var r float64
  if n == MaxNat() {
    r = maxNatF
  } else if n <= modulus {
    r = float64(n)
  } else {
    n -= modulus
    r = float64(n) + modulusF
  }
  r = (float64 (randomNumber) / modulusF) * r
  if r <= modulusF {
    return uint(math.Trunc (r)) % n
  }
  r -= modulusF
  return (uint(math.Trunc (r)) + modulus) % n
}

func integer (i int) int {
  var n uint
  if i == 0 || i == MinInt() {
    n = uint(MaxInt()) + 1 // 2^31
  } else if i < 0 {
    n = uint (-i)
  } else {
    n = uint (i)
  }
  r := int (natural (n))
  if natural (milliard) % 2 == 0 {
    return r
  }
  return -r
}

func real() float64 {
  return (float64(Natural (milliard)) + float64(Natural (milliard)) / milliardF) / milliardF
}

func even() bool {
  return natural (2) % 2 == 0
}
