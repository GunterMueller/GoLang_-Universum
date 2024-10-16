package errh

// (c) Christian Maurer   v. 240925 - license see µU.go

import (
  "strconv"
  "µU/env"
  "µU/char"
  "µU/str"
  "µU/kbd"
  "µU/col"
  . "µU/scr"
  "µU/N"
)
var (
  transparent bool
//           1         2         3         4         5         6         7         8         9
// 012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012
  license = []string {
  " Das Mikrouniversum µU ist nur zum Einsatz in der Lehre konstruiert  und hat deshalb einen ",
  " rein akademischen Charakter. Es liefert u. a. eine Reihe von Beispielen für mein Lehrbuch ",
  " \"Nichtsequentielle und Verteilte Programmierung mit Go\" (Springer Vieweg 2019) und dessen ",
  " Übersetzung  \"Nonsequential and Distributed Programming with Go\"  (Springer Nature 2021). ",
  " Auch alle Projekte im Buch  \"Objektbasierte Programmierung mit Go\" (Springer Vieweg 2023) ",
  " machen intensiven Gebrauch von diversen Paketen aus dem Mikrouniversum.                   ",
  " Für Zwecke der Lehre an Universitäten und in Schulen sind die Quellen des Mikrouniversums ",
  " uneingeschränkt verwendbar; jede Form weitergehender Nutzung ist jedoch strikt untersagt. ",
  "                                                                                           ",
  " THIS SOFTWARE IS PROVIDED BY THE AUTHORS  \"AS IS\"  AND ANY EXPRESS OR IMPLIED WARRANTIES, ",
  " INCLUDING,  BUT NOT LIMITED TO,  THE IMPLIED WARRANTIES  OF MERCHANTABILITY  AND  FITNESS ",
  " FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY ",
  " DIRECT, INDIRECT,  INCIDENTAL, SPECIAL,  EXEMPLARY, OR CONSEQUENTIAL DAMAGES  (INCLUDING, ",
  " BUT NOT LIMITED TO,  PROCUREMENT OF SUBSTITUTE GOODS  OR SERVICES;  LOSS OF USE, DATA, OR ",
  " PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER ",
  " IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY ",
  " WAY OUT OF THE USE OF THIS SOFTWARE,  EVEN IF ADVISED  OF THE POSSIBILITY OF SUCH DAMAGE. ",
  "                                                                                           ",
  " APART FROM THIS  THE TEXT IN GERMAN ABOVE AND BELOW  IS A MANDATORY PART  OF THE LICENSE. ",
  "                                                                                           ",
  " Die Quelltexte von µU sind äußerst sorgfältig entwickelt und werden laufend überarbeitet. ",
  " ABER:  Es gibt keine fehlerfreie Software - dies gilt natürlich auch für _diese_ Quellen. ",
  " Ihre Verwendung in Programmen könnte zu SCHÄDEN führen, z. B. zum Abfackeln von Rechnern, ",
  " zur Entgleisung von Eisenbahnen, zum GAU in Atomkraftwerken  oder zum Absturz des Mondes. ",
  " Deshalb wird vor der Verwendung irgendwelcher Quellen von µU in Programmen zu ernsthaften ",
  " Zwecken AUSDRÜCKLICH GEWARNT! (Ausgenommen sind Demo-Programme zum Einsatz in der Lehre.) ",
  "                                                                                           ",
  " Meldungen entdeckter Fehler und Hinweise auf Unklarheiten werden sehr dankbar angenommen. "}
  ok = true
)

