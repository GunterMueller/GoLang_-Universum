package dgra

// (c) Christian Maurer   v. 231220 - license see nU.go

func (x *distributedGraph) Dfs() {
  x.connect (x.time) // Netzkanäle sind vom Typ uint
  defer x.fin()
  if x.me == x.root {
    x.parent = x.root
    x.time = 0
    x.ch[0].Send(x.time) // (a) root sendet als Erster
    x.child[0] = true
    x.visited[0] = true
  }
  x.distance, x.diameter = x.n, inf
  for i := uint(0); i < x.n; i++ { // (x.n == Anzahl der Nachbarn !)
    go func (j uint) {
      t := x.ch[j].Recv().(uint)
      mutex.Lock()
      if x.distance == j && x.diameter == t { // t unverändert
                           // aus dem j-ten Netzkanal zurück,
        x.child[j] = false // also ist x.nr[j] kein Kind von x.me
      }
      u := x.next(j) // == x.n genau dann, wenn alle
                     // Netzkanäle != j schon markiert sind
      k := u // Kanal für die nächste Sendung
      if x.visited[j] { // d.h. Echo
        if u == x.n { // kein Netzkanal ist mehr unmarkiert
          t++
          x.time1 = t
          if x.me == x.root { // root darf kein Echo mehr senden
            mutex.Unlock()
            done <- 0
            return
          }
          k = x.channel(x.parent) // (b) t == x.time1, Echo zum Vater
        } else {
          // k == u < x.n, t unverändert als Probe an x.nr[u]
        }
      } else { // ! x.visited[j], d.h. Probe
        x.visited[j] = true
        if x.parent < inf { // Vater schon definiert
          k = j // (d) t unverändert als Echo zum Absender
        } else { // x.parent == inf, d.h. Vater noch
                 // undefiniert (nicht für root!)
          x.parent = x.nr[j]
          t++
          x.time = t // wenn u < x.n Probe x.time weiter an x.nr[u]
          if u == x.n { // alle Netzkanäle markiert
            t++
            x.time1 = t // Echo x.time1 zum Absender (Vater)
            k = j
          }
        }
      }
      x.visited[k] = true
      if k == u {
        x.distance, x.diameter = k, t // k und t für o.a. Prüfung retten
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
}
