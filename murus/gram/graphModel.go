package gram

// (c) Christian Maurer   v. 170810 - license see murus.go

import (
  . "murus/obj"
  "murus/str"
  "murus/kbd"
  "murus/col"
  "murus/scr"
  "murus/errh"
  "murus/img"
  "murus/gra"
  "murus/vtx"
  "murus/edg"
)
type
  graphModel struct {
                    gra.Graph
             vertex vtx.Vertex
               edge edg.Edge
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

func new_(d bool, v, e Any) GraphModel {
  x := new(graphModel)
// TODO  checkType vertex
  x.vertex = Clone(v).(vtx.Vertex)
  if e == nil {
    e = edg.New(d, uint(1))
  }
  x.edge = edg.New (d, e)
  x.f, x.b = scr.StartCols()
  x.Colours (x.f, x.b)
  x.fa, x.ba = scr.StartColsA()
  x.ColoursA (x.fa, x.ba)
  x.Graph = gra.New (d, x.vertex, x.edge)
  x.Graph.SetWrite (vtx.W, edg.W)
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
  x.edge.Colours (f, b)
  x.vertex.Colours (f, b)
}

func (x *graphModel) ColoursA (f, b col.Colour) {
  x.fa, x.ba = f, b
  x.edge.ColoursA (f, b)
  x.vertex.ColoursA (f, b)
}

func (x *graphModel) underMouse (a Any) bool {
  return a.(vtx.Vertex).UnderMouse()
}

func (x *graphModel) selected (v vtx.Vertex) bool {
  loop: for {
    c, _ := kbd.Command()
    scr.MousePointer (true)
    switch c {
    case kbd.Esc:
      break loop
    case kbd.Enter:
      v = x.Get().(vtx.Vertex)
      return true
    case kbd.Here:
      if x.ExPred (x.underMouse) {
        v = x.Get().(vtx.Vertex)
        return true
      }
    }
  }
  return false
}

func (x *graphModel) VerticesSelected() bool {
  scr.MousePointer (true)
  s := false
  for {
    errh.Hint ("Start auswählen")
    if x.selected (x.vertex.(vtx.Vertex)) { // vertex local
      x.Locate (true) // vertex colocal
      x.Graph.Write()
      errh.Hint ("Ziel auswählen")
      if x.selected (x.vertex.(vtx.Vertex)) { // vertex local
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
  x.Graph.Write()
  errh.DelHint()
  scr.MousePointer (false)
  return s
}

func (x *graphModel) VertexSelected() bool {
  scr.MousePointer (true)
  errh.Hint ("Ecke auswählen")
  s := false
  if x.selected (x.vertex.(vtx.Vertex)) { // n local
    x.Locate (true) // n colocal
    s = true
    x.Graph.Write()
  }
  errh.DelHint()
  scr.MousePointer (false)
  return s
}

func (x *graphModel) write (v Any, a bool) {
  v.(vtx.Vertex).Colours (x.f, x.b)
  v.(vtx.Vertex).ColoursA (x.fa, x.ba)
  v.(vtx.Vertex).Write1 (a)
}

func (x *graphModel) write1 (e Any, a bool) {
  e.(edg.Edge).Colours (x.f, x.b)
  e.(edg.Edge).Write1 (a)
}

func (x *graphModel) Write() {
  scr.Buf (true)
  if x.bool { scr.Restore1() }
  x.Graph.Write()
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
    case kbd.Here: // new vertex or change name of existing vertex:
      if x.ExPred (x.underMouse) {
        if i > 0 {
          x.vertex = x.Get().(vtx.Vertex) // local: vertex
          x.vertex.Edit()
          x.Put (x.vertex)
          x.Write()
        }
      } else {
        x.vertex.Clr()
        x.vertex.Mouse()
        x.Ins (x.vertex)
        x.Write()
        x.vertex.Edit()
        x.Put (x.vertex)
        x.Write()
      }
    case kbd.Del: // remove vertex
      if x.ExPred (x.underMouse) {
        x.Del()
      }
      x.Write()
    case kbd.There: // move vertex
      switch i {
      case 0:
        if x.ExPred (x.underMouse) { // vertex local
          loop1: for {
            kk, _ := kbd.Command()
            scr.MousePointer (true)
            switch kk {
            case kbd.Push:
              x.vertex = x.Get().(vtx.Vertex)
              xu, yu := x.vertex.Pos()
              x.vertex.Mouse()
              x.Put (x.vertex)
              xv, yv := x.vertex.Pos()
              x.Graph.Trav1Loc (func (a Any) {
                                  x0, y0 := a.(edg.Edge).Pos0()
                                  if x0 == xu && y0 == yu {
                                    a.(edg.Edge).SetPos0 (xv, yv)
                                  }
                                  x1, y1 := a.(edg.Edge).Pos1()
                                  if x1 == xu && y1 == yu {
                                    a.(edg.Edge).SetPos1 (xv, yv)
                                  }
                                })
              x.Write()
            case kbd.Thither:
              break loop1
            }
          }
        }
      default: // remove vertex
        if x.ExPred (x.underMouse) {
          x.Del()
          x.Write()
        }
      }
    case kbd.This: // connect vertices / remove edge:
      if x.ExPred (x.underMouse) {
        x.vertex = x.Get().(vtx.Vertex) // vertex local
        xm, ym := x.vertex.Pos()
        xm1, ym1 := xm, ym
        x.Locate (true) // vertex also colocal
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
            if x.ExPred (x.underMouse) {
              if x.Located () {
                errh.Error0("Schleife - geht nicht")
              } else {
                xm1, ym1 := x.vertex.Pos()
                x.edge.SetPos0(xm, ym)
                x.edge.SetPos1(xm1, ym1)
                x.edge.Edit()
                x.Edge (x.edge)
/*
                if x.Edge.Val() == 0 { // XXX
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
  n := x.NumNeighboursOut()
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
  x.SetDemo (gra.Breadth) // >>> ist eine DijkstraDemo, wenn x bewertet ist
  if x.Conn() {
    x.Act()
    x.Write()
    errh.Error ("ein kürzester Weg in Hinrichtung der Länge", x.NumSub1())
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
      errh.Error ("ein (kürzester) Weg in Rückrichtung der Länge", x.NumSub1())
    } else {
      errh.Error0("es gibt keinen Weg in Rückrichtung")
    }
  }
}

func (x *graphModel) Hamilton (ready, ok Cond, w *bool) {
  x.Mark (true)
  if ready() {
    x.nWays++
    x.Graph.Write()
    errh.Error ("Hamiltonweg", x.nWays)
  } else {
    n := x.NumNeighboursOut()
    for i := uint(0); i < n; i++ {
      x.Step (i, true)
      if x.Marked() || ! ok() {
        x.Step (0, false)
      } else {
        x.Graph.Write(); wait (w)
        x.Hamilton (ready, ok, w)
        x.Mark (false)
        x.Step (0, false)
        x.Graph.Write(); wait (w)
      }
    }
  }
}

func (x *graphModel) Demo (d gra.Demo) {
  x.SetDemo (d)
  x.SetWrite (x.write, x.write1)
}
