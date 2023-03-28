package rob

// (c) Christian Maurer   v. 230309 - license see µU.go

/* Verwaltet Roboter, die in einer rechteckigen Welt
   aus schachbrettförmig angeordneten Plätzen leben.
   Die Welt ist M Plätze breit und hoch. Auf den Plätzen
   können Klötze liegen oder sie können zugemauert sein.

   Jeder Roboter steht immer auf einem der Plätze, der als
   "R's Platz" bezeichnet wird. Er steht immer in einer der vier
   Himmelsrichtungen, die als "R's Richtung" bezeichnet wird.

   Jeder Roboter hat eine Tasche mit anfangs Max Klötzen
   und jederzeit Zugriff auf Mauersteine.

   Die Plätze der Roboter, ihre Richtungen und die Anzahl der
   Klötze in ihren Taschen sind die gleichen wie beim letzten
   Programmlauf mit dieser Welt. Wenn die Welt neu ist,
   ist sie leer und ein Roboter steht in Richtung Süd in der
   nordwestlichen Ecke und hat Max Klötze, abzgl. derjenigen,
   die er schon in die Welt gelegt hat - in seiner Tasche.

   Anfangs ist das Protokoll nicht eingeschaltet.

   Der aufrufende Roboter wird immer mit R bezeichnet. */

import
  . "µU/obj"
const (
  M = 24 // Anzahl der Plätze pro Reihe und Spalte der Welt
  Max = 999
)
type
  Roboter interface {

  Coder

// Liefert die Nummer von R.
  Nummer() uint

// R hat sich um 90 Grad nach links gedreht.
  LinksDrehen()

// R hat sich ist um 90 Grad nach rechts gedreht.
  RechtsDrehen()

// Liefert genau dann true,
// wenn R in der nordwestlichen Ecke steht.
  InLinkerObererEcke() bool

// Liefert genau dann true, wenn R's Platz
// in R's Richtung keinen Nachbarplatz hat.
  VorRand() bool

// Vor.: R steht nicht am Rand und der Nachbarplatz
//       in R's Richtung ist nicht zugemauert.
// R steht in der gleichen Richtung wie vorher
// auf diesem Nachbarplatz.
  Laufen()

// Vor.: R's Platz hat entgegen R's Richtung einen Nachbarplatz,
//       der nicht zugemauert ist.
// R steht in der gleichen Richtung wie vorher
// auf diesem Nachbarplatz.
  Zurücklaufen()

// Liefert genau dann true, wenn auf R's Platz kein Klotz liegt.
  Leer() bool

// Liefert genau dann true, wenn R's Platz
// in R's Richtung einen Nachbarplatz hat
// und auf diesem Platz kein Klotz liegt.
  NachbarLeer() bool

// Liefert genau dann true, wenn R's Tasche nicht leer ist.
  HatKlötze() bool

// Liefert die Anzahl der Klötze in R's Tasche.
  AnzahlKlötze() uint

// Vor.: R's Tasche ist nicht leer. 
// Auf R's Platz liegt ein Klotz mehr als vorher,
// in seiner Tasche ist einer weniger.
  Ablegen()

// Vor.: Auf R's Platz liegt mindestens ein Klotz.
// Auf R's Platz liegt ein Klotz weniger als vorher,
// in seiner Tasche ist einer mehr.
  Aufnehmen()

// Liefert genau dann true, wenn R's Platz in R's Richtung einen
// Nachbarplatz hat, auf dem genau ein Klotz liegt und dieser
// Nachbarplatz wiederum einen Nachbarplatz in R's Richtung hat,
// der leer und nicht zugemauert ist. In diesem Fall steht R in
// der gleichen Richtung wie vorher auf dem vorherigen Nachbarplatz
// in seiner Richtung und der Klotz, der vorher auf ihm lag,
// liegen jetzt auf dem Nachbarplatz in R's Richtung.
  Geschoben() bool

// Vor.: R's Platz hat in R's Richtung einen Nachbarplatz,
//       auf dem genau ein Klotz liegt. Dieser Klotz ist so
//       weit wie möglich (d.h. ohne Kollision mit Robotern,
//       Klötzen oder Mauern) in R's Richtung weitergeschoben.
  Schießen()

// R's Platz ist markiert.
  Markieren()

// R's Platz ist nicht markiert.
  Entmarkieren()

// Liefert genau dann true, wenn R's Platz markiert ist.
  Markiert() bool

// Liefert genau dann true, wenn R's Platz in R's Richtung
// einen Nachbarplatz hat, der markiert ist.
  NachbarMarkiert() bool

// Liefert genau dann true, wenn R's Platz in R's Richtung
// einen Nachbarplatz hat, der zugemauert ist.
  VorMauer() bool

// Vor.: R's Platz hat in R's Richtung einen Nachbarplatz,
//       der nicht zugemauert ist.
// R steht in der gleichen Richtung wie vorher auf diesem
// Nachbarplatz. Wenn auf dem Platz, auf dem R vorher
// gestanden hat, Klötze gelegen haben, liegen sie dort
// jetzt nicht mehr, sondern sind in seiner Tasche;
// dafür ist dieser Platz jetzt zugemauert. Eine vorher
// dort etwa vorhandene Markierung ist jetzt entfernt.
  Zumauern()

// Vor.: R's Platz hat in R's Richtung einen Nachbarplatz,
//       der zugemauert ist.
// R steht in der gleichen Richtung wie vorher auf diesem
// Nachbarplatz und der ist jetzt nicht mehr zugemauert.
  Entmauern()

// Die Roboterwelt ist auf dem Bildschirm ausgegeben.
  Ausgeben()

// Liefert R's Position.
  Pos() (uint, uint)

// R's Position ist (x, y).
  Set (x, y uint)
}

