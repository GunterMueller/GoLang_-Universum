package dgra

// (c) Christian Maurer   v. 231229 - license see µU.go

// >>> Helary, J. M., Raynal, M.:
//     Depth-first traversal and virtual ring construction in distributed Systems.
//     Research Rreport RR-0704, INRIA, 1987

import
  . "µU/obj"
const (
  DISCOVER = uint(iota)
  RETURN
)
var
  chanus = make(chan UintStream)

func (x *distributedGraph) HelaryRaynal() {
  x.connect (nil)
  defer x.fin()
  if x.me == x.root {
    x.parent = x.root
    us := append(UintStream {DISCOVER}, x.me)
    x.send (0, us)
    x.child[0] = true
  }
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      loop:
      for {
        t := x.recv (j) // dirty trick to terminate the algorithm:
        if t == nil {   // no more input available or use of closed network connection:
          chanus <- nil // then force termination of for-loop below
          break loop    // and terminate process
        }
        us := Decode (UintStream{}, t.(Stream)).(UintStream)
        chanus <- append (UintStream{j}, us...)
      }
    }(i)
  }
  for {
    us := <-chanus
    if us == nil {
      break
    }
    j := us[0]
    us = us[1:]
    neighbours := us[1:]
    existUnvisitedNeighbours := false
    for i := uint(0); i < x.n; i++ {
      for _, n := range neighbours {
        if n == x.nr[i] {
          x.visited[i] = true
        }
      }
      if ! x.visited[i] {
        existUnvisitedNeighbours = true
      }
    }
    k := x.n // channel number of first unvisited neighbour
    if existUnvisitedNeighbours {
      for i := uint(0); i < x.n; i++ {
        if ! x.visited[i] {
          k = i
          break
        }
      }
    }
    if us[0] == DISCOVER {
      x.parent = x.nr[j]
      us = append(us, x.me)
      if ! existUnvisitedNeighbours {
        us[0] = RETURN
        x.send (j, us)
      } else { // existUnvisitedNeighbours
        x.send (k, us) // DISCOVER
        x.child[k] = true
      }
    } else { // us[0] == RETURN
      if existUnvisitedNeighbours {
        us[0] = DISCOVER
        x.send (k, us)
        x.child[k] = true
      } else { // ! existUnvisitedNeighbours
        if x.parent == x.me {
          return
        } else {
          x.send (x.channel(x.parent), us)
        }
      }
    }
  }
}
