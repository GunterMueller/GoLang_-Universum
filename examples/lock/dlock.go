package main

// (c) Christian Maurer   v. 241023 - license see µU.go

// open N windows and start the program with "dlock n",
// where 0 <= n < N is the number of the corresponding window

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
    N = 6
  type
    state byte; const (
    outsideCS = state(iota)
    waiting
    insideCS
    nStates
  )
  var (
    f = [nStates]col.Colour { col.LightGreen(), col.Yellow(), col.LightRed() }
    b = col.Black()
    t = [nStates]string {"  I am outside the c.s.  ",
                         "I am waiting for the c.s.",
                         "     I am in the c.s.    "}
    e = [nStates]string {"enter c.s.: enter-key",
                         "... please wait ...",
                         "leave cs: backspace-key"}
  )
  st := outsideCS
  me := ego.Ego(N)
  w := uint(25 * 8 + 4)
  scr.NewWH ((w + 8) * (me % N), 0, w, 4 * 16); defer scr.Fin()
  hosts := make([]string, N)
  for i := uint(0); i < N; i++ { hosts[i] = env.Localhost() }
  l := dlock.New (me, hosts, 0)
  n := true
  for {
    if n {
      scr.Colours (f[st], b); scr.Write (t[st], 1, 0); errh.Hint (e[st])
    }
    n = true
    switch c, _ := kbd.Command(); c {
    case kbd.Enter:
      if st == outsideCS {
        st = waiting
        scr.Colours (f[st], b); scr.Write (t[st], 1, 0); errh.Hint (e[st])
        l.Lock()
        st = insideCS
      }
    case kbd.Back:
      if st == insideCS {
        st = outsideCS
        l.Unlock()
      }
    case kbd.Esc:
      return // unsauber, andere Prozesse müssen erst gekillt werden TODO
    default:
      n = false
    }
  }
}
