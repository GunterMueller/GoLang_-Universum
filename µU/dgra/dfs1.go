package dgra

// (c) Christian Maurer   v. 231229 - license see µU.go

import (
  . "µU/obj"
  "µU/scr"
)

func (x *distributedGraph) Dfs1() {
  scr.Cls()
  x.connect (nil)
  defer x.fin()
  x.tree.Clr()
  x.parent = inf
  if x.me == x.root { // root sends the first message
    x.parent = x.root
    x.child[0] = true
    x.visited[0] = true
    x.tree.Ins (x.actVertex)
    x.tree.Mark (x.actVertex)
    x.tree.Write()
    pause()
    x.send (0, x.tree)
  }
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      s := x.recv (j).(Stream)
      x.tree = x.decodedGraph (s)
      if x.distance == j && x.tree.Eq (x.tmpGraph) { // tree unchanged back
        x.child[j] = false // from j-th netchannel, so assumption refused
      }
      x.tree.Write()
      pause()
      u := x.next(j) // == x.n, iff all neighbours != j are visited
      k := u
      if x.visited[j] { // echo
        if u == x.n {
          if x.me == x.root { // root must not reply any more
            done <- 0
            return
          }
          k = x.channel(x.parent) // send echo back to parent
        } else { // k == u < x.n, x.tree unchanged as probe to x.nr[u]
        }
      } else { // ! visited[j], i.e. probe
        if ! x.tree.Ex (x.actVertex) {
          x.tree.Ex (x.nb[j]) // nb[j] local in x.tree
          x.tree.Ins (x.actVertex) // actVertex local, nb[j] colocal in x.tree
          x.tree.Edge (x.directedEdge(x.nb[j], x.actVertex))
          x.tree.Write()
          pause()
        }
        x.visited[j] = true
        if x.parent == inf { // not for root
          x.parent = x.nr[j]
          if u == x.n { // all neighbours visited
            k = j // == x.channel(x.parent) echo back to sender == parent
          }
        } else {
          k = j // send echo back to sender
        }
      }
      x.visited[k] = true
      if k == u {
        x.distance = k // save for test
        x.tmpGraph.Copy (x.tree)
        x.child[k] = true // just an assumption
      }
      x.send (k, x.tree)
      done <- 0
    }(i)
  }
  for i := uint(0); i < x.n; i++ {
    <-done
  }
  x.tree.Ex (x.actVertex) // now x.actVertex is local in x.tree
  x.tree.Write()
  var s Stream
  if x.me == x.root {
    s = x.tree.Encode()
  } else {
    s = x.recv (x.channel(x.parent)).(Stream)
    x.tree = x.decodedGraph (s)
    x.tree.Write()
    pause()
  }
  for k := uint(0); k < x.n; k++ {
    if x.child[k] {
      x.send (k, s)
    }
  }
  x.tree.Write()
}
