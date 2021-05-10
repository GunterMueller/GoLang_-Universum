package mouse

// (c) Christian Maurer   v. 210505 - license see µU.go

import (
  "os"
  "µU/ker"
)
const (
  go_ = iota         // mousemove without any button pressed
  here; this; that   // left mousebutton
  drag; drop; move   // right mousebutton
  to_; there; hither // middle mousebutton
)
type
  button byte; const (
  none = iota
  left
  right
  middle
)
var (
  file *os.File
  mousepipe chan Command
  lastCommand Command
  butt, oldButt button
  yy, // vertical swap
  x0, y0, x1, y1, // boundaries
  xm, ym uint // location of mouse pointer
)

func init() {
  mousedev := "/dev" + Mouse
  var e error
  if file, e = os.Open (mousedev); e == nil {
    Def (0, 0, 1600, 1200) // TODO
    lastCommand = go_
    oldButt = none
    mousepipe = make (chan Command)
    go catch()
  } else {
    ker.Panic (mousedev + " nicht lesbar !")
  }
}

func ex() bool {
  return mousepipe != (chan Command)(nil)
}

func channel() chan Command {
  return mousepipe
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
    dragged bool
    c Command
  )
  for {
    i, _:= file.Read (bs[:])
    if i < 3 { continue }
    a = uint(bs[0])
    switch a % 8 {
    case 0:
      butt = none
    case 1, 5: // left, left and middle
      butt = left
    case 2, 6: // right, right and middle
      butt = right
    default:   // left and right, middle, all three
      butt = middle
    }
    dx = uint(bs[1])
    dy = uint(bs[2])
    dragged = dx > 0 || dy > 0
    if butt == oldButt && ! dragged { continue }
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
      dx = 256 - dx
      dy = 256 - dy
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
    switch butt {
    case none:
      switch lastCommand {
      case go_:
        if dragged {
        c = go_
      } else {
        continue
      }
    case here, this:
      c = that
    case drag, drop:
      c = move
    case to_, there:
      c = hither
    case that, move, hither:
      c = go_
    }
    case left:
      switch lastCommand {
      case go_, that, move, hither:
        c = here
      case here, this:
        if dragged {
          c = this
        } else {
          continue
        }
      default:
        c = lastCommand
      }
    case right:
      switch lastCommand {
      case go_, that, move, hither:
        c = drag
      case drag, drop:
        if dragged {
          c = drop
        } else {
          continue
        }
      default:
        c = lastCommand
      }
    case middle:
      switch lastCommand {
      case go_, that, move, hither:
      c = to_
      case to_, there:
        if dragged {
          c = there
        } else {
          continue
        }
      default:
        c = lastCommand
      }
    }
    oldButt = butt
    lastCommand = c
    mousepipe <- c
  }
}

func pos() (int, int) {
  return int(xm), int(yy - ym)
}
