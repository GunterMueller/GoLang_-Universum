package main

// (c) Christian Maurer   v. 241009 - license see µU.go

// 8 Fenster aufmachen und das Programm mit "dlock n" starten,
// wobei 0 <= n < 8 die Nummer des entsprechenden Fensters ist.

import (
  "µU/env"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/errh"
  "µU/ego"
  "µU/dlock"
)

func main() {
  const
    N = 8
  type
    state byte; const (
    outside = state(iota)
    waiting
    inCS
    nStates
  )
  var (
    f = [nStates]col.Colour { col.LightGreen(), col.Yellow(), col.LightRed() }
    b = col.Black()
    t = [nStates]string { "I am outside the cs.",
                          " I wait for the cs. ",
                          "   I am in the cs.  " }
    h = [nStates]string { "enter cs: enter-key ",
                          "... wait please ... ",
                          "leave cs: back-key  " }
  )
  s := outside
  me := ego.Ego(N)
  w := uint(20 * 8 + 4)
  scr.NewWH ((w + 4) * (me % N), 0, w, 3 * 16); defer scr.Fin()
  hosts := make([]string, N)
  for i := uint(0); i < N; i++ { hosts[i] = env.Localhost() }
  l := dlock.New (me, hosts, 0)
  n := true
  for {
    if n {
      scr.Colours (f[s], b); scr.Write (t[s], 0, 0); errh.Hint (h[s])
    }
    c, _ := kbd.Command()
    n = true
    switch c {
    case kbd.Enter:
      if s == outside {
        s = waiting
        scr.Colours (f[s], b); scr.Write (t[s], 0, 0); errh.Hint (h[s])
        l.Lock()
        s = inCS
      }
    case kbd.Back:
      if s == inCS {
        s = outside
        l.Unlock()
      }
    case kbd.Esc:
      return // unsauber, andere Prozesse müssen erst gekillt werden TODO
    default:
      n = false
    }
  }
}
