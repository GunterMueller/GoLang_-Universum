package str

// (c) Christian Maurer   v. 210212 - license see µU.go

import
  "µU/char"
const
  spc = byte(' ')

func devilsDung (s string) bool {
  return char.DevilsDung (s)
}

func const_(c byte, n uint) string {
  b := make([]byte, n)
  for i := uint(0); i < n; i++ {
    b[i] = c
  }
  return string(b)
}

func new_(n uint) string {
  if n == 0 {
    return ""
  }
  return const_(spc, n)
}

func clr (s *string) {
  *s = const_(spc, uint(len(*s)))
}

func lat1 (s string) string {
  char.ToHellWithUTF8 (&s)
  return s
}

func utf8 (s string) string {
  for i := len(s) - 1; i >= 0; i-- {
    c := s[i]
    if char.IsLatin1 (c) {
      s = s[:i] + string(c) + s[i+1:]
    }
  }
  return s
}

func letter (c byte) bool {
  return byte('A') <= c && c <= byte('Z') ||
         byte('a') <= c && c <= byte('z')
}

func digit (c byte) bool {
  return byte('0') <= c && c <= byte('9')
}

func letterOrDigit (c byte) bool {
  return byte('A') <= c && c <= byte('Z') ||
         byte('a') <= c && c <= byte('z') ||
         byte('0') <= c && c <= byte('9')
}

func lit (s string) bool {
  n := len (s)
  if n == 0 { return false }
  if ! letter (s[0]) { return false }
  for i := 1; i < n; i++ {
    if ! letterOrDigit (s[i]) { return false }
  }
  return true
}

func empty (s string) bool {
  for i := 0; i < len (s); i++ {
    if s[i] != spc {
      return false
    }
  }
  return true
}

func properLen (s string) uint {
  n := len (s)
  for {
    if n == 0 { break }
    if s[n-1] == spc {
      n --
    } else {
      break
    }
  }
  return uint(n)
}

func eq (s, t string) bool {
  n, k := properLen (s), properLen (t)
  if n != k { return false }
  if n == 0 { return true }
  return s[0:n] == t[0:n]
}

func toUpper (s *string) {
  n := properLen (*s)
  if n == 0 { return }
  b := make ([]byte, n)
  for i := uint(0); i < n; i++ {
    b[i] = char.Upper ((*s)[i])
  }
  *s = string(b)
}

func toLower (s *string) {
  n := properLen (*s)
  if n == 0 { return }
  b := make ([]byte, n)
  for i := uint(0); i < n; i++ {
    b[i] = char.Lower ((*s)[i])
  }
  *s = string(b)
}

func toUpper0 (s *string) {
  if len (*s) == 0 { return }
  *s = string(char.Upper ((*s)[0])) + (*s)[1:]
}

func toLower0 (s *string) {
  if len (*s) == 0 { return }
  *s = string(char.Lower ((*s)[0])) + (*s)[1:]
}

func cap0 (s string) bool {
  if s == "" { return false }
  return s[0] == char.Upper (s[0])
}

func equiv (s, t string) bool {
  n := properLen (s) // len(s)
  if properLen (t) /* len(t) */ != n {
    return false
  }
  for i := uint(0); i < n; i++ {
    if ! char.Equiv (s[i], t[i]) {
      return false
    }
  }
  return true
}

func less (s, t string) bool {
  n, n1 := len (s), len (t)
  i := 0
  for {
    if i == n {
      return n < n1
    }
    if i == n1 {
      return false
    }
    if char.Less (s[i], t[i]) {
      return true
    }
    if char.Less (t[i], s[i]) {
      return false
    }
    i++
  }
  return false
}

func equivLess (s, t string) bool {
  toUpper (&s)
  toUpper (&t)
  return less (s, t)
}

func pos (s string, b byte) (uint, bool) {
  n := uint(len (s))
  for i := uint(0); i < n; i++ {
    if s[i] == b {
      return i, true
    }
  }
  return n, false
}

func equivPos (s string, b byte) (uint, bool) {
  n := uint(len (s))
  for i := uint(0); i < n; i++ {
    if char.Equiv (s[i], b) {
      return i, true
    }
  }
  return n, false
}

func sub (s, t string) (uint, bool) {
//  char.ToHellWithUTF8 (&s) // sicher ist sicher
//  char.ToHellWithUTF8 (&t)
  n := properLen (s)
  if n == 0 { return 0, true }
  k, m := uint(len (t)), properLen (t)
  if n > m {
    return k, false
  }
  s = s[:n]
  for i := uint(0); i + n <= m; i++ {
    if s == t[i:i+n] {
      return i, true
    }
  }
  return k, false
}

