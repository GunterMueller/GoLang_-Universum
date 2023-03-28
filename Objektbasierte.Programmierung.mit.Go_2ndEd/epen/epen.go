package main

// (c) Christian Maurer   v. 230112 - license see µU.go

// >>> TODO: Bei Änderung von Kurven die Punkte stehen lassen. Aber: Wann verschwinden sie wieder?
import (
  "µU/time"
  . "µU/obj"
  "µU/env"
  "µU/fontsize"
  "µU/str"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/errh"
  "µU/sel"
  "µU/seq"
  "µU/stk"
  "µU/pseq"
  "µU/files"
  "µU/fig2"
  "µU/psp"
  "µU/prt"
)
const (
  suffix = ".epn"
  lenName = 10
  delta = 10
)
var (
  X, Y int
  tx, ty = 16, 16
  bx = box.New()
  newFigure fig2.Figure2
  figures seq.Sequence
  trash seq.Sequence
  externalFigures  seq.Sequence
  stack stk.Stack
  cF, cB col.Colour
  help = []string {
// 0         1         2         3         4         5         6         7
// 01234567890123456789012345678901234567890123456789012345678901234567890123456789
  "               Art für neue Figur auswählen: Eingabetaste (Enter)               ",
  "                        neue Figur erzeugen: linke Maustaste                    ",
  "                               Figur färben: F3                                 ",
  "             Farbe für neue Figur auswählen: Umschalttaste + F3                 ",
  "                               Figur ändern: Umschalt- + linke Maustaste, Punkte",
  "                                             mit rechter Maustaste verschieben  ",
  "                              Figur löschen: Entf                               ",
  "         letzte gelöschte Figur zurückholen: Rücktaste                          ",
  "                  markierte Figuren löschen: Umschalttaste + Entf               ",
  "                 Hintergrundfarbe auswählen: F4                                 ",
  "                          Figur verschieben: rechte Maustaste                   ",
  "                             Figur kopieren: mittlere Maustaste                 ",
  "             Figur markieren / entmarkieren: F5 / F6                            ",
  "      alle Figuren markieren / entmarkieren: Umschalttaste + F5 / F6            ",
  "        Figuren aus anderer eTafel einfügen: Eingabetaste (Enter)               ",
  "neue eTafel aus markierten Figuren erzeugen: Umschalt- + Eingabetaste           ",
  "        Figur in den Papierkorb und löschen: F7                                 ",
  "           Figur in den Papierkorb schieben: F8                                 ",
  "      alle Figuren aus dem Papierkorb holen: F9                                 ",
  "                         eTafel verschieben: Pfeiltasten                        ",
  " eTafel an die Startposition zurückschieben: Pos1                               ",
  "                          eTafel ausdrucken: Drucken                            ",
  "                               epen beenden: Esc                                ",
  "                                                                                ",
  "                              (Strg wirkt wie Umschalttaste)                    "}
)

func bgColour (c col.Colour) {
  scr.ScrColourB (c)
  scr.Cls()
  write()
}

func load (s seq.Sequence, name string) {
  name += suffix
  n := pseq.Length (name)
  if n == 0 { return }
  buf := make (Stream, n)
  file := pseq.New (buf)
  file.Name (name)
  buf = file.Get().(Stream)
  file.Fin()
  s.Decode (buf)
}

func store (s seq.Sequence, name string) {
  pseq.Erase (name)
  n := s.Codelen()
  buf := make (Stream, n)
  buf = s.Encode()
  file := pseq.New (buf)
  file.Name (name + suffix)
  file.Put (buf)
  file.Fin()
}

func loadPos (name string) {
  file := pseq.New (int(0))
  file.Name (name + ".pos")
  if file.Empty() {
    X, Y = 0, 0
  } else {
    X = file.Get().(int)
    file.Seek (1)
    Y = file.Get().(int)
    file.Fin()
  }
}

func storePos (name string) {
  file := pseq.New (int(0))
  file.Name (name + ".pos")
  file.Clr()
  file.Ins (X)
  file.Ins (Y)
  file.Fin()
}

func write() {
  scr.Buf (true)
  figures.Trav (func (a any) { a.(fig2.Figure2).Write() })
  scr.Buf (false)
}

func found (x, y int) (fig2.Figure2, bool) {
  if figures.ExPred (func (a any) bool { return a.(fig2.Figure2).On (x, y, delta) }, true) {
    f := figures.Get().(fig2.Figure2)
    return f, true
  }
  return nil, false
}

func bl (a any) {
  f := a.(fig2.Figure2)
  if f.Marked() {
    go blink (f)
  }
}

func showMarked() {
  figures.Trav (bl)
}

