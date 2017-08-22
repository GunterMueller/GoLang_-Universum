package main

// (c) murus.org  v. 170814 - license see murus.go

// >>> E. W. Dijkstra, Hierarchical Ordering of Sequential Processes
//     Acta Informatica 1 (1971), p. 115-138 (section 6.)
// >>> Nichtsequentielle Programmierung mit Go 1 kompakt, S. 95ff., 163ff., 187ff.

import (
  "murus/ker"
  "murus/rand"
  "murus/kbd"
  "murus/col"
  "murus/scr"
  "murus/errh"
  . "murus/menue"
  . "murus/phil"
)
var
  menu = New ("Die speisenden Philosophen")

func wait (t uint) {
  ker.Sleep ((t + rand.Natural (3 * t)))
}

func eat (ph Philos, i uint) {
  ph.Lock (i)
  switch i { case 0, 2:
    wait (5) // they like to eat ...
  default:
    wait (2)
  }
  ph.Unlock (i)
}

func run (ph Philos) {
  errh.Hint ("Philosoph will essen: Mausklick auf seinen Platz           Ende: Esc")
  Start()
  loop: for {
    scr.MousePointer (true)
    switch c, _ := kbd.Command(); c {
    case kbd.Esc:
      break loop
    case kbd.Here:
      if i, ok := WannaEat(); ok {
        go eat (ph, i)
      }
    }
  }
}

var
  stop bool

func runNaive (ph Philos, i uint, ch chan bool) {
  for ! stop {
    eat (ph, i)
    switch i { case 0, 2:
      wait (2) // they like to eat ...
    default:
      wait (8)
    }
  }
  ch <- true
}

func naive() {
  Start()
  go func() { ker.Sleep (60); stop = true }()
  ph := NewNaive()
  halt := make (chan bool)
  for i := uint(0); i < NPhilos; i++ {
    go func (q Philos, j uint, c chan bool) { runNaive (q, j, c) } (ph, i, halt)
  }
  for i := uint(0); i < NPhilos; i++ {
    <-halt
  }
}

func main() {
  fg, bg:= col.LightWhite, col.Blue; scr.Colours (fg, bg)
  na := New ("Naive Lösung mit Verklemmung")
              na.Leaf (func() { naive() }, false); menu.Ins (na)
  bo := New ("Nur n-1 Philosophen dürfen gleichzeitig an den Tisch")
              bo.Leaf (func() { run (NewBounded()) }, false); menu.Ins (bo)
  un := New ("Unsymmetrie: nicht alle zuerst nach der linken Gabel")
              un.Leaf (func() { run (NewUnsymmetric()) }, false); menu.Ins (un)
  su := New ("Unfaire Lösung mit Semaphoren, aushungerungsgefährdet")
              su.Leaf (func() { run (NewSemaphoreUnfair()) }, false); menu.Ins (su)
  sf := New ("Faire Lösung mit Semaphoren nach Dijkstra")
              sf.Leaf (func() { run (NewSemaphoreFair()) }, false); menu.Ins (sf)
  cs := New ("Lösung mit universellen kritischen Abschnitten")
              cs.Leaf (func() { run (NewCriticalSection()) }, false); menu.Ins (cs)
//  co := New ("Lösung mit Conditions")
//              co.Leaf (func() { run (NewCondition()) }, false); menu.Ins (co)
  ca := New ("Faire Lösung mit universellen kritischen Abschnitten und \"Aging\"")
              ca.Leaf (func() { run (NewCriticalSectionAging()) }, false); menu.Ins (ca)
  mo := New ("Lösung mit universellem Monitor")
              mo.Leaf (func() { run (NewMonitor()) }, false); menu.Ins (mo)
  mu := New ("Unfaire Monitor-Lösung nach Dijkstra")
              mu.Leaf (func() { run (NewMonitorUnfair()) }, false); menu.Ins (mu)
  mf := New ("Faire Monitor-Lösung nach Dijkstra")
              mf.Leaf (func() { run (NewMonitorFair()) }, false); menu.Ins (mf)
  cm := New ("Lösung mit konditioniertem universellen Monitor")
              cm.Leaf (func() { run (NewCondMonitor()) }, false); menu.Ins (cm)
  cb := New ("Lösung mit Botschaftenaustausch nach Ben-Ari")
              cb.Leaf (func() { run (NewChannel()) }, false); menu.Ins (cb)
  cu := New ("Unsymmetrische Lösung mit Botschaftenaustausch")
              cu.Leaf (func() { run (NewChannelUnsymmetric()) }, false); menu.Ins (cu)
  menu.Exec()
}
