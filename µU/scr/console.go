package scr

// (c) Christian Maurer   v. 211222 - license see µU.go

//#include <stdlib.h>
//#include <fcntl.h>
//#include <unistd.h>
//#include <sys/ioctl.h>
//#include <sys/mman.h>
//#include <linux/vt.h>
//#include <linux/fb.h>
/*
void * framebuffer (int *x, int *y, int *b, int *a) {
  int fd;
  struct fb_var_screeninfo v;
  struct fb_fix_screeninfo f;
  struct vt_stat s;
  void *M = NULL;
  int offset;
  *x = 0;
  *y = 0;
  *b = 0;
  *a = 0;
  if ((fd = open ("/dev/fb0", O_RDWR)) == -1) { return NULL; }
  if (ioctl (fd, FBIOGET_VSCREENINFO, &v) == -1) { close (fd); return NULL; }
  *x = v.xres;
  *y = v.yres;
  *b = v.bits_per_pixel;
  if (ioctl (fd, FBIOGET_FSCREENINFO, &f) == -1) { close (fd); return NULL; }
  if (f.type != FB_TYPE_PACKED_PIXELS) { close (fd); return NULL; }
  if (ioctl (0, VT_GETSTATE, &s) == -1) { close (fd); return NULL; }
  *a = s.v_active;
  ioctl (0, VT_ACTIVATE, *a);
  ioctl (0, VT_WAITACTIVE, *a);
  offset = (unsigned long)(f.smem_start) & 4095UL;
  M = mmap (NULL, f.smem_len + offset, PROT_WRITE, MAP_SHARED, fd, 0);
  if ((long)M == -1L) { M = NULL; }
  close (fd);
  return M;
}
*/
import
  "C"
import (
	"reflect"
  "unsafe"
  "runtime"
  "syscall"
  "sync"
  "strconv"
  "math"
  "µU/ker"
  "µU/time"
  "µU/char"
  "µU/obj"
  "µU/font"
  "µU/col"
  "µU/mode"
  "µU/linewd"
  "µU/mouse"
  "µU/scr/shape"
  "µU/scr/ptr"
)
type
  console struct {
                 mode.Mode
            x, y int
          wd, ht uint
          nLines,
        nColumns uint
         archive obj.Stream
          shadow []obj.Stream
            buff bool
        wd1, ht1 uint
          cF, cB,
        cFA, cBA col.Colour
    codeF, codeB obj.Stream
      scrF, scrB col.Colour
          lineWd linewd.Linewidth
        fontsize font.Size
     transparent bool
     cursorShape,
    consoleShape,
      blinkShape shape.Shape
  blinkX, blinkY uint
      blinkMutex sync.Mutex
        blinking bool
         mouseOn bool
         pointer ptr.Pointer
  xMouse, yMouse int
         polygon [][]bool // to fill polygons
            done [][]bool // to fill polygons
        incident bool
   xx_, yy_, tt_ int // for incidence tests
       ppmheader string
              lh uint // length of ppmheader
                 }
var (
  actMutex sync.Mutex
  actualC, mouseConsole *console
  mouseIndex int
  width, height uint
  fullScreen mode.Mode
)

func (x *console) Fin() {
  ker.Fin()
}

func (x *console) Flush() {
}

func (x *console) Name (n string) {
  // TODO
}

func (x *console) ActMode() mode.Mode {
  return x.Mode
}

func (x *console) X() uint {
  return uint(x.x)
}

func (x *console) Y() uint {
  return uint(x.y)
}

func (x *console) Wd() uint {
  return x.wd
}

func (x *console) Ht() uint {
  return x.ht
}

func (x *console) Proportion() float64 {
  return float64(x.wd) / float64(x.ht)
}

func (X *console) ok() bool {
  return uint(X.x) + X.wd <= width && uint(X.y) + X.ht <= height
}

func (X *console) OnFocus() bool {
  return true
}

func (X *console) OffFocus() bool {
  return false
}

func (X *console) Win2buf() {
}

// colours /////////////////////////////////////////////////////////////

const
  colC = esc1 + "3%d;4%d"
var
  colourdepth uint

func (X *console) ScrColours (f, b col.Colour) {
  X.scrF = f
  X.ScrColourB (b)
}

func (X *console) ScrColourF (f col.Colour) {
  X.scrF = f
}

func (X *console) ScrColourB (b col.Colour) {
  X.scrB = b.Clone().(col.Colour)
  c := b.EncodeInv()
  a := uint(0)
  for x := uint(0); x < width; x++ {
    for y := uint(0); y < height; y++ {
      copy (emptyBackground[a:a+colourdepth], c)
      a += colourdepth
    }
  }
}

func (X *console) ScrCols() (col.Colour, col.Colour) {
  return X.scrF, X.scrB
}

func (X *console) StartCols() (col.Colour, col.Colour) {
  return X.cF, X.cB
}

func (X *console) StartColsA() (col.Colour, col.Colour) {
  return X.cFA, X.cBA
}

func (X *console) ScrColF() col.Colour {
  return X.scrF
}

func (X *console) ScrColB() col.Colour {
  return X.scrB
}

func (X *console) Colours (f, b col.Colour) {
  X.cF, X.codeF = f, f.EncodeInv()
  X.cB, X.codeB = b, b.EncodeInv()
}

func (X *console) ColourF (f col.Colour) {
  X.cF, X.codeF = f, f.EncodeInv()
}

func (X *console) ColourB (b col.Colour) {
  X.cB, X.codeB = b, b.EncodeInv()
}

func (X *console) Cols() (col.Colour, col.Colour) {
  return X.cF, X.cB
}

func (X *console) ColF() col.Colour {
  return X.cF
}

func (X *console) ColB() col.Colour {
  return X.cB
}

func (X *console) Colour (x, y uint) col.Colour {
  if x >= X.wd || y >= X.ht || ! visible {
    return X.scrB
  }
  x += uint(X.x)
  y += uint(X.y)
  i := int(width * y + x) * int(colourdepth)
  s := fbcop [i:i+int(colourdepth)]
  return col.New3 (s[0], s[1], s[2])
}

// ranges //////////////////////////////////////////////////////////////

func (X *console) rectangOk (x, y, x1, y1 *int) bool {
  if ! visible { return false }
  intord (x, y, x1, y1)
  if *x >= int(X.wd) || *y >= int(X.ht) {
    return false
  }
  if *x1 >= int(X.wd) { *x1 = int(X.wd) - 1 }
  if *y1 >= int(X.ht) { *y1 = int(X.ht) - 1 }
  return true
}

func (X *console) urectangOk (x, y, x1, y1 *uint) bool {
  if ! visible { return false }
  if *x > *x1 { *x, *x1 = *x1, *x }
  if *y > *y1 { *y, *y1 = *y1, *y }
  if *x >= X.wd || *y >= X.ht {
    return false
  }
  if *x1 >= X.wd { *x1 = X.wd - 1 }
  if *y1 >= X.ht { *y1 = X.ht - 1 }
  return true
}

func (X *console) Cls() {
  if ! visible { return }
  a := uint(0)
  c := X.scrB.EncodeInv()
  for j := uint(0); j < X.ht; j++ {
    for i := uint(0); i < X.wd; i++ {
      copy (emptyBackground[a:a+colourdepth], c)
      a += colourdepth
    }
  }
  a = (uint(X.y) * width + uint(X.x)) * colourdepth
  l := colourdepth * X.wd
  for j := uint(0); j < X.ht; j++ {
    copy (fbmem[a:a+l], emptyBackground)
    copy (fbcop[a:a+l], emptyBackground)
    a += width * colourdepth
  }
}

func u (n uint) string { return strconv.Itoa(int(n)) }
func i (n  int) string { return strconv.Itoa(    n ) }

func (X *console) Clr (l, c, w, h uint) {
  x, y := int(c * X.wd1), int(l * X.ht1)
  X.ClrGr (x, y, x + int(w * X.wd1), y + int(h * X.ht1))
}

func (X *console) ClrGr (x, y, x1, y1 int) {
  if ! X.rectangOk (&x, &y, &x1, &y1) { return }
  if ! visible { return }
  x += X.x; x1 += X.x // y's diff !
  da := uint(x1 - x) * colourdepth
  a := uint(0)
/*
  c := col.ColStream (X.ScrColB())
  for j := uint(0); j < da; j++ {
    copy (emptyBackground[a:a+colourdepth], c)
    a += colourdepth
  }
*/
  w := width * colourdepth
  a = uint(X.y + y) * w + uint(x) * colourdepth
  for z := 0; z <= y1 - y; z++ {
    copy (fbmem[a:a+da], emptyBackground[:da])
    copy (fbcop[a:a+da], emptyBackground[:da])
    a += w
  }
}

