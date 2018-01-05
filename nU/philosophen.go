package main

// (c) Christian Maurer   v. 171229 - license see nU.go

// >>> "philosophen" muss auf einer GUI mit maximal großem Fenster aufgerufen werden !
//     Nach einem Aufruf auf einer tty-Konsole hilft nur - wenn die Magic-SysRq-Keys
//     aktiviert sind - das gleichzeitige Drücken der Tasten <Alt>, <SysRq> und <R>
//     - wenn nicht, muss der Rechner mit dem Reset-Knopf neu gestartet werden.
//
//     Ein Philosoph wird durch das Tippen seiner Nummer (0..4) oder des Anfangs-
//     buchstabens seines Namens an den Tisch gesetzt (Programmende mit <Esc>).

import ("time"; "nU/term"; "nU/col"; "nU/scr"; . "nU/phil")

var ph Philos

func eat (i uint) {
  ph.Lock (i)
  time.Sleep (10e9)
  ph.Unlock (i)
}

func main() {
/*
  ph = NewBounded()
  ph = NewUnsymmetric()
  ph = NewSemaphoreUnfair()
  ph = NewSemaphoreFair()
  ph = NewCriticalSection()
  ph = NewMonitorUnfair()
  ph = NewMonitorFair()
  ph = NewCondMonitor()
  ph = NewChannel()
  ph = NewChannelUnsymmetric()
*/
  ph = NewCriticalSection()
  scr.New(); defer scr.Fin()
  term.New(); defer term.Fin()
  Start()
  scr.ColourF (col.LightWhite())
  scr.Write ("Philosoph", 0, 0)
  loop: for {
    scr.Switch (true)
    scr.Warp (0, 10)
    b := term.Read()
    switch b {
    case '0', 'p', 'P':
      go eat(0)
    case '1', 's', 'S':
      go eat(1)
    case '2', 'a', 'A':
      go eat(2)
    case '3', 'c', 'C':
      go eat(3)
    case '4', 'h', 'H':
      go eat(4)
    case term.Esc:
      break loop
    }
  }
}