func init() {
  for i, l := range (license) { license[i] = str.Lat1 (l) }
//                                           1         2         3         4         5         6         7
//                                 012345678901234567890123456789012345678901234567890123456789012345678901234567
  if env.E() {
    ToWait            = str.Lat1 ("please be patien ...")
    ToContinue        = str.Lat1 ("next: Enter")
    ToContinueOrNot   = str.Lat1 ("next: Enter                                                      finished: Esc")
    ToCancel          = str.Lat1 ("cancel: Esc")
    ToScroll          = str.Lat1 ("browse: arrow keys                                                 cancel: Esc")
    ToSelect          = str.Lat1 ("browse: arrow keys                 select: Enter                   cancel: Esc")
    ToChange          = str.Lat1 ("browse: arrow keys                 change: Enter                   cancel: Esc")
    ToSwitch          = str.Lat1 ("browse: arrow keys        select: Enter        toggle: Tab         cancel: Esc")
    ToSelectWithPrint = str.Lat1 ("browse: arrow keys        select: Enter        print: Druck        cancel: Esc")
    ToPrint           = str.Lat1 ("print: Prt                                                 finished: other key")
  } else {
    ToWait            = str.Lat1 ("bitte etwas Geduld ...")
    ToContinue        = str.Lat1 ("weiter: Enter")
    ToContinueOrNot   = str.Lat1 ("weiter: Enter                                                      fertig: Esc")
    ToCancel          = str.Lat1 ("abbrechen: Esc")
    ToScroll          = str.Lat1 ("blättern: Pfeiltasten                                           abbrechen: Esc")
    ToSelect          = str.Lat1 ("blättern: Pfeiltasten              auswählen: Enter             abbrechen: Esc")
    ToChange          = str.Lat1 ("blättern: Pfeiltasten       ändern: Enter       abbrechen: Esc")
    ToSwitch          = str.Lat1 ("blättern: Pfeiltasten    auswählen: Enter    umschalten: Tab    abbrechen: Esc")
    ToSelectWithPrint = str.Lat1 ("blättern: Pfeiltasten    auswählen: Enter    drucken: Druck     abbrechen: Esc")
    ToPrint           = str.Lat1 ("ausdrucken: Druck                                         fertig: andere Taste")
  }
}

func pre() {
  transparent = Transparent()
  if transparent { Transparence (false) }
}

func post() {
  if transparent { Transparence (true) }
}

func delHead() {
  pre()
  Restore (0, 0, NColumns(), 1)
  post()
}

func head (s string) {
  pre()
  w := NColumns()
  Save (0, 0, w, 1)
  Lock()
  s = str.Lat1 (s)
  str.Norm (&s, w)
  Colours (col.HeadF(), col.HeadB())
  Write (s, 0, 0)
  Unlock()
  post()
}

func delHint() {
  if ok {
    return
  }
  ok = true
  pre()
  Restore (NLines() - 1, 0, NColumns(), 1)
  post()
}

func hint (s string) {
  s = str.Lat1 (s)
  if ! ok {
    delHint()
  }
  ok = false
  pre()
  l, w := NLines() - 1, NColumns()
  Save (l, 0, w, 1)
  str.Center (&s, w)
  Lock()
  Colours (col.HintF(), col.HintB())
  Write (s, l, 0)
  Unlock()
  post()
}

func hint1 (s string, k uint) {
  hint (s + " " + N.String (k))
}

func hint2 (s string, k uint, s1 string, k1 uint) {
  hint (s + " " + N.String (k) + " " + s1 + " " + N.String (k1))
}

func delHintPos (s string, l, c uint) {
  Restore (l, c, uint(len (s)), 1)
}

func hintPos (s string, l, c uint) {
  pre()
  if l >= NLines() { l = NLines() - 1 }
  w := uint(len (s))
  if c + w >= NColumns() { c = NColumns() - w }
  Save (l, c, w, 1)
  Lock()
  Colours (col.HintF(), col.HintB())
  Write (s, l, c)
  Unlock()
  post()
}

func do (s string, enter bool) {
  pre()
  s = str.Lat1 (s)
  str.Center (&s, NColumns())
  l := NLines() - 1
  Save (l, 0, NColumns(), 1)
  Lock()
  Colours (col.ErrorF(), col.ErrorB())
  Write (s, l, 0)
//  Colours (col.ErrorF(), col.ErrorB())
  Unlock()
  kbd.Wait (enter)
  Restore (l, 0, NColumns(), 1)
  post()
  Flush()
}

func errorZ (s string, i int) {
  if i < 0 {
    do (s + " -" + N.String (uint(-i)), false)
  } else {
    error (s, uint(i))
  }
}

func errorF (s string, f float64) {
  do (s + " " + strconv.FormatFloat (f, 'e', 10, 64), false)
}