func (X *console) Buf (on bool) {
  if X.buff == on { return }
  X.buff = on
/*
  a := 0
  c := col.ColStream (X.ScrColB())
  for x := 0; x < int(X.wd); x++ {
    copy (emptyBackground[a:a+int(colourdepth)], c)
    a += int(colourdepth)
  }
*/
  da := int(X.wd) * int(colourdepth)
  w := int(width) * int(colourdepth)
  a := (int(width) * X.y + X.x) * int(colourdepth)
  for y := 0; y < int(X.ht); y++ {
    if on {
      copy (fbcop[a:a+da], emptyBackground[:da])
    } else {
      copy (fbmem[a:a+da], fbcop[a:a+da])
    }
    a += w
  }
/*/
  if on {
    copy (fbcop, emptyBackground)
  } else {
    copy (fbmem, fbcop)
  }
/*/
}

func (X *console) Buffered () bool {
  return X.buff
}

func (X *console) Save (l, c, w, h uint) {
  x, y := int(X.wd1 * c), int(X.ht1 * l)
  X.SaveGr (x, y, x + int(X.wd1 * w), y + int(X.ht1 * h))
}

func (X *console) SaveGr (x, y, x1, y1 int) {
  if ! X.rectangOk (&x, &y, &x1, &y1) { return }
  w, h := x1 - x + 1, y1 - y + 1
  x0, y0 := X.x + x, X.y + y
  a, da := x * int(colourdepth), w * int(colourdepth)
  if X.mouseOn { X.MousePointer (false) }
  for i := 0; i < h; i++ {
    b := (int(width) * (y0 + i) + x0) * int(colourdepth)
    copy (X.shadow[i][a:a+da], fbmem[b:b+da])
  }
  if X.mouseOn { X.MousePointer (true) }
}

func (X *console) Save1 () {
  X.SaveGr (0, 0, int(X.wd) - 1, int(X.ht) - 1)
}

func (X *console) Restore (l, c, w, h uint) {
  x, y := int(X.wd1 * c), int(X.ht1 * l)
  X.RestoreGr (x, y, x + int(X.wd1 * w), y + int(X.ht1 * h))
}

func (X *console) RestoreGr (x, y, x1, y1 int) {
  if ! X.rectangOk (&x, &y, &x1, &y1) { return }
  w, h := x1 - x + 1, y1 - y + 1
  x0, y0 := X.x + x, X.y + y
  a, da := x * int(colourdepth), w * int(colourdepth)
  for i := 0; i < h; i++ {
    b := (int(width) * (y0 + i) + x0) * int(colourdepth)
    copy (fbmem[b:b+da], X.shadow[i][a:a+da])
    copy (fbcop[b:b+da], X.shadow[i][a:a+da])
  }
}

func (X *console) Restore1() {
  X.RestoreGr (0, 0, int(X.wd) - 1, int(X.ht) - 1)
}

var
  finished bool

func (c *console) blink() {
  var s shape.Shape
  for {
    c.blinkMutex.Lock()
    if c.cursorShape == shape.Off {
      s = c.blinkShape
    } else {
      s = shape.Off
    }
    c.cursor (c.blinkX, c.blinkY, s)
    c.blinkMutex.Unlock()
    if finished {
      break
    }
    time.Msleep (250)
  }
  runtime.Goexit()
}

func (c *console) doBlink() {
  if c.blinking { return }
  c.blinking = true
  go c.blink()
}

func (c *console) cursor (x, y uint, s shape.Shape) {
  y0, y1 := shape.Cursor (x, y, c.ht1, c.cursorShape, s)
  if y0 + y1 == 0 { return }
  c.cursorShape = s
//  Lock()
  x -= uint(c.x); y -= uint(c.y)
  c.RectangleFullInv (int(x), int(y + y0), int(x + c.wd1) - 1, int(y + y1))
//  Unlock()
}

func (cons *console) Warp (l, c uint, s shape.Shape) {
  cons.WarpGr (cons.wd1 * c, cons.ht1 * l, s)
}

func (c *console) WarpGr (x, y uint, s shape.Shape) {
  x += uint(c.x); y += uint(c.y)
  c.blinkMutex.Lock()
  c.blinkX, c.blinkY = x, y
  c.blinkShape = s
  c.cursor (x, y, c.blinkShape)
  c.blinkMutex.Unlock()
}

// text ////////////////////////////////////////////////////////////////

func (X *console) Transparent() bool {
  return X.transparent
}

func (X *console) Transparence (on bool) {
  X.transparent = on
}

func (X *console) Write1 (b byte, l, c uint) {
  if ! visible { return }
  if l >= X.nLines || c >= X.nColumns { return }
  f := X.codeF
//  w := X.lineWd
//  X.lineWd = linewd.Thin
  for i := uint(0); i < X.ht1; i++ {
    for j := uint(0); j < X.wd1; j++ {
      if X.pointed (X.fontsize, b, i, j) {
        X.codeF = f
      } else {
        X.codeF = X.codeB
      }
      X.Point (int(X.wd1 * c + j), int(X.ht1 * l + i))
    }
  }
  X.codeF = f
//  X.lineWd = w
}

func (X *console) Write (s string, l, c uint) {
  if len(s) == 0 || ! visible { return }
  n := len (s)
  if c + uint(n) > X.nColumns { n = int(X.nColumns - c) }
  for i := 0; i < n; i++ {
    X.Write1 (s[i], l, c + uint(i))
  }
}

func (X *console) WriteNat (n, l, c uint) {
  t := "00"
  if n > 0 {
    const M = 10
    bs := make (obj.Stream, M)
    for i := 0; i < M; i++ {
      bs[M - 1 - i] = byte('0' + n % 10)
      n = n / M
    }
    s := 0
    for s < M && bs[s] == '0' {
      s++
    }
    t = ""
    if s == M - 1 { s = M - 2 }
    for i := s; i < M - int(n); i++ {
      t += string(bs[i])
    }
  }
  X.Write (t, l, c)
}

func (X *console) WriteNatGr (n uint, x, y int) {
}

func (X *console) Write1Gr (b byte, x, y int) {
  if ! visible { return }
  f := X.codeF
//  w := X.lineWd
//  X.lineWd = linewd.Thin
  for i := uint(0); i < X.ht1; i++ {
    for j := uint(0); j < X.wd1; j++ {
      if X.pointed (X.fontsize, b, i, j) {
        X.codeF = f
        X.Point (x + int(j), y + int(i))
      } else if ! X.transparent {
        X.codeF = X.codeB
        X.Point (x + int(j), y + int(i))
      }
    }
  }
//  X.lineWd = w
  X.codeF = f
}

func (X *console) WriteGr (s string, x, y int) {
  n := len (s)
  if n == 0 || ! visible { return }
  if x < X.x || y < X.y { return }
  n = len(s)
  for i := 0; i < n; i++ {
    X.Write1Gr (s[i], x + i * int(X.wd1), y)
  }
}

func (X *console) Write1InvGr (b byte, x, y int) {
  if ! visible { return }
  if x < X.x || x >= X.x + int(X.wd - X.wd1) || y < X.y || y >= X.y + int(X.ht - X.ht1) { return }
  for i := uint(0); i < X.ht1; i++ {
    for j := uint(0); j < X.wd1; j++ {
      if X.pointed (X.fontsize, b, i, j) {
        X.PointInv (x + int(j), y + int(i))
      } else if ! X.transparent {
        X.PointInv (x + int(j), y + int(i))
      }
    }
  }
}

func (X *console) WriteInvGr (s string, x, y int) {
  n := len (s)
  if n == 0 || ! visible { return }
  if x < 0 || y < 0 { return }
  for i := 0; i < n; i++ {
    X.Write1InvGr (s[i], x + i * int(X.wd1), y)
  }
}

// font ////////////////////////////////////////////////////////////////

func (X *console) ActFontsize() font.Size {
  return X.fontsize
}

func (X *console) SetFontsize (f font.Size) {
  X.fontsize = f
  X.ht1, X.wd1 = font.Ht (X.fontsize), font.Wd (X.fontsize)
  X.nLines, X.nColumns = X.ht / X.ht1, X.wd / X.wd1
}

