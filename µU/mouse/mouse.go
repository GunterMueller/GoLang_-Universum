package mouse

// (c) Christian Maurer   v. 211226 - license see µU.go

import (
  "os"
  "µU/ker"
)
type
  buttons byte; const (
  none = iota
  left
  right
  middle
)
var (
  file *os.File
  lastCommand = None
  button = none
  lastButton = none
  yy, // vertical swap
  x0, y0, x1, y1, // boundaries
  xm, ym uint // position of mouse pointer
)

func init() {
  dev := "/dev/input/mice"
  var e error
  file, e = os.Open (dev)
  if e != nil { ker.Panic (dev + " is not readable !") }
  Def (0, 0, 1600, 1200)
  Pipe = make (chan Command)
  go catch()
}

func def (x, y, w, h uint) {
  x0, y0 = x, y
  x1, y1 = x0 + w - 1, y0 + h - 1
  yy = y1
  xm, ym = x0 + w / 2, y0 + h / 2
}

func warp (x, y uint) {
  if x > x1 {
    x = x1
  }
  if y > y1 {
    y = y1
  }
  xm, ym = x, yy - y
}

func catch() {
  var (
    bs [3]byte
    a, dx, dy uint
    moved bool
    cmd Command
  )
  for {
    i, _ := file.Read (bs[:])
    if i < 3 { continue }
    a = uint(bs[0])
    switch a % 8 {
    case 0:
      button = none
    case 1, 5: // left, left and middle
      button = left
    case 2, 6: // right, right and middle
      button = right
    default:   // left and right, middle, all three
      button = middle
    }
    dx, dy = uint(bs[1]), uint(bs[2])
    moved = dx > 0 || dy > 0
    if button == lastButton && ! moved { continue }
    a /= 8
    if a == 0 { break }
    a = (a - 1) / 2
    switch a {
    case 0:
      xm += dx
      ym += dy
    case 1:
      dx = 256 - dx
      ym += dy
      if xm > dx { xm -= dx } else { xm = 0 }
    case 2:
      dy = 256 - dy
      xm += dx
      if ym > dy { ym -= dy } else { ym = 0 }
    case 3:
      dx, dy = 256 - dx, 256 - dy
      if xm > dx { xm -= dx } else { xm = 0 }
      if ym > dy { ym -= dy } else { ym = 0 }
    default:
      break
    }
    if xm < x0 {
      xm = x0
    } else if xm > x1 {
      xm = x1
    }
    if ym < y0 {
      ym = y0
    } else if ym > y1 {
      ym = y1
    }
    switch button {
    case none:
      switch lastCommand {
      case Go:
        if moved {
        cmd = Go
      } else {
        continue
      }
    case Here, Drag:
      cmd = To
    case This, Drop:
      cmd = There
    case That, Move:
      cmd = Hither
    case To, There, Hither:
      cmd = Go
    }
    case left:
      switch lastCommand {
      case Go, To, There, Hither:
        cmd = Here
      case Here, Drag:
        if moved {
          cmd = Drag
        } else {
          continue
        }
      default:
        cmd = lastCommand
      }
    case right:
      switch lastCommand {
      case Go, To, There, Hither:
        cmd = This
      case This, Drop:
        if moved {
          cmd = Drop
        } else {
          continue
        }
      default:
        cmd = lastCommand
      }
    case middle:
      switch lastCommand {
      case Go, To, There, Hither:
      cmd = That
      case That, Move:
        if moved {
          cmd = Move
        } else {
          continue
        }
      default:
        cmd = lastCommand
      }
    }
    lastButton = button
    lastCommand = cmd
    Pipe <- cmd
  }
}

func pos() (int, int) {
  return int(xm), int(yy - ym)
}
