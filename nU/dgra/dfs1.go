package dgra

// (c) Christian Maurer   v. 171227 - license see nU.go

import . "nU/obj"

func (x *distributedGraph) dfs1 (o Op) {
  x.connect (nil)
  defer x.fin()
  x.Op = o
  x.tree.Clr()
  x.parent = inf
  if x.me == x.root {
    x.parent = x.root
    x.child[0] = true
    x.visited[0] = true
    x.tree.Ins (x.actVertex)
    x.tree.Mark (x.actVertex)
    x.tree.Write()
    pause()
    x.ch[0].Send (x.tree)
  }
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      bs := x.ch[j].Recv().(Stream)
      x.tree = x.decodedGraph(bs)
      if x.distance == j && x.tree.Eq (x.tmpGraph) {
        x.child[j] = false
      }
      x.tree.Write()
      pause()
      u := x.next(j)
      k := u
      if ! x.visited[j] {
        if ! x.tree.Ex (x.actVertex) {
          x.tree.Ex (x.nb[j])
          x.tree.Ins (x.actVertex)
          x.tree.Edge (x.edge(x.nb[j], x.actVertex))
          x.tree.Write()
          pause()
        }
        x.visited[j] = true
        if x.parent == inf {
          x.parent = x.nr[j]
          if u == x.n {
            k = j // == x.channel(x.parent)
          }
        } else {
          k = j
        }
      } else {
        if u == x.n {
          if x.me == x.root {
            done <- 0
            return
          }
          k = x.channel(x.parent)
        }
      }
      x.visited[k] = true
      if k == u {
        x.distance = k
        x.tmpGraph.Copy (x.tree)
        x.child[k] = true
      }
      x.ch[k].Send (x.tree)
      done <- 0
    }(i)
  }
  for i := uint(0); i < x.n; i++ {
    <-done
  }
  x.tree.Ex (x.actVertex)
  var bs Stream
  if x.me == x.root {
    bs = x.tree.Encode()
//    x.parent = x.me
  } else {
    bs = x.ch[x.channel(x.parent)].Recv().(Stream)
    x.tree = x.decodedGraph (bs)
    x.tree.Write()
    pause()
  }
  for k := uint(0); k < x.n; k++ {
    if x.child[k] {
      x.ch[k].Send (bs)
    }
  }
  x.Op (x.me)
  x.tree.Write()
}
