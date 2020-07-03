package main

// >>> to be applied in small windows on a GUI;
//     the far monitor is started with the argument 0
//     and the clients with different arguments 1, 2, 3, ...


import ("time"; "math/rand"; "nU/ego"; "nU/env"; "nU/term"; "nU/col"; "nU/scr"; . "nU/lr")

var (
  x LeftRight
  sL, sR = "", ""
  s0 = "                          "
  uL = "  start left proces: Tab  "
  uR = "start right process: Enter"
  uE = "               stop: Esc  "
  uM = "terminate far monitor: Esc"
)

func drive() {
  const t = 5e9; time.Sleep (t + time.Duration(rand.Intn(t)))
}

func write() {
  for {
    scr.Lock()
    scr.ColourF (col.LightRed())
    scr.Write (s0, 0, 0)
    scr.Write (sL, 0, 0)
    scr.ColourF (col.LightGreen())
    scr.Write (s0, 1, 0)
    scr.Write (sR, 1, 0)
    scr.Unlock()
    time.Sleep (1e8)
  }
}

func plusL() {
  scr.Lock()
  sL += "<"
  scr.Unlock()
}

func minusL() {
  scr.Lock()
  sL = sL[1:]
  scr.Unlock()
}

func plusR() {
  scr.Lock()
  sR += ">"
  scr.Unlock()
}

func minusR() {
  scr.Lock()
  sR = sR[1:]
  scr.Unlock()
}

func left() {
  x.LeftIn()
  plusL()
  drive()
  minusL()
  x.LeftOut()
}

func right() {
  x.RightIn()
  plusR()
  drive()
  minusR()
  x.RightOut()
}

func main() {
  scr.New(); defer scr.Fin()
  term.New(); defer term.Fin()
  me := ego.Me()
  x = NewFarMonitor (env.Localhost(), 50000, me == 0)
  if me == 0 {
    scr.Write (uM, 5, 0)
    for {
      switch term.Read() {
      case term.Esc:
        return
      }
    }
  } else {
    scr.Write (uL, 3, 0)
    scr.Write (uR, 4, 0)
    scr.Write (uE, 5, 0)
    go write()
    for {
      switch term.Read() {
      case term.Esc:
        return
      case term.Back, term.Tab:
        go left()
      case term.Enter:
        go right()
      }
    }
  }
}
