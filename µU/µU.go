package main

/* (c) 1986-2021  Christian Maurer       maurer-berlin.eu proprietary - all rights reserved

  Das Mikrouniversum µU ist nur zum Einsatz in der Lehre konstruiert  und hat deshalb einen
  rein akademischen Charakter.  Es liefert u.a. eine Reihe von Beispielen für mein Lehrbuch
  "Nichtsequentielle und Verteilte Programmierung mit Go" (Springer Vieweg 2019) und dessen
  Übersetzung  "Nonsequential and Distributed Programming with Go"  (Springer Nature 2021).
  Für Zwecke der Lehre an Universitäten und in Schulen sind die Quellen des Mikrouniversums
  uneingeschränkt verwendbar; jede Form weitergehender Nutzung ist jedoch strikt untersagt.

  THIS SOFTWARE IS PROVIDED BY the authors  "AS IS"  AND ANY EXPRESS OR IMPLIED WARRANTIES,
  INCLUDING,  BUT NOT LIMITED TO,  THE IMPLIED WARRANTIES  OF MERCHANTABILITY  AND  FITNESS
  FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY
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
  Zwecken AUSDRÜCKLICH GEWARNT! (Ausgenommen sind Demo-Programme zum Einsatz in der Lehre.) */

import (
  "µU/achan"; "µU/audio"; "µU/barb"; "µU/barr"; "µU/bbuf"; "µU/book"; "µU/bpqu"; "µU/br"
  "µU/bytes"; "µU/car"; "µU/cdrom"; "µU/col"; "µU/collop"; "µU/cntry"; "µU/comp"; "µU/day"
  "µU/date"; "µU/dgra"; "µU/dlock"; "µU/env"; "µU/errh"; "µU/euro"; "µU/f"; "µU/fig2"
  "µU/files"; "µU/fuday"; "µU/gram"; "µU/host"; "µU/ieee"; "µU/kbd"; "µU/li"; "µU/lock2"
  "µU/lr"; "µU/macc"; "µU/mbbuf"; "µU/mbuf"; "µU/mcorn"; "µU/menue"; "µU/mstk"; "µU/pbar"
  "µU/persaddr"; "µU/phil"; "µU/piset"; "µU/pos"; "µU/ppm"; "µU/q"; "µU/r"; "µU/reg"
  "µU/rpc"; "µU/rw"; "µU/scale"; "µU/schan"; "µU/scr"; "µU/smok"; "µU/term"; "µU/texts"
  "µU/time"; "µU/vnset"; "µU/z"
)
const (
  yy = 2021
  mm =    7
  dd =   23
)
var (
  red = col.FlashRed()
  green = col.FlashGreen()
  nx, nx1, ny1, wdtext, wd, ht int
)

func circ (x int, c col.Colour) {
  scr.ColourF (c)
  scr.Circle (x, ht / 2, uint(ht) / 2 - 1)
}

func dr (x0, x1, y int, c col.Colour, f bool) {
  const dx = 2
  y1 := 0
  for x := x0; x < x1; x += dx {
    scr.SaveGr (x, y, x + car.W, y + car.H)
    car.Draw (true, c, x, y)
    time.Msleep (20)
    scr.RestoreGr (x, y, x + car.W, y + car.H)
    if f && x > x0 + 46 * nx1 && x % 8 == 0 && y + 2 * car.H < ht {
      y1++
      y += y1
    }
  }
}

func moon (x int, c col.Colour) {
  const r = 40
  y, y1 := r, 0
  for y < int(scr.Ht()) - r {
    scr.SaveGr (x - r, y - r, x + r, y + r)
    scr.ColourF (c)
    scr.CircleFull (x, y, r)
    scr.Flush()
    time.Msleep (33)
    scr.RestoreGr (x - r, y - r, x + r, y + r)
    y1++
    y += y1
  }
}

