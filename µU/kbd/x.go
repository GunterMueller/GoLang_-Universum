package kbd

// (c) Christian Maurer   v. 220530 - license see µU.go

// #cgo LDFLAGS: -lX11
// #include <X11/X.h>
import
  "C"
import (
  "os"
  "µU/time"
  "µU/scr"
)
var (
  xpipe chan scr.Event
  ch chan int = make (chan int, 1)
)

func lock() {
  <-ch
}

func unlock() {
  ch <- 0
}

// Pre: scr.x.initialized == true
func catchX() {
  for scr.Eventpipe == nil {
    time.Msleep (10)
  }
//  unlock()
//  println ("keyboard.catchX: Eventpipe != nil")
  for p := range scr.Eventpipe {
    xpipe <- p
    lock()
  }
  close (xpipe)
}

func isSet (bit, x uint) bool {
  return x >> bit % 2 == 1
}

func inputX (B *byte, C *Comm, D *uint) {
  const (
    shiftBit     =  0
    shiftLockBit =  1
    ctrlBit      =  2
    altBit       =  3
    altGrBit     =  7
    mouseBitL    =  8
    mouseBitM    =  9
    mouseBitR    = 10
  )
  var event scr.Event
  ok := false
loop:
  for {
    *B, *C, *D = 0, None, 0
    event, ok = <-xpipe
    if ! ok { panic ("kbd/x.go inputX: ! ok") }
    unlock()
    if event.S == 64 { continue } // d(o,o)f-key
    shift := isSet (shiftBit, event.S)
    shiftLock := isSet (shiftLockBit, event.S)
    if shiftLock { shift = false } // weg isser
    ctrl := isSet (ctrlBit, event.S)
    alt := isSet (altBit, event.S)
    altGr := isSet (altGrBit, event.S)
    mouseL := isSet (mouseBitL, event.S)
    mouseM := isSet (mouseBitM, event.S)
    mouseR := isSet (mouseBitR, event.S)
    if shift || ctrl {
      *D = 1
    } else if alt {
      *D = 2
    }
    switch event.T {
/*/
    case C.Expose:
      *B = 0
      *C = Expose
      *D = 0
      break loop
/*/
    case C.KeyPress:
      if event.C <= 8 {
        println ("oops, kbd/x.go C.Keypress keycode ", event.C, " <= 8")
      } else {
        event.C -= 8
        switch {
        case event.C == esc:
          *C = Esc
        case event.C == shiftL || event.C == shiftR:
          shift = true
        case event.C == ctrlL || event.C ==  ctrlR:
          ctrl = true
        case event.C == altL:
          alt = true
        case event.C == altR:
          altGr = true
        case isAlpha (event.C):
          if ctrl && (event.C == 'C' || event.C == 'Q' ) {
            // finX () // TODO
            os.Exit (0)
          }
          switch *D {
          case 0:
            if altGr {
              switch event.C {
              case 3: // 2
                *B = twoSup
              case 4: // 3
                *B = threeSup
              case 8: // 7
                *B = braceL
              case 9: // 8
                *B = bracketL
              case 10: // 9
                *B = bracketR
              case 11: // 0
                *B = braceR
              case 12: // ß
                *B = backslash
              case 16: // Q
                *B = '@'
              case 18: // E
                *B = euro
              case 19: // R
                *B = registered
              case 27: // +
                *B = '~'
              case 41:
                *B = degree
              case 46: // C
                *B = copyright
              case 50: // M
                *B = mu
              case 52: // .
                *B = division
              }
            } else {
              *B = bb[event.C]
            }
          case 1:
            if altGr {
/*/
              switch event.C {
              case 26:
                *B = Ü
              case 39:
                *B = Ö
              case 40:
                *B = Ä
              case 86:
                *B = '|'
              }
/*/
            } else {
              *B = aa[event.C]
            }
          case 2:
            *B = aa[event.C]
          }
        case isCmd (event.C):
          *C = kK[event.C]
          switch event.C {
          case left:
          if alt {
            *C = PgLeft
            *D = 0
            if shift {
              *D = 1
            }
            break loop
          }
          case right:
            if alt {
              alt = false
              *C = PgRight
              *D = 0; if shift {
                *D = 1
              }
              break loop
            }
          case pgUp:
            *C = PgUp
//            *D = 2
            break loop
          case pgDown:
            *C = PgDown
//            *D = 2
            break loop
          }
//          if (e.C == left || e.C == right) && e.S == 64 { *D += 2 }
          if event.C == back && *D > 2 {
            *C, *D = None, 0
          } // doesn't help: wm crashes
        case event.C == numOnOff:
//        // TODO
        case isKeypad (event.C):
          switch *D {
          case 0:
            *B = bb[event.C]
          default:
            *C = kK[event.C]
          }
        case event.C == 127:
          *B = backslash
        default:
          println ("kbd/x.go C.KeyPress: keycode ", event.C, "/ state ", event.S) // XXX
        }
      }
      if *B > 0 || *C > 0 {
        break loop
      }
    case C.KeyRelease:
      ; // is ignored
    case C.ButtonPress:
      if *D > 1 {
        *D = 1 // because the WM eats up everything else
      }
      switch event.C {
      case 1:
        *C = Here
      case 2:
        *C = That
      case 3:
        *C = This
      case 4:
        *C = ScrollUp
      case 5:
        *C = ScrollDown
      default:
        println ("kbd/x.go C.ButtonPress: button ", event.C ,"/ state ", event.S) // XXX
      }
      if *C > 0 {
        break loop
      }
    case C.ButtonRelease:
      if *D > 1 {
        *D = 1 // because WM eats everything else up
      }
      ctrl = false
      alt = false
      altGr = false
      switch event.C {
      case 1:
        if mouseL {
//          mouseL = false
          *C = To
        }
      case 2:
        if mouseM {
//          mouseM = false
          *C = Thither
        }
      case 3:
        if mouseR {
//          mouseR = false
          *C = There
        }
      case 4:
        *C = ScrollUp
      case 5:
        *C = ScrollDown
      default:
        println ("kbd/x.go C.ButtonRelease: button ", event.C ,"/ state ", event.S) // XXX
      }
      if *C > 0 {
        break loop
      }
    case C.MotionNotify:
//      *D = 0
      if mouseL {
        *C = Drag
      } else if mouseM {
        *C = Move
      } else if mouseR {
        *C = Drop
      } else {
        *C = Go
      }
      break loop
    case C.ClientMessage:
      ; // break loop // navi
    default:
      *B, *C, *D = 0, None, 0
      break loop
    }
  }
  lastbyte, lastcommand, lastdepth = *B, *C, *D
}
