package kbd

// (c) Christian Maurer   v. 220206 - license see µU.go

import (
  "os"
  "µU/ker"
  "µU/kbd/internal"
  "µU/mouse"
//  "µU/navi"
)
var (
  keypipe chan byte
//  navipipe chan navi.Command
)

func catch() {
  shift, ctrl, alt, altGr = false, false, false, false
  defer ker.Fin() // hilft nich
  for {
    b := internal.Read()
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
//    nc navi.Command
    ok bool
  )
  loop:
  for {
    *B, *C, *D = 0, None, 0
    select {
    case mc = <-mouse.Pipe:
      switch mc {
      case mouse.None:
      case mouse.Go:
        *C = Go
      case mouse.Here:
        *C = Here
      case mouse.Drag:
        *C = Drag
      case mouse.To:
        *C = To
      case mouse.This:
        *C = This
      case mouse.Drop:
        *C = Drop
      case mouse.There:
        *C = There
      case mouse.That:
        *C = That
      case mouse.Move:
        *C = Move
      case mouse.Thither:
        *C = Thither
      }
      if shift || ctrl {
        *D = 1
      }
      if alt || altGr {
        *D = 2
      }
      break loop
/*/
    case nc = <-navi.Pipe:
      *C, *D = Go + Comm (nc), 0
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
        ker.Panic1 ("kbd.console", 10000 + k) // e.g. for k == 125
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
  lastdepth = *D
  lastcommand = *C
}

func initConsole() {
  internal.New()
  ker.InstallTerm (func() { internal.Fin() } )
  keypipe = make (chan byte, 256)
//  navipipe = navi.Channel()
  go catch()
}
