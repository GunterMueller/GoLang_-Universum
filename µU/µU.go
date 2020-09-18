package main

/* (c) 1986-2020  Christian Maurer
       christian.maurer-berlin.eu proprietary - all rights reserved

  Das Mikrouniversum µU ist nur zum Einsatz in der Lehre konstruiert  und hat deshalb einen
  rein akademischen Charakter.  Es liefert u.a. eine Reihe von Beispielen für mein Lehrbuch
  "Nichtsequentielle und Verteilte Programmierung mit Go"  (Springer Vieweg 2018 und 2019).
  Für Zwecke der Lehre an Universitäten und in Schulen sind die Quellen des Mikrouniversums
  uneingeschränkt verwendbar; jede Form weitergehender Nutzung ist jedoch strikt untersagt.

  THIS SOFTWARE IS PROVIDED BY the authors  "AS IS"  AND ANY EXPRESS OR IMPLIED WARRANTIES,
  INCLUDING,  BUT NOT LIMITED TO,  THE IMPLIED WARRANTIES  OF MERCHANTABILITY  AND  FITNESS
  FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL the authors BE LIABLE FOR ANY
  DIRECT, INDIRECT,  INCIDENTAL, SPECIAL,  EXEMPLARY, OR CONSEQUENTIAL DAMAGES  (INCLUDING,
  BUT NOT LIMITED TO,  PROCUREMENT OF SUBSTITUTE GOODS  OR SERVICES;  LOSS OF USE, DATA, OR
  PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER
  IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY
  WAY OUT OF THE USE OF THIS SOFTWARE,  EVEN IF ADVISED  OF THE POSSIBILITY OF SUCH DAMAGE.

  APART FROM THIS  THE TEXT IN GERMAN ABOVE AND BELOW  IS A MANDATORY PART  OF THE LICENSE.

  Die Quelltexte von µU sind äußerst sorgfältig entwickelt und werden laufend überarbeitet.
  ABER:  Es gibt keine fehlerfreie Software - dies gilt natürlich auch für _diese_ Quellen.
  Ihre Verwendung in Programmen könnte zu SCHÄDEN führen, z. B. zum Abfackeln von Rechnern,
  zur Entgleisung von Eisenbahnen, zum GAU in Atomkraftwerken  oder zum Absturz des Mondes.
  Deshalb wird vor der Verwendung irgendwelcher Quellen von µU in Programmen zu ernsthaften
  Zwecken AUSDRÜCKLICH GEWARNT! (Ausgenommen sind Demo-Programme zum Einsatz in der Lehre.)

  Meldungen entdeckter Fehler und Hinweise auf Unklarheiten werden sehr dankbar angenommen. */

import (
  "µU/env"
  "µU/ker"
  "µU/time"
  . "µU/obj"
  "µU/sort"
  "µU/cdrom";
//  "µU/mode"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/vect"
  "µU/gl"
  "µU/errh"
  "µU/scale"
  "µU/pbar"
  "µU/files"
  "µU/integ"
  "µU/lint"
  "µU/brat"
  "µU/real"
  "µU/stk"
  "µU/buf"
  "µU/bbuf"
  "µU/bpqu"
  "µU/menue"
  "µU/date"
  "µU/fuday"
  "µU/img"
  "µU/fig2"
  "µU/piset"
  "µU/persaddr"
  "µU/pset"
//  "µU/words"
  "µU/schol"
  "µU/gram"
  "µU/audio"
  "µU/v"
  "µU/car"
  "µU/schan"
  "µU/achan"
  "µU/lock"
  "µU/asem"
  "µU/barr"
  "µU/rw"
  "µU/lr"
  "µU/lock2"
  "µU/lockn"
  "µU/phil"
  "µU/smok"
  "µU/barb"
  "µU/mstk"
  "µU/mbuf"
  "µU/mbbuf"
  "µU/macc"
  "µU/nchan"
  "µU/naddr"
  "µU/dlock"
  "µU/dgra"
  "µU/rpc"
  "µU/vnset"
)
var (
  red = col.FlashRed()
  screen scr.Screen
  nx, nx1, ny1, wdtext, wd, ht int
)

func circ (x int, c col.Colour) {
  screen.ColourF (c)
  screen.Circle (x, ht / 2, uint(ht) / 2 - 1)
}

func dr (x0, x1, y int, c col.Colour, f bool) {
  const dx = 2
  y1 := 0
  for x := x0; x < x1; x += dx {
    screen.SaveGr (x, y, x + car.W, y + car.H)
    car.Draw (true, c, x, y)
    time.Msleep (10)
    screen.RestoreGr (x, y, x + car.W, y + car.H)
    if f && x > x0 + 46 * nx1 && x % 8 == 0 && y + 2 * car.H < ht {
      y1++
      y += y1
    }
  }
}

