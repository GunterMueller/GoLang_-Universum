package main

/* (c) 1986-2016  murus.org
       dr-maurer.eu proprietary - all rights reserved

  Die Quellen von murus sind lediglich zum Einsatz in der Lehre konstruiert und haben demzufolge
  einen rein akademischen Charakter; sie liefern u.a. eine Reihe von Beispielen für das Lehrbuch
  "Nichtsequentielle Programmierung mit Go 1 kompakt" (Springer, 2. Aufl. 2012, 223 S. 14 Abb.).
  Für Lehrzwecke an Universitäten und in Schulen sind die Quelltexte uneingeschränkt verwendbar;
  jegliche Form weitergehender (insbesondere kommerzieller) Nutzung ist jedoch strikt untersagt.
  Davon abweichende Bedingungen sind der schriftlichen Vereinbarung mit dem Urheber vorbehalten.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER AND THE CONTRIBUTORS "AS IS" AND ANY EXPRESS
  OR IMPLIED WARRANTIES, INCLUDING BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
  AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE COPYRIGHT OWNER OR
  ANY CONTRIBUTOR BE LIABLE  FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSE-
  QUENTIAL DAMAGES  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)  HOWEVER CAUSED  AND ON ANY THEORY OF
  LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT  (INCLUDING NEGLIGENCE OR OTHERWISE)
  ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH
  DAMAGE. APART FROM THIS THE TEXT IN GERMAN ABOVE AND BELOW IS A MANDATORY PART OF THE LICENSE.

  Die Quelltexte von murus sind mit größter Sorgfalt entwickelt und werden laufend überarbeitet.
  ABER: Es gibt keine fehlerfreie Software - dies gilt natürlich auch für die Quellen von murus.
  Ihre Verwendung in Programmen könnte zu SCHÄDEN führen, z. B. zur Inbrandsetzung von Rechnern,
  zur Entgleisung von Eisenbahnzügen, zum GAU in Atomkraftwerken oder zum Absturz des Mondes ...
  Deshalb wird vor der Einbindung irgendwelcher Quelltexte von murus in Programme zu ernsthaften
  Zwecken AUSDRÜCKLICH GEWARNT ! (Ausgenommen sind nur Demo-Programme zum Einsatz in der Lehre.)

  Meldungen entdeckter Fehler und Hinweise auf Unklarheiten werden jederzeit dankbar angenommen. */

import (
  "murus/env"
  "murus/ker"; . "murus/obj"; "murus/sort"; "murus/cdrom";
  . "murus/mode"
  "murus/kbd"
  "murus/col"; "murus/scr"; "murus/errh"; "murus/scale"; "murus/pbar"
  "murus/files"
  "murus/integ"; "murus/lint"; "murus/brat"; "murus/real"
  "murus/stk"; "murus/buf"; "murus/bpqu"
  "murus/menue"
  "murus/date"; "murus/fuday"
  "murus/img"
  "murus/fig2"
  "murus/piset"
  "murus/persaddr"
  "murus/pset"
  "murus/schol"
  "murus/gram"
  "murus/fig"
  "murus/eye"
  "murus/audio"
  "murus/v"
  "murus/car"
  "murus/chanm"
  "murus/lock"
  "murus/asem"; "murus/barr"
  "murus/rw"; "murus/lr"
  "murus/lockp"
  "murus/mstk"; "murus/mqu"; "murus/mbuf"
  "murus/macc"
  "murus/nchan"
  "murus/naddr"
  "murus/dlock"
  "murus/lan"; "murus/lans"
  "murus/ntop"
  "murus/nelect"
  "murus/ntrav"
)
var
  Scr scr.Screen

func circ (c col.Colour, x, y int) {
  Scr.ColourF (c)
  Scr.Circle (x, y, uint(y) - 1)
}

func dr (x0, x1, y int, c col.Colour, f bool) {
  const dx = 2
  nx1, ny, y1 := int(Scr.Wd1()), int(Scr.Ht()), 0
  for x := x0; x < x1; x += dx {
    Scr.SaveGr (x, y, x + car.W, y + car.H)
    car.Draw (true, c, x, y)
    ker.Msleep (10)
    Scr.RestoreGr (x, y, x + car.W, y + car.H)
    if f && x > x0 + 26 * nx1 && x % 8 == 0 && y + car.H < ny {
      y1++
      y += y1
    }
  }
}

func moon (x0 int) {
  const r = 40
  x, y, y1, ny := x0, r, 0, int(Scr.Ht())
  for y < ny - r {
    Scr.SaveGr (x - r, y - r, x + r, y + r)
    Scr.ColourF (col.LightGray)
    Scr.CircleFull (x, y, r)
    Scr.Flush()
    ker.Msleep (33)
    Scr.RestoreGr (x - r, y - r, x + r, y + r)
    y1 ++
    y += y1
  }
}

