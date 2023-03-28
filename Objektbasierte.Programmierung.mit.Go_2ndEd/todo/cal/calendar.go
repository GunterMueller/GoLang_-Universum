package cal

// (c) Christian Maurer   v. 211215 - license see µU.go

import (
  . "µU/obj"
  "µU/env"
  "µU/kbd"
  "µU/errh"
  "µU/day"
  "µU/piset"
  "µU/files"
  "todo/word"
  "todo/page"
)
var (
  globalDay = day.New()
  globalPage = page.New()
  content = piset.New (globalPage.(Indexer))
)

func init() {
  files.Cd (env.Gosrc() + "/todo/")
  content.Name ("Termine")
}

func setFormat (p day.Period) {
  globalPage.SetFormat (p)
}

func seek (d day.Calendarday) {
  globalPage.Set (d)
  if content.Ex (globalPage) { // richtige Seite gefunden
    globalPage = content.Get().(page.Page)
  } else {
    globalPage.Clr()
  }
}

func writeDay (l, c uint) {
  globalPage.Write (l, c)
}

func edit (d day.Calendarday, l, c uint) {
  globalPage.Set (d)
  globalPage.SetFormat (day.Daily)
  exists := content.Ex (globalPage)
  if exists { // haben wir an diesem Tag Termine
    errh.Hint (errh.ToSelect)
    loop:
    for {
      globalPage = content.Get().(page.Page)
      globalPage.Write (l, c)
      switch k, _ := kbd.Command(); k {
      case kbd.Enter:
        break loop
      case kbd.Esc:
        errh.DelHint()
        return
      case kbd.Down, kbd.Up:
        content.Step (k == kbd.Down)
      case kbd.Pos1, kbd.End:
        content.Jump (k == kbd.End)
      case kbd.Print:
        globalPage.Print (0, 0)
      }
    }
    errh.DelHint()
    globalDay = globalPage.Day()
  }
  globalPage.Edit (l, c)
  if globalPage.Empty() {
    if exists {
      content.Del()
    }
  } else if exists {
    content.Put (globalPage)
  } else {
    content.Ins (globalPage)
  }
}

func editWord (l, c uint) {
  word.EditActual (l, c)
}

func print (l, c uint) {
  globalPage.Print (l, c)
}

func fin() {
  globalPage.Fin()
}