func error0 (s string) {
  do (s, false)
}

func error (s string, k uint) {
  do (s + " " + N.String(k), false)
}

func concm (s string, k uint, s1 string) string {
  s += " " + N.String (k)
  return s + " " + s1
}

func conc2 (s string, k uint, s1 string, k1 uint) string {
  s += " " + N.String (k)
  s1 += " " + N.String(k1)
  return s + " " + s1
}

func conc2m (s string, k uint, s1 string, k1 uint, s2 string) string {
  s += " " + N.String (k)
  s1 += " " + N.String(k1)
  return s + " " + s1 + " " + s2
}

func conc3 (s string, k uint, s1 string, k1 uint, s2 string, k2 uint) string {
  s += " " + N.String (k)
  s1 += " " + N.String(k1)
  s2 += " " + N.String(k2)
  return s + " " + s1 + " " + s2
}

func conc4 (s string, k uint, s1 string, k1 uint, s2 string, k2 uint, s3 string, k3 uint) string {
  s += " " + N.String (k)
  s1 += " " + N.String(k1)
  s2 += " " + N.String(k2)
  s3 += " " + N.String(k3)
  return s + " " + s1 + " " + s2 + " " + s3
}

func errorm (s string, n uint, s1 string) {
  do (concm (s, n, s1), false)
}

func error2 (s string, n uint, s1 string, n1 uint) {
  do (conc2 (s, n, s1, n1), false)
}

func error2m (s string, n uint, s1 string, n1 uint, s2 string) {
  do (conc2m (s, n, s1, n1, s2), false)
}

func error3 (s string, n uint, s1 string, n1 uint, s2 string, n2 uint) {
  do (conc3 (s, n, s1, n1, s2, n2), false)
}

func error4 (s string, n uint, s1 string, n1 uint, s2 string, n2 uint, s3 string, n3 uint) {
  do (conc4 (s, n, s1, n1, s2, n2, s3, n3), false)
}

func error0Pos (s string, l, c uint) {
  pre()
  s = str.Lat1 (s)
  if l >= NLines() { l = NLines() - 1 }
  w := uint(len (s))
  if c + w >= NColumns() { c = NColumns() - w }
  Save (l, c, w, 1)
  Lock()
  Colours (col.ErrorF(), col.ErrorB())
  Write (s, l, c)
  Unlock()
  kbd.Wait (false)
  Restore (l, c, w, 1)
  post()
}

func errorPos (s string, k, l, c uint) {
  s += " " + N.String (k)
  error0Pos(s, l, c)
}

func error2Pos (s string, n uint, s1 string, n1 uint, l, c uint) {
  error0Pos(conc2(s, n, s1, n1), l, c)
}

func confirmed() bool {
  pre()
  s := ""
  if env.E() {
    s = "Are you sure?  y(es / n(o"
  } else {
    s = "Sind Sie sicher?  j(a / n(ein"
  }
  w := NColumns()
  str.Center (&s, w)
  l := NLines() - 1
  Save (l, 0, w, 1)
  Lock()
  Colours (col.ErrorF(), col.ErrorB())
  Write (s, l, 0)
  Unlock()
  b, _, _ := kbd.Read()
  var a bool
  if env.E() {
    a = char.Lower(b) == 'y'
  } else {
    a = char.Lower(b) == 'j'
  }
  Restore (l, 0, w, 1)
  post()
  return a
}

func confirmedYou() bool {
  pre()
  s := ""
  if env.E() {
    s = "Are you sure?  y(es / n(o"
  } else {
    s = "Bist Du sicher?  j(a / n(ein"
  }
  w := NColumns()
  str.Center (&s, w)
  l := NLines() - 1
  Save (l, 0, w, 1)
  Lock()
  Colours (col.ErrorF(), col.ErrorB())
  Write (s, l, 0)
  Unlock()
  b, _, _ := kbd.Read()
  var a bool
  if env.E() {
    a = char.Lower(b) == 'y'
  } else {
    a = char.Lower(b) == 'j'
  }
  Restore (l, 0, w, 1)
  post()
  return a
}

