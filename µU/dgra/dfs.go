package dgra

// (c) Christian Maurer   v. 171206 - license see µU.go

import
  . "µU/obj"

func (x *distributedGraph) dfs (o Op) {
  x.connect (x.time) // Netzkanäle sind vom Typ uint
  defer x.fin()
  x.Op = o
  if x.me == x.root {
    x.parent = x.root
    x.time = 0
    x.child[0] = true
    x.visited[0] = true
    x.ch[0].Send(x.time) // root sendet als Erster
  }
  x.distance, x.diameter = x.n, inf
  for i := uint(0); i < x.n; i++ {
    go func (j uint) {
      t := x.ch[j].Recv().(uint)
// x.log2("recv", t, "from", x.nr[j])
      mutex.Lock()
      if x.distance == j && x.diameter == t { // t unverändert
                           // aus dem j-ten Netzkanal zurück,
        x.child[j] = false // deshalb ist x.nr[j] kein Kind von x.me
      }
      u := x.next(j) // == x.n genau dann, wenn alle
                     // Netzkanäle != j schon markiert sind
      k := u // Kanal für die nächste Sendung
      if x.visited[j] { // d.h. Echo
        if u == x.n { // alle Netzkanäle markiert
          t++
          x.time1 = t
          if x.me == x.root { // root darf kein Echo mehr senden
            mutex.Unlock()
            done <- 0
            return
          }
          k = x.channel(x.parent) // Echo x.time1 zurück an Absender
        } else {
          // k == u < x.n, t unverändert als Probe weiter an x.nr[u]
        }
      } else { // ! x.visited[j], d.h. Probe
        x.visited[j] = true
        if x.parent < inf { // Vater schon definiert
          k = j // t unverändert als Echo zurück zum Absender
        } else { // x.parent == inf, d.h. Vater noch
                 // undefiniert (nicht für root!)
          x.parent = x.nr[j]
          t++
          x.time = t // wenn u < x.n Probe x.time weiter an x.nr[u]
          if u == x.n { // alle Netzkanäle markiert
            t++
            x.time1 = t
            k = j // Echo x.time1 zurück zum Absender (Vater)
          }
        }
      }
      x.visited[k] = true
      if k == u {
        x.distance, x.diameter = k, t // k, t für o.a. Prüfung retten
        x.child[k] = true // versuchsweise
      }
      x.ch[k].Send(t) // für k == u Probe, sonst Echo
      mutex.Unlock()
      done <- 0
    }(i)
  }
  for i := uint(0); i < x.n; i++ { // Beendigung aller
    <-done                         // Goroutinen abwarten
  }
  x.Op(x.me) // übergebene Operation ausführen
}
