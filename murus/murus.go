package main

/* (c) 1986-2017  Christian Maurer
       dr.maurer-berlin.eu proprietary - all rights reserved

  Die Quellen von murus sind nur zum Einsatz in der Lehre konstruiert  und haben demzufolge
  einen rein akademischen Charakter;  sie liefern u.a. etliche Beispiele  für mein Lehrbuch
  "Nichtsequentielle Programmierung mit Go 1 kompakt"  (Springer, 2. Auflage 2012, 223 S.).
  Für Lehrzwecke in Universitäten und Schulen  sind die Quellen uneingeschränkt verwendbar;
  jegliche weitergehende - insbesondere kommerzielle - Nutzung ist jedoch strikt untersagt.

  THIS SOFTWARE IS PROVIDED BY the authors  "AS IS"  AND ANY EXPRESS OR IMPLIED WARRANTIES,
  INCLUDING,  BUT NOT LIMITED TO,  THE IMPLIED WARRANTIES  OF MERCHANTABILITY  AND  FITNESS
  FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL the authors BE LIABLE FOR ANY
  DIRECT, INDIRECT,  INCIDENTAL, SPECIAL,  EXEMPLARY, OR CONSEQUENTIAL DAMAGES  (INCLUDING,
  BUT NOT LIMITED TO,  PROCUREMENT OF SUBSTITUTE GOODS  OR SERVICES;  LOSS OF USE, DATA, OR
  PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER
  IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY
  WAY OUT OF THE USE OF THIS SOFTWARE,  EVEN IF ADVISED  OF THE POSSIBILITY OF SUCH DAMAGE.

  APART FROM THIS  THE TEXT IN GERMAN ABOVE AND BELOW  IS A MANDATORY PART  OF THE LICENSE.

  Die Quelltexte von murus sind sehr sorgfältig entwickelt und werden laufend überarbeitet.
  ABER: Es gibt keine fehlerfreie Software - dies gilt natürlich auch für diese Quelltexte.
  Ihre Verwendung in Programmen könnte zu SCHÄDEN führen, z. B. zum Abfackeln von Rechnern,
  zur Entgleisung von Eisenbahnen, zum GAU in Atomkraftwerken  oder zum Absturz des Mondes.
  Deshalb wird  vor dem Einbau irgendwelcher Quellen von murus  in Programme zu ernsthaften
  Zwecken AUSDRÜCKLICH GEWARNT! (Ausgenommen sind Demo-Programme zum Einsatz in der Lehre.)

  Meldungen entdeckter Fehler und Hinweise auf Unklarheiten werden sehr dankbar angenommen. */

import (
  "murus/env"
  "murus/ker"
  . "murus/obj"
  "murus/sort"
  "murus/cdrom";
  . "murus/mode"
  "murus/kbd"
  "murus/col"
  "murus/scr"
  "murus/gl"
  "murus/errh"
  "murus/scale"
  "murus/pbar"
  "murus/files"
  "murus/integ"
  "murus/lint"
  "murus/brat"
  "murus/real"
  "murus/stk"
  "murus/buf"
  "murus/bpqu"
  "murus/menue"
  "murus/date"
  "murus/fuday"
  "murus/img"
  "murus/fig"
  "murus/piset"
  "murus/persaddr"
  "murus/pset"
  "murus/schol"
  "murus/gram"
  "murus/audio"
  "murus/v"
  "murus/car"
  "murus/chanm"
  "murus/lock"
  "murus/asem"
  "murus/barr"
  "murus/rw"
  "murus/lr"
  "murus/lockp"
  "murus/phil"
  "murus/smok"
  "murus/barb"
  "murus/mstk"
  "murus/mqu"
  "murus/mbuf"
  "murus/macc"
  "murus/nchan"
  "murus/naddr"
  "murus/dlock"
  "murus/dgra"
  "murus/dgras"
  "murus/vnset"
)
var
  screen scr.Screen

func circ (c col.Colour, x, y int) {
  screen.ColourF (c)
  screen.Circle (x, y, uint(y) - 1)
}

func dr (x0, x1, y int, c col.Colour, f bool) {
  const dx = 2
  nx1, ny, y1 := int(screen.Wd1()), int(screen.Ht()), 0
  for x := x0; x < x1; x += dx {
    screen.SaveGr (x, y, x + car.W, y + car.H)
    car.Draw (true, c, x, y)
    ker.Msleep (10)
    screen.RestoreGr (x, y, x + car.W, y + car.H)
    if f && x > x0 + 47 * nx1 && x % 8 == 0 && y + car.H < ny {
      y1++
      y += y1
    }
  }
}

