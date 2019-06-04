package kbd

// (c) Christian Maurer   v. 171217 - license see µU.go

// #cgo LDFLAGS: -lX11
// #include <X11/X.h>
import
  "C"
import (
  "os"
  "µU/time"
  "µU/xwin"
)
var (
  underX bool
  xpipe chan xwin.Event
  ch chan int = make (chan int, 1)
)

// Pre: xwin.x.initialized == true
func catchX () {
  for xwin.Eventpipe == nil {
    time.Msleep (10)
  }
//  ch <- 0
//  println ("keyboard.catchX: Eventpipe != nil")
  for p := range xwin.Eventpipe {
    xpipe <- p
    <-ch
  }
  close (xpipe)
}

func isSet (bit, x uint) bool {
  return x >> bit % 2 == 1
}

func inputX (B *byte, C *Comm, T *uint) {
  const (
    shiftBit     = 0
    shiftLockBit = 1
    ctrlBit      = 2
    altBit       = 3
    altGrBit     = 7
    mouseBitL    = 8
    mouseBitM    = 9
    mouseBitR   = 10
  )
  var e xwin.Event
//  var k uint
  ok := false
  loop: for {
    *B, *C, *T = 0, None, 0
    e, ok = <-xpipe
    ch <- 0
    if ! ok { println ("x.inputX: ! ok") }
if e.S == 64 { panic("x.inputX: shriek !") }
    shiftX := isSet (shiftBit, e.S)
    shiftFix := isSet (shiftLockBit, e.S)
    if shiftFix { shiftX = false } // weg isser
    ctrlX := isSet (ctrlBit, e.S)
    altX := isSet (altBit, e.S)
    altGrX := isSet (altGrBit, e.S)
    lBut := isSet (mouseBitL, e.S)
    mBut := isSet (mouseBitM, e.S)
    rBut := isSet (mouseBitR, e.S)
    if shiftX || ctrlX { *T ++ }
    if altX { *T += 2 }
    switch e.T {
    case C.KeyPress:
      if e.C < 9 {
        println ("oops, got keycode ", e.C, " < 9") // XXX ?
      } else {
        e.C -= 8
        switch {
        case e.C == esc:
          *C = Esc
        case e.C == shiftL || e.C == shiftR:
          shiftX = true
        case e.C == ctrlL || e.C ==  ctrlR:
          ctrlX = true
        case e.C == altL:
          altX = true
        case e.C == altR:
          altGrX = true
        case isAlpha (e.C):
          if ctrlX && (e.C == 46 || e.C == 16 ) { // Ctrl C, Ctrl Q
            // finX ()
            os.Exit (0)
          }
          switch *T {
          case 0:
            if altGrX {
              switch e.C {
              case 16, 19, 33, 46, 50: // Q, R, F, C, M ->
                *B = aA[e.C] // @, z.Registered, z.Female, z.Copyright, z.Male
              }
            } else {
              *B = bb[e.C]
            }
          case 1:
            *B = bB[e.C]
          case 2:
            *B = aa[e.C]
          }
        case isCmd (e.C):
          *C = kK[e.C]
/*
          if e.C == pgUp || e.C == pgDown {
//          if e.C == pgUp + 8 || e.C == pgDown + 8 { // 112/117 -> 104/109
            *T += 2
          }
*/
          if e.C == left && altX {
            *C = PgLeft
            *T = 0; if shiftX { *T = 1 }
            break loop
          }
          if e.C == right && altX {
            altX = false
            *C = PgRight
            *T = 0; if shiftX { *T = 1 }
            break loop
          }
          if e.C == pgUp {
            *C = PgUp
            break loop
          }
          if e.C == pgDown {
            *C = PgDown
            break loop
          }
//          if (e.C == left || e.C == right) && e.S == 64 { *T += 2 }
          if e.C == back && *T > 2 { *C = None; *T = 0 } // doesn't help: wm crashes
//        case k == numOnOff:
//          ; // TODO
        case isKeypad (e.C):
          switch *T { case 0:
            *B = bb[e.C]
          default:
            *C = kK[e.C]
          }
//        case isFunction (e.C):
//          ; // TODO
        default:
          println ("C.KeyPress: keycode ", e.C, "/ state ", e.S)
        }
      }
      if *B > 0 || *C > 0 {
        break loop
      }
    case C.KeyRelease:
      ;
    case C.ButtonPress:
      if *T > 1 { *T = 1 } // because the bloody WM eats everything else up
      switch e.C { case 1:
        *C = Here
      case 2:
        *C = This
      case 3:
        *C = There
      case 4:
        *C = Up
      case 5:
        *C = Down
      default:
        println ("xwin.ButtonPress: button ", e.C ,"/ state ", e.S)
      }
      if *C > 0 {
        break loop
      }
    case C.ButtonRelease:
      if *T > 1 { *T = 1 } // because the bloody WM eats everything else up
      ctrlX = false
      altX = false
      altGrX = false
      switch e.C { case 1:
        if lBut {
//          lBut = false
          *C = Hither
        }
      case 2:
        if mBut {
//          mBut = false
          *C = Thus
        }
      case 3:
        if rBut {
//          rBut = false
          *C = Thither
        }
      case 4:
        *C = Up
      case 5:
        *C = Down
      default:
        println ("xwin.ButtonRelease: button ", e.C ,"/ state ", e.S)
      }
      if *C > 0 {
        break loop
      }
    case C.MotionNotify:
      *T = 0
      if lBut {
        *C = Pull
      } else if mBut {
        *C = Move
      } else if rBut {
        *C = Push
      } else {
        *C = Go
      }
      break loop
    case C.ClientMessage:
      ; // break loop // navi
    default:
      *B, *C, *T = 0, None, 0
      break loop
    }
  }
  lastbyte, lastcommand, lastdepth = *B, *C, *T
}