func subAll (s, t string) (uint, []uint) {
  p := make([]uint, 0)
  k := properLen (s)
  s = s[:k]
  if k == 0 { return 0, p }
  m := properLen (t)
  if k > m {
    return 0, p
  }
  n := uint(0)
  for i := uint(0); i + k <= m; i++ {
    if s == t[i:i+k] {
      p = append (p, i)
      n++
    }
  }
  return n, p
}

func sub0 (s, t string) bool {
  n := properLen(s)
  if int(n) > len(t) { return false }
  return s[:n] == t[:n]
}

func equivSub (s, t string) (uint, bool) {
  if properLen (s) == 0 { return 0, true }
  toUpper (&s)
  toUpper (&t)
  return sub (s, t)
}

func ins1 (s *string, c byte, p uint) {
  t := string(c)
  char.ToHellWithUTF8 (&t)
  ins (s, t, p)
}

func ins (s *string, t string, p uint) {
  if len (t) == 0 || p > uint(len (*s)) { return }
  *s = (*s)[:p] + t + (*s)[p:]
}

func insAll (s *string, v, t string) {
  offSpc (s)
  j := uint(len(v))
  n, p := subAll (v, *s)
  if n == 0 {
    return
  }
  s1 := (*s)[:p[0]] + t
  for i := uint(1); i < n; i++ {
    s1 += (*s)[p[i-1]+j:p[i]] + t
  }
  s1 += (*s)[p[n-1]+j:]
  *s = s1
}

func replace1 (s *string, p uint, c byte) {
  n := len (*s)
  if int(p) >= n { return }
  t := string(c)
  *s = (*s)[:p] + t + (*s)[p+1:]
  char.ToHellWithUTF8 (s)
}

func replaceAll (s *string, b byte, t string) {
  n := len (*s)
  bs := make([]byte, 0)
  for i := 0; i < n; i++ {
    c := (*s)[i]
    if c == b {
      bs = append (bs, []byte(t)...)
    } else {
      bs = append (bs, c)
    }
  }
  *s = string(bs)
}

func replace (s *string, p uint, t string) {
  m := uint(len(t))
  n := len (*s)
  if p + m >= uint(n) { return }
  *s = (*s)[:p] + t + (*s)[p+m:]
}

func app (s *string, b byte) {
  n := uint(len(*s))
  bs := make ([]byte, n + 1)
  copy (bs[:n], []byte(*s))
  bs[n] = b
  *s = string(bs)
}

func rem (s *string, p, n uint) {
  if n == 0 { return }
  l := uint(len (*s))
  if p >= l { return }
  if p + n >= l {
    n = l - p
  }
  *s = (*s)[:p] + (*s)[p+n:]
}

func part (s string, p, n uint) string {
  if n == 0 { return "" }
  l := uint(len(s))
  if p >= l { return s }
  if p + n > l { n = l - p }
  return s[p:p+n]
}

func norm (s *string, n uint) {
  if n == 0 { *s = ""; return }
  k := uint(len (*s))
  if k > n {
    *s = (*s)[:n]
    return
  }
  for i := k; i < n; i++ { // k <= n
    *s += " "
  }
}

func offSpc (s *string) {
  offBytes (s, ' ')
}

func offSpc1 (s *string) {
  n := properLen (*s)
  *s = (*s)[:n]
}

func offBytes (s *string, b byte) {
  n := len (*s)
  if n == 0 { return }
  ss := make ([]byte, n)
  i, j := 0, 0
  loop:
  for j < n {
    if j == n { break }
    for (*s)[j] == b {
      j++
      if j == n {
        break loop
      }
    }
    ss[i] = (*s)[j]
    i++
    j++
  }
  *s = string(ss[0:i])
}

func move (s *string, left bool) {
  l := uint(len (*s))
  if l == 0 { return }
  if left {
    n := l
    for n > 0 && (*s)[0] == spc {
      *s = (*s)[1:]
      n--
    }
    for n < l {
      *s = *s + " "
      n++
    }
  } else {
    n := l
    for n = l; n >= 1; n-- {
      if (*s)[n - 1] != spc {
        break
      }
    }
    *s = (*s)[:n]
    for i := n; i < l; i++ {
      *s = " " + *s
    }
  }
}

func insSpace (s *string, p uint) {
  l := uint(len (*s))
  if l == 0 || p >= l { return }
  *s = (*s)[:p] + " " + (*s)[p:]
}

func shift (s *string, p uint) {
  l := uint(len (*s))
  if l <= 1 || p + 1 >= l { return }
  if (*s)[l-1] != spc { return }
  *s = (*s)[0:p] + " " + (*s)[p:l-1]
}

