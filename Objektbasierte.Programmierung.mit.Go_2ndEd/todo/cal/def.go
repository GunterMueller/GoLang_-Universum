package cal

// (c) Christian Maurer   v. 210329 - license see µU.go

import
  "µU/day"

func SetFormat (p day.Period) { setFormat(p) }

// Das aktuelle Kalenderblatt ist das vom Tag d.
func Seek (d day.Calendarday) { seek(d) }

// Die Termine im Wochenkalender und alle Terminatribute
// sind jeweils auf dem Bildschirm ausgegeben.
func WriteDay (l, c uint) { writeDay(l,c) }

// func ClearDay (d day.Calendarday, l, c uint) { clearDay(d,l,c) }

// Die Folge der Kalenderblätter ist durch Editieren verändert, wobei bei d begonnen wird.
// d ist danach das Datum des zuletzt editierten Kalenderblattes.
func Edit (d day.Calendarday, l, c uint) { edit (d,l,c) }

// Der aktuelle Suchbegriff ist der an Position (l, c) editierte.
func EditWord (l, c uint) { editWord(l,c) }

// Das aktuelle Kalenderblatt ist ab (Zeile, Spalte) == (l, c) ausgedruckt.
func Print (l, c uint) { print(l,c) }

// Das aktuelle Kalenderblatt ist persistent gesichert.
func Fin() { fin() }
