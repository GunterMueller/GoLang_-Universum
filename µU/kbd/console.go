package kbd

// (c) Christian Maurer   v. 210105 - license see µU.go

import (
  "os"
  "µU/ker"
  "µU/terminal"
  "µU/mouse"
//  "µU/navi"
)
var (
  keypipe chan byte
  mousepipe chan mouse.Command = nil
//  navipipe chan navi.Command
)

func catch() {
//  shift, ctrl, alt, altGr, fn = false, false, false, false, false
  shift, ctrl, alt, altGr = false, false, false, false
  defer ker.Fin() // hilft nich
  for {
    b := terminal.Read()
    switch b {
    // case 0:
      // ker.Oops() // Fn-key combination !
    case shiftL, shiftR, shiftLock:
      shift = true
    case ctrlL, ctrlR:
      ctrl = true
    case altL:
      alt = true
//    case doofR:
//      backslash
    case altR:
      altGr = true
    case shiftLoff, shiftRoff, shiftLockoff:
      shift = false
    case ctrlLoff, ctrlRoff:
      ctrl = false
    case altLoff:
      alt = false
    case altRoff:
      altGr = false
//    case function:
      // println ("Fn-Key")
//      fn = true
    default:
      if ctrl && // (alt || altGr) && b == pause ||
                  b == 46 { // C
        ker.Fin()
        os.Exit (1)
      } else if b < off && ctrl && (alt || altGr) {
        switch b {
        case left, right:
          ker.SwitchConsole (b == right)
        case f1, f2, f3, f4, f5, f6, f7, f8, f9, f10:
          ker.Console (b - f1 + 1)
        case f11, f12:
          ker.Console (b - f11 + 11)
        case esc, back, tab, enter, roll, numEnter, pos1, up, pgUp, end, down, pgDown, ins, del:
          keypipe <- b
        }
      } else {
        keypipe <- b
      }
    }
  }
}

func inputC (B *byte, C *Comm, D *uint) {
  var (
    b0 byte
    k, k1 uint
    mc mouse.Command
//    m3c navi.Command
    ok bool
  )
loop:
  for {
    *B, *C, *D = 0, None, 0
    select {
    case mc = <-mousepipe:
      *C, *D = Go + Comm (mc), 0
      if shift || ctrl {
        *D = 1
      }
      if alt || altGr {
        *D = 2
      }
      break loop
/*/
    case m3c = <-navipipe:
      *C, *D = Go + Comm (m3c), 0
      if shift || ctrl {
        *D = 1
      }
      if alt || altGr {
        *D = 2
      }
      break loop
/*/
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
    if shift || ctrl {
      *D = 1
    }
    if alt || altGr {
      *D += 2
    }
    switch b0 {
    case pgUp, pgDown:
      *D = 2
    }
    switch {
    case isAlpha (k):
      if *D == 0 {
        *B = bb[k]
      } else {
        *B = aa[k]
      }
    case k == esc || k == numEnter || isCmd (k):
      *C = kK[k]
    case k == shiftLock:
      shift = true
    case isKeypad (k):
      if shift {
        *C = kK[k]
        switch k {
        case num9, num3:
          *D = 2
        }
      } else {
        *B = bb[k]
      }
    default:
      switch k {
      case 0:
        ; // ignore
      case doofL:
        ;
      case doofR:
        ;
      case windoof:
        ;
      default:
        ker.Panic1 ("kbd.console", 10000 + k) // z.B. für k == 125
      }
    }
    if k1 < off { // key pressed, not released
      if *B == 0 {
        if *C > None {
          break loop
        }
      } else {
        lastbyte = *B
        *C = None
        break loop
      }
    }
  }
  lastcommand = *C
  lastdepth = *D
}

func initConsole() {
  terminal.New()
  ker.InstallTerm (func() { terminal.Fin() } )
  keypipe = make (chan byte, 256)
  if mouse.Ex() {
    mousepipe = mouse.Channel()
  }
//  navipipe = navi.Channel()
  go catch()
}
