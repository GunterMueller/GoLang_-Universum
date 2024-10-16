package main

// (c) Christian Maurer   v. 241005 - license see µU.go

// >>> Codewandler als Filter: Caesars Chiffrierverfahren

import (
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/errh"
)

func cap (b byte) byte {
  if b >= 'a' { return b - 'a' + 'A' }
  return b
}

func dictate (t chan byte) {
  b := byte(0)
  S := uint(0)
  for b != '.' {
    b = kbd.Byte()
    scr.Lock()
    scr.Colours (col.FlashWhite(), col.Green())
    scr.Write1 (b, 0, S)
    scr.Unlock()
    S++
    t <- b
  }
}

func encrypt (t, c chan byte) {
  b := byte(0)
  for b != '.' {
    b = <-t
    if 'A' <= cap (b) && cap (b) <= 'W' {
      b += 3
    } else if 'X' <= cap (b) && cap (b) <= 'Z' {
      b -= 26 - 3
    }
    c <- b
  }
}

func send (c chan byte, d chan bool) {
  i := uint(0)
  b := byte(0)
  for b != '.' {
    b = <-c
    scr.Lock()
    scr.Colours (col.FlashWhite(), col.Red())
    scr.Write1 (b, 2, i)
    scr.Unlock()
    i++
  }
  d <- true
}

func main() {
  scr.NewWH (0, 0, 800, 600); defer scr.Fin()
  textchan := make (chan byte)
  cryptchan := make (chan byte)
  done:= make (chan bool)
  errh.Hint ("Tippen Sie einen Satz ein und beenden Sie ihn mit einem Punkt.")
  go dictate (textchan) // Caesar
  go encrypt (textchan, cryptchan) // Offizier
  go send (cryptchan, done) // Bote
  <-done
  errh.Error0 ("fertig")
}
