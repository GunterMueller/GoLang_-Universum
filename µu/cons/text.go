package cons

// (c) Christian Maurer   v. 140713 - license see µu.go

//import
//  "µu/linewd"

func (X *console) Transparent() bool {
  return X.transparent
}

func (X *console) Transparence (on bool) {
  X.transparent = on
}

func (X *console) Write1 (b byte, l, c uint) {
  if ! visible { return }
  if l >= X.nLines || c >= X.nColumns { return }
  f := X.codeF
//  w := X.lineWd
//  X.lineWd = linewd.Thin
  for i := uint(0); i < X.ht1; i++ {
    for j := uint(0); j < X.wd1; j++ {
      if pointed (X.fontsize, b, i, j) {
        X.codeF = f
      } else {
        X.codeF = X.codeB
      }
      X.Point (int(X.wd1 * c + j), int(X.ht1 * l + i))
    }
  }
  X.codeF = f
//  X.lineWd = w
}

func (X *console) Write (s string, l, c uint) {
  if len(s) == 0 || ! visible { return }
  n := len (s)
  if c + uint(n) > X.nColumns { n = int(X.nColumns - c) }
  for i := 0; i < n; i++ {
    X.Write1 (s[i], l, c + uint(i))
  }
}

func (X *console) WriteNat (n, l, c uint) {
  t := "00"
  if n > 0 {
    const M = 10
    bs := make ([]byte, M)
    for i := 0; i < M; i++ {
      bs[M - 1 - i] = byte('0' + n % 10)
      n = n / M
    }
    s := 0
    for s < M && bs[s] == '0' {
      s++
    }
    t = ""
    if s == M - 1 { s = M - 2 }
    for i := s; i < M - int(n); i++ {
      t += string(bs[i])
    }
  }
  X.Write (t, l, c)
}

func (X *console) Write1Gr (b byte, x, y int) {
  if ! visible { return }
  f := X.codeF
//  w := X.lineWd
//  X.lineWd = linewd.Thin
  for i := uint(0); i < X.ht1; i++ {
    for j := uint(0); j < X.wd1; j++ {
      if pointed (X.fontsize, b, i, j) {
        X.codeF = f
        X.Point (x + int(j), y + int(i))
      } else if ! X.transparent {
        X.codeF = X.codeB
        X.Point (x + int(j), y + int(i))
      }
    }
  }
//  X.lineWd = w
  X.codeF = f
}

func (X *console) WriteGr (s string, x, y int) {
  n := len (s)
  if n == 0 || ! visible { return }
  if x < X.x || y < X.y { return }
  n = len(s)
  for i := 0; i < n; i++ {
    X.Write1Gr (s[i], x + i * int(X.wd1), y)
  }
}

func (X *console) Write1InvGr (b byte, x, y int) {
  if ! visible { return }
  if x < X.x || x >= X.x + int(X.wd - X.wd1) || y < X.y || y >= X.y + int(X.ht - X.ht1) { return }
  for i := uint(0); i < X.ht1; i++ {
    for j := uint(0); j < X.wd1; j++ {
      if pointed (X.fontsize, b, i, j) {
        X.PointInv (x + int(j), y + int(i))
      } else if ! X.transparent {
        X.PointInv (x + int(j), y + int(i))
      }
    }
  }
}

func (X *console) WriteInvGr (s string, x, y int) {
  n := len (s)
  if n == 0 || ! visible { return }
  if x < 0 || y < 0 { return }
  for i := 0; i < n; i++ {
    X.Write1InvGr (s[i], x + i * int(X.wd1), y)
  }
}
