package fuday

// (c) murus.org  v. 161216 - license see murus.go

import
  "murus/day"
// Vorlesungszeiten an der Freien Universität Berlin

type FUDay interface { // Kalendertage

// Das Semester von x ist jetzt dasjenige, in dem der Tag x liegt.
  Set (d day.Calendarday)

// b und e sind der erste und letzte Vorlesungstag des Semesters von x.
  Lectures (b, e day.Calendarday)

// Liefert die Kurzbezeichnung des Semesters von x (z.B. "SS 13" oder "WS 13/14").
  String () string

// Liefert genau dann true, wenn d in der Vorlesungszeit
// des Semesters von x liegt und kein Feiertag ist.
  LectureDay (d day.Calendarday) bool

// Liefert die Anzahl der Vorlesungswochen im Semster von x.
  NumWeeks () uint

// Wenn die n-te Woche (beginnend mit n == 1) ab Beginn der Vorlesungszeit
// in der Vorlesungszeit liegt, ist d der Montag dieser Woche.
// Andernfalls ist d leer.
  Monday (d day.Calendarday, n uint)
}
func New() FUDay { return newFU() }
// Liefert das Systemdatum des aufrufenden Rechners als neuen FU-Kalendertag.
// Sein Semester ist das Semester, in dem er liegt.