// func writePos (x, y int) {
//   scr.Colours (col.HeadB(), col.HeadF())
//   scr.Colours (cF, cB)
//   scr.Write ("          ", 0, 0)
//   scr.WriteInt (x, 0, 0)
//   scr.WriteInt (y, 0, 6)
// }

func translateX (x int) {
  figures.Trav (func (a any) { a.(fig2.Figure2).Move (x, 0) })
  X += x
  write()
//  writePos (X, Y)
}

func translateY (y int) {
  figures.Trav (func (a any) { a.(fig2.Figure2).Move(0, y) })
  Y += y
  write()
//  writePos (X, Y)
}

func ins (f fig2.Figure2) {
  figures.Seek (figures.Num())
  figures.Ins (f); write()
}

func setColour (figure fig2.Figure2, c col.Colour) {
  figure.SetColour (c)
  figures.Put (figure)
  figure.Write()
}

func blink (f fig2.Figure2) {
  if f.Marked() {
    for n := 0; n < 3; n++ {
      f.WriteInv(); time.Msleep (100)
      f.Write(); time.Msleep (100)
    }
  }
}

func mark (x, y int, m bool) {
  if f, ok := found (x, y); ok {
    f.Mark (m)
    if m { blink (f) }
    figures.Put (f)
  }
}

func markAll (m bool) {
  figures.Trav (func (a any) { a.(fig2.Figure2).Mark (m) }); write()
}

func editedName() string {
  name1 := str.New(lenName)
  errh.HintPos ("Name der anderen eTafel: ", 0, 0)
  bx.Wd (lenName)
  bx.Colours (col.HintF(), col.HintB())
  bx.Edit (&name1, 0, 26)
  errh.DelHint()
  str.OffSpc (&name1)
  return name1
}

