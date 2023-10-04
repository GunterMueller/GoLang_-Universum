package main // mergesort

// (c) Christian Maurer   v. 230924 - license see µU.go

import (
  "µU/rand"
  "µU/time"
  "µU/col"
  "µU/scr"
  "µU/errh"
)

const (
  N = 32 // 2-er Potenz !
  max = 2 * N - 2
)
var (
  c [max+1]chan int
  done chan bool
  cb, cw = col.Black(), col.FlashWhite()
)

func P() { scr.Lock() }

func V() { scr.Unlock() }

func pause() { time.Sleep (3 + rand.Natural(5)) }

func pos0 (i, d, y, x, n int) (uint, uint) {
  if i == 2 * N - 1 { return 20, 4 * uint(i) }
  if i / n == 0 { return uint(4 * y), uint(d) / 2 - 2 + uint(d * i) + 1 }
  return pos0 (i - n, 2 * d, y + 1, x, n / 2)
}

func pos (i int) (uint, uint) {
  return pos0 (i, 4, 0, 0, N)
}

func next0 (i, j, d, n int) int { // Pre: i <= max
  if i < n { return n + j / 2 }
  return next0 (i, j - d, d / 2, n + d / 2)
}

func next (i int) (int) {
  return next0 (i, i, N, N)
}

func show (k, i int, l bool, c col.Colour) {
  P()
  y, x := pos (i); if i < N || i >= N && ! l { y++ }
  scr.Colours (c, cb); scr.Write ("  ", y, x); scr.WriteNat (uint(k), y, x)
  V()
}

func del (i int, l bool) {
  P()
  y, x := pos (i); if i >= N && ! l { y++ }
  scr.Colours (col.FlashWhite(), cb); scr.Write ("  ", y, x)
  V()
  pause()
//XXX
}

func generate() {
  var (
    number [N]int
    schonda [N]bool
  )
  rand.Init()
  for i := 0; i < N; i++ {
    for {
      n := int(rand.Natural(N))
      if ! schonda[n] {
        schonda[n] = true
        number[i] = n
        break
      }
    }
    y, x := pos (i)
    x0, y0 := 8 * int(x + 1), 16 * int(y + 1)
    P()
    scr.Circle (x0, y0 + 8, 12)
    y, x = pos (next (i))
    x1, y1 := 8 * int(x + 1), 16 * int(y + 1)
    scr.Line (x0, y0, x1, y1)
    V()
    show (number[i], i, i % 2 == 0, cw)

  }
  for i := 0; i < N; i++ {
    c[i] <- number[i] // zu den Sortierern
    show (number[i], i, i % 2 == 0, col.LightRed()) // del (i, true)
  }
}

func nAwaitedNumbers (i int) int {
  e, d, m := N / 2, 1, N - 2
  for m > i { d *= 2; m -= d; e /= 2 }
  return e
}

func sort (i int) { // i = Nummer des Sortierprozesses
  rL, rR := 2 * i, 2 * i + 1 // Nummer des linken bzw. rechten Empfangskanals
  s := N + i                 // Nummer des Sendekanals
  y, x := pos (s)
  x0, y0 := 8 * int(x + 1), 16 * int(y + 1)
  P()
  r := uint(20); if i == max { r = 50 }
  scr.Colours (cw, cb)
  scr.Circle (x0, y0, r)
  if i < N - 2 {
    y, x = pos (next (s))
    x1, y1 := 8 * int(x + 1), 16 * int(y + 1)
    scr.Line (x0, y0, x1, y1)
  }
  V()
  e := nAwaitedNumbers(i) // Anzahl der zu erwartenden Botschaften
  if i >= N / 2 { pause() }
//                XXX
  nL  := <-c[rL]; show (nL, s, true, cw)
  if i >= N / 2 { pause() }
//                XXX
  nR := <-c[rR]; show (nR, s, false, cw)
  pause()
//XXX
  eL, eR := 1, 1 // Anzahl der von links bzw. rechts empfangenen Botschaften
  for eL <= e && eR <= e {
    if nL <= nR {
      c[s] <- nL; del (s, true)
      eL++
      if eL <= e {
        pause(); nL = <-c[rL]; show (nL, s, true, cw); pause()
//      XXX                                            XXX
      }
    } else {
      c[s] <- nR; del (s, false)
      eR++
      if eR <= e {
        pause(); nR = <-c[rR]; show (nR, s, false, cw); pause()
//      XXX                                             XXX
      }
    }
  }
  for eL <= e {
    c[s] <- nL; del (s, true)
    eL++
    if eL <= e {
      pause(); nL = <-c[rL]; show (nL, s, true, cw); pause()
//    XXX                                            XXX
    }
  }
  for eR <= e {
    c[s] <- nR; del (s, false)
    eR++
    if eR <= e {
      pause(); nR = <-c[rR]; show (nR, s, false, cw); pause()
//    XXX                                             XXX
    }
  }
}

func write() {
  y, _ := pos (2 * N - 2)
  for i := 0; i < N; i++ {
    pause()
//  XXX
    n := uint(<-c[max])
    P()
    scr.Colours (col.LightGreen(), cb); scr.WriteNat (n, y + 4, 4 * uint(i) + 1)
    V()
  }
  done <- true
}

func main() {
  scr.NewWH (0, 0, 1024, 432); defer scr.Fin()
  for i := 0; i <= max; i++ { c[i] = make(chan int) }
  done = make(chan bool)
  go generate()
  for i := 0; i < N - 1; i++ { go sort (i) }
  go write()
  <-done
  errh.Error0 ("fertig")
}
