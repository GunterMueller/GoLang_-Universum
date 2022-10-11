package n

// (c) Christian Maurer   v. 220715 - license see µU.go

import
  "µU/col"

// Returns the number n of connected substrings of s, that consist only
// of digits, and slices s, p, l, s.t. s[i] is the i-th such string,
// p[i] its start position in s and l[i] its length (i < n).
func DigitSequences (s string) (uint, []string, []uint, []uint) { return digSeqs(s) }

// Returns the number of digits of n.
func Wd (n uint) uint { return wd(n) }

// Returns (n, true), Iff s represents the natural number n in the range of uint;
// returns otherwise (0, false).
func Natural (s string) (uint, bool) { return natural(s) }

func String (n uint) string { return string_(n) }

// Returns the representation of n formatted to w digits,
// with trailing 0's, iff z == true, otherwise trailing spaces.
func StringFmt (n, w uint, z bool) string { return stringFmt(n,w,z) }

// Next time, a natural number is written to the screen,
// that is done in the colours (foreground, background) = (f, b).
func Colours (f, b col.Colour) { colours(f,b) }

// Pre: l < scr.NLines(); c + Len(n) < scr.NColumns().
// n is written to the screen starting at (line, column) = (l, c).
func Write (n, l, c uint) { write(n,l,c) }
func WriteGr (n uint, x, y int) { writeGr(n,x,y) }

// TODO Spec
func SetWd (w uint) { setWd(w) }

// Pre: s. Write.
// n is the natural number that was edited at (line, column) = (l, c).
func Edit (n *uint, l, c uint) { edit(n,l,c) }
func EditGr (n *uint, x, y int) { editGr(n,x,y) }

// Returns the sum of the digits of n.
func SumDigits (n uint) uint { return sumDigits(n) }

// Returns 0, if n == 0 or k == 0;
// returns otherwise the greatest common divisor of n and k.
func Gcd (n, k uint) uint { return gcd(n,k) }

// Returns 0, if n == 0 or k == 0;
// returns otherwise the least common multiple of n and k.
func Lcm (n, k uint) uint { return lcm(n,k) }

// Returns n!, the number of bijective mappings between n-element-sets. 
func Fak (n uint) uint { return fak(n) }

// Returns n^k, the number of mappings from a k-element-set to a n-element-set.
func Pow (n, k uint) uint { return pow(n,k) }

// Returns n over k, the number of k-element-subsets of an n-element-set.
func Binom (n, k uint) uint { return binom(n,k) }

// Returns the falling factorial (n, k), the number of injective mappings
// from a k-element-set into an n-element-set.
// One has LowFak (n, k) == Fak (k) * Binom (n, k)
func LowFak (n, k uint) uint { return lowFak(n,k) }

// Returns Stirling2 (n, k), the number of k-partitions of an n-element-set.
// The function f, defined by
//   func f (k, n uint) uint { a:= uint(0)
//     for i:= 0; i <= n; i++ { a = a + Stirl2 (n, i) * LowFak (k, i) }; return a
//   }
// coincides with Pow.
func Stirl2 (n, k uint) uint { return stirl2(n,k) }