func (x *console) Wd1() uint{
  return x.wd1
}

func (x *console) Ht1() uint{
  return x.ht1
}

func (x *console) NLines() uint{
  return x.nLines
}

func (x *console) NColumns() uint{
  return x.nColumns
}

// graphics ////////////////////////////////////////////////////////////

type
  pointFunc func (int, int)

func (X *console) SetLinewidth (w linewd.Linewidth) {
  X.lineWd = w
}

func (X *console) ActLinewidth() linewd.Linewidth {
  return X.lineWd
}

func (X *console) iok (x, y int) bool {
  if ! visible { return false }
  if x < 0 || y < 0 { return false }
  return x < int(X.wd) && y < int(X.ht)
//  return x < X.x + int(X.wd) && y < X.y + int(X.ht)
}

func (X *console) iok4 (x, y, x1, y1 int) bool {
  if ! visible { return false }
  if x < 0 || y < 0 || x1 < 0 || y1 < 0 { return false }
  return true
}

func intord (x, y, x1, y1 *int) {
  if *x > *x1 { *x, *x1 = *x1, *x }
  if *y > *y1 { *y, *y1 = *y1, *y }
}

func (X *console) Point (x, y int) {
  if ! visible || ! X.iok (x, y) { return }
  x += X.x; y += X.y
//  ux, uy := uint(x), uint(y)
  a := (int(width) * y + x) * int(colourdepth)
  copy (fbcop[a:a+int(colourdepth)], X.codeF)
  if ! X.buff {
    copy (fbmem[a:a+int(colourdepth)], X.codeF)
  }
/* TODO
  if X.lineWd > Thin && ux + 1 < X.wd && uy + 1 < X.ht {
    if ux + 1 < X.ht {
      a += int(colourdepth)
      copy (fbcop[a:a+int(colourdepth)], X.codeF)
      if ! X.buff {
        copy (fbmem[a:a+int(colourdepth)], X.codeF)
      }
    }
    if uy + 1 < X.wd {
      a += int(width - 1) * int(colourdepth)
      copy (fbcop[a:a+int(colourdepth)], X.codeF)
      if ! X.buff {
        copy (fbmem[a:a+int(colourdepth)], X.codeF)
      }
    }
    if X.lineWd == Thick {
      a += int(colourdepth)
      copy (fbcop[a:a+int(colourdepth)], X.codeF)
      if ! X.buff {
        copy (fbmem[a:a+int(colourdepth)], X.codeF)
      }
    } else { // Thicker
      if ux > 0 && uy > 0 {
        a -= int(width * 2 * colourdepth)
        copy (fbcop[a:a+int(colourdepth)], X.codeF)
        if ! X.buff {
          copy (fbmem[a:a+int(colourdepth)], X.codeF)
        }
        a += int(width - 1) * int(colourdepth)
        copy (fbcop[a:a+int(colourdepth)], X.codeF)
        if ! X.buff {
          copy (fbmem[a:a+int(colourdepth)], X.codeF)
        }
      }
    }
  }
  if X.lineWd > Thin && ux + 1 < X.wd && uy + 1 < X.ht { // still buggy TODO
    a += int(colourdepth)
    copy (fbcop[a:a+int(colourdepth)], X.codeF)
    a += int(width - 1) * int(colourdepth)
    copy (fbcop[a:a+int(colourdepth)], X.codeF)
    if X.lineWd == Thick {
      a += int(colourdepth)
      copy (fbcop[a:a+int(colourdepth)], X.codeF)
    } else { // Thicker
      if ux > 0 && uy > 0 {
        a -= int(width * 2 * colourdepth)
        copy (fbcop[a:a+int(colourdepth)], X.codeF)
        a += int(width - 1) * int(colourdepth)
        copy (fbcop[a:a+int(colourdepth)], X.codeF)
      }
    }
  }
*/
}

func (X *console) PointInv (x, y int) {
  if ! X.iok (x, y) { return }
  c := X.Colour (uint(x), uint(y))
  c.Invert()
  X.ColourF (c)
  X.Point (x, y)
  X.ColourF (X.cF)
}

func (X *console) OnPoint (x, y, a, b int, d uint) bool {
  dx, dy := x - a, y - b
  return dx * dx + dy * dy <= int(d * d)
}

// Returns true, iff m is up to tolerance t between i and k.
func between (i, k, m, t int) bool {
  return i <= m + t && m <= k + t || k <= m + t && m <= i + t
}

func ok2 (xs, ys []int) bool {
  if ! visible { return false }
  n := len (xs)
  return n != 0 && n == len (ys)
}

func ok4 (xs, ys, xs1, ys1 []int) bool {
  if ! visible { return false }
  n := len (xs)
  return n != 0 && n == len (ys) && n == len (xs1) && len (xs1) == len (ys1)
}

func (X *console) Points (xs, ys []int) {
  if ! ok2 (xs, ys) { return }
  for i := 0; i < len (xs); i++ {
    X.Point (xs[i], ys[i])
  }
}

func (X *console) PointsInv (xs, ys []int) {
  if ! ok2 (xs, ys) { return }
  for i := 0; i < len (xs); i++ {
    X.PointInv (xs[i], ys[i])
  }
}

func (X *console) OnPoints (xs, ys []int, a, b int, d uint) bool {
  n := len(xs)
  for i := 0; i < n; i++ {
    if X.OnPoint (xs[i], ys[i], a, b, d) {
      return true
    }
  }
  return false
}

// Pre: x <= x1 < Wd, y < Ht.
func (X *console) horizontal (x, y, x1 int, f pointFunc) {
  if x == x1 { f (x, y); return }
  if x > x1 { x, x1 = x1, x }
//  if x >= X.wd { return }
//  if x1 >= int(X.wd) { x1 = int(X.wd) - 1 }
  x0 := x
  for x := x0; x <= x1; x++ {
    f (x, y)
  }
/*
  if X.lineWd > Thin && y + 1 <= int(X.ht) {
    for x := x0; x <= x1; x++ {
      f (x, y + 1)
    }
  }
  if X.lineWd > Thick && y > 0 {
    for x := x0; x <= x1; x++ {
      f (x, y - 1)
    }
  }
*/
}

// Pre: x < Wd, y <= y1 < Ht.
func (X *console) vertical (x, y, y1 int, f pointFunc) {
  if y > y1 { y, y1 = y1, y }
//  if y1 >= int(X.ht) { y1 = int(X.ht) - 1 }
  y0 := y
  for y := y0; y <= y1; y++ {
    f (x, y)
  }
/*
  if X.lineWd > Thin && x + 1 < int(X.wd) {
    for y := y0; y <= y1; y++ {
      f (x + 1, y)
    }
  }
  if X.lineWd > Thick && x > 0 {
    for y := y0; y <= y1; y++ {
      f (x - 1, y)
    }
  }
*/
}

// Pre: 0 <= x <= x1 < NColumns, 0 <= y != y1 < NLines.
func (X *console) bresenham (x, y, x1, y1 int, f pointFunc) {
  dx := x1 - x
  Fehler, dy := 0, 0
  if y <= y1 { // gradient positive
    dy = y1 - y
    if dy <= dx { // gradient <= 45°
      for {
        f (x, y)
        if x == x1 { break }
        x++
        Fehler += 2 * dy
        if Fehler > dx {
          y++
          Fehler -= 2 * dx
        }
      }
    } else { // gradient > 45°
      for {
        f (x, y)
        if y == y1 { break }
        y++
        Fehler += 2 * dx
        if Fehler > dy {
          x++
          Fehler -= 2 * dy
        }
      }
    }
  } else { // gradient negative
    dy = y - y1
    if dy <= dx { // gradient >= -45°
      for {
        f (x, y)
        if x == x1 { break }
        x++
        Fehler += 2 * dy
        if Fehler > dx {
          y--
          Fehler -= 2 * dx
        }
      }
    } else { // gradient < -45°
      for {
        f (x, y)
        if y == y1 { break }
        y--
        Fehler += 2 * dx
        if Fehler > dy {
          x++
          Fehler -= 2 * dy
        }
      }
    }
  }
}

