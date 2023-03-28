package dayattr

// (c) Christian Maurer   v. 210723 - license see µU.go

import
  "µU/day"

// Verwaltet eine Menge von Tagesattributen, von denen eins aktuell ist,
// und zu jedem der Attribute die persistente Menge derjenigen Tage,
// die dieses Attribut haben. Anfangs ist das erste Attribut
// in der Datei "Tagesattribute.kfg" das aktuelle Attribut und
// für jedes Attribut ist die persistente Menge der Tage, die es
// haben, diejenige, die es beim vorigen Aufruf des Programms war.

// Das aktuelle Attribut ist das erste in der Datei "Tagesattribute.kfg".
func Normalize() { normalize() }

// Für w == true ist das aktuelle Attribut in der zyklischen Reihenfolge
// der Attribute um eins weiter und für w == false um eins zurückgesetzt.
func Change (w bool) { change(w) }

// Das aktuelle Attribut ist auf dem Bildschirm
// in der Position (Zeile, Spalte) = (l, c) ausgegeben.
func WriteActual (l, c uint) { writeActual(l,c) }

// Für b == true trifft das aktuelle Attribut auf d zu, für b == false
// nicht; die Menge der Tage, die es haben, ist entsprechend geändert.
func Actualize (d day.Calendarday, b bool) { actualize (d,b) }

// Die Menge der Tage, die das erste Attribut
// aus der Datei "Tagesattribute.kfg" haben, ist leer.
func Clr() { clr() }

// Farbe und Font von d sind - je nachdem, ob d in der Menge
// der Tage, die das aktuellen Attributs haben, enthalten ist
// oder nicht und ob d ein Feiertag ist oder nicht - gesetzt.
func Attrib (d day.Calendarday) { attrib(d) }

// Für jedes Attribut ist die Menge der Tage, die es haben,
// in der Datei "Tagesattribute.kfg" gesichert.
func Fin() { fin() }