func main() {
//  scr.NewMax(); defer scr.Fin()
  scr.NewWH (300, 0, 1200, 800); defer scr.Fin()
  scr.SetPointer (scr.Crosshair)
  files.Cdp()
  for i, h := range (help) { help[i] = str.Lat1 (h) }
  cF, cB = scr.ScrCols()
  newFigure = fig2.New()
  newFigure.SetColour (cF)
//  newFigure.SetFont (font.Normal)
  figures = seq.New (newFigure)
  trash = seq.New (newFigure)
  externalFigures = seq.New (newFigure)
  stack = stk.New (newFigure)
  name := env.Arg(1)
  if str.Empty (name) { name = "temp" }
  bx.Wd (lenName)
  bx.Colours (col.HeadF(), col.HeadB())
  bx.Edit (&name, 0, 0)
  str.OffSpc (&name)
  load (figures, name)
  loadPos (name)
  write()
  var figure fig2.Figure2
  movable := false
  x0, y0 := 0, 0
  loop:
  for {
    scr.MousePointer (true)
    xm, ym := scr.MousePosGr()
    b, cmd, d := kbd.Read()
    switch cmd {
    case kbd.None, kbd.Enter, kbd.Ins:
      if b == ' ' || b == 'A' || cmd == kbd.Enter || cmd == kbd.Ins {
        newFigure.Select()
      }
      write()
    case kbd.Esc:
      break loop
    case kbd.Back:
      if d == 0 {
        if ! stack.Empty() {
          f := stack.Pop().(fig2.Figure2)
          f.Write()
          ins (f)
        }
      } else {
        for ! stack.Empty() {
          f := stack.Pop().(fig2.Figure2)
          f.Write()
          ins (f)
        }
      }
    case kbd.Left:
      translateX (-tx)
    case kbd.Right:
      translateX (tx)
    case kbd.Up:
      translateY (-ty)
    case kbd.Down:
      translateY (ty)
    case kbd.Pos1:
      translateX (-X)
      translateY (-Y)
//    case kbd.End: case kbd.PgLeft: case kbd.PgRight: case kbd.PgUp: case kbd.PgDown:
    case kbd.Tab:
      showMarked()
    case kbd.Del:
      if d == 0 {
        if f, ok := found (xm, ym); ok {
          figures.Del()
          write()
          stack.Push (f)
        }
      } else {
        figures.Trav (func (a any) { f := a.(fig2.Figure2); if f.Marked() { stack.Push (f) } })
        figures.ClrPred (func (a any) bool { return a.(fig2.Figure2).Marked() })
        write()
      }
    case kbd.Help:
      scr.SetFontsize (fontsize.Normal)
      errh.Help (help)
      write()
    case kbd.Search:
      if f, ok := found (xm, ym); ok {
        if f.Marked() {
          blink (f)
        }
      }
    case kbd.Act:
      f, ok := found (xm, ym)
      if d == 0 {
        if ok {
          if co, oc := sel.Colour (uint(xm) / 8, uint(ym) / 16, 3); oc {
            setColour (f, co)
          }
        }
      } else {
        if co, oc := sel.Colour (uint(xm) / 8, uint(ym) / 16, 3); oc {
          cF = co
          newFigure.SetColour (cF)
        }
      }
      write()
    case kbd.Cfg:
      if co, oc := sel.Colour (uint(xm) / 8, uint(ym) / 16, 3); oc {
        cB = co
        bgColour (cB)
      }
    case kbd.Mark:
      if d == 0 {
        mark (xm, ym, true)
      } else {
        markAll (true)
        showMarked()
      }
    case kbd.Unmark:
      if d == 0 {
        mark (xm, ym, false)
      } else {
        markAll (false)
      }
    case kbd.Cut:
      if f, ok := found (xm, ym); ok {
        trash.Ins (f)
        figures.Del()
        write()
      }
    case kbd.Copy:
      if f, ok := found (xm, ym); ok {
        trash.Ins (f)
      }
    case kbd.Paste:
      if d == 0 {
        figures.Join (trash)
        write()
      } else {
        trash.Clr()
      }
    case kbd.Print:
      p := psp.New()
      p.Name (name)
      bgColour (col.LightWhite())
      figures.Trav (func (a any) { a.(fig2.Figure2).Print (p) })
      bgColour (cB)
      p.Fin()
      prt.PrintImage (name + ".ps")
    case kbd.Roll:
      name1 := editedName()
      if str.Empty (name1) { continue }
			// ? TODO same workaround as with kbd.Drag
      if c, _ := kbd.LastCommand(); c == kbd.Esc { continue }
      if d == 0 {
        if files.IsFile (name1 + suffix) {
          load (externalFigures, name1)
          figures.Join (externalFigures)
          write()
        }
      } else {
        markedFigures := seq.New (newFigure)
        figures.Trav (func (a any) { f := a.(fig2.Figure2); if f.Marked() { markedFigures.Ins (f) } })
        store (markedFigures, name1)
      }
    case kbd.Go:
      scr.WriteMousePosGr (0, 0)
    case kbd.Here:
      if d == 0 {
        newFigure.Clr()
        if newFigure.Typ() == fig2.Image {
          errh.Hint ("Name des Bildes eingeben")
          const lenImageName = 20
          imageName := str.New (lenImageName)
          bx.Wd (lenImageName)
          bx.Colours (col.HeadF(), col.HeadB())
          bx.EditGr (&imageName, xm, ym)
          errh.DelHint()
          str.OffSpc (&imageName)
          if files.IsFile (imageName + ".ppm") {
            if newFigure.ImageEdited (imageName) {
              ins (newFigure)
            }
          }
          write()
        } else {
          errh.Hint ("aktuelle Figur: " + newFigure.String())
          newFigure.Edit()
          errh.DelHint()
          if ! newFigure.Empty() {
            ins (newFigure)
          }
        }
      } else {
        if f, ok := found (xm, ym); ok {
          f.Change()
          figures.Put (f)
          write()
        }
      }
//    case kbd.Drag, kbd.To: // caught in fig2
    case kbd.This:
      figure, movable = found (xm, ym)
    case kbd.Drop:
      scr.WriteMousePosGr (0, 0)
      x0, y0 = xm, ym
      xm, ym = scr.MousePosGr()
      a := uint(figures.Num())
      if movable {
        figure.Move (xm - x0, ym - y0)
        a1 := uint(figures.Num())
        for a1 > a { figures.Seek (a); figures.Del(); write() }
      }
    case kbd.There:
      if movable {
        x0, y0 = scr.MousePosGr()
        figure.Move (xm - x0, ym - y0)
        figures.Put (figure)
        figure.Write()
      }
      write()
    case kbd.That:
      movable = false
      if f, ok := found (xm, ym); ok {
        figure = f.Clone().(fig2.Figure2)
        movable = ok
      }
    case kbd.Move:
      scr.WriteMousePosGr (0, 0)
      x0, y0 = xm, ym
      xm, ym = scr.MousePosGr()
      if movable {
        figure.Move (xm - x0, ym - y0)
      }
    case kbd.Thither:
      if movable {
        x0, y0 = xm, ym
        xm, ym = scr.MousePosGr()
        figure.Move (xm - x0, ym - y0)
        figures.Ins (figure)
        write()
      }
    }
  }
  bx.Wd (lenName)
  bx.Colours (col.HeadF(), col.HeadB())
  bx.Edit (&name, 0, 0)
  str.OffSpc (&name)
  store (figures, name)
  storePos (name)
}