// Pre: 0 <= x <= x1 < xx, y != y1, 0 <= y, y1 < yy.
func (X *console) bresenhamInf (xx, yy, x, y, x1, y1 int, f pointFunc) {
  x0, y0 := x, y
  dx := x1 - x
  Fehler, dy := 0, 0
  if y <= y1 { // gradient positive
    dy = y1 - y
    if dy <= dx { // gradient <= 45°
      for {
        f (x, y)
        if x == xx - 1 || y == yy - 1 { break }
        x++
        Fehler += 2 * dy
        if Fehler > dx {
          y++
          Fehler -= 2 * dx
        }
      }
      x, y = x0, y0
      Fehler = 0
      for {
        Fehler += 2 * dy
        if Fehler > dx {
          y--
          Fehler -= 2 * dx
        }
        f (x, y)
        if x == 0 || y == 0 { break }
        x--
      }
    } else { // gradient > 45°
      for {
        f (x, y)
        if y == yy - 1 || x == xx - 1 { break }
        y++
        Fehler += 2 * dx
        if Fehler > dy {
          x++
          Fehler -= 2 * dy
        }
      }
      x, y = x0, y0
      Fehler = 0
      for {
        Fehler += 2 * dx
        if Fehler > dy {
          x--
          Fehler -= 2 * dy
        }
        f (x, y)
        if x == 0 || y == 0 { break }
        y--
      }
    }
  } else {
    dy = y - y1 // gradient negative
    if dy <= dx { // gradient >= -45°
      for {
        f (x, y)
        if x == xx - 1 || y == 0 { break }
        x++
        Fehler += 2 * dy
        if Fehler > dx {
          y--
          Fehler -= 2 * dx
        }
      }
      x, y = x0, y0
      Fehler = 0
      for {
        f (x, y)
        if x == 0 || y == yy - 1 { break }
        x--
        Fehler += 2 * dy
        if Fehler > dx {
          y++
          Fehler -= 2 * dx
        }
      }
    } else { // gradient < -45°
      for {
        f (x, y)
        if x == xx - 1 || y == 0 { break }
        y--
        Fehler += 2 * dx
        if Fehler > dy {
          x++
          Fehler -= 2 * dy
        }
      }
      x, y = x0, y0
      Fehler = 0
      for {
        f (x, y)
        if x == 0 || y == yy - 1 { break }
        y++
        Fehler += 2 * dx
        if Fehler > dy {
          x--
          Fehler -= 2 * dy
        }
      }
    }
  }
}

func nat (x, y int) bool {
  return x >= 0 && y >= 0
}

func (X *console) line (x, y, x1, y1 int, f pointFunc) {
  if x1 < x { x, x1 = x1, x; y, y1 = y1, y }
  if ! X.iok4 (x, y, x1, y1) {
    return
  }
  if y == y1 {
    X.horizontal (x, y, x1, f)
    return
  }
  if x == x1 {
    X.vertical (x, y, y1, f)
    return
  }
  X.bresenham (x, y, x1, y1, f)
}

func (X *console) Line (x, y, x1, y1 int) {
  X.line (x, y, x1, y1, X.Point)
}

func (X *console) LineInv (x, y, x1, y1 int) {
  X.line (x, y, x1, y1, X.PointInv)
}

func (X *console) OnLine (x, y, x1, y1, a, b int, t uint) bool {
  if x1 < x { x, x1 = x1, x; y, y1 = y1, y }
  if ! (between (x, x1, a, int(t)) && between (y, y1, b, int(t))) {
    return false
  }
  if x == x1 {
    return between (x, x, a, int(t)) && between (y, y1, b, int(t))
  }
  if y == y1 {
    return between (y, y, b, int(t)) && between (x, x1, a, int(t))
  }
  X.xx_, X.yy_, X.tt_, X.incident = a, b, int(t * t), false
  X.bresenham (x, y, x1, y1, X.onPoint)
  return X.incident
}

func (X *console) lines (xs, ys, xs1, ys1 []int, f pointFunc) {
  if ! ok4 (xs, ys, xs1, ys1) { return }
  for i := 0; i < len (xs); i++ {
    if X.iok (xs[i], ys[i]) && X.iok (xs1[i], ys1[i]) {
      X.line (xs[i], ys[i], xs1[i], ys1[i], f)
    }
  }
}

func (X *console) Lines (xs, ys, xs1, ys1 []int) {
  X.lines (xs, ys, xs1, ys1, X.Point)
}

func (X *console) LinesInv (xs, ys, xs1, ys1 []int) {
  X.lines (xs, ys, xs1, ys1, X.PointInv)
}

func (X *console) OnLines (xs, ys, xs1, ys1 []int, a, b int, t uint) bool {
  if ! ok4 (xs, ys, xs1, ys1) { return false }
  if len (xs) == 1 {
    return between (xs[0], xs[0], a, int(t)) && between (ys[0], ys[0], b, int(t))
  }
  for i := 0; i < len (xs); i++ {
    if X.OnLine (xs[i], ys[i], xs1[i], ys1[i], a, b, t) {
      return true
    }
  }
  return false
}

func (X *console) segs (xs, ys []int, f pointFunc) {
  if ! ok2 (xs, ys) { return }
  n := len (xs)
  for i := 0; i < n; i++ {
    if ! X.iok (xs[i], ys[i]) {
      return
    }
  }
  if n == 0 {
    f (xs[0], ys[0])
  } else {
    for i := 1; i < len (xs); i++ {
      X.line (xs[i-1], ys[i-1], xs[i], ys[i], f)
    }
  }
}

func (X *console) Segments (xs, ys []int) {
  if ! ok2 (xs, ys) { return }
  X.segs (xs, ys, X.Point)
}

func (X *console) SegmentsInv (xs, ys []int) {
  if ! ok2 (xs, ys) { return }
  X.segs (xs, ys, X.PointInv)
  if len (xs) > 1 {
    for i := 1; i < len (xs); i++ {
      X.PointInv (xs[i], ys[i])
    }
  }
}

func (X *console) OnSegments (xs, ys []int, a, b int, t uint) bool {
  if ! ok2 (xs, ys) { return false }
  if len (xs) == 1 {
    return xs[0] == a && ys[0] == b // TODO, weil das noch Blödsinn ist
  }
  for i := 1; i < len (xs); i++ {
    if X.OnLine (xs[i-1], ys[i-1], xs[i], ys[i], a, b, t) {
      return true
    }
  }
  return false
}

func (X *console) onPoint (x, y int) {
  X.incident = X.incident || (x - X.xx_) * (x - X.xx_) + (y - X.yy_) * (y - X.yy_) <= X.tt_
}

func (X *console) infLine (x, y, x1, y1 int, f pointFunc) {
  if x == x1 && y == y1 { return }
  if x1 < x { x, x1 = x1, x; y, y1 = y1, y }
  if ! visible { return }
  if y == y1 {
    X.horizontal (0, y, int(width) - 1, f)
    return
  }
  if x == x1 {
    X.vertical (x, 0, int(height), f)
    return
  }
  X.bresenhamInf (int(width), int(height), x, y, x1, y1, f)
}

func (X *console) InfLine (x, y, x1, y1 int) {
  X.infLine (x, y, x1, y1, X.Point)
}

func (X *console) InfLineInv (x, y, x1, y1 int) {
  X.infLine (x, y, x1, y1, X.PointInv)
}

func (X *console) OnInfLine (x, y, x1, y1, a, b int, t uint) bool {
  if x1 < x { x, x1 = x1, x; y, y1 = y1, y }
  X.xx_, X.yy_, X.tt_, X.incident = a, b, int(t * t), false
  X.bresenhamInf (int(width), int(height), x, y, x1, y1, X.onPoint)
  return X.incident
}

func (X *console) Triangle (x, y, x1, y1, x2, y2 int) {
  X.Line (x, y, x1, y1)
  X.Line (x1, y1, x2, y2)
  X.Line (x2, y2, x, y)
}

func (X *console) TriangleInv (x, y, x1, y1, x2, y2 int) {
  X.LineInv (x, y, x1, y1)
  X.LineInv (x1, y1, x2, y2)
  X.LineInv (x2, y2, x, y)
}

func (X *console) TriangleFull (x, y, x1, y1, x2, y2 int) {
  X.PolygonFull ([]int{x, x1, x2}, []int{y, y1, y2})
}

func (X *console) TriangleFullInv (x, y, x1, y1, x2, y2 int) {
  X.PolygonFullInv ([]int{x, x1, x2}, []int{y, y1, y2})
}