// Die Roboterwelt ist die, deren Name beim
// Programmaufruf als Parameter übergeben wurde.
func Laden (s ...string) { laden(s...) }

func AnzahlRoboter() uint { return uint(nRoboter) }

// Die Klötze und Mauersteine in der Roboterwelt liegen
// auf den von der benutzenden Person festgelegten Plätzen.
// Wenn R auf einem Platz steht, auf dem Klötze liegen,
// ist deren Anzahl angezeigt. R's Platz und Richtung,
// die Anzahl der Klötze in seiner Tasche und der Klötze
// auf allen Plätzen sind beim nächsten Programmlauf mit
// dieser Welt die gleichen wie beim Aufruf dieser Methode.
// Wenn das Protokoll eingeschaltet ist,
// ist der Editiervorgang in einem Go-Quelltext (unter dem
// Namen der Roboterwelt mit dem Suffix ".go") protokolliert.
// Das aus diesem Quelltext durch Übersetzung erzeugte Programm
// simuliert schrittweise den Editiervorgang.
func Editieren() { editieren() }

// Vor.: x < M, y < M.
// Liefert auf der Position (x, y) einen neuen Roboter
// mit Max Klötzen in der Tasche.
func NeuerRoboter (x, y uint) Roboter { return neuerRoboter(x,y) }

// Vor.: n > 0. Es gibt mindestensn Roboter in der Welt.
// Liefert den Roboter mit der Nummer n.
func Nr (n uint) Roboter { return alle[n] }

// Die Roboterwelt ist auf dem Bildschirm ausgegeben.
func WeltAusgeben() { weltAusgeben() }

// Nach dem Editieren ist das Programm erzeugt,
// das den Editierverlauf wiedergibt.
func ProgrammErzeugen() { programmErzeugen() }

// Das Protokoll ist genau dann eingeschaltet,
// wenn ein == true (siehe editieren).
func ProtokollSchalten (ein bool) { protokollSchalten(ein) }

// Für ein == true ist das Verhalten des Editors
// gemäß den Anforderungen an das Spiel Sokoban vereinfacht.
func SokobanSchalten (ein bool) { sokobanSchalten(ein) }

// n ist in der untersten Bildschirmzeile ausgegeben.
// Der aufrufende Prozess war danach solange angehalten,
// bis Benutzerix die Ausgabe mit <Esc> quittiert hatte.
func Ausgabe (n uint) { ausgabe(n) }

// Liefert die Zahl, die von der benutzenden Person
// in der untersten Bildschirmzeile eingegeben wurde.
func Eingabe() uint { return eingabe() }

// s und n sind in einer Zeile am unteren Bildschirmrand
// ausgegeben. Der aufrufende Prozess war danach solange
// angehalten, bis die Ausgabe mit <Esc> quittiert wurde.
// Jetzt ist die Meldung wieder vom Bildschirm entfernt.
func FehlerMelden (s string, n uint) { fehlerMelden(s,n) }

// s und n sind in einer Zeile
// am unteren Bildschirmrand ausgegeben.
func HinweisAusgeben (s string, n uint) { hinweisAusgeben(s,n) }

// Das Programm ist mit der Fehlermeldung ("Programm beendet")
// angehalten.
func Fertig() { fertig() }
