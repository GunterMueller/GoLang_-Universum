package ntrav

// (c) murus.org  v. 170106 - license see murus.go

import (
  . "murus/obj"
  "murus/nchan"
)

// Pre: scr.New (constr. of channels, log)
//      nh is sorted by the identities of the nodes hosts (is done in murus/lan!).
func (x *netTraversal) dfs (op Op) (time0, time1 uint) {
  ch := make ([]nchan.NetChannel, x.n)
  for i := uint(0); i < x.n; i++ {
    ch[i] = nchan.New (uint(0), x.me, x.nr[i], x.h[i], x.port[i])
  }
  visited := make ([]bool, x.n)
  parent := x.n
  done := make (chan int)
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      probe := true // false means echo
      if x.me == x.root && j == 0 { // the root node sends the first message
        op(nil)
        time0 = 1
        parent = x.n + 1 // trick no. 17, see below
        visited[0] = true
        ch[0].Send (time0)
        probe = true
        x.log (j, time0, true, probe)
      }
      t := ch[j].Recv().(uint)
      probe = ! visited[j]
      x.log (j, t, false, probe)
      u := x.n // channel to next unvisited neighbour == x.n, if all visited
      for k := uint(0); k < x.n; k++ {
        if k != j && ! visited[k] {
          u = k
          break
        }
      }
      var k uint // channel to neighbour for the next send operation
      if probe { // i.e. ! visited[j]
        visited[j] = true
        if parent == x.n { // not for x.root (trick no. 17)
          op(nil)
          parent = j
          t++; time0 = t
          if u == x.n { // all neighbours visited
            t++; time1 = t
            probe = false
            k = parent // send echo back to parent
          } else {
            probe = true
            k = u // send probe to next unvisited neighbour
          }
        } else {
          probe = false
          k = j // send echo back to sender
        }
      } else { // received echo, i.e. visited[j]
        if u == x.n { // all neighbours visited
          t++; time1 = t
          if x.me == x.root { // the root node must not reply any more
            done <- 0
            return
          }
          probe = false
          k = parent // send echo back to parent
        } else {
          probe = true
          k = u // send probe to next unvisited neighbour
        }
      }
      visited[k] = true
      ch[k].Send (t)
      x.log (k, t, true, probe)
      done <- 0
//      return // XXX ? ? ? ?
    }(i)
  }
  for i := uint(0); i < x.n; i++ {
    <-done
  }
  return
}