func moon (x int, c col.Colour) {
  const r = 40
  y, y1 := r, 0
  for y < int(screen.Ht()) - r {
    screen.SaveGr (x - r, y - r, x + r, y + r)
    screen.ColourF (c)
    screen.CircleFull (x, y, r)
    screen.Flush()
    time.Msleep (33)
    screen.RestoreGr (x - r, y - r, x + r, y + r)
    y1++
    y += y1
  }
}

func joke (x, x1, y, imx, imy, imw int, c col.Colour, s string) {
  x2 := x + imx * nx1
  y1, y2 := imy * ny1, (imy + 13) * ny1
  a := int(screen.NLines() - scr.MaxY() / screen.Ht1() / 2) / 2
  y1 += a * ny1
  y2 += a * ny1
  if s == "nsp4" {
    y1 += 6 * ny1;
    y2 += 17 * ny1;
  }
  y11 := y1
  dr (x, x2, y + imy * ny1, c, false)
  if s == "mca" {
    y11 -= 5 * ny1
  }
  screen.SaveGr (x2 - 4, y11, x + imx * nx1 + imw * nx1, y2)
  img.Get (s, uint(x2) - 4, uint(y11))
  time.Sleep (uint(imw) / 6)
  screen.RestoreGr (x2 - 4, y11, x2 + imw * nx1, y2)
  dr (x2 + imw * nx1, x1, y + imy * ny1, c, false)
}

func drive (cl, cf, cb col.Colour, d chan bool) {
  x := (nx - wdtext) / 2
//  y := ((int(screen.NLines()) - 30) / 2 + 3) * ny1
  y := ((int(screen.NLines()) - 31) / 2 + 3) * ny1
  x1 := x + wdtext - car.W
  dr (x, x1, y +  0 * ny1, cl, false)
  dr (x, x1, y +  2 * ny1, cf, false)
  dr (x, x1, y +  3 * ny1, cf, false)
  joke (x, x1, y, 58, 4, 34, cf, "nsp4")
  dr (x, x1, y +  5 * ny1, cf, false)
  dr (x, x1, y +  6 * ny1, cf, false)
  dr (x, x1, y + 19 * ny1, cf, false)
  dr (x, x + 42 * nx1, y + 20 * ny1, cf, false)
  dr (x + 43 * nx1, nx + 31 * nx1, y + 20 * ny1, red, true)
  joke (x, x1, y, 67, 21, 14, cf, "fire")
  joke (x, x1, y, 38, 22, 22, cf, "mca")
  moon (x + 85 * nx1, col.LightGray())
  dr (x, x1, y + 23 * ny1, cf, false)
  dr (x, x1, y + 24 * ny1, cf, false)
  dr (x, x1, y + 26 * ny1, red, false)
  moon (x + 85 * nx1, red)
  d <- true
}

func input() { for { _, _ = kbd.Command() } }

func main() { // get all packages compiled and show the license
  if scr.UnderX() {
    xm, ym := scr.MaxRes()
    ht = int(ym) - 56
    screen = scr.NewWH (0, 0, xm, uint(ht))
  } else {
    screen = scr.NewMax()
    ht = int(screen.Ht())
  }
  wd = int(screen.Wd())
  defer screen.Fin()
  nx = int(screen.Wd())
  nx1, ny1 = int(screen.Wd1()), int(screen.Ht1())
  wdtext = 91 * nx1 // 91 == width of license text lines + 2
  files.Cd (env.Gosrc() + "/" + ker.Mu)
  sort.Sort(make([]Any, 0))
  vect.New()
  gl.Touch()
  if cdrom.MaxVol == 0 {}
  scale.Lim(0,0,0,0,0)
  pbar.Touch()
  integ.String(0)
  lint.New(0)
  brat.New()
  real.String(0)
  stk.New(0)
  buf.New(0)
  bbuf.New(nil, 0)
  bpqu.New(0, 1)
  menue.Touch()
  date.New()
  fuday.New()
  fig2.Touch()
  piset.Touch()
  pset.New(persaddr.New())
//  words.New(0, 0)
  schol.New()
  gram.Touch()
  audio.New()
  schan.New(0)
  achan.New(0)
  lock.NewMutex()
  asem.New(2)
  barr.New(2)
  rw.New1()
  lr.NewMutex()
  lock2.NewPeterson()
  lockn.NewTiebreaker(2)
  phil.TouchPhil()
  smok.TouchSmok()
  barb.NewDir()
  mstk.New(0)
  mbuf.New(0)
  mbbuf.New(nil, 1)
  macc.New()
  naddr.New(nchan.Port0)
  dgra.Touch()
  dlock.New(0, nil, 0)
  rpc.Touch()
  vnset.EmptySet()
  go input()
  cl, cf, cb := v.Colours()
  circ (ht / 2, cf);
  circ (wd - ht / 2, cl)
  errh.MuLicense (ker.Mu, v.String(), "1986-2020  Christian Maurer   https://maurer-berlin.eu/mu", cl, cf, cb)
  screen.ScrColourB (cb)
  done := make(chan bool)
  go drive (cl, cf, cb, done)
  <-done
}
