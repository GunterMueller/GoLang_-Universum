package main

// >>> "philosophen" must be called on a GUI with maximal big window !
//     After a call on a tty-console only helps - if the Magic-SysRq-Keys are
//     activate - the simultanous pressing of the keys <Alt>, <SysRq> and <R>
//     - if not, the computer has to be restarted with the reset-button.
//
//     A philosopher is set to the table by pressing its number (0..4)
//     or the first letter of its name (program end with <Esc>).

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
  scr.ColourF (col.FlashWhite())
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