func (X *console) rectang (x, y, x1, y1 int, f pointFunc) {
  if ! X.rectangOk (&x, &y, &x1, &y1) { return }
  if x == x1 {
    if y == y1 {
      f (x, y)
    } else {
      X.vertical (int(x), int(y), int(y1), f)
    }
    return
  }
  X.horizontal (x, y, x1, f)
  if y == y1 {
    return
  }
  X.horizontal (x, y1, x1, f)
  X.vertical (x, y, y1, f)
  X.vertical (x1, y, y1, f)
}

func (X *console) Rectangle (x, y, x1, y1 int) {
  X.rectang (x, y, x1, y1, X.Point)
}

func (X *console) RectangleInv (x, y, x1, y1 int) {
  X.rectang (x, y, x1, y1, X.PointInv)
  X.PointInv (x, y)
  X.PointInv (x1, y)
  X.PointInv (x, y1)
  X.PointInv (x1, y1)
}

func (X *console) RectangleFull (x, y, x1, y1 int) {
  intord (&x, &y, &x1, &y1)
  if x1 >= int(X.wd) { x1 = int(X.wd) - 1 }
  if y1 >= int(X.ht) { y1 = int(X.ht) - 1 }
  for y <= y1 {
    X.horizontal (x, y, x1, X.Point)
    y++
  }
}

func (X *console) RectangleFullInv (x, y, x1, y1 int) {
  intord (&x, &y, &x1, &y1)
  if x1 >= int(X.wd) { x1 = int(X.wd) - 1 }
  if y1 >= int(X.ht) { y1 = int(X.ht) - 1 }
  for y <= y1 {
    X.horizontal (x, y, x1, X.PointInv)
    y++
  }
}

func (X *console) OnRectangle (x, y, x1, y1, a, b int, t uint) bool {
  if ! (between (x, x1, a, int(t)) && between (y, y1, b, int(t))) {
    return false
  }
  return between (x, x, a, int(t)) || between (x1, x1, a, int(t)) ||
         between (y, y, b, int(t)) || between (y1, y1, b, int(t))
}

func (X *console) InRectangle (x, y, x1, y1, a, b int, t uint) bool {
  return between (x, x1, a, int(t)) && between (y, y1, b, int(t))
}

func (X *console) Polygon (xs, ys []int) {
  if ! ok2 (xs, ys) { return }
  n := len (xs)
  if n < 3 { return }
  X.segs (xs, ys, X.Point)
  X.line (xs[n-1], ys[n-1], xs[0], ys[0], X.Point)
}

func (X *console) PolygonInv (xs, ys []int) {
  if ! ok2 (xs, ys) { return }
  n := len (xs)
  if n < 3 { return }
  X.segs (xs, ys, X.PointInv)
  X.line (xs[n-1], ys[n-1], xs[0], ys[0], X.PointInv)
  X.PointInv (xs[0], ys[0])
  X.PointInv (xs[n-1], ys[n-1])
}

func (X *console) interior (x, y int, xs, ys []int) bool {
  return false // TODO winding number algorithm
}

func (X *console) mark (x, y int) {
  X.polygon[x][y] = true
  X.Point (x, y)
}

func (X *console) markInv (x, y int) {
  X.polygon[x][y] = true
  X.PointInv (x, y)
}

func (X *console) demark (x, y int) {
  X.polygon[x][y] = false
}

func (X *console) dedone() {
  for x := uint(0); x < X.wd; x++ {
    for y := uint(0); y < X.ht; y++ {
      X.done[x][y] = false
    }
  }
}

func (X *console) st (x, y int, f pointFunc) {
  if X.polygon[x][y] {
    return
  }
  if ! X.done[x][y] {
    X.done[x][y] = true
    f (x, y)
    if y > 0 { X.st (x, y - 1, f) }
    if x > 0 { X.st (x - 1, y, f) }
    if y + 1 < int(X.ht) { X.st (x, y + 1, f) }
    if x + 1 < int(X.wd) { X.st (x + 1, y, f) }
  }
}

func (X *console) setInv (x, y int) {
  X.st (x, y, X.PointInv)
}

func (X *console) set (x, y int) {
  X.st (x, y, X.Point)
}

func (X *console) polygonFull (xs, ys []int, m, s pointFunc) {
  if ! ok2 (xs, ys) { return }
  n := len (xs)
  if n < 2 { return }
  X.segs (xs, ys, m)
  xx, yy := 0, 0
  xMin, yMin := int(X.wd), int(X.ht)
  xMax, yMax := 0, 0
  for i := 0; i < int(n); i++ {
    xx += xs[i]; yy += ys[i]
    if xs[i] < xMin { xMin = xs[i] }
    if ys[i] < yMin { yMin = ys[i] }
    if xs[i] > xMax { xMax = xs[i] }
    if ys[i] > yMax { yMax = ys[i] }
  }
  s (xx / n, yy / n)
  X.segs (xs, ys, X.demark)
  X.dedone()
}

func (X *console) PolygonFull (xs, ys []int) {
  X.polygonFull (xs, ys, X.mark, X.set)
}

func (X *console) PolygonFullInv (xs, ys []int) {
  X.polygonFull (xs, ys, X.markInv, X.setInv)
}

func (X *console) OnPolygon (xs, ys []int, a, b int, t uint) bool {
  n := len (xs)
  if n == 0 { return false }
  if ! ok2 (xs, ys) { return false }
  if n == 1 { return xs[0] == a && ys[0] == b }
  for i := 1; i < int(n); i++ {
    if X.OnLine (xs[i-1], ys[i-1], xs[i], ys[i], a, b, t) {
      return true
    }
  }
  return X.OnLine (xs[n-1], ys[n-1], xs[0], ys[0], a, b, t)
}

func (X *console) circ (x, y int, r uint, filled bool, f pointFunc) {
// Algorithm of Bresenham (Fellner: Computer Grafik, 5.5)
  if ! visible { return }
  if x >= int(X.wd) || y >= int(X.ht) || r >= X.wd {
    return
  }
  if r == 0 {
    f (x, y)
    return
  }
  x1, y1 := 0, int(r)
  Fehler := 3
  Fehler -= 2 * int(r)
/*
  if filled {
    X.horizontal (x - r, y, x + r, b)
    X.Point (x, y - r)
    X.Point (x, y + r)
  } else {
    f (x - r, y    )
    f (x + r, y    )
    f (x    , y - r)
    f (x    , y + r)
  }
  x1++
  if Fehler >= 0 {
    y1--
    Fehler -= 4 * y1
  }
  Fehler += 6
*/
  y0 := y1 + 1
  for x1 <= y1 {
    if filled {
      X.horizontal (x - y1, y - x1, x + y1, f)
      if x1 > 0 {
        X.horizontal (x - y1, y + x1, x + y1, f)
      }
      if y1 < y0 { // not yet correct, but a bit better than the above code
        y0 = y1
        X.horizontal (x - x1, y - y1, x + x1, f)
        X.horizontal (x - x1, y + y1, x + x1, f)
      }
    } else {
      f (x - y1, y - x1)
      f (x + y1, y - x1)
      f (x - y1, y + x1)
      f (x + y1, y + x1)
      f (x - x1, y - y1)
      f (x + x1, y - y1)
      f (x - x1, y + y1)
      f (x + x1, y + y1)
    }
    x1++
    if Fehler >= 0 {
      y1--
      Fehler -= 4 * y1
    }
    Fehler += 4 * x1 + 2
  }
}

func (X *console) Circle (x, y int, r uint) {
  if ! X.iok (x, y) { return }
  if uint(x) >= r && uint(y) >= r {
    X.circ (x, y, r, false, X.Point)
  }
}

func (X *console) CircleInv (x, y int, r uint) {
  if ! X.iok (x, y) { return }
  if uint(x) >= r && uint(y) >= r {
    X.circ (x, y, r, false, X.PointInv)
  }
}

func (X *console) CircleFull (x, y int, r uint) {
  if ! X.iok (x, y) { return }
  if uint(x) >= r && uint(y) >= r {
    X.circ (x, y, r, true, X.Point)
  }
}

func (X *console) CircleFullInv (x, y int, r uint) {
  if ! X.iok (x, y) { return }
  if uint(x) >= r && uint(y) >= r {
    X.circ (x, y, r, true, X.PointInv)
  }
}

