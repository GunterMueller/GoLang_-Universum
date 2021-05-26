package collop

// (c) Christian Maurer   v. 210525 - license see µU.go

import (
  . "µU/obj"
  "µU/str"
  "µU/kbd"
  "µU/errh"
)

func operate (c Collector, o Indexer, f func (x, y Indexer) bool) {
  help := []string {" vor-/rückwärts: Pfeiltaste ab-/aufwärts",
                    "zum Anfang/Ende: Pos1/Ende              ",
                    " Eintrag ändern: Enter                  ",
                    "       einfügen: Einfg                  ",
                    "      entfernen: Entf                   ",
                    "       umordnen: F3                     ",
                    "         suchen: F2                     ",
                    "   Programmende: Esc                    " }
  for i, h := range (help) { help[i] = str.Lat1 (h) }
  c.Jump (false)
  if c.Empty() {
    for {
      o.Edit (0, 0)
      if ! o.Empty() {
        c.Ins (o)
        break
      }
    }
  }
  tex := ""
  loop:
  for {
    o = c.Get().(Indexer)
    o.Write (0, 0)
    errh.Hint ("Hilfe: F1                  Ende: Esc")
    switch k, _ := kbd.Command(); k {
    case kbd.Enter:
      o1 := o.Clone().(Indexer)
      o.(Indexer).Edit (0, 0)
      if o.Empty() {
        c.Del()
      } else {
        if ! o.Eq (o1) {
          c.Put (o)
        }
      }
    case kbd.Esc:
      break loop
    case kbd.Up:
      c.Step (false)
    case kbd.Down:
      c.Step (true)
    case kbd.Pos1:
      c.Jump (false)
    case kbd.End:
      c.Jump (true)
    case kbd.Ins:
      errh.DelHint()
      o.Clr()
      o.Edit (0, 0)
      if ! o.Empty() {
        c.Ins (o)
      }
    case kbd.Del:
      if errh.Confirmed() {
        c.Del()
      }
    case kbd.Help:
      errh.Help (help)
    case kbd.Search:
      o.Clr()
      o.Edit (0, 0)
      if ! o.Empty() {
        c.Jump (false)
        loop1:
        for {
          o1 := c.Get().(Indexer)
          if f (o, o1) {
            o1.Write (0, 0)
            for {
              switch k, _ := kbd.Command(); k {
              case kbd.Esc:
                break loop1
              case kbd.Down:
                for ! c.Eoc (true) {
                  c.Step (true)
                  o2 := c.Get().(Indexer)
                  if f (o, o2) {
                    o2.Write (0, 0)
                    break
                  }
                }
              case kbd.Up:
                for ! c.Eoc (false) {
                  c.Step (false)
                  o2 := c.Get().(Indexer)
                  if f (o, o2) {
                    o2.Write (0, 0)
                    break
                  }
                }
              }
            }
          }
          if c.Eoc (true) {
            c.Jump (false)
            break
          }
          c.Step (true)
        }
      }
    case kbd.Act:
      o.(Rotator).Rotate()
      c.Sort()
    case kbd.Print:
      tex = ""
      c.Trav (func (a Any) { tex += a.(TeXer).TeX() })
      break loop
    }
  }
  errh.DelHint()
}
