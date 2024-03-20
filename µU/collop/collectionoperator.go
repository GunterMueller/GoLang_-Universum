package collop

// (c) Christian Maurer   v. 240317 - license see µU.go

import (
  "µU/env"
  . "µU/obj"
  "µU/str"
  "µU/kbd"
  "µU/errh"
)

func operate (x Collector, o Rotator, f func (x, y Rotator) bool) {
  var help []string
  if env.E() {
    help = []string {"for-/backwards: arrow key down-/upwards",
                     "  to start/end: Pos1/End               ",
                     "  change entry: shift Enter            ",
                     "        insert: Ins                    ",
                     "        delete: Del                    ",
                     "        search: F2                     ",
                     "  change order: F3                     ",
                     "   end program: Esc                    "}
  } else {
    help = []string {"  vor-/rückwärts: Pfeiltaste ab-/aufwärts",
                     " zum Anfang/Ende: Pos1/Ende              ",
                     "  Eintrag ändern: Umschalt-Enter         ",
                     "        einfügen: Einfg                  ",
                     "       entfernen: Entf                   ",
                     "          suchen: F2                     ",
                     "        umordnen: F3                     ",
                     "Programm beenden: Esc                    "}
  }
  for i, h := range (help) { help[i] = str.Lat1 (h) }
  x.Jump (false)
  if x.Empty() {
    for {
      o.(Editor).Edit (0, 0)
      if ! o.Empty() {
        x.Ins (o)
        break
      }
    }
  }
  if env.E() {
    errh.Hint ("help: F1                  end: Esc")
  } else {
    errh.Hint ("Hilfe: F1                  Ende: Esc")
  }
  loop:
  for {
    o = x.Get().(Rotator)
    o.(Editor).Write (0, 0)
    switch c, _ := kbd.Command(); c {
    case kbd.Esc:
      break loop
    case kbd.Enter:
      errh.DelHint()
      o1 := o.Clone().(Rotator)
      o.(Editor).Edit (0, 0)
      if o.Empty() {
        x.Del()
      } else {
        if ! o.Eq (o1) {
          x.Put (o)
        }
      }
      if env.E() {
        errh.Hint ("help: F1                  end: Esc")
      } else {
        errh.Hint ("Hilfe: F1                  Ende: Esc")
      }
    case kbd.Up:
      x.Step (false)
    case kbd.Down:
      x.Step (true)
    case kbd.Pos1:
      x.Jump (false)
    case kbd.End:
      x.Jump (true)
    case kbd.Ins:
      errh.DelHint()
      o.Clr()
      o.(Editor).Edit (0, 0)
      if ! o.Empty() {
        x.Ins (o)
      }
    case kbd.Del:
      if errh.Confirmed() {
        x.Del()
      }
    case kbd.Help:
      errh.Help (help)
    case kbd.Search:
      o.Clr()
      o.(Editor).Edit (0, 0)
      if ! o.Empty() {
        x.Jump (false)
        loop1:
        for {
          o1 := x.Get().(Rotator)
          if f (o, o1) {
            o1.(Editor).Write (0, 0)
            for {
              switch c1, _ := kbd.Command(); c1 {
              case kbd.Esc:
                break loop1
              case kbd.Down:
                for ! x.Eoc (true) {
                  x.Step (true)
                  o2 := x.Get().(Rotator)
                  if f (o, o2) {
                    o2.(Editor).Write (0, 0)
                    break
                  } else {
                  }
                }
              case kbd.Up:
                for ! x.Eoc (false) {
                  x.Step (false)
                  o2 := x.Get().(Rotator)
                  if f (o, o2) {
                    o2.(Editor).Write (0, 0)
                    break
                  }
                }

              case kbd.Del:
                if errh.Confirmed() {
                  x.Del()
                }
              }
            }
          }
          if x.Eoc (true) {
            x.Jump (false)
            break
          }
          x.Step (true)
        }
      }
    case kbd.Act:
      o.(Rotator).Rotate()
      x.Sort()
    }
  }
  errh.DelHint()
}