func (X *console) arc (x, y int, r uint, a, b float64, filled bool, f pointFunc) {
  if filled { ker.Panic ("filled arcs not yet implemented") }
// lousy implementation, but better than nothing
  a0, b0, r0, db := a / 180 * math.Pi, b / 180 * math.Pi, float64(r), 1.0 / 180 * math.Pi
  a1 := a0; if b0 > 0 { a1 += b0 } else { a0 += b0 }
  var x1, y1 []int
  for alpha := a0; alpha < a1; alpha += db {
    x1, y1 = append (x1, x + int(r0 * math.Cos(alpha))), append(y1, y - int(r0 * math.Sin(alpha)))
  }
  x1, y1 = append (x1, x + int(r0 * math.Cos(a1))), append(y1, y - int(r0 * math.Sin(a1)))
  for i := 1; i < len(x1); i+= 1 {
    X.line (x1[i-1], y1[i-1], x1[i], y1[i], f)
  }
}

func (X *console) Arc (x, y int, r uint, a, b float64) {
  if ! X.iok (x, y) { return }
  if uint(x) >= r && uint(y) >= r {
    X.arc (x, y, r, a, b, false, X.Point)
  }
}

func (X *console) ArcInv (x, y int, r uint, a, b float64) {
  if ! X.iok (x, y) { return }
  if uint(x) >= r && uint(y) >= r {
    X.arc (x, y, r, a, b, false, X.PointInv)
  }
}

func (X *console) ArcFull (x, y int, r uint, a, b float64) {
  if ! X.iok (x, y) { return }
  if uint(x) >= r && uint(y) >= r {
    X.arc (x, y, r, a, b, true, X.Point)
  }
}

func (X *console) ArcFullInv (x, y int, r uint, a, b float64) {
  if ! X.iok (x, y) { return }
  if uint(x) >= r && uint(y) >= r {
    X.arc (x, y, r, a, b, true, X.PointInv)
  }
}

func (X *console) OnCircle (x, y int, r uint, a, b int, t uint) bool {
  d := uint(dist2 (x, y, a, b))
  if d >= r {
    return d <= t + r
  }
  return d + t > r
}

func (X *console) InCircle (x, y int, r uint, a, b int, t uint) bool {
  return uint(dist2 (x, y, a, b)) <= r + t
}

func (X *console) ell (x, y int, a, b uint, filled bool, f pointFunc) {
  if ! X.iok (x, y) { return }
  if a == b {
    X.circ (x, y, a, filled, f)
    return
  }
  if a == 0 {
    if b == 0 {
      f (x, y)
    } else {
      X.vertical (x, y - int(b), y + int(b), f)
    }
    return
  } else {
    if b == 0 {
      X.horizontal (x - int(a), y, x + int(a), f)
      return
    }
  }
  a1, b1 := 2 * a * a, 2 * b * b
  i := int (a * b * b)
  x2, y2 := int(2 * a * b * b), 0
  xi, x1 := x - int(a), x + int(a)
  yi, y1 := y, y
  var xl int
  if xi < 0 {
    xl = 0
  } else {
    xl = xi
  }
  if filled {
    X.horizontal (xl, y, x1, f)
  } else {
    f (xl, y)
    f (int(x1), y)
  }
  var yo int
  if a == 0 {
    if y < int(b) {
      yo = 0
    } else {
      yo = y - int(b)
    }
    X.vertical (xi, yo, y + int(b), f)
    return
  }
  for { // a > uint(0) {
    if i > 0 {
      yi--
      y1++
      y2 += int(a1)
      i -= int(y2)
    }
    if i <= 0 {
      xi++
      x1--
      x2 -= int(b1)
      i += int(x2)
      a--
    }
    if xi < 0 {
      xl = 0
    } else {
      xl = xi
    }
    if yi < 0 {
      yo = 0
    } else {
      yo = yi
    }
    var xr int
    if x1 < int(X.wd) {
      xr = int(x1)
    } else {
      xr = int(X.wd) - 1
    }
    var yu int
    if y1 < int(X.ht) {
      yu = int(y1)
    } else {
      yu = int(X.ht) - 1
    }
    if filled {
      X.horizontal (xl, yo, xr, f)
      X.horizontal (xl, yu, xr, f)
    } else {
      f (xl, yo)
      f (xr, yo)
      f (xl, yu)
      f (xr, yu)
    }
    if a == uint(0) {
      break
    }
  }
}

func (X *console) Ellipse (x, y int, a, b uint) {
  if ! X.iok (x, y) { return }
  if uint(x) >= a && uint(y) >= b {
    X.ell (x, y, a, b, false, X.Point)
  }
}

func (X *console) EllipseInv (x, y int, a, b uint) {
  if ! X.iok (x, y) { return }
   if uint(x) >= a && uint(y) >= b {
    X.ell (x, y, a, b, false, X.PointInv)
  }
}

func (X *console) EllipseFull (x, y int, a, b uint) {
  if ! X.iok (x, y) { return }
  if uint(x) >= a && uint(y) >= b {
    X.ell (x, y, a, b, true, X.Point)
  }
}


func (X *console) EllipseFullInv (x, y int, a, b uint) {
  if ! X.iok (x, y) { return }
  if uint(x) >= a && uint(y) >= b {
    X.ell (x, y, a, b, true, X.PointInv)
  }
}

func (X *console) OnEllipse (x, y int, a, b uint, A, B int, t uint) bool {
  if ! X.iok (x, y) { return false }
  X.xx_, X.yy_, X.tt_, X.incident = A, B, int(t * t), false
  X.ell (x, y, a, b, false, X.onPoint)
  return X.incident
}

func (X *console) curve (xs, ys []int, f pointFunc) {
  m := len (xs)
  if m == 0 || m != len (ys) {
panic ("curve: wrong m")
    return
  }
  n := ker.ArcLen (xs, ys)
  xs1, ys1 := make ([]int, n), make ([]int, n)
  for i := uint(0); i < n; i++ {
    xs1[i], ys1[i] = ker.Bezier (xs, ys, uint(m), n, i)
  }
  f (xs[0], ys[0])
  for i := 0; i < len(xs1); i++ {
    f (xs1[i], ys1[i])
  }
}

func (X *console) Curve (xs, ys []int) {
  X.curve (xs, ys, X.Point)
}

func (X *console) CurveInv (xs, ys []int) {
  X.curve (xs, ys, X.PointInv)
}

func (X *console) OnCurve (xs, ys []int, a, b int, t uint) bool {
  if ! ok2 (xs, ys) {
panic ("OnCurve: ! ok2")
    return false
  }
  X.xx_, X.yy_, X.tt_, X.incident = a, b, int(t * t), false
  X.curve (xs, ys, X.onPoint)
  return X.incident
}

// mouse ///////////////////////////////////////////////////////////////

func (X *console) MousePos() (uint, uint) {
  xm, ym := mouse.Pos()
  return uint(ym - X.y) / X.ht1, uint(xm - X.x) / X.wd1 // Offset
}

func (X *console) MousePosGr() (int, int) {
  xm, ym := mouse.Pos()
  return xm - X.x, ym - X.y // Offset
}

func (X *console) SetPointer (p ptr.Pointer) {
  X.pointer = p
}

