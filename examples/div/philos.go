package main

// (c) Christian Maurer   v. 241005 - license see µU.go

// >>> E. W. Dijkstra, Hierarchical Ordering of Sequential Processes
//     Acta Informatica 1 (1971), p. 115-138 (section 6.)
// >>> Nichtsequentielle und Verteilte Programmierung mit Go, S. 180, 253

// >>> Klick on the philosophers, who want to eat.

import (
  "µU/time"
  "µU/rand"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/errh"
  . "µU/phil"
)

func wait (t uint) {
  time.Sleep ((t + rand.Natural (3 * t)))
}

func eat (ph Philos, i uint) {
  ph.Lock (i)
  switch i {
  case 0, 2, 5, 8:
    wait (5) // they like to eat ...
  default:
    wait (3)
  }
  ph.Unlock (i)
}

func run (ph Philos) {
  Start()
  var helpon = false
  loop:
  for {
    scr.MousePointer (true)
    switch c, _ := kbd.Command(); c {
    case kbd.Esc:
      scr.Cls()
      break loop
    case kbd.Help:
      helpon = ! helpon
      if helpon {
        errh.Hint ("Philosoph will essen: Mausklick auf seinen Platz           Ende: Esc")
      } else {
        errh.DelHint()
      }
    case kbd.Here:
      if i, ok := SitDownAtTable(); ok {
        go eat (ph, i)
      }
    }
  }
}

func main() {
  fg, bg := col.FlashWhite(), col.Blue(); scr.Colours (fg, bg)
// choose one of the following implementations (see µU/phil/def.go):
/*/
  run (NewNaive())
  run (NewSemaphore())
  run (NewBounded())
  run (NewUnsymmetric())
  run (NewSemaphoreUnfair())
  run (NewSemaphoreFair())
  run (NewCriticalSection())
  run (NewMonitor())
  run (NewMonitorFair())
  run (NewMonitorUnfair())
  run (NewCondMonitor())
  run (NewChannel())
  run (NewChannelUnsymmetric())
//  run (NewFarMonitor ("s", p, env.N(1) == 0)) // s = name of the server, p = used port // XXX
/*/
  run (NewCondMonitor())
}
