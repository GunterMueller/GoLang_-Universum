package dgra

// (c) Christian Maurer   v. 171129 - license see µU.go

import . "nU/obj"

func (x *distributedGraph) dfs1 (o Op) {
  x.connect (nil)
  defer x.fin()
  x.Op = o
  x.tmpGraph.Copy (x.Graph)
  x.tree.Clr()
  x.parent = inf
  if x.me == x.root {
    x.parent = x.root
    x.tree.Ins (x.me)
    x.tree.Mark (x.me)
    x.visited[0] = true
    x.child[0] = true
    x.ch[0].Send (x.tree.Encode())
  }
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      bs := x.ch[j].Recv().(Stream)
      mutex.Lock()
      x.tree = x.decodedGraph(bs)



      u := x.next(j)

      k := u
      if x.visited[j] {
        if u == x.n {


          if x.me == x.root {
            mutex.Unlock()
            done <- 0
            return
          }
          k = x.channel(x.parent)
        }
      } else {
        x.tree.Ex (x.nr[j])
        if ! x.tree.Ex (x.me) {
          x.tree.Ins (x.me)
          x.tree.Edge (x.edge(x.nr[j], x.me))
        }
        x.visited[j] = true
        if x.parent == inf {

          x.parent = x.nr[j]


          if u == x.n {


            k = x.channel(x.parent)
          }
        } else {
          k = j
        }
      }
      x.visited[k] = true




      x.ch[k].Send (x.tree.Encode())
      x.tree.Ex2 (x.me, x.nr[j])
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
  x.tree.Ex(x.me)
  var bs Stream
  if x.me == x.root {
    x.Write()
    bs = x.tree.Encode()
    x.parent = x.me
  } else {
    bs = x.ch[x.channel(x.parent)].Recv().(Stream)
    x.tree = x.decodedGraph (bs)
  }
  x.tree.Ex (x.me)
  for k := uint(0); k < x.n; k++ {
    if x.child[k] {
      x.ch[k].Send(bs)
    }
  }
  x.Op (x.me)
  x.tree.Write()
}
