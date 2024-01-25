package dgra

// (c) Christian Maurer   v. 231220 - license see nU.go

const (
  label = uint(iota)
  keepon
  stop
  end
  term
)

// Liefert die Anzahl der Nachbarn, an die
// noch label-Botschaften gesendet werden müssen.
func (x *distributedGraph) numSendTos() uint {
  s := uint(0)
  for k := uint(0); k < x.n; k++ {
    if x.sendTo[k] {
      s++
    }
  }
  return s
}

// Liefert genau dann true, wenn für alle Nachbarn gilt,
// dass keine label-Botschaften mehr an ihn zu senden sind
// oder er schon auf eine label-Botschaft reagiert hat.
func (x *distributedGraph) allSendTosEchoed() bool {
  for k := uint(0); k < x.n; k++ {
    if x.sendTo[k] && ! x.echoed[k] {
      return false
    }
  }
  return true
}

func (x *distributedGraph) Bfs() {
  x.connect (uint(0)) // alle Botschaften sind vom Typ uint
  defer x.fin() // Schliessen aller Kanäle am Ende vorbereiten
  m := inf * x.me // um nicht immer dieses Produkt auszurechnen
  if x.me == x.root {
    x.parent = x.root // root ist sein eigener Vater, der
    x.labeled = true  // muss nicht auf label-Botschaften warten
    x.distance = 0    // Abstand von sich selbst
    for i := uint(0); i < x.n; i++ { // für alle Nachbarn
      x.child[i] = false // - die sind (noch) kein Kind - eine
      x.ch[i].Send (label + 8 * x.distance + m) // label-Botschaft
      x.echoed[i] = false // senden, aber die Reaktion steht noch aus
      x.sendTo[i] = true // später noch weitere label-Botschaften senden
    }
  }
  done = make(chan int, x.n)
  for j := uint(0); j < x.n; j++ { // für jeden Nachbarkanal
    go func (i uint) { // eine Goroutine abzweigen
      loop: // die in einer Endlosschleife arbeitet
      for {
        t := x.ch[i].Recv().(uint)
        if t % 8 == term { // wenn eine term-Botschaft da ist,
          break loop       // Schleife verlassen,
        } else {           // andernfalls die Botschaft
          x.chan1 <- t     // über chan1 weitersenden
        }
      }
      done <- 1 // eine Botschaft zum Beendigen schicken
    }(j)
  }
  for {
    t := <-x.chan1 // Botschaft aus chan1 auslesen
    j := x.channel (t / inf) // Kanalnummer des Absenders
    t %= inf
    x.distance = t / 8
    switch t % 8 { // Art der Botschaft
    case label: // Fall 1
      if ! x.labeled { // das war die erste label-Botschaft
        x.labeled = true
        x.parent = x.nr[j] // Absender ist Vater
        x.distance++       // Zeit = empfangene Zeit + 1
        for k := uint(0); k < x.n; k++ { // an alle Nachbarn
          x.sendTo[k] = k != j           // außer dem Absender
        }             // sind noch label-Botschaften zu senden
        if x.n == 1 { if x.numSendTos() > 0 { panic("Pisse") }
          x.ch[j].Send (end + m) // keine weiteren Nachbarn mehr
        } else {
          x.ch[j].Send (keepon + m) // mach weiter, Vater!
        }
      } else { // schon vorher label-Botschaften erhalten
        if x.parent == x.nr[j] { // Absender ist Vater
          for k := uint(0); k < x.n; k++ { // an alle Nachbarn,
            if x.sendTo[k] { // an die noch label-Botschaften
              // zu senden sind, wird eine weitere gesendet
              x.ch[k].Send (label + 8 * x.distance + m)
              x.echoed[k] = false // aber deren Reaktion
            }                     // steht noch aus
          }
        } else { // Absender jemand anders als Vater
          x.ch[j].Send (stop + m) // Absender: aufhören!
        }
      }
    case keepon: // Fall 2
      x.echoed[j] = true // Absender hat reagiert
      x.child[j] = true  // Absender ist Kind
    case stop: // Fall 3
      x.echoed[j] = true // Absender hat reagiert
      if x.nr[j] == x.parent { // wenn Absender der Vater ist,
        for k := uint(0); k < x.n; k++ { // dann an alle
          if x.child[k] { // Kinder die Botschaft senden:
            x.ch[k].Send (stop + m) // aufhören
          }
        }
        for k := uint(0); k < x.n; k++ { // für alle Nachbarn die
          x.ch[k].Send (term) // Beendigung der Goroutine veranlsssen
        }
        for k := uint(0); k < x.n; k++ { // warten,
          <-done // bis alle Nachbarn fertig sind
        } // dann verlässt der aufrufende Prozess
        return // die Schleife, d.h., er terminiert
      } else { // Absender - nicht der Vater - hat reagiert,
        x.echoed[j] = true
        x.sendTo[j] = false // also keine labels mehr an ihn
      }
    case end: // Fall 4
      x.echoed[j] = true  // Absender hat reagiert
      x.child[j] = true   // Absender ist Kind
      x.sendTo[j] = false // keine labels mehr an ihn
    }
    if x.numSendTos() == 0 { // im Grunde fertig
      if x.me == x.root {
        for k := uint(0); k < x.n; k++ {
          if x.child[k] {
            x.ch[k].Send (stop + m)
          }
        }
        for k := uint(0); k < x.n; k++ { // s. Fall 3
          x.ch[k].Send(term)
        }
        for k := uint(0); k < x.n; k++ {
          <-done
        }
        return // Algorithmus terminiert für root
      } else { // aufrufender Prozess ist nicht root
        k := x.channel(x.parent) // Kanal zum Vater
        x.ch[k].Send (end + m) // Botschaft an ihn: Ende
      }
    } else { // x.numSendTos() > 0: weiter in der Tiefensuche
      if x.allSendTosEchoed() { // für alle Nachbarn gilt,
        // dass keine label-Botschaften mehr an sie zu senden sind
        // oder sie schon auf eine label-Botschaft reagiert haben
        if x.me == x.root {
          for k := uint(0); k < x.n; k++ { // an alle Nachbarn,
            if x.sendTo[k] { // an die noch label-Botschaften
              // zu senden sind, wird eine weitere gesendet
              x.ch[k].Send (label + 8 * x.distance + m)
              x.echoed[k] = false // Reaktions steht noch aus
            }
          }
        } else { // aufrufender Prozess ist nicht root
          k := x.channel(x.parent) // Kanal zum Vater
          x.ch[k].Send (keepon + m) // Botschaft an ihn: weitermachen
        }
      }
    }
  }
}