func joke (x, x1, y, imx, imy, imw int, c col.Colour, s string) {
  x2 := x + imx * nx1
  y1, y2 := imy * ny1, (imy + 13) * ny1
  _, my := scr.MaxRes()
  a := int(scr.NLines() - my / scr.Ht1() / 2) / 2
  y1 += a * ny1
  y2 += a * ny1
  switch s {
  case "nsp2", "nsp4", "nspe":
    y1 += 5 * ny1;
    y2 += 17 * ny1;
  }
  y11 := y1
  dr (x, x2, y + imy * ny1, c, false)
  switch s {
  case "fire":
    y11 -= 1 * ny1
  case "mca":
    y11 -= 6 * ny1
  }
  scr.SaveGr (x2 - 4, y11, x + imx * nx1 + imw * nx1, y2)
  ppm.Get (s, uint(x2) - 4, uint(y11))
  time.Sleep (uint(imw) / 6)
  scr.RestoreGr (x2 - 4, y11, x2 + imw * nx1, y2)
  dr (x2 + imw * nx1, x1, y + imy * ny1, c, false)
}

func drive (cl, cf, cb col.Colour, d chan bool) {
  x := (nx - wdtext) / 2
  y := ((int(scr.NLines()) - 32) / 2 + 3) * ny1
  x1 := x + wdtext - car.W
  dr (x, x1, y +  0 * ny1, cl, false)
  dr (x, x1, y +  2 * ny1, cf, false)
  dr (x, x1, y +  3 * ny1, cf, false)
  joke (x, x1, y, 14, 4, 23, cf, "nsp2")
  joke (x, x1, y, 18, 5, 32, cf, "nsp4")
  joke (x, x1, y, 16, 6, 32, cf, "nspe")
  dr (x, x1, y +  7 * ny1, cf, false)
  dr (x, x1, y +  8 * ny1, cf, false)
  dr (x, x1, y + 21 * ny1, cf, false)
  dr (x, x + 42 * nx1, y + 22 * ny1, cf, false)
  dr (x + 43 * nx1, nx + 33 * nx1, y + 22 * ny1, red, true)
  joke (x, x1, y, 67, 23, 14, cf, "fire")
  joke (x, x1, y, 38, 24, 22, cf, "mca")
  moon (x + 85 * nx1, col.LightGray())
  dr (x, x1, y + 25 * ny1, cf, false)
  dr (x, x1, y + 26 * ny1, cf, false)
  dr (x, x1, y + 28 * ny1, green, false)
  d <- true
}

func input() { for { _, _ = kbd.Command() } }

func main() { // get all packages compiled and show the license
  if env.UnderX() {
    xm, ym := scr.MaxRes()
    ht = int(ym) - 56
    scr.NewWH (0, 0, xm, uint(ht))
  } else {
    scr.NewMax()
    ht = int(scr.Ht())
  }
  defer scr.Fin()
  achan.New(0); audio.New(); barb.NewDir(); barr.New(2); bbuf.New(nil, 0); book.New()
  bpqu.New(0, 1); br.New(1); bytes.Touch(); if cdrom.MaxVol == 0 {}; cntry.New()
  collop.Touch(); comp.Touch(); date.New(); euro.New(); dgra.Touch(); dlock.New(0, nil, 0)
  f.Touch(); fig2.Touch(); fuday.New(); gram.Touch(); host.New(); ieee.New(); li.New(0)
  lock2.NewPeterson(); lr.NewMutex(); macc.New(); mbbuf.New(nil, 1); mbuf.New(0)
  mcorn.New(0); menue.Touch(); mstk.New(0); pbar.Touch(); persaddr.New(); phil.TouchPhil()
  piset.Touch(); pos.Touch(); q.New(); r.String(0); reg.Touch(); rpc.Touch(); rw.New1()
  scale.Lim(0,0,0,0,0); schan.New(0); smok.TouchSmok(); term.New(""); texts.Touch()
  vnset.EmptySet(); z.String(0)
  var v day.Calendarday = day.New()
  v.Set (dd, mm, yy)
  v.SetFormat (day.Yymmdd)
  wd = int(scr.Wd())
  nx = int(scr.Wd())
  nx1, ny1 = int(scr.Wd1()), int(scr.Ht1())
  wdtext = 91 * nx1 // 91 == width of license text lines + 2
  files.Cd (env.Gosrc() + "/µU")
  go input()
  cl, cf, cb := col.LightWhite(), col.LightGreen(), col.BlackGreen()
  circ (ht / 2, cf);
  circ (wd - ht / 2, cl)
  errh.MuLicense ("µU", v.String(), "1986-2021  Christian Maurer   https://maurer-berlin.eu/mU", cl, cf, cb)
  scr.ScrColourB (cb)
  done := make(chan bool)
  go drive (cl, cf, cb, done)
  <-done
}
