package errh

// (c) Christian Maurer   v. 171007 - license see µu.go

import (
//  "µu/env"
  "µu/z"
  "µu/str"
  "µu/kbd"
  "µu/scr"
  "µu/col"
  . "µu/scr"
  "µu/box"
  "µu/nat"
)
var (
  errorbox, headbox, hintbox, licenseBox, choiceBox = box.New(), box.New(), box.New(), box.New(), box.New()
  headWritten, hintWritten, hintPosWritten /* , DocExists */ bool
  hintwidth uint
  transparent bool
//           1         2         3         4         5         6         7         8         9
// 012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012
  license = []string {
  " Die Quellen von µu sind nur zum Einsatz in der Lehre konstruiert  und haben deshalb einen ",
  " rein akademischen Charakter. Sie liefern u.a. eine Reihe von Beispielen für mein Lehrbuch ",
  " \"Nichtsequentielle Programmierung mit Go 1 kompakt\"  (Springer, 2. Auflage 2012, 223 S.). ",
  " Für Lehrzwecke in Universitäten und Schulen  sind die Quellen uneingeschränkt verwendbar; ",
  " jegliche weitergehende - insbesondere kommerzielle - Nutzung ist jedoch strikt untersagt. ",
  "                                                                                           ",
  " THIS SOFTWARE IS PROVIDED BY THE AUTHORS  \"AS IS\"   AND ANY EXPRESS OR IMPLIED WARRANTIES, ",
  " INCLUDING,  BUT NOT LIMITED TO,  THE IMPLIED WARRANTIES  OF MERCHANTABILITY  AND  FITNESS ",
  " FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY ",
  " DIRECT, INDIRECT,  INCIDENTAL, SPECIAL,  EXEMPLARY, OR CONSEQUENTIAL DAMAGES  (INCLUDING, ",
  " BUT NOT LIMITED TO,  PROCUREMENT OF SUBSTITUTE GOODS  OR SERVICES;  LOSS OF USE, DATA, OR ",
  " PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER ",
  " IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY ",
  " WAY OUT OF THE USE OF THIS SOFTWARE,  EVEN IF ADVISED  OF THE POSSIBILITY OF SUCH DAMAGE. ",
  " APART FROM THIS  THE TEXT IN GERMAN ABOVE AND BELOW  IS A MANDATORY PART  OF THE LICENSE. ",
  "                                                                                           ",
  " Die Quelltexte von µu sind äußerst sorgfältig entwickelt und werden laufend überarbeitet. ",
  " ABER: Es gibt keine fehlerfreie Software - dies gilt natürlich auch für diese Quelltexte. ",
  " Ihre Verwendung in Programmen könnte zu SCHÄDEN führen, z. B. zum Abfackeln von Rechnern, ",
  " zur Entgleisung von Eisenbahnen, zum GAU in Atomkraftwerken  oder zum Absturz des Mondes. ",
  " Deshalb wird vor der Verwendung irgendwelcher Quellen von µu in Programmen zu ernsthaften ",
  " Zwecken AUSDRÜCKLICH GEWARNT! (Ausgenommen sind Demo-Programme zum Einsatz in der Lehre.) ",
  "                                                                                           ",
  " Meldungen entdeckter Fehler und Hinweise auf Unklarheiten werden sehr dankbar angenommen. " }
//  actualFontsize FontSizes
  first bool = true
)

func init() {
  for i, l := range (license) { license[i] = str.Lat1 (l) }
  headbox.Colours (col.HeadF(), col.HeadB())
  hintbox.Colours (col.HintF(), col.HintB())
//  errorbox.Colours (col.ErrorF, col.ErrorB)
//  pre() TODO theScreen not yet defined
//  post() TODO theScreen not yet defined
//                                         1         2         3         4         5         6         7
//                               012345678901234567890123456789012345678901234567890123456789012345678901234567
  ToWait            = str.Lat1 ("einen Augenblick bitte ...")
  ToContinue        = str.Lat1 ("weiter: Einter")
  ToContinueOrNot   = str.Lat1 ("weiter: Einter                                                     fertig: Esc")
  ToCancel          = str.Lat1 ("                                                                abbrechen: Esc")
  ToScroll          = str.Lat1 ("blättern: Pfeiltasten                                           abbrechen: Esc")
  ToSelect          = str.Lat1 ("blättern/auswählen/abbrechen: Pfeiltasten/Enter/Esc, Maus bewegen/links/rechts")
  ToChange          = str.Lat1 ("blättern: Pfeiltasten       ändern: Enter       abbrechen: Esc")
  ToSwitch          = str.Lat1 ("blättern: Pfeiltasten    auswählen: Enter    umschalten: Tab    abbrechen: Esc")
  ToSelectWithPrint = str.Lat1 ("blättern: Pfeiltasten    auswählen: Enter    drucken: Druck     abbrechen: Esc")
  ToPrint           = str.Lat1 ("ausdrucken: Druck                                         fertig: andere Taste")
}

