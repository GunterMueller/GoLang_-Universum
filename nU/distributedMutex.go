package main

import ("nU/env"; "nU/term"; "nU/col"; "nU/scr"; "nU/ego"; "nU/dlock")

func main() {
  const N = 6
  type state byte; const (outside = state(iota); waiting; inCS; nStates)

  var (
    f = [nStates]col.Colour { col.LightGreen(), col.Yellow(), col.LightRed() }
    b = col.Black()
    t = [nStates]string { "  I am outside the c.s.  ",
                          "I am waiting for the c.s.",
                          "    I am in the c.s.     " }
    h = [nStates]string { " enter c.s.: Return key  ",
                          "...... please wait ......",
                          "leave c.s.: Backspace key" }
  )
  s := outside
  me := ego.Ego (N)
  scr.New(); defer scr.Fin()
  term.New(); defer term.Fin()
  hosts := make([]string, N)
  for i := uint(0); i < N; i++ {
    hosts[i] = env.Localhost()
  }
  l := dlock.New (me, hosts, 0)
  n := true
  for {
    if n {
      scr.Colours (f[s], b); scr.Write (t[s], 0, 0); scr.Write (h[s], 1, 0)
    }
    scr.Warp (1, 25)
    scr.Switch (true)
    c := term.Read()
    n = true
    switch c {
    case term.Enter:
      if s == outside {
        s = waiting
        scr.Colours (f[s], b); scr.Write (t[s], 0, 0); scr.Write (h[s], 1, 0)
        l.Lock()
        s = inCS
      }
    case term.Back:
      if s == inCS {
        s = outside
        l.Unlock()
      }
    case term.Esc:
      return // XXX unclean, other processes have to be killed before
    default:
      n = false
    }
  }
}
