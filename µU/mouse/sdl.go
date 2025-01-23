package mouse

// (c) Christian Maurer   v. 220206 - license see µU.go

/*/
import
  "os"
/*/

func catch1() {
/*/
  var (
    bs [3]byte
    a, dx, dy uint
    moved bool
    cmd Command
    lastCommand = Go
    button = noButton
    lastButton = noButton
  )
  for {
    i, _ := file.Read (bs[:])
    if i < 3 { continue }
    a = uint(bs[0])
    switch a % 8 {
    case 0:
      button = noButton
    case 1, 5: // leftButton, left and middleButton
      button = leftButton
    case 2, 6: // rightButton, right and middleButton
      button = rightButton
    default:   // leftButton and rightButton, middleButton, all three
      button = middleButton
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
    case noButton:
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
        cmd = Thither
      case To, There, Thither:
        cmd = Go
      }
    case leftButton:
      switch lastCommand {
      case Go, To, There, Thither:
        cmd = Here
      case Here, Drag:
        if moved {
          cmd = Drag
        } else {
          continue
        }
//      case None, This, That, Drop, Move:
      case This, That, Drop, Move:
        cmd = lastCommand
      }
    case rightButton:
      switch lastCommand {
      case Go, To, There, Thither:
        cmd = This
      case This, Drop:
        if moved {
          cmd = Drop
        } else {
          continue
        }
 //     case None, Here, That, Drag, Move:
      case Here, That, Drag, Move:
        cmd = lastCommand
      }
    case middleButton:
      switch lastCommand {
      case Go, To, There, Thither:
      cmd = That
      case That, Move:
        if moved {
          cmd = Move
        } else {
          continue
        }
//        case None, Here, This, Drag, Drop:
        case Here, This, Drag, Drop:
        cmd = lastCommand
      }
    }
    lastButton = button
    lastCommand = cmd
    if cmd != None {
      Pipe <- cmd
    }
  }
/*/
}

func init() {
/*/
  gfx.Fenster (1600, 1200)
/*/
/*/
  dev := "/dev/input/mice"
  var e error
  file, e = os.Open (dev)
  if e != nil { panic (dev + " is not readable !") }
  Def (0, 0, 1600, 1200)
  Pipe = make (chan Command)
/*/
  go catch1()
}