func moon (x0 int) {
  const r = 40
  x, y, y1, ny := x0, r, 0, int(screen.Ht())
  for y < ny - r {
    screen.SaveGr (x - r, y - r, x + r, y + r)
    screen.ColourF (col.LightGray)
    screen.CircleFull (x, y, r)
    screen.Flush()
    ker.Msleep (33)
    screen.RestoreGr (x - r, y - r, x + r, y + r)
    y1 ++
    y += y1
  }
}

func joke (x0, x1, y0, nx1, ny1, x, y, w int, cl col.Colour, s string, b bool) {
  x2 := x0 + x * nx1
  y1, y2, t := (y + 0) * ny1, (y + 13) * ny1, uint(1)
  a := int(screen.NLines() - scr.MaxY() / screen.Ht1() / 2) / 2 // fehlerdrumrum TODO
  y1 += a * ny1; y2 += a * ny1
  if b { y1 += 6 * ny1; y2 += 17 * ny1; t += 1 }
  dr (x0, x2, y0 + y * ny1, cl, false)
  screen.SaveGr (x2 - 4, y1, x0 + x * nx1 + w * nx1, y2)
  img.Get (s, uint(x2) - 4, uint(y1))
  ker.Sleep (t)
  screen.RestoreGr (x2 - 4, y1, x2 + w * nx1, y2)
  if b { w = 2 * w / 3 }
  dr (x2 + w * nx1, x1, y0 + y * ny1, cl, false)
}

func drive (cf, cl, cb col.Colour, d chan bool) {
  nx, nx1, ny1 := int(screen.Wd()), int(screen.Wd1()), int(screen.Ht1())
  dw := 91 * nx1
  x0 := (nx - dw) / 2
  x1 := x0 + dw - car.W
  y0 := ((int(screen.NLines()) - 30) / 2 + 3) * ny1
  dr (x0, x1, y0,            cf, false)
  dr (x0, x1, y0 +  2 * ny1, cl, false)
  dr (x0, x1, y0 +  3 * ny1, cl, false)
  joke (x0, x1, y0, nx1, ny1, 2, 5, 32, cl, "nsp", true)
  dr (x0, x1, y0 + 18 * ny1, cl, false)
  dr (x0, x0 + 42 * nx1, y0 + 19 * ny1, cl, false)
  b := screen.ScrColB(); screen.ScrColourB (col.Black)
  dr (x0 + 43 * nx1, nx + 31 * nx1, y0 + 19 * ny1, col.FlashRed, true)
  screen.ScrColourB (b)
  joke (x0, x1, y0, nx1, ny1, 67, 20, 14, cl, "fire", false)
  joke (x0, x1, y0, nx1, ny1, 38, 21, 22, cl, "mca", false)
  moon (x0 + 85 * nx1)
  dr (x0, x1, y0 + 25 * ny1, cl, false)
  d <- true
}

func input() { for { _, _ = kbd.Command() } }

func main() { // just to get all stuff compiled
  if scr.UnderX() {
    xm, ym := scr.MaxRes()
    m := scr.MaxMode() - 3
    if m < XGA { m = XGA }
    x, y := Res(m)
    screen = scr.New ((xm - x) / 2, (ym - y) / 2, m)
  } else {
    screen = scr.NewMax()
  }
  defer screen.Fin()
  env.Call()
  files.Cd(ker.Murus)
  sort.Sort(make([]Any, 0))
  gl.Touch()
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
  fig.Touch()
  piset.Touch()
  pset.New(persaddr.New())
  schol.New()
  gram.Touch()
  audio.New()
  chanm.New()
  lock.NewMutex()
  asem.New(2)
  barr.New(2)
  rw.New()
  lr.New()
  lockp.NewPeterson()
  phil.NewNaive()
  smok.NewNaive()
  barb.NewDir()
  mstk.New(0)
  mqu.New(0)
  mbuf.New(0, 1)
  macc.New()
  naddr.New(nchan.Port0)
  dgra.Touch()
  dgras.Touch()
  dlock.New(0, nil, 0)
  vnset.EmptySet()
  go input()
  x, y := int(screen.Wd()), int(screen.Ht()) / 2
  cf, cl, cb := v.Colours()
  circ (cb, x / 2, y); circ (cl, x - y, y); circ (cf, y, y)
  errh.MurusLicense ("murus", v.String(),
                     "1986-2017  Christian Maurer   https://maurer-berlin.eu/murus", cf, cl, cb)
  screen.ScrColourB (cb)
  done := make(chan bool)
  go drive (cf, cl, cb, done)
  <-done
}