func wait() { // TODO -> kbd, other name
  loop: for {
    _, c, _  := kbd.Read()
    switch c { case kbd.Enter, kbd.Esc, kbd.Back, kbd.Here, kbd.There:
      break loop
    }
  }
}

func pre() {
  transparent = Transparent()
  if transparent { Transparence (false) }
  if first {
    first = false
    errorbox = box.New()
    errorbox.Colours (col.ErrorF(), col.ErrorB())
  }
//  actualFontsize = ActFontsize()
//  if actualFontsize # Normal {
//    SwitchFontsize (Normal)
//  }
//  hintbox.Wd (width)
//  errorbox.Wd (width)
}

func post() {
  if transparent { Transparence (true) }
//  if actualFontsize # Normal {
//    SwitchFontsize (Normal)
//  }
}

func head (s string) {
  delHead()
  pre()
  w := NColumns()
  Save (0, 0, w, 1)
  headbox.Wd (w)
  str.Norm (&s, w)
  headbox.Write (s, 0, 0)
  headWritten = true
  post()
}

func delHead() {
  pre()
  if headWritten {
    headWritten = false
    Restore (0, 0, NColumns(), 1)
  }
  post()
}

func hint (s string) {
  delHint()
  pre()
  w := NColumns()
  s = str.Lat1 (s)
  str.Center (&s, w)
  l := NLines() - 1
  Save (l, 0, w, 1)
  hintbox.Wd (w)
  hintbox.Write (s, l, 0)
  hintWritten = true
  post()
}

func hint1 (s string, n uint) {
  hint (s + " " + nat.String (n))
}

func hint2 (s string, n uint, s1 string, n1 uint) {
  hint (s + " " + nat.String (n) + " " + s1 + " " + nat.String (n1))
}

func delHint() {
  pre()
  if hintWritten {
    hintWritten = false
    Restore (NLines() - 1, 0, NColumns(), 1)
  }
  post()
}

func hintPos (s string, l, c uint) {
//  delHintPos (s)
  pre()
  if l >= NLines() { l = NLines() - 1 }
  w := uint(len (s))
  if c + w >= NColumns() { c = NColumns() - w }
  Save (l, c, w, 1)
  hintbox.Wd (w)
  hintbox.Write (s, l, c)
  hintPosWritten = true
  post()
}

func delHintPos (s string, l, c uint) {
  if hintPosWritten {
    hintPosWritten = false
    Restore (l, c, uint(len (s)), 1)
  }
}

func do (s string, enter bool) {
  pre()
  s = str.Lat1 (s)
  str.Center (&s, NColumns())
  l := NLines() - 1
  Save (l, 0, NColumns(), 1)
  errorbox.Wd (NColumns())
  errorbox.Write (s, l, 0)
  errorbox.Colours(col.ErrorF(), col.ErrorB())
  kbd.Wait (enter)
  Restore (l, 0, NColumns(), 1)
  post()
  Flush()
}

func proceed0 (s string) {
  do (s, true)
}

func error0 (s string) {
  do (s, false)
}

func proceed (s string, n uint) {
  do (s + " " + nat.String(n), true)
}

func error (s string, n uint) {
  do (s + " " + nat.String(n), false)
}

func conc2 (s string, n uint, s1 string, n1 uint) string {
  s += " " + nat.String (n)
  s1 += " " + nat.String(n1)
  return s + " " + s1
}

func conc3 (s string, n uint, s1 string, n1 uint, s2 string, n2 uint) string {
  s += " " + nat.String (n)
  s1 += " " + nat.String(n1)
  s2 += " " + nat.String(n2)
  return s + " " + s1 + " " + s2
}

func proceed2 (s string, n uint, s1 string, n1 uint) {
  do (conc2(s, n, s1, n1), true)
}

func error2 (s string, n uint, s1 string, n1 uint) {
  do (conc2(s, n, s1, n1), false)
}

func error3 (s string, n uint, s1 string, n1 uint, s2 string, n2 uint) {
  do (conc3(s, n, s1, n1, s2, n2), false)
}