func center (s *string, n uint) {
  if n == 0 {
    return
  }
  move (s, false)
  l := ProperLen (*s)
  if n < l {
    *s = (*s)[:n]
    return
  }
  if l == n { return
  }
  if n == l + 1 {
    *s += " "
    return
  }
  k := (n - l) / 2 // + (n - l) % 2
  *s = new_(k) + *s + new_(n - (l + k))
}

func offNondigits (s *string) {
  l := uint(len (*s))
  if l == 0 { return }
  b := make ([]byte, l)
  i, j := uint(0), uint(0)
  loop: for j < l {
    if j == l { break }
    for ! digit ((*s)[j]) {
      j ++
      if j == l {
        break loop
      }
    }
    b[i] = (*s)[j]
    i ++
    j ++
  }
  *s = string(b[:i]) + new_(l - i)
}

func split (s string) (uint, []string, []uint) {
  char.ToHellWithUTF8 (&s)
  var t []string
  var p []uint
  l := properLen (s)
  spaceBefore := true
  n := uint(0)
  for i := uint(0); i < l; i++ {
    if s[i] == spc {
      spaceBefore = true
    } else {
      if spaceBefore {
        t = append (t, string(s[i]))
        p = append (p, i)
        n ++
        spaceBefore = false
      } else {
        t[n - 1] += string(s[i])
      }
    }
  }
  return n, t, p
}

func appendLF (s *string) {
  *s += "\n"
}

func appendLine (s *string, t string) {
  *s += (t + "\n")
}

func splitLine (s *string) string {
  l := uint(len (*s))
  if l == 0 { return "" }
  n := uint(0)
  for n = 0; n < l; n++ {
    if (*s)[n] == byte('\n') {
      break
    }
  }
  t := (*s)[:n]
  n ++
  *s = (*s)[n:]
  return t
}

func splitBrackets (s string, sep, b byte) []string {
  var b1 byte
  switch b {
  case byte('('): b1 = byte(')')
  case byte('['): b1 = byte(']')
  case byte('{'): b1 = byte('}')
  case byte('<'): b1 = byte('>')
  default:
    return nil
  }
  p, l := make([]uint, 0), make([]uint, 0)
  n := uint(len(s))
  if s[0] != b || s[n-1] != b1 {
    return nil
  }
  p = append(p, 0)
  c, j := 0, uint(1)
  for i := uint(1); i < n - 1; i++ {
    switch s[i] {
    case b:
      c++
    case b1:
      if c > 0 {
        c--
      } else {
        return nil
      }
    case sep:
      if c == 0 {
        p, l = append(p, i), append(l, i - j)
        j = i + 1
      }
    }
  }
  p, l = append(p, n - 1), append(l, n - 1 - j)
  n = uint(len(p)) - 1
  ss := make([]string, n)
  for i := uint(0); i < n; i++ {
    ss[i] = part (s, p[i]+1, l[i])
  }
  return ss
}

func splitByte (s string, b byte) ([]string, uint) {
  ss, n := make([]string, 0), uint(0)
  if s == "" {
    return ss, n
  }
  if s[0] == '/' {
    s = s[1:]
  }
  l := len(s)
  if l > 1 && s[l-1] == '/' {
    s = s[:l-1]
  }
  for {
    p, c := pos (s, b)
    if ! c {
      ss = append (ss, s)
      break
    }
    ss = append (ss, s[:p])
    s = s[p+1:]
    n++
  }
  return ss, n
}

func rightBr (s string) uint {
  n := 0
  for i := 1; i < len(s); i++ {
    b := s[i]
    if b == '(' {
      n++
    }
    if b == ')' {
      if n > 0 {
        n--
      } else {
        return uint(i)
      }
    }
  }
  return 0
}

func isVarString (s string) bool {
  if ! char.IsLetter (s[0]) {
    return false
  }
  n := properLen (s)
  if n == 1 {
    return true
  }
  for i := uint(1); i < n; i++ {
    if ! char.IsLetterOrDigit (s[i]) {
      return false
    }
  }
  return true
}

func startsWithVar (s string) (string, uint, bool) {
  n := uint(len(s))
  if n == 1 && char.IsLetter(s[0]) {
    return string(s[0]), 1, true
  }
  if n > 0 {
    for p := uint(1); p < n; p++ {
      if char.IsLetterOrDigit (s[p]) {
      } else {
        t := s[:p]
        if isVarString (t) {
          return t, p, true
        } else {
          break
        }
      }
      if p == n - 1 {
        return s, p + 1, true
      }
    }
  }
  return "", 0, false
}
