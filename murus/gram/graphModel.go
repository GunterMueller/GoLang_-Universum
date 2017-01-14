package gram

// (c) murus.org  v. 170107 - license see murus.go

import (
  . "murus/obj"; "murus/str"; "murus/kbd"
  "murus/col"; "murus/scr"; "murus/errh"; "murus/img"
  "murus/gra"; "murus/node"; "murus/edge"
)
type
  graphModel struct {
                    gra.Graph
                  n node.Node
                  e edge.Edge
              nWays uint
                    string "name"
                    bool "has background"
               f, b,
             fa, ba col.Colour
                    }
var (
  h = [...]string { "        neue Ecke: linke Maustaste              ",
                    " Ecke verschieben: rechte Maustaste             ",
                    "  Ecken verbinden: mittlere Maustaste           ",
                    "                                                ",
                    "Eckennamen ändern: Vorwahl- und linke Maustaste ",
                    "     Ecke löschen: Vorwahl- und rechte Maustaste",
                    " Graph ausdrucken: Drucktaste                   ",
                    "                                                ",
                    "(Vorwahltaste = Umschalt-, Strg- oder Alt-Taste)",
                    "                                                ",
                    "             Done: Abbruchtaste (Esc)           " }
  help = make([]string, len(h))
)

func init() {
  for i, l := range (h) { help[i] = str.Lat1(l) }
}

func New (d bool, n, v Any) GraphModel {
  x := new(graphModel)
// TODO  checkType node
  x.n = Clone(n).(node.Node)
  if v == nil {
    x.e = nil
  } else {
    x.e = edge.New (v)
  }
  x.Graph = gra.New (d, x.n, x.e)
  x.f, x.b = col.StartCols()
  x.Colours (x.f, x.b)
  x.fa, x.ba = col.Red, x.b
  x.ColoursA (x.fa, x.ba)
  return x
}

func (x *graphModel) Background (n string) {
  x.bool = ! str.Empty (n)
  x.string = n
  img.Get (x.string, 0, 0)
  scr.Save1()
}

func (x *graphModel) Colours (f, b col.Colour) {
  x.f, x.b = f, b
}

func (x *graphModel) ColoursA (f, b col.Colour) {
  x.fa, x.ba = f, b
}

func (x *graphModel) underMouse (a Any) bool {
  return a.(node.Node).UnderMouse()
}

func (x *graphModel) selected (n node.Node) bool {
  loop: for {
    c, _ := kbd.Command()
    scr.MousePointer (true)
    switch c {
    case kbd.Esc:
      break loop
    case kbd.Enter:
      n = x.Get().(node.Node)
      return true
    case kbd.Here:
      if x.Graph.ExPred (x.underMouse) {
        n = x.Get().(node.Node)
        return true
      }
    }
  }
  return false
}

func (x *graphModel) NodesSelected() bool {
  scr.MousePointer (true)
  s := false
  for {
    errh.Hint ("Start auswählen")
    if x.selected (x.n.(node.Node)) { // n local
      x.Locate (true) // n colocal
      x.Write()
      errh.Hint ("Ziel auswählen")
      if x.selected (x.n.(node.Node)) { // n local
        errh.DelHint()
        if x.Located() {
          errh.Error0("Fehler: Start und Ziel sind gleich !")
        } else {
          s = true
          break
        }
      }
    } else {
      break
    }
  }
  x.Write()
  errh.DelHint()
  scr.MousePointer (false)
  return s
}

func (x *graphModel) NodeSelected() bool {
  scr.MousePointer (true)
  errh.Hint ("Ecke auswählen")
  s := false
  if x.selected (x.n.(node.Node)) { // n local
    x.Locate (true) // n colocal
    s = true
    x.Write()
  }
  errh.DelHint()
  scr.MousePointer (false)
  return s
}

func (x *graphModel) write1 (n Any, a bool) {
  n.(node.Node).Colours (x.f, x.b)
  n.(node.Node).ColoursA (x.fa, x.ba)
  n.(node.Node).Write1 (a)
}

func (g *graphModel) write3 (n, e, n1 Any, a bool) {
  w, h := n.(node.Node).Size()
  if w + h == 0 { return }
  x, y := n.(node.Node).Pos()
  x1, y1 := n1.(node.Node).Pos()
  f, b := g.f, g.b; if a { f, b = g.fa, g.ba }
  scr.ColourF (f)
  scr.Line (x, y, x1, y1)
  if e == nil { return }
  e.(edge.Edge).Colours (f, b)
  e.(edge.Edge).Write (x, y, x1, y1, a)
  if g.Graph.Directed() {
    x0, y0 := (x + 4 * x1) / 5, (y + 4 * y1) / 5
    scr.CircleFull (x0, y0, 4)
  }
}

func (x *graphModel) Write() {
  scr.Buf (true)
  if x.bool { scr.Restore1() }
  x.Trav3Cond (x.write1, x.write3)
  scr.Buf (false)
}

