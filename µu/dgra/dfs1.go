package dgra

// (c) Christian Maurer   v. 170506 - license see µu.go

import
  . "µu/obj"
const (
  probe = uint(iota)
  echo
)

func (x *distributedGraph) dfs1 (o Op) {
  x.connect (nil)
  defer x.fin()
  x.Op = o
  x.tmpGraph.Copy (x.Graph)
  x.tree.Clr()
  x.parent = inf
  if x.me == x.root { // root sends the first message
    x.parent = inf + 1 // trick, see below
    x.tree.Ins (x.actVertex)
    x.tree.Ex(x.actVertex)
    x.tree.SubLocal()
    x.visited[0] = true
    x.child[0] = true
    x.ch[0].Send (x.tree.Encode())
  }
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      bs := x.ch[j].Recv().([]byte)
      mutex.Lock()
      x.tree = x.decodedGraph(bs)
      u := x.next(j) // == x.n, iff all neighbours != j are visited
      k := u
      if ! x.visited[j] { // probe
        x.tree.Ex(x.nb[j]) // nb[j] local in x.tree
        if ! x.tree.Ex(x.actVertex) {
          x.tree.Ins(x.actVertex) // MeVertex local, nb[j] colocal in x.tree
          x.tree.Edge (x.directedEdge(x.nb[j], x.actVertex))
        }
        x.visited[j] = true
        if x.parent == inf { // not for root - see trick
          x.parent = x.nr[j]
          if u == x.n { // all neighbours visited
            k = x.channel(x.parent) // send echo back to parent
          }
        } else {
          k = j // send echo back to sender
        }
      } else { // visited[j], i.e. received echo
        if u == x.n {
          if x.me == x.root { // root must not reply any more
            mutex.Unlock()
            done <- 0
            return
          }
          k = x.channel(x.parent) // send echo back to parent
        }
      }
      x.visited[k] = true
      x.ch[k].Send(x.tree.Encode())
      x.tree.Ex2(x.actVertex, x.nb[j])
      if x.tree.Edged() {
        x.child[j] = true
      }
      mutex.Unlock()
      done <- 0
    }(i)
  }
  for i := uint(0); i < x.n; i++ {
    <-done
  }
  x.tree.Ex(x.actVertex) // now x.actVertex is local in x.tree
  var bs []byte
  if x.me == x.root {
    x.Write()
    bs = x.tree.Encode()
    x.parent = x.me
  } else {
    bs = x.ch[x.channel(x.parent)].Recv().([]byte)
    x.tree = x.decodedGraph(bs)
  }
  x.tree.Ex(x.actVertex) // now x.actVertex is local in x.tree
  for k := uint(0); k < x.n; k++ {
    if x.child[k] {
      x.ch[k].Send(bs)
    }
  }
  x.Op (x.actVertex)
  x.tree.Write()
}
