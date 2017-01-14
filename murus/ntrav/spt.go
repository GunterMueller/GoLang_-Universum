package ntrav

// (c) murus.org  v. 170107 - license see murus.go

import (
//  "murus/ker"
  . "murus/obj"
  "murus/scr"
  "murus/errh"
  "murus/nchan"
  "murus/bnat"
  "murus/node"
  "murus/gra"
)

func (x *netTraversal) log (k, t uint, sent, probe bool) {
  if ! x.bool { return }
  a, b := "recv", "from"; if sent { a, b = "send", " to " }
/* TODO
  if sent {
    if probe {
//      x.graph.Step (k, true) // XXX  x.path[0] != x.colocal
    } else { // echo
      x.graph.Step (k, false)
    }
  } else { // received
    if probe {
//      x.graph.Step (k, true) // XXX  x.path[0] != x.colocal
      x.graph.Colocate()
    } else { // echo
      x.graph.Step (k, false)
    }
  }
*/
  x.Write()
  if sent {
    errh.Hint2 (a, t, b, x.nr[k])
  } else {
    errh.Error2 (a, t, b, x.nr[k])
  }
}

func (x *netTraversal) showTree() {
  scr.Cls(); x.tree.Trav3Cond (node.O, node.O3)
}

func (x *netTraversal) next(i uint) uint {
  return x.nn[i].Val()
}

// Pre: scr.New (constr. of channels, log)
//      nh is sorted by the identities of the nodes hosts (is done in murus/lan!).
func (x *netTraversal) spt (op Op) {
  ch := make ([]nchan.NetChannel, x.n)
  for i := uint(0); i < x.n; i++ {
    ch[i] = nchan.New (nil, x.me, x.nr[i], x.h[i], x.port[i])
  }
  known := make([]bool, x.n)
  var s string
  parent := x.n
  done := make (chan int)
  x.Write()
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      probe := true // false means echo
      if x.me == x.root && // the root node sends the first message
         j == 0 { // but only once
        op(nil)
        parent = x.n + 1 // trick no. 17
        known[0] = true
        x.tree.Ins (x.meNode)
        bs := x.tree.Encode()
        probe = true
        ch[0].Send (bs)
        errh.Hint1("sent probe to", x.next(0))
      }
      bs := ch[j].Recv().([]byte)
      if len(bs) == 0 {
        errh.Error0(">>> partner off") // XXX
      }
      x.tree = Decode(gra.New (true, x.tmpNode, nil), bs).(gra.Graph)
      x.showTree()
//      probe = ! x.tree.Ex(x.meNode)
      probe = ! known[j]
      if x.bool {
        s = "echo"; if probe { s = "probe" }
        errh.Error("received " + s + " from", x.next(j))
      }
      u := x.n // channel to next unvisited neighbour == x.n, if all visited
      for k := uint(0); k < x.n; k++ { // x.tree.Trav...(...) ?
        if k != j && ! known[k] {
          u = k
          break
        }
      }
      var k uint // channel to neighbour for the next send operation
      if probe { // i.e. ! x.tree.Ex(x.meNode)
        x.tree.Ex(x.nb[j]) // x.nb[j] local in x.tree
        if ! x.tree.Ex(x.meNode) {
          x.tree.Ins(x.meNode)
          x.tree.Edge()
        }
        known[j] = true
        x.showTree()
        if parent == x.n {
          op(nil)
          parent = j
          if u == x.n { // all neighbours visited
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
      } else { // x.tree.Ex(x.meNode), i.e. received echo
        if u == x.n {
        if x.bool { errh.Error0("all neighbours visited") }
          if x.me == x.root { // the root node must not reply any more
            x.showTree()
            if x.bool { errh.Error ("the spanning tree", j) }
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
      known[k] = true
      ch[k].Send (x.tree.Encode())
      if x.bool {
        s = "echo"; if probe { s = "probe" }
        errh.Hint1("sent " + s + " to", x.next(k))
      }
      done <- 0
    }(i)
  }
  for i := uint(0); i < x.n; i++ {
    <-done
  }

  var bs []byte
  if x.me == x.root { bs = x.tree.Encode() }
  x.tree.Ex (x.meNode) // now x.meNode is local in x.tree
  in := x.tree.NumLocInv()
  if in == 1 { // x.me != x.root
    if x.bool { errh.Hint("waiting to receive bs") }
    inNr := x.tree.CoNeighbour(0).(node.Node).Content().(bnat.Natural).Val()
    for i := uint(0); i < x.n; i++ {
      if x.next(i) == inNr {
        bs = ch[i].Recv().([]byte)
        x.tree = Decode(gra.New (true, x.tmpNode, nil), bs).(gra.Graph)
        x.showTree()
        if x.bool { errh.Error("received from", x.next(i)) }
        break
      }
    }
  }
  x.tree.Ex (x.meNode) // now x.meNode is local in x.tree
  out := x.tree.NumLoc()
  if out > 0 { // there are outgoing edges
    outNr := make([]uint, out)
    for j := uint(0); j < out; j++ {
      outNr[j] = x.tree.Neighbour(j).(node.Node).Content().(bnat.Natural).Val()
      for i := uint(0); i < x.n; i++ {
        if x.next(i) == outNr[j] {
          ch[i].Send (bs)
          if x.bool { errh.Error("sent to", x.next(i)) }
        }
      }
    }
  }
}
