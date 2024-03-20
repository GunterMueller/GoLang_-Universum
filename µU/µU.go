package main
/*/
  (c) 1986-2024  Christian Maurer        maurer-berlin.eu proprietary - all rights reserved

  Das Mikrouniversum µU ist nur zum Einsatz in der Lehre konstruiert  und hat deshalb einen
  rein akademischen Charakter. Es liefert u. a. eine Reihe von Beispielen für mein Lehrbuch
  "Nichtsequentielle und Verteilte Programmierung mit Go" (Springer Vieweg 2019) und dessen
  Übersetzung  "Nonsequential and Distributed Programming with Go"  (Springer Nature 2021).
  Auch alle Projekte im Buch  "Objektbasierte Programmierung mit Go" (Springer-Vieweg 2023)
  machen intensiven Gebrauch von diversen Paketen aus dem Mikrouniversum.
  Für Zwecke der Lehre an Universitäten und in Schulen sind die Quellen des Mikrouniversums
  uneingeschränkt verwendbar; jede Form weitergehender Nutzung ist jedoch strikt untersagt.

  THIS SOFTWARE  IS PROVIDED BY THE AUTHOR  "AS IS"  AND ANY EXPRESS OR IMPLIED WARRANTIES,
  INCLUDING,  BUT NOT LIMITED TO,  THE IMPLIED WARRANTIES  OF MERCHANTABILITY  AND  FITNESS
  FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL  THE AUTHOR BE LIABLE FOR ANY
  DIRECT, INDIRECT,  INCIDENTAL, SPECIAL,  EXEMPLARY, OR CONSEQUENTIAL DAMAGES  (INCLUDING,
  BUT NOT LIMITED TO,  PROCUREMENT OF SUBSTITUTE GOODS  OR SERVICES;  LOSS OF USE, DATA, OR
  PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER
  IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY
  WAY OUT OF THE USE OF THIS SOFTWARE,  EVEN IF ADVISED  OF THE POSSIBILITY OF SUCH DAMAGE.

  APART FROM THE ABOVE, THE GERMAN TEXT ABOVE AND BELOW IS A MANDATORY PART OF THE LICENSE.

  Die Quelltexte von µU sind äußerst sorgfältig entwickelt und werden laufend überarbeitet.
  ABER: Es gibt keine fehlerfreie Weichware - dies gilt natürlich auch für _diese_ Quellen.
  Ihre Verwendung in Programmen könnte zu SCHÄDEN führen, z. B. zum Abfackeln von Rechnern,
  zur Entgleisung von Eisenbahnen, zum GAU in Atomkraftwerken  oder zum Absturz des Mondes.
  Deshalb wird vor der Verwendung irgendwelcher Quellen von µU in Programmen zu ernsthaften
  Zwecken ausdrücklich gewarnt! (Ausgenommen sind Demo-Programme zum Einsatz in der Lehre.)
  Meldungen entdeckter Fehler und Hinweise auf Unklarheiten werden sehr dankbar angenommen.
/*/
import (
  "µU/achan"; "µU/audio"; "µU/barb"; "µU/barr"; "µU/book"; "µU/bpqu"; "µU/br"; "µU/bytes"
  "µU/car"; "µU/cdrom"; "µU/char"; "µU/col"; "µU/collop"; "µU/comp"; "µU/date"; "µU/day"
  "µU/dgra"; "µU/dlock"; "µU/env"; "µU/errh"; "µU/fig2"; "µU/fig3"; "µU/files"; "µU/gram"
  "µU/ieee"; "µU/kbd"; "µU/li"; "µU/lock"; "µU/lock2"; "µU/lockn"; "µU/lr"; "µU/macc"
  "µU/masks"; "µU/mbbuf"; "µU/mbuf"; "µU/mcorn"; "µU/menue"; "µU/mol"; "µU/mstk"
  "µU/pbar"; "µU/pat"; "µU/piset"; "µU/pos"; "µU/ppm"; "µU/pstk"; "µU/qmat"; "µU/reg"
  "µU/rpc"; "µU/rn"; "µU/rw"; "µU/scale"; "µU/schan"; "µU/scr"; "µU/smok"; "µU/term"
  "µU/texts"; "µU/time"; "µU/tval"; "µU/vnset"; "µU/Z"
)
const (
  yy = 2024
  mm =    2
  dd =   19
)
var (
  red, green = col.FlashRed(), col.FlashGreen()
  wd, ht, wd1, ht1, wdtext int
)

func circ (x int, c col.Colour) {
  scr.ColourF (c)
  scr.Circle (x, ht / 2, uint(ht) / 2 - 1)
}

func dr (x0, x1, y int, c col.Colour, f bool) {
  const dx = 2
  y1 := 0
  time.Msleep (100)
  for x := x0; x < x1; x += dx {
    scr.SaveGr (x, y, car.W, car.H)
    car.Draw (true, c, x, y)
    time.Msleep (20)
    scr.RestoreGr (x, y, car.W, car.H)
    if f && x > x0 + 46 * wd1 && x % 8 == 0 && y + 2 * car.H < ht {
      y1++
      y += y1
      if y + car.H >= ht { return }
    }
  }
}