func joke (x0, x1, y0, nx1, ny1, x, y, w int, cl col.Colour, s string, b bool) {
  x2 := x0 + x * nx1
  y1, y2, t := (y + 0) * ny1, (y + 13) * ny1, uint(1)
  a := int(Scr.NLines() - scr.MaxY() / Scr.Ht1() / 2) / 2 // fehlerdrumrum TODO
  y1 += a * ny1; y2 += a * ny1
  if b { y1 += 6 * ny1; y2 += 17 * ny1; t += 1 }
  dr (x0, x2, y0 + y * ny1, cl, false)
  Scr.SaveGr (x2, y1, x0 + x * nx1 + w * nx1, y2)
  img.Get (s, uint(x2), uint(y1))
  ker.Sleep (t)
  Scr.RestoreGr (x2, y1, x2 + w * nx1, y2)
  if b { w = 2 * w / 3 }
  dr (x2 + w * nx1, x1, y0 + y * ny1, cl, false)
}

func drive (cf, cl, cb col.Colour, d chan bool) {
  nx, nx1, ny1 := int(Scr.Wd()), int(Scr.Wd1()), int(Scr.Ht1())
  dw := 96 * nx1
  x0 := (nx - dw) / 2
  x1 := x0 + dw - car.W
  y0 := ((int(Scr.NLines()) - 31) / 2 + 3) * ny1 // 240
  dr (x0, x1, y0,            cf, false)
  dr (x0, x1, y0 +  2 * ny1, cl, false)
  dr (x0, x1, y0 +  3 * ny1, cl, false)
  joke (x0, x1, y0, nx1, ny1, 2, 4, 48, cl, "nsp", true)
  dr (x0, x1, y0 + 19 * ny1, cl, false)
  dr (x0, x0 + 70 * nx1, y0 + 20 * ny1, cl, false)
  b := Scr.ScrColB(); Scr.ScrColourB (col.Black)
  dr (x0 + 71 * nx1, nx, y0 + 20 * ny1, col.FlashRed, true)
  Scr.ScrColourB (b)
  joke (x0, x1, y0, nx1, ny1, 67, 21, 14, cl, "fire", false)
  joke (x0, x1, y0, nx1, ny1, 41, 22, 22, cl, "mca", false)
  moon (x0 + 90 * nx1)
  dr (x0, x1, y0 + 26 * ny1, cl, false)
  d <- true
}

func input() { for { _, _ = kbd.Command() } }

func main() { // just to get all stuff compiled
  if scr.UnderX() {
    xm, ym := scr.MaxRes()
    m := scr.MaxMode() - 3
    if m < WVGApp { m = WVGApp }
    x, y := Res(m)
    Scr = scr.New ((xm - x) / 2, (ym - y) / 2, m)
  } else {
    Scr = scr.NewMax()
  }
  defer Scr.Fin()
  files.Cd(env.Val("GOSRC"))
  files.Cd("murus")
  sort.Sort(make([]Any, 0))
  if cdrom.MaxVol == 0 {}
  scale.Lim(0,0,0,0,0)
  pbar.Touch()
  integ.String(0)
  lint.New(0)
  brat.New()
  real.String(0)
  stk.New(0)
  buf.New(nil, 0)
  bpqu.New(0, 1)
  menue.Touch()
  date.New()
  fuday.New()
  fig2.Touch()
  piset.Touch()
  pset.New(persaddr.New())
  schol.New()
  gram.Touch()
  fig.Touch()
  eye.New()
  audio.New()
  chanm.New()
  lock.NewMutex()
  asem.New(2)
  barr.New(2)
  rw.New()
  lr.New()
  lockp.NewPeterson()
  mstk.New(0)
  mqu.New(0)
  mbuf.New(0, 1)
  macc.New()
  naddr.New(nchan.Port0)
  lan.Touch()
  lans.Touch()
  dlock.New(0, nil, 0)
  ntop.Touch()
  nelect.Touch()
  ntrav.Touch()
  go input()
  x, y := int(Scr.Wd()), int(Scr.Ht()) / 2
  cf, cl, cb := v.Colours()
  circ (cb, x / 2, y); circ (cl, x - y, y); circ (cf, y, y)
  errh.MurusLicense ("murus", v.String(), "1986-2016  Christian Maurer   http://murus.org", cf, cl, cb)
  Scr.ScrColourB (cb); done := make (chan bool); go drive (cf, cl, cb, done); <-done
}