var (
  pointer = [ptr.NPointers][]string {
    []string { // Ptr
      "# . . . . . . . . . . . . ",
      "# # . . . . . . . . . . . ",
      "# o # . . . . . . . . . . ",
      "# * o # . . . . . . . . . ",
      "# * * o # . . . . . . . . ",
      "# * * * o # . . . . . . . ",
      "# * * * * o # . . . . . . ",
      "# * * * * * o # . . . . . ",
      "# * * * * * * o # . . . . ",
      "# * * * * * * * o # . . . ",
      "# * * * * * * * * o # . . ",
      "# * * * * * * * * * o # . ",
      "# * * * * * # # # # # # # ",
      "# * * * o * o # . . . . . ",
      "# * * # * * * # . . . . . ",
      "# * # # # o * o # . . . . ",
      "# # . . . # * * # . . . . ",
      "# . . . . # o * o . . . . ",
      ". . . . . . # o # . . . . " },
    []string { // Gumby
      ". . # # # # # # . . . . . . . . ",
      "* * o # * * * * # . . . . . . . ",
      "# # * o # * * * * # . . . . . . ",
      "# # # * # * . * . * # . . . . . ",
      "# # * o # * * * * * # . . . . . ",
      "# # * o # * . . . * # * * * . . ",
      "# # # # # * * * * * # # # # * . ",
      "* * # # # * * * * * # # # # # # ",
      ". . * * # * * * * * # . * # # # ",
      ". . . . # * * * * * # . * # # # ",
      ". . . . # * * # * * # * # # # # ",
      ". . . . # * * # * * # . * # # # ",
      ". . . . # * * # * * # . . * * * ",
      ". . . # * * * # * * * # . . . . ",
      ". . # * * * * # * * * * # . . . ",
      ". . # # # # # . # # # # # . . . " },
    []string { // Hand
      ". . . o # # o . . . . . . . . . ",
      ". . . # * * # . . . . . . . . . ",
      ". . . # * * # . . . . . . . . . ",
      ". . . # * * # . . . . . . . . . ",
      ". . . # * * # # # o . . . . . . ",
      ". . . # * * # o * o # # o . . . ",
      ". . . # * * # * * # o * # # # o ",
      ". # # # * * # * * * * * # o * # ",
      "# o * # * * # * * * * * * * * # ",
      "# * * # * * * * * * * * * * * # ",
      "# * * # * * * * * * * * * * * # ",
      "# * * * * o o * o o * o o * * # ",
      "# * * * * * * * * * * * * * * # ",
      "# * * * * * * * * * * * * * * # ",
      "# o * * * * * * * * * * * * o # ",
      ". # # # # # # # # # # # # # # . " },
    []string { // Gobbler
      ". . . . . . . . * * * * * * . . ",
      ". . . . . . . . * # # # # * . . ",
      "* * . . . . . . * # # # * * * * ",
      "# * * * * * * * * * # # * * # # ",
      "# * * # # # # # # * # # * * * * ",
      "# # # # # # # # # # # # * * . . ",
      "# # # # # # # * * * # # * * . . ",
      "# # # # # # * * * * # # * * . . ",
      "* # # * * * * * * * # # # * . . ",
      "* * * * * * * * # # # # * * . . ",
      ". * * # * # # # # # # * * . . . ",
      ". . . * # * * * * * * * . . . . ",
      ". . . * # * . . . . . . . . . . ",
      ". . . * # * . . . . . . . . . . ",
      ". . * * # * * * . . . . . . . . ",
      ". . * # # # # * . . . . . . . . " },
    []string { // Watch
      ". . . . o # # # # # o . . . . . ",
      ". . . # o * * * * * # # . . . . ",
      ". . # * * * * * * * * * # . . . ",
      ". # * * * * * * * * * * * # . . ",
      "o o * * * * * * * * * * * o o . ",
      "# * * * * * * * * * * * * * # . ",
      "# * * * * * o # o * * * * * # . ",
      "# * * o # # # # # * * * * * # . ",
      "# * # # o * o # o * * * * * # . ",
      "# * * * * * * * # * * * * * # . ",
      "o o * * * * * * # * * * * o o . ",
      ". # * * * * * * * # * * * # . . ",
      ". . # * * * * * * # * * # . . . ",
      ". . . # o * * * * * # # . . . . ",
      ". . . . o # # # # # o . . . . . ",
      ". . . . . . . . . . . . . . . . " },
  }
  pointerHt = [ptr.NPointers]int { 18, 16, 16, 16, 16 }
  pointerWd = [ptr.NPointers]int { 12, 16, 16, 16, 16 }
)

func (X *console) initMouse () {
//  X.mouseOn = false
  mouse.Def (uint(0), uint(0), width, height)
  X.xMouse, X.yMouse = int(X.wd) / 2, int(X.ht) / 2
  mouse.Warp (uint(X.xMouse), uint(X.yMouse))
}

func (X *console) restore (x, y int) {
  a := (int(width) * y + x) * int(colourdepth)
  da := pointerWd[X.pointer] * int(colourdepth)
  w := width * colourdepth
// TODO limit to right screen border ???
  h1, ht := pointerHt[X.pointer], int(X.ht)
  if y + h1 > ht { h1 = ht - y }
  for h := 0; h < h1; h++ {
    copy (fbmem[a:a+da], fbcop[a:a+da])
    a += int(w)
  }
}

func (X *console) writePointer (x, y int) {
  cB := col.Black().EncodeInv()
  cW := col.LightWhite().EncodeInv()
  cG := col.LightGray().EncodeInv()
  var p obj.Stream
  for h := 0; h < pointerHt[X.pointer]; h++ {
    for w := 0; w < pointerWd[X.pointer]; w++ {
      switch pointer[X.pointer][h][2 * w] {
      case '#':
        p = cB
      case '*':
        p = cW
      case 'o':
        p = cG
      default:
        continue
      }
      if x + w < X.x || x + w >= X.x + int(X.wd) ||
         y + h < X.y || y + h >= X.y + int(X.ht) {
        continue
      }
      a := (int(width) * (y + h) + (x + w)) * int(colourdepth)
      copy (fbmem[a:a+int(colourdepth)], p)
//      copy (fbcop[a:a+int(colourdepth)], p) // No, we don't want that
    }
  }
}

func (X *console) MousePointer (on bool) {
  X.mouseOn = on
//  if ! mouse.Ex() || ! X.mouseOn || ! visible { return }
  if ! X.mouseOn || ! visible { return }
  if X.x <= X.xMouse + pointerWd[X.pointer] && X.xMouse < X.x + int(X.wd) &&
     X.y <= X.yMouse + pointerHt[X.pointer] && X.yMouse < X.y + int(X.ht) {
    X.restore (X.xMouse, X.yMouse)
  }
  if X == mouseConsole {
//    X.restore (X.xMouse, X.yMouse)
    X.xMouse, X.yMouse = mouse.Pos()
    X.writePointer (X.xMouse, X.yMouse)
  } else {
    // TODO full screen as root window
  }
}

func (X *console) MousePointerOn() bool {
//  if ! mouse.Ex() { return false }
  return X.mouseOn
}

func (X *console) WarpMouse (l, c uint) {
  mouse.Warp (uint(X.y) + l * X.ht, uint(X.x) + c * X.wd) // Offset
  X.MousePointer (true)
}

func (X *console) WarpMouseGr (x, y int) {
  mouse.Warp (uint(x + X.x), uint(y + X.y)) // Offset
  X.MousePointer (true)
}

func (X *console) UnderMouse (l, c, w, h uint) bool {
  lm, cm := X.MousePos()
  return l <= lm && lm < l + h && c <= cm && cm < c + w
}

func (X *console) UnderMouseGr (x, y, x1, y1 int, t uint) bool {
  intord (&x, &y, &x1, &y1)
  xm, ym := X.MousePosGr()
  return x <= int(xm) + int(t) && int(xm) <= x1 + int(t) &&
         y <= int(ym) + int(t) && int(ym) <= y1 + int(t)
}

func (X *console) UnderMouse1 (x, y int, d uint) bool {
  xm, ym := X.MousePosGr()
  return (x - xm) * (x - xm) + (y - ym) * (y - ym) <= int(d * d)
}

// serialisation ///////////////////////////////////////////////////////

func (X *console) Codelen (w, h uint) uint {
  return 2 * uint(4) + colourdepth * w * h
}

func (X *console) Encode (x, y, w, h uint) []byte {
  s := make (obj.Stream, X.Codelen (w, h))
  i := 2 * uint(4)
  copy (s[:i], obj.Encode4 (uint16(x), uint16(y), uint16(w), uint16(h)))
  for l := X.y; l < X.y + int(h); l++ {
    j := colourdepth * width * uint(l)
    for c := X.x; c < X.x + int(w); c++ {
      copy (s[i:i+3], fbmem[j:j+3])
      i += 3
      j += colourdepth
    }
  }
  return s
}

func (X *console) Decode (s obj.Stream) {
  if s == nil { return }
  if ! visible { return }
  i := 2 * uint(4)
  x0, y0, w, h := obj.Decode4 (s[:i])
  c := col.New()
  for y := int(y0); y < int(y0 + h); y++ {
    for x := int(x0); x < int(x0 + w); x++ {
      c.Set (s[i], s[i+1], s[i+2])
      X.cF, X.codeF = c, c.EncodeInv()
      X.Point (x, y)
      i += 3
    }
  }
}

// ppm-serialisation ///////////////////////////////////////////////////

func string_(n uint) string {
  if n == 0 { return "0" }
  var s string
  for s = ""; n > 0; n /= 10 {
    s = string(n % 10 + '0') + s
  }
  return s
}

func number (s obj.Stream) (uint, int) {
  n := uint(0)
  i := 0
  for char.IsDigit (s[i]) { i++ }
  for j := 0; j < i; j++ {
    n = 10 * n + uint(s[j] - '0')
  }
  return n, i
}

