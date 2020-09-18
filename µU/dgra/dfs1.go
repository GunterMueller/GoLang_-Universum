package dgra

// (c) Christian Maurer   v. 200728 - license see µU.go

import
  . "µU/obj"

func (x *distributedGraph) dfs1 (o Op) {
  x.connect (nil)
  defer x.fin()
  x.Op = o
  x.tree.Clr()
  x.parent = inf
  if x.me == x.root { // root sends the first message
    x.parent = x.root
    x.child[0] = true
    x.visited[0] = true
    x.tree.Ins (x.actVertex)
    x.tree.Mark (x.actVertex)
    x.tree.Write()
//    x.ch[0].Send (x.tree)
    x.send (0, x.tree)
// x.log("sent to", x.nr[0])
  }
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      bs := x.ch[j].Recv().(Stream)
      x.tree = x.decodedGraph(bs)
      if x.distance == j && x.tree.Eq (x.tmpGraph) { // tree unchanged back
        x.child[j] = false // from j-th netchannel, so assumption refused
      }
      x.tree.Write()
// x.log("rec from", x.nr[j])
      u := x.next(j) // == x.n, iff all neighbours != j are visited
      k := u
      if x.visited[j] { // echo
        if u == x.n {
          if x.me == x.root { // root must not reply any more
            done <- 0
            return
          }
          k = x.channel(x.parent) // send echo back to parent
        } else {
          // k == u < x.n, x.tree unchanged as probe to x.nr[u]
        }
      } else { // ! visited[j], i.e. probe
        if ! x.tree.Ex (x.actVertex) {
          x.tree.Ex (x.nb[j]) // nb[j] local in x.tree
          x.tree.Ins (x.actVertex) // actVertex local, nb[j] colocal in x.tree
          x.tree.Edge (x.directedEdge(x.nb[j], x.actVertex))
          x.tree.Write()
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
// x.log("send to", x.nr[k])
//      x.ch[k].Send (x.tree)
      x.send (k, x.tree)
      if k == u {
        x.distance = k // save for test
        x.tmpGraph.Copy (x.tree)
        x.child[k] = true // just an assumption
      }
// x.l("sent to", x.nr[k])
      done <- 0
    }(i)
  }
  for i := uint(0); i < x.n; i++ {
    <-done
  }
  x.tree.Ex (x.actVertex) // now x.actVertex is local in x.tree
  x.tree.Write()
  var bs Stream
  if x.me == x.root {
    bs = x.tree.Encode()
  } else {
    bs = x.ch[x.channel(x.parent)].Recv().(Stream)
    x.tree = x.decodedGraph (bs)
//    j := nrLocal (x.tree)
// x.log("Recv", j) // from", x.nr[j])
  }
//  x.tree.Write()
//  x.tree.Ex (x.actVertex) // now x.actVertex is local in x.tree
  for k := uint(0); k < x.n; k++ {
    if x.child[k] {
// x.log("Send to", x.nr[k])
//      x.ch[k].Send (bs)
      x.send (k, bs)
    }
  }
  x.Op (x.actVertex)
  x.tree.Write()
}
