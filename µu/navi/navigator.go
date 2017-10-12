package navi

// (c) Christian Maurer   v. 121204 - license see µu.go

// Pre: /dev/input/navi is readable for world.
// If there is e.g. a rule in /etc/udev/rules.d with the line:
// SYSFS{idVendor}=="046d", SYSFS{idProduct}=="c626", MODE="444", SYMLINK+="input/navi"
// then a Space Navigator of 3dconnexion is initialized.

import (
  "os"
  . "µu/obj"
  . "µu/spc"
)
type
  button byte; const (
  none = iota
  right
  left
)
const (
  buttonReleased = 0
  buttonPressed = 1
)
const ( // see mouse: Go, ...
  dummy0 = iota
  here; dummy2; drop // left button
  there; dummy5; drop1 // right button
  dummy7; dummy8; dummy9
  navigate
)
var (
  file *os.File
  navipipe chan Command
  but button
  data [2 * NDirs]int16
  index [NDirs]int16 // 0..2
  sign [NDirs]int16 // +1, -1
)

func init() {
  for d := D0; d < NDirs; d++ {
    sign[d], index[d] = 1, int16(d)
  }
// The 3dconnexion SpaceNavigator has the rightoriented trihedral (x, y, z) = (right, back, bottomn):
// it delivers the movements in 0..2 and the rotations around the corresponding axes (consequently - viewed
// in opposition to the direction of the axes - in mathematical positive sense) in NDirs+0..2 = 3..5.
// In µu this is translated into the trihedral (x, y, z) = (right, front, top) defined in the package spc:
  sign[Front], sign[Top] = -1, -1
  f, err := os.OpenFile ("/dev/input/navi", os.O_RDONLY, 0444)
  if err == nil {
    file = f
    navipipe = make (chan Command)
    go catch()
  } else {
    file = nil
    navipipe = (chan Command)(nil)
  }
//  "Move Right"
//  "Move Front"
//  "Move   Top"
//  "Rot Right"
//  "Rot Front"
//  "Rot   Top"
}

func exists() bool {
  return file != nil
}

/*
func channel() (chan Command) {
  return navipipe
}
*/
/*
func isSet (bit, x uint) bool {
  return x >> bit % 2 == 1
}
*/

func catch() {
  var (
    value int16
    typ, code byte
  )
  const
    M = 8
  var (
    B = make ([]byte, M)
    C Command
  )
  for {
    i, _ := file.Read (B[:])
    if i != M { continue }
    value = Decode (int16(0), B[4:6]).(int16)
    typ = B[6]
    code = B[7]
    C = dummy0
    switch typ { case 0:
      println ("navi.catch(): What the hell is going on HERE ?")
    case 1: // key
      switch code { case 0:
        but = left // 2
      case 1:
        but = right // 1
      default:
        but = none // 0
        println ("unknown navi-input_event code", code)
        continue
      }
      if value == buttonReleased {
        if but == left {
          C = drop
        } else { // but == right 
          C = drop1
        }
      } else { // value == buttonPressed
        if but == left {
          C = here
        } else { // but == right
          C = there
        }
      }
    case 2: // movement, rotation
      C = navigate
      if Direction(code) >= 2 * NDirs {
        println ("navi-input_event code too big: ", code)
        continue
      }
      if -1 <= value && value <= 1 { // suppress small movements
        data [code] = 0
      } else {
        data [code] = value
      }
    case 129, 130:
      C = navigate
      continue
    default:
      println ("unknown navi-input_event type ", typ)
      continue
    }
    navipipe <- C
  }
}

func read() (GridCoord, GridCoord) {
  var mov, rot GridCoord
  for d := D0; d < NDirs; d++ {
    mov[d] = sign[d] * data[index [d]]
    rot[d] = sign[d] * data[int16(NDirs) + index [d]]
  }
  return mov, rot
}