func (X *console) PPMHeader (w, h uint) string {
  s := "P6 " + string_(w) + " " + string_(h) + " 255" + string(byte(10))
  X.ppmheader = s
  X.lh = uint(len(s))
  return s
}

func (X *console) PPMCodelen (w, h uint) uint {
  X.PPMHeader (w, h)
  return X.lh + 3 * w * h
}

func (X *console) PPMSize (s obj.Stream) (uint, uint) {
  w, h, _, _ := X.PPMHeaderData (s)
  return w, h
}

func (X *console) PPMEncode (x0, y0, w, h uint) obj.Stream {
  s := X.Encode (x0, y0, w, h)
  return append (obj.Stream(X.PPMHeader (w, h)), s[2*4:]...)
}

func (X *console) PPMHeaderData (s obj.Stream) (uint, uint, uint, int) {
  p := string(s[:2]); if p != "P6" { panic ("wrong ppm-header: " + p) }
  i := 3
  if s[i] == '#' {
    for {
      i++
      if s[i] == byte(10) {
        i++
        break
      }
    }
  }
  w, dw := number (s[i:])
  i += dw + 1
  h, dh := number (s[i:])
  i += dh + 1
  m, dm := number (s[i:])
  i += dm
  return w, h, m, i + 1
}

func (X *console) PPMDecode (st obj.Stream, x0, y0 uint) {
  w, h, _, j := X.PPMHeaderData (st)
  if w == 0 || h == 0 || w > X.Wd() || h > X.Ht() { return }
  i := 4 * uint(2)
  l := i + 3 * w * h
  e := make(obj.Stream, l)
  copy (e[:i], obj.Encode4 (uint16(x0), uint16(y0), uint16(w), uint16(h)))
  if under_X {
    c := col.New()
    for y := uint(0); y < h; y++ {
      for x := uint(0); x < w; x++ {
        c.Decode (st[j:j+3])
        copy (e[i:i+3], obj.Encode (c.Code()))
        i += 3
        j += 3
      }
    }
  } else { // under_C, i.e. console
    copy (e[i:], st[j:])
  }
  X.Decode (e)
}

// cut buffer //////////////////////////////////////////////////////////

var
  buffer obj.Stream

func (x *console) Cut (s *string) {
// TODO
}

func (x *console) Copy (s string) {
  buffer = make(obj.Stream, len(s))
  copy (buffer[:], obj.Stream(s))
}

func (x *console) Paste() string {
  return string(buffer[:])
}

// framebuffer /////////////////////////////////////////////////////////

const (
  esc1 = "\x1b["
  ClearScreen = esc1 + "H" + esc1 + "J"
  home = esc1 + "?25h" + esc1 + "?0c"
)
var (
  fbmemsize uint
  fbmem, fbcop,
  emptyBackground obj.Stream
  visible bool // only for console switching
)

func consoleOn() {
  ker.ActivateConsole()
  n := width * height * uint(colourdepth)
  copy (fbmem[:n], fbcop[:n])
  visible = true
  c := actualC
  c.Warp (c.ht1 * c.blinkY, c.wd1 * c.blinkX, c.consoleShape)
}

func consoleOff() {
  visible = false
  c := actualC
  c.consoleShape = c.blinkShape
  c.Warp (c.ht1 * c.blinkY, c.wd1 * c.blinkX, shape.Off)
  ker.DeactivateConsole()
}

func consoleFin() {
// TODO wait (blink())
// TODO fin (blink())
  c := actualC
  finished = true
  time.Msleep (250) // provisorial
  c.cursorShape = shape.Off
  print (ClearScreen + home)
}

var
  initialized bool

func MaxResC() (uint, uint) {
  if framebufferOk() {
    return width, height
  }
  return 0, 0
}

func framebuffer() (x, y, b uint, fb obj.Stream) {
  var xc, yc, bc, ac C.int
  f := C.framebuffer (&xc, &yc, &bc, &ac)
  x, y, b = uint(xc), uint(yc), uint(bc)
  h := (*reflect.SliceHeader)((unsafe.Pointer(&fb)))
  m := int(x * y * (b / 8))
  h.Cap, h.Len, h.Data = m, m, uintptr(f)
  return
}

func framebufferOk() bool {
  if initialized {
    return true
  }
  initialized = true
  colbits := uint(0)
  width, height, colbits, fbmem = framebuffer()
  if colbits < 24 { ker.Panic ("µU does not support less than 24 bits per pixel") }
  if fbmem == nil {
    return false
  }
  fullScreen = mode.ModeOf (width, height)
  if mode.Wd (fullScreen) != width || mode.Ht (fullScreen) != height { ker.Panic ("fullScreen bug") }
  colourdepth = colbits / 8
  fbmemsize = width * height * colourdepth
  if uint(len (fbmem)) != fbmemsize {
    ker.Panic ("len (fbmem) == " + strconv.Itoa(len(fbmem)) +
               " != fbmemsize == " + strconv.Itoa(int(fbmemsize)))
  }
  fbcop = make(obj.Stream, fbmemsize)
  emptyBackground = make(obj.Stream, fbmemsize)
  ker.ConsoleInit()
  ker.SetAction (syscall.SIGUSR1, consoleOff)
  ker.SetAction (syscall.SIGUSR2, consoleOn)
  ker.InstallTerm (consoleFin)
  go ker.CatchSignals()
  initConsoleFonts()
  print (esc1 + "2J" + esc1 + "?1c" + esc1 + "?25l")
  visible = true
  return true
}

func (X *console) Go (m int, draw func(), ox, oy, oz, fx, fy, fz, tx, ty, tz float64) {
  ker.Panic ("the method Go does not work on a console")
}

////////////////////////////////////////////////////////////////////////

func (X *console) init_(x, y uint) {
  actualC = X
  mouseConsole = X
  X.x, X.y = int(x), int(y)
  X.cF, X.cB = col.StartCols()
  X.cFA, X.cBA = col.StartColsA()
  if ! framebufferOk() {
    ker.Panic ("µU does not support far tty-consoles")
  }
  if ! X.ok() { a, b, c := strconv.Itoa(X.x), strconv.Itoa(int(X.wd)), strconv.Itoa (int(width))
                d, e, f := strconv.Itoa(X.y), strconv.Itoa(int(X.ht)), strconv.Itoa (int(height))
    ker.Panic ("new console too large: " + a + " + " + b + " > " + c + " or " +
                                           d + " + " + e + " > " + f)
  }
  X.archive = make(obj.Stream, fbmemsize)
  X.shadow = make([]obj.Stream, X.ht)
  for i := 0; i < int(X.ht); i++ {
    X.shadow[i] = make(obj.Stream, X.wd * colourdepth)
  }
  X.initMouse()
  X.SetLinewidth (linewd.Thin)
  wm, _ := MaxResC()
  if wm > mode.Wd (mode.UHD) {
    X.SetFontsize (font.Huge)
  } else {
    X.SetFontsize (font.Normal)
  }
  X.Transparence (false)
  X.Colours (col.StartCols())
  X.ScrColours (col.StartCols())
  X.polygon = make([][]bool, X.wd)
  X.done = make([][]bool, X.wd)
  for i := 0; i < int(X.wd); i++ {
    X.polygon[i] = make([]bool, X.ht)
    X.done[i] = make([]bool, X.ht)
  }
  X.ScrColours (X.cF, X.cB)
  X.Cls()
  X.SetFontsize (font.Normal)
  X.doBlink()
  X.Cls()
}

func NewC (x, y uint, m mode.Mode) Screen {
  X := new(console)
  X.wd, X.ht = mode.Wd(m), mode.Ht(m)
  X.Mode = m
  X.init_(x, y)
  return X
}

func NewWHC (x, y, w, h uint) Screen {
  X := new(console)
  X.wd, X.ht = w, h
  X.Mode = mode.None
  X.init_(x, y)
  return X
}

func NewMaxC() Screen {
  return NewC (0, 0, mode.ModeOf (MaxResC()))
}

func MaxModeC() mode.Mode {
  width, height = MaxResC()
  return mode.ModeOf (width, height)
}

// func MaxResC() (uint, uint) {
//   return mode.Res (maxMode())
// }

func OkC (m mode.Mode) bool {
  fullScreen = MaxModeC()
  return mode.Wd (m) <= mode.Wd (fullScreen) &&
         mode.Ht (m) <= mode.Ht (fullScreen)
}
