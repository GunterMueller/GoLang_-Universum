package dgra

// (c) Christian Maurer   v. 241013 - license see µU.go

import (
  "os"
//  "math"
  . "µU/dgra/st"
  "µU/perm"
  "µU/time"
  "µU/col"
  "µU/scr"
  "µU/errh"
  "µU/N"
)

func (x *distributedGraph) Misra() {
  const (
    M = 8 // number of involved processes
    h1 = 16 // scr.Ht1()
  )
  w := uint(160)
  r := int(w / 2)
  h := w + h1
  x0 := x.me * (w + 8) // 8 = 2 x border of window
  var X, Y [M] int
  X[0], Y[0] = int(x0) - 8, int(h) - r - 8
  r0 := 64.0
  for i := uint(0); i < M; i++ {
    a := 90 - float64(360 * i) / M
    X[i], Y[i] = r + int(r0 * cos(a)), r - int(r0 * sin(a))
  }
  scr.NewWH (x0, 0, w, h); defer scr.Fin()
  scr.Name ("rana " + N.String(x.me))
  scr.ColourF (col.White())
  scr.Circle (r, r, uint(r0))
  for i := uint(0); i < M; i++ {
    scr.ColourF (col.Black()); scr.CircleFull (X[i], Y[i], 10)
    scr.ColourF (col.White()); scr.Circle (X[i], Y[i], 10)
    scr.WriteNatGr (i, X[i] - 3, Y[i] - 7)
  }
  x.connect (uint(0))
  in, out := uint(0), uint(1)
  if x.Graph.Outgoing(0) { in, out = out, in }
  var c [NStates]col.Colour
  c[Active], c[Ready] = col.LightRed(), col.LightGreen()
  state := Active
  write (x.me, X[x.me], Y[x.me], c[state])
  p := perm.New (2 * M)
  n := uint(0)
  if x.me == x.root {
    x.send (out, n)
    errh.Hint1 ("sent", n)
  }
  round := uint(0)
  for {
    round++
    time.Sleep (5)
    n = x.recv (in).(uint)
    errh.Hint1 ("received", n)
    if state == Active && round == 1 + p.F(x.me) {
      state = Ready
      write (x.me, X[x.me], Y[x.me], c[state])
      errh.Hint ("Ready")
    }
    if state == Ready {
      if n < M {
        n++
      }
    } else {
      n = 0
    }
    x.send (out, n)
    errh.Hint1 ("sent", n)
    if n == M {
      errh.Error0 ("termination detected")
      os.Exit (0)
    }
  }
}