func (x *graphModel) Edit() {
  x.Write()
  errh.Hint ("Graph editieren - Hilfe: F1, fertig: Esc")
  loop: for {
    c, i := kbd.Command()
    scr.MousePointer (true)
    errh.DelHint()
    switch c {
    case kbd.Esc:
      break loop
    case kbd.Help:
      errh.Help (help)
    case kbd.Here: // new node or change name of existing node:
      if x.Graph.ExPred (x.underMouse) {
        if i > 0 {
          x.n = x.Get().(node.Node) // local: node
          x.n.Edit()
          x.Put (x.n)
          x.Write()
        }
      } else {
        x.n.Clr()
        x.n.Mouse()
        x.Ins (x.n)
        x.Write()
        x.n.Edit()
        x.Put (x.n)
        x.Write()
      }
    case kbd.Del: // remove node
      if x.Graph.ExPred (x.underMouse) {
        x.Del()
      }
      x.Write()
    case kbd.There: // move node
      switch i {
      case 0:
        if x.Graph.ExPred (x.underMouse) { // n local
          loop1: for {
            kk, _ := kbd.Command()
            scr.MousePointer (true)
            switch kk {
            case kbd.Push:
              x.n = x.Get().(node.Node)
              x.n.Mouse()
              x.Put (x.n)
              x.Write()
            case kbd.Thither:
              break loop1
            }
          }
        }
      default: // remove node
        if x.Graph.ExPred (x.underMouse) {
          x.Del()
          x.Write()
        }
      }
    case kbd.This: // connect nodes / remove edge:
      xm, ym := scr.MousePosGr()
      xm1, ym1 := xm, ym
      if x.Graph.ExPred (x.underMouse) {
        x.n = x.Get().(node.Node) // n local
        x.Locate (true) // n also colocal
        loop2: for {
          kk, _ := kbd.Command()
          scr.MousePointer (true)
          switch kk {
          case kbd.Move:
            scr.LineInv (xm, ym, xm1, ym1)
            xm1, ym1 = scr.MousePosGr()
            scr.LineInv (xm, ym, xm1, ym1)
          case kbd.Thus:
            scr.LineInv (xm, ym, xm1, ym1)
            if x.Graph.ExPred (x.underMouse) {
              n1 := x.Get().(node.Node) // n1 local, n colocal
              if x.Located () {
                errh.Error0("Schleife - geht nicht")
              } else {
                if x.e != nil {
                  x0, y0 := x.n.Pos()
                  x1, y1 := n1.Pos()
                  x.e.Edit (x0, y0, x1, y1)
                }
                x.Edge1 (x.e)
/*
                if x.e.Val() == 0 { // XXX
                  x.Del1()
                }
*/
                x.Write()
              }
              break loop2
            }
          }
        }
      }
    case kbd.Print:
      errh.DelHint()
      img.Put1 (".tmp.g")
      img.Print1()
      errh.Hint ("Graph editieren - Hilfe: F1, fertig: Esc")
    }
  }
  errh.DelHint()
}

func (x *graphModel) ResetNWays() {
  x.nWays = 0
}

func (x *graphModel) NWays() uint {
  return x.nWays
}

func wait (w *bool) {
  *w = *w && kbd.Wait (true)
}

func (x *graphModel) DFS (all bool) {
  x.Mark (true)
  x.Write()
//  kbd.Wait (true)
  n := x.NumLoc()
  if n > 0 {
    for i := uint(0); i < n; i++ {
      x.Step (i, true)
      if x.Marked() {
        x.Step (0, false)
      } else {
        x.DFS (all)
        if all { x.Mark (false) } // for _all_ ways
        x.Step (0, false)
        x.Write()
//        kbd.Wait (true)
      }
    }
  }
}

func (x *graphModel) BFS (all bool) {
  x.Graph.SetDemo (gra.Breadth) // >>> ist eine DijkstraDemo, wenn x bewertet ist
  if x.Conn() {
    x.Act()
    x.Write()
    errh.Error ("ein kürzester Weg in Hinrichtung der Länge", x.Num1Act())
  } else {
    x.Write()
    errh.Error0("es gibt keinen Weg in Hinrichtung")
  }
  if x.Directed() {
    x.Relocate()
    x.Write()
    if x.Conn() {
      x.Act()
      x.Write()
      errh.Error ("ein (kürzester) Weg in Rückrichtung der Länge", x.Num1Act())
    } else {
      errh.Error0("es gibt keinen Weg in Rückrichtung")
    }
  }
}

func (x *graphModel) Hamilton (ready, ok Cond, w *bool) {
  x.Mark (true)
  if ready() {
    x.nWays ++
    x.Write()
    errh.Error ("Hamiltonweg", x.nWays)
  } else {
    n := x.NumLoc()
    for i := uint(0); i < n; i++ {
      x.Step (i, true)
      if x.Marked() || ! ok() {
        x.Step (0, false)
      } else {
        x.Write(); wait (w)
        x.Hamilton (ready, ok, w)
        x.Mark (false)
        x.Step (0, false)
        x.Write(); wait (w)
      }
    }
  }
}

func (x *graphModel) Demo (d gra.Demo) {
  x.Graph.SetDemo (d)
  x.Install (x.write1, x.write3)
}