func error0Pos (s string, l, c uint) {
  pre()
  s = str.Lat1 (s)
  if l >= NLines() { l = NLines() - 1 }
  w := uint(len (s))
  if c + w >= NColumns() { c = NColumns() - w }
  Save (l, c, w, 1)
  errorbox.Wd (w)
  errorbox.Write (s, l, c)
  kbd.Wait (false)
  Restore (l, c, w, 1)
  post()
}

func errorPos (s string, n, l, c uint) {
  s += " " + nat.String(n)
  error0Pos(s, l, c)
}

func error2Pos (s string, n uint, s1 string, n1 uint, l, c uint) {
  error0Pos(conc2(s, n, s1, n1), l, c)
}

func confirmed() bool {
  pre()
  s := "Sind Sie sicher?  j(a / n(ein"
  w := NColumns()
  str.Center (&s, w)
  l := NLines() - 1
  Save (l, 0, w, 1)
  errorbox.Wd (w)
  errorbox.Write (s, l, 0)
  b, _, _ := kbd.Read()
  a := z.Lower(b) == 'j'
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
  scr.Save (l0, c0, w, h)
  licenseBox.Wd (w)
  licenseBox.Colours (cl, b)
  emptyLine := str.Clr (w)
  licenseBox.Write (emptyLine, l, c); l ++
  s := str.Lat1 (project + " v. " + version)
  *t = s
  str.Center (&s, w)
  licenseBox.Write (s, l, c); l ++
  licenseBox.Write (emptyLine, l, c); l ++
  s = str.Lat1 ("(c) " + author)
  str.Center (&s, w)
  licenseBox.Colours (f, b)
  licenseBox.Write (s, l, c); l ++ // l, c = 30, 52
  licenseBox.Colours (cl, b)
  licenseBox.Write (emptyLine, l, c); l ++
  for i := 0; i < len (g); i++ {
    licenseBox.Write (g[i], l, c); l ++
  }
  licenseBox.Write (emptyLine, l, c); l ++
  licenseBox.Colours (f, b)
/*
  var line string
  if DocExists {
    line = str.Lat1 ("ausführliche Bedienungshinweise: siehe Dokumentation")
  } else {
    line = env.Parameter (0)
    if line == "µu" {
      line = str.Clr (w)
    } else {
      line = str.Lat1 ("kurze Bedienungshinweise: F1-Taste")
    }
  }
  if ! str.Empty (line) { str.Center (&line, w) }
  licenseBox.Write (line, l, c); l ++
  licenseBox.Write (emptyLine, l, c)
*/
//  kbd.Wait (true)
//  scr.Restore (l0, c0, w, h)
}

func µuLicense (project, version, author string, f, l, b col.Colour) {
  t := ""; writeLicense (project, version, author, f, l, b, license, &t)
}

func headline (project, version, author string, f, b col.Colour) {
  pre()
  n := NColumns()
  Text := project + "       (c) " + author + "  v. " + version
  str.Center (&Text, n)
  licenseBox.Wd (n)
  licenseBox.Colours (f, b)
  licenseBox.Write (Text, 0, 0)
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
  hintbox.Wd (w + 4)
  T := str.Clr (w + 4)
  for i := uint(0); i <= h + 1; i++ {
    hintbox.Write (T, l + i, c)
  }
  hintbox.Wd (w)
  for i := uint(0); i < h; i++ {
    hintbox.Write (H[i], l + 1 + i, c + 2)
  }
  wait()
  Restore (l, c, w + 4, h + 2)
  if mouseOn { MousePointer (true) }
  post()
}

func help1() {
  pre()
  s := "kurze Bedienungshinweise: F1-Taste"
  w := uint(len (s))
//  mouseOn := MousePointerOn()
  var l, c uint
  if false { // mouseOn {
    l, c = MousePos()
    if l >= NLines() - 2 { l = NLines() - 3 }
    if c > NColumns() - w { c = NColumns() - w }
    MousePointer (false)
  } else {
    l = (NLines() - 3) / 2
    c = (NColumns() - w - 4) / 2
  }
  hintbox.Wd (w + 4)
  t := str.Clr (w + 4)
  Save (l, c, w + 4, 3)
  for i := uint(0); i <= 2; i++ { hintbox.Write (t, l + i, c) }
  hintbox.Wd (w)
  hintbox.Write (s, l + 1, c + 2)
  wait()
  Restore (l, c, w + 4, 3)
//  if mouseOn { MousePointer (true) }
  post()
}