func moon (x int, c col.Colour) {
  const r = 40
  y, y1 := r, 0
  for y < int(scr.Ht()) - r {
    scr.SaveGr (x - r, y - r, 2 * r, 2 * r)
    scr.ColourF (c)
    scr.CircleFull (x, y, r)
    scr.Flush()
    time.Msleep (33)
    scr.RestoreGr (x - r, y - r, 2 * r, 2 * r)
    y1++
    y += y1
  }
}

func joke (x, x1, y, imx, imy, imw int, c col.Colour, s string) {
  x2 := x + imx * wd1
  y1, y2 := imy * ht1, (imy + 13) * ht1
  _, my := scr.MaxRes()
  a := int(scr.NLines() - my / scr.Ht1() / 2) / 2
  y1 += a * ht1
  y2 += a * ht1
  switch s {
  case "nsp4", "nspe", "obp":
    y1 +=  4 * ht1
    y2 += 16 * ht1
  }
  y11 := y1
  dr (x, x2, y + imy * ht1, c, false)
  switch s {
  case "fire":
    y11 -= 2 * ht1
  case "mca":
    y11 -= 7 * ht1
  }
  scr.SaveGr (x2 - 4, y11, uint(x + imx * wd1 + imw * wd1 - x2 + 4), uint(y2 - y11))
  image := ppm.New()
  image.Load (s)
  scr.WriteImage (image.Colours(), x2 - 4, y11 + ht1)
  time.Sleep (uint(imw) / 6)
  scr.RestoreGr (x2 - 4, y11, uint(x + imx * wd1 + imw * wd1 - x2 + 4), uint(y2 - y11))
  dr (x2 + imw * wd1, x1, y + imy * ht1, c, false)
}

func drive (cl, cf, cb col.Colour, d chan int) {
  x := (wd - wdtext) / 2
  y := ((int(scr.NLines()) - 34) / 2 + 3) * ht1
  x1 := x + wdtext - car.W
  dr (x, x1, y +  0 * ht1, cl, false)
  dr (x, x1, y +  2 * ht1, cf, false)
  dr (x, x1, y +  3 * ht1, cf, false)
  joke (x, x1, y,  1, 4, 32, cf, "nsp4")
  joke (x, x1, y, 12, 5, 32, cf, "nspe")
  joke (x, x1, y, 30, 6, 32, cf, "obp")
  dr (x, x1, y +  7 * ht1, cf, false)
  dr (x, x1, y +  8 * ht1, cf, false)
  dr (x, x1, y +  9 * ht1, cf, false)
  dr (x, x1, y + 22 * ht1, cf, false)
  dr (x, x + 42 * wd1, y + 23 * ht1, cf, false)
  dr (x + 43 * wd1, wd + 33 * wd1, y + 23 * ht1, red, true)
  joke (x, x1, y, 67, 24, 14, cf, "fire")
  joke (x, x1, y, 38, 25, 22, cf, "mca")
  moon (x + 85 * wd1, col.LightGray())
  dr (x, x1, y + 26 * ht1, cf, false)
  dr (x, x1, y + 27 * ht1, cf, false)
  dr (x, x1, y + 29 * ht1, col.FlashYellow(), false)
  d <- 0
}

func input() { for { _, _ = kbd.Command() } }

func main() {
  scr.NewWH (0, 0, 1250, 800); defer scr.Fin()
  scr.Name (string(char.Mu)[1:] + "U")
  wd, ht = int(scr.Wd()), int(scr.Ht())
  wd1, ht1 = int(scr.Wd1()), int(scr.Ht1())
  achan.New(0); audio.New(); barb.NewDir(); barr.New(2); book.New(); bpqu.New(0, 1); br.New(3)
  bytes.Touch(); cdrom.Touch(); char.Touch(); collop.Touch(); comp.Touch(); date.New()
  dgra.Touch(); dlock.New(0, nil, 0); fig2.Touch(); fig3.Touch(); gram.Touch(); ieee.New()
  li.New(0); lock.NewChannel(); lock2.NewPeterson(); lockn.NewDijkstra(0); lr.NewMutex()
  macc.New(); masks.New(); mbbuf.New(nil, 2); mbuf.New(0); mcorn.New(0); menue.Touch()
  mol.New(); mstk.New(0); pbar.Touch(); pat.New(); piset.Touch(); pos.Touch(); pstk.Touch()
  qmat.Touch(); reg.Touch(); rpc.Touch(); rn.New0(); rw.New1(); scale.Touch(); schan.New(0)
  smok.Touch(); term.Touch(); texts.Touch(); tval.New(); vnset.EmptySet(); Z.String(0)
  var v day.Calendarday = day.New()
  v.Set (dd, mm, yy); v.SetFormat (day.Yymmdd)
  wdtext = 91 * wd1 // 91 == width of license text lines + 2
  files.Cd (env.Gosrc() + "/µU")
  go input()
  cl, cf, cb := col.FlashWhite(), col.LightGreen(), col.DarkGreen()
  circ (ht / 2, cf); circ (wd - ht / 2, cl)
  errh.MuLicense ("µU", v.String(),
                  "1986-2023  Christian Maurer   https://maurer-berlin.eu/mU", cl, cf, cb)
  scr.ScrColourB (cb)
  done := make(chan int)
  go drive (cl, cf, cb, done)
  <-done
}
