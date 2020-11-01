package pids

// Christian Maurer   v. 201017 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/str"
  "µU/errh"
  "µU/piset"
  "µU/files"
)
type
  persistentIndexerSequence struct {
                                   Any
                              name string
                              f, b col.Colour
                               all piset.PersistentIndexedSet
                                   }

var
  help = []string {" vor-/rückwärts: Pfeiltaste auf-/abwärts",
                   "zum Anfang/Ende: Pos1/End               ",
                   " Eintrag ändern: Enter                  ",
                   "       einfügen: Einfg                  ",
                   "      entfernen: Entf                   ",
//                   "         suchen: F2                     ",
                   "Einträge umordnen: F3                   ",
                   "   Programmende: Esc                    " }

func init() {
  for i, h := range (help) { help[i] = str.Lat1 (h) }
}

func new_(a Any, n string) PersistentIndexerSequence {
  x := new(persistentIndexerSequence)
  x.Any = Clone(a)
  if ! IsIndexer (x.Any) || ! col.IsColourer (x.Any) { ker.Panic ("geht nicht") }
  x.name = n
  x.f, x.b = scr.ScrCols()
  x.all = piset.New (a.(Object), a.(Indexer).Index())
  files.Cd0()
  x.all.Name (n)
  return x
}

func (x *persistentIndexerSequence) Colours (f, b col.Colour) {
  x.f, x.b = f, b
}

func (x *persistentIndexerSequence) Operate (l, c uint) {
  hint := "Kommandotaste drücken       Hilfe: F1       Programm beenden: Esc"
  x.all.Jump (false)
  if x.all.Empty() {
    a := x.Any.(Indexer)
    a.Clr()
    for {
      a.(Editor).Edit (l, c)
      if a.Empty() {
        errh.Error0 ("ist leer")
      } else {
        x.all.Ins (a)
        return
      }
    }
  }
  loop:
  for {
    a := x.all.Get().(Indexer)
    a0 := Clone(a).(Indexer)
    a.(Editor).Write (l, c)
    errh.Hint (hint)
    switch k, _ := kbd.Command(); k {
    case kbd.Enter:
      a.(Editor).Edit (l, c)
      if a.Empty() {
        x.all.Del()
      } else {
        if ! Eq (a, a0) {
          x.all.Put (a)
        }
      }
    case kbd.Esc:
      break loop
    case kbd.Up, kbd.Left:
      x.all.Step (false)
    case kbd.Down, kbd.Right:
      x.all.Step (true)
    case kbd.Pos1:
      x.all.Jump (false)
    case kbd.End:
      x.all.Jump (true)
    case kbd.Ins:
      errh.DelHint()
      a.Clr()
      a.(Editor).Edit (l, c)
      if ! a.Empty() {
        x.all.Ins (a)
      }
      errh.Hint (hint)
    case kbd.Del:
      if errh.Confirmed() {
        x.all.Del()
      }
    case kbd.Act:
      x.Any.(Orderer).RotOrder()
      x.all.Sort()
    case kbd.Help:
      errh.Help (help)
      errh.Hint (hint)
    case kbd.Search:
      a.(Editor).Edit (l, c)
/*/
      x.all.TravPred ()
/*/
    }
    errh.DelHint()
  }
}
