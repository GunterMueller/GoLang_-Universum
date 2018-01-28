package kbd

// (c) Christian Maurer   v. 180106 - license see µU.go

import (
  "os"
  "µU/ker"
  "µU/term"
  "µU/mouse"
  "µU/navi"
)
var (
  keypipe chan byte
  mousepipe chan mouse.Command = nil
  navipipe chan navi.Command
  shiftC, ctrlC, altC, altGrC, fn bool
)

func catch() {
  shiftC, ctrlC, altC, altGrC, fn = false, false, false, false, false
  defer ker.Fin() // Hilft nix. Warum nicht ???
  for {
    b := term.Read()
    switch b { // case 0:
      // ker.Oops() // Fn-key combination !
    case shiftL, shiftR, shiftLock:
      shiftC = true
    case ctrlL, ctrlR:
      ctrlC = true
    case altL:
      altC = true
//    case doofL, doofM, doofR:
//      ctrlC = true
    case altR:
      altGrC = true
    case shiftLoff, shiftRoff, shiftLockoff:
      shiftC = false
    case ctrlLoff, ctrlRoff:
      ctrlC = false
    case altLoff:
      altC = false
    case altRoff:
      altGrC = false
//    case doofLoff, doofMoff, doofRoff:
//      altC = false
    case function:
      // println ("Fn-Key")
      fn = true
    default:
      if ctrlC && // (alt || altGr) && b == pause ||
                  b == 46 { // 'C'
        ker.Fin()
        os.Exit (1)
      } else if b < off && ctrlC && (altC || altGrC) {
        switch b {
        case left, right:
          ker.Console1 (b == right)
        case f1, f2, f3, f4, f5, f6, f7, f8, f9, f10:
          ker.Console (b - f1 + 1)
        case f11, f12:
          ker.Console (b - f11 + 11)
        case esc, back, tab, enter, roll, numEnter, pos1,
             up, pgUp, end, down, pgDown, ins, del:
          keypipe <- b
        }
      } else {
        keypipe <- b
      }
    }
  }
}

func input (b *byte, c *Comm, d *uint) {
  var (
    b0 byte
    k, k1 uint
    mc mouse.Command
    m3c navi.Command
    ok bool
  )
  loop: for {
    *b, *c, *d = 0, None, 0
    select {
    case mc = <-mousepipe:
      *c, *d = Go + Comm (mc), 0
      if shiftC || ctrlC { *d ++ }
      if altC || altGrC { *d += 2 }
      break loop
    case m3c = <-navipipe:
      *c, *d = Go + Comm (m3c), 0
      if shiftC || ctrlC { *d ++ }
      if altC || altGrC { *d += 2 }
      break loop
    case b0, ok = <-keypipe:
      if ok {
        k = uint(b0)
      } else {
        ker.Shit()
      }
    }
//    if k == 0 { ker.Shit() }
    k1 = k
    k = k % off
    if shiftC || ctrlC { *d ++ }
    if altC || altGrC { *d += 2 }
    switch b0 {
    case pgUp, pgDown:
      *d += 2
    }
    switch {
    case isAlpha (k):
      switch *d { case 0:
        *b = bb[k]
      case 1:
        *b = bB[k]
      case 2:
        *b = aa[k]
      default:
//        if altGr {
//          *b = aA[k]
//        } else {
        *b = aa[k]
//        }
      }
    case k == esc || k == numEnter || isCmd (k):
      *c = kK[k]
    case k == shiftLock:
      shiftC = true
    case isKeypad (k):
      if shiftC {
        *c = kK[k]
        switch k {
        case num9, num3:
          *d = 2
        }
      } else {
        *b = bb[k]
      }
    case k == function:
      // println ("Fn-Key")
    default:
      if k == 0 {
        // ignore
      } else {
        ker.Panic1 ("kbd.console", 10000 + k)
      }
    }
    if k1 < off { // key pressed, not released
      if *b == 0 {
        if *c > None {
          break loop
        }
      } else {
        lastbyte = *b
        *c = None
        break loop
      }
    }
  }
  lastcommand = *c
  lastdepth = *d
}

func initConsole() {
  term.New()
  ker.InstallTerm (func() { term.Fin() } )
  keypipe = make (chan byte, 256)
  if mouse.Ex() {
    mousepipe = mouse.Channel()
  }
  navipipe = navi.Channel()
  go catch()
}
