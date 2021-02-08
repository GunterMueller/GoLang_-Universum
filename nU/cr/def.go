package cr

// (c) Christian Maurer   v. 210123 - license see nU.go

/* Objekte, die den Zugang zu kritischen Abschnitten synchronisieren
   (d.h. Betriebsystem-Ressourcen oder gemeinesame Daten) für p Prozessklassen
   P(0), P(1), ..., P(p-1) (p > 1), r Ressourcen R(0), R(1), ... R(r-1)
   (r > 0) und p * r Zahlen m(k,n) (k < p, n < r) so, dass jede Ressource
   unter gegenseitigem Ausschluss nur von Prozessen verschiedener Klassen
   betreten werden kann und die Ressource R(n) (n < r) nur von höchstens
   m(k,n) Prozessen der Klasse P(k) (k < p) nebenläufig betreten werden kann.
   Beispiele:
   a) Leser-Schreiber-Problem:
      p = 2, P(0) = Leser and P(1) = Schreiber; r = 1 (R(0) = gemeinsame Daten),
      m(0,0) = MaxNat, m(1,0) = 1.
   b) Links-Rechts-Problem:
      p = 2, P(1) = Linksfahrende und P(2) = Rechtsfahrende; r = 1
      (R(0) = gemeinsame Spur), m(0,0) = m(1,0) = MaxNat.
   c) beschränkter Links-Rechts-Problem:
      wie oben mit Schrankens m(0,0), m(1,0) < MaxNat.
   d) Bowling Problem: 
      p = Anzahl der teilnehmenden Vereine (P(k) = Spieler von Verein k);
      r = Anzahl der verfügbaren Bowlingbahnen (R(a) = Bowlingbahn a),
      m(c,a) = Maximalzahjl der Spieler des Vereins k auf der Bahn alley a.
*/
type
  CriticalResource interface {

// Vor.: m[i][r] ist für alle i < Anzahl der Klassen
//       und für alle r < Anzahl der Ressourcen von x definiert.
// Auf die Ressource r von x kann von höchstens m[i][r]
// Prozessen der Klasse i zugegriffen werden.
  Limit (m [][]uint)

// Vor.: i < Anzahl der Klassen von x. Der aufrufende Prozess
//       hat keinen Zugriff auf eine Ressource von x.
// Liefert die Anzahl der Ressourcen, auf die
// der aufrufende Prozess jetzt zugreifen kann.
// Er war ggf. solange blockiert, bis das möglich war.
  Enter (i uint) uint

// Vor.: i < Anzahl der Klassen von x. Der aufrufende Prozess
//       hat Zugriff auf eine Ressource von x.
// Er hat jetzt nicht mehr den Zugriff.
  Leave (i uint)
}

// Liefert eine neue kritische Ressource
// mit k Klassen und r Ressourcen.
func New (k, r uint) CriticalResource { return new_(k,r) }