func writeLicense (project, version, author string, f, cl, b col.Colour, g []string, t *string) {
  pre()
  post()
  w, h := uint(len (g[0])), uint(len(license)) /* == len (license), see func init */ + 6
  l, c := (NLines() - h) / 2, (NColumns() - w) / 2
  l0, c0 := l, c
  Save (l0, c0, w, h)
  emptyLine := str.New (w)
  Lock()
  Colours (cl, b)
  Write (emptyLine, l, c)
  Unlock()
  l++
  s := str.Lat1 (project + " v. " + version)
  *t = s
  str.Center (&s, w)
  Lock()
  Colours (cl, b)
  Write (s, l, c)
  l++
  Write (emptyLine, l, c)
  Unlock()
  l++
  s = str.Lat1 ("(c) " + author)
  str.Center (&s, w)
  Lock()
  Colours (f, b)
  Write (s, l, c)
  Unlock()
  l++
// l, c = 30, 52
  Lock()
  Colours (cl, b)
  Write (emptyLine, l, c)
  l++
  for i := 0; i < len (g); i++ {
    Write (g[i], l, c)
    l++
  }
  Write (emptyLine, l, c)
  Unlock()
/*
  l++
  var line string
  if DocExists {
    line = str.Lat1 ("ausführliche Bedienungshinweise: siehe Dokumentation")
  } else {
    line = env.Parameter (0)
    if line == "µU" {
      line = str.New (w)
    } else {
      line = str.Lat1 ("kurze Bedienungshinweise: F1-Taste")
    }
  }
  if ! str.Empty (line) { str.Center (&line, w) }
  Lock()
  Colours (f, b)
  Write (line, l, c); l ++
  Write (emptyLine, l, c)
  Unlock()
*/
//  kbd.Wait (true)
//  scr.Restore (l0, c0, w, h)
}

func µULicense (project, version, author string, f, l, b col.Colour) {
  t := ""; writeLicense (project, version, author, f, l, b, license, &t)
}

func headline (project, version, author string, f, b col.Colour) {
  pre()
  n := NColumns()
  Text := project + "       (c) " + author + "  v. " + version
  str.Center (&Text, n)
  Lock()
  Colours (f, b)
  Write (Text, 0, 0)
  Unlock()
  post()
}

func help (H []string) {
  pre()
  h := uint(len (H))
  var w, l, c uint
  for i := uint(0); i < h; i++ {
    c = uint(len (H[i]))
    if c > w { w = c }
  }
  if h + 2 > NLines() { h = NLines() - 2 }
  if w + 4 > NColumns() { w = NColumns() - 4 }
  mouseOn := MousePointerOn()
  if false { // mouseOn {
    l, c = MousePos()
    if l >= NLines() - h - 1 { l = NLines() - h - 2 }
    if c > NColumns() - w - 4 { c = NColumns() - w - 4 }
    MousePointer (false)
  } else {
    l, c = (NLines() - h - 2) / 2, (NColumns() - w - 4) / 2
  }
  Save (l, c, w + 4, h + 2)
  Lock()
  Colours (col.HintF(), col.HintB())
  t := str.New (w + 4)
  for i := uint(0); i <= h + 1; i++ {
    Write (t, l + i, c)
  }
  for i := uint(0); i < h; i++ {
    str.Center (&H[i], w) 
    Write (H[i], l + 1 + i, c + 2)
  }
  Unlock()
  kbd.Wait (false)
  Restore (l, c, w + 4, h + 2)
  if mouseOn { MousePointer (true) }
  post()
}

func help1() {
  pre()
  s := ""
  if env.E() {
    s = "short operating instructions: F1-key"
  } else {
    s = "kurze Bedienungshinweise: F1-Taste"
  }
  w := uint(len (s))
  l, c := (NLines() - 3) / 2, (NColumns() - w - 4) / 2
  t := str.New (w + 4)
  Save (l, c, w + 4, 3)
  Lock()
  Colours (col.HintF(), col.HintB())
  for i := uint(0); i <= 2; i++ {
    Write (t, l + i, c)
  }
  str.Center (&s, w)
  Write (s, l + 1, c + 2)
  Unlock()
  kbd.Quit()
  Restore (l, c, w + 4, 3)
  post()
}
