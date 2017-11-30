package dgra

// (c) Christian Maurer   v. 171128 - license see nU.go

import . "nU/obj"

func (x *distributedGraph) dfs (o Op) {
  x.connect (x.time)
  defer x.fin()
  x.Op = o



  if x.me == x.root { // root sendet die erste Probe
    x.time = 0
    x.parent = x.root
    x.ch[0].Send (x.time)
    x.child[0] = true
    x.visited[0] = true
  }
  x.distance, x.diameter = x.n, inf
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      t := x.ch[j].Recv().(uint)
      mutex.Lock()
      if x.distance == j && x.diameter == t { // t unverändert aus
                                              // diesem Kanal zurück,
        x.child[j] = false // deshalb ist x.nr[j] kein Kind von x.me
      }
      u := x.next (j) // == x.n genau dann, wenn
                      // alle Nachbarn != j besucht sind
      k := u // Kanal für die nächste Botschaft
      if x.visited[j] { // d.h. Echo
        if u == x.n { // alle Nachbarn besucht
          t++
          x.time1 = t
          if x.me == x.root { // root darf nicht mehr antworten
            mutex.Unlock()
            done <- 0
            return
          }
          k = x.channel(x.parent) // Echo zurück zum Vater
        }
      } else { // ! x.visited[j], d.h. Probe





        x.visited[j] = true
        if x.parent == inf { // Vater ist undefiniert
                             // (nicht für root)
          x.parent = x.nr[j]
          t++
          x.time = t
          if u == x.n { // alle Nachbarn besucht
            t++
            x.time1 = t
            k = x.channel(x.parent) // Echo zurück zum Vater
          }
        } else { // Vater ist schon definiert
          k = j // Echo über Kanal j zurück, d.h. zum Absender
        }
      }
      x.visited[k] = true
      if k == u { // Probe senden
        x.distance, x.diameter = k, t
        x.child[k] = true // versuchsweise
      }
      x.ch[k].Send(t)




      mutex.Unlock()
      done <- 0
    }(i)
  }
  for i := uint(0); i < x.n; i++ {
    <-done
  }
  x.Op(x.me)
}
