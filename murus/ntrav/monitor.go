package ntrav

// (c) murus.org  v. 170106 - license see murus.go
//
// >>> highly experimental stuff - do not yet take that serious !

import (
  . "murus/obj"
//  "murus/errh"
  "murus/fmon"
)
type
  Monitor interface {

  Probe (i uint) uint
}
type
  monitor struct {
         visited []bool
                 fmon.FarMonitor
                 }

func (x *netTraversal) channel (id uint) uint {
  j := x.n
  for i := uint(0); i < x.n; i++ {
    if x.nr[i] == id {
      j = i
      break
    }
  }
  return j
}

func NewMonitor (x *netTraversal, id uint, s bool) Monitor {
  m := new(monitor)
  m.visited = make([]bool, x.n)
  f := func (a Any, i uint) Any {
         parent := x.n
         if x.me == x.root {
           parent = uint(x.n + 1)
         }
         j := x.channel(a.(uint)) // a.(uint) == x.nr[j]
         probe := ! m.visited[j]
s := "echo"; if probe { s = "probe" }
println(s + " from", a.(uint))
         u := x.n
         for k := uint(0); k < x.n; k++ {
           if k != j && ! m.visited[k] {
             u = k
             break
           }
         }
         var k uint
         if probe {
           m.visited[j] = true
           if parent == x.n {
println("parent == x.n")
             parent = j
             if u == x.n {
println("all visited, echo to parent", u)
               probe = false
               k = parent
             } else { // ex. unvisited's
               probe = true
               k = u
println("next unvisited is", x.nr[k])
             }
           } else { // parent != x.n
             probe = false
             k = j
println("parent is", x.nr[k])
           }
         } else { // echo
           if u == x.n {
             if x.me == x.root {
println("x.me == root")
//               done <- 0
               return 0 // what ?
             }
             probe = false
             k = parent
           } else { // ex. unvisited's
println("ex unvisiteds")
             probe = true
             k = u
           }
         }
         m.visited[k] = true
// println("will call mon", k, "Probe", x.me)
println(x.me, "will call Probe to", x.nr[k])
         b := x.mon[k].Probe (x.me)
println(x.me, "'s Probe to", x.nr[k], "got result", b)
         return x.me
       }
  m.FarMonitor = fmon.New (uint(0), 1, f, AllTrueSp, x.Host, x.port1[id], s)
  return m
}

func (x *monitor) Probe (i uint) uint {
  println("waiting for return of Probe", i)
  r := x.F(i, 0).(uint)
  println("Probe", i, "returned", r)
  return r
}
