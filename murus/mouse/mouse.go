package mouse

// (c) murus.org  170814 - license see murus.go

import
(
  "os"
  "murus/ker"
  "murus/xwin"
)
const (
  move = iota         // mousemove without any button pressed
  here; drag; drop    // left mousebutton
  there; drag1; drop1 // right mousebutton
  this; drag2; drop2  // middle mousebutton
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
  mX, mY uint // location of mouse pointer
)

func init() {
  if xwin.UnderX() {
    mousepipe = (chan Command)(nil)
  } else {
    var err error
    file, err = os.Open ("/dev/input/mice")
    if err == nil {
      Def (0, 0, 1000, 1000) // TODO
      lastCommand = move
      oldButt = none
      mousepipe = make (chan Command)
      go catch()
    } else {
      ker.Panic ("/dev/input/mice nicht lesbar !")
    }
  }
}

func Ex() bool {
//  return file != nil
  return mousepipe != (chan Command)(nil)
}

func Channel() chan Command {
  return mousepipe
}

func Def (x, y, w, h uint) {
  x0, y0 = x, y
  x1, y1 = x0 + w - 1, y0 + h - 1
  yy = y1
  mX, mY = x0 + w / 2, y0 + h / 2
}

func Warp (x, y uint) {
  if x > x1 {
    x = x1
  }
  if y > y1 {
    y = y1
  }
  mX, mY = x, yy - y
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
    switch a % 8 { case 0:
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
    switch a { case 0:
      mX += dx
      mY += dy
    case 1:
      dx = 256 - dx
      mY += dy
      if mX > dx { mX -= dx } else { mX = 0 }
    case 2:
      dy = 256 - dy
      mX += dx
      if mY > dy { mY -= dy } else { mY = 0 }
    case 3:
      dx = 256 - dx
      dy = 256 - dy
      if mX > dx { mX -= dx } else { mX = 0 }
      if mY > dy { mY -= dy } else { mY = 0 }
    default:
      break
    }
    if mX < x0 {
      mX = x0
    } else if mX > x1 {
      mX = x1
    }
    if mY < y0 {
      mY = y0
    } else if mY > y1 {
      mY = y1
    }
    switch butt { case none:
      switch lastCommand { case move:
        if dragged {
        c = move
      } else {
        continue
      }
    case here, drag:
      c = drop
    case there, drag1:
      c = drop1
    case this, drag2:
      c = drop2
    case drop, drop1, drop2:
      c = move
    }
    case left:
      switch lastCommand { case move, drop, drop1, drop2:
        c = here
      case here, drag:
        if dragged {
          c = drag
        } else {
          continue
        }
      default:
        c = lastCommand
      }
    case right:
      switch lastCommand { case move, drop, drop1, drop2:
        c = there
      case there, drag1:
        if dragged {
          c = drag1
        } else {
          continue
        }
      default:
        c = lastCommand
      }
    case middle:
      switch lastCommand { case move, drop, drop1, drop2:
      c = this
      case this, drag2:
        if dragged {
          c = drag2
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

func Pos() (int, int) {
  return int(mX), int(yy - mY)
}
