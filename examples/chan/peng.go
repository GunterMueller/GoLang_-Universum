package main

// Zwei Fenster aufmachen, in einem das Programm _ohne_ Parameter starten,
// im anderen mit _irgendeinem_ Parameter (völlig egal, was).

import (
  "µU/time"
  "µU/rand"
  "µU/env"
  "µU/scr"
  "µU/errh"
  "µU/host"
  "µU/nchan"
)

func pause() {
  time.Sleep (rand.Natural(10))
}

func main() {
  eins := env.NArgs() > 0
  x := uint(0); if eins { x = 14 * 8 + 8}
  scr.NewWH (x, 0, 14 * 8, 4 * 16); defer scr.Fin()
  hostname := "terra"
  h := host.New(); h.Defined(hostname)
  me, i := uint(0), uint(1)
  if eins { me, i = 1, 0 }
  c := nchan.New("wort", me, i, h.String(), 50000)
  for i:= uint(0); i < 3; i++ {
    if eins {
      errh.Hint ("will senden")
      c.Send ("ping")
      errh.Hint ("habe gesendet")
      pause()
      errh.Hint ("will empfangen")
      scr.Write (c.Recv().(string), i, 0)
      errh.Hint ("habe empfangen")
      pause()
    } else {
      errh.Hint ("will empfangen")
      scr.Write (c.Recv().(string), i, 0)
      errh.Hint ("habe empfangen")
      pause()
      errh.Hint ("will senden")
      c.Send ("pong")
      errh.Hint ("habe gesendet")
      pause()
    }
  }
  errh.Error0 ("fertig")
}
/*
Ist deine Idee vollduplex, also kann ein Prozess darauf Senden _und_ Empfangen,
sodass folgendes Szenario nicht geschehen kann?

  Kanal: []
  Prozess0 sendet "hi"
  Kanal: ["hi"]
  Prozess0 liest: "hi"
  Kanal: []
  Prozess1 liest: <am Warten auf eine Nachricht>

Könnte mir gut vorstellen, dass, wenn man einen einzelnen Port verwendet,
aufgrund des Zugriffs auf den gleichen Speicher ein Prozess seine eigene
Nachricht auslesen könnte (Ports sind ja eigentlich schon eine Form
des Multiplexens, den Port zu Multiplexen kommt mir sehr merkwürdig vor,
zumindest für diese Art der Anwendung).)
>     
Falls das nicht passieren kann, kannst du mir nochmals sagen, was du da getan hast?
(Hoffe ich habe nicht falsch in Erinnerung, dass du da was ausprobiert hattest.)
*/
