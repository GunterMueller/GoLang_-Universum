package scr

// (c) Christian Maurer   v. 211127 - license see µU.go

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

// consolefonts ////////////////////////////////////////////////////////

var (
  f07[256][ 7]string // ter-d16b
  f10[256][10]string // ter-d18b
  f16[256][16]string // ter-d20b
  f24[256][24]string // ter-d24b
  f28[256][28]string // ter-d28b
  f32[256][32]string // ter-d32b
)

func (X *console) pointed (s font.Size, b byte, l, c uint) bool {
  switch s {
  case font.Tiny:
    return f07[b][l][c] != ' '
  case font.Small:
    return f10[b][l][c] != ' '
  case font.Normal:
    break
  case font.Big:
    return f24[b][l][c] != ' '
  case font.Large:
    return f28[b][l][c] != ' '
  case font.Huge:
    return f32[b][l][c] != ' '
  }
  return f16[b][l][c] != ' '
}

func initConsoleFonts() {
  for b := 0; b <= 255; b++ {
    for l := 0; l < 7; l++ {
      f07[b][l] = "     "
    }
  }

  f07[ 33][0] = "  !  "
  f07[ 33][1] = "  !  "
  f07[ 33][2] = "  !  "
  f07[ 33][3] = "  !  "
  f07[ 33][4] = "     "
  f07[ 33][5] = "  !  "

  f07[ 34][0] = " * * "
  f07[ 34][1] = " * * "

  f07[ 35][0] = " # # "
  f07[ 35][1] = "#####"
  f07[ 35][2] = " # # "
  f07[ 35][3] = " # # "
  f07[ 35][4] = "#####"
  f07[ 35][5] = " # # "

  f07[ 36][0] = "  $  "
  f07[ 36][1] = " $$$$"
  f07[ 36][2] = "$ $  "
  f07[ 36][3] = " $$$ "
  f07[ 36][4] = "  $ $"
  f07[ 36][5] = "$$$$ "
  f07[ 36][6] = "  $  "

  f07[ 37][0] = "%  % "
  f07[ 37][1] = "   % "
  f07[ 37][2] = "  %  "
  f07[ 37][3] = " %   "
  f07[ 37][4] = "%    "
  f07[ 37][5] = "%  % "

  f07[ 38][0] = " & & "
  f07[ 38][1] = "& & &"
  f07[ 38][2] = "& & &"
  f07[ 38][3] = " && &"
  f07[ 38][4] = "& & &"
  f07[ 38][5] = " & &&"

  f07[ 39][0] = "  '  "
  f07[ 39][1] = " '   "

  f07[ 40][0] = "  (  "
  f07[ 40][1] = " (   "
  f07[ 40][2] = " (   "
  f07[ 40][3] = " (   "
  f07[ 40][4] = " (   "
  f07[ 40][5] = "  (  "

  f07[ 41][0] = "  )  "
  f07[ 41][1] = "   ) "
  f07[ 41][2] = "   ) "
  f07[ 41][3] = "   ) "
  f07[ 41][4] = "   ) "
  f07[ 41][5] = "  )  "

  f07[ 42][1] = "*  * "
  f07[ 42][2] = " **  "
  f07[ 42][3] = "**** "
  f07[ 42][4] = " **  "
  f07[ 42][5] = "*  * "

  f07[ 43][1] = "  +  "
  f07[ 43][2] = "  +  "
  f07[ 43][3] = "+++++"
  f07[ 43][4] = "  +  "
  f07[ 43][5] = "  +  "

  f07[ 44][5] = "  ,  "
  f07[ 44][6] = " ,   "

  f07[ 45][3] = "---- "

  f07[ 46][5] = "  .  "

  f07[ 47][1] = "    /"
  f07[ 47][2] = "   / "
  f07[ 47][3] = "  /  "
  f07[ 47][4] = " /   "
  f07[ 47][5] = "/    "

  f07[ 48][0] = " 00  "
  f07[ 48][1] = "0  0 "
  f07[ 48][2] = "0  0 "
  f07[ 48][3] = "0  0 "
  f07[ 48][4] = "0  0 "
  f07[ 48][5] = " 00  "

  f07[ 49][0] = "  1  "
  f07[ 49][1] = " 11  "
  f07[ 49][2] = "1 1  "
  f07[ 49][3] = "  1  "
  f07[ 49][4] = "  1  "
  f07[ 49][5] = " 111 "

  f07[ 50][0] = " 22  "
  f07[ 50][1] = "2  2 "
  f07[ 50][2] = "   2 "
  f07[ 50][3] = "  2  "
  f07[ 50][4] = " 2   "
  f07[ 50][5] = "2222 "

  f07[ 51][0] = "3333 "
  f07[ 51][1] = "   3 "
  f07[ 51][2] = " 33  "
  f07[ 51][3] = "   3 "
  f07[ 51][4] = "3  3 "
  f07[ 51][5] = " 33  "

  f07[ 52][0] = "  4  "
  f07[ 52][1] = " 44  "
  f07[ 52][2] = "4 4  "
  f07[ 52][3] = "4444 "
  f07[ 52][4] = "  4  "
  f07[ 52][5] = "  4  "

  f07[ 53][0] = "5555 "
  f07[ 53][1] = "5    "
  f07[ 53][2] = "555  "
  f07[ 53][3] = "   5 "
  f07[ 53][4] = "5  5 "
  f07[ 53][5] = " 55  "

  f07[ 54][0] = " 66  "
  f07[ 54][1] = "6    "
  f07[ 54][2] = "666  "
  f07[ 54][3] = "6  6 "
  f07[ 54][4] = "6  6 "
  f07[ 54][5] = " 66  "

  f07[ 55][0] = "7777 "
  f07[ 55][1] = "   7 "
  f07[ 55][2] = "  7  "
  f07[ 55][3] = "  7  "
  f07[ 55][4] = " 7   "
  f07[ 55][5] = " 7   "

  f07[ 56][0] = " 88  "
  f07[ 56][1] = "8  8 "
  f07[ 56][2] = " 88  "
  f07[ 56][3] = "8  8 "
  f07[ 56][4] = "8  8 "
  f07[ 56][5] = " 88  "

  f07[ 57][0] = " 99  "
  f07[ 57][1] = "9  9 "
  f07[ 57][2] = " 999 "
  f07[ 57][3] = "   9 "
  f07[ 57][4] = "   9 "
  f07[ 57][5] = " 99  "

  f07[ 58][2] = "  :  "
  f07[ 58][3] = "     "
  f07[ 58][4] = "  :  "

  f07[ 59][3] = "  ,  "
  f07[ 59][4] = "  ,  "
  f07[ 59][5] = " ,,  "

  f07[ 60][1] = "   < "
  f07[ 60][2] = "  <  "
  f07[ 60][3] = " <   "
  f07[ 60][4] = "  <  "
  f07[ 60][5] = "   < "

  f07[ 61][2] = "==== "
  f07[ 61][3] = "     "
  f07[ 61][4] = "==== "

  f07[ 62][1] = " >   "
  f07[ 62][2] = "  >  "
  f07[ 62][3] = "   > "
  f07[ 62][4] = "  >  "
  f07[ 62][5] = " >   "

  f07[ 63][0] = "???  "
  f07[ 63][1] = "   ? "
  f07[ 63][2] = "  ?  "
  f07[ 63][3] = " ??  "
  f07[ 63][4] = "     "
  f07[ 63][5] = " ?   "

  f07[ 64][0] = " @@  "
  f07[ 64][1] = "@ @@ "
  f07[ 64][2] = "@@@@ "
  f07[ 64][3] = "@    "
  f07[ 64][4] = "@  @ "
  f07[ 64][5] = " @@@ "

  f07[ 65][0] = " AA  "
  f07[ 65][1] = "A  A "
  f07[ 65][2] = "A  A "
  f07[ 65][3] = "AAAA "
  f07[ 65][4] = "A  A "
  f07[ 65][5] = "A  A "

  f07[196][0] = "A  A "
  f07[196][1] = " AA  "
  f07[196][2] = "A  A "
  f07[196][3] = "AAAA "
  f07[196][4] = "A  A "
  f07[196][5] = "A  A "

  f07[ 66][0] = "BBB  "
  f07[ 66][1] = "B  B "
  f07[ 66][2] = "BBB  "
  f07[ 66][3] = "B  B "
  f07[ 66][4] = "B  B "
  f07[ 66][5] = "BBB  "

  f07[ 67][0] = " CC  "
  f07[ 67][1] = "C  C "
  f07[ 67][2] = "C    "
  f07[ 67][3] = "C    "
  f07[ 67][4] = "C  C "
  f07[ 67][5] = " CC  "

  f07[ 68][0] = "DDD  "
  f07[ 68][1] = "D  D "
  f07[ 68][2] = "D  D "
  f07[ 68][3] = "D  D "
  f07[ 68][4] = "D  D "
  f07[ 68][5] = "DDD  "

  f07[ 69][0] = "EEEE "
  f07[ 69][1] = "E    "
  f07[ 69][2] = "EEE  "
  f07[ 69][3] = "E    "
  f07[ 69][4] = "E    "
  f07[ 69][5] = "EEEE "

  f07[ 70][0] = "FFFF "
  f07[ 70][1] = "F    "
  f07[ 70][2] = "FFF  "
  f07[ 70][3] = "F    "
  f07[ 70][4] = "F    "
  f07[ 70][5] = "F    "

  f07[ 71][0] = " GG  "
  f07[ 71][1] = "G  G "
  f07[ 71][2] = "G    "
  f07[ 71][3] = "G GG "
  f07[ 71][4] = "G  G "
  f07[ 71][5] = " GGG "

  f07[ 72][0] = "H  H "
  f07[ 72][1] = "H  H "
  f07[ 72][2] = "HHHH "
  f07[ 72][3] = "H  H "
  f07[ 72][4] = "H  H "
  f07[ 72][5] = "H  H "

  f07[ 73][0] = " III "
  f07[ 73][1] = "  I  "
  f07[ 73][2] = "  I  "
  f07[ 73][3] = "  I  "
  f07[ 73][4] = "  I  "
  f07[ 73][5] = " III "

  f07[ 74][0] = "JJJJ "
  f07[ 74][1] = "   J "
  f07[ 74][2] = "   J "
  f07[ 74][3] = "   J "
  f07[ 74][4] = "J  J "
  f07[ 74][5] = " JJ  "

  f07[ 75][0] = "K  K "
  f07[ 75][1] = "K K  "
  f07[ 75][2] = "KK   "
  f07[ 75][3] = "KK   "
  f07[ 75][4] = "K K  "
  f07[ 75][5] = "K  K "

  f07[ 76][0] = "L    "
  f07[ 76][1] = "L    "
  f07[ 76][2] = "L    "
  f07[ 76][3] = "L    "
  f07[ 76][4] = "L    "
  f07[ 76][5] = "LLLL "

  f07[ 77][0] = "M  MM"
  f07[ 77][1] = "MMMM "
  f07[ 77][2] = "MMMM "
  f07[ 77][3] = "M  M "
  f07[ 77][4] = "M  M "
  f07[ 77][5] = "M  M "

  f07[ 78][0] = "N  N "
  f07[ 78][1] = "NN N "
  f07[ 78][2] = "NNNN "
  f07[ 78][3] = "N NN "
  f07[ 78][4] = "N  N "
  f07[ 78][5] = "N  N "

  f07[ 79][0] = " OO  "
  f07[ 79][1] = "O  O "
  f07[ 79][2] = "O  O "
  f07[ 79][3] = "O  O "
  f07[ 79][4] = "O  O "
  f07[ 79][5] = " OO  "

  f07[214][0] = "O  O "
  f07[214][1] = " OO  "
  f07[214][2] = "O  O "
  f07[214][3] = "O  O "
  f07[214][4] = "O  O "
  f07[214][5] = " OO  "

  f07[ 80][0] = "PPP  "
  f07[ 80][1] = "P  P "
  f07[ 80][2] = "P  P "
  f07[ 80][3] = "PPP  "
  f07[ 80][4] = "P    "
  f07[ 80][5] = "P    "

  f07[ 81][0] = " QQ  "
  f07[ 81][1] = "Q  Q "
  f07[ 81][2] = "Q  Q "
  f07[ 81][3] = "Q  Q "
  f07[ 81][4] = "QQ Q "
  f07[ 81][5] = " QQ  "
  f07[ 81][6] = "  QQ "

  f07[ 82][0] = "RRR  "
  f07[ 82][1] = "R  R "
  f07[ 82][2] = "R  R "
  f07[ 82][3] = "RRR  "
  f07[ 82][4] = "R R  "
  f07[ 82][5] = "R  R "

  f07[ 83][0] = " SS  "
  f07[ 83][1] = "S  S "
  f07[ 83][2] = " S   "
  f07[ 83][3] = "  S  "
  f07[ 83][4] = "S  S "
  f07[ 83][5] = " SS  "

  f07[ 84][0] = " TTT "
  f07[ 84][1] = "  T  "
  f07[ 84][2] = "  T  "
  f07[ 84][3] = "  T  "
  f07[ 84][4] = "  T  "
  f07[ 84][5] = "  T  "

  f07[ 85][0] = "U  U "
  f07[ 85][1] = "U  U "
  f07[ 85][2] = "U  U "
  f07[ 85][3] = "U  U "
  f07[ 85][4] = "U  U "
  f07[ 85][5] = " UU  "

  f07[220][0] = "U  U "
  f07[220][1] = "     "
  f07[220][2] = "U  U "
  f07[220][3] = "U  U "
  f07[220][4] = "U  U "
  f07[220][5] = " UU  "

  f07[ 86][0] = "V  V "
  f07[ 86][1] = "V  V "
  f07[ 86][2] = "V  V "
  f07[ 86][3] = "V  V "
  f07[ 86][4] = " VV  "
  f07[ 86][5] = " VV  "

  f07[ 87][0] = "W  W "
  f07[ 87][1] = "W  W "
  f07[ 87][2] = "W  W "
  f07[ 87][3] = "WWWW "
  f07[ 87][4] = "WWWW "
  f07[ 87][5] = "W  W "

  f07[ 88][0] = "X  X "
  f07[ 88][1] = "X  X "
  f07[ 88][2] = " XX  "
  f07[ 88][3] = " XX  "
  f07[ 88][4] = "X  X "
  f07[ 88][5] = "X  X "

  f07[ 89][0] = " Y Y "
  f07[ 89][1] = " Y Y "
  f07[ 89][2] = " Y Y "
  f07[ 89][3] = "  Y  "
  f07[ 89][4] = "  Y  "
  f07[ 89][5] = "  Y  "

  f07[ 90][0] = "ZZZZ "
  f07[ 90][1] = "   Z "
  f07[ 90][2] = "  Z  "
  f07[ 90][3] = " Z   "
  f07[ 90][4] = "Z    "
  f07[ 90][5] = "ZZZZ "

  f07[ 91][0] = " [[[ "
  f07[ 91][1] = " [   "
  f07[ 91][2] = " [   "
  f07[ 91][3] = " [   "
  f07[ 91][4] = " [   "
  f07[ 91][5] = " [[[ "

  f07[ 92][1] = "/    "
  f07[ 92][2] = " /   "
  f07[ 92][3] = "  /  "
  f07[ 92][4] = "   / "

  f07[ 93][0] = " ]]] "
  f07[ 93][1] = "   ] "
  f07[ 93][2] = "   ] "
  f07[ 93][3] = "   ] "
  f07[ 93][4] = "   ] "
  f07[ 93][5] = " ]]] "

  f07[ 94][0] = "  ^  "
  f07[ 94][1] = " ^ ^ "

  f07[ 95][5] = "____ "

  f07[ 96][0] = " `   "
  f07[ 96][1] = "  `  "

  f07[ 97][2] = " aaa "
  f07[ 97][3] = "a  a "
  f07[ 97][4] = "a  a "
  f07[ 97][5] = " aaa "

  f07[228][0] = " a a "
  f07[228][1] = "     "
  f07[228][2] = " aaa "
  f07[228][3] = "a  a "
  f07[228][4] = "a  a "
  f07[228][5] = " aaa "

  f07[ 98][0] = "b    "
  f07[ 98][1] = "b    "
  f07[ 98][2] = "bbb  "
  f07[ 98][3] = "b  b "
  f07[ 98][4] = "b  b "
  f07[ 98][5] = "bbb  "

  f07[ 99][2] = " cc  "
  f07[ 99][3] = "c    "
  f07[ 99][4] = "c    "
  f07[ 99][5] = " ccc "

  f07[100][0] = "   d "
  f07[100][1] = "   d "
  f07[100][2] = " ddd "
  f07[100][3] = "d  d "
  f07[100][4] = "d  d "
  f07[100][5] = " ddd "

  f07[101][2] = " ee  "
  f07[101][3] = "eeee "
  f07[101][4] = "e    "
  f07[101][5] = " eee "

  f07[102][0] = "  ff "
  f07[102][1] = " f   "
  f07[102][2] = "fff  "
  f07[102][3] = " f   "
  f07[102][4] = " f   "
  f07[102][5] = " f   "
  f07[102][6] = " f   "

  f07[103][2] = " ggg "
  f07[103][3] = "g  g "
  f07[103][4] = " ggg "
  f07[103][5] = "   g "
  f07[103][6] = "ggg  "

  f07[104][0] = "h    "
  f07[104][1] = "h    "
  f07[104][2] = "hhh  "
  f07[104][3] = "h  h "
  f07[104][4] = "h  h "
  f07[104][5] = "h  h "

  f07[105][0] = " i   "
  f07[105][1] = "     "
  f07[105][2] = " i   "
  f07[105][3] = " i   "
  f07[105][4] = " i   "
  f07[105][5] = "  ii "

  f07[106][0] = "  j  "
  f07[106][1] = "     "
  f07[106][2] = "  j  "
  f07[106][3] = "  j  "
  f07[106][4] = "  j  "
  f07[106][5] = "  j  "
  f07[106][6] = " jj  "

  f07[107][0] = "k    "
  f07[107][1] = "k    "
  f07[107][2] = "k  k "
  f07[107][3] = "kkk  "
  f07[107][4] = "k  k "
  f07[107][5] = "k  k "

  f07[108][0] = " l   "
  f07[108][1] = " l   "
  f07[108][2] = " l   "
  f07[108][3] = " l   "
  f07[108][4] = " l   "
  f07[108][5] = "  ll "

  f07[109][2] = "mmmmm"
  f07[109][3] = "m m m"
  f07[109][4] = "m m m"
  f07[109][5] = "m m m"

  f07[110][2] = "nnn  "
  f07[110][3] = "n  n "
  f07[110][4] = "n  n "
  f07[110][5] = "n  n "

  f07[111][2] = " oo  "
  f07[111][3] = "o  o "
  f07[111][4] = "o  o "
  f07[111][5] = " oo  "

  f07[246][0] = "o  o "
  f07[246][1] = "     "
  f07[246][2] = " oo  "
  f07[246][3] = "o  o "
  f07[246][4] = "o  o "
  f07[246][5] = " oo  "

  f07[112][2] = "ppp  "
  f07[112][3] = "p  p "
  f07[112][4] = "ppp  "
  f07[112][5] = "p    "
  f07[112][6] = "p    "

  f07[113][2] = " qqq "
  f07[113][3] = "q  q "
  f07[113][4] = " qqq "
  f07[113][5] = "   q "
  f07[113][6] = "   q "

  f07[114][2] = " rr  "
  f07[114][3] = " r   "
  f07[114][4] = " r   "
  f07[114][5] = " r   "

  f07[115][2] = " sss "
  f07[115][3] = "ss   "
  f07[115][4] = "  ss "
  f07[115][5] = "sss  "

  f07[116][0] = " t   "
  f07[116][1] = " t   "
  f07[116][2] = "ttt  "
  f07[116][3] = " t   "
  f07[116][4] = " t   "
  f07[116][5] = "  tt "

  f07[117][2] = "u  u "
  f07[117][3] = "u  u "
  f07[117][4] = "u  u "
  f07[117][5] = " uuu "

  f07[252][0] = "u  u "
  f07[252][1] = "     "
  f07[252][2] = "u  u "
  f07[252][3] = "u  u "
  f07[252][4] = "u  u "
  f07[252][5] = " uuu "

  f07[118][2] = " v v "
  f07[118][3] = " v v "
  f07[118][4] = " v v "
  f07[118][5] = "  v  "

  f07[119][2] = "w w w"
  f07[119][3] = "w w w"
  f07[119][4] = "w w w"
  f07[119][5] = " w w "

  f07[120][2] = "x  x "
  f07[120][3] = " xx  "
  f07[120][4] = " xx  "
  f07[120][5] = "x  x "

  f07[121][2] = "y  y "
  f07[121][3] = "y  y "
  f07[121][4] = " yyy "
  f07[121][5] = "   y "
  f07[121][6] = "yyy  "

  f07[122][2] = "zzzz "
  f07[122][3] = "  z  "
  f07[122][3] = " z   "
  f07[122][4] = "z    "
  f07[122][5] = "zzzz "

  f07[123][0] = "  {{ "
  f07[123][1] = " {   "
  f07[123][2] = " {   "
  f07[123][3] = "{    "
  f07[123][4] = " {   "
  f07[123][5] = " {   "
  f07[123][6] = "  {{ "

  f07[124][0] = "  |  "
  f07[124][1] = "  |  "
  f07[124][2] = "  |  "
  f07[124][3] = "  |  "
  f07[124][4] = "  |  "
  f07[124][5] = "  |  "

  f07[125][0] = "}}   "
  f07[125][1] = "  }  "
  f07[125][2] = "  }  "
  f07[125][3] = "   } "
  f07[125][4] = "  }  "
  f07[125][5] = "  }  "
  f07[125][6] = "}}   "

  f07[126][0] = " ~~ ~"
  f07[126][1] = "~ ~~ "

  f07[164][0] = "  eee"
  f07[164][1] = " e   "
  f07[164][2] = "eeee "
  f07[164][3] = " e   "
  f07[164][4] = "eeee "
  f07[164][5] = " e   "
  f07[164][6] = "  eee"

  f07[167][0] = " pp  "
  f07[167][1] = "p    "
  f07[167][2] = " pp  "
  f07[167][3] = "p  p "
  f07[167][4] = " pp  "
  f07[167][5] = "   p "
  f07[167][6] = " pp  "

  f07[176][0] = "  o  "
  f07[176][1] = " o o "
  f07[176][2] = "  o  "

  f07[181][2] = "m  m "
  f07[181][3] = "m  m "
  f07[181][4] = "mmm m"
  f07[181][5] = "m    "
  f07[181][6] = "m    "

  f07[223][0] = " ss  "
  f07[223][1] = "s  s "
  f07[223][2] = "s s  "
  f07[223][3] = "s  s "
  f07[223][4] = "s  s "
  f07[223][5] = "sss  "
  f07[223][6] = "s    "

  for b := 0; b <= 255; b++ {
    for l := 0; l < 10; l++ {
      f10[b][l] = "      "
    }
  }

  f10[ 33][1] = "  !   "
  f10[ 33][2] = " !!!  "
  f10[ 33][3] = " !!!  "
  f10[ 33][4] = "  !   "
  f10[ 33][5] = "  !   "
  f10[ 33][6] = "      "
  f10[ 33][7] = "  !   "

  f10[ 34][0] = " * *  "
  f10[ 34][1] = " * *  "

  f10[ 35][1] = " != #  "
  f10[ 35][2] = " != #  "
  f10[ 35][3] = "##### "
  f10[ 35][4] = " != #  "
  f10[ 35][5] = "##### "
  f10[ 35][6] = " != #  "
  f10[ 35][7] = " != #  "

  f10[ 36][1] = "  $   "
  f10[ 36][2] = " $$$$ "
  f10[ 36][3] = "$ $   "
  f10[ 36][4] = " $$$  "
  f10[ 36][5] = "  $ $ "
  f10[ 36][6] = "$$$$  "
  f10[ 36][7] = "  $   "

  f10[ 37][1] = "%%  % "
  f10[ 37][2] = "%%  % "
  f10[ 37][3] = "   %  "
  f10[ 37][4] = "  %   "
  f10[ 37][5] = " %    "
  f10[ 37][6] = "%  %% "
  f10[ 37][7] = "%  %% "

  f10[ 38][1] = " &&   "
  f10[ 38][2] = "&   &  &   "
  f10[ 38][3] = "&   &  &   "
  f10[ 38][4] = " &&   "
  f10[ 38][5] = "&  &  &  & "
  f10[ 38][6] = "&   &  &   "
  f10[ 38][7] = " &&  &  &  "

  f10[ 39][0] = "  ''  "
  f10[ 39][1] = "  ''  "
  f10[ 39][2] = "  '   "

  f10[ 40][0] = "   (  "
  f10[ 40][1] = "  (   "
  f10[ 40][2] = " (    "
  f10[ 40][3] = " (    "
  f10[ 40][4] = " (    "
  f10[ 40][5] = " (    "
  f10[ 40][6] = "  (   "
  f10[ 40][7] = "   (  "

  f10[ 41][0] = "  )   "
  f10[ 41][1] = "   )  "
  f10[ 41][2] = "    ) "
  f10[ 41][3] = "    ) "
  f10[ 41][4] = "    ) "
  f10[ 41][5] = "    ) "
  f10[ 41][6] = "   )  "
  f10[ 41][7] = "  )   "

  f10[ 42][2] = "*   * "
  f10[ 42][3] = " * *  "
  f10[ 42][4] = "***** "
  f10[ 42][5] = " * *  "
  f10[ 42][6] = "*   * "

  f10[ 43][2] = "  +   "
  f10[ 43][3] = "  +   "
  f10[ 43][4] = "+++++ "
  f10[ 43][5] = "  +   "
  f10[ 43][6] = "  +   "

  f10[ 44][6] = "  ,,  "
  f10[ 44][7] = "  ,,  "
  f10[ 44][8] = "  ,,  "
  f10[ 44][9] = "  ,   "

  f10[ 45][4] = "----- "

  f10[ 46][6] = "  ..  "
  f10[ 46][7] = "  ..  "

  f10[ 47][1] = "    / "
  f10[ 47][2] = "    / "
  f10[ 47][3] = "   /  "
  f10[ 47][4] = "  /   "
  f10[ 47][5] = " /    "
  f10[ 47][6] = "/     "
  f10[ 47][7] = "/     "

  f10[ 48][1] = " 000  "
  f10[ 48][2] = "0   0 "
  f10[ 48][3] = "0  00 "
  f10[ 48][4] = "0 0 0 "
  f10[ 48][5] = "00  0 "
  f10[ 48][6] = "0   0 "
  f10[ 48][7] = " 000  "

  f10[ 49][1] = "   1  "
  f10[ 49][2] = "  11  "
  f10[ 49][3] = " 111  "
  f10[ 49][4] = "11 1  "
  f10[ 49][5] = "   1  "
  f10[ 49][6] = "   1  "
  f10[ 49][7] = "   1  "

  f10[ 50][1] = " 222  "
  f10[ 50][2] = "2   2 "
  f10[ 50][3] = "    2 "
  f10[ 50][4] = "  22  "
  f10[ 50][5] = " 2    "
  f10[ 50][6] = "2     "
  f10[ 50][7] = "22222 "

  f10[ 51][1] = " 333  "
  f10[ 51][2] = "3   3 "
  f10[ 51][3] = "    3 "
  f10[ 51][4] = "  33  "
  f10[ 51][5] = "    3 "
  f10[ 51][6] = "3   3 "
  f10[ 51][7] = " 333  "

  f10[ 52][1] = "4  4  "
  f10[ 52][2] = "4  4  "
  f10[ 52][3] = "4  4  "
  f10[ 52][4] = "44444 "
  f10[ 52][5] = "   4  "
  f10[ 52][6] = "   4  "
  f10[ 52][7] = "   4  "

  f10[ 53][1] = "55555 "
  f10[ 53][2] = "5     "
  f10[ 53][3] = "5     "
  f10[ 53][4] = "5555  "
  f10[ 53][5] = "    5 "
  f10[ 53][6] = "5   5 "
  f10[ 53][7] = " 555  "

  f10[ 54][1] = "  66  "
  f10[ 54][2] = " 6    "
  f10[ 54][3] = "6     "
  f10[ 54][4] = "6666  "
  f10[ 54][5] = "6   6 "
  f10[ 54][6] = "6   6 "
  f10[ 54][7] = " 666  "

  f10[ 55][1] = "77777 "
  f10[ 55][2] = "    7 "
  f10[ 55][3] = "   7  "
  f10[ 55][4] = " 7777 "
  f10[ 55][5] = "  7   "
  f10[ 55][6] = " 7    "
  f10[ 55][7] = "7     "

  f10[ 56][1] = " 888  "
  f10[ 56][2] = "8   8 "
  f10[ 56][3] = "8   8 "
  f10[ 56][4] = " 888  "
  f10[ 56][5] = "8   8 "
  f10[ 56][6] = "8   8 "
  f10[ 56][7] = " 888  "

  f10[ 57][1] = " 999  "
  f10[ 57][2] = "9   9 "
  f10[ 57][3] = "9   9 "
  f10[ 57][4] = " 9999 "
  f10[ 57][5] = "    9 "
  f10[ 57][6] = "   9  "
  f10[ 57][7] = " 99   "

  f10[ 58][2] = "  ::  "
  f10[ 58][3] = "  ::  "
  f10[ 58][4] = "      "
  f10[ 58][5] = "  ::  "
  f10[ 58][6] = "  ::  "

  f10[ 59][4] = "  ,,  "
  f10[ 59][5] = "  ,,  "
  f10[ 59][6] = "      "
  f10[ 59][7] = "  ,,  "
  f10[ 59][8] = "  ,,  "
  f10[ 59][9] = "  ,   "

  f10[ 60][2] = "   << "
  f10[ 60][3] = "  <<  "
  f10[ 60][4] = " <<   "
  f10[ 60][5] = "  <<  "
  f10[ 60][6] = "   << "

  f10[ 61][3] = "===== "
  f10[ 61][4] = "      "
  f10[ 61][5] = "===== "

  f10[ 62][1] = "      "
  f10[ 62][2] = " >>   "
  f10[ 62][3] = "  >>  "
  f10[ 62][4] = "   >> "
  f10[ 62][5] = "  >>  "
  f10[ 62][6] = " >>   "
  f10[ 62][7] = "      "

  f10[ 63][1] = " ???  "
  f10[ 63][2] = "?   ? "
  f10[ 63][3] = "    ? "
  f10[ 63][4] = "   ?  "
  f10[ 63][5] = "  ?   "
  f10[ 63][6] = "      "
  f10[ 63][7] = "  ?   "

  f10[ 64][1] = " @@@  "
  f10[ 64][2] = "@   @ "
  f10[ 64][3] = "@ @@@ "
  f10[ 64][4] = "@ @ @ "
  f10[ 64][5] = "@ @@@ "
  f10[ 64][6] = "@     "
  f10[ 64][7] = " @@@@ "

  f10[ 65][1] = " AAA  "
  f10[ 65][2] = "A   A "
  f10[ 65][3] = "A   A "
  f10[ 65][4] = "AAAAA "
  f10[ 65][5] = "A   A "
  f10[ 65][6] = "A   A "
  f10[ 65][7] = "A   A "

  f10[196][0] = "A   A "
  f10[196][1] = " AAA  "
  f10[196][2] = "A   A "
  f10[196][3] = "A   A "
  f10[196][4] = "AAAAA "
  f10[196][5] = "A   A "
  f10[196][6] = "A   A "
  f10[196][7] = "A   A "

  f10[ 66][1] = "BBBB  "
  f10[ 66][2] = "B   B "
  f10[ 66][3] = "B   B "
  f10[ 66][4] = "BBBB  "
  f10[ 66][5] = "B   B "
  f10[ 66][6] = "B   B "
  f10[ 66][7] = "BBBB  "

  f10[ 67][1] = " CCC  "
  f10[ 67][2] = "C   C "
  f10[ 67][3] = "C     "
  f10[ 67][4] = "C     "
  f10[ 67][5] = "C     "
  f10[ 67][6] = "C   C "
  f10[ 67][7] = " CCC  "

  f10[ 68][1] = "DDD   "
  f10[ 68][2] = "D  D  "
  f10[ 68][3] = "D   D "
  f10[ 68][4] = "D   D "
  f10[ 68][5] = "D   D "
  f10[ 68][6] = "D  D  "
  f10[ 68][7] = "DDD   "

  f10[ 69][1] = "EEEEE "
  f10[ 69][2] = "E     "
  f10[ 69][3] = "E     "
  f10[ 69][4] = "EEE   "
  f10[ 69][5] = "E     "
  f10[ 69][6] = "E     "
  f10[ 69][7] = "EEEEE "

  f10[ 70][1] = "FFFFF "
  f10[ 70][2] = "F     "
  f10[ 70][3] = "F     "
  f10[ 70][4] = "FFF   "
  f10[ 70][5] = "F     "
  f10[ 70][6] = "F     "
  f10[ 70][7] = "F     "

  f10[ 71][1] = " GGG  "
  f10[ 71][2] = "G   G "
  f10[ 71][3] = "G     "
  f10[ 71][4] = "G GGG "
  f10[ 71][5] = "G   G "
  f10[ 71][6] = "G   G "
  f10[ 71][7] = " GGG  "

  f10[ 72][1] = "H   H "
  f10[ 72][2] = "H   H "
  f10[ 72][3] = "H   H "
  f10[ 72][4] = "HHHHH "
  f10[ 72][5] = "H   H "
  f10[ 72][6] = "H   H "
  f10[ 72][7] = "H   H "

  f10[ 73][1] = " III  "
  f10[ 73][2] = "  I   "
  f10[ 73][3] = "  I   "
  f10[ 73][4] = "  I   "
  f10[ 73][5] = "  I   "
  f10[ 73][6] = "  I   "
  f10[ 73][7] = " III  "

  f10[ 74][1] = "JJJJJ "
  f10[ 74][2] = "    J "
  f10[ 74][3] = "    J "
  f10[ 74][4] = "    J "
  f10[ 74][5] = "    J "
  f10[ 74][6] = "J   J "
  f10[ 74][7] = " JJJ  "

  f10[ 75][0] = "      "
  f10[ 75][1] = "K   K "
  f10[ 75][2] = "K  K  "
  f10[ 75][3] = "K K   "
  f10[ 75][4] = "KK    "
  f10[ 75][5] = "K K   "
  f10[ 75][6] = "K  K  "
  f10[ 75][7] = "K   K "

  f10[ 76][1] = "L     "
  f10[ 76][2] = "L     "
  f10[ 76][3] = "L     "
  f10[ 76][4] = "L     "
  f10[ 76][5] = "L     "
  f10[ 76][6] = "L     "
  f10[ 76][7] = "LLLLL "

  f10[ 77][1] = "M   M "
  f10[ 77][2] = "MM MM "
  f10[ 77][3] = "MMMMM "
  f10[ 77][4] = "M M M "
  f10[ 77][5] = "M M M "
  f10[ 77][6] = "M   M "
  f10[ 77][7] = "M   M "

  f10[ 78][1] = "N   N "
  f10[ 78][2] = "NN  N "
  f10[ 78][3] = "NNN N "
  f10[ 78][4] = "N NNN "
  f10[ 78][5] = "N  NN "
  f10[ 78][6] = "N   N "
  f10[ 78][7] = "N   N "

  f10[ 79][1] = " OOO  "
  f10[ 79][2] = "O   O "
  f10[ 79][3] = "O   O "
  f10[ 79][4] = "O   O "
  f10[ 79][5] = "O   O "
  f10[ 79][6] = "O   O "
  f10[ 79][7] = " OOO  "

  f10[214][0] = "O   O "
  f10[214][1] = " OOO  "
  f10[214][2] = "O   O "
  f10[214][3] = "O   O "
  f10[214][4] = "O   O "
  f10[214][5] = "O   O "
  f10[214][6] = "O   O "
  f10[214][7] = " OOO  "

  f10[ 80][1] = "PPPP  "
  f10[ 80][2] = "P   P "
  f10[ 80][3] = "P   P "
  f10[ 80][4] = "PPPP  "
  f10[ 80][5] = "P     "
  f10[ 80][6] = "P     "
  f10[ 80][7] = "P     "

  f10[ 81][1] = " QQQ  "
  f10[ 81][2] = "Q   Q "
  f10[ 81][3] = "Q   Q "
  f10[ 81][4] = "Q   Q "
  f10[ 81][5] = "QQ  Q "
  f10[ 81][6] = "Q Q Q "
  f10[ 81][7] = " QQQ  "
  f10[ 81][8] = "   QQ "

  f10[ 82][1] = "RRRR  "
  f10[ 82][2] = "R   R "
  f10[ 82][3] = "R   R "
  f10[ 82][4] = "RRRR  "
  f10[ 82][5] = "R R   "
  f10[ 82][6] = "R  R  "
  f10[ 82][7] = "R   R "

  f10[ 83][1] = " SSS  "
  f10[ 83][2] = "S   S "
  f10[ 83][3] = "S     "
  f10[ 83][4] = " SSS  "
  f10[ 83][5] = "    S "
  f10[ 83][6] = "S   S "
  f10[ 83][7] = " SSS  "

  f10[ 84][1] = "TTTTT "
  f10[ 84][2] = "  T   "
  f10[ 84][3] = "  T   "
  f10[ 84][4] = "  T   "
  f10[ 84][5] = "  T   "
  f10[ 84][6] = "  T   "
  f10[ 84][7] = "  T   "

  f10[ 85][1] = "U   U "
  f10[ 85][2] = "U   U "
  f10[ 85][3] = "U   U "
  f10[ 85][4] = "U   U "
  f10[ 85][5] = "U   U "
  f10[ 85][6] = "U   U "
  f10[ 85][7] = " UUU  "

  f10[220][0] = "U   U "
  f10[220][1] = "      "
  f10[220][2] = "U   U "
  f10[220][3] = "U   U "
  f10[220][4] = "U   U "
  f10[220][5] = "U   U "
  f10[220][6] = "U   U "
  f10[220][7] = " UUU  "

  f10[ 86][1] = "V   V "
  f10[ 86][2] = "V   V "
  f10[ 86][3] = "V   V "
  f10[ 86][4] = " V V  "
  f10[ 86][5] = " V V  "
  f10[ 86][6] = "  V   "
  f10[ 86][7] = "  V   "

  f10[ 87][1] = "W   W "
  f10[ 87][2] = "W   W "
  f10[ 87][3] = "W   W "
  f10[ 87][4] = "W W W "
  f10[ 87][5] = "W W W "
  f10[ 87][6] = "W W W "
  f10[ 87][7] = " W W  "

  f10[ 88][1] = "X   X "
  f10[ 88][2] = "X   X "
  f10[ 88][3] = " X X  "
  f10[ 88][4] = "  X   "
  f10[ 88][5] = " X X  "
  f10[ 88][6] = "X   X "
  f10[ 88][7] = "X   X "

  f10[ 89][1] = "Y   Y "
  f10[ 89][2] = "Y   Y "
  f10[ 89][3] = "Y   Y "
  f10[ 89][4] = " Y Y  "
  f10[ 89][5] = "  Y   "
  f10[ 89][6] = "  Y   "
  f10[ 89][7] = "  Y   "

  f10[ 90][1] = "ZZZZZ "
  f10[ 90][2] = "    Z "
  f10[ 90][3] = "   Z  "
  f10[ 90][4] = "  Z   "
  f10[ 90][5] = " Z    "
  f10[ 90][6] = "Z     "
  f10[ 90][7] = "ZZZZZ "

  f10[ 91][1] = "  [[[ "
  f10[ 91][2] = "  [   "
  f10[ 91][3] = "  [   "
  f10[ 91][4] = "  [   "
  f10[ 91][5] = "  [   "
  f10[ 91][6] = "  [   "
  f10[ 91][7] = "  [[[ "

  f10[ 92][2] = "/     "
  f10[ 92][3] = " /    "
  f10[ 92][4] = "  /   "
  f10[ 92][5] = "   /  "
  f10[ 92][6] = "    / "

  f10[ 93][1] = " ]]]  "
  f10[ 93][2] = "   ]  "
  f10[ 93][3] = "   ]  "
  f10[ 93][4] = "   ]  "
  f10[ 93][5] = "   ]  "
  f10[ 93][6] = "   ]  "
  f10[ 93][7] = " ]]]  "

  f10[ 94][0] = "  ^   "
  f10[ 94][1] = " ^ ^  "

  f10[ 95][7] = "_____ "

  f10[ 96][0] = "`     "
  f10[ 96][1] = " `    "
  f10[ 96][2] = "  `   "

  f10[ 97][3] = " aa a "
  f10[ 97][4] = "a  aa "
  f10[ 97][5] = "a   a "
  f10[ 97][6] = "a  aa "
  f10[ 97][7] = " aa a "

  f10[228][1] = "a   a "
  f10[228][2] = "      "
  f10[228][3] = " aa a "
  f10[228][4] = "a  aa "
  f10[228][5] = "a   a "
  f10[228][6] = "a  aa "
  f10[228][7] = " aa a "

  f10[ 98][1] = "b     "
  f10[ 98][2] = "b     "
  f10[ 98][3] = "b bb  "
  f10[ 98][4] = "bb  b "
  f10[ 98][5] = "b   b "
  f10[ 98][6] = "bb  b "
  f10[ 98][7] = "b bb  "

  f10[ 99][3] = " ccc  "
  f10[ 99][4] = "c     "
  f10[ 99][5] = "c     "
  f10[ 99][6] = "c   c "
  f10[ 99][7] = " ccc  "

  f10[100][1] = "    d "
  f10[100][2] = "    d "
  f10[100][3] = " dd d "
  f10[100][4] = "d  dd "
  f10[100][5] = "d   d "
  f10[100][6] = "d  dd "
  f10[100][7] = " dd d "

  f10[101][3] = " eee  "
  f10[101][4] = "e   e "
  f10[101][5] = "eeeee "
  f10[101][6] = "e     "
  f10[101][7] = " eeee "

  f10[102][1] = "  ff  "
  f10[102][2] = " f  f "
  f10[102][3] = " f    "
  f10[102][4] = "fffff "
  f10[102][5] = " f    "
  f10[102][6] = " f    "
  f10[102][7] = " f    "
  f10[102][8] = " f    "

  f10[103][3] = " ggg  "
  f10[103][4] = "g   g "
  f10[103][5] = "g   g "
  f10[103][6] = "g  gg "
  f10[103][7] = " gg g "
  f10[103][8] = "    g "
  f10[103][9] = " ggg  "

  f10[104][1] = "h     "
  f10[104][2] = "h     "
  f10[104][3] = "h hh  "
  f10[104][4] = "hh  h "
  f10[104][5] = "h   h "
  f10[104][6] = "h   h "
  f10[104][7] = "h   h "

  f10[105][1] = " i    "
  f10[105][2] = "      "
  f10[105][3] = " i    "
  f10[105][4] = " i    "
  f10[105][5] = " i    "
  f10[105][6] = " i  i "
  f10[105][7] = "  ii  "

  f10[106][1] = "   j  "
  f10[106][2] = "      "
  f10[106][3] = "   j  "
  f10[106][4] = "   j  "
  f10[106][5] = "   j  "
  f10[106][6] = "   j  "
  f10[106][7] = "   j  "
  f10[106][8] = "j  j  "
  f10[106][9] = " jj   "

  f10[107][1] = "k     "
  f10[107][2] = "k     "
  f10[107][3] = "k   k "
  f10[107][4] = "k  k  "
  f10[107][5] = "kkk   "
  f10[107][6] = "k  k  "
  f10[107][7] = "k   k "

  f10[108][1] = " l    "
  f10[108][2] = " l    "
  f10[108][3] = " l    "
  f10[108][4] = " l    "
  f10[108][5] = " l    "
  f10[108][6] = " l  l "
  f10[108][7] = "  ll  "

  f10[109][3] = "mm m  "
  f10[109][4] = "m m m "
  f10[109][5] = "m m m "
  f10[109][6] = "m m m "
  f10[109][7] = "m m m "

  f10[110][3] = "n nn  "
  f10[110][4] = "nn  n "
  f10[110][5] = "n   n "
  f10[110][6] = "n   n "
  f10[110][7] = "n   n "

  f10[111][3] = " ooo  "
  f10[111][4] = "o   o "
  f10[111][5] = "o   o "
  f10[111][6] = "o   o "
  f10[111][7] = " ooo  "

  f10[246][1] = "o   o "
  f10[246][2] = "      "
  f10[246][3] = " ooo  "
  f10[246][4] = "o   o "
  f10[246][5] = "o   o "
  f10[246][6] = "o   o "
  f10[246][7] = " ooo  "

  f10[112][3] = "p pp  "
  f10[112][4] = "pp  p "
  f10[112][5] = "p   p "
  f10[112][6] = "pp  p "
  f10[112][7] = "p pp  "
  f10[112][8] = "p     "
  f10[112][9] = "p     "

  f10[113][3] = " qq q "
  f10[113][4] = "q  qq "
  f10[113][5] = "q   q "
  f10[113][6] = "q  qq "
  f10[113][7] = " qq q "
  f10[113][8] = "    q "
  f10[113][9] = "    q "

  f10[114][3] = "r rr  "
  f10[114][4] = "rr  r "
  f10[114][5] = "r     "
  f10[114][6] = "r     "
  f10[114][7] = "r     "

  f10[115][3] = " sss  "
  f10[115][4] = "s     "
  f10[115][5] = " sss  "
  f10[115][6] = "    s "
  f10[115][7] = "ssss  "

  f10[116][1] = " t    "
  f10[116][2] = " t    "
  f10[116][3] = "ttttt "
  f10[116][4] = " t    "
  f10[116][5] = " t    "
  f10[116][6] = " t  t "
  f10[116][7] = "  tt  "

  f10[117][3] = "u   u "
  f10[117][4] = "u   u "
  f10[117][5] = "u   u "
  f10[117][6] = "u  uu "
  f10[117][7] = " uu u "

  f10[252][1] = "u   u "
  f10[252][2] = "      "
  f10[252][3] = "u   u "
  f10[252][4] = "u   u "
  f10[252][5] = "u   u "
  f10[252][6] = "u  uu "
  f10[252][7] = " uu u "

  f10[118][3] = "v   v "
  f10[118][4] = "v   v "
  f10[118][5] = "v   v "
  f10[118][6] = " v v  "
  f10[118][7] = "  v   "

  f10[119][3] = "w   w "
  f10[119][4] = "w   w "
  f10[119][5] = "w w w "
  f10[119][6] = "w w w "
  f10[119][7] = " w w  "

  f10[120][3] = "x   x "
  f10[120][4] = " x x  "
  f10[120][5] = "  x   "
  f10[120][6] = " x x  "
  f10[120][7] = "x   x "

  f10[121][3] = "y   y "
  f10[121][4] = "y   y "
  f10[121][5] = "y   y "
  f10[121][6] = "y  yy "
  f10[121][7] = " yy y "
  f10[121][8] = "    y "
  f10[121][9] = "yyyy  "

  f10[122][3] = "zzzzz "
  f10[122][4] = "   z  "
  f10[122][5] = "  z   "
  f10[122][6] = " z    "
  f10[122][7] = "zzzzz "

  f10[123][1] = "   {{ "
  f10[123][2] = "  {   "
  f10[123][3] = "  {   "
  f10[123][4] = " {    "
  f10[123][5] = "  {   "
  f10[123][6] = "  {   "
  f10[123][7] = "   {{ "

  f10[124][1] = "  |   "
  f10[124][2] = "  |   "
  f10[124][3] = "  |   "
  f10[124][4] = "  |   "
  f10[124][5] = "  |   "
  f10[124][6] = "  |   "
  f10[124][7] = "  |   "

  f10[125][1] = " }}   "
  f10[125][2] = "   }  "
  f10[125][3] = "   }  "
  f10[125][4] = "    } "
  f10[125][5] = "   }  "
  f10[125][6] = "   }  "
  f10[125][7] = " }}   "

  f10[126][0] = " ~~  ~"
  f10[126][1] = "~  ~~ "

  f10[164][1] = "   ee "
  f10[164][2] = "  e  e"
  f10[164][3] = "eeee  "
  f10[164][4] = " e    "
  f10[164][5] = "eeee  "
  f10[164][6] = "  e  e"
  f10[164][7] = "   ee "

  f10[167][0] = "  pp  "
  f10[167][1] = " p    "
  f10[167][2] = " pp   "
  f10[167][3] = "p  p  "
  f10[167][4] = "p   p "
  f10[167][5] = " p  p "
  f10[167][6] = "  pp  "
  f10[167][7] = "   p  "
  f10[167][8] = " pp   "

  f10[176][0] = "  oo  "
  f10[176][1] = " o  o "
  f10[176][2] = "  oo  "

  f10[181][3] = "m   m "
  f10[181][4] = "m   m "
  f10[181][5] = "m   m "
  f10[181][6] = "mm  m "
  f10[181][7] = "m mm m"
  f10[181][8] = "m     "
  f10[181][9] = "m     "

  f10[223][1] = " ss   "
  f10[223][2] = "s  s  "
  f10[223][3] = "s  s  "
  f10[223][4] = "s ss  "
  f10[223][5] = "s   s "
  f10[223][6] = "s   s "
  f10[223][7] = "s ss  "
  f10[223][8] = "s     "
  f10[223][9] = "s     "

  for b := 0; b <= 255; b++ {
    for l := 0; l < 16; l++ {
      f16[b][l] = "        "
    }
  }

  f16[ 33][ 2] = "   !!   "
  f16[ 33][ 3] = "  !!!!  "
  f16[ 33][ 4] = "  !!!!  "
  f16[ 33][ 5] = "  !!!!  "
  f16[ 33][ 6] = "   !!   "
  f16[ 33][ 7] = "   !!   "
  f16[ 33][ 8] = "   !!   "
  f16[ 33][ 9] = "        "
  f16[ 33][10] = "   !!   "
  f16[ 33][11] = "   !!   "

  f16[ 34][ 2] = " **  ** "
  f16[ 34][ 3] = " **  ** "
  f16[ 34][ 4] = " **  ** "
  f16[ 34][ 5] = "  *  *  "

  f16[ 35][ 3] = " ## ##  "
  f16[ 35][ 4] = " ## ##  "
  f16[ 35][ 5] = "####### "
  f16[ 35][ 6] = " ## ##  "
  f16[ 35][ 7] = " ## ##  "
  f16[ 35][ 8] = "####### "
  f16[ 35][ 9] = " ## ##  "
  f16[ 35][10] = " ## ##  "
  f16[ 35][11] = " ## ##  "

  f16[ 36][ 2] = "   $    "
  f16[ 36][ 3] = "   $    "
  f16[ 36][ 4] = " $$$$$  "
  f16[ 36][ 5] = "$$ $ $$ "
  f16[ 36][ 6] = "$$ $    "
  f16[ 36][ 7] = "$$ $    "
  f16[ 36][ 8] = " $$$$$  "
  f16[ 36][ 9] = "   $ $$ "
  f16[ 36][10] = "   $ $$ "
  f16[ 36][11] = "$$ $ $$ "
  f16[ 36][12] = " $$$$$  "
  f16[ 36][13] = "   $    "
  f16[ 36][14] = "   $    "

  f16[ 37][ 3] = "%%    % "
  f16[ 37][ 4] = "%%   %% "
  f16[ 37][ 5] = "    %%% "
  f16[ 37][ 6] = "   %%%  "
  f16[ 37][ 7] = "  %%%   "
  f16[ 37][ 8] = " %%%    "
  f16[ 37][ 9] = "%%%     "
  f16[ 37][10] = "%%   %% "
  f16[ 37][11] = "%    %% "

  f16[ 38][ 3] = "  &&&   "
  f16[ 38][ 4] = " && &&  "
  f16[ 38][ 5] = " && &&  "
  f16[ 38][ 6] = "  &&&   "
  f16[ 38][ 7] = " &&& && "
  f16[ 38][ 8] = "&& &&&  "
  f16[ 38][ 9] = "&&  &&  "
  f16[ 38][10] = "&&  &&  "
  f16[ 38][11] = " &&& && "

  f16[ 39][ 2] = "  ''    "
  f16[ 39][ 3] = "  ''    "
  f16[ 39][ 4] = "  ''    "
  f16[ 39][ 5] = " ''     "

  f16[ 40][ 3] = "    ((  "
  f16[ 40][ 4] = "   ((   "
  f16[ 40][ 5] = "  ((    "
  f16[ 40][ 6] = "  ((    "
  f16[ 40][ 7] = "  ((    "
  f16[ 40][ 8] = "  ((    "
  f16[ 40][ 9] = "  ((    "
  f16[ 40][10] = "   ((   "
  f16[ 40][11] = "    ((  "

  f16[ 41][ 3] = " ))     "
  f16[ 41][ 4] = "  ))    "
  f16[ 41][ 5] = "   ))   "
  f16[ 41][ 6] = "   ))   "
  f16[ 41][ 7] = "   ))   "
  f16[ 41][ 8] = "   ))   "
  f16[ 41][ 9] = "   ))   "
  f16[ 41][10] = "  ))    "
  f16[ 41][11] = " ))     "

  f16[ 42][ 5] = " ** **  "
  f16[ 42][ 6] = "  ***   "
  f16[ 42][ 7] = "******* "
  f16[ 42][ 8] = "  ***   "
  f16[ 42][ 9] = " ** **  "

  f16[ 43][ 5] = "   ++   "
  f16[ 43][ 6] = "   ++   "
  f16[ 43][ 7] = " ++++++ "
  f16[ 43][ 8] = "   ++   "
  f16[ 43][ 9] = "   ++   "

  f16[ 44][ 9] = "   ,,   "
  f16[ 44][10] = "   ,,   "
  f16[ 44][11] = "   ,,   "
  f16[ 44][12] = "  ,,    "

  f16[ 45][ 7] = " ------ "

  f16[ 46][10] = "   ..   "
  f16[ 46][11] = "   ..   "

  f16[ 47][ 3] = "      / "
  f16[ 47][ 4] = "     // "
  f16[ 47][ 5] = "    /// "
  f16[ 47][ 6] = "   ///  "
  f16[ 47][ 7] = "  ///   "
  f16[ 47][ 8] = " ///    "
  f16[ 47][ 9] = "///     "
  f16[ 47][10] = "//      "
  f16[ 47][11] = "/       "

  f16[ 48][ 2] = " 00000  "
  f16[ 48][ 3] = "00   00 "
  f16[ 48][ 4] = "00   00 "
  f16[ 48][ 5] = "00   00 "
  f16[ 48][ 6] = "00 0 00 "
  f16[ 48][ 7] = "00 0 00 "
  f16[ 48][ 8] = "00   00 "
  f16[ 48][ 9] = "00   00 "
  f16[ 48][10] = "00   00 "
  f16[ 48][11] = " 00000  "

  f16[ 49][ 2] = "   11   "
  f16[ 49][ 3] = "  111   "
  f16[ 49][ 4] = " 1111   "
  f16[ 49][ 5] = "   11   "
  f16[ 49][ 6] = "   11   "
  f16[ 49][ 7] = "   11   "
  f16[ 49][ 8] = "   11   "
  f16[ 49][ 9] = "   11   "
  f16[ 49][10] = "   11   "
  f16[ 49][11] = " 111111 "

  f16[ 50][ 2] = " 22222  "
  f16[ 50][ 3] = "22   22 "
  f16[ 50][ 4] = "     22 "
  f16[ 50][ 5] = "    22  "
  f16[ 50][ 6] = "   22   "
  f16[ 50][ 7] = "  22    "
  f16[ 50][ 8] = " 22     "
  f16[ 50][ 9] = "22      "
  f16[ 50][10] = "22   22 "
  f16[ 50][11] = "2222222 "

  f16[ 51][ 2] = " 33333  "
  f16[ 51][ 3] = "33   33 "
  f16[ 51][ 4] = "     33 "
  f16[ 51][ 5] = "     33 "
  f16[ 51][ 6] = "   333  "
  f16[ 51][ 7] = "     33 "
  f16[ 51][ 8] = "     33 "
  f16[ 51][ 9] = "     33 "
  f16[ 51][10] = "33   33 "
  f16[ 51][11] = " 33333  "

  f16[ 52][ 2] = "    44  "
  f16[ 52][ 3] = "   444  "
  f16[ 52][ 4] = "  4444  "
  f16[ 52][ 5] = " 44 44  "
  f16[ 52][ 6] = "44  44  "
  f16[ 52][ 7] = "4444444 "
  f16[ 52][ 8] = "    44  "
  f16[ 52][ 9] = "    44  "
  f16[ 52][10] = "    44  "
  f16[ 52][11] = "   4444 "

  f16[ 53][ 2] = "5555555 "
  f16[ 53][ 3] = "55      "
  f16[ 53][ 4] = "55      "
  f16[ 53][ 5] = "55      "
  f16[ 53][ 6] = "555555  "
  f16[ 53][ 7] = "     55 "
  f16[ 53][ 8] = "     55 "
  f16[ 53][ 9] = "     55 "
  f16[ 53][10] = "55   55 "
  f16[ 53][11] = " 55555  "

  f16[ 54][ 2] = "  666   "
  f16[ 54][ 3] = " 66     "
  f16[ 54][ 4] = "66      "
  f16[ 54][ 5] = "66      "
  f16[ 54][ 6] = "666666  "
  f16[ 54][ 7] = "66   66 "
  f16[ 54][ 8] = "66   66 "
  f16[ 54][ 9] = "66   66 "
  f16[ 54][10] = "66   66 "
  f16[ 54][11] = " 66666  "

  f16[ 55][ 2] = "7777777 "
  f16[ 55][ 3] = "77   77 "
  f16[ 55][ 4] = "     77 "
  f16[ 55][ 5] = "     77 "
  f16[ 55][ 6] = "    77  "
  f16[ 55][ 7] = "   77   "
  f16[ 55][ 8] = "  77    "
  f16[ 55][ 9] = "  77    "
  f16[ 55][10] = "  77    "
  f16[ 55][11] = "  77    "

  f16[ 56][ 2] = " 88888  "
  f16[ 56][ 3] = "88   88 "
  f16[ 56][ 4] = "88   88 "
  f16[ 56][ 5] = "88   88 "
  f16[ 56][ 6] = " 88888  "
  f16[ 56][ 7] = "88   88 "
  f16[ 56][ 8] = "88   88 "
  f16[ 56][ 9] = "88   88 "
  f16[ 56][10] = "88   88 "
  f16[ 56][11] = " 88888  "

  f16[ 57][ 2] = " 99999  "
  f16[ 57][ 3] = "99   99 "
  f16[ 57][ 4] = "99   99 "
  f16[ 57][ 5] = "99   99 "
  f16[ 57][ 6] = " 999999 "
  f16[ 57][ 7] = "     99 "
  f16[ 57][ 8] = "     99 "
  f16[ 57][ 9] = "     99 "
  f16[ 57][10] = "    99  "
  f16[ 57][11] = "  999   "

  f16[ 58][ 5] = "   ::   "
  f16[ 58][ 6] = "   ::   "
  f16[ 58][ 7] = "        "
  f16[ 58][ 8] = "        "
  f16[ 58][ 9] = "   ::   "
  f16[ 58][10] = "   ::   "

  f16[ 59][ 5] = "   ,;   "
  f16[ 59][ 6] = "   ,;   "
  f16[ 59][ 7] = "        "
  f16[ 59][ 8] = "        "
  f16[ 59][ 9] = "   ,;   "
  f16[ 59][10] = "   ,;   "
  f16[ 59][11] = "   ,;   "
  f16[ 59][12] = "  ,;    "

  f16[ 60][ 3] = "    <<  "
  f16[ 60][ 4] = "   <<   "
  f16[ 60][ 5] = "  <<    "
  f16[ 60][ 6] = " <<     "
  f16[ 60][ 7] = "<<      "
  f16[ 60][ 8] = " <<     "
  f16[ 60][ 9] = "  <<    "
  f16[ 60][10] = "   <<   "
  f16[ 60][11] = "    <<  "

  f16[ 61][ 6] = "======= "
  f16[ 61][ 7] = "        "
  f16[ 61][ 8] = "        "
  f16[ 61][ 9] = "======= "

  f16[ 62][ 3] = " >>     "
  f16[ 62][ 4] = "  >>    "
  f16[ 62][ 5] = "   >>   "
  f16[ 62][ 6] = "    >>  "
  f16[ 62][ 7] = "     >> "
  f16[ 62][ 8] = "    >>  "
  f16[ 62][ 9] = "   >>   "
  f16[ 62][10] = "  >>    "
  f16[ 62][11] = " >>     "

  f16[ 63][ 2] = " ?????  "
  f16[ 63][ 3] = "??   ?? "
  f16[ 63][ 4] = "??   ?? "
  f16[ 63][ 5] = "    ??  "
  f16[ 63][ 6] = "   ??   "
  f16[ 63][ 7] = "   ??   "
  f16[ 63][ 8] = "   ??   "
  f16[ 63][ 9] = "        "
  f16[ 63][10] = "   ??   "
  f16[ 63][11] = "   ??   "

  f16[ 64][ 2] = " @@@@@  "
  f16[ 64][ 3] = "@@   @@ "
  f16[ 64][ 4] = "@@   @@ "
  f16[ 64][ 5] = "@@   @@ "
  f16[ 64][ 6] = "@@ @@@@ "
  f16[ 64][ 7] = "@@ @@@@ "
  f16[ 64][ 8] = "@@ @@@@ "
  f16[ 64][ 9] = "@@ @@@  "
  f16[ 64][10] = "@@      "
  f16[ 64][11] = " @@@@@  "

  f16[ 65][ 2] = "   A    "
  f16[ 65][ 3] = "  AAA   "
  f16[ 65][ 4] = " AA AA  "
  f16[ 65][ 5] = "AA   AA "
  f16[ 65][ 6] = "AA   AA "
  f16[ 65][ 7] = "AAAAAAA "
  f16[ 65][ 8] = "AA   AA "
  f16[ 65][ 9] = "AA   AA "
  f16[ 65][10] = "AA   AA "
  f16[ 65][11] = "AA   AA "

  f16[196][ 1] = "AA   AA "
  f16[196][ 2] = "        "
  f16[196][ 3] = "  AAA   "
  f16[196][ 4] = " AA AA  "
  f16[196][ 5] = "AA   AA "
  f16[196][ 6] = "AA   AA "
  f16[196][ 7] = "AAAAAAA "
  f16[196][ 8] = "AA   AA "
  f16[196][ 9] = "AA   AA "
  f16[196][10] = "AA   AA "
  f16[196][11] = "AA   AA "

  f16[ 66][ 2] = "BBBBBB  "
  f16[ 66][ 3] = " BB  BB "
  f16[ 66][ 4] = " BB  BB "
  f16[ 66][ 5] = " BB  BB "
  f16[ 66][ 6] = " BBBBB  "
  f16[ 66][ 7] = " BB  BB "
  f16[ 66][ 8] = " BB  BB "
  f16[ 66][ 9] = " BB  BB "
  f16[ 66][10] = " BB  BB "
  f16[ 66][11] = "BBBBBB  "

  f16[ 67][ 2] = "  CCCC  "
  f16[ 67][ 3] = " CC  CC "
  f16[ 67][ 4] = "CC    C "
  f16[ 67][ 5] = "CC      "
  f16[ 67][ 6] = "CC      "
  f16[ 67][ 7] = "CC      "
  f16[ 67][ 8] = "CC      "
  f16[ 67][ 9] = "CC    C "
  f16[ 67][10] = " CC  CC "
  f16[ 67][11] = "  CCCC  "

  f16[ 68][ 2] = "DDDDD   "
  f16[ 68][ 3] = " DD DD  "
  f16[ 68][ 4] = " DD  DD "
  f16[ 68][ 5] = " DD  DD "
  f16[ 68][ 6] = " DD  DD "
  f16[ 68][ 7] = " DD  DD "
  f16[ 68][ 8] = " DD  DD "
  f16[ 68][ 9] = " DD  DD "
  f16[ 68][10] = " DD DD  "
  f16[ 68][11] = "DDDDD   "

  f16[ 69][ 2] = "EEEEEEE "
  f16[ 69][ 3] = " EE  EE "
  f16[ 69][ 4] = " EE   E "
  f16[ 69][ 5] = " EE E   "
  f16[ 69][ 6] = " EEEE   "
  f16[ 69][ 7] = " EE E   "
  f16[ 69][ 8] = " EE     "
  f16[ 69][ 9] = " EE   E "
  f16[ 69][10] = " EE  EE "
  f16[ 69][11] = "EEEEEEE "

  f16[ 70][ 2] = "FFFFFFF "
  f16[ 70][ 3] = " FF  FF "
  f16[ 70][ 4] = " FF   F "
  f16[ 70][ 5] = " FF F   "
  f16[ 70][ 6] = " FFFF   "
  f16[ 70][ 7] = " FF F   "
  f16[ 70][ 8] = " FF     "
  f16[ 70][ 9] = " FF     "
  f16[ 70][10] = " FF     "
  f16[ 70][11] = "FFFF    "

  f16[ 71][ 2] = "  GGGG  "
  f16[ 71][ 3] = " GG  GG "
  f16[ 71][ 4] = "GG    G "
  f16[ 71][ 5] = "GG      "
  f16[ 71][ 6] = "GG      "
  f16[ 71][ 7] = "GG GGGG "
  f16[ 71][ 8] = "GG   GG "
  f16[ 71][ 9] = "GG   GG "
  f16[ 71][10] = " GG  GG "
  f16[ 71][11] = "  GGG G "

  f16[ 72][ 2] = "HH   HH "
  f16[ 72][ 3] = "HH   HH "
  f16[ 72][ 4] = "HH   HH "
  f16[ 72][ 5] = "HH   HH "
  f16[ 72][ 6] = "HHHHHHH "
  f16[ 72][ 7] = "HH   HH "
  f16[ 72][ 8] = "HH   HH "
  f16[ 72][ 9] = "HH   HH "
  f16[ 72][10] = "HH   HH "
  f16[ 72][11] = "HH   HH "

  f16[ 73][ 2] = "  IIII  "
  f16[ 73][ 3] = "   II   "
  f16[ 73][ 4] = "   II   "
  f16[ 73][ 5] = "   II   "
  f16[ 73][ 6] = "   II   "
  f16[ 73][ 7] = "   II   "
  f16[ 73][ 8] = "   II   "
  f16[ 73][ 9] = "   II   "
  f16[ 73][10] = "   II   "
  f16[ 73][11] = "  IIII  "

  f16[ 74][ 2] = "   JJJJ "
  f16[ 74][ 3] = "    JJ  "
  f16[ 74][ 4] = "    JJ  "
  f16[ 74][ 5] = "    JJ  "
  f16[ 74][ 6] = "    JJ  "
  f16[ 74][ 7] = "    JJ  "
  f16[ 74][ 8] = "JJ  JJ  "
  f16[ 74][ 9] = "JJ  JJ  "
  f16[ 74][10] = "JJ  JJ  "
  f16[ 74][11] = " JJJJ   "

  f16[ 75][ 2] = "KKK  KK "
  f16[ 75][ 3] = " KK  KK "
  f16[ 75][ 4] = " KK  KK "
  f16[ 75][ 5] = " KK KK  "
  f16[ 75][ 6] = " KKKK   "
  f16[ 75][ 7] = " KKKK   "
  f16[ 75][ 8] = " KK KK  "
  f16[ 75][ 9] = " KK  KK "
  f16[ 75][10] = " KK  KK "
  f16[ 75][11] = "KKK  KK "

  f16[ 76][ 2] = "LLLL    "
  f16[ 76][ 3] = " LL     "
  f16[ 76][ 4] = " LL     "
  f16[ 76][ 5] = " LL     "
  f16[ 76][ 6] = " LL     "
  f16[ 76][ 7] = " LL     "
  f16[ 76][ 8] = " LL     "
  f16[ 76][ 9] = " LL   L "
  f16[ 76][10] = " LL  LL "
  f16[ 76][11] = "LLLLLLL "

  f16[ 77][ 2] = "MM   MM "
  f16[ 77][ 3] = "MMM MMM "
  f16[ 77][ 4] = "MMMMMMM "
  f16[ 77][ 5] = "MMMMMMM "
  f16[ 77][ 6] = "MM M MM "
  f16[ 77][ 7] = "MM   MM "
  f16[ 77][ 8] = "MM   MM "
  f16[ 77][ 9] = "MM   MM "
  f16[ 77][10] = "MM   MM "
  f16[ 77][11] = "MM   MM "

  f16[ 78][ 2] = "NN   NN "
  f16[ 78][ 3] = "NNN  NN "
  f16[ 78][ 4] = "NNNN NN "
  f16[ 78][ 5] = "NNNNNNN "
  f16[ 78][ 6] = "NN NNNN "
  f16[ 78][ 7] = "NN  NNN "
  f16[ 78][ 8] = "NN   NN "
  f16[ 78][ 9] = "NN   NN "
  f16[ 78][10] = "NN   NN "
  f16[ 78][11] = "NN   NN "

  f16[ 79][ 2] = " OOOOO  "
  f16[ 79][ 3] = "OO   OO "
  f16[ 79][ 4] = "OO   OO "
  f16[ 79][ 5] = "OO   OO "
  f16[ 79][ 6] = "OO   OO "
  f16[ 79][ 7] = "OO   OO "
  f16[ 79][ 8] = "OO   OO "
  f16[ 79][ 9] = "OO   OO "
  f16[ 79][10] = "OO   OO "
  f16[ 79][11] = " OOOOO  "

  f16[214][ 1] = "OO   OO "
  f16[214][ 2] = "        "
  f16[214][ 3] = " OOOOO  "
  f16[214][ 4] = "OO   OO "
  f16[214][ 5] = "OO   OO "
  f16[214][ 6] = "OO   OO "
  f16[214][ 7] = "OO   OO "
  f16[214][ 8] = "OO   OO "
  f16[214][ 9] = "OO   OO "
  f16[214][10] = "OO   OO "
  f16[214][11] = " OOOOO  "

  f16[ 80][ 2] = "PPPPPP  "
  f16[ 80][ 3] = " PP  PP "
  f16[ 80][ 4] = " PP  PP "
  f16[ 80][ 5] = " PP  PP "
  f16[ 80][ 6] = " PPPPP  "
  f16[ 80][ 7] = " PP     "
  f16[ 80][ 8] = " PP     "
  f16[ 80][ 9] = " PP     "
  f16[ 80][10] = " PP     "
  f16[ 80][11] = "PPPP    "

  f16[ 81][ 2] = " QQQQQ  "
  f16[ 81][ 3] = "QQ   QQ "
  f16[ 81][ 4] = "QQ   QQ "
  f16[ 81][ 5] = "QQ   QQ "
  f16[ 81][ 6] = "QQ   QQ "
  f16[ 81][ 7] = "QQ   QQ "
  f16[ 81][ 8] = "QQQ  QQ "
  f16[ 81][ 9] = "QQQQ QQ "
  f16[ 81][10] = "QQ QQQQ "
  f16[ 81][11] = " QQQQQQ "

  f16[ 82][ 2] = "RRRRRR  "
  f16[ 82][ 3] = " RR  RR "
  f16[ 82][ 4] = " RR  RR "
  f16[ 82][ 5] = " RR  RR "
  f16[ 82][ 6] = " RRRRR  "
  f16[ 82][ 7] = " RR RR  "
  f16[ 82][ 8] = " RR  RR "
  f16[ 82][ 9] = " RR  RR "
  f16[ 82][10] = " RR  RR "
  f16[ 82][11] = "RRR  RR "

  f16[ 83][ 2] = " SSSSS  "
  f16[ 83][ 3] = "SS   SS "
  f16[ 83][ 4] = "SS   SS "
  f16[ 83][ 5] = " SS     "
  f16[ 83][ 6] = "  SSS   "
  f16[ 83][ 7] = "    SS  "
  f16[ 83][ 8] = "     SS "
  f16[ 83][ 9] = "SS   SS "
  f16[ 83][10] = "SS   SS "
  f16[ 83][11] = " SSSSS  "

  f16[ 84][ 2] = " TTTTTT "
  f16[ 84][ 3] = " TTTTTT "
  f16[ 84][ 4] = " T TT T "
  f16[ 84][ 5] = "   TT   "
  f16[ 84][ 6] = "   TT   "
  f16[ 84][ 7] = "   TT   "
  f16[ 84][ 8] = "   TT   "
  f16[ 84][ 9] = "   TT   "
  f16[ 84][10] = "   TT   "
  f16[ 84][11] = "  TTTT  "

  f16[ 85][ 2] = "UU   UU "
  f16[ 85][ 3] = "UU   UU "
  f16[ 85][ 4] = "UU   UU "
  f16[ 85][ 5] = "UU   UU "
  f16[ 85][ 6] = "UU   UU "
  f16[ 85][ 7] = "UU   UU "
  f16[ 85][ 8] = "UU   UU "
  f16[ 85][ 9] = "UU   UU "
  f16[ 85][10] = "UU   UU "
  f16[ 85][11] = " UUUUU  "

  f16[220][ 1] = "UU   UU "
  f16[220][ 2] = "        "
  f16[220][ 3] = "UU   UU "
  f16[220][ 4] = "UU   UU "
  f16[220][ 5] = "UU   UU "
  f16[220][ 6] = "UU   UU "
  f16[220][ 7] = "UU   UU "
  f16[220][ 8] = "UU   UU "
  f16[220][ 9] = "UU   UU "
  f16[220][10] = "UU   UU "
  f16[220][11] = " UUUUU  "

  f16[ 86][ 2] = "VV   VV "
  f16[ 86][ 3] = "VV   VV "
  f16[ 86][ 4] = "VV   VV "
  f16[ 86][ 5] = "VV   VV "
  f16[ 86][ 6] = "VV   VV "
  f16[ 86][ 7] = "VV   VV "
  f16[ 86][ 8] = "VV   VV "
  f16[ 86][ 9] = " VV VV  "
  f16[ 86][10] = "  VVV   "
  f16[ 86][11] = "   V    "

  f16[ 87][ 2] = "WW   WW "
  f16[ 87][ 3] = "WW   WW "
  f16[ 87][ 4] = "WW   WW "
  f16[ 87][ 5] = "WW   WW "
  f16[ 87][ 6] = "WW W WW "
  f16[ 87][ 7] = "WW W WW "
  f16[ 87][ 8] = "WW W WW "
  f16[ 87][ 9] = " WWWWW  "
  f16[ 87][10] = " WW WW  "
  f16[ 87][11] = " WW WW  "

  f16[ 88][ 2] = "XX   XX "
  f16[ 88][ 3] = "XX   XX "
  f16[ 88][ 4] = " XX XX  "
  f16[ 88][ 5] = " XX XX  "
  f16[ 88][ 6] = "  XXX   "
  f16[ 88][ 7] = "  XXX   "
  f16[ 88][ 8] = " XX XX  "
  f16[ 88][ 9] = " XX XX  "
  f16[ 88][10] = "XX   XX "
  f16[ 88][11] = "XX   XX "

  f16[ 89][ 2] = " YY  YY "
  f16[ 89][ 3] = " YY  YY "
  f16[ 89][ 4] = " YY  YY "
  f16[ 89][ 5] = " YY  YY "
  f16[ 89][ 6] = "  YYYY  "
  f16[ 89][ 7] = "   YY   "
  f16[ 89][ 8] = "   YY   "
  f16[ 89][ 9] = "   YY   "
  f16[ 89][10] = "   YY   "
  f16[ 89][11] = "  YYYY  "

  f16[ 90][ 2] = "ZZZZZZZ "
  f16[ 90][ 3] = "ZZ   ZZ "
  f16[ 90][ 4] = "Z    ZZ "
  f16[ 90][ 5] = "    ZZ  "
  f16[ 90][ 6] = "   ZZ   "
  f16[ 90][ 7] = "  ZZ    "
  f16[ 90][ 8] = " ZZ     "
  f16[ 90][ 9] = "ZZ    Z "
  f16[ 90][10] = "ZZ   ZZ "
  f16[ 90][11] = "ZZZZZZZ "

  f16[ 91][ 2] = " [[[[[  "
  f16[ 91][ 3] = " [[     "
  f16[ 91][ 4] = " [[     "
  f16[ 91][ 5] = " [[     "
  f16[ 91][ 6] = " [[     "
  f16[ 91][ 7] = " [[     "
  f16[ 91][ 8] = " [[     "
  f16[ 91][ 9] = " [[     "
  f16[ 91][10] = " [[     "
  f16[ 91][11] = " [[[[[  "

  f16[ 92][ 3] = "/       "
  f16[ 92][ 4] = "//      "
  f16[ 92][ 5] = "///     "
  f16[ 92][ 6] = " ///    "
  f16[ 92][ 7] = "  ///   "
  f16[ 92][ 8] = "   ///  "
  f16[ 92][ 9] = "    /// "
  f16[ 92][10] = "     // "
  f16[ 92][11] = "      / "

  f16[ 93][ 2] = " ]]]]]  "
  f16[ 93][ 3] = "    ]]  "
  f16[ 93][ 4] = "    ]]  "
  f16[ 93][ 5] = "    ]]  "
  f16[ 93][ 6] = "    ]]  "
  f16[ 93][ 7] = "    ]]  "
  f16[ 93][ 8] = "    ]]  "
  f16[ 93][ 9] = "    ]]  "
  f16[ 93][10] = "    ]]  "
  f16[ 93][11] = " ]]]]]  "

  f16[ 94][ 0] = "   ^    "
  f16[ 94][ 1] = "  ^^^   "
  f16[ 94][ 2] = " ^^ ^^  "
  f16[ 94][ 3] = "^^   ^^ "

  f16[ 95][13] = "_______ "

  f16[ 96][ 2] = "  ``    "
  f16[ 96][ 3] = "  ``    "
  f16[ 96][ 4] = "   ``   "

  f16[ 97][ 5] = " aaaa   "
  f16[ 97][ 6] = "    aa  "
  f16[ 97][ 7] = " aaaaa  "
  f16[ 97][ 8] = "aa  aa  "
  f16[ 97][ 9] = "aa  aa  "
  f16[ 97][10] = "aa  aa  "
  f16[ 97][11] = " aaa aa "

  f16[228][ 3] = " aa aa  "
  f16[228][ 4] = "        "
  f16[228][ 5] = " aaaa   "
  f16[228][ 6] = "    aa  "
  f16[228][ 7] = " aaaaa  "
  f16[228][ 8] = "aa  aa  "
  f16[228][ 9] = "aa  aa  "
  f16[228][10] = "aa  aaa "
  f16[228][11] = " aaa aa "

  f16[ 98][ 2] = "bbb     "
  f16[ 98][ 3] = " bb     "
  f16[ 98][ 4] = " bb     "
  f16[ 98][ 5] = " bbbb   "
  f16[ 98][ 6] = " bb bb  "
  f16[ 98][ 7] = " bb  bb "
  f16[ 98][ 8] = " bb  bb "
  f16[ 98][ 9] = " bb  bb "
  f16[ 98][10] = " bb  bb "
  f16[ 98][11] = " bbbbb  "

  f16[ 99][ 5] = " ccccc  "
  f16[ 99][ 6] = "cc   cc "
  f16[ 99][ 7] = "cc      "
  f16[ 99][ 8] = "cc      "
  f16[ 99][ 9] = "cc      "
  f16[ 99][10] = "cc   cc "
  f16[ 99][11] = " ccccc  "

  f16[100][ 2] = "   ddd  "
  f16[100][ 3] = "    dd  "
  f16[100][ 4] = "    dd  "
  f16[100][ 5] = "  dddd  "
  f16[100][ 6] = " dd dd  "
  f16[100][ 7] = "dd  dd  "
  f16[100][ 8] = "dd  dd  "
  f16[100][ 9] = "dd  dd  "
  f16[100][10] = "dd  dd  "
  f16[100][11] = " ddd dd "

  f16[101][ 5] = " eeeee  "
  f16[101][ 6] = "ee   ee "
  f16[101][ 7] = "eeeeeee "
  f16[101][ 8] = "ee      "
  f16[101][ 9] = "ee      "
  f16[101][10] = "ee   ee "
  f16[101][11] = " eeeee  "

  f16[102][ 2] = "  fff   "
  f16[102][ 3] = " ff ff  "
  f16[102][ 4] = " ff  f  "
  f16[102][ 5] = " ff     "
  f16[102][ 6] = "ffff    "
  f16[102][ 7] = " ff     "
  f16[102][ 8] = " ff     "
  f16[102][ 9] = " ff     "
  f16[102][10] = " ff     "
  f16[102][11] = "ffff    "

  f16[103][ 5] = " ggg gg "
  f16[103][ 6] = "gg  gg  "
  f16[103][ 7] = "gg  gg  "
  f16[103][ 8] = "gg  gg  "
  f16[103][ 9] = "gg  gg  "
  f16[103][10] = " ggggg  "
  f16[103][11] = "    gg  "
  f16[103][12] = "gg  gg  "
  f16[103][13] = " gggg   "

  f16[104][ 2] = "hh      "
  f16[104][ 3] = " hh     "
  f16[104][ 4] = " hh     "
  f16[104][ 5] = " hh hh  "
  f16[104][ 6] = " hhh hh "
  f16[104][ 7] = " hh  hh "
  f16[104][ 8] = " hh  hh "
  f16[104][ 9] = " hh  hh "
  f16[104][10] = " hh  hh "
  f16[104][11] = "hhh  hh "

  f16[105][ 2] = "   ii   "
  f16[105][ 3] = "   ii   "
  f16[105][ 4] = "        "
  f16[105][ 5] = "  iii   "
  f16[105][ 6] = "   ii   "
  f16[105][ 7] = "   ii   "
  f16[105][ 8] = "   ii   "
  f16[105][ 9] = "   ii   "
  f16[105][10] = "   ii   "
  f16[105][11] = "  iiii  "

  f16[106][ 2] = "     jj "
  f16[106][ 3] = "     jj "
  f16[106][ 4] = "        "
  f16[106][ 5] = "    jjj "
  f16[106][ 6] = "     jj "
  f16[106][ 7] = "     jj "
  f16[106][ 8] = "     jj "
  f16[106][ 9] = "     jj "
  f16[106][10] = "     jj "
  f16[106][11] = "     jj "
  f16[106][12] = " jj  jj "
  f16[106][13] = " jj  jj "
  f16[106][14] = "  jjjj  "

  f16[107][ 2] = "kkk     "
  f16[107][ 3] = " kk     "
  f16[107][ 4] = " kk     "
  f16[107][ 5] = " kk  kk "
  f16[107][ 6] = " kk kk  "
  f16[107][ 7] = " kkkk   "
  f16[107][ 8] = " kkkk   "
  f16[107][ 9] = " kk kk  "
  f16[107][10] = " kk  kk "
  f16[107][11] = "kkk  kk "

  f16[108][ 2] = "  lll   "
  f16[108][ 3] = "   ll   "
  f16[108][ 4] = "   ll   "
  f16[108][ 5] = "   ll   "
  f16[108][ 6] = "   ll   "
  f16[108][ 7] = "   ll   "
  f16[108][ 8] = "   ll   "
  f16[108][ 9] = "   ll   "
  f16[108][10] = "   ll   "
  f16[108][11] = "  llll  "

  f16[109][ 5] = "mmm mm  "
  f16[109][ 6] = "mmmmmmm "
  f16[109][ 7] = "mm m mm "
  f16[109][ 8] = "mm m mm "
  f16[109][ 9] = "mm m mm "
  f16[109][10] = "mm m mm "
  f16[109][11] = "mm m mm "

  f16[110][ 5] = "nn nnn  "
  f16[110][ 6] = " nn  nn "
  f16[110][ 7] = " nn  nn "
  f16[110][ 8] = " nn  nn "
  f16[110][ 9] = " nn  nn "
  f16[110][10] = " nn  nn "
  f16[110][11] = " nn  nn "

  f16[111][ 5] = " ooooo  "
  f16[111][ 6] = "oo   oo "
  f16[111][ 7] = "oo   oo "
  f16[111][ 8] = "oo   oo "
  f16[111][ 9] = "oo   oo "
  f16[111][10] = "oo   oo "
  f16[111][11] = " ooooo  "

  f16[246][ 3] = "oo   oo "
  f16[246][ 4] = "        "
  f16[246][ 5] = " ooooo  "
  f16[246][ 6] = "oo   oo "
  f16[246][ 7] = "oo   oo "
  f16[246][ 8] = "oo   oo "
  f16[246][ 9] = "oo   oo "
  f16[246][10] = "oo   oo "
  f16[246][11] = " ooooo  "

  f16[112][ 5] = "pp ppp  "
  f16[112][ 6] = " pp  pp "
  f16[112][ 7] = " pp  pp "
  f16[112][ 8] = " pp  pp "
  f16[112][ 9] = " pp  pp "
  f16[112][10] = " pp  pp "
  f16[112][11] = " ppppp  "
  f16[112][12] = " pp     "
  f16[112][13] = " pp     "
  f16[112][14] = "pppp    "

  f16[113][ 5] = " qqq qq "
  f16[113][ 6] = "qq  qq  "
  f16[113][ 7] = "qq  qq  "
  f16[113][ 8] = "qq  qq  "
  f16[113][ 9] = "qq  qq  "
  f16[113][10] = "qq  qq  "
  f16[113][11] = " qqqqq  "
  f16[113][12] = "    qq  "
  f16[113][13] = "    qq  "
  f16[113][14] = "   qqqq "

  f16[114][ 5] = "rr rrr  "
  f16[114][ 6] = " rrr rr "
  f16[114][ 7] = " rr  rr "
  f16[114][ 8] = " rr     "
  f16[114][ 9] = " rr     "
  f16[114][10] = " rr     "
  f16[114][11] = "rrrr    "

  f16[115][ 5] = " sssss  "
  f16[115][ 6] = "ss   ss "
  f16[115][ 7] = " ss     "
  f16[115][ 8] = "  sss   "
  f16[115][ 9] = "    ss  "
  f16[115][10] = "ss   ss "
  f16[115][11] = " sssss  "

  f16[116][ 2] = "   t    "
  f16[116][ 3] = "  tt    "
  f16[116][ 4] = "  tt    "
  f16[116][ 5] = "tttttt  "
  f16[116][ 6] = "  tt    "
  f16[116][ 7] = "  tt    "
  f16[116][ 8] = "  tt    "
  f16[116][ 9] = "  tt    "
  f16[116][10] = "  tt tt "
  f16[116][11] = "   ttt  "

  f16[117][ 5] = "uu  uu  "
  f16[117][ 6] = "uu  uu  "
  f16[117][ 7] = "uu  uu  "
  f16[117][ 8] = "uu  uu  "
  f16[117][ 9] = "uu  uu  "
  f16[117][10] = "uu  uu  "
  f16[117][11] = " uuu uu "

  f16[252][ 3] = "uu  uu  "
  f16[252][ 4] = "        "
  f16[252][ 5] = "uu  uu  "
  f16[252][ 6] = "uu  uu  "
  f16[252][ 7] = "uu  uu  "
  f16[252][ 8] = "uu  uu  "
  f16[252][ 9] = "uu  uu  "
  f16[252][10] = "uu  uu  "
  f16[252][11] = " uuu uu "

  f16[118][ 5] = " vv  vv "
  f16[118][ 6] = " vv  vv "
  f16[118][ 7] = " vv  vv "
  f16[118][ 8] = " vv  vv "
  f16[118][ 9] = " vv  vv "
  f16[118][10] = "  vvvv  "
  f16[118][11] = "   vv   "

  f16[119][ 5] = "ww   ww "
  f16[119][ 6] = "ww   ww "
  f16[119][ 7] = "ww w ww "
  f16[119][ 8] = "ww w ww "
  f16[119][ 9] = "ww w ww "
  f16[119][10] = "wwwwwww "
  f16[119][11] = " ww ww  "

  f16[120][ 5] = "xx   xx "
  f16[120][ 6] = " xx xx  "
  f16[120][ 7] = "  xxx   "
  f16[120][ 8] = "  xxx   "
  f16[120][ 9] = "  xxx   "
  f16[120][10] = " xx xx  "
  f16[120][11] = "xx   xx "

  f16[121][ 5] = "yy   yy "
  f16[121][ 6] = "yy   yy "
  f16[121][ 7] = "yy   yy "
  f16[121][ 8] = "yy   yy "
  f16[121][ 9] = "yy   yy "
  f16[121][10] = "yy   yy "
  f16[121][11] = " yyyyyy "
  f16[121][12] = "     yy "
  f16[121][13] = "    yy  "
  f16[121][14] = "yyyyy   "

  f16[122][ 5] = "zzzzzzz "
  f16[122][ 6] = "zz  zz  "
  f16[122][ 7] = "   zz   "
  f16[122][ 8] = "  zz    "
  f16[122][ 9] = " zz     "
  f16[122][10] = "zz   zz "
  f16[122][11] = "zzzzzzz "

  f16[123][ 2] = "    {{{ "
  f16[123][ 3] = "   {{   "
  f16[123][ 4] = "   {{   "
  f16[123][ 5] = "   {{   "
  f16[123][ 6] = " {{{    "
  f16[123][ 7] = "   {{   "
  f16[123][ 8] = "   {{   "
  f16[123][ 9] = "   {{   "
  f16[123][10] = "   {{   "
  f16[123][11] = "    {{{ "

  f16[124][ 2] = "   ||   "
  f16[124][ 3] = "   ||   "
  f16[124][ 4] = "   ||   "
  f16[124][ 5] = "   ||   "
  f16[124][ 6] = "        "
  f16[124][ 7] = "   ||   "
  f16[124][ 8] = "   ||   "
  f16[124][ 9] = "   ||   "
  f16[124][10] = "   ||   "
  f16[124][11] = "   ||   "

  f16[125][ 2] = " }}}    "
  f16[125][ 3] = "   }}   "
  f16[125][ 4] = "   }}   "
  f16[125][ 5] = "   }}   "
  f16[125][ 6] = "    }}} "
  f16[125][ 7] = "   }}   "
  f16[125][ 8] = "   }}   "
  f16[125][ 9] = "   }}   "
  f16[125][10] = "   }}   "
  f16[125][11] = " }}}    "

  f16[126][ 2] = " ~~~ ~~ "
  f16[126][ 3] = "~~ ~~~  "

  f16[164][ 3] = "  eeee  "
  f16[164][ 4] = " ee   e "
  f16[164][ 5] = " e      "
  f16[164][ 6] = "eeeeee  "
  f16[164][ 7] = " e      "
  f16[164][ 8] = "eeeeee  "
  f16[164][ 9] = " e      "
  f16[164][10] = " ee   e "
  f16[164][11] = "  eeee  "

  f16[167][ 1] = "  ppp   "
  f16[167][ 2] = " pp pp  "
  f16[167][ 3] = " pp     "
  f16[167][ 4] = "  ppp   "
  f16[167][ 5] = " pp pp  "
  f16[167][ 6] = "pp   pp "
  f16[167][ 7] = "pp   pp "
  f16[167][ 8] = " pp pp  "
  f16[167][ 9] = "  ppp   "
  f16[167][10] = "    pp  "
  f16[167][11] = " pp pp  "
  f16[167][12] = "  ppp   "

  f16[176][ 2] = "  ooo   "
  f16[176][ 3] = " oo oo  "
  f16[176][ 4] = " oo oo  "
  f16[176][ 5] = "  ooo   "
  f16[176][ 6] = "        "
  f16[176][ 7] = "        "
  f16[176][ 8] = "        "
  f16[176][ 9] = "        "
  f16[176][10] = "        "
  f16[176][11] = "        "

  f16[178][ 1] = "  222   "
  f16[178][ 2] = " 2   2  "
  f16[178][ 3] = "     2  "
  f16[178][ 4] = "   22   "
  f16[178][ 5] = "  2     "
  f16[178][ 6] = " 2      "
  f16[178][ 7] = " 22222  "
  f16[178][ 8] = "        "
  f16[178][ 9] = "        "
  f16[178][10] = "        "
  f16[178][11] = "        "

  f16[179][ 1] = "  333   "
  f16[179][ 2] = " 3   3  "
  f16[179][ 3] = "     3  "
  f16[179][ 4] = "   33   "
  f16[179][ 5] = "     3  "
  f16[179][ 6] = " 3   3  "
  f16[179][ 7] = "  333   "
  f16[179][ 8] = "        "
  f16[179][ 9] = "        "
  f16[179][10] = "        "
  f16[179][11] = "        "

  f16[181][ 5] = "mm  mm  "
  f16[181][ 6] = "mm  mm  "
  f16[181][ 7] = "mm  mm  "
  f16[181][ 8] = "mm  mm  "
  f16[181][ 9] = "mm  mm  "
  f16[181][10] = "mm  mm  "
  f16[181][11] = "mmmm  m "
  f16[181][12] = "mm      "
  f16[181][13] = "mm      "
  f16[181][14] = "mm      "

  f16[223][ 2] = " sssss  "
  f16[223][ 3] = "ss   ss "
  f16[223][ 4] = "ss   ss "
  f16[223][ 5] = "ss   ss "
  f16[223][ 6] = "ss  ss  "
  f16[223][ 7] = "ss   ss "
  f16[223][ 8] = "ss   ss "
  f16[223][ 9] = "ss   ss "
  f16[223][10] = "ss s ss "
  f16[223][11] = "ss sss  "
  f16[223][12] = "ss      "
  f16[223][13] = "ss      "
  f16[223][14] = "ss      "

  for b := 0; b <= 255; b++ {
    for l := 0; l < 24; l++ {
      f24[b][l] = "            "
    }
  }

  f24[ 33][ 4] = "    !!!     "
  f24[ 33][ 5] = "    !!!     "
  f24[ 33][ 6] = "    !!!     "
  f24[ 33][ 7] = "    !!!     "
  f24[ 33][ 8] = "    !!!     "
  f24[ 33][ 9] = "    !!!     "
  f24[ 33][10] = "    !!!     "
  f24[ 33][11] = "    !!!     "
  f24[ 33][12] = "    !!!     "
  f24[ 33][13] = "    !!!     "
  f24[ 33][14] = "            "
  f24[ 33][15] = "            "
  f24[ 33][16] = "    !!!     "
  f24[ 33][17] = "    !!!     "
  f24[ 33][18] = "    !!!     "

  f24[ 34][ 1] = "   **  **   "
  f24[ 34][ 2] = "   **  **   "
  f24[ 34][ 3] = "   **  **   "
  f24[ 34][ 4] = "   **  **   "
  f24[ 34][ 5] = "   **  **   "

  f24[ 35][ 4] = "   ##  ##   "
  f24[ 35][ 5] = "   ##  ##   "
  f24[ 35][ 6] = "   ##  ##   "
  f24[ 35][ 7] = "   ##  ##   "
  f24[ 35][ 8] = " ########## "
  f24[ 35][ 9] = "   ##  ##   "
  f24[ 35][10] = "   ##  ##   "
  f24[ 35][11] = "   ##  ##   "
  f24[ 35][12] = "   ##  ##   "
  f24[ 35][13] = "   ##  ##   "
  f24[ 35][14] = " ########## "
  f24[ 35][15] = "   ##  ##   "
  f24[ 35][16] = "   ##  ##   "
  f24[ 35][17] = "   ##  ##   "
  f24[ 35][18] = "   ##  ##   "

  f24[ 36][ 3] = "     $$     "
  f24[ 36][ 4] = "     $$     "
  f24[ 36][ 5] = "   $$$$$$   "
  f24[ 36][ 6] = "  $$ $$ $$  "
  f24[ 36][ 7] = " $$  $$  $$ "
  f24[ 36][ 8] = " $$  $$     "
  f24[ 36][ 9] = " $$  $$     "
  f24[ 36][10] = "  $$ $$     "
  f24[ 36][11] = "   $$$$$$   "
  f24[ 36][12] = "     $$ $$  "
  f24[ 36][13] = "     $$  $$ "
  f24[ 36][14] = "     $$  $$ "
  f24[ 36][15] = " $$  $$  $$ "
  f24[ 36][16] = "  $$ $$ $$  "
  f24[ 36][17] = "   $$$$$$   "
  f24[ 36][18] = "     $$     "
  f24[ 36][19] = "     $$     "

  f24[ 37][ 5] = "  %%%   %%  "
  f24[ 37][ 6] = " %% %%  %%  "
  f24[ 37][ 7] = " %% %% %%   "
  f24[ 37][ 8] = "  %%%  %%   "
  f24[ 37][ 9] = "      %%    "
  f24[ 37][10] = "      %%    "
  f24[ 37][11] = "     %%     "
  f24[ 37][12] = "     %%     "
  f24[ 37][13] = "    %%      "
  f24[ 37][14] = "    %%      "
  f24[ 37][15] = "   %%  %%%  "
  f24[ 37][16] = "   %% %% %% "
  f24[ 37][17] = "  %%  %% %% "
  f24[ 37][18] = "  %%   %%%  "

  f24[ 38][ 4] = "    &&&     "
  f24[ 38][ 5] = "   && &&    "
  f24[ 38][ 6] = "  &&   &&   "
  f24[ 38][ 7] = "  &&   &&   "
  f24[ 38][ 8] = "  &&   &&   "
  f24[ 38][ 9] = "   && &&    "
  f24[ 38][10] = "    &&&     "
  f24[ 38][11] = "   &&&&  && "
  f24[ 38][12] = "  &&  && && "
  f24[ 38][13] = " &&    &&&  "
  f24[ 38][14] = " &&     &&  "
  f24[ 38][15] = " &&     &&  "
  f24[ 38][16] = " &&    &&&& "
  f24[ 38][17] = "  &&  && && "
  f24[ 38][18] = "   &&&&  && "

  f24[ 39][ 1] = "     ''     "
  f24[ 39][ 2] = "     ''     "
  f24[ 39][ 3] = "     ''     "
  f24[ 39][ 4] = "     ''     "
  f24[ 39][ 5] = "    ''      "

  f24[ 40][ 4] = "      ((    "
  f24[ 40][ 5] = "     ((     "
  f24[ 40][ 6] = "    ((      "
  f24[ 40][ 7] = "    ((      "
  f24[ 40][ 8] = "   ((       "
  f24[ 40][ 9] = "   ((       "
  f24[ 40][10] = "   ((       "
  f24[ 40][11] = "   ((       "
  f24[ 40][12] = "   ((       "
  f24[ 40][13] = "   ((       "
  f24[ 40][14] = "   ((       "
  f24[ 40][15] = "    ((      "
  f24[ 40][16] = "    ((      "
  f24[ 40][17] = "     ((     "
  f24[ 40][18] = "      ((    "

  f24[ 41][ 4] = "   ))       "
  f24[ 41][ 5] = "    ))      "
  f24[ 41][ 6] = "     ))     "
  f24[ 41][ 7] = "     ))     "
  f24[ 41][ 8] = "      ))    "
  f24[ 41][ 9] = "      ))    "
  f24[ 41][10] = "      ))    "
  f24[ 41][11] = "      ))    "
  f24[ 41][12] = "      ))    "
  f24[ 41][13] = "      ))    "
  f24[ 41][14] = "      ))    "
  f24[ 41][15] = "     ))     "
  f24[ 41][16] = "     ))     "
  f24[ 41][17] = "    ))      "
  f24[ 41][18] = "   ))       "

  f24[ 42][ 7] = " **     **  "
  f24[ 42][ 8] = "  **   **   "
  f24[ 42][ 9] = "   ** **    "
  f24[ 42][10] = "    ***     "
  f24[ 42][11] = "*********** "
  f24[ 42][12] = "    ***     "
  f24[ 42][13] = "   ** **    "
  f24[ 42][14] = "  **   **   "
  f24[ 42][15] = " **     **  "

  f24[ 43][ 7] = "     ++     "
  f24[ 43][ 8] = "     ++     "
  f24[ 43][ 9] = "     ++     "
  f24[ 43][10] = "     ++     "
  f24[ 43][11] = " ++++++++++ "
  f24[ 43][12] = "     ++     "
  f24[ 43][13] = "     ++     "
  f24[ 43][14] = "     ++     "
  f24[ 43][15] = "     ++     "

  f24[ 44][15] = "     ,,     "
  f24[ 44][16] = "     ,,     "
  f24[ 44][17] = "     ,,     "
  f24[ 44][18] = "     ,,     "
  f24[ 44][19] = "    ,,      "

  f24[ 45][11] = " ---------- "

  f24[ 46][16] = "     ..     "
  f24[ 46][17] = "     ..     "
  f24[ 46][18] = "     ..     "

  f24[ 47][ 5] = "        //  "
  f24[ 47][ 6] = "        //  "
  f24[ 47][ 7] = "       //   "
  f24[ 47][ 8] = "       //   "
  f24[ 47][ 9] = "      //    "
  f24[ 47][10] = "      //    "
  f24[ 47][11] = "     //     "
  f24[ 47][12] = "     //     "
  f24[ 47][13] = "    //      "
  f24[ 47][14] = "    //      "
  f24[ 47][15] = "   //       "
  f24[ 47][16] = "   //       "
  f24[ 47][17] = "  //        "
  f24[ 47][18] = "  //        "

  f24[ 48][ 4] = "   000000   "
  f24[ 48][ 5] = "  00    00  "
  f24[ 48][ 6] = " 00      00 "
  f24[ 48][ 7] = " 00      00 "
  f24[ 48][ 8] = " 00     000 "
  f24[ 48][ 9] = " 00    0000 "
  f24[ 48][10] = " 00   00 00 "
  f24[ 48][11] = " 00  00  00 "
  f24[ 48][12] = " 00 00   00 "
  f24[ 48][13] = " 0000    00 "
  f24[ 48][14] = " 000     00 "
  f24[ 48][15] = " 00      00 "
  f24[ 48][16] = " 00      00 "
  f24[ 48][17] = "  00    00  "
  f24[ 48][18] = "   000000   "

  f24[ 49][ 4] = "     11     "
  f24[ 49][ 5] = "    111     "
  f24[ 49][ 6] = "   1111     "
  f24[ 49][ 7] = "  11 11     "
  f24[ 49][ 8] = "     11     "
  f24[ 49][ 9] = "     11     "
  f24[ 49][10] = "     11     "
  f24[ 49][11] = "     11     "
  f24[ 49][12] = "     11     "
  f24[ 49][13] = "     11     "
  f24[ 49][14] = "     11     "
  f24[ 49][15] = "     11     "
  f24[ 49][16] = "     11     "
  f24[ 49][17] = "     11     "
  f24[ 49][18] = "  11111111  "

  f24[ 50][ 4] = "   222222   "
  f24[ 50][ 5] = "  22    22  "
  f24[ 50][ 6] = " 22      22 "
  f24[ 50][ 7] = " 22      22 "
  f24[ 50][ 8] = " 22      22 "
  f24[ 50][ 9] = "         22 "
  f24[ 50][10] = "        22  "
  f24[ 50][11] = "       22   "
  f24[ 50][12] = "      22    "
  f24[ 50][13] = "     22     "
  f24[ 50][14] = "    22      "
  f24[ 50][15] = "   22       "
  f24[ 50][16] = "  22        "
  f24[ 50][17] = " 22         "
  f24[ 50][18] = " 222222222  "

  f24[ 51][ 4] = "   333333   "
  f24[ 51][ 5] = "  33    33  "
  f24[ 51][ 6] = " 33      33 "
  f24[ 51][ 7] = "         33 "
  f24[ 51][ 8] = "         33 "
  f24[ 51][ 9] = "         33 "
  f24[ 51][10] = "        33  "
  f24[ 51][11] = "    33333   "
  f24[ 51][12] = "        33  "
  f24[ 51][13] = "         33 "
  f24[ 51][14] = "         33 "
  f24[ 51][15] = "         33 "
  f24[ 51][16] = " 33      33 "
  f24[ 51][17] = "  33    33  "
  f24[ 51][18] = "   333333   "

  f24[ 52][ 4] = "         44 "
  f24[ 52][ 5] = "        444 "
  f24[ 52][ 6] = "       4444 "
  f24[ 52][ 7] = "      44 44 "
  f24[ 52][ 8] = "     44  44 "
  f24[ 52][ 9] = "    44   44 "
  f24[ 52][10] = "   44    44 "
  f24[ 52][11] = "  44     44 "
  f24[ 52][12] = " 44      44 "
  f24[ 52][13] = " 44      44 "
  f24[ 52][14] = " 44      44 "
  f24[ 52][15] = " 4444444444 "
  f24[ 52][16] = "         44 "
  f24[ 52][17] = "         44 "
  f24[ 52][18] = "         44 "

  f24[ 53][ 4] = " 5555555555 "
  f24[ 53][ 5] = " 55         "
  f24[ 53][ 6] = " 55         "
  f24[ 53][ 7] = " 55         "
  f24[ 53][ 8] = " 55         "
  f24[ 53][ 9] = " 55         "
  f24[ 53][10] = " 55555555   "
  f24[ 53][11] = "        55  "
  f24[ 53][12] = "         55 "
  f24[ 53][13] = "         55 "
  f24[ 53][14] = "         55 "
  f24[ 53][15] = "         55 "
  f24[ 53][16] = " 55      55 "
  f24[ 53][17] = "  55    55  "
  f24[ 53][18] = "   555555   "

  f24[ 54][ 4] = "   6666666  "
  f24[ 54][ 5] = "  66        "
  f24[ 54][ 6] = " 66         "
  f24[ 54][ 7] = " 66         "
  f24[ 54][ 8] = " 66         "
  f24[ 54][ 9] = " 66         "
  f24[ 54][10] = " 66666666   "
  f24[ 54][11] = " 66     66  "
  f24[ 54][12] = " 66      66 "
  f24[ 54][13] = " 66      66 "
  f24[ 54][14] = " 66      66 "
  f24[ 54][15] = " 66      66 "
  f24[ 54][16] = " 66      66 "
  f24[ 54][17] = "  66    66  "
  f24[ 54][18] = "   666666   "

  f24[ 55][ 4] = " 7777777777 "
  f24[ 55][ 5] = " 77      77 "
  f24[ 55][ 6] = " 77      77 "
  f24[ 55][ 7] = "         77 "
  f24[ 55][ 8] = "        77  "
  f24[ 55][ 9] = "        77  "
  f24[ 55][10] = "       77   "
  f24[ 55][11] = "       77   "
  f24[ 55][12] = "      77    "
  f24[ 55][13] = "      77    "
  f24[ 55][14] = "     77     "
  f24[ 55][15] = "     77     "
  f24[ 55][16] = "     77     "
  f24[ 55][17] = "     77     "
  f24[ 55][18] = "     77     "

  f24[ 56][ 4] = "   888888   "
  f24[ 56][ 5] = "  88    88  "
  f24[ 56][ 6] = " 88      88 "
  f24[ 56][ 7] = " 88      88 "
  f24[ 56][ 8] = " 88      88 "
  f24[ 56][ 9] = " 88      88 "
  f24[ 56][10] = "  88    88  "
  f24[ 56][11] = "   888888   "
  f24[ 56][12] = "  88    88  "
  f24[ 56][13] = " 88      88 "
  f24[ 56][14] = " 88      88 "
  f24[ 56][15] = " 88      88 "
  f24[ 56][16] = " 88      88 "
  f24[ 56][17] = "  88    88  "
  f24[ 56][18] = "   888888   "

  f24[ 57][ 4] = "   999999   "
  f24[ 57][ 5] = "  99    99  "
  f24[ 57][ 6] = " 99      99 "
  f24[ 57][ 7] = " 99      99 "
  f24[ 57][ 8] = " 99      99 "
  f24[ 57][ 9] = " 99      99 "
  f24[ 57][10] = " 99      99 "
  f24[ 57][11] = "  99     99 "
  f24[ 57][12] = "   99999999 "
  f24[ 57][13] = "         99 "
  f24[ 57][14] = "         99 "
  f24[ 57][15] = "         99 "
  f24[ 57][16] = " 99      99 "
  f24[ 57][17] = "  99    99  "
  f24[ 57][18] = "   999999   "

  f24[ 58][ 8] = "     ::     "
  f24[ 58][ 9] = "     ::     "
  f24[ 58][10] = "     ::     "
  f24[ 58][11] = "            "
  f24[ 58][12] = "            "
  f24[ 58][13] = "            "
  f24[ 58][14] = "            "
  f24[ 58][15] = "     ::     "
  f24[ 58][16] = "     ::     "
  f24[ 58][17] = "     ::     "

  f24[ 59][ 8] = "     ,;     "
  f24[ 59][ 9] = "     ,;     "
  f24[ 59][10] = "     ,;     "
  f24[ 59][11] = "            "
  f24[ 59][12] = "            "
  f24[ 59][13] = "            "
  f24[ 59][14] = "            "
  f24[ 59][15] = "     ,;     "
  f24[ 59][16] = "     ,;     "
  f24[ 59][17] = "     ,;     "
  f24[ 59][18] = "     ,;     "
  f24[ 59][19] = "    ,;      "

  f24[ 60][ 4] = "        <<  "
  f24[ 60][ 5] = "       <<   "
  f24[ 60][ 6] = "      <<    "
  f24[ 60][ 7] = "     <<     "
  f24[ 60][ 8] = "    <<      "
  f24[ 60][ 9] = "   <<       "
  f24[ 60][10] = "  <<        "
  f24[ 60][11] = " <<         "
  f24[ 60][12] = "  <<        "
  f24[ 60][13] = "   <<       "
  f24[ 60][14] = "    <<      "
  f24[ 60][15] = "     <<     "
  f24[ 60][16] = "      <<    "
  f24[ 60][17] = "       <<   "
  f24[ 60][18] = "        <<  "

  f24[ 61][ 9] = " ========== "
  f24[ 61][10] = "            "
  f24[ 61][11] = "            "
  f24[ 61][12] = "            "
  f24[ 61][13] = "            "
  f24[ 61][14] = " ========== "

  f24[ 62][ 4] = " >>         "
  f24[ 62][ 5] = "  >>        "
  f24[ 62][ 6] = "   >>       "
  f24[ 62][ 7] = "    >>      "
  f24[ 62][ 8] = "     >>     "
  f24[ 62][ 9] = "      >>    "
  f24[ 62][10] = "       >>   "
  f24[ 62][11] = "        >>  "
  f24[ 62][12] = "       >>   "
  f24[ 62][13] = "      >>    "
  f24[ 62][14] = "     >>     "
  f24[ 62][15] = "    >>      "
  f24[ 62][16] = "   >>       "
  f24[ 62][17] = "  >>        "
  f24[ 62][18] = " >>         "

  f24[ 63][ 4] = "   ??????   "
  f24[ 63][ 5] = "  ??    ??  "
  f24[ 63][ 6] = " ??      ?? "
  f24[ 63][ 7] = " ??      ?? "
  f24[ 63][ 8] = " ??      ?? "
  f24[ 63][ 9] = "        ??  "
  f24[ 63][10] = "       ??   "
  f24[ 63][11] = "      ??    "
  f24[ 63][12] = "     ??     "
  f24[ 63][13] = "     ??     "
  f24[ 63][14] = "     ??     "
  f24[ 63][15] = "            "
  f24[ 63][16] = "     ??     "
  f24[ 63][17] = "     ??     "
  f24[ 63][18] = "     ??     "

  f24[ 64][ 4] = "   @@@@@@   "
  f24[ 64][ 5] = "  @@    @@  "
  f24[ 64][ 6] = " @@      @@ "
  f24[ 64][ 7] = " @@    @@@@ "
  f24[ 64][ 8] = " @@   @@ @@ "
  f24[ 64][ 9] = " @@  @@  @@ "
  f24[ 64][10] = " @@  @@  @@ "
  f24[ 64][11] = " @@  @@  @@ "
  f24[ 64][12] = " @@  @@  @@ "
  f24[ 64][13] = " @@  @@  @@ "
  f24[ 64][14] = " @@   @@ @@ "
  f24[ 64][15] = " @@    @@@@ "
  f24[ 64][16] = " @@         "
  f24[ 64][17] = "  @@        "
  f24[ 64][18] = "   @@@@@@@@ "

  f24[ 65][ 4] = "   AAAAAA   "
  f24[ 65][ 5] = "  AA    AA  "
  f24[ 65][ 6] = " AA      AA "
  f24[ 65][ 7] = " AA      AA "
  f24[ 65][ 8] = " AA      AA "
  f24[ 65][ 9] = " AA      AA "
  f24[ 65][10] = " AA      AA "
  f24[ 65][11] = " AA      AA "
  f24[ 65][12] = " AAAAAAAAAA "
  f24[ 65][13] = " AA      AA "
  f24[ 65][14] = " AA      AA "
  f24[ 65][15] = " AA      AA "
  f24[ 65][16] = " AA      AA "
  f24[ 65][17] = " AA      AA "
  f24[ 65][18] = " AA      AA "

  f24[196][ 0] = "  AA    AA  "
  f24[196][ 1] = "  AA    AA  "
  f24[196][ 2] = "  AA    AA  "
  f24[196][ 3] = "            "
  f24[196][ 4] = "   AAAAAA   "
  f24[196][ 5] = "  AA    AA  "
  f24[196][ 6] = " AA      AA "
  f24[196][ 7] = " AA      AA "
  f24[196][ 8] = " AA      AA "
  f24[196][ 9] = " AA      AA "
  f24[196][10] = " AA      AA "
  f24[196][11] = " AA      AA "
  f24[196][12] = " AAAAAAAAAA "
  f24[196][13] = " AA      AA "
  f24[196][14] = " AA      AA "
  f24[196][15] = " AA      AA "
  f24[196][16] = " AA      AA "
  f24[196][17] = " AA      AA "
  f24[196][18] = " AA      AA "

  f24[ 66][ 4] = " BBBBBBBB   "
  f24[ 66][ 5] = " BB     BB  "
  f24[ 66][ 6] = " BB      BB "
  f24[ 66][ 7] = " BB      BB "
  f24[ 66][ 8] = " BB      BB "
  f24[ 66][ 9] = " BB     BB  "
  f24[ 66][10] = " BBBBBBBB   "
  f24[ 66][11] = " BB     BB  "
  f24[ 66][12] = " BB      BB "
  f24[ 66][13] = " BB      BB "
  f24[ 66][14] = " BB      BB "
  f24[ 66][15] = " BB      BB "
  f24[ 66][16] = " BB      BB "
  f24[ 66][17] = " BB     BB  "
  f24[ 66][18] = " BBBBBBBB   "

  f24[ 67][ 4] = "   CCCCCC   "
  f24[ 67][ 5] = "  CC    CC  "
  f24[ 67][ 6] = " CC      CC "
  f24[ 67][ 7] = " CC      CC "
  f24[ 67][ 8] = " CC         "
  f24[ 67][ 9] = " CC         "
  f24[ 67][10] = " CC         "
  f24[ 67][11] = " CC         "
  f24[ 67][12] = " CC         "
  f24[ 67][13] = " CC         "
  f24[ 67][14] = " CC         "
  f24[ 67][15] = " CC      CC "
  f24[ 67][16] = " CC      CC "
  f24[ 67][17] = "  CC    CC  "
  f24[ 67][18] = "   CCCCCC   "

  f24[ 68][ 4] = " DDDDDDDD   "
  f24[ 68][ 5] = " DD     DD  "
  f24[ 68][ 6] = " DD      DD "
  f24[ 68][ 7] = " DD      DD "
  f24[ 68][ 8] = " DD      DD "
  f24[ 68][ 9] = " DD      DD "
  f24[ 68][10] = " DD      DD "
  f24[ 68][11] = " DD      DD "
  f24[ 68][12] = " DD      DD "
  f24[ 68][13] = " DD      DD "
  f24[ 68][14] = " DD      DD "
  f24[ 68][15] = " DD      DD "
  f24[ 68][16] = " DD      DD "
  f24[ 68][17] = " DD     DD  "
  f24[ 68][18] = " DDDDDDDD   "

  f24[ 69][ 4] = " EEEEEEEEEE "
  f24[ 69][ 5] = " EE         "
  f24[ 69][ 6] = " EE         "
  f24[ 69][ 7] = " EE         "
  f24[ 69][ 8] = " EE         "
  f24[ 69][ 9] = " EE         "
  f24[ 69][10] = " EE         "
  f24[ 69][11] = " EEEEEEEE   "
  f24[ 69][12] = " EE         "
  f24[ 69][13] = " EE         "
  f24[ 69][14] = " EE         "
  f24[ 69][15] = " EE         "
  f24[ 69][16] = " EE         "
  f24[ 69][17] = " EE         "
  f24[ 69][18] = " EEEEEEEEEE "

  f24[ 70][ 4] = " FFFFFFFFFF "
  f24[ 70][ 5] = " FF         "
  f24[ 70][ 6] = " FF         "
  f24[ 70][ 7] = " FF         "
  f24[ 70][ 8] = " FF         "
  f24[ 70][ 9] = " FF         "
  f24[ 70][10] = " FF         "
  f24[ 70][11] = " FFFFFFFF   "
  f24[ 70][12] = " FF         "
  f24[ 70][13] = " FF         "
  f24[ 70][14] = " FF         "
  f24[ 70][15] = " FF         "
  f24[ 70][16] = " FF         "
  f24[ 70][17] = " FF         "
  f24[ 70][18] = " FF         "

  f24[ 71][ 4] = "   GGGGGG   "
  f24[ 71][ 5] = "  GG    GG  "
  f24[ 71][ 6] = " GG      GG "
  f24[ 71][ 7] = " GG      GG "
  f24[ 71][ 8] = " GG         "
  f24[ 71][ 9] = " GG         "
  f24[ 71][10] = " GG         "
  f24[ 71][11] = " GG   GGGGG "
  f24[ 71][12] = " GG      GG "
  f24[ 71][13] = " GG      GG "
  f24[ 71][14] = " GG      GG "
  f24[ 71][15] = " GG      GG "
  f24[ 71][16] = " GG      GG "
  f24[ 71][17] = "  GG    GG  "
  f24[ 71][18] = "   GGGGGG   "

  f24[ 72][ 4] = " HH      HH "
  f24[ 72][ 5] = " HH      HH "
  f24[ 72][ 6] = " HH      HH "
  f24[ 72][ 7] = " HH      HH "
  f24[ 72][ 8] = " HH      HH "
  f24[ 72][ 9] = " HH      HH "
  f24[ 72][10] = " HH      HH "
  f24[ 72][11] = " HHHHHHHHHH "
  f24[ 72][12] = " HH      HH "
  f24[ 72][13] = " HH      HH "
  f24[ 72][14] = " HH      HH "
  f24[ 72][15] = " HH      HH "
  f24[ 72][16] = " HH      HH "
  f24[ 72][17] = " HH      HH "
  f24[ 72][18] = " HH      HH "

  f24[ 73][ 4] = "   IIIIII   "
  f24[ 73][ 5] = "     II     "
  f24[ 73][ 6] = "     II     "
  f24[ 73][ 7] = "     II     "
  f24[ 73][ 8] = "     II     "
  f24[ 73][ 9] = "     II     "
  f24[ 73][10] = "     II     "
  f24[ 73][11] = "     II     "
  f24[ 73][12] = "     II     "
  f24[ 73][13] = "     II     "
  f24[ 73][14] = "     II     "
  f24[ 73][15] = "     II     "
  f24[ 73][16] = "     II     "
  f24[ 73][17] = "     II     "
  f24[ 73][18] = "   IIIIII   "

  f24[ 74][ 4] = "     JJJJJJ "
  f24[ 74][ 5] = "       JJ   "
  f24[ 74][ 6] = "       JJ   "
  f24[ 74][ 7] = "       JJ   "
  f24[ 74][ 8] = "       JJ   "
  f24[ 74][ 9] = "       JJ   "
  f24[ 74][10] = "       JJ   "
  f24[ 74][11] = "       JJ   "
  f24[ 74][12] = "       JJ   "
  f24[ 74][13] = "       JJ   "
  f24[ 74][14] = "       JJ   "
  f24[ 74][15] = "       JJ   "
  f24[ 74][16] = "JJ     JJ   "
  f24[ 74][17] = " JJ   JJ    "
  f24[ 74][18] = "  JJJJJ     "

  f24[ 75][ 4] = " KK      KK "
  f24[ 75][ 5] = " KK     KK  "
  f24[ 75][ 6] = " KK    KK   "
  f24[ 75][ 7] = " KK   KK    "
  f24[ 75][ 8] = " KK  KK     "
  f24[ 75][ 9] = " KK KK      "
  f24[ 75][10] = " KKKK       "
  f24[ 75][11] = " KKK        "
  f24[ 75][12] = " KKKK       "
  f24[ 75][13] = " KK KK      "
  f24[ 75][14] = " KK  KK     "
  f24[ 75][15] = " KK   KK    "
  f24[ 75][16] = " KK    KK   "
  f24[ 75][17] = " KK     KK  "
  f24[ 75][18] = " KK      KK "

  f24[ 76][ 4] = " LL         "
  f24[ 76][ 5] = " LL         "
  f24[ 76][ 6] = " LL         "
  f24[ 76][ 7] = " LL         "
  f24[ 76][ 8] = " LL         "
  f24[ 76][ 9] = " LL         "
  f24[ 76][10] = " LL         "
  f24[ 76][11] = " LL         "
  f24[ 76][12] = " LL         "
  f24[ 76][13] = " LL         "
  f24[ 76][14] = " LL         "
  f24[ 76][15] = " LL         "
  f24[ 76][16] = " LL         "
  f24[ 76][17] = " LL         "
  f24[ 76][18] = " LLLLLLLLLL "

  f24[ 77][ 4] = "M         M "
  f24[ 77][ 5] = "MM       MM "
  f24[ 77][ 6] = "MMM     MMM "
  f24[ 77][ 7] = "MMMM   MMMM "
  f24[ 77][ 8] = "MM MM MM MM "
  f24[ 77][ 9] = "MM  MMM  MM "
  f24[ 77][10] = "MM   M   MM "
  f24[ 77][11] = "MM       MM "
  f24[ 77][12] = "MM       MM "
  f24[ 77][13] = "MM       MM "
  f24[ 77][14] = "MM       MM "
  f24[ 77][15] = "MM       MM "
  f24[ 77][16] = "MM       MM "
  f24[ 77][17] = "MM       MM "
  f24[ 77][18] = "MM       MM "

  f24[ 78][ 4] = " NN      NN "
  f24[ 78][ 5] = " NN      NN "
  f24[ 78][ 6] = " NN      NN "
  f24[ 78][ 7] = " NN      NN "
  f24[ 78][ 8] = " NNN     NN "
  f24[ 78][ 9] = " NNNN    NN "
  f24[ 78][10] = " NN NN   NN "
  f24[ 78][11] = " NN  NN  NN "
  f24[ 78][12] = " NN   NN NN "
  f24[ 78][13] = " NN    NNNN "
  f24[ 78][14] = " NN     NNN "
  f24[ 78][15] = " NN      NN "
  f24[ 78][16] = " NN      NN "
  f24[ 78][17] = " NN      NN "
  f24[ 78][18] = " NN      NN "

  f24[ 79][ 4] = "   OOOOOO   "
  f24[ 79][ 5] = "  OO    OO  "
  f24[ 79][ 6] = " OO      OO "
  f24[ 79][ 7] = " OO      OO "
  f24[ 79][ 8] = " OO      OO "
  f24[ 79][ 9] = " OO      OO "
  f24[ 79][10] = " OO      OO "
  f24[ 79][11] = " OO      OO "
  f24[ 79][12] = " OO      OO "
  f24[ 79][13] = " OO      OO "
  f24[ 79][14] = " OO      OO "
  f24[ 79][15] = " OO      OO "
  f24[ 79][16] = " OO      OO "
  f24[ 79][17] = "  OO    OO  "
  f24[ 79][18] = "   OOOOOO   "

  f24[214][ 0] = "  OO   OO   "
  f24[214][ 1] = "  OO   OO   "
  f24[214][ 2] = "  OO   OO   "
  f24[214][ 3] = "            "
  f24[214][ 4] = "   OOOOOO   "
  f24[214][ 5] = "  OO    OO  "
  f24[214][ 6] = " OO      OO "
  f24[214][ 7] = " OO      OO "
  f24[214][ 8] = " OO      OO "
  f24[214][ 9] = " OO      OO "
  f24[214][10] = " OO      OO "
  f24[214][11] = " OO      OO "
  f24[214][12] = " OO      OO "
  f24[214][13] = " OO      OO "
  f24[214][14] = " OO      OO "
  f24[214][15] = " OO      OO "
  f24[214][16] = " OO      OO "
  f24[214][17] = "  OO    OO  "
  f24[214][18] = "   OOOOOO   "

  f24[ 80][ 4] = " PPPPPPPP   "
  f24[ 80][ 5] = " PP     PP  "
  f24[ 80][ 6] = " PP      PP "
  f24[ 80][ 7] = " PP      PP "
  f24[ 80][ 8] = " PP      PP "
  f24[ 80][ 9] = " PP      PP "
  f24[ 80][10] = " PP     PP  "
  f24[ 80][11] = " PPPPPPPP   "
  f24[ 80][12] = " PP         "
  f24[ 80][13] = " PP         "
  f24[ 80][14] = " PP         "
  f24[ 80][15] = " PP         "
  f24[ 80][16] = " PP         "
  f24[ 80][17] = " PP         "
  f24[ 80][18] = " PP         "

  f24[ 81][ 4] = "   QQQQQQ   "
  f24[ 81][ 5] = "  QQ    QQ  "
  f24[ 81][ 6] = " QQ      QQ "
  f24[ 81][ 7] = " QQ      QQ "
  f24[ 81][ 8] = " QQ      QQ "
  f24[ 81][ 9] = " QQ      QQ "
  f24[ 81][10] = " QQ      QQ "
  f24[ 81][11] = " QQ      QQ "
  f24[ 81][12] = " QQ      QQ "
  f24[ 81][13] = " QQ      QQ "
  f24[ 81][14] = " QQ      QQ "
  f24[ 81][15] = " QQ      QQ "
  f24[ 81][16] = " QQ  QQ  QQ "
  f24[ 81][17] = "  QQ  QQQQ  "
  f24[ 81][18] = "   QQQQQQ   "
  f24[ 81][19] = "        QQ  "
  f24[ 81][20] = "         QQ "

  f24[ 82][ 4] = " RRRRRRRR   "
  f24[ 82][ 5] = " RR     RR  "
  f24[ 82][ 6] = " RR      RR "
  f24[ 82][ 7] = " RR      RR "
  f24[ 82][ 8] = " RR      RR "
  f24[ 82][ 9] = " RR      RR "
  f24[ 82][10] = " RR     RR  "
  f24[ 82][11] = " RRRRRRRR   "
  f24[ 82][12] = " RRRR       "
  f24[ 82][13] = " RR RR      "
  f24[ 82][14] = " RR  RR     "
  f24[ 82][15] = " RR   RR    "
  f24[ 82][16] = " RR    RR   "
  f24[ 82][17] = " RR     RR  "
  f24[ 82][18] = " RR      RR "

  f24[ 83][ 4] = "   SSSSSS   "
  f24[ 83][ 5] = "  SS    SS  "
  f24[ 83][ 6] = " SS      SS "
  f24[ 83][ 7] = " SS         "
  f24[ 83][ 8] = " SS         "
  f24[ 83][ 9] = " SS         "
  f24[ 83][10] = "  SS        "
  f24[ 83][11] = "   SSSSSS   "
  f24[ 83][12] = "        SS  "
  f24[ 83][13] = "         SS "
  f24[ 83][14] = "         SS "
  f24[ 83][15] = "         SS "
  f24[ 83][16] = " SS      SS "
  f24[ 83][17] = "  SS    SS  "
  f24[ 83][18] = "   SSSSSS   "

  f24[ 84][ 4] = " TTTTTTTTTT "
  f24[ 84][ 5] = "     TT     "
  f24[ 84][ 6] = "     TT     "
  f24[ 84][ 7] = "     TT     "
  f24[ 84][ 8] = "     TT     "
  f24[ 84][ 9] = "     TT     "
  f24[ 84][10] = "     TT     "
  f24[ 84][11] = "     TT     "
  f24[ 84][12] = "     TT     "
  f24[ 84][13] = "     TT     "
  f24[ 84][14] = "     TT     "
  f24[ 84][15] = "     TT     "
  f24[ 84][16] = "     TT     "
  f24[ 84][17] = "     TT     "
  f24[ 84][18] = "     TT     "

  f24[ 85][ 4] = " UU      UU "
  f24[ 85][ 5] = " UU      UU "
  f24[ 85][ 6] = " UU      UU "
  f24[ 85][ 7] = " UU      UU "
  f24[ 85][ 8] = " UU      UU "
  f24[ 85][ 9] = " UU      UU "
  f24[ 85][10] = " UU      UU "
  f24[ 85][11] = " UU      UU "
  f24[ 85][12] = " UU      UU "
  f24[ 85][13] = " UU      UU "
  f24[ 85][14] = " UU      UU "
  f24[ 85][15] = " UU      UU "
  f24[ 85][16] = " UU      UU "
  f24[ 85][17] = "  UU    UU  "
  f24[ 85][18] = "   UUUUUU   "

  f24[220][ 0] = "  UU    UU  "
  f24[220][ 1] = "  UU    UU  "
  f24[220][ 2] = "  UU    UU  "
  f24[220][ 3] = "            "
  f24[220][ 4] = " UU      UU "
  f24[220][ 5] = " UU      UU "
  f24[220][ 6] = " UU      UU "
  f24[220][ 7] = " UU      UU "
  f24[220][ 8] = " UU      UU "
  f24[220][ 9] = " UU      UU "
  f24[220][10] = " UU      UU "
  f24[220][11] = " UU      UU "
  f24[220][12] = " UU      UU "
  f24[220][13] = " UU      UU "
  f24[220][14] = " UU      UU "
  f24[220][15] = " UU      UU "
  f24[220][16] = " UU      UU "
  f24[220][17] = "  UU    UU  "
  f24[220][18] = "   UUUUUU   "

  f24[ 86][ 4] = " VV      VV "
  f24[ 86][ 5] = " VV      VV "
  f24[ 86][ 6] = " VV      VV "
  f24[ 86][ 7] = " VV      VV "
  f24[ 86][ 8] = "  VV    VV  "
  f24[ 86][ 9] = "  VV    VV  "
  f24[ 86][10] = "  VV    VV  "
  f24[ 86][11] = "   VV  VV   "
  f24[ 86][12] = "   VV  VV   "
  f24[ 86][13] = "   VV  VV   "
  f24[ 86][14] = "    VVVV    "
  f24[ 86][15] = "    VVVV    "
  f24[ 86][16] = "    VVVV    "
  f24[ 86][17] = "     VV     "
  f24[ 86][18] = "     VV     "

  f24[ 87][ 4] = "WW       WW "
  f24[ 87][ 5] = "WW       WW "
  f24[ 87][ 6] = "WW       WW "
  f24[ 87][ 7] = "WW       WW "
  f24[ 87][ 8] = "WW       WW "
  f24[ 87][ 9] = "WW       WW "
  f24[ 87][10] = "WW       WW "
  f24[ 87][11] = "WW       WW "
  f24[ 87][12] = "WW   W   WW "
  f24[ 87][13] = "WW  WWW  WW "
  f24[ 87][14] = "WW WW WW WW "
  f24[ 87][15] = "WWWW   WWWW "
  f24[ 87][16] = "WWW     WWW "
  f24[ 87][17] = "WW       WW "
  f24[ 87][18] = "W         W "

  f24[ 88][ 4] = " XX      XX "
  f24[ 88][ 5] = " XX      XX "
  f24[ 88][ 6] = "  XX    XX  "
  f24[ 88][ 7] = "  XX    XX  "
  f24[ 88][ 8] = "   XX  XX   "
  f24[ 88][ 9] = "   XX  XX   "
  f24[ 88][10] = "    XXXX    "
  f24[ 88][11] = "     XX     "
  f24[ 88][12] = "    XXXX    "
  f24[ 88][13] = "   XX  XX   "
  f24[ 88][14] = "   XX  XX   "
  f24[ 88][15] = "  XX    XX  "
  f24[ 88][16] = "  XX    XX  "
  f24[ 88][17] = " XX      XX "
  f24[ 88][18] = " XX      XX "

  f24[ 89][ 4] = " YY      YY "
  f24[ 89][ 5] = " YY      YY "
  f24[ 89][ 6] = "  YY    YY  "
  f24[ 89][ 7] = "  YY    YY  "
  f24[ 89][ 8] = "   YY  YY   "
  f24[ 89][ 9] = "   YY  YY   "
  f24[ 89][10] = "    YYYY    "
  f24[ 89][11] = "     YY     "
  f24[ 89][12] = "     YY     "
  f24[ 89][13] = "     YY     "
  f24[ 89][14] = "     YY     "
  f24[ 89][15] = "     YY     "
  f24[ 89][16] = "     YY     "
  f24[ 89][17] = "     YY     "
  f24[ 89][18] = "     YY     "

  f24[ 90][ 4] = " ZZZZZZZZZZ "
  f24[ 90][ 5] = "         ZZ "
  f24[ 90][ 6] = "         ZZ "
  f24[ 90][ 7] = "         ZZ "
  f24[ 90][ 8] = "        ZZ  "
  f24[ 90][ 9] = "       ZZ   "
  f24[ 90][10] = "      ZZ    "
  f24[ 90][11] = "     ZZ     "
  f24[ 90][12] = "    ZZ      "
  f24[ 90][13] = "   ZZ       "
  f24[ 90][14] = "  ZZ        "
  f24[ 90][15] = " ZZ         "
  f24[ 90][16] = " ZZ         "
  f24[ 90][17] = " ZZ         "
  f24[ 90][18] = " ZZZZZZZZZZ "

  f24[ 91][ 4] = "   [[[[[    "
  f24[ 91][ 5] = "   [[       "
  f24[ 91][ 6] = "   [[       "
  f24[ 91][ 7] = "   [[       "
  f24[ 91][ 8] = "   [[       "
  f24[ 91][ 9] = "   [[       "
  f24[ 91][10] = "   [[       "
  f24[ 91][11] = "   [[       "
  f24[ 91][12] = "   [[       "
  f24[ 91][13] = "   [[       "
  f24[ 91][14] = "   [[       "
  f24[ 91][15] = "   [[       "
  f24[ 91][16] = "   [[       "
  f24[ 91][17] = "   [[       "
  f24[ 91][18] = "   [[[[[    "

  f24[ 92][ 5] = "  //        "
  f24[ 92][ 6] = "  //        "
  f24[ 92][ 7] = "   //       "
  f24[ 92][ 8] = "   //       "
  f24[ 92][ 9] = "    //      "
  f24[ 92][10] = "    //      "
  f24[ 92][11] = "     //     "
  f24[ 92][12] = "     //     "
  f24[ 92][13] = "      //    "
  f24[ 92][14] = "      //    "
  f24[ 92][15] = "       //   "
  f24[ 92][16] = "       //   "
  f24[ 92][17] = "        //  "
  f24[ 92][18] = "        //  "

  f24[ 93][ 4] = "   ]]]]]    "
  f24[ 93][ 5] = "      ]]    "
  f24[ 93][ 6] = "      ]]    "
  f24[ 93][ 7] = "      ]]    "
  f24[ 93][ 8] = "      ]]    "
  f24[ 93][ 9] = "      ]]    "
  f24[ 93][10] = "      ]]    "
  f24[ 93][11] = "      ]]    "
  f24[ 93][12] = "      ]]    "
  f24[ 93][13] = "      ]]    "
  f24[ 93][14] = "      ]]    "
  f24[ 93][15] = "      ]]    "
  f24[ 93][16] = "      ]]    "
  f24[ 93][17] = "      ]]    "
  f24[ 93][18] = "   ]]]]]    "

  f24[ 94][ 1] = "     ^^     "
  f24[ 94][ 2] = "    ^^^^    "
  f24[ 94][ 3] = "   ^^  ^^   "
  f24[ 94][ 4] = "  ^^    ^^  "
  f24[ 94][ 5] = " ^^      ^^ "

  f24[ 95][20] = " __________ "

  f24[ 96][ 3] = "     ``     "
  f24[ 96][ 4] = "     ``     "
  f24[ 96][ 5] = "     ``     "
  f24[ 96][ 6] = "     ``     "
  f24[ 96][ 7] = "      ``    "

  f24[ 97][ 8] = "  aaaaaaa   "
  f24[ 97][ 9] = "        aa  "
  f24[ 97][10] = "         aa "
  f24[ 97][11] = "         aa "
  f24[ 97][12] = "   aaaaaaaa "
  f24[ 97][13] = "  aa     aa "
  f24[ 97][14] = " aa      aa "
  f24[ 97][15] = " aa      aa "
  f24[ 97][16] = " aa      aa "
  f24[ 97][17] = "  aa     aa "
  f24[ 97][18] = "   aaaaaaaa "

  f24[228][ 4] = "  aa    aa  "
  f24[228][ 5] = "  aa    aa  "
  f24[228][ 6] = "  aa    aa  "
  f24[228][ 7] = "            "
  f24[228][ 8] = "  aaaaaaa   "
  f24[228][ 9] = "         aa "
  f24[228][10] = "         aa "
  f24[228][11] = "         aa "
  f24[228][12] = "   aaaaaaaa "
  f24[228][13] = "  aa     aa "
  f24[228][14] = " aa      aa "
  f24[228][15] = " aa      aa "
  f24[228][16] = " aa      aa "
  f24[228][17] = "  aa     aa "
  f24[228][18] = "    aaaaaaa "

  f24[ 98][ 4] = " bb         "
  f24[ 98][ 5] = " bb         "
  f24[ 98][ 6] = " bb         "
  f24[ 98][ 7] = " bb         "
  f24[ 98][ 8] = " bbbbbbbb   "
  f24[ 98][ 9] = " bb     bb  "
  f24[ 98][10] = " bb      bb "
  f24[ 98][11] = " bb      bb "
  f24[ 98][12] = " bb      bb "
  f24[ 98][13] = " bb      bb "
  f24[ 98][14] = " bb      bb "
  f24[ 98][15] = " bb      bb "
  f24[ 98][16] = " bb      bb "
  f24[ 98][17] = " bb     bb  "
  f24[ 98][18] = " bbbbbbbb   "

  f24[ 99][ 8] = "   cccccc   "
  f24[ 99][ 9] = "  cc    cc  "
  f24[ 99][10] = " cc      cc "
  f24[ 99][11] = " cc         "
  f24[ 99][12] = " cc         "
  f24[ 99][13] = " cc         "
  f24[ 99][14] = " cc         "
  f24[ 99][15] = " cc         "
  f24[ 99][16] = " cc      cc "
  f24[ 99][17] = "  cc    cc  "
  f24[ 99][18] = "   cccccc   "

  f24[100][ 4] = "         dd "
  f24[100][ 5] = "         dd "
  f24[100][ 6] = "         dd "
  f24[100][ 7] = "         dd "
  f24[100][ 8] = "   dddddddd "
  f24[100][ 9] = "  dd     dd "
  f24[100][10] = " dd      dd "
  f24[100][11] = " dd      dd "
  f24[100][12] = " dd      dd "
  f24[100][13] = " dd      dd "
  f24[100][14] = " dd      dd "
  f24[100][15] = " dd      dd "
  f24[100][16] = " dd      dd "
  f24[100][17] = "  dd     dd "
  f24[100][18] = "   dddddddd "

  f24[101][ 8] = "   eeeeee   "
  f24[101][ 9] = "  ee    ee  "
  f24[101][10] = " ee      ee "
  f24[101][11] = " ee      ee "
  f24[101][12] = " ee      ee "
  f24[101][13] = " eeeeeeeeee "
  f24[101][14] = " ee         "
  f24[101][15] = " ee         "
  f24[101][16] = " ee         "
  f24[101][17] = "  ee    ee  "
  f24[101][18] = "   eeeeee   "

  f24[102][ 4] = "      fffff "
  f24[102][ 5] = "     ff     "
  f24[102][ 6] = "     ff     "
  f24[102][ 7] = "     ff     "
  f24[102][ 8] = "  ffffffff  "
  f24[102][ 9] = "     ff     "
  f24[102][10] = "     ff     "
  f24[102][11] = "     ff     "
  f24[102][12] = "     ff     "
  f24[102][13] = "     ff     "
  f24[102][14] = "     ff     "
  f24[102][15] = "     ff     "
  f24[102][16] = "     ff     "
  f24[102][17] = "     ff     "
  f24[102][18] = "     ff     "

  f24[103][ 8] = "   gggggggg "
  f24[103][ 9] = "  gg     gg "
  f24[103][10] = " gg      gg "
  f24[103][11] = " gg      gg "
  f24[103][12] = " gg      gg "
  f24[103][13] = " gg      gg "
  f24[103][14] = " gg      gg "
  f24[103][15] = " gg      gg "
  f24[103][16] = " gg      gg "
  f24[103][17] = "  gg     gg "
  f24[103][18] = "   gggggggg "
  f24[103][19] = "         gg "
  f24[103][20] = "         gg "
  f24[103][21] = "        gg  "
  f24[103][22] = "  ggggggg   "

  f24[104][ 4] = " hh         "
  f24[104][ 5] = " hh         "
  f24[104][ 6] = " hh         "
  f24[104][ 7] = " hh         "
  f24[104][ 8] = " hhhhhhhh   "
  f24[104][ 9] = " hh     hh  "
  f24[104][10] = " hh      hh "
  f24[104][11] = " hh      hh "
  f24[104][12] = " hh      hh "
  f24[104][13] = " hh      hh "
  f24[104][14] = " hh      hh "
  f24[104][15] = " hh      hh "
  f24[104][16] = " hh      hh "
  f24[104][17] = " hh      hh "
  f24[104][18] = " hh      hh "

  f24[105][ 4] = "     ii     "
  f24[105][ 5] = "     ii     "
  f24[105][ 6] = "     ii     "
  f24[105][ 7] = "            "
  f24[105][ 8] = "   iiii     "
  f24[105][ 9] = "     ii     "
  f24[105][10] = "     ii     "
  f24[105][11] = "     ii     "
  f24[105][12] = "     ii     "
  f24[105][13] = "     ii     "
  f24[105][14] = "     ii     "
  f24[105][15] = "     ii     "
  f24[105][16] = "     ii     "
  f24[105][17] = "     ii     "
  f24[105][18] = "   iiiiii   "

  f24[106][ 4] = "        jj  "
  f24[106][ 5] = "        jj  "
  f24[106][ 6] = "        jj  "
  f24[106][ 7] = "            "
  f24[106][ 8] = "            "
  f24[106][ 9] = "      jjjj  "
  f24[106][10] = "        jj  "
  f24[106][11] = "        jj  "
  f24[106][12] = "        jj  "
  f24[106][13] = "        jj  "
  f24[106][14] = "        jj  "
  f24[106][15] = "        jj  "
  f24[106][16] = "        jj  "
  f24[106][17] = "        jj  "
  f24[106][18] = "        jj  "
  f24[106][19] = "  jj    jj  "
  f24[106][20] = "  jj    jj  "
  f24[106][21] = "   jj  jj   "
  f24[106][22] = "    jjjj    "

  f24[107][ 4] = " kk         "
  f24[107][ 5] = " kk         "
  f24[107][ 6] = " kk         "
  f24[107][ 7] = " kk         "
  f24[107][ 8] = " kk     kk  "
  f24[107][ 9] = " kk    kk   "
  f24[107][10] = " kk   kk    "
  f24[107][11] = " kk  kk     "
  f24[107][12] = " kk kk      "
  f24[107][13] = " kkkk       "
  f24[107][14] = " kk kk      "
  f24[107][15] = " kk  kk     "
  f24[107][16] = " kk   kk    "
  f24[107][17] = " kk    kk   "
  f24[107][18] = " kk     kk  "

  f24[108][ 4] = "   llll     "
  f24[108][ 5] = "     ll     "
  f24[108][ 5] = "     ll     "
  f24[108][ 6] = "     ll     "
  f24[108][ 7] = "     ll     "
  f24[108][ 8] = "     ll     "
  f24[108][ 9] = "     ll     "
  f24[108][10] = "     ll     "
  f24[108][11] = "     ll     "
  f24[108][12] = "     ll     "
  f24[108][13] = "     ll     "
  f24[108][14] = "     ll     "
  f24[108][15] = "     ll     "
  f24[108][16] = "     ll     "
  f24[108][17] = "     ll     "
  f24[108][18] = "   lllll    "

  f24[109][ 8] = " mmmmmmmm   "
  f24[109][ 9] = " mm  mm mm  "
  f24[109][10] = " mm  mm  mm "
  f24[109][11] = " mm  mm  mm "
  f24[109][12] = " mm  mm  mm "
  f24[109][13] = " mm  mm  mm "
  f24[109][14] = " mm  mm  mm "
  f24[109][15] = " mm  mm  mm "
  f24[109][16] = " mm  mm  mm "
  f24[109][17] = " mm  mm  mm "
  f24[109][18] = " mm  mm  mm "

  f24[110][ 8] = " nnnnnnnn   "
  f24[110][ 9] = " nn     nn  "
  f24[110][10] = " nn      nn "
  f24[110][11] = " nn      nn "
  f24[110][12] = " nn      nn "
  f24[110][13] = " nn      nn "
  f24[110][14] = " nn      nn "
  f24[110][15] = " nn      nn "
  f24[110][16] = " nn      nn "
  f24[110][17] = " nn      nn "
  f24[110][18] = " nn      nn "

  f24[111][ 8] = "   oooooo   "
  f24[111][ 9] = "  oo    oo  "
  f24[111][10] = " oo      oo "
  f24[111][11] = " oo      oo "
  f24[111][12] = " oo      oo "
  f24[111][13] = " oo      oo "
  f24[111][14] = " oo      oo "
  f24[111][15] = " oo      oo "
  f24[111][16] = " oo      oo "
  f24[111][17] = "  oo    oo  "
  f24[111][18] = "   oooooo   "

  f24[246][ 4] = "  oo   oo   "
  f24[246][ 5] = "  oo   oo   "
  f24[246][ 6] = "  oo   oo   "
  f24[246][ 7] = "            "
  f24[246][ 8] = "   oooooo   "
  f24[246][ 9] = "  oo    oo  "
  f24[246][10] = " oo      oo "
  f24[246][11] = " oo      oo "
  f24[246][12] = " oo      oo "
  f24[246][13] = " oo      oo "
  f24[246][14] = " oo      oo "
  f24[246][15] = " oo      oo "
  f24[246][16] = " oo      oo "
  f24[246][17] = "  oo    oo  "
  f24[246][18] = "   oooooo   "

  f24[112][ 8] = " pppppppp   "
  f24[112][ 9] = " pp     pp  "
  f24[112][10] = " pp      pp "
  f24[112][11] = " pp      pp "
  f24[112][12] = " pp      pp "
  f24[112][13] = " pp      pp "
  f24[112][14] = " pp      pp "
  f24[112][15] = " pp      pp "
  f24[112][16] = " pp      pp "
  f24[112][17] = " pp     pp  "
  f24[112][18] = " pppppppp   "
  f24[112][19] = " pp         "
  f24[112][20] = " pp         "
  f24[112][21] = " pp         "
  f24[112][22] = " pp         "

  f24[113][ 8] = "   qqqqqqqq "
  f24[113][ 9] = "  qq     qq "
  f24[113][10] = " qq      qq "
  f24[113][11] = " qq      qq "
  f24[113][12] = " qq      qq "
  f24[113][13] = " qq      qq "
  f24[113][14] = " qq      qq "
  f24[113][15] = " qq      qq "
  f24[113][16] = " qq      qq "
  f24[113][17] = "  qq     qq "
  f24[113][18] = "   qqqqqqqq "
  f24[113][19] = "         qq "
  f24[113][20] = "         qq "
  f24[113][21] = "         qq "
  f24[113][22] = "         qq "

  f24[114][ 8] = " rr  rrrrrr "
  f24[114][ 9] = " rr rr      "
  f24[114][10] = " rrrr       "
  f24[114][11] = " rrr        "
  f24[114][12] = " rr         "
  f24[114][13] = " rr         "
  f24[114][14] = " rr         "
  f24[114][15] = " rr         "
  f24[114][16] = " rr         "
  f24[114][17] = " rr         "
  f24[114][18] = " rr         "

  f24[115][ 8] = "   sssssss  "
  f24[115][ 9] = " ss      ss "
  f24[115][10] = " ss         "
  f24[115][11] = " ss         "
  f24[115][12] = " ss         "
  f24[115][13] = "  ssssssss  "
  f24[115][14] = "         ss "
  f24[115][15] = "         ss "
  f24[115][16] = "         ss "
  f24[115][17] = " ss      ss "
  f24[115][18] = "  ssssssss  "

  f24[116][ 4] = "    tt      "
  f24[116][ 5] = "    tt      "
  f24[116][ 6] = "    tt      "
  f24[116][ 7] = "    tt      "
  f24[116][ 8] = " tttttttt   "
  f24[116][ 9] = "    tt      "
  f24[116][10] = "    tt      "
  f24[116][11] = "    tt      "
  f24[116][12] = "    tt      "
  f24[116][13] = "    tt      "
  f24[116][14] = "    tt      "
  f24[116][15] = "    tt      "
  f24[116][16] = "    tt      "
  f24[116][17] = "    tt      "
  f24[116][18] = "     ttttt  "

  f24[117][ 8] = " uu      uu "
  f24[117][ 9] = " uu      uu "
  f24[117][10] = " uu      uu "
  f24[117][11] = " uu      uu "
  f24[117][12] = " uu      uu "
  f24[117][13] = " uu      uu "
  f24[117][14] = " uu      uu "
  f24[117][15] = " uu      uu "
  f24[117][16] = " uu      uu "
  f24[117][17] = "  uu     uu "
  f24[117][18] = "   uuuuuuuu "

  f24[252][ 4] = "  uu    uu  "
  f24[252][ 5] = "  uu    uu  "
  f24[252][ 6] = "  uu    uu  "
  f24[252][ 7] = "            "
  f24[252][ 8] = " uu      uu "
  f24[252][ 9] = " uu      uu "
  f24[252][10] = " uu      uu "
  f24[252][11] = " uu      uu "
  f24[252][12] = " uu      uu "
  f24[252][13] = " uu      uu "
  f24[252][14] = " uu      uu "
  f24[252][15] = " uu      uu "
  f24[252][16] = " uu      uu "
  f24[252][17] = "  uu     uu "
  f24[252][18] = "   uuuuuuuu "

  f24[118][ 8] = " vv      vv "
  f24[118][ 9] = " vv      vv "
  f24[118][10] = " vv      vv "
  f24[118][11] = "  vv    vv  "
  f24[118][12] = "  vv    vv  "
  f24[118][13] = "   vv  vv   "
  f24[118][14] = "   vv  vv   "
  f24[118][15] = "    vvvv    "
  f24[118][16] = "    vvvv    "
  f24[118][17] = "     vv     "
  f24[118][18] = "     vv     "

  f24[119][ 8] = " ww      ww "
  f24[119][ 9] = " ww      ww "
  f24[119][10] = " ww      ww "
  f24[119][11] = " ww  ww  ww "
  f24[119][12] = " ww  ww  ww "
  f24[119][13] = " ww  ww  ww "
  f24[119][14] = " ww  ww  ww "
  f24[119][15] = " ww  ww  ww "
  f24[119][16] = " ww  ww  ww "
  f24[119][17] = " wwwwwwwwww "
  f24[119][18] = "  wwwwwwww  "

  f24[120][ 8] = " xx      xx "
  f24[120][ 9] = " xx      xx "
  f24[120][10] = "  xx    xx  "
  f24[120][11] = "   xx  xx   "
  f24[120][12] = "    xxxx    "
  f24[120][13] = "     xx     "
  f24[120][14] = "    xxxx    "
  f24[120][15] = "   xx  xx   "
  f24[120][16] = "  xx    xx  "
  f24[120][17] = " xx      xx "
  f24[120][18] = " xx      xx "

  f24[121][ 8] = " yy      yy "
  f24[121][ 9] = " yy      yy "
  f24[121][10] = " yy      yy "
  f24[121][11] = " yy      yy "
  f24[121][12] = " yy      yy "
  f24[121][13] = " yy      yy "
  f24[121][14] = " yy      yy "
  f24[121][15] = " yy      yy "
  f24[121][16] = " yy      yy "
  f24[121][17] = "  yy    yyy "
  f24[121][18] = "   yyyyyyyy "
  f24[121][19] = "         yy "
  f24[121][20] = "         yy "
  f24[121][21] = "        yy  "
  f24[121][22] = "  yyyyyyy   "

  f24[122][ 8] = " zzzzzzzzzz "
  f24[122][ 9] = "         zz "
  f24[122][10] = "        zz  "
  f24[122][11] = "       zz   "
  f24[122][12] = "      zz    "
  f24[122][13] = "     zz     "
  f24[122][14] = "    zz      "
  f24[122][15] = "   zz       "
  f24[122][16] = "  zz        "
  f24[122][17] = " zz         "
  f24[122][18] = " zzzzzzzzzz "

  f24[123][ 4] = "      {{{   "
  f24[123][ 5] = "     {{     "
  f24[123][ 6] = "    {{      "
  f24[123][ 7] = "    {{      "
  f24[123][ 8] = "    {{      "
  f24[123][ 9] = "    {{      "
  f24[123][10] = "    {{      "
  f24[123][11] = "  {{{       "
  f24[123][12] = "    {{      "
  f24[123][13] = "    {{      "
  f24[123][14] = "    {{      "
  f24[123][15] = "    {{      "
  f24[123][16] = "    {{      "
  f24[123][17] = "     {{     "
  f24[123][18] = "      {{{   "

  f24[124][ 4] = "     ||     "
  f24[124][ 5] = "     ||     "
  f24[124][ 6] = "     ||     "
  f24[124][ 7] = "     ||     "
  f24[124][ 8] = "     ||     "
  f24[124][ 9] = "     ||     "
  f24[124][10] = "     ||     "
  f24[124][11] = "     ||     "
  f24[124][12] = "     ||     "
  f24[124][13] = "     ||     "
  f24[124][14] = "     ||     "
  f24[124][15] = "     ||     "
  f24[124][16] = "     ||     "
  f24[124][17] = "     ||     "
  f24[124][18] = "     ||     "

  f24[125][ 4] = "  }}}       "
  f24[125][ 5] = "    }}      "
  f24[125][ 6] = "     }}     "
  f24[125][ 7] = "     }}     "
  f24[125][ 8] = "     }}     "
  f24[125][ 9] = "     }}     "
  f24[125][10] = "     }}     "
  f24[125][11] = "      }}}   "
  f24[125][12] = "     }}     "
  f24[125][13] = "     }}     "
  f24[125][14] = "     }}     "
  f24[125][15] = "     }}     "
  f24[125][16] = "     }}     "
  f24[125][17] = "    }}      "
  f24[125][18] = "  }}}       "

  f24[126][ 1] = "  ~~~~   ~~ "
  f24[126][ 2] = " ~~  ~~  ~~ "
  f24[126][ 3] = " ~~  ~~  ~~ "
  f24[126][ 4] = " ~~  ~~  ~~ "
  f24[126][ 5] = " ~~   ~~~~  "

  f24[164][ 4] = "     eee    "
  f24[164][ 5] = "    eeeee   "
  f24[164][ 6] = "   ee   ee  "
  f24[164][ 7] = "  ee    ee  "
  f24[164][ 8] = " ee         "
  f24[164][ 9] = " ee         "
  f24[164][10] = "eeeeeeeeeee "
  f24[164][11] = " ee         "
  f24[164][12] = " ee         "
  f24[164][13] = "eeeeeeeeeee "
  f24[164][14] = " ee         "
  f24[164][15] = " ee         "
  f24[164][16] = "  ee    ee  "
  f24[164][17] = "   ee   ee  "
  f24[164][18] = "    eeeee   "
  f24[164][19] = "     eee    "

  f24[167][ 3] = "  ppppppp   "
  f24[167][ 4] = " pp     pp  "
  f24[167][ 5] = " pp         "
  f24[167][ 6] = " pp         "
  f24[167][ 7] = "  pppppp    "
  f24[167][ 8] = " pp    pp   "
  f24[167][ 9] = " pp     pp  "
  f24[167][10] = " pp     pp  "
  f24[167][11] = " pp     pp  "
  f24[167][12] = " pp     pp  "
  f24[167][13] = "  pp    pp  "
  f24[167][14] = "   pppppp   "
  f24[167][15] = "        pp  "
  f24[167][16] = "        pp  "
  f24[167][17] = " pp     pp  "
  f24[167][18] = "  ppppppp   "

  f24[176][ 1] = "   oooooo   "
  f24[176][ 2] = "  oo    oo  "
  f24[176][ 3] = "  oo    oo  "
  f24[176][ 4] = "  oo    oo  "
  f24[176][ 5] = "  oo    oo  "
  f24[176][ 6] = "   oooooo   "

  f24[181][ 8] = " mm      mm "
  f24[181][ 9] = " mm      mm "
  f24[181][10] = " mm      mm "
  f24[181][11] = " mm      mm "
  f24[181][12] = " mm      mm "
  f24[181][13] = " mm      mm "
  f24[181][14] = " mm      mm "
  f24[181][15] = " mm      mm "
  f24[181][16] = " mmm     mm "
  f24[181][17] = " mmmm  mmmm "
  f24[181][18] = " mm mmmm mm "
  f24[181][19] = " mm         "
  f24[181][20] = " mm         "
  f24[181][21] = " mm         "
  f24[181][22] = " mm         "

  f24[223][ 4] = "  ssssss    "
  f24[223][ 5] = " ss    ss   "
  f24[223][ 6] = " ss     ss  "
  f24[223][ 7] = " ss     ss  "
  f24[223][ 8] = " ss     ss  "
  f24[223][ 9] = " ss    ss   "
  f24[223][10] = " ss sssss   "
  f24[223][11] = " ss     ss  "
  f24[223][12] = " ss      ss "
  f24[223][13] = " ss      ss "
  f24[223][14] = " ss      ss "
  f24[223][15] = " ss      ss "
  f24[223][16] = " sss     ss "
  f24[223][17] = " ssss   ss  "
  f24[223][18] = " ss sssss   "
  f24[223][19] = " ss         "
  f24[223][20] = " ss         "
  f24[223][21] = " ss         "

  for b := 0; b <= 255; b++ {
    for l := 0; l < 28; l++ {
      f28[b][l] = "              "
    }
  }

  f28[ 33][ 5] = "     !!!      "
  f28[ 33][ 6] = "     !!!      "
  f28[ 33][ 7] = "     !!!      "
  f28[ 33][ 8] = "     !!!      "
  f28[ 33][ 9] = "     !!!      "
  f28[ 33][10] = "     !!!      "
  f28[ 33][11] = "     !!!      "
  f28[ 33][12] = "     !!!      "
  f28[ 33][13] = "     !!!      "
  f28[ 33][14] = "     !!!      "
  f28[ 33][15] = "     !!!      "
  f28[ 33][16] = "     !!!      "
  f28[ 33][17] = "     !!!      "
  f28[ 33][18] = "     !!!      "
  f28[ 33][19] = "              "
  f28[ 33][20] = "              "
  f28[ 33][21] = "     !!!      "
  f28[ 33][22] = "     !!!      "
  f28[ 33][23] = "     !!!      "

  f28[ 34][ 3] = "  ***   ***   "
  f28[ 34][ 4] = "  ***   ***   "
  f28[ 34][ 5] = "  ***   ***   "
  f28[ 34][ 6] = "  ***   ***   "
  f28[ 34][ 7] = "  ***   ***   "
  f28[ 34][ 8] = "  ***   ***   "

  f28[ 35][ 5] = "   ###  ###   "
  f28[ 35][ 6] = "   ###  ###   "
  f28[ 35][ 7] = "   ###  ###   "
  f28[ 35][ 8] = "   ###  ###   "
  f28[ 35][ 9] = " ############ "
  f28[ 35][10] = " ############ "
  f28[ 35][11] = "   ###  ###   "
  f28[ 35][12] = "   ###  ###   "
  f28[ 35][13] = "   ###  ###   "
  f28[ 35][14] = "   ###  ###   "
  f28[ 35][15] = "   ###  ###   "
  f28[ 35][16] = "   ###  ###   "
  f28[ 35][17] = " ############ "
  f28[ 35][18] = " ############ "
  f28[ 35][19] = "   ###  ###   "
  f28[ 35][20] = "   ###  ###   "
  f28[ 35][21] = "   ###  ###   "
  f28[ 35][22] = "   ###  ###   "

  f28[ 36][ 5] = "      $$      "
  f28[ 36][ 6] = "      $$      "
  f28[ 36][ 7] = "   $$$$$$$$   "
  f28[ 36][ 8] = "  $$$$$$$$$$  "
  f28[ 36][ 9] = " $$$  $$  $$$ "
  f28[ 36][10] = " $$$  $$  $$$ "
  f28[ 36][11] = " $$$  $$      "
  f28[ 36][12] = " $$$  $$      "
  f28[ 36][13] = "  $$$$$$$$$   "
  f28[ 36][14] = "   $$$$$$$$$  "
  f28[ 36][15] = "      $$  $$$ "
  f28[ 36][16] = "      $$  $$$ "
  f28[ 36][17] = " $$$  $$  $$$ "
  f28[ 36][18] = " $$$  $$  $$$ "
  f28[ 36][19] = "  $$$$$$$$$$  "
  f28[ 36][21] = "   $$$$$$$$   "
  f28[ 36][22] = "      $$      "
  f28[ 36][22] = "      $$      "

  f28[ 37][ 6] = "  %%%%   %%%  "
  f28[ 37][ 9] = " %%%%%%  %%%  "
  f28[ 37][ 8] = " %%  %% %%%   "
  f28[ 37][ 9] = " %%  %% %%%   "
  f28[ 37][10] = " %%%%%%%%%%   "
  f28[ 37][11] = "  %%%% %%%    "
  f28[ 37][12] = "      %%%     "
  f28[ 37][13] = "      %%%     "
  f28[ 37][16] = "     %%%      "
  f28[ 37][17] = "     %%%      "
  f28[ 37][18] = "    %%%       "
  f28[ 37][19] = "    %%%       "
  f28[ 37][20] = "   %%% %%%%   "
  f28[ 37][21] = "   %%%%%%%%%  "
  f28[ 37][22] = "  %%% %%  %%  "
  f28[ 37][23] = "  %%% %%  %%  "
  f28[ 37][24] = " %%%  %%%%%%  "
  f28[ 37][25] = " %%%   %%%%   "

  f28[ 38][ 5] = "    &&&&      "
  f28[ 38][ 6] = "   &&&&&&     "
  f28[ 38][ 7] = "  &&&  &&&    "
  f28[ 38][ 8] = "  &&&  &&&    "
  f28[ 38][ 9] = "  &&&  &&&    "
  f28[ 38][10] = "  &&&  &&&    "
  f28[ 38][11] = "   &&&&&&     "
  f28[ 38][12] = "    &&&&      "
  f28[ 38][13] = "    &&&       "
  f28[ 38][14] = "   &&&&&  &&& "
  f28[ 38][15] = "  &&& &&& &&& "
  f28[ 38][16] = " &&&   &&&&&  "
  f28[ 38][17] = " &&&    &&&   "
  f28[ 38][18] = " &&&    &&&   "
  f28[ 38][19] = " &&&    &&&   "
  f28[ 38][10] = " &&&   &&&&&  "
  f28[ 38][21] = "  &&&&&&& &&& "
  f28[ 38][22] = "   &&&&&  &&& "

  f28[ 39][ 3] = "      '''     "
  f28[ 39][ 4] = "      '''     "
  f28[ 39][ 5] = "      '''     "
  f28[ 39][ 6] = "      '''     "
  f28[ 39][ 7] = "     '''      "
  f28[ 39][ 8] = "     '''      "

  f28[ 40][ 5] = "       (((    "
  f28[ 40][ 6] = "      (((     "
  f28[ 40][ 7] = "     (((      "
  f28[ 40][ 8] = "    (((       "
  f28[ 40][ 9] = "    (((       "
  f28[ 40][10] = "   (((        "
  f28[ 40][11] = "   (((        "
  f28[ 40][12] = "   (((        "
  f28[ 40][13] = "   (((        "
  f28[ 40][14] = "   (((        "
  f28[ 40][15] = "   (((        "
  f28[ 40][16] = "   (((        "
  f28[ 40][17] = "   (((        "
  f28[ 40][18] = "    (((       "
  f28[ 40][19] = "    (((       "
  f28[ 40][20] = "     (((      "
  f28[ 40][21] = "      (((     "
  f28[ 40][22] = "       (((    "

  f28[ 41][ 5] = "   )))        "
  f28[ 41][ 6] = "    )))       "
  f28[ 41][ 7] = "     )))      "
  f28[ 41][ 8] = "      )))     "
  f28[ 41][ 9] = "      )))     "
  f28[ 41][10] = "       )))    "
  f28[ 41][11] = "       )))    "
  f28[ 41][12] = "       )))    "
  f28[ 41][13] = "       )))    "
  f28[ 41][14] = "       )))    "
  f28[ 41][15] = "       )))    "
  f28[ 41][16] = "       )))    "
  f28[ 41][17] = "       )))    "
  f28[ 41][18] = "      )))     "
  f28[ 41][19] = "      )))     "
  f28[ 41][20] = "     )))      "
  f28[ 41][21] = "    )))       "
  f28[ 41][22] = "   )))        "

  f28[ 42][ 8] = " ***     ***  "
  f28[ 42][ 9] = "  ***   ***   "
  f28[ 42][10] = "   *** ***    "
  f28[ 42][11] = "    *****     "
  f28[ 42][12] = "     ***      "
  f28[ 42][13] = " ***********  "
  f28[ 42][14] = " ***********  "
  f28[ 42][15] = "     ***      "
  f28[ 42][16] = "    *****     "
  f28[ 42][17] = "   *** ***    "
  f28[ 42][18] = "  ***   ***   "
  f28[ 42][19] = " ***     ***  "

  f28[ 43][ 8] = "     +++      "
  f28[ 43][ 9] = "     +++      "
  f28[ 43][10] = "     +++      "
  f28[ 43][11] = "     +++      "
  f28[ 43][12] = "     +++      "
  f28[ 43][13] = " ++++++++++++ "
  f28[ 43][14] = " ++++++++++++ "
  f28[ 43][15] = "     +++      "
  f28[ 43][16] = "     +++      "
  f28[ 43][17] = "     +++      "
  f28[ 43][18] = "     +++      "
  f28[ 43][19] = "     +++      "

  f28[ 44][22] = "     ,,,      "
  f28[ 44][23] = "     ,,,      "
  f28[ 44][24] = "     ,,,      "
  f28[ 44][25] = "     ,,,      "
  f28[ 44][26] = "    ,,,       "
  f28[ 44][27] = "   ,,,        "

  f28[ 45][15] = " -----------  "
  f28[ 45][16] = " -----------  "

  f28[ 46][19] = "     ...      "
  f28[ 46][20] = "    .....     "
  f28[ 46][21] = "    .....     "
  f28[ 46][22] = "     ...      "

  f28[ 47][ 5] = "         ///  "
  f28[ 47][ 6] = "         ///  "
  f28[ 47][ 7] = "        ///   "
  f28[ 47][ 8] = "        ///   "
  f28[ 47][ 9] = "       ///    "
  f28[ 47][10] = "       ///    "
  f28[ 47][11] = "      ///     "
  f28[ 47][12] = "      ///     "
  f28[ 47][13] = "     ///      "
  f28[ 47][14] = "     ///      "
  f28[ 47][15] = "    ///       "
  f28[ 47][16] = "    ///       "
  f28[ 47][17] = "   ///        "
  f28[ 47][18] = "   ///        "
  f28[ 47][19] = "  ///         "
  f28[ 47][10] = "  ///         "
  f28[ 47][21] = " ///          "
  f28[ 47][22] = " ///          "

  f28[ 48][ 5] = "   0000000    "
  f28[ 48][ 6] = "  000000000   "
  f28[ 48][ 7] = " 000     000  "
  f28[ 48][ 8] = " 000     000  "
  f28[ 48][ 9] = " 000     000  "
  f28[ 48][10] = " 000    0000  "
  f28[ 48][11] = " 000   00000  "
  f28[ 48][12] = " 000  000000  "
  f28[ 48][13] = " 000 000 000  "
  f28[ 48][14] = " 000000  000  "
  f28[ 48][15] = " 00000   000  "
  f28[ 48][16] = " 0000    000  "
  f28[ 48][17] = " 000     000  "
  f28[ 48][18] = " 000     000  "
  f28[ 48][19] = " 000     000  "
  f28[ 48][20] = " 000     000  "
  f28[ 48][21] = "  000000000   "
  f28[ 48][22] = "   0000000    "

  f28[ 49][ 5] = "     111      "
  f28[ 49][ 6] = "    1111      "
  f28[ 49][ 7] = "   11111      "
  f28[ 49][ 8] = "  111111      "
  f28[ 49][ 9] = "  111111      "
  f28[ 49][10] = "     111      "
  f28[ 49][11] = "     111      "
  f28[ 49][12] = "     111      "
  f28[ 49][13] = "     111      "
  f28[ 49][14] = "     111      "
  f28[ 49][15] = "     111      "
  f28[ 49][16] = "     111      "
  f28[ 49][17] = "     111      "
  f28[ 49][18] = "     111      "
  f28[ 49][19] = "     111      "
  f28[ 49][20] = "     111      "
  f28[ 49][21] = "  111111111   "
  f28[ 49][22] = "  111111111   "

  f28[ 50][ 5] = "   2222222    "
  f28[ 50][ 6] = "  222222222   "
  f28[ 50][ 7] = " 222     222  "
  f28[ 50][ 8] = " 222     222  "
  f28[ 50][ 9] = " 222     222  "
  f28[ 50][10] = " 222     222  "
  f28[ 50][11] = "         222  "
  f28[ 50][12] = "         222  "
  f28[ 50][13] = "        222   "
  f28[ 50][14] = "       222    "
  f28[ 50][15] = "      222     "
  f28[ 50][16] = "     222      "
  f28[ 50][17] = "    222       "
  f28[ 50][18] = "   222        "
  f28[ 50][19] = "  222         "
  f28[ 50][20] = " 222          "
  f28[ 50][21] = " 22222222222  "
  f28[ 50][22] = " 22222222222  "

  f28[ 51][ 5] = "   3333333    "
  f28[ 51][ 6] = "  333333333   "
  f28[ 51][ 7] = " 333     333  "
  f28[ 51][ 8] = " 333     333  "
  f28[ 51][ 9] = "         333  "
  f28[ 51][10] = "         333  "
  f28[ 51][11] = "         333  "
  f28[ 51][12] = "         333  "
  f28[ 51][13] = "    3333333   "
  f28[ 51][14] = "    3333333   "
  f28[ 51][15] = "         333  "
  f28[ 51][16] = "         333  "
  f28[ 51][17] = "         333  "
  f28[ 51][18] = "         333  "
  f28[ 51][19] = " 333     333  "
  f28[ 51][20] = " 333     333  "
  f28[ 51][21] = "  333333333   "
  f28[ 51][22] = "   3333333    "

  f28[ 52][ 5] = " 444    444   "
  f28[ 52][ 6] = " 444    444   "
  f28[ 52][ 7] = " 444    444   "
  f28[ 52][ 8] = " 444    444   "
  f28[ 52][ 9] = " 444    444   "
  f28[ 52][10] = " 444    444   "
  f28[ 52][11] = " 444    444   "
  f28[ 52][12] = " 444    444   "
  f28[ 52][13] = " 44444444444  "
  f28[ 52][14] = " 44444444444  "
  f28[ 52][15] = "        444   "
  f28[ 52][16] = "        444   "
  f28[ 52][17] = "        444   "
  f28[ 52][18] = "        444   "
  f28[ 52][19] = "        444   "
  f28[ 52][20] = "        444   "
  f28[ 52][21] = "        444   "
  f28[ 52][22] = "        444   "

  f28[ 53][ 5] = " 55555555555  "
  f28[ 53][ 6] = " 55555555555  "
  f28[ 53][ 7] = " 555          "
  f28[ 53][ 8] = " 555          "
  f28[ 53][19] = " 555          "
  f28[ 53][10] = " 555          "
  f28[ 53][11] = " 555          "
  f28[ 53][12] = " 555          "
  f28[ 53][13] = " 555555555    "
  f28[ 53][14] = " 5555555555   "
  f28[ 53][15] = "         555  "
  f28[ 53][16] = "         555  "
  f28[ 53][17] = "         555  "
  f28[ 53][18] = "         555  "
  f28[ 53][19] = " 555     555  "
  f28[ 53][20] = " 555     555  "
  f28[ 53][21] = "  555555555   "
  f28[ 53][22] = "   5555555    "

  f28[ 54][ 5] = "   66666666   "
  f28[ 54][ 6] = "  666666666   "
  f28[ 54][ 7] = " 666          "
  f28[ 54][ 8] = " 666          "
  f28[ 54][ 9] = " 666          "
  f28[ 54][10] = " 666          "
  f28[ 54][11] = " 666          "
  f28[ 54][12] = " 666          "
  f28[ 54][13] = " 666666666    "
  f28[ 54][14] = " 6666666666   "
  f28[ 54][15] = " 666     666  "
  f28[ 54][16] = " 666     666  "
  f28[ 54][17] = " 666     666  "
  f28[ 54][18] = " 666     666  "
  f28[ 54][19] = " 666     666  "
  f28[ 54][20] = " 666     666  "
  f28[ 54][21] = "  666666666   "
  f28[ 54][22] = "   6666666    "

  f28[ 55][ 5] = " 77777777777  "
  f28[ 55][ 6] = " 77777777777  "
  f28[ 55][ 7] = " 777     777  "
  f28[ 55][ 8] = " 777     777  "
  f28[ 55][ 9] = "        777   "
  f28[ 55][10] = "        777   "
  f28[ 55][12] = "       777    "
  f28[ 55][13] = "       777    "
  f28[ 55][14] = "      777     "
  f28[ 55][15] = "      777     "
  f28[ 55][16] = "     777      "
  f28[ 55][17] = "     777      "
  f28[ 55][18] = "     777      "
  f28[ 55][19] = "     777      "
  f28[ 55][20] = "     777      "
  f28[ 55][21] = "     777      "
  f28[ 55][22] = "     777      "

  f28[ 56][ 5] = "   8888888    "
  f28[ 56][ 6] = "  888888888   "
  f28[ 56][ 7] = " 888     888  "
  f28[ 56][ 8] = " 888     888  "
  f28[ 56][ 9] = " 888     888  "
  f28[ 56][10] = " 888     888  "
  f28[ 56][11] = " 888     888  "
  f28[ 56][12] = " 888     888  "
  f28[ 56][13] = "  888888888   "
  f28[ 56][14] = "  888888888   "
  f28[ 56][15] = " 888     888  "
  f28[ 56][16] = " 888     888  "
  f28[ 56][17] = " 888     888  "
  f28[ 56][18] = " 888     888  "
  f28[ 56][19] = " 888     888  "
  f28[ 56][20] = " 888     888  "
  f28[ 56][21] = "  888888888   "
  f28[ 56][22] = "   8888888    "

  f28[ 57][ 5] = "   9999999    "
  f28[ 57][ 6] = "  999999999   "
  f28[ 57][ 7] = " 999     999  "
  f28[ 57][ 8] = " 999     999  "
  f28[ 57][ 9] = " 999     999  "
  f28[ 57][10] = " 999     999  "
  f28[ 57][11] = " 999     999  "
  f28[ 57][12] = " 999     999  "
  f28[ 57][13] = "  9999999999  "
  f28[ 57][14] = "   999999999  "
  f28[ 57][15] = "         999  "
  f28[ 57][16] = "         999  "
  f28[ 57][17] = "         999  "
  f28[ 57][18] = "         999  "
  f28[ 57][19] = "         999  "
  f28[ 57][20] = "         999  "
  f28[ 57][21] = "  999999999   "
  f28[ 57][22] = "  99999999    "

  f28[ 58][13] = "     :::      "
  f28[ 58][14] = "    :::::     "
  f28[ 58][15] = "    :::::     "
  f28[ 58][16] = "     :::      "
  f28[ 58][17] = "              "
  f28[ 58][18] = "              "
  f28[ 58][19] = "     :::      "
  f28[ 58][20] = "    :::::     "
  f28[ 58][21] = "    :::::     "
  f28[ 58][22] = "     :::      "

  f28[ 59][13] = "     ;;;      "
  f28[ 59][14] = "    ;;;;;     "
  f28[ 59][15] = "     ;;;      "
  f28[ 59][16] = "              "
  f28[ 59][17] = "              "
  f28[ 59][18] = "              "
  f28[ 59][19] = "              "
  f28[ 59][20] = "              "
  f28[ 59][21] = "     ;;;      "
  f28[ 59][22] = "     ;;;      "
  f28[ 59][23] = "     ;;;      "
  f28[ 59][24] = "     ;;;      "
  f28[ 59][25] = "     ;;;      "
  f28[ 59][26] = "    ;;;       "
  f28[ 59][27] = "   ;;;        "

  f28[ 60][ 5] = "         <<<  "
  f28[ 60][ 6] = "        <<<   "
  f28[ 60][ 7] = "       <<<    "
  f28[ 60][ 8] = "      <<<     "
  f28[ 60][ 9] = "     <<<      "
  f28[ 60][10] = "    <<<       "
  f28[ 60][11] = "   <<<        "
  f28[ 60][12] = "  <<<         "
  f28[ 60][13] = " <<<          "
  f28[ 60][14] = " <<<          "
  f28[ 60][15] = "  <<<         "
  f28[ 60][16] = "   <<<        "
  f28[ 60][17] = "    <<<       "
  f28[ 60][18] = "     <<<      "
  f28[ 60][19] = "      <<<     "
  f28[ 60][20] = "       <<<    "
  f28[ 60][21] = "        <<<   "
  f28[ 60][22] = "         <<<  "

  f28[ 61][12] = " ===========  "
  f28[ 61][13] = " ===========  "
  f28[ 61][14] = "              "
  f28[ 61][15] = "              "
  f28[ 61][16] = "              "
  f28[ 61][17] = "              "
  f28[ 61][18] = " ===========  "
  f28[ 61][19] = " ===========  "

  f28[ 62][ 5] = " >>>          "
  f28[ 62][ 6] = "  >>>         "
  f28[ 62][ 7] = "   >>>        "
  f28[ 62][ 8] = "    >>>       "
  f28[ 62][ 9] = "     >>>      "
  f28[ 62][10] = "      >>>     "
  f28[ 62][11] = "       >>>    "
  f28[ 62][12] = "        >>>   "
  f28[ 62][13] = "         >>>  "
  f28[ 62][14] = "         >>>  "
  f28[ 62][15] = "        >>>   "
  f28[ 62][16] = "       >>>    "
  f28[ 62][17] = "      >>>     "
  f28[ 62][18] = "     >>>      "
  f28[ 62][19] = "    >>>       "
  f28[ 62][20] = "   >>>        "
  f28[ 62][21] = "  >>>         "
  f28[ 62][22] = " >>>          "

  f28[ 63][ 5] = "   ???????    "
  f28[ 63][ 6] = "  ?????????   "
  f28[ 63][ 7] = " ???     ???  "
  f28[ 63][ 8] = " ???     ???  "
  f28[ 63][ 9] = " ???     ???  "
  f28[ 63][10] = " ???     ???  "
  f28[ 63][11] = "         ???  "
  f28[ 63][12] = "        ???   "
  f28[ 63][13] = "       ???    "
  f28[ 63][14] = "      ???     "
  f28[ 63][15] = "      ???     "
  f28[ 63][16] = "      ???     "
  f28[ 63][17] = "      ???     "
  f28[ 63][18] = "              "
  f28[ 63][19] = "              "
  f28[ 63][20] = "      ???     "
  f28[ 63][21] = "     ?????    "
  f28[ 63][22] = "      ???     "

  f28[ 64][ 6] = "   @@@@@@@@   "
  f28[ 64][ 7] = "  @@@@@@@@@@  "
  f28[ 64][ 8] = " @@@      @@@ "
  f28[ 64][ 9] = " @@@       @@ "
  f28[ 64][10] = " @@@   @@@@@@ "
  f28[ 64][11] = " @@@  @@@@@@@ "
  f28[ 64][12] = " @@@ @@@  @@@ "
  f28[ 64][13] = " @@@ @@@  @@@ "
  f28[ 64][14] = " @@@ @@@  @@@ "
  f28[ 64][15] = " @@@ @@@  @@@ "
  f28[ 64][16] = " @@@ @@@  @@@ "
  f28[ 64][17] = " @@@ @@@  @@@ "
  f28[ 64][18] = " @@@ @@@  @@@ "
  f28[ 64][19] = " @@@ @@@ @@@@ "
  f28[ 64][20] = " @@@  @@@@@@@ "
  f28[ 64][21] = " @@@   @@@ @@ "
  f28[ 64][22] = " @@@          "
  f28[ 64][23] = " @@@          "
  f28[ 64][24] = "  @@@@@@@@@@@ "
  f28[ 64][25] = "   @@@@@@@@@@ "

  f28[ 65][ 5] = "      A       "
  f28[ 65][ 6] = "     AAA      "
  f28[ 65][ 7] = "    AAAAA     "
  f28[ 65][ 8] = "   AAA AAA    "
  f28[ 65][ 9] = "  AAA   AAA   "
  f28[ 65][10] = " AAA     AAA  "
  f28[ 65][11] = " AAA     AAA  "
  f28[ 65][12] = " AAA     AAA  "
  f28[ 65][13] = " AAA     AAA  "
  f28[ 65][14] = " AAAAAAAAAAA  "
  f28[ 65][15] = " AAAAAAAAAAA  "
  f28[ 65][16] = " AAA     AAA  "
  f28[ 65][17] = " AAA     AAA  "
  f28[ 65][18] = " AAA     AAA  "
  f28[ 65][19] = " AAA     AAA  "
  f28[ 65][20] = " AAA     AAA  "
  f28[ 65][21] = " AAA     AAA  "
  f28[ 65][22] = " AAA     AAA  "

  f28[196][ 1] = "  AAA   AAA   "
  f28[196][ 2] = "  AAA   AAA   "
  f28[196][ 3] = "  AAA   AAA   "
  f28[196][ 4] = "              "
  f28[196][ 5] = "      A       "
  f28[196][ 6] = "     AAA      "
  f28[196][ 7] = "    AAAAA     "
  f28[196][ 8] = "   AAA AAA    "
  f28[196][ 9] = "  AAA   AAA   "
  f28[196][10] = " AAA     AAA  "
  f28[196][11] = " AAA     AAA  "
  f28[196][12] = " AAA     AAA  "
  f28[196][13] = " AAA     AAA  "
  f28[196][14] = " AAAAAAAAAAA  "
  f28[196][15] = " AAAAAAAAAAA  "
  f28[196][16] = " AAA     AAA  "
  f28[196][17] = " AAA     AAA  "
  f28[196][18] = " AAA     AAA  "
  f28[196][19] = " AAA     AAA  "
  f28[196][20] = " AAA     AAA  "
  f28[196][21] = " AAA     AAA  "
  f28[196][22] = " AAA     AAA  "

  f28[ 66][ 5] = " BBBBBBBBB    "
  f28[ 66][ 6] = " BBBBBBBBBB   "
  f28[ 66][ 7] = " BBB     BBB  "
  f28[ 66][ 8] = " BBB     BBB  "
  f28[ 66][ 9] = " BBB     BBB  "
  f28[ 66][10] = " BBB     BBB  "
  f28[ 66][11] = " BBB     BBB  "
  f28[ 66][12] = " BBB    BBB   "
  f28[ 66][13] = " BBBBBBBBB    "
  f28[ 66][14] = " BBBBBBBBBB   "
  f28[ 66][15] = " BBB     BBB  "
  f28[ 66][16] = " BBB     BBB  "
  f28[ 66][17] = " BBB     BBB  "
  f28[ 66][18] = " BBB     BBB  "
  f28[ 66][19] = " BBB     BBB  "
  f28[ 66][20] = " BBB     BBB  "
  f28[ 66][21] = " BBBBBBBBBB   "
  f28[ 66][22] = " BBBBBBBBB    "

  f28[ 67][ 5] = "   CCCCCCC    "
  f28[ 67][ 6] = "  CCCCCCCCC   "
  f28[ 67][ 7] = " CCC     CCC  "
  f28[ 67][ 8] = " CCC     CCC  "
  f28[ 67][ 9] = " CCC     CCC  "
  f28[ 67][10] = " CCC          "
  f28[ 67][11] = " CCC          "
  f28[ 67][12] = " CCC          "
  f28[ 67][13] = " CCC          "
  f28[ 67][14] = " CCC          "
  f28[ 67][15] = " CCC          "
  f28[ 67][16] = " CCC          "
  f28[ 67][17] = " CCC     CCC  "
  f28[ 67][18] = " CCC     CCC  "
  f28[ 67][19] = " CCC     CCC  "
  f28[ 67][20] = " CCC     CCC  "
  f28[ 67][21] = "  CCCCCCCCC   "
  f28[ 67][22] = "   CCCCCCC    "

  f28[ 68][ 5] = " DDDDDDDD     "
  f28[ 68][ 6] = " DDDDDDDDD    "
  f28[ 68][ 7] = " DDD    DDD   "
  f28[ 68][ 8] = " DDD     DDD  "
  f28[ 68][ 9] = " DDD     DDD  "
  f28[ 68][10] = " DDD     DDD  "
  f28[ 68][11] = " DDD     DDD  "
  f28[ 68][12] = " DDD     DDD  "
  f28[ 68][13] = " DDD     DDD  "
  f28[ 68][14] = " DDD     DDD  "
  f28[ 68][15] = " DDD     DDD  "
  f28[ 68][16] = " DDD     DDD  "
  f28[ 68][17] = " DDD     DDD  "
  f28[ 68][18] = " DDD     DDD  "
  f28[ 68][19] = " DDD     DDD  "
  f28[ 68][20] = " DDD    DDD   "
  f28[ 68][21] = " DDDDDDDDD    "
  f28[ 68][22] = " DDDDDDDD     "

  f28[ 69][ 5] = " EEEEEEEEEEE  "
  f28[ 69][ 6] = " EEEEEEEEEEE  "
  f28[ 69][ 7] = " EEE          "
  f28[ 69][ 8] = " EEE          "
  f28[ 69][ 9] = " EEE          "
  f28[ 69][10] = " EEE          "
  f28[ 69][11] = " EEE          "
  f28[ 69][12] = " EEE          "
  f28[ 69][13] = " EEEEEEEE     "
  f28[ 69][14] = " EEEEEEEE     "
  f28[ 69][15] = " EEE          "
  f28[ 69][16] = " EEE          "
  f28[ 69][17] = " EEE          "
  f28[ 69][18] = " EEE          "
  f28[ 69][19] = " EEE          "
  f28[ 69][20] = " EEE          "
  f28[ 69][21] = " EEEEEEEEEEE  "
  f28[ 69][22] = " EEEEEEEEEEE  "

  f28[ 70][ 5] = " FFFFFFFFFFF  "
  f28[ 70][ 6] = " FFFFFFFFFFF  "
  f28[ 70][ 7] = " FFF          "
  f28[ 70][ 8] = " FFF          "
  f28[ 70][ 9] = " FFF          "
  f28[ 70][10] = " FFF          "
  f28[ 70][11] = " FFF          "
  f28[ 70][12] = " FFF          "
  f28[ 70][13] = " FFFFFFFF     "
  f28[ 70][14] = " FFFFFFFF     "
  f28[ 70][15] = " FFF          "
  f28[ 70][16] = " FFF          "
  f28[ 70][17] = " FFF          "
  f28[ 70][18] = " FFF          "
  f28[ 70][19] = " FFF          "
  f28[ 70][20] = " FFF          "
  f28[ 70][21] = " FFF          "
  f28[ 70][22] = " FFF          "

  f28[ 71][ 5] = "   GGGGGGG    "
  f28[ 71][ 6] = "  GGGGGGGGG   "
  f28[ 71][ 7] = " GGG     GGG  "
  f28[ 71][ 8] = " GGG     GGG  "
  f28[ 71][19] = " GGG          "
  f28[ 71][10] = " GGG          "
  f28[ 71][11] = " GGG          "
  f28[ 71][12] = " GGG          "
  f28[ 71][13] = " GGG   GGGGG  "
  f28[ 71][14] = " GGG   GGGGG  "
  f28[ 71][15] = " GGG     GGG  "
  f28[ 71][16] = " GGG     GGG  "
  f28[ 71][17] = " GGG     GGG  "
  f28[ 71][18] = " GGG     GGG  "
  f28[ 71][19] = " GGG     GGG  "
  f28[ 71][20] = " GGG     GGG  "
  f28[ 71][21] = "  GGGGGGGGG   "
  f28[ 71][22] = "   GGGGGGG    "

  f28[ 72][ 5] = " HHH     HHH  "
  f28[ 72][ 6] = " HHH     HHH  "
  f28[ 72][ 7] = " HHH     HHH  "
  f28[ 72][ 8] = " HHH     HHH  "
  f28[ 72][ 9] = " HHH     HHH  "
  f28[ 72][10] = " HHH     HHH  "
  f28[ 72][11] = " HHH     HHH  "
  f28[ 72][12] = " HHH     HHH  "
  f28[ 72][13] = " HHHHHHHHHHH  "
  f28[ 72][14] = " HHHHHHHHHHH  "
  f28[ 72][15] = " HHH     HHH  "
  f28[ 72][16] = " HHH     HHH  "
  f28[ 72][17] = " HHH     HHH  "
  f28[ 72][18] = " HHH     HHH  "
  f28[ 72][19] = " HHH     HHH  "
  f28[ 72][20] = " HHH     HHH  "
  f28[ 72][21] = " HHH     HHH  "
  f28[ 72][22] = " HHH     HHH  "

  f28[ 73][ 5] = "   IIIIIII    "
  f28[ 73][ 6] = "   IIIIIII    "
  f28[ 73][ 7] = "     III      "
  f28[ 73][ 8] = "     III      "
  f28[ 73][ 9] = "     III      "
  f28[ 73][10] = "     III      "
  f28[ 73][11] = "     III      "
  f28[ 73][12] = "     III      "
  f28[ 73][13] = "     III      "
  f28[ 73][14] = "     III      "
  f28[ 73][15] = "     III      "
  f28[ 73][16] = "     III      "
  f28[ 73][17] = "     III      "
  f28[ 73][18] = "     III      "
  f28[ 73][19] = "     III      "
  f28[ 73][20] = "     III      "
  f28[ 73][21] = "   IIIIIII    "
  f28[ 73][22] = "   IIIIIII    "

  f28[ 74][ 5] = "      JJJJJJJ "
  f28[ 74][ 6] = "      JJJJJJJ "
  f28[ 74][ 7] = "        JJJ   "
  f28[ 74][ 8] = "        JJJ   "
  f28[ 74][ 9] = "        JJJ   "
  f28[ 74][10] = "        JJJ   "
  f28[ 74][11] = "        JJJ   "
  f28[ 74][12] = "        JJJ   "
  f28[ 74][13] = "        JJJ   "
  f28[ 74][14] = "        JJJ   "
  f28[ 74][15] = "        JJJ   "
  f28[ 74][16] = "        JJJ   "
  f28[ 74][17] = "        JJJ   "
  f28[ 74][18] = " JJJ    JJJ   "
  f28[ 74][19] = " JJJ    JJJ   "
  f28[ 74][20] = " JJJ    JJJ   "
  f28[ 74][21] = "  JJJJJJJJ    "
  f28[ 74][22] = "   JJJJJJ     "

  f28[ 75][ 5] = " KKK      KK  "
  f28[ 75][ 6] = " KKK     KKK  "
  f28[ 75][ 7] = " KKK    KKK   "
  f28[ 75][ 8] = " KKK   KKK    "
  f28[ 75][ 9] = " KKK  KKK     "
  f28[ 75][10] = " KKK KKK      "
  f28[ 75][11] = " KKKKKK       "
  f28[ 75][12] = " KKKKK        "
  f28[ 75][13] = " KKKK         "
  f28[ 75][14] = " KKKK         "
  f28[ 75][15] = " KKKKK        "
  f28[ 75][16] = " KKKKKK       "
  f28[ 75][17] = " KKK KKK      "
  f28[ 75][18] = " KKK  KKK     "
  f28[ 75][19] = " KKK   KKK    "
  f28[ 75][20] = " KKK    KKK   "
  f28[ 75][21] = " KKK     KKK  "
  f28[ 75][22] = " KKK      KK  "

  f28[ 76][ 5] = " LLL          "
  f28[ 76][ 6] = " LLL          "
  f28[ 76][ 7] = " LLL          "
  f28[ 76][ 8] = " LLL          "
  f28[ 76][ 9] = " LLL          "
  f28[ 76][10] = " LLL          "
  f28[ 76][11] = " LLL          "
  f28[ 76][12] = " LLL          "
  f28[ 76][13] = " LLL          "
  f28[ 76][14] = " LLL          "
  f28[ 76][15] = " LLL          "
  f28[ 76][16] = " LLL          "
  f28[ 76][17] = " LLL          "
  f28[ 76][18] = " LLL          "
  f28[ 76][19] = " LLL          "
  f28[ 76][20] = " LLL          "
  f28[ 76][21] = " LLLLLLLLLLL  "
  f28[ 76][22] = " LLLLLLLLLLL  "

  f28[ 77][ 5] = " MM       MM  "
  f28[ 77][ 6] = " MM       MM  "
  f28[ 77][ 7] = " MMM     MMM  "
  f28[ 77][ 8] = " MMM     MMM  "
  f28[ 77][ 9] = " MMMM   MMMM  "
  f28[ 77][10] = " MMMM   MMMM  "
  f28[ 77][11] = " MMMMM MMMMM  "
  f28[ 77][12] = " MMMMM MMMMM  "
  f28[ 77][13] = " MMMMMMMMMMM  "
  f28[ 77][14] = " MMMMMMMMMMM  "
  f28[ 77][15] = " MMM MMM MMM  "
  f28[ 77][16] = " MMM MMM MMM  "
  f28[ 77][17] = " MMM  M  MMM  "
  f28[ 77][18] = " MMM  M  MMM  "
  f28[ 77][19] = " MMM     MMM  "
  f28[ 77][20] = " MMM     MMM  "
  f28[ 77][21] = " MMM     MMM  "
  f28[ 77][22] = " MMM     MMM  "

  f28[ 78][ 5] = " NNN     NNN  "
  f28[ 78][ 6] = " NNN     NNN  "
  f28[ 78][ 7] = " NNNN    NNN  "
  f28[ 78][ 8] = " NNNN    NNN  "
  f28[ 78][19] = " NNNNN   NNN  "
  f28[ 78][10] = " NNNNN   NNN  "
  f28[ 78][11] = " NNNNNN  NNN  "
  f28[ 78][12] = " NNNNNN  NNN  "
  f28[ 78][13] = " NNN NNN NNN  "
  f28[ 78][14] = " NNN NNN NNN  "
  f28[ 78][15] = " NNN  NNNNNN  "
  f28[ 78][16] = " NNN  NNNNNN  "
  f28[ 78][17] = " NNN   NNNNN  "
  f28[ 78][18] = " NNN   NNNNN  "
  f28[ 78][19] = " NNN    NNNN  "
  f28[ 78][20] = " NNN    NNNN  "
  f28[ 78][21] = " NNN     NNN  "
  f28[ 78][22] = " NNN     NNN  "

  f28[ 79][ 5] = "   OOOOOOO    "
  f28[ 79][ 6] = "  OOOOOOOOO   "
  f28[ 79][ 7] = " OOO     OOO  "
  f28[ 79][ 8] = " OOO     OOO  "
  f28[ 79][ 9] = " OOO     OOO  "
  f28[ 79][10] = " OOO     OOO  "
  f28[ 79][11] = " OOO     OOO  "
  f28[ 79][12] = " OOO     OOO  "
  f28[ 79][13] = " OOO     OOO  "
  f28[ 79][14] = " OOO     OOO  "
  f28[ 79][15] = " OOO     OOO  "
  f28[ 79][16] = " OOO     OOO  "
  f28[ 79][17] = " OOO     OOO  "
  f28[ 79][18] = " OOO     OOO  "
  f28[ 79][19] = " OOO     OOO  "
  f28[ 79][20] = " OOO     OOO  "
  f28[ 79][21] = "  OOOOOOOOO   "
  f28[ 79][22] = "   OOOOOOO    "

  f28[214][ 1] = "  OOO   OOO   "
  f28[214][ 2] = "  OOO   OOO   "
  f28[214][ 3] = "  OOO   OOO   "
  f28[214][ 4] = "              "
  f28[214][ 5] = "   OOOOOOO    "
  f28[214][ 6] = "  OOOOOOOOO   "
  f28[214][ 7] = " OOO     OOO  "
  f28[214][ 8] = " OOO     OOO  "
  f28[214][ 9] = " OOO     OOO  "
  f28[214][10] = " OOO     OOO  "
  f28[214][11] = " OOO     OOO  "
  f28[214][12] = " OOO     OOO  "
  f28[214][13] = " OOO     OOO  "
  f28[214][14] = " OOO     OOO  "
  f28[214][15] = " OOO     OOO  "
  f28[214][16] = " OOO     OOO  "
  f28[214][17] = " OOO     OOO  "
  f28[214][18] = " OOO     OOO  "
  f28[214][19] = " OOO     OOO  "
  f28[214][20] = " OOO     OOO  "
  f28[214][21] = "  OOOOOOOOO   "
  f28[214][22] = "   OOOOOOO    "

  f28[ 80][ 5] = " PPPPPPPPP    "
  f28[ 80][ 6] = " PPPPPPPPPP   "
  f28[ 80][ 7] = " PPP     PPP  "
  f28[ 80][ 8] = " PPP     PPP  "
  f28[ 80][ 9] = " PPP     PPP  "
  f28[ 80][10] = " PPP     PPP  "
  f28[ 80][11] = " PPP     PPP  "
  f28[ 80][12] = " PPP     PPP  "
  f28[ 80][13] = " PPPPPPPPPP   "
  f28[ 80][14] = " PPPPPPPPP    "
  f28[ 80][15] = " PPP          "
  f28[ 80][16] = " PPP          "
  f28[ 80][17] = " PPP          "
  f28[ 80][18] = " PPP          "
  f28[ 80][19] = " PPP          "
  f28[ 80][20] = " PPP          "
  f28[ 80][21] = " PPP          "
  f28[ 80][22] = " PPP          "

  f28[ 81][ 5] = "   QQQQQQQ    "
  f28[ 81][ 6] = "  QQQQQQQQQ   "
  f28[ 81][ 7] = " QQQ     QQQ  "
  f28[ 81][ 8] = " QQQ     QQQ  "
  f28[ 81][ 9] = " QQQ     QQQ  "
  f28[ 81][10] = " QQQ     QQQ  "
  f28[ 81][11] = " QQQ     QQQ  "
  f28[ 81][12] = " QQQ     QQQ  "
  f28[ 81][13] = " QQQ     QQQ  "
  f28[ 81][14] = " QQQ     QQQ  "
  f28[ 81][15] = " QQQ     QQQ  "
  f28[ 81][16] = " QQQ     QQQ  "
  f28[ 81][17] = " QQQ     QQQ  "
  f28[ 81][18] = " QQQ     QQQ  "
  f28[ 81][19] = " QQQ QQQ QQQ  "
  f28[ 81][20] = " QQQ  QQQQQQ  "
  f28[ 81][21] = "  QQQQQQQQQ   "
  f28[ 81][22] = "   QQQQQQQ    "
  f28[ 81][23] = "        QQQ   "
  f28[ 81][24] = "         QQQ  "

  f28[ 82][ 5] = " RRRRRRRRR    "
  f28[ 82][ 6] = " RRRRRRRRRR   "
  f28[ 82][ 7] = " RRR     RRR  "
  f28[ 82][ 8] = " RRR     RRR  "
  f28[ 82][ 9] = " RRR     RRR  "
  f28[ 82][10] = " RRR     RRR  "
  f28[ 82][11] = " RRR     RRR  "
  f28[ 82][12] = " RRR     RRR  "
  f28[ 82][13] = " RRRRRRRRRR   "
  f28[ 82][14] = " RRRRRRRRR    "
  f28[ 82][15] = " RRRR         "
  f28[ 82][16] = " RRRRR        "
  f28[ 82][17] = " RRRRRR       "
  f28[ 82][18] = " RRR RRR      "
  f28[ 82][19] = " RRR  RRR     "
  f28[ 82][20] = " RRR   RRR    "
  f28[ 82][21] = " RRR    RRR   "
  f28[ 82][22] = " RRR     RRR  "

  f28[ 83][ 5] = "   SSSSSSS    "
  f28[ 83][ 6] = "  SSSSSSSSS   "
  f28[ 83][ 7] = " SSS     SSS  "
  f28[ 83][ 8] = " SSS     SSS  "
  f28[ 83][ 9] = " SSS          "
  f28[ 83][10] = " SSS          "
  f28[ 83][11] = " SSS          "
  f28[ 83][12] = " SSS          "
  f28[ 83][13] = "  SSSSSSSS    "
  f28[ 83][14] = "   SSSSSSSS   "
  f28[ 83][15] = "         SSS  "
  f28[ 83][16] = "         SSS  "
  f28[ 83][17] = "         SSS  "
  f28[ 83][18] = "         SSS  "
  f28[ 83][19] = " SSS     SSS  "
  f28[ 83][20] = " SSS     SSS  "
  f28[ 83][21] = "  SSSSSSSSS   "
  f28[ 83][22] = "   SSSSSSS    "

  f28[ 84][ 5] = " TTTTTTTTTTT  "
  f28[ 84][ 6] = " TTTTTTTTTTT  "
  f28[ 84][ 7] = "     TTT      "
  f28[ 84][ 8] = "     TTT      "
  f28[ 84][ 9] = "     TTT      "
  f28[ 84][10] = "     TTT      "
  f28[ 84][11] = "     TTT      "
  f28[ 84][12] = "     TTT      "
  f28[ 84][13] = "     TTT      "
  f28[ 84][14] = "     TTT      "
  f28[ 84][15] = "     TTT      "
  f28[ 84][16] = "     TTT      "
  f28[ 84][17] = "     TTT      "
  f28[ 84][18] = "     TTT      "
  f28[ 84][19] = "     TTT      "
  f28[ 84][20] = "     TTT      "
  f28[ 84][21] = "     TTT      "
  f28[ 84][22] = "     TTT      "

  f28[ 85][ 5] = " UUU     UUU  "
  f28[ 85][ 6] = " UUU     UUU  "
  f28[ 85][ 7] = " UUU     UUU  "
  f28[ 85][ 8] = " UUU     UUU  "
  f28[ 85][ 9] = " UUU     UUU  "
  f28[ 85][10] = " UUU     UUU  "
  f28[ 85][11] = " UUU     UUU  "
  f28[ 85][12] = " UUU     UUU  "
  f28[ 85][13] = " UUU     UUU  "
  f28[ 85][14] = " UUU     UUU  "
  f28[ 85][15] = " UUU     UUU  "
  f28[ 85][16] = " UUU     UUU  "
  f28[ 85][17] = " UUU     UUU  "
  f28[ 85][18] = " UUU     UUU  "
  f28[ 85][19] = " UUU     UUU  "
  f28[ 85][20] = " UUU     UUU  "
  f28[ 85][21] = "  UUUUUUUUU   "
  f28[ 85][22] = "   UUUUUUU    "

  f28[220][ 1] = "  UUU   UUU   "
  f28[220][ 2] = "  UUU   UUU   "
  f28[220][ 3] = "  UUU   UUU   "
  f28[220][ 4] = "              "
  f28[220][ 5] = " UUU     UUU  "
  f28[220][ 6] = " UUU     UUU  "
  f28[220][ 7] = " UUU     UUU  "
  f28[220][ 8] = " UUU     UUU  "
  f28[220][ 9] = " UUU     UUU  "
  f28[220][10] = " UUU     UUU  "
  f28[220][11] = " UUU     UUU  "
  f28[220][12] = " UUU     UUU  "
  f28[220][13] = " UUU     UUU  "
  f28[220][14] = " UUU     UUU  "
  f28[220][15] = " UUU     UUU  "
  f28[220][16] = " UUU     UUU  "
  f28[220][17] = " UUU     UUU  "
  f28[220][18] = " UUU     UUU  "
  f28[220][19] = " UUU     UUU  "
  f28[220][20] = " UUU     UUU  "
  f28[220][21] = "  UUUUUUUUU   "
  f28[220][22] = "   UUUUUUU    "

  f28[ 86][ 5] = " VVV     VVV  "
  f28[ 86][ 6] = " VVV     VVV  "
  f28[ 86][ 7] = " VVV     VVV  "
  f28[ 86][ 8] = " VVV     VVV  "
  f28[ 86][ 9] = "  VVV   VVV   "
  f28[ 86][10] = "  VVV   VVV   "
  f28[ 86][11] = "  VVV   VVV   "
  f28[ 86][12] = "  VVV   VVV   "
  f28[ 86][13] = "   VVV VVV    "
  f28[ 86][14] = "   VVV VVV    "
  f28[ 86][15] = "   VVV VVV    "
  f28[ 86][16] = "   VVV VVV    "
  f28[ 86][17] = "    VVVVV     "
  f28[ 86][18] = "    VVVVV     "
  f28[ 86][19] = "    VVVVV     "
  f28[ 86][20] = "    VVVVV     "
  f28[ 86][21] = "     VVV      "
  f28[ 86][22] = "     VVV      "

  f28[ 87][ 5] = " WWW     WWW  "
  f28[ 87][ 6] = " WWW     WWW  "
  f28[ 87][ 7] = " WWW     WWW  "
  f28[ 87][ 8] = " WWW     WWW  "
  f28[ 87][ 9] = " WWW     WWW  "
  f28[ 87][10] = " WWW     WWW  "
  f28[ 87][11] = " WWW  W  WWW  "
  f28[ 87][12] = " WWW  W  WWW  "
  f28[ 87][13] = " WWW WWW WWW  "
  f28[ 87][14] = " WWW WWW WWW  "
  f28[ 87][15] = " WWWWWWWWWWW  "
  f28[ 87][16] = " WWWWW WWWWW  "
  f28[ 87][17] = " WWWWW WWWWW  "
  f28[ 87][18] = " WWWW   WWWW  "
  f28[ 87][19] = " WWW     WWW  "
  f28[ 87][20] = " WWW     WWW  "
  f28[ 87][21] = " WW       WW  "
  f28[ 87][22] = " WW       WW  "

  f28[ 88][ 5] = " XXX     XXX  "
  f28[ 88][ 6] = " XXX     XXX  "
  f28[ 88][ 7] = "  XXX   XXX   "
  f28[ 88][ 8] = "  XXX   XXX   "
  f28[ 88][19] = "   XXX XXX    "
  f28[ 88][10] = "   XXX XXX    "
  f28[ 88][11] = "    XXXXX     "
  f28[ 88][12] = "    XXXXX     "
  f28[ 88][13] = "     XXX      "
  f28[ 88][14] = "     XXX      "
  f28[ 88][15] = "    XXXXX     "
  f28[ 88][16] = "    XXXXX     "
  f28[ 88][17] = "   XXX XXX    "
  f28[ 88][18] = "   XXX XXX    "
  f28[ 88][19] = "  XXX   XXX   "
  f28[ 88][20] = "  XXX   XXX   "
  f28[ 88][21] = " XXX     XXX  "
  f28[ 88][22] = " XXX     XXX  "

  f28[ 89][ 5] = " YYY     YYY  "
  f28[ 89][ 6] = " YYY     YYY  "
  f28[ 89][ 7] = "  YYY   YYY   "
  f28[ 89][ 8] = "  YYY   YYY   "
  f28[ 89][ 9] = "   YYY YYY    "
  f28[ 89][10] = "   YYY YYY    "
  f28[ 89][11] = "    YYYYY     "
  f28[ 89][12] = "    YYYYY     "
  f28[ 89][13] = "     YYY      "
  f28[ 89][14] = "     YYY      "
  f28[ 89][15] = "     YYY      "
  f28[ 89][16] = "     YYY      "
  f28[ 89][17] = "     YYY      "
  f28[ 89][18] = "     YYY      "
  f28[ 89][19] = "     YYY      "
  f28[ 89][20] = "     YYY      "
  f28[ 89][21] = "     YYY      "
  f28[ 89][22] = "     YYY      "

  f28[ 90][ 5] = " ZZZZZZZZZZZ  "
  f28[ 90][ 6] = " ZZZZZZZZZZZ  "
  f28[ 90][ 7] = "        ZZZ   "
  f28[ 90][ 8] = "        ZZZ   "
  f28[ 90][ 9] = "       ZZZ    "
  f28[ 90][10] = "       ZZZ    "
  f28[ 90][11] = "      ZZZ     "
  f28[ 90][12] = "      ZZZ     "
  f28[ 90][13] = "     ZZZ      "
  f28[ 90][14] = "     ZZZ      "
  f28[ 90][15] = "    ZZZ       "
  f28[ 90][16] = "    ZZZ       "
  f28[ 90][17] = "   ZZZ        "
  f28[ 90][18] = "   ZZZ        "
  f28[ 90][19] = "  ZZZ         "
  f28[ 90][20] = "  ZZZ         "
  f28[ 90][21] = " ZZZZZZZZZZZ  "
  f28[ 90][22] = " ZZZZZZZZZZZ  "

  f28[ 91][ 5] = "   [[[[[[[[   "
  f28[ 91][ 6] = "   [[[[[[[[   "
  f28[ 91][ 7] = "   [[[        "
  f28[ 91][ 8] = "   [[[        "
  f28[ 91][ 9] = "   [[[        "
  f28[ 91][10] = "   [[[        "
  f28[ 91][11] = "   [[[        "
  f28[ 91][12] = "   [[[        "
  f28[ 91][13] = "   [[[        "
  f28[ 91][14] = "   [[[        "
  f28[ 91][15] = "   [[[        "
  f28[ 91][16] = "   [[[        "
  f28[ 91][17] = "   [[[        "
  f28[ 91][18] = "   [[[        "
  f28[ 91][19] = "   [[[        "
  f28[ 91][20] = "   [[[        "
  f28[ 91][21] = "   [[[[[[[[   "
  f28[ 91][22] = "   [[[[[[[[   "

  f28[ 92][ 5] = " ///          "
  f28[ 92][ 6] = " ///          "
  f28[ 92][ 7] = "  ///         "
  f28[ 92][ 8] = "  ///         "
  f28[ 92][ 9] = "   ///        "
  f28[ 92][10] = "   ///        "
  f28[ 92][11] = "    ///       "
  f28[ 92][12] = "    ///       "
  f28[ 92][13] = "     ///      "
  f28[ 92][14] = "     ///      "
  f28[ 92][15] = "      ///     "
  f28[ 92][16] = "      ///     "
  f28[ 92][17] = "       ///    "
  f28[ 92][18] = "       ///    "
  f28[ 92][19] = "        ///   "
  f28[ 92][20] = "        ///   "
  f28[ 92][21] = "         ///  "
  f28[ 92][22] = "         ///  "

  f28[ 93][ 5] = "   ]]]]]]]]   "
  f28[ 93][ 6] = "   ]]]]]]]]   "
  f28[ 93][ 7] = "        ]]]   "
  f28[ 93][ 8] = "        ]]]   "
  f28[ 93][ 9] = "        ]]]   "
  f28[ 93][10] = "        ]]]   "
  f28[ 93][11] = "        ]]]   "
  f28[ 93][12] = "        ]]]   "
  f28[ 93][13] = "        ]]]   "
  f28[ 93][14] = "        ]]]   "
  f28[ 93][15] = "        ]]]   "
  f28[ 93][16] = "        ]]]   "
  f28[ 93][17] = "        ]]]   "
  f28[ 93][18] = "        ]]]   "
  f28[ 93][19] = "        ]]]   "
  f28[ 93][20] = "        ]]]   "
  f28[ 93][21] = "   ]]]]]]]]   "
  f28[ 93][22] = "   ]]]]]]]]   "

  f28[ 94][ 1] = "      ^       "
  f28[ 94][ 2] = "     ^^^      "
  f28[ 94][ 3] = "    ^^^^^     "
  f28[ 94][ 4] = "   ^^^ ^^^    "
  f28[ 94][ 5] = "  ^^^   ^^^   "

  f28[ 95][25] = " ___________  "
  f28[ 95][26] = " ___________  "

  f28[ 96][ 1] = "    ```       "
  f28[ 96][ 2] = "    ```       "
  f28[ 96][ 3] = "    ```       "
  f28[ 96][ 4] = "     ```      "

  f28[ 97][11] = "   aaaaaaaaa  "
  f28[ 97][12] = "  aaaaaaaaaa  "
  f28[ 97][13] = " aaa     aaa  "
  f28[ 97][14] = " aaa     aaa  "
  f28[ 97][15] = " aaa     aaa  "
  f28[ 97][16] = " aaa     aaa  "
  f28[ 97][17] = " aaa     aaa  "
  f28[ 97][18] = " aaa     aaa  "
  f28[ 97][19] = " aaa     aaa  "
  f28[ 97][20] = " aaa     aaa  "
  f28[ 97][21] = "  aaaaaaaaaa  "
  f28[ 97][22] = "   aaaaaaaaaa "

  f28[228][ 5] = "   aaa  aaa   "
  f28[228][ 6] = "   aaa  aaa   "
  f28[228][ 7] = "   aaa  aaa   "
  f28[228][ 8] = "              "
  f28[228][ 9] = "              "
  f28[228][10] = "              "
  f28[228][11] = "   aaaaaaaaa  "
  f28[228][12] = "  aaaaaaaaaa  "
  f28[228][13] = " aaa     aaa  "
  f28[228][14] = " aaa     aaa  "
  f28[228][15] = " aaa     aaa  "
  f28[228][16] = " aaa     aaa  "
  f28[228][17] = " aaa     aaa  "
  f28[228][18] = " aaa     aaa  "
  f28[228][19] = " aaa     aaa  "
  f28[228][20] = " aaa     aaa  "
  f28[228][21] = "  aaaaaaaaaa  "
  f28[228][22] = "   aaaaaaaaaa "

  f28[ 98][ 5] = " bbb          "
  f28[ 98][ 6] = " bbb          "
  f28[ 98][ 7] = " bbb          "
  f28[ 98][ 8] = " bbb          "
  f28[ 98][ 9] = " bbb          "
  f28[ 98][10] = " bbb          "
  f28[ 98][11] = " bbbbbbbbb    "
  f28[ 98][12] = " bbbbbbbbbb   "
  f28[ 98][13] = " bbb     bbb  "
  f28[ 98][14] = " bbb     bbb  "
  f28[ 98][15] = " bbb     bbb  "
  f28[ 98][16] = " bbb     bbb  "
  f28[ 98][17] = " bbb     bbb  "
  f28[ 98][18] = " bbb     bbb  "
  f28[ 98][19] = " bbb     bbb  "
  f28[ 98][20] = " bbb     bbb  "
  f28[ 98][21] = " bbbbbbbbbb   "
  f28[ 98][22] = "bbbbbbbbbb    "

  f28[ 99][11] = "   ccccccc    "
  f28[ 99][12] = "  ccccccccc   "
  f28[ 99][13] = " ccc     ccc  "
  f28[ 99][14] = " ccc     ccc  "
  f28[ 99][15] = " ccc          "
  f28[ 99][16] = " ccc          "
  f28[ 99][17] = " ccc          "
  f28[ 99][18] = " ccc          "
  f28[ 99][19] = " ccc     ccc  "
  f28[ 99][20] = " ccc     ccc  "
  f28[ 99][21] = "  ccccccccc   "
  f28[ 99][22] = "   ccccccc    "

  f28[100][ 5] = "         ddd  "
  f28[100][ 6] = "         ddd  "
  f28[100][ 7] = "         ddd  "
  f28[100][ 8] = "         ddd  "
  f28[100][ 9] = "         ddd  "
  f28[100][10] = "         ddd  "
  f28[100][11] = "   ddddddddd  "
  f28[100][12] = "  dddddddddd  "
  f28[100][13] = " ddd     ddd  "
  f28[100][14] = " ddd     ddd  "
  f28[100][15] = " ddd     ddd  "
  f28[100][16] = " ddd     ddd  "
  f28[100][17] = " ddd     ddd  "
  f28[100][18] = " ddd     ddd  "
  f28[100][19] = " ddd     ddd  "
  f28[100][20] = " ddd     ddd  "
  f28[100][21] = "  dddddddddd  "
  f28[100][22] = "   dddddddddd "

  f28[101][11] = "   eeeeeee    "
  f28[101][12] = "  eeeeeeeee   "
  f28[101][13] = " eee     eee  "
  f28[101][14] = " eee     eee  "
  f28[101][15] = " eee     eee  "
  f28[101][16] = " eeeeeeeeeee  "
  f28[101][17] = " eeeeeeeeeee  "
  f28[101][18] = " eee          "
  f28[101][19] = " eee          "
  f28[101][20] = " eee     eee  "
  f28[101][21] = "  eeeeeeeee   "
  f28[101][22] = "   eeeeeee    "

  f28[102][ 5] = "      ffffff  "
  f28[102][ 6] = "     fffffff  "
  f28[102][ 7] = "    fff       "
  f28[102][ 8] = "    fff       "
  f28[102][ 9] = "    fff       "
  f28[102][10] = "    fff       "
  f28[102][11] = "    fff       "
  f28[102][12] = "    fff       "
  f28[102][13] = "    fff       "
  f28[102][14] = " fffffffff    "
  f28[102][15] = " fffffffff    "
  f28[102][16] = "    fff       "
  f28[102][17] = "    fff       "
  f28[102][18] = "    fff       "
  f28[102][19] = "    fff       "
  f28[102][20] = "    fff       "
  f28[102][21] = "    fff       "
  f28[102][22] = "    fff       "
  f28[102][23] = "    fff       "
  f28[102][24] = "    fff       "
  f28[102][25] = "    fff       "
  f28[102][26] = "    fff       "
  f28[102][27] = "    fff       "

  f28[103][11] = "   gggggggggg "
  f28[103][12] = "  gggggggggg  "
  f28[103][13] = " ggg     ggg  "
  f28[103][14] = " ggg     ggg  "
  f28[103][15] = " ggg     ggg  "
  f28[103][16] = " ggg     ggg  "
  f28[103][17] = " ggg     ggg  "
  f28[103][18] = " ggg     ggg  "
  f28[103][19] = " ggg     ggg  "
  f28[103][20] = " ggg     ggg  "
  f28[103][21] = "  gggggggggg  "
  f28[103][22] = "   ggggggggg  "
  f28[103][23] = "         ggg  "
  f28[103][24] = "         ggg  "
  f28[103][25] = "         ggg  "
  f28[103][26] = "  ggggggggg   "
  f28[103][27] = "  gggggggg    "

  f28[104][ 5] = " hhh          "
  f28[104][ 6] = " hhh          "
  f28[104][ 7] = " hhh          "
  f28[104][ 8] = " hhh          "
  f28[104][19] = " hhh          "
  f28[104][10] = " hhh          "
  f28[104][11] = " hhhhhhhhh    "
  f28[104][12] = " hhhhhhhhhh   "
  f28[104][13] = " hhh     hhh  "
  f28[104][14] = " hhh     hhh  "
  f28[104][15] = " hhh     hhh  "
  f28[104][16] = " hhh     hhh  "
  f28[104][17] = " hhh     hhh  "
  f28[104][18] = " hhh     hhh  "
  f28[104][19] = " hhh     hhh  "
  f28[104][20] = " hhh     hhh  "
  f28[104][21] = " hhh     hhh  "
  f28[104][22] = " hhh     hhh  "

  f28[105][ 6] = "     iii      "
  f28[105][ 7] = "    iiiii     "
  f28[105][ 8] = "     iii      "
  f28[105][ 9] = "              "
  f28[105][10] = "              "
  f28[105][11] = "   iiiii      "
  f28[105][12] = "   iiiii      "
  f28[105][13] = "     iii      "
  f28[105][14] = "     iii      "
  f28[105][15] = "     iii      "
  f28[105][16] = "     iii      "
  f28[105][17] = "     iii      "
  f28[105][18] = "     iii      "
  f28[105][19] = "     iii      "
  f28[105][20] = "     iii      "
  f28[105][21] = "   iiiiiii    "
  f28[105][22] = "   iiiiiii    "

  f28[106][ 6] = "        jjj   "
  f28[106][ 7] = "       jjjjj  "
  f28[106][ 8] = "        jjj   "
  f28[106][ 9] = "              "
  f28[106][10] = "              "
  f28[106][11] = "      jjjjj   "
  f28[106][12] = "      jjjjj   "
  f28[106][13] = "        jjj   "
  f28[106][14] = "        jjj   "
  f28[106][15] = "        jjj   "
  f28[106][16] = "        jjj   "
  f28[106][17] = "        jjj   "
  f28[106][18] = "        jjj   "
  f28[106][19] = "        jjj   "
  f28[106][20] = "        jjj   "
  f28[106][21] = "        jjj   "
  f28[106][22] = "        jjj   "
  f28[106][23] = " jjj    jjj   "
  f28[106][24] = " jjj    jjj   "
  f28[106][25] = " jjj    jjj   "
  f28[106][26] = "  jjjjjjjj    "
  f28[106][27] = "   jjjjjj     "

  f28[107][ 5] = " kkk          "
  f28[107][ 6] = " kkk          "
  f28[107][ 7] = " kkk          "
  f28[107][ 8] = " kkk          "
  f28[107][ 9] = " kkk          "
  f28[107][10] = " kkk          "
  f28[107][11] = " kkk     kkk  "
  f28[107][12] = " kkk    kkk   "
  f28[107][13] = " kkk   kkk    "
  f28[107][14] = " kkk  kkk     "
  f28[107][15] = " kkk kkk      "
  f28[107][16] = " kkkkkk       "
  f28[107][17] = " kkkkkk       "
  f28[107][18] = " kkkkkkk      "
  f28[107][19] = " kkk  kkk     "
  f28[107][20] = " kkk   kkk    "
  f28[107][21] = " kkk    kkk   "
  f28[107][22] = " kkk     kkk  "

  f28[108][ 5] = "   lllll      "
  f28[108][ 6] = "   lllll      "
  f28[108][ 7] = "     lll      "
  f28[108][ 8] = "     lll      "
  f28[108][ 9] = "     lll      "
  f28[108][10] = "     lll      "
  f28[108][11] = "     lll      "
  f28[108][12] = "     lll      "
  f28[108][13] = "     lll      "
  f28[108][14] = "     lll      "
  f28[108][15] = "     lll      "
  f28[108][16] = "     lll      "
  f28[108][17] = "     lll      "
  f28[108][18] = "     lll      "
  f28[108][19] = "     lll      "
  f28[108][20] = "     lll      "
  f28[108][21] = "   lllllll    "
  f28[108][22] = "   lllllll    "

  f28[109][11] = "mmmmmmmmmmm   "
  f28[109][12] = " mmmmmmmmmmm  "
  f28[109][13] = " mmm  mm  mmm "
  f28[109][14] = " mmm  mm  mmm "
  f28[109][15] = " mmm  mm  mmm "
  f28[109][16] = " mmm  mm  mmm "
  f28[109][17] = " mmm  mm  mmm "
  f28[109][18] = " mmm  mm  mmm "
  f28[109][19] = " mmm  mm  mmm "
  f28[109][20] = " mmm  mm  mmm "
  f28[109][21] = " mmm  mm  mmm "
  f28[109][22] = " mmm  mm  mmm "

  f28[110][11] = "nnnnnnnnnn    "
  f28[110][12] = " nnnnnnnnnn   "
  f28[110][13] = " nnn     nnn  "
  f28[110][14] = " nnn     nnn  "
  f28[110][15] = " nnn     nnn  "
  f28[110][16] = " nnn     nnn  "
  f28[110][17] = " nnn     nnn  "
  f28[110][18] = " nnn     nnn  "
  f28[110][19] = " nnn     nnn  "
  f28[110][20] = " nnn     nnn  "
  f28[110][21] = " nnn     nnn  "
  f28[110][22] = " nnn     nnn  "

  f28[111][11] = "   ooooooo    "
  f28[111][12] = "  ooooooooo   "
  f28[111][13] = " ooo     ooo  "
  f28[111][14] = " ooo     ooo  "
  f28[111][15] = " ooo     ooo  "
  f28[111][16] = " ooo     ooo  "
  f28[111][17] = " ooo     ooo  "
  f28[111][18] = " ooo     ooo  "
  f28[111][19] = " ooo     ooo  "
  f28[111][20] = " ooo     ooo  "
  f28[111][21] = "  ooooooooo   "
  f28[111][22] = "   ooooooo    "

  f28[246][ 5] = "  aaa   aaa   "
  f28[246][ 6] = "  aaa   aaa   "
  f28[246][ 7] = "  aaa   aaa   "
  f28[246][ 8] = "              "
  f28[246][ 9] = "              "
  f28[246][10] = "              "
  f28[246][11] = "   ooooooo    "
  f28[246][12] = "  ooooooooo   "
  f28[246][13] = " ooo     ooo  "
  f28[246][14] = " ooo     ooo  "
  f28[246][15] = " ooo     ooo  "
  f28[246][16] = " ooo     ooo  "
  f28[246][17] = " ooo     ooo  "
  f28[246][18] = " ooo     ooo  "
  f28[246][19] = " ooo     ooo  "
  f28[246][20] = " ooo     ooo  "
  f28[246][21] = "  ooooooooo   "
  f28[246][22] = "   ooooooo    "

  f28[112][11] = "pppppppppp    "
  f28[112][12] = " pppppppppp   "
  f28[112][13] = " ppp     ppp  "
  f28[112][14] = " ppp     ppp  "
  f28[112][15] = " ppp     ppp  "
  f28[112][16] = " ppp     ppp  "
  f28[112][17] = " ppp     ppp  "
  f28[112][18] = " ppp     ppp  "
  f28[112][19] = " ppp     ppp  "
  f28[112][20] = " ppp     ppp  "
  f28[112][21] = " pppppppppp   "
  f28[112][22] = " ppppppppp    "
  f28[112][23] = " ppp          "
  f28[112][24] = " ppp          "
  f28[112][25] = " ppp          "
  f28[112][26] = " ppp          "
  f28[112][27] = " ppp          "

  f28[113][11] = "   qqqqqqqqqq "
  f28[113][12] = "  qqqqqqqqqq  "
  f28[113][13] = " qqq     qqq  "
  f28[113][14] = " qqq     qqq  "
  f28[113][15] = " qqq     qqq  "
  f28[113][16] = " qqq     qqq  "
  f28[113][17] = " qqq     qqq  "
  f28[113][18] = " qqq     qqq  "
  f28[113][19] = " qqq     qqq  "
  f28[113][20] = " qqq     qqq  "
  f28[113][21] = "  qqqqqqqqqq  "
  f28[113][22] = "   qqqqqqqqq  "
  f28[113][23] = "         qqq  "
  f28[113][24] = "         qqq  "
  f28[113][25] = "         qqq  "
  f28[113][26] = "         qqq  "
  f28[113][27] = "         qqq  "

  f28[114][11] = "rrrr  rrrrrr  "
  f28[114][12] = " rrr rrrrrrr  "
  f28[114][13] = " rrrrrr       "
  f28[114][14] = " rrrrr        "
  f28[114][15] = " rrrr         "
  f28[114][16] = " rrr          "
  f28[114][17] = " rrr          "
  f28[114][18] = " rrr          "
  f28[114][19] = " rrr          "
  f28[114][20] = " rrr          "
  f28[114][21] = " rrr          "
  f28[114][22] = " rrr          "

  f28[115][11] = "   sssssss    "
  f28[115][12] = "  sssssssss   "
  f28[115][13] = " sss     sss  "
  f28[115][14] = " sss          "
  f28[115][15] = " sss          "
  f28[115][16] = "  ssssssss    "
  f28[115][17] = "   ssssssss   "
  f28[115][18] = "         sss  "
  f28[115][19] = "         sss  "
  f28[115][20] = " sss     sss  "
  f28[115][21] = "  sssssssss   "
  f28[115][22] = "   sssssss    "

  f28[116][ 5] = "    ttt       "
  f28[116][ 6] = "    ttt       "
  f28[116][ 7] = "    ttt       "
  f28[116][ 8] = "    ttt       "
  f28[116][ 9] = "    ttt       "
  f28[116][10] = "    ttt       "
  f28[116][11] = " ttttttttt    "
  f28[116][12] = " ttttttttt    "
  f28[116][13] = "    ttt       "
  f28[116][14] = "    ttt       "
  f28[116][15] = "    ttt       "
  f28[116][16] = "    ttt       "
  f28[116][17] = "    ttt       "
  f28[116][18] = "    ttt       "
  f28[116][19] = "    ttt       "
  f28[116][20] = "    ttt       "
  f28[116][21] = "     ttttttt  "
  f28[116][22] = "      tttttt  "

  f28[117][11] = " uuu     uuu  "
  f28[117][12] = " uuu     uuu  "
  f28[117][13] = " uuu     uuu  "
  f28[117][14] = " uuu     uuu  "
  f28[117][15] = " uuu     uuu  "
  f28[117][16] = " uuu     uuu  "
  f28[117][17] = " uuu     uuu  "
  f28[117][18] = " uuu     uuu  "
  f28[117][19] = " uuu     uuu  "
  f28[117][20] = " uuu     uuu  "
  f28[117][21] = "  uuuuuuuuuu  "
  f28[117][22] = "   uuuuuuuuuu "

  f28[252][ 5] = "  uuu   uuu   "
  f28[252][ 6] = "  uuu   uuu   "
  f28[252][ 7] = "  uuu   uuu   "
  f28[252][ 8] = "              "
  f28[252][ 9] = "              "
  f28[252][10] = "              "
  f28[252][11] = " uuu     uuu  "
  f28[252][12] = " uuu     uuu  "
  f28[252][13] = " uuu     uuu  "
  f28[252][14] = " uuu     uuu  "
  f28[252][15] = " uuu     uuu  "
  f28[252][16] = " uuu     uuu  "
  f28[252][17] = " uuu     uuu  "
  f28[252][18] = " uuu     uuu  "
  f28[252][19] = " uuu     uuu  "
  f28[252][20] = " uuu     uuu  "
  f28[252][21] = "  uuuuuuuuuu  "
  f28[252][22] = "   uuuuuuuuu  "

  f28[118][11] = " vvv     vvv  "
  f28[118][12] = " vvv     vvv  "
  f28[118][13] = " vvv     vvv  "
  f28[118][14] = "  vvv   vvv   "
  f28[118][15] = "  vvv   vvv   "
  f28[118][16] = "  vvv   vvv   "
  f28[118][17] = "   vvv vvv    "
  f28[118][18] = "   vvv vvv    "
  f28[118][19] = "   vvv vvv    "
  f28[118][20] = "    vvvvv     "
  f28[118][21] = "     vvv      "
  f28[118][22] = "     vvv      "

  f28[119][11] = " www      www "
  f28[119][12] = " www      www "
  f28[119][13] = " www      www "
  f28[119][14] = " www      www "
  f28[119][15] = " www      www "
  f28[119][16] = " www  ww  www "
  f28[119][17] = " www  ww  www "
  f28[119][18] = " www  ww  www "
  f28[119][19] = " www  ww  www "
  f28[119][20] = " www  ww  www "
  f28[119][21] = "  wwwwwwwww   "
  f28[119][22] = "   wwwwwww    "

  f28[120][11] = " xxx     xxx  "
  f28[120][12] = " xxx     xxx  "
  f28[120][13] = "  xxx   xxx   "
  f28[120][14] = "   xxx xxx    "
  f28[120][15] = "    xxxxx     "
  f28[120][16] = "     xxx      "
  f28[120][17] = "     xxx      "
  f28[120][18] = "    xxxxx     "
  f28[120][19] = "   xxx xxx    "
  f28[120][20] = "  xxx   xxx   "
  f28[120][21] = " xxx     xxx  "
  f28[120][22] = " xxx     xxx  "

  f28[121][11] = " yyy     yyy  "
  f28[121][12] = " yyy     yyy  "
  f28[121][13] = " yyy     yyy  "
  f28[121][14] = " yyy     yyy  "
  f28[121][15] = " yyy     yyy  "
  f28[121][16] = " yyy     yyy  "
  f28[121][17] = " yyy     yyy  "
  f28[121][18] = " yyy     yyy  "
  f28[121][19] = " yyy     yyy  "
  f28[121][20] = " yyy     yyy  "
  f28[121][21] = "  yyyyyyyyyy  "
  f28[121][22] = "   yyyyyyyyy  "
  f28[121][23] = "         yyy  "
  f28[121][24] = "         yyy  "
  f28[121][25] = "         yyy  "
  f28[121][26] = "  yyyyyyyyy   "
  f28[121][27] = "  yyyyyyyy    "

  f28[122][11] = " zzzzzzzzzzz  "
  f28[122][12] = " zzzzzzzzzzz  "
  f28[122][13] = "         zzz  "
  f28[122][14] = "        zzz   "
  f28[122][15] = "       zzz    "
  f28[122][16] = "      zzz     "
  f28[122][17] = "     zzz      "
  f28[122][18] = "   zzz        "
  f28[122][19] = "  zzz         "
  f28[122][20] = " zzz          "
  f28[122][21] = " zzzzzzzzzzz  "
  f28[122][22] = " zzzzzzzzzzz  "

  f28[123][ 5] = "      {{{{{   "
  f28[123][ 6] = "     {{{{{{   "
  f28[123][ 7] = "    {{{       "
  f28[123][ 8] = "    {{{       "
  f28[123][ 9] = "    {{{       "
  f28[123][10] = "    {{{       "
  f28[123][11] = "    {{{       "
  f28[123][12] = "    {{{       "
  f28[123][13] = "  {{{{        "
  f28[123][14] = "  {{{{        "
  f28[123][15] = "    {{{       "
  f28[123][16] = "    {{{       "
  f28[123][17] = "    {{{       "
  f28[123][18] = "    {{{       "
  f28[123][19] = "    {{{       "
  f28[123][20] = "    {{{       "
  f28[123][21] = "     {{{{{{   "
  f28[123][22] = "      {{{{{   "

  f28[124][ 5] = "     |||      "
  f28[124][ 6] = "     |||      "
  f28[124][ 7] = "     |||      "
  f28[124][ 8] = "     |||      "
  f28[124][ 9] = "     |||      "
  f28[124][10] = "     |||      "
  f28[124][11] = "     |||      "
  f28[124][12] = "     |||      "
  f28[124][13] = "     |||      "
  f28[124][14] = "     |||      "
  f28[124][15] = "     |||      "
  f28[124][16] = "     |||      "
  f28[124][17] = "     |||      "
  f28[124][18] = "     |||      "
  f28[124][19] = "     |||      "
  f28[124][20] = "     |||      "
  f28[124][21] = "     |||      "
  f28[124][22] = "     |||      "

  f28[125][ 5] = "  }}}}}       "
  f28[125][ 6] = "  }}}}}}      "
  f28[125][ 7] = "      }}}     "
  f28[125][ 8] = "      }}}     "
  f28[125][ 9] = "      }}}     "
  f28[125][10] = "      }}}     "
  f28[125][11] = "      }}}     "
  f28[125][12] = "      }}}     "
  f28[125][13] = "       }}}}   "
  f28[125][14] = "       }}}}   "
  f28[125][15] = "      }}}     "
  f28[125][16] = "      }}}     "
  f28[125][17] = "      }}}     "
  f28[125][18] = "      }}}     "
  f28[125][19] = "      }}}     "
  f28[125][20] = "      }}}     "
  f28[125][21] = "  }}}}}}      "
  f28[125][22] = "  }}}}}       "

  f28[126][ 3] = "   ~~~   ~~~  "
  f28[126][ 4] = "  ~~~~   ~~~  "
  f28[126][ 5] = " ~~~ ~~~ ~~~  "
  f28[126][ 6] = " ~~~ ~~~ ~~~  "
  f28[126][ 7] = " ~~~  ~~~~~   "
  f28[126][ 8] = " ~~~   ~~~    "

  f28[164][ 5] = "      eeee    "
  f28[164][ 6] = "     eeeeee   "
  f28[164][ 7] = "    eee  eee  "
  f28[164][ 8] = "   eee    eee "
  f28[164][ 9] = "  eee         "
  f28[164][10] = "  eee         "
  f28[164][11] = "  eee         "
  f28[164][12] = "  eee         "
  f28[164][13] = " eeeeeeeeee   "
  f28[164][14] = " eeeeeeeeee   "
  f28[164][15] = "  eee         "
  f28[164][16] = "  eee         "
  f28[164][17] = "  eee         "
  f28[164][18] = "  eee         "
  f28[164][19] = "   eee    eee "
  f28[164][20] = "    eee  eee  "
  f28[164][21] = "     eeeeee   "
  f28[164][22] = "      eeee    "

  f28[167][ 5] = "   ppppppp    "
  f28[167][ 6] = "  ppppppppp   "
  f28[167][ 7] = " ppp     ppp  "
  f28[167][ 8] = " ppp          "
  f28[167][ 9] = " pppp         "
  f28[167][10] = "  ppppppp     "
  f28[167][11] = "  pppppppp    "
  f28[167][12] = " ppp    ppp   "
  f28[167][13] = " ppp     ppp  "
  f28[167][14] = " ppp     ppp  "
  f28[167][15] = "  ppp    ppp  "
  f28[167][16] = "   pppppppp   "
  f28[167][17] = "    ppppppp   "
  f28[167][18] = "        pppp  "
  f28[167][19] = "         ppp  "
  f28[167][20] = " ppp     ppp  "
  f28[167][21] = "  ppppppppp   "
  f28[167][22] = "   ppppppp    "

  f28[176][ 3] = "   ooooooo    "
  f28[176][ 4] = "  ooooooooo   "
  f28[176][ 5] = "  ooo   ooo   "
  f28[176][ 6] = "  ooo   ooo   "
  f28[176][ 7] = "  ooo   ooo   "
  f28[176][ 8] = "  ooo   ooo   "
  f28[176][ 9] = "  ooooooooo   "
  f28[176][10] = "   ooooooo    "

  f28[178][ 3] = "    222222    "
  f28[178][ 4] = "   22222222   "
  f28[178][ 5] = "   222  222   "
  f28[178][ 6] = "   222  222   "
  f28[178][ 7] = "       222    "
  f28[178][ 8] = "      222     "
  f28[178][ 9] = "     222      "
  f28[178][10] = "    2222222   "
  f28[178][11] = "   22222222   "

  f28[179][ 3] = "   3333333    "
  f28[179][ 4] = "  333333333   "
  f28[179][ 5] = "  333   333   "
  f28[179][ 6] = "     33333    "
  f28[179][ 7] = "     33333    "
  f28[179][ 8] = "        333   "
  f28[179][ 9] = "  333   333   "
  f28[179][10] = "  333333333   "
  f28[179][11] = "   3333333    "

  f28[181][13] = " mmm     mmm  "
  f28[181][14] = " mmm     mmm  "
  f28[181][15] = " mmm     mmm  "
  f28[181][16] = " mmm     mmm  "
  f28[181][17] = " mmm     mmm  "
  f28[181][18] = " mmm     mmm  "
  f28[181][19] = " mmmm   mmmm  "
  f28[181][20] = " mmmmm mmmmm  "
  f28[181][21] = " mmmmmmmmmmm  "
  f28[181][22] = " mmmm    mmm  "
  f28[181][23] = " mmm          "
  f28[181][25] = " mmm          "
  f28[181][25] = " mmm          "
  f28[181][26] = " mmm          "
  f28[181][27] = " mmm          "

  f28[223][ 5] = "   ssssss     "
  f28[223][ 6] = "  sssssssss   "
  f28[223][ 7] = " sss    sss   "
  f28[223][ 8] = " sss    sss   "
  f28[223][ 9] = " sss    sss   "
  f28[223][10] = " sss    sss   "
  f28[223][11] = " sss   sss    "
  f28[223][12] = " sss ssss     "
  f28[223][13] = " sss sssss    "
  f28[223][14] = " sss    sss   "
  f28[223][15] = " sss     sss  "
  f28[223][16] = " sss     sss  "
  f28[223][17] = " sss     sss  "
  f28[223][18] = " sss     sss  "
  f28[223][19] = " sss     sss  "
  f28[223][20] = " sss ss  sss  "
  f28[223][21] = " sss ssssss   "
  f28[223][22] = " sss  ssss    "
  f28[223][23] = " sss          "
  f28[223][24] = " sss          "
  f28[223][25] = " sss          "
  f28[223][26] = " sss          "
  f28[223][27] = " sss          "

  for b := 0; b <= 255; b++ {
    for l := 0; l < 32; l++ {
      f32[b][l] = "                "
    }
  }

  f32[ 33][ 6] = "      !!!       "
  f32[ 33][ 7] = "      !!!       "
  f32[ 33][ 8] = "      !!!       "
  f32[ 33][ 9] = "      !!!       "
  f32[ 33][10] = "      !!!       "
  f32[ 33][11] = "      !!!       "
  f32[ 33][12] = "      !!!       "
  f32[ 33][13] = "      !!!       "
  f32[ 33][14] = "      !!!       "
  f32[ 33][15] = "      !!!       "
  f32[ 33][16] = "      !!!       "
  f32[ 33][17] = "      !!!       "
  f32[ 33][18] = "      !!!       "
  f32[ 33][19] = "      !!!       "
  f32[ 33][20] = "                "
  f32[ 33][21] = "                "
  f32[ 33][22] = "      !!!       "
  f32[ 33][23] = "      !!!       "
  f32[ 33][24] = "      !!!       "
  f32[ 33][25] = "      !!!       "

  f32[ 34][ 3] = "   ***   ***    "
  f32[ 34][ 4] = "   ***   ***    "
  f32[ 34][ 5] = "   ***   ***    "
  f32[ 34][ 6] = "   ***   ***    "
  f32[ 34][ 7] = "   ***   ***    "
  f32[ 34][ 8] = "   ***   ***    "

  f32[ 35][ 6] = "   ###   ###    "
  f32[ 35][ 7] = "   ###   ###    "
  f32[ 35][ 8] = "   ###   ###    "
  f32[ 35][ 9] = "   ###   ###    "
  f32[ 35][10] = "   ###   ###    "
  f32[ 35][11] = " #############  "
  f32[ 35][12] = " #############  "
  f32[ 35][13] = "   ###   ###    "
  f32[ 35][14] = "   ###   ###    "
  f32[ 35][15] = "   ###   ###    "
  f32[ 35][16] = "   ###   ###    "
  f32[ 35][17] = "   ###   ###    "
  f32[ 35][18] = "   ###   ###    "
  f32[ 35][19] = " #############  "
  f32[ 35][20] = " #############  "
  f32[ 35][21] = "   ###   ###    "
  f32[ 35][22] = "   ###   ###    "
  f32[ 35][23] = "   ###   ###    "
  f32[ 35][24] = "   ###   ###    "
  f32[ 35][25] = "   ###   ###    "

  f32[ 36][ 5] = "      $$$       "
  f32[ 36][ 6] = "      $$$       "
  f32[ 36][ 7] = "   $$$$$$$$$    "
  f32[ 36][ 8] = "  $$$$$$$$$$$   "
  f32[ 36][ 9] = " $$$  $$$  $$$  "
  f32[ 36][10] = " $$$  $$$  $$$  "
  f32[ 36][11] = " $$$  $$$       "
  f32[ 36][12] = " $$$  $$$       "
  f32[ 36][13] = " $$$  $$$       "
  f32[ 36][14] = " $$$  $$$       "
  f32[ 36][15] = "  $$$$$$$$$$    "
  f32[ 36][16] = "   $$$$$$$$$$   "
  f32[ 36][17] = "      $$$  $$$  "
  f32[ 36][18] = "      $$$  $$$  "
  f32[ 36][19] = "      $$$  $$$  "
  f32[ 36][20] = "      $$$  $$$  "
  f32[ 36][21] = " $$$  $$$  $$$  "
  f32[ 36][22] = " $$$  $$$  $$$  "
  f32[ 36][23] = "  $$$$$$$$$$$   "
  f32[ 36][24] = "   $$$$$$$$$    "
  f32[ 36][25] = "      $$$       "
  f32[ 36][26] = "      $$$       "

  f32[ 37][ 6] = "   %%%%%   %%%  "
  f32[ 37][ 7] = "  %%%%%%%  %%%  "
  f32[ 37][ 8] = "  %%% %%% %%%   "
  f32[ 37][ 9] = "  %%% %%% %%%   "
  f32[ 37][10] = "  %%%%%%%%%%    "
  f32[ 37][11] = "   %%%%% %%%    "
  f32[ 37][12] = "        %%%     "
  f32[ 37][13] = "        %%%     "
  f32[ 37][14] = "       %%%      "
  f32[ 37][15] = "       %%%      "
  f32[ 37][16] = "      %%%       "
  f32[ 37][17] = "      %%%       "
  f32[ 37][18] = "     %%%        "
  f32[ 37][19] = "     %%%        "
  f32[ 37][20] = "    %%% %%%%%   "
  f32[ 37][21] = "    %%%%%%%%%%  "
  f32[ 37][22] = "   %%% %%% %%%  "
  f32[ 37][23] = "   %%% %%% %%%  "
  f32[ 37][24] = "  %%%  %%%%%%%  "
  f32[ 37][25] = "  %%%   %%%%%   "

  f32[ 38][ 6] = "    &&&&&&      "
  f32[ 38][ 7] = "   &&&&&&&&     "
  f32[ 38][ 8] = "  &&&    &&&    "
  f32[ 38][ 9] = "  &&&    &&&    "
  f32[ 38][10] = "  &&&    &&&    "
  f32[ 38][11] = "  &&&    &&&    "
  f32[ 38][12] = "  &&&    &&&    "
  f32[ 38][13] = "   &&&  &&&     "
  f32[ 38][14] = "    &&&&&&      "
  f32[ 38][15] = "    &&&&&       "
  f32[ 38][16] = "   &&&&&&&  &&& "
  f32[ 38][17] = "  &&&   &&& &&& "
  f32[ 38][18] = " &&&     &&&&&  "
  f32[ 38][19] = " &&&      &&&   "
  f32[ 38][20] = " &&&      &&&   "
  f32[ 38][21] = " &&&      &&&   "
  f32[ 38][22] = " &&&      &&&   "
  f32[ 38][23] = " &&&     &&&&&  "
  f32[ 38][24] = "  &&&&&&&&& &&& "
  f32[ 38][25] = "   &&&&&&&  &&& "

  f32[ 39][ 3] = "       '''      "
  f32[ 39][ 4] = "       '''      "
  f32[ 39][ 5] = "       '''      "
  f32[ 39][ 6] = "       '''      "
  f32[ 39][ 7] = "      '''       "
  f32[ 39][ 8] = "      '''       "

  f32[ 40][ 6] = "        (((     "
  f32[ 40][ 7] = "       (((      "
  f32[ 40][ 8] = "      (((       "
  f32[ 40][ 9] = "     (((        "
  f32[ 40][10] = "     (((        "
  f32[ 40][11] = "    (((         "
  f32[ 40][12] = "    (((         "
  f32[ 40][13] = "    (((         "
  f32[ 40][14] = "    (((         "
  f32[ 40][15] = "    (((         "
  f32[ 40][16] = "    (((         "
  f32[ 40][17] = "    (((         "
  f32[ 40][18] = "    (((         "
  f32[ 40][19] = "    (((         "
  f32[ 40][20] = "    (((         "
  f32[ 40][21] = "     (((        "
  f32[ 40][22] = "     (((        "
  f32[ 40][23] = "      (((       "
  f32[ 40][24] = "       (((      "
  f32[ 40][25] = "        (((     "

  f32[ 41][ 6] = "    )))         "
  f32[ 41][ 7] = "     )))        "
  f32[ 41][ 8] = "      )))       "
  f32[ 41][ 9] = "       )))      "
  f32[ 41][10] = "       )))      "
  f32[ 41][11] = "        )))     "
  f32[ 41][12] = "        )))     "
  f32[ 41][13] = "        )))     "
  f32[ 41][14] = "        )))     "
  f32[ 41][15] = "        )))     "
  f32[ 41][16] = "        )))     "
  f32[ 41][17] = "        )))     "
  f32[ 41][18] = "        )))     "
  f32[ 41][19] = "        )))     "
  f32[ 41][20] = "        )))     "
  f32[ 41][21] = "       )))      "
  f32[ 41][22] = "       )))      "
  f32[ 41][23] = "      )))       "
  f32[ 41][24] = "     )))        "
  f32[ 41][25] = "    )))         "

  f32[ 42][10] = "  ***     ***   "
  f32[ 42][11] = "   ***   ***    "
  f32[ 42][12] = "    *** ***     "
  f32[ 42][13] = "     *****      "
  f32[ 42][14] = "      ***       "
  f32[ 42][15] = " *************  "
  f32[ 42][16] = " *************  "
  f32[ 42][17] = "      ***       "
  f32[ 42][18] = "     *****      "
  f32[ 42][19] = "    *** ***     "
  f32[ 42][20] = "   ***   ***    "
  f32[ 42][21] = "  ***     ***   "

  f32[ 43][10] = "      +++       "
  f32[ 43][11] = "      +++       "
  f32[ 43][12] = "      +++       "
  f32[ 43][13] = "      +++       "
  f32[ 43][14] = "      +++       "
  f32[ 43][15] = " +++++++++++++  "
  f32[ 43][16] = " +++++++++++++  "
  f32[ 43][17] = "      +++       "
  f32[ 43][18] = "      +++       "
  f32[ 43][19] = "      +++       "
  f32[ 43][20] = "      +++       "
  f32[ 43][21] = "      +++       "

  f32[ 44][22] = "      ,,,       "
  f32[ 44][23] = "      ,,,       "
  f32[ 44][24] = "      ,,,       "
  f32[ 44][25] = "      ,,,       "
  f32[ 44][26] = "     ,,,        "
  f32[ 44][27] = "    ,,,         "

  f32[ 45][15] = " -------------  "
  f32[ 45][16] = " -------------  "

  f32[ 46][22] = "      ...       "
  f32[ 46][23] = "      ...       "
  f32[ 46][24] = "      ...       "
  f32[ 46][25] = "      ...       "

  f32[ 47][ 7] = "           ///  "
  f32[ 47][ 8] = "           ///  "
  f32[ 47][ 9] = "          ///   "
  f32[ 47][10] = "          ///   "
  f32[ 47][11] = "         ///    "
  f32[ 47][12] = "         ///    "
  f32[ 47][13] = "        ///     "
  f32[ 47][14] = "        ///     "
  f32[ 47][15] = "       ///      "
  f32[ 47][16] = "       ///      "
  f32[ 47][17] = "      ///       "
  f32[ 47][18] = "      ///       "
  f32[ 47][19] = "     ///        "
  f32[ 47][20] = "     ///        "
  f32[ 47][21] = "    ///         "
  f32[ 47][22] = "    ///         "
  f32[ 47][23] = "   ///          "
  f32[ 47][24] = "   ///          "
  f32[ 47][25] = "  ///           "
  f32[ 47][26] = "  ///           "

  f32[ 48][ 6] = "   000000000    "
  f32[ 48][ 7] = "  00000000000   "
  f32[ 48][ 8] = " 000       000  "
  f32[ 48][ 9] = " 000       000  "
  f32[ 48][10] = " 000       000  "
  f32[ 48][11] = " 000      0000  "
  f32[ 48][12] = " 000     00000  "
  f32[ 48][13] = " 000    000000  "
  f32[ 48][14] = " 000   000 000  "
  f32[ 48][15] = " 000  000  000  "
  f32[ 48][16] = " 000 000   000  "
  f32[ 48][17] = " 000000    000  "
  f32[ 48][18] = " 00000     000  "
  f32[ 48][19] = " 0000      000  "
  f32[ 48][20] = " 000       000  "
  f32[ 48][21] = " 000       000  "
  f32[ 48][22] = " 000       000  "
  f32[ 48][23] = " 000       000  "
  f32[ 48][24] = "  00000000000   "
  f32[ 48][25] = "   000000000    "

  f32[ 49][ 6] = "      111       "
  f32[ 49][ 7] = "     1111       "
  f32[ 49][ 8] = "    11111       "
  f32[ 49][ 9] = "   111111       "
  f32[ 49][10] = "   111111       "
  f32[ 49][11] = "      111       "
  f32[ 49][12] = "      111       "
  f32[ 49][13] = "      111       "
  f32[ 49][14] = "      111       "
  f32[ 49][15] = "      111       "
  f32[ 49][16] = "      111       "
  f32[ 49][17] = "      111       "
  f32[ 49][18] = "      111       "
  f32[ 49][19] = "      111       "
  f32[ 49][20] = "      111       "
  f32[ 49][21] = "      111       "
  f32[ 49][22] = "      111       "
  f32[ 49][23] = "      111       "
  f32[ 49][24] = "   111111111    "
  f32[ 49][25] = "   111111111    "

  f32[ 50][ 6] = "   222222222    "
  f32[ 50][ 7] = "  22222222222   "
  f32[ 50][ 8] = " 222       222  "
  f32[ 50][ 9] = " 222       222  "
  f32[ 50][10] = " 222       222  "
  f32[ 50][11] = " 222       222  "
  f32[ 50][12] = " 222       222  "
  f32[ 50][13] = "           222  "
  f32[ 50][14] = "          222   "
  f32[ 50][15] = "         222    "
  f32[ 50][16] = "        222     "
  f32[ 50][17] = "       222      "
  f32[ 50][18] = "      222       "
  f32[ 50][19] = "     222        "
  f32[ 50][20] = "    222         "
  f32[ 50][21] = "   222          "
  f32[ 50][22] = "  222           "
  f32[ 50][23] = " 222            "
  f32[ 50][24] = " 2222222222222  "
  f32[ 50][25] = " 2222222222222  "

  f32[ 51][ 6] = "   333333333    "
  f32[ 51][ 7] = "  33333333333   "
  f32[ 51][ 8] = " 333       333  "
  f32[ 51][ 9] = " 333       333  "
  f32[ 51][10] = "           333  "
  f32[ 51][11] = "           333  "
  f32[ 51][12] = "           333  "
  f32[ 51][13] = "           333  "
  f32[ 51][14] = "           333  "
  f32[ 51][15] = "    333333333   "
  f32[ 51][16] = "    333333333   "
  f32[ 51][17] = "           333  "
  f32[ 51][18] = "           333  "
  f32[ 51][19] = "           333  "
  f32[ 51][20] = "           333  "
  f32[ 51][21] = "           333  "
  f32[ 51][22] = " 333       333  "
  f32[ 51][23] = " 333       333  "
  f32[ 51][24] = "  33333333333   "
  f32[ 51][25] = "   333333333    "

  f32[ 52][ 6] = "           444  "
  f32[ 52][ 7] = "          4444  "
  f32[ 52][ 8] = "         44444  "
  f32[ 52][ 9] = "        444444  "
  f32[ 52][10] = "       444 444  "
  f32[ 52][11] = "      444  444  "
  f32[ 52][12] = "     444   444  "
  f32[ 52][13] = "    444    444  "
  f32[ 52][14] = "   444     444  "
  f32[ 52][15] = "  444      444  "
  f32[ 52][16] = " 444       444  "
  f32[ 52][17] = " 444       444  "
  f32[ 52][18] = " 444       444  "
  f32[ 52][19] = " 4444444444444  "
  f32[ 52][20] = " 4444444444444  "
  f32[ 52][21] = "           444  "
  f32[ 52][22] = "           444  "
  f32[ 52][23] = "           444  "
  f32[ 52][24] = "           444  "
  f32[ 52][25] = "           444  "

  f32[ 53][ 6] = " 5555555555555  "
  f32[ 53][ 7] = " 5555555555555  "
  f32[ 53][ 8] = " 555            "
  f32[ 53][ 9] = " 555            "
  f32[ 53][10] = " 555            "
  f32[ 53][11] = " 555            "
  f32[ 53][12] = " 555            "
  f32[ 53][13] = " 555            "
  f32[ 53][14] = " 55555555555    "
  f32[ 53][15] = " 555555555555   "
  f32[ 53][16] = "           555  "
  f32[ 53][17] = "           555  "
  f32[ 53][18] = "           555  "
  f32[ 53][19] = "           555  "
  f32[ 53][20] = "           555  "
  f32[ 53][21] = "           555  "
  f32[ 53][22] = " 555       555  "
  f32[ 53][23] = " 555       555  "
  f32[ 53][24] = "  55555555555   "
  f32[ 53][25] = "   555555555    "

  f32[ 54][ 6] = "   6666666666   "
  f32[ 54][ 7] = "  66666666666   "
  f32[ 54][ 8] = " 666            "
  f32[ 54][ 9] = " 666            "
  f32[ 54][10] = " 666            "
  f32[ 54][11] = " 666            "
  f32[ 54][12] = " 666            "
  f32[ 54][13] = " 666            "
  f32[ 54][14] = " 66666666666    "
  f32[ 54][15] = " 666666666666   "
  f32[ 54][16] = " 666       666  "
  f32[ 54][17] = " 666       666  "
  f32[ 54][18] = " 666       666  "
  f32[ 54][19] = " 666       666  "
  f32[ 54][20] = " 666       666  "
  f32[ 54][21] = " 666       666  "
  f32[ 54][22] = " 666       666  "
  f32[ 54][23] = " 666       666  "
  f32[ 54][24] = "  66666666666   "
  f32[ 54][25] = "   666666666    "

  f32[ 55][ 6] = " 7777777777777  "
  f32[ 55][ 7] = " 7777777777777  "
  f32[ 55][ 8] = " 777       777  "
  f32[ 55][ 9] = " 777       777  "
  f32[ 55][10] = " 777       777  "
  f32[ 55][11] = " 777      777   "
  f32[ 55][12] = "          777   "
  f32[ 55][13] = "         777    "
  f32[ 55][14] = "         777    "
  f32[ 55][15] = "         777    "
  f32[ 55][16] = "        777     "
  f32[ 55][17] = "        777     "
  f32[ 55][18] = "       777      "
  f32[ 55][19] = "       777      "
  f32[ 55][20] = "      777       "
  f32[ 55][21] = "      777       "
  f32[ 55][22] = "      777       "
  f32[ 55][23] = "      777       "
  f32[ 55][24] = "      777       "
  f32[ 55][25] = "      777       "

  f32[ 56][ 6] = "   888888888    "
  f32[ 56][ 7] = "  88888888888   "
  f32[ 56][ 8] = " 888       888  "
  f32[ 56][ 9] = " 888       888  "
  f32[ 56][10] = " 888       888  "
  f32[ 56][11] = " 888       888  "
  f32[ 56][12] = " 888       888  "
  f32[ 56][13] = " 888       888  "
  f32[ 56][14] = " 888       888  "
  f32[ 56][15] = "  88888888888   "
  f32[ 56][16] = "  88888888888   "
  f32[ 56][17] = " 888       888  "
  f32[ 56][18] = " 888       888  "
  f32[ 56][19] = " 888       888  "
  f32[ 56][20] = " 888       888  "
  f32[ 56][21] = " 888       888  "
  f32[ 56][22] = " 888       888  "
  f32[ 56][23] = " 888       888  "
  f32[ 56][24] = "  88888888888   "
  f32[ 56][25] = "   888888888    "

  f32[ 57][ 6] = "   999999999    "
  f32[ 57][ 7] = "  99999999999   "
  f32[ 57][ 8] = " 999       999  "
  f32[ 57][ 9] = " 999       999  "
  f32[ 57][10] = " 999       999  "
  f32[ 57][11] = " 999       999  "
  f32[ 57][12] = " 999       999  "
  f32[ 57][13] = " 999       999  "
  f32[ 57][14] = " 999       999  "
  f32[ 57][15] = " 999       999  "
  f32[ 57][16] = "  999999999999  "
  f32[ 57][17] = "   99999999999  "
  f32[ 57][18] = "           999  "
  f32[ 57][19] = "           999  "
  f32[ 57][20] = "           999  "
  f32[ 57][21] = "           999  "
  f32[ 57][22] = "           999  "
  f32[ 57][23] = "           999  "
  f32[ 57][24] = "  99999999999   "
  f32[ 57][25] = "  9999999999    "

  f32[ 58][12] = "      :::       "
  f32[ 58][13] = "      :::       "
  f32[ 58][14] = "      :::       "
  f32[ 58][15] = "      :::       "
  f32[ 58][16] = "                "
  f32[ 58][17] = "                "
  f32[ 58][18] = "                "
  f32[ 58][19] = "                "
  f32[ 58][20] = "                "
  f32[ 58][21] = "                "
  f32[ 58][22] = "      :::       "
  f32[ 58][23] = "      :::       "
  f32[ 58][24] = "      :::       "
  f32[ 58][25] = "      :::       "

  f32[ 59][12] = "      ,;;       "
  f32[ 59][13] = "      ,;;       "
  f32[ 59][14] = "      ,;;       "
  f32[ 59][15] = "      ,;;       "
  f32[ 59][16] = "                "
  f32[ 59][17] = "                "
  f32[ 59][18] = "                "
  f32[ 59][19] = "                "
  f32[ 59][20] = "                "
  f32[ 59][21] = "                "
  f32[ 59][22] = "      ,;;       "
  f32[ 59][23] = "      ,;;       "
  f32[ 59][24] = "      ,;;       "
  f32[ 59][25] = "      ,;;       "
  f32[ 59][26] = "     ,;;        "
  f32[ 59][27] = "    ,;;         "

  f32[ 60][ 6] = "           <<<  "
  f32[ 60][ 7] = "          <<<   "
  f32[ 60][ 8] = "         <<<    "
  f32[ 60][ 9] = "        <<<     "
  f32[ 60][10] = "       <<<      "
  f32[ 60][11] = "      <<<       "
  f32[ 60][12] = "     <<<        "
  f32[ 60][13] = "    <<<         "
  f32[ 60][14] = "   <<<          "
  f32[ 60][15] = "  <<<           "
  f32[ 60][16] = "  <<<           "
  f32[ 60][17] = "   <<<          "
  f32[ 60][18] = "    <<<         "
  f32[ 60][19] = "     <<<        "
  f32[ 60][20] = "      <<<       "
  f32[ 60][21] = "       <<<      "
  f32[ 60][22] = "        <<<     "
  f32[ 60][23] = "         <<<    "
  f32[ 60][24] = "          <<<   "
  f32[ 60][25] = "           <<<  "

  f32[ 61][12] = " =============  "
  f32[ 61][13] = " =============  "
  f32[ 61][14] = "                "
  f32[ 61][15] = "                "
  f32[ 61][16] = "                "
  f32[ 61][17] = "                "
  f32[ 61][18] = " =============  "
  f32[ 61][19] = " =============  "

  f32[ 62][ 6] = "  >>>           "
  f32[ 62][ 7] = "   >>>          "
  f32[ 62][ 8] = "    >>>         "
  f32[ 62][ 9] = "     >>>        "
  f32[ 62][10] = "      >>>       "
  f32[ 62][11] = "       >>>      "
  f32[ 62][12] = "        >>>     "
  f32[ 62][13] = "         >>>    "
  f32[ 62][14] = "          >>>   "
  f32[ 62][15] = "           >>>  "
  f32[ 62][16] = "           >>>  "
  f32[ 62][17] = "          >>>   "
  f32[ 62][18] = "         >>>    "
  f32[ 62][19] = "        >>>     "
  f32[ 62][20] = "       >>>      "
  f32[ 62][21] = "      >>>       "
  f32[ 62][22] = "     >>>        "
  f32[ 62][23] = "    >>>         "
  f32[ 62][24] = "   >>>          "
  f32[ 62][25] = "  >>>           "

  f32[ 63][ 6] = "   ?????????    "
  f32[ 63][ 7] = "  ???????????   "
  f32[ 63][ 8] = " ???       ???  "
  f32[ 63][ 9] = " ???       ???  "
  f32[ 63][10] = " ???       ???  "
  f32[ 63][11] = " ???       ???  "
  f32[ 63][12] = " ???       ???  "
  f32[ 63][13] = "          ???   "
  f32[ 63][14] = "         ???    "
  f32[ 63][15] = "        ???     "
  f32[ 63][16] = "       ???      "
  f32[ 63][17] = "      ???       "
  f32[ 63][18] = "      ???       "
  f32[ 63][19] = "      ???       "
  f32[ 63][20] = "                "
  f32[ 63][21] = "                "
  f32[ 63][22] = "      ???       "
  f32[ 63][23] = "      ???       "
  f32[ 63][24] = "      ???       "
  f32[ 63][25] = "      ???       "

  f32[ 64][ 6] = "   @@@@@@@@@@   "
  f32[ 64][ 7] = "  @@@@@@@@@@@@  "
  f32[ 64][ 8] = " @@@        @@@ "
  f32[ 64][ 9] = " @@@         @@ "
  f32[ 64][10] = " @@@   @@@@@@@@ "
  f32[ 64][11] = " @@@  @@@@@@@@@ "
  f32[ 64][12] = " @@@ @@@    @@@ "
  f32[ 64][13] = " @@@ @@@    @@@ "
  f32[ 64][14] = " @@@ @@@    @@@ "
  f32[ 64][15] = " @@@ @@@    @@@ "
  f32[ 64][16] = " @@@ @@@    @@@ "
  f32[ 64][17] = " @@@ @@@    @@@ "
  f32[ 64][18] = " @@@ @@@    @@@ "
  f32[ 64][19] = " @@@ @@@   @@@@ "
  f32[ 64][20] = " @@@  @@@@@@@@@ "
  f32[ 64][21] = " @@@   @@@@@ @@ "
  f32[ 64][22] = " @@@            "
  f32[ 64][23] = " @@@            "
  f32[ 64][24] = "  @@@@@@@@@@@@@ "
  f32[ 64][25] = "   @@@@@@@@@@@@ "

  f32[ 65][ 6] = "   AAAAAAAAA    "
  f32[ 65][ 7] = "  AAAAAAAAAAA   "
  f32[ 65][ 8] = " AAA       AAA  "
  f32[ 65][ 9] = " AAA       AAA  "
  f32[ 65][10] = " AAA       AAA  "
  f32[ 65][11] = " AAA       AAA  "
  f32[ 65][12] = " AAA       AAA  "
  f32[ 65][13] = " AAA       AAA  "
  f32[ 65][14] = " AAA       AAA  "
  f32[ 65][15] = " AAA       AAA  "
  f32[ 65][16] = " AAAAAAAAAAAAA  "
  f32[ 65][17] = " AAAAAAAAAAAAA  "
  f32[ 65][18] = " AAA       AAA  "
  f32[ 65][19] = " AAA       AAA  "
  f32[ 65][20] = " AAA       AAA  "
  f32[ 65][21] = " AAA       AAA  "
  f32[ 65][22] = " AAA       AAA  "
  f32[ 65][23] = " AAA       AAA  "
  f32[ 65][24] = " AAA       AAA  "
  f32[ 65][25] = " AAA       AAA  "

  f32[196][ 1] = "  AAA     AAA   "
  f32[196][ 2] = "  AAA     AAA   "
  f32[196][ 3] = "  AAA     AAA   "
  f32[196][ 4] = "  AAA     AAA   "
  f32[196][ 5] = "                "
  f32[196][ 6] = "   AAAAAAAAA    "
  f32[196][ 7] = "  AAAAAAAAAAA   "
  f32[196][ 8] = " AAA       AAA  "
  f32[196][ 9] = " AAA       AAA  "
  f32[196][10] = " AAA       AAA  "
  f32[196][11] = " AAA       AAA  "
  f32[196][12] = " AAA       AAA  "
  f32[196][13] = " AAA       AAA  "
  f32[196][14] = " AAA       AAA  "
  f32[196][15] = " AAA       AAA  "
  f32[196][16] = " AAAAAAAAAAAAA  "
  f32[196][17] = " AAAAAAAAAAAAA  "
  f32[196][18] = " AAA       AAA  "
  f32[196][19] = " AAA       AAA  "
  f32[196][20] = " AAA       AAA  "
  f32[196][21] = " AAA       AAA  "
  f32[196][22] = " AAA       AAA  "
  f32[196][23] = " AAA       AAA  "
  f32[196][24] = " AAA       AAA  "
  f32[196][25] = " AAA       AAA  "

  f32[ 66][ 6] = " BBBBBBBBBBB    "
  f32[ 66][ 7] = " BBBBBBBBBBBB   "
  f32[ 66][ 8] = " BBB       BBB  "
  f32[ 66][ 9] = " BBB       BBB  "
  f32[ 66][10] = " BBB       BBB  "
  f32[ 66][11] = " BBB       BBB  "
  f32[ 66][12] = " BBB       BBB  "
  f32[ 66][13] = " BBB      BBB   "
  f32[ 66][14] = " BBBBBBBBBBB    "
  f32[ 66][15] = " BBBBBBBBBBBB   "
  f32[ 66][16] = " BBB       BBB  "
  f32[ 66][17] = " BBB       BBB  "
  f32[ 66][18] = " BBB       BBB  "
  f32[ 66][19] = " BBB       BBB  "
  f32[ 66][20] = " BBB       BBB  "
  f32[ 66][21] = " BBB       BBB  "
  f32[ 66][22] = " BBB       BBB  "
  f32[ 66][23] = " BBB       BBB  "
  f32[ 66][24] = " BBBBBBBBBBBB   "
  f32[ 66][25] = " BBBBBBBBBBB    "

  f32[ 67][ 6] = "   CCCCCCCCC    "
  f32[ 67][ 7] = "  CCCCCCCCCCC   "
  f32[ 67][ 8] = " CCC       CCC  "
  f32[ 67][ 9] = " CCC       CCC  "
  f32[ 67][10] = " CCC       CCC  "
  f32[ 67][11] = " CCC            "
  f32[ 67][12] = " CCC            "
  f32[ 67][13] = " CCC            "
  f32[ 67][14] = " CCC            "
  f32[ 67][15] = " CCC            "
  f32[ 67][16] = " CCC            "
  f32[ 67][17] = " CCC            "
  f32[ 67][18] = " CCC            "
  f32[ 67][19] = " CCC            "
  f32[ 67][20] = " CCC            "
  f32[ 67][21] = " CCC       CCC  "
  f32[ 67][22] = " CCC       CCC  "
  f32[ 67][23] = " CCC       CCC  "
  f32[ 67][24] = "  CCCCCCCCCCC   "
  f32[ 67][25] = "   CCCCCCCCC    "

  f32[ 68][ 6] = " DDDDDDDDDD     "
  f32[ 68][ 7] = " DDDDDDDDDDD    "
  f32[ 68][ 8] = " DDD      DDD   "
  f32[ 68][ 9] = " DDD       DDD  "
  f32[ 68][10] = " DDD       DDD  "
  f32[ 68][11] = " DDD       DDD  "
  f32[ 68][12] = " DDD       DDD  "
  f32[ 68][13] = " DDD       DDD  "
  f32[ 68][14] = " DDD       DDD  "
  f32[ 68][15] = " DDD       DDD  "
  f32[ 68][16] = " DDD       DDD  "
  f32[ 68][17] = " DDD       DDD  "
  f32[ 68][18] = " DDD       DDD  "
  f32[ 68][19] = " DDD       DDD  "
  f32[ 68][20] = " DDD       DDD  "
  f32[ 68][21] = " DDD       DDD  "
  f32[ 68][22] = " DDD       DDD  "
  f32[ 68][23] = " DDD      DDD   "
  f32[ 68][24] = " DDDDDDDDDDD    "
  f32[ 68][25] = " DDDDDDDDDD     "

  f32[ 69][ 6] = " EEEEEEEEEEEEE  "
  f32[ 69][ 7] = " EEEEEEEEEEEEE  "
  f32[ 69][ 8] = " EEE            "
  f32[ 69][ 9] = " EEE            "
  f32[ 69][10] = " EEE            "
  f32[ 69][11] = " EEE            "
  f32[ 69][12] = " EEE            "
  f32[ 69][13] = " EEE            "
  f32[ 69][14] = " EEE            "
  f32[ 69][15] = " EEEEEEEEEE     "
  f32[ 69][16] = " EEEEEEEEEE     "
  f32[ 69][17] = " EEE            "
  f32[ 69][18] = " EEE            "
  f32[ 69][19] = " EEE            "
  f32[ 69][20] = " EEE            "
  f32[ 69][21] = " EEE            "
  f32[ 69][22] = " EEE            "
  f32[ 69][23] = " EEE            "
  f32[ 69][24] = " EEEEEEEEEEEEE  "
  f32[ 69][25] = " EEEEEEEEEEEEE  "

  f32[ 70][ 6] = " FFFFFFFFFFFFF  "
  f32[ 70][ 7] = " FFFFFFFFFFFFF  "
  f32[ 70][ 8] = " FFF            "
  f32[ 70][ 9] = " FFF            "
  f32[ 70][10] = " FFF            "
  f32[ 70][11] = " FFF            "
  f32[ 70][12] = " FFF            "
  f32[ 70][13] = " FFF            "
  f32[ 70][14] = " FFF            "
  f32[ 70][15] = " FFFFFFFFFF     "
  f32[ 70][16] = " FFFFFFFFFF     "
  f32[ 70][17] = " FFF            "
  f32[ 70][18] = " FFF            "
  f32[ 70][19] = " FFF            "
  f32[ 70][20] = " FFF            "
  f32[ 70][21] = " FFF            "
  f32[ 70][22] = " FFF            "
  f32[ 70][23] = " FFF            "
  f32[ 70][24] = " FFF            "
  f32[ 70][25] = " FFF            "

  f32[ 71][ 6] = "   GGGGGGGGG    "
  f32[ 71][ 7] = "  GGGGGGGGGGG   "
  f32[ 71][ 8] = " GGG       GGG  "
  f32[ 71][ 9] = " GGG       GGG  "
  f32[ 71][10] = " GGG       GGG  "
  f32[ 71][11] = " GGG            "
  f32[ 71][12] = " GGG            "
  f32[ 71][13] = " GGG            "
  f32[ 71][14] = " GGG            "
  f32[ 71][15] = " GGG   GGGGGGG  "
  f32[ 71][16] = " GGG   GGGGGGG  "
  f32[ 71][17] = " GGG       GGG  "
  f32[ 71][18] = " GGG       GGG  "
  f32[ 71][19] = " GGG       GGG  "
  f32[ 71][20] = " GGG       GGG  "
  f32[ 71][21] = " GGG       GGG  "
  f32[ 71][22] = " GGG       GGG  "
  f32[ 71][23] = " GGG       GGG  "
  f32[ 71][24] = "  GGGGGGGGGGG   "
  f32[ 71][25] = "   GGGGGGGGG    "

  f32[ 72][ 6] = " HHH       HHH  "
  f32[ 72][ 7] = " HHH       HHH  "
  f32[ 72][ 8] = " HHH       HHH  "
  f32[ 72][ 9] = " HHH       HHH  "
  f32[ 72][10] = " HHH       HHH  "
  f32[ 72][11] = " HHH       HHH  "
  f32[ 72][12] = " HHH       HHH  "
  f32[ 72][13] = " HHH       HHH  "
  f32[ 72][14] = " HHH       HHH  "
  f32[ 72][15] = " HHHHHHHHHHHHH  "
  f32[ 72][16] = " HHHHHHHHHHHHH  "
  f32[ 72][17] = " HHH       HHH  "
  f32[ 72][18] = " HHH       HHH  "
  f32[ 72][19] = " HHH       HHH  "
  f32[ 72][20] = " HHH       HHH  "
  f32[ 72][21] = " HHH       HHH  "
  f32[ 72][22] = " HHH       HHH  "
  f32[ 72][23] = " HHH       HHH  "
  f32[ 72][24] = " HHH       HHH  "
  f32[ 72][25] = " HHH       HHH  "

  f32[ 73][ 6] = "    IIIIIII     "
  f32[ 73][ 7] = "    IIIIIII     "
  f32[ 73][ 8] = "      III       "
  f32[ 73][ 9] = "      III       "
  f32[ 73][10] = "      III       "
  f32[ 73][11] = "      III       "
  f32[ 73][12] = "      III       "
  f32[ 73][13] = "      III       "
  f32[ 73][14] = "      III       "
  f32[ 73][15] = "      III       "
  f32[ 73][16] = "      III       "
  f32[ 73][17] = "      III       "
  f32[ 73][18] = "      III       "
  f32[ 73][19] = "      III       "
  f32[ 73][20] = "      III       "
  f32[ 73][21] = "      III       "
  f32[ 73][22] = "      III       "
  f32[ 73][23] = "      III       "
  f32[ 73][24] = "    IIIIIII     "
  f32[ 73][25] = "    IIIIIII     "

  f32[ 74][ 6] = "        JJJJJJJ "
  f32[ 74][ 7] = "        JJJJJJJ "
  f32[ 74][ 8] = "          JJJ   "
  f32[ 74][ 9] = "          JJJ   "
  f32[ 74][10] = "          JJJ   "
  f32[ 74][11] = "          JJJ   "
  f32[ 74][12] = "          JJJ   "
  f32[ 74][13] = "          JJJ   "
  f32[ 74][14] = "          JJJ   "
  f32[ 74][15] = "          JJJ   "
  f32[ 74][16] = "          JJJ   "
  f32[ 74][17] = "          JJJ   "
  f32[ 74][18] = "          JJJ   "
  f32[ 74][19] = "          JJJ   "
  f32[ 74][20] = " JJJ      JJJ   "
  f32[ 74][21] = " JJJ      JJJ   "
  f32[ 74][22] = " JJJ      JJJ   "
  f32[ 74][23] = " JJJ      JJJ   "
  f32[ 74][24] = "  JJJJJJJJJJ    "
  f32[ 74][25] = "   JJJJJJJJ     "

  f32[ 75][ 6] = " KKK        KK  "
  f32[ 75][ 7] = " KKK       KKK  "
  f32[ 75][ 8] = " KKK      KKK   "
  f32[ 75][ 9] = " KKK     KKK    "
  f32[ 75][10] = " KKK    KKK     "
  f32[ 75][11] = " KKK   KKK      "
  f32[ 75][12] = " KKK  KKK       "
  f32[ 75][13] = " KKK KKK        "
  f32[ 75][14] = " KKKKKK         "
  f32[ 75][15] = " KKKKK          "
  f32[ 75][16] = " KKKKK          "
  f32[ 75][17] = " KKKKKK         "
  f32[ 75][18] = " KKK KKK        "
  f32[ 75][19] = " KKK  KKK       "
  f32[ 75][20] = " KKK   KKK      "
  f32[ 75][21] = " KKK    KKK     "
  f32[ 75][22] = " KKK     KKK    "
  f32[ 75][23] = " KKK      KKK   "
  f32[ 75][24] = " KKK       KKK  "
  f32[ 75][25] = " KKK        KK  "

  f32[ 76][ 6] = " LLL            "
  f32[ 76][ 7] = " LLL            "
  f32[ 76][ 8] = " LLL            "
  f32[ 76][ 9] = " LLL            "
  f32[ 76][10] = " LLL            "
  f32[ 76][11] = " LLL            "
  f32[ 76][12] = " LLL            "
  f32[ 76][13] = " LLL            "
  f32[ 76][14] = " LLL            "
  f32[ 76][15] = " LLL            "
  f32[ 76][16] = " LLL            "
  f32[ 76][17] = " LLL            "
  f32[ 76][18] = " LLL            "
  f32[ 76][19] = " LLL            "
  f32[ 76][20] = " LLL            "
  f32[ 76][21] = " LLL            "
  f32[ 76][22] = " LLL            "
  f32[ 76][23] = " LLL            "
  f32[ 76][24] = " LLLLLLLLLLLLL  "
  f32[ 76][25] = " LLLLLLLLLLLLL  "

  f32[ 77][ 6] = " MM          MM "
  f32[ 77][ 7] = " MMM        MMM "
  f32[ 77][ 8] = " MMMM      MMMM "
  f32[ 77][ 9] = " MMMMM    MMMMM "
  f32[ 77][10] = " MMMMMM  MMMMMM "
  f32[ 77][11] = " MMM MMMMMM MMM "
  f32[ 77][12] = " MMM  MMMM  MMM "
  f32[ 77][13] = " MMM   MM   MMM "
  f32[ 77][14] = " MMM        MMM "
  f32[ 77][15] = " MMM        MMM "
  f32[ 77][16] = " MMM        MMM "
  f32[ 77][17] = " MMM        MMM "
  f32[ 77][18] = " MMM        MMM "
  f32[ 77][19] = " MMM        MMM "
  f32[ 77][20] = " MMM        MMM "
  f32[ 77][21] = " MMM        MMM "
  f32[ 77][22] = " MMM        MMM "
  f32[ 77][23] = " MMM        MMM "
  f32[ 77][24] = " MMM        MMM "
  f32[ 77][25] = " MMM        MMM "

  f32[ 78][ 6] = " NNN       NNN  "
  f32[ 78][ 7] = " NNN       NNN  "
  f32[ 78][ 8] = " NNN       NNN  "
  f32[ 78][ 9] = " NNN       NNN  "
  f32[ 78][10] = " NNN       NNN  "
  f32[ 78][11] = " NNNN      NNN  "
  f32[ 78][12] = " NNNNN     NNN  "
  f32[ 78][13] = " NNNNNN    NNN  "
  f32[ 78][14] = " NNN NNN   NNN  "
  f32[ 78][15] = " NNN  NNN  NNN  "
  f32[ 78][16] = " NNN   NNN NNN  "
  f32[ 78][17] = " NNN    NNNNNN  "
  f32[ 78][18] = " NNN     NNNNN  "
  f32[ 78][19] = " NNN      NNNN  "
  f32[ 78][20] = " NNN       NNN  "
  f32[ 78][21] = " NNN       NNN  "
  f32[ 78][22] = " NNN       NNN  "
  f32[ 78][23] = " NNN       NNN  "
  f32[ 78][24] = " NNN       NNN  "
  f32[ 78][25] = " NNN       NNN  "

  f32[ 79][ 6] = "   OOOOOOOOO    "
  f32[ 79][ 7] = "  OOOOOOOOOOO   "
  f32[ 79][ 8] = " OOO       OOO  "
  f32[ 79][ 9] = " OOO       OOO  "
  f32[ 79][10] = " OOO       OOO  "
  f32[ 79][11] = " OOO       OOO  "
  f32[ 79][12] = " OOO       OOO  "
  f32[ 79][13] = " OOO       OOO  "
  f32[ 79][14] = " OOO       OOO  "
  f32[ 79][15] = " OOO       OOO  "
  f32[ 79][16] = " OOO       OOO  "
  f32[ 79][17] = " OOO       OOO  "
  f32[ 79][18] = " OOO       OOO  "
  f32[ 79][19] = " OOO       OOO  "
  f32[ 79][20] = " OOO       OOO  "
  f32[ 79][21] = " OOO       OOO  "
  f32[ 79][22] = " OOO       OOO  "
  f32[ 79][23] = " OOO       OOO  "
  f32[ 79][24] = "  OOOOOOOOOOO   "
  f32[ 79][25] = "   OOOOOOOOO    "

  f32[214][ 0] = "                "
  f32[214][ 1] = "   OOO   OOO    "
  f32[214][ 2] = "   OOO   OOO    "
  f32[214][ 3] = "   OOO   OOO    "
  f32[214][ 4] = "   OOO   OOO    "
  f32[214][ 5] = "                "
  f32[214][ 6] = "   OOOOOOOOO    "
  f32[214][ 7] = "  OOOOOOOOOOO   "
  f32[214][ 8] = " OOO       OOO  "
  f32[214][ 9] = " OOO       OOO  "
  f32[214][10] = " OOO       OOO  "
  f32[214][11] = " OOO       OOO  "
  f32[214][12] = " OOO       OOO  "
  f32[214][13] = " OOO       OOO  "
  f32[214][14] = " OOO       OOO  "
  f32[214][15] = " OOO       OOO  "
  f32[214][16] = " OOO       OOO  "
  f32[214][17] = " OOO       OOO  "
  f32[214][18] = " OOO       OOO  "
  f32[214][19] = " OOO       OOO  "
  f32[214][20] = " OOO       OOO  "
  f32[214][21] = " OOO       OOO  "
  f32[214][22] = " OOO       OOO  "
  f32[214][23] = " OOO       OOO  "
  f32[214][24] = "  OOOOOOOOOOO   "
  f32[214][25] = "   OOOOOOOOO    "

  f32[ 80][ 6] = " PPPPPPPPPPP    "
  f32[ 80][ 7] = " PPPPPPPPPPPP   "
  f32[ 80][ 8] = " PPP       PPP  "
  f32[ 80][ 9] = " PPP       PPP  "
  f32[ 80][10] = " PPP       PPP  "
  f32[ 80][11] = " PPP       PPP  "
  f32[ 80][12] = " PPP       PPP  "
  f32[ 80][13] = " PPP       PPP  "
  f32[ 80][14] = " PPP       PPP  "
  f32[ 80][15] = " PPP       PPP  "
  f32[ 80][16] = " PPPPPPPPPPPP   "
  f32[ 80][17] = " PPPPPPPPPPP    "
  f32[ 80][18] = " PPP            "
  f32[ 80][19] = " PPP            "
  f32[ 80][20] = " PPP            "
  f32[ 80][21] = " PPP            "
  f32[ 80][22] = " PPP            "
  f32[ 80][23] = " PPP            "
  f32[ 80][24] = " PPP            "
  f32[ 80][25] = " PPP            "

  f32[ 81][ 6] = "   QQQQQQQQQ    "
  f32[ 81][ 7] = "  QQQQQQQQQQQ   "
  f32[ 81][ 8] = " QQQ       QQQ  "
  f32[ 81][ 9] = " QQQ       QQQ  "
  f32[ 81][10] = " QQQ       QQQ  "
  f32[ 81][11] = " QQQ       QQQ  "
  f32[ 81][12] = " QQQ       QQQ  "
  f32[ 81][13] = " QQQ       QQQ  "
  f32[ 81][14] = " QQQ       QQQ  "
  f32[ 81][15] = " QQQ       QQQ  "
  f32[ 81][16] = " QQQ       QQQ  "
  f32[ 81][17] = " QQQ       QQQ  "
  f32[ 81][18] = " QQQ       QQQ  "
  f32[ 81][19] = " QQQ       QQQ  "
  f32[ 81][20] = " QQQ       QQQ  "
  f32[ 81][21] = " QQQ       QQQ  "
  f32[ 81][22] = " QQQ  QQQ  QQQ  "
  f32[ 81][23] = " QQQ   QQQ QQQ  "
  f32[ 81][24] = "  QQQQQQQQQQQ   "
  f32[ 81][25] = "   QQQQQQQQQ    "
  f32[ 81][26] = "          QQQ   "
  f32[ 81][27] = "           QQQ  "

  f32[ 82][ 6] = " RRRRRRRRRRR    "
  f32[ 82][ 7] = " RRRRRRRRRRRR   "
  f32[ 82][ 8] = " RRR       RRR  "
  f32[ 82][ 9] = " RRR       RRR  "
  f32[ 82][10] = " RRR       RRR  "
  f32[ 82][11] = " RRR       RRR  "
  f32[ 82][12] = " RRR       RRR  "
  f32[ 82][13] = " RRR       RRR  "
  f32[ 82][14] = " RRR       RRR  "
  f32[ 82][15] = " RRR       RRR  "
  f32[ 82][16] = " RRRRRRRRRRRR   "
  f32[ 82][17] = " RRRRRRRRRRR    "
  f32[ 82][18] = " RRRRRR         "
  f32[ 82][19] = " RRR RRR        "
  f32[ 82][20] = " RRR  RRR       "
  f32[ 82][21] = " RRR   RRR      "
  f32[ 82][22] = " RRR    RRR     "
  f32[ 82][23] = " RRR     RRR    "
  f32[ 82][24] = " RRR      RRR   "
  f32[ 82][25] = " RRR       RRR  "

  f32[ 83][ 6] = "   SSSSSSSSS    "
  f32[ 83][ 7] = "  SSSSSSSSSSS   "
  f32[ 83][ 8] = " SSS       SSS  "
  f32[ 83][ 9] = " SSS       SSS  "
  f32[ 83][10] = " SSS            "
  f32[ 83][11] = " SSS            "
  f32[ 83][12] = " SSS            "
  f32[ 83][13] = " SSS            "
  f32[ 83][14] = " SSS            "
  f32[ 83][15] = "  SSSSSSSSSS    "
  f32[ 83][16] = "   SSSSSSSSSS   "
  f32[ 83][17] = "           SSS  "
  f32[ 83][18] = "           SSS  "
  f32[ 83][19] = "           SSS  "
  f32[ 83][20] = "           SSS  "
  f32[ 83][21] = "           SSS  "
  f32[ 83][22] = " SSS       SSS  "
  f32[ 83][23] = " SSS       SSS  "
  f32[ 83][24] = "  SSSSSSSSSSS   "
  f32[ 83][25] = "   SSSSSSSSS    "

  f32[ 84][ 6] = " TTTTTTTTTTTTT  "
  f32[ 84][ 7] = " TTTTTTTTTTTTT  "
  f32[ 84][ 8] = "      TTT       "
  f32[ 84][ 9] = "      TTT       "
  f32[ 84][10] = "      TTT       "
  f32[ 84][11] = "      TTT       "
  f32[ 84][12] = "      TTT       "
  f32[ 84][13] = "      TTT       "
  f32[ 84][14] = "      TTT       "
  f32[ 84][15] = "      TTT       "
  f32[ 84][16] = "      TTT       "
  f32[ 84][17] = "      TTT       "
  f32[ 84][18] = "      TTT       "
  f32[ 84][19] = "      TTT       "
  f32[ 84][20] = "      TTT       "
  f32[ 84][21] = "      TTT       "
  f32[ 84][22] = "      TTT       "
  f32[ 84][23] = "      TTT       "
  f32[ 84][24] = "      TTT       "
  f32[ 84][25] = "      TTT       "

  f32[ 85][ 6] = " UUU       UUU  "
  f32[ 85][ 7] = " UUU       UUU  "
  f32[ 85][ 8] = " UUU       UUU  "
  f32[ 85][ 9] = " UUU       UUU  "
  f32[ 85][10] = " UUU       UUU  "
  f32[ 85][11] = " UUU       UUU  "
  f32[ 85][12] = " UUU       UUU  "
  f32[ 85][13] = " UUU       UUU  "
  f32[ 85][14] = " UUU       UUU  "
  f32[ 85][15] = " UUU       UUU  "
  f32[ 85][16] = " UUU       UUU  "
  f32[ 85][17] = " UUU       UUU  "
  f32[ 85][18] = " UUU       UUU  "
  f32[ 85][19] = " UUU       UUU  "
  f32[ 85][20] = " UUU       UUU  "
  f32[ 85][21] = " UUU       UUU  "
  f32[ 85][22] = " UUU       UUU  "
  f32[ 85][23] = " UUU       UUU  "
  f32[ 85][24] = "  UUUUUUUUUUU   "
  f32[ 85][25] = "   UUUUUUUUU    "

  f32[220][ 1] = "   UUU   UUU    "
  f32[220][ 2] = "   UUU   UUU    "
  f32[220][ 3] = "   UUU   UUU    "
  f32[220][ 4] = "   UUU   UUU    "
  f32[220][ 5] = "                "
  f32[220][ 6] = " UUU       UUU  "
  f32[220][ 7] = " UUU       UUU  "
  f32[220][ 8] = " UUU       UUU  "
  f32[220][ 9] = " UUU       UUU  "
  f32[220][10] = " UUU       UUU  "
  f32[220][11] = " UUU       UUU  "
  f32[220][12] = " UUU       UUU  "
  f32[220][13] = " UUU       UUU  "
  f32[220][14] = " UUU       UUU  "
  f32[220][15] = " UUU       UUU  "
  f32[220][16] = " UUU       UUU  "
  f32[220][17] = " UUU       UUU  "
  f32[220][18] = " UUU       UUU  "
  f32[220][19] = " UUU       UUU  "
  f32[220][20] = " UUU       UUU  "
  f32[220][21] = " UUU       UUU  "
  f32[220][22] = " UUU       UUU  "
  f32[220][23] = " UUU       UUU  "
  f32[220][24] = "  UUUUUUUUUUU   "
  f32[220][25] = "   UUUUUUUUU    "

  f32[ 86][ 6] = " VVV       VVV  "
  f32[ 86][ 7] = " VVV       VVV  "
  f32[ 86][ 8] = " VVV       VVV  "
  f32[ 86][ 9] = " VVV       VVV  "
  f32[ 86][10] = "  VVV     VVV   "
  f32[ 86][11] = "  VVV     VVV   "
  f32[ 86][12] = "  VVV     VVV   "
  f32[ 86][13] = "  VVV     VVV   "
  f32[ 86][14] = "   VVV   VVV    "
  f32[ 86][15] = "   VVV   VVV    "
  f32[ 86][16] = "   VVV   VVV    "
  f32[ 86][17] = "   VVV   VVV    "
  f32[ 86][18] = "    VVV VVV     "
  f32[ 86][19] = "    VVV VVV     "
  f32[ 86][20] = "    VVV VVV     "
  f32[ 86][21] = "    VVV VVV     "
  f32[ 86][22] = "     VVVVV      "
  f32[ 86][23] = "     VVVVV      "
  f32[ 86][24] = "     VVVVV      "
  f32[ 86][25] = "      VVV       "

  f32[ 87][ 6] = " WWW        WWW "
  f32[ 87][ 7] = " WWW        WWW "
  f32[ 87][ 8] = " WWW        WWW "
  f32[ 87][ 9] = " WWW        WWW "
  f32[ 87][10] = " WWW        WWW "
  f32[ 87][11] = " WWW        WWW "
  f32[ 87][12] = " WWW        WWW "
  f32[ 87][13] = " WWW        WWW "
  f32[ 87][14] = " WWW        WWW "
  f32[ 87][15] = " WWW        WWW "
  f32[ 87][16] = " WWW        WWW "
  f32[ 87][17] = " WWW        WWW "
  f32[ 87][18] = " WWW   WW   WWW "
  f32[ 87][19] = " WWW  WWWW  WWW "
  f32[ 87][20] = " WWW WWWWWW WWW "
  f32[ 87][21] = " WWWWWW  WWWWWW "
  f32[ 87][22] = " WWWWW    WWWWW "
  f32[ 87][23] = " WWWW      WWWW "
  f32[ 87][24] = " WWW        WWW "
  f32[ 87][25] = " WW          WW "

  f32[ 88][ 6] = " XXX       XXX  "
  f32[ 88][ 7] = " XXX       XXX  "
  f32[ 88][ 8] = "  XXX     XXX   "
  f32[ 88][ 9] = "  XXX     XXX   "
  f32[ 88][10] = "   XXX   XXX    "
  f32[ 88][11] = "   XXX   XXX    "
  f32[ 88][12] = "    XXX XXX     "
  f32[ 88][13] = "    XXX XXX     "
  f32[ 88][14] = "     XXXXX      "
  f32[ 88][15] = "     XXXXX      "
  f32[ 88][16] = "     XXXXX      "
  f32[ 88][17] = "     XXXXX      "
  f32[ 88][18] = "    XXX XXX     "
  f32[ 88][19] = "    XXX XXX     "
  f32[ 88][20] = "   XXX   XXX    "
  f32[ 88][21] = "   XXX   XXX    "
  f32[ 88][22] = "  XXX     XXX   "
  f32[ 88][23] = "  XXX     XXX   "
  f32[ 88][24] = " XXX       XXX  "
  f32[ 88][25] = " XXX       XXX  "

  f32[ 89][ 6] = " YYY       YYY  "
  f32[ 89][ 7] = " YYY       YYY  "
  f32[ 89][ 8] = " YYY       YYY  "
  f32[ 89][ 9] = "  YYY     YYY   "
  f32[ 89][10] = "  YYY     YYY   "
  f32[ 89][11] = "   YYY   YYY    "
  f32[ 89][12] = "   YYY   YYY    "
  f32[ 89][13] = "    YYY YYY     "
  f32[ 89][14] = "    YYY YYY     "
  f32[ 89][15] = "     YYYYY      "
  f32[ 89][16] = "     YYYYY      "
  f32[ 89][17] = "      YYY       "
  f32[ 89][18] = "      YYY       "
  f32[ 89][19] = "      YYY       "
  f32[ 89][20] = "      YYY       "
  f32[ 89][21] = "      YYY       "
  f32[ 89][22] = "      YYY       "
  f32[ 89][23] = "      YYY       "
  f32[ 89][24] = "      YYY       "
  f32[ 89][25] = "      YYY       "

  f32[ 90][ 6] = " ZZZZZZZZZZZZZ  "
  f32[ 90][ 7] = " ZZZZZZZZZZZZZ  "
  f32[ 90][ 8] = "           ZZZ  "
  f32[ 90][ 9] = "           ZZZ  "
  f32[ 90][10] = "           ZZZ  "
  f32[ 90][11] = "          ZZZ   "
  f32[ 90][12] = "         ZZZ    "
  f32[ 90][13] = "        ZZZ     "
  f32[ 90][14] = "       ZZZ      "
  f32[ 90][15] = "      ZZZ       "
  f32[ 90][16] = "     ZZZ        "
  f32[ 90][17] = "    ZZZ         "
  f32[ 90][18] = "   ZZZ          "
  f32[ 90][19] = "  ZZZ           "
  f32[ 90][20] = " ZZZ            "
  f32[ 90][21] = " ZZZ            "
  f32[ 90][22] = " ZZZ            "
  f32[ 90][23] = " ZZZ            "
  f32[ 90][24] = " ZZZZZZZZZZZZZ  "
  f32[ 90][25] = " ZZZZZZZZZZZZZ  "

  f32[ 91][ 6] = "    [[[[[[[[    "
  f32[ 91][ 7] = "    [[[[[[[[    "
  f32[ 91][ 8] = "    [[[         "
  f32[ 91][ 9] = "    [[[         "
  f32[ 91][10] = "    [[[         "
  f32[ 91][11] = "    [[[         "
  f32[ 91][12] = "    [[[         "
  f32[ 91][13] = "    [[[         "
  f32[ 91][14] = "    [[[         "
  f32[ 91][15] = "    [[[         "
  f32[ 91][16] = "    [[[         "
  f32[ 91][17] = "    [[[         "
  f32[ 91][18] = "    [[[         "
  f32[ 91][19] = "    [[[         "
  f32[ 91][20] = "    [[[         "
  f32[ 91][21] = "    [[[         "
  f32[ 91][22] = "    [[[         "
  f32[ 91][23] = "    [[[         "
  f32[ 91][24] = "    [[[[[[[[    "
  f32[ 91][25] = "    [[[[[[[[    "

  f32[ 92][ 6] = "  ///           "
  f32[ 92][ 7] = "  ///           "
  f32[ 92][ 8] = "   ///          "
  f32[ 92][ 9] = "   ///          "
  f32[ 92][10] = "    ///         "
  f32[ 92][11] = "    ///         "
  f32[ 92][12] = "     ///        "
  f32[ 92][13] = "     ///        "
  f32[ 92][14] = "      ///       "
  f32[ 92][15] = "      ///       "
  f32[ 92][16] = "       ///      "
  f32[ 92][17] = "       ///      "
  f32[ 92][18] = "        ///     "
  f32[ 92][19] = "        ///     "
  f32[ 92][20] = "         ///    "
  f32[ 92][21] = "         ///    "
  f32[ 92][22] = "          ///   "
  f32[ 92][23] = "          ///   "
  f32[ 92][24] = "           ///  "
  f32[ 92][25] = "           ///  "

  f32[ 93][ 6] = "    ]]]]]]]]    "
  f32[ 93][ 7] = "    ]]]]]]]]    "
  f32[ 93][ 8] = "         ]]]    "
  f32[ 93][ 9] = "         ]]]    "
  f32[ 93][10] = "         ]]]    "
  f32[ 93][11] = "         ]]]    "
  f32[ 93][12] = "         ]]]    "
  f32[ 93][13] = "         ]]]    "
  f32[ 93][14] = "         ]]]    "
  f32[ 93][15] = "         ]]]    "
  f32[ 93][16] = "         ]]]    "
  f32[ 93][17] = "         ]]]    "
  f32[ 93][18] = "         ]]]    "
  f32[ 93][19] = "         ]]]    "
  f32[ 93][20] = "         ]]]    "
  f32[ 93][21] = "         ]]]    "
  f32[ 93][22] = "         ]]]    "
  f32[ 93][23] = "         ]]]    "
  f32[ 93][24] = "    ]]]]]]]]    "
  f32[ 93][25] = "    ]]]]]]]]    "

  f32[ 94][ 3] = "      ^^^       "
  f32[ 94][ 4] = "     ^^^^^      "
  f32[ 94][ 5] = "    ^^^ ^^^     "
  f32[ 94][ 6] = "   ^^^   ^^^    "
  f32[ 94][ 7] = "  ^^^     ^^^   "
  f32[ 94][ 7] = " ^^^       ^^^  "

  f32[ 95][27] = " _____________  "
  f32[ 95][28] = " _____________  "

  f32[ 96][ 1] = "     ```        "
  f32[ 96][ 2] = "     ```        "
  f32[ 96][ 3] = "     ```        "
  f32[ 96][ 4] = "      ```       "

  f32[ 97][12] = "  aaaaaaaaaa    "
  f32[ 97][13] = "  aaaaaaaaaaa   "
  f32[ 97][14] = "           aaa  "
  f32[ 97][15] = "           aaa  "
  f32[ 97][16] = "           aaa  "
  f32[ 97][17] = "   aaaaaaaaaaa  "
  f32[ 97][18] = "  aaaaaaaaaaaa  "
  f32[ 97][19] = " aaa       aaa  "
  f32[ 97][20] = " aaa       aaa  "
  f32[ 97][21] = " aaa       aaa  "
  f32[ 97][22] = " aaa       aaa  "
  f32[ 97][23] = " aaa       aaa  "
  f32[ 97][24] = "  aaaaaaaaaaaa  "
  f32[ 97][25] = "   aaaaaaaaaaa  "

  f32[228][ 6] = "   aaa   aaa    "
  f32[228][ 7] = "   aaa   aaa    "
  f32[228][ 8] = "   aaa   aaa    "
  f32[228][ 9] = "   aaa   aaa    "
  f32[228][10] = "                "
  f32[228][11] = "                "
  f32[228][12] = "  aaaaaaaaaa    "
  f32[228][13] = "  aaaaaaaaaaa   "
  f32[228][14] = "           aaa  "
  f32[228][15] = "           aaa  "
  f32[228][16] = "           aaa  "
  f32[228][17] = "   aaaaaaaaaaa  "
  f32[228][18] = "  aaaaaaaaaaaa  "
  f32[228][19] = " aaa       aaa  "
  f32[228][20] = " aaa       aaa  "
  f32[228][21] = " aaa       aaa  "
  f32[228][22] = " aaa       aaa  "
  f32[228][23] = " aaa       aaa  "
  f32[228][24] = "  aaaaaaaaaaaa  "
  f32[228][25] = "   aaaaaaaaaaa  "

  f32[ 98][ 6] = " bbb            "
  f32[ 98][ 7] = " bbb            "
  f32[ 98][ 8] = " bbb            "
  f32[ 98][ 9] = " bbb            "
  f32[ 98][10] = " bbb            "
  f32[ 98][11] = " bbb            "
  f32[ 98][12] = " bbbbbbbbbbb    "
  f32[ 98][13] = " bbbbbbbbbbbb   "
  f32[ 98][14] = " bbb       bbb  "
  f32[ 98][15] = " bbb       bbb  "
  f32[ 98][16] = " bbb       bbb  "
  f32[ 98][17] = " bbb       bbb  "
  f32[ 98][18] = " bbb       bbb  "
  f32[ 98][19] = " bbb       bbb  "
  f32[ 98][20] = " bbb       bbb  "
  f32[ 98][21] = " bbb       bbb  "
  f32[ 98][22] = " bbb       bbb  "
  f32[ 98][23] = " bbb       bbb  "
  f32[ 98][24] = " bbbbbbbbbbbb   "
  f32[ 98][25] = " bbbbbbbbbbb    "

  f32[ 99][12] = "   ccccccccc    "
  f32[ 99][13] = "  ccccccccccc   "
  f32[ 99][14] = " ccc       ccc  "
  f32[ 99][15] = " ccc       ccc  "
  f32[ 99][16] = " ccc            "
  f32[ 99][17] = " ccc            "
  f32[ 99][18] = " ccc            "
  f32[ 99][19] = " ccc            "
  f32[ 99][20] = " ccc            "
  f32[ 99][21] = " ccc            "
  f32[ 99][22] = " ccc       ccc  "
  f32[ 99][23] = " ccc       ccc  "
  f32[ 99][24] = "  ccccccccccc   "
  f32[ 99][25] = "   ccccccccc    "

  f32[100][ 6] = "           ddd  "
  f32[100][ 7] = "           ddd  "
  f32[100][ 8] = "           ddd  "
  f32[100][ 9] = "           ddd  "
  f32[100][10] = "           ddd  "
  f32[100][11] = "           ddd  "
  f32[100][12] = "   ddddddddddd  "
  f32[100][13] = "  dddddddddddd  "
  f32[100][14] = " ddd       ddd  "
  f32[100][15] = " ddd       ddd  "
  f32[100][16] = " ddd       ddd  "
  f32[100][17] = " ddd       ddd  "
  f32[100][18] = " ddd       ddd  "
  f32[100][19] = " ddd       ddd  "
  f32[100][20] = " ddd       ddd  "
  f32[100][21] = " ddd       ddd  "
  f32[100][22] = " ddd       ddd  "
  f32[100][23] = " ddd       ddd  "
  f32[100][24] = "  dddddddddddd  "
  f32[100][25] = "   ddddddddddd  "

  f32[101][12] = "   eeeeeeeee    "
  f32[101][13] = "  eeeeeeeeeee   "
  f32[101][14] = " eee       eee  "
  f32[101][15] = " eee       eee  "
  f32[101][16] = " eee       eee  "
  f32[101][17] = " eee       eee  "
  f32[101][18] = " eeeeeeeeeeeee  "
  f32[101][19] = " eeeeeeeeeeeee  "
  f32[101][20] = " eee            "
  f32[101][21] = " eee            "
  f32[101][22] = " eee            "
  f32[101][23] = " eee       eee  "
  f32[101][24] = "  eeeeeeeeeee   "
  f32[101][25] = "   eeeeeeeee    "

  f32[102][ 6] = "       fffffff  "
  f32[102][ 7] = "      ffffffff  "
  f32[102][ 8] = "     fff        "
  f32[102][ 9] = "     fff        "
  f32[102][10] = "     fff        "
  f32[102][11] = "     fff        "
  f32[102][12] = " fffffffffff    "
  f32[102][13] = " fffffffffff    "
  f32[102][14] = "     fff        "
  f32[102][15] = "     fff        "
  f32[102][16] = "     fff        "
  f32[102][17] = "     fff        "
  f32[102][18] = "     fff        "
  f32[102][19] = "     fff        "
  f32[102][20] = "     fff        "
  f32[102][21] = "     fff        "
  f32[102][22] = "     fff        "
  f32[102][23] = "     fff        "
  f32[102][24] = "     fff        "
  f32[102][25] = "     fff        "

  f32[103][12] = "   ggggggggggg  "
  f32[103][13] = "  gggggggggggg  "
  f32[103][14] = " ggg       ggg  "
  f32[103][15] = " ggg       ggg  "
  f32[103][16] = " ggg       ggg  "
  f32[103][17] = " ggg       ggg  "
  f32[103][18] = " ggg       ggg  "
  f32[103][19] = " ggg       ggg  "
  f32[103][20] = " ggg       ggg  "
  f32[103][21] = " ggg       ggg  "
  f32[103][22] = " ggg       ggg  "
  f32[103][23] = " ggg       ggg  "
  f32[103][24] = "  gggggggggggg  "
  f32[103][25] = "   ggggggggggg  "
  f32[103][26] = "           ggg  "
  f32[103][27] = "           ggg  "
  f32[103][28] = "           ggg  "
  f32[103][29] = "  ggggggggggg   "
  f32[103][30] = "  gggggggggg    "

  f32[104][ 6] = " hhh            "
  f32[104][ 7] = " hhh            "
  f32[104][ 8] = " hhh            "
  f32[104][ 9] = " hhh            "
  f32[104][10] = " hhh            "
  f32[104][11] = " hhh            "
  f32[104][12] = " hhhhhhhhhhh    "
  f32[104][13] = " hhhhhhhhhhhh   "
  f32[104][14] = " hhh       hhh  "
  f32[104][15] = " hhh       hhh  "
  f32[104][16] = " hhh       hhh  "
  f32[104][17] = " hhh       hhh  "
  f32[104][18] = " hhh       hhh  "
  f32[104][19] = " hhh       hhh  "
  f32[104][20] = " hhh       hhh  "
  f32[104][21] = " hhh       hhh  "
  f32[104][22] = " hhh       hhh  "
  f32[104][23] = " hhh       hhh  "
  f32[104][24] = " hhh       hhh  "
  f32[104][25] = " hhh       hhh  "

  f32[105][ 6] = "      iii       "
  f32[105][ 7] = "      iii       "
  f32[105][ 8] = "      iii       "
  f32[105][ 9] = "      iii       "
  f32[105][10] = "                "
  f32[105][11] = "                "
  f32[105][12] = "    iiiii       "
  f32[105][13] = "    iiiii       "
  f32[105][14] = "      iii       "
  f32[105][15] = "      iii       "
  f32[105][16] = "      iii       "
  f32[105][17] = "      iii       "
  f32[105][18] = "      iii       "
  f32[105][19] = "      iii       "
  f32[105][20] = "      iii       "
  f32[105][21] = "      iii       "
  f32[105][22] = "      iii       "
  f32[105][23] = "      iii       "
  f32[105][24] = "    iiiiiii     "
  f32[105][25] = "    iiiiiii     "

  f32[106][ 6] = "          jjj   "
  f32[106][ 7] = "          jjj   "
  f32[106][ 8] = "          jjj   "
  f32[106][ 9] = "          jjj   "
  f32[106][10] = "                "
  f32[106][11] = "                "
  f32[106][12] = "        jjjjj   "
  f32[106][13] = "        jjjjj   "
  f32[106][14] = "          jjj   "
  f32[106][15] = "          jjj   "
  f32[106][16] = "          jjj   "
  f32[106][17] = "          jjj   "
  f32[106][18] = "          jjj   "
  f32[106][19] = "          jjj   "
  f32[106][20] = "          jjj   "
  f32[106][21] = "          jjj   "
  f32[106][22] = "          jjj   "
  f32[106][23] = "          jjj   "
  f32[106][24] = "          jjj   "
  f32[106][25] = "          jjj   "
  f32[106][26] = "  jjj     jjj   "
  f32[106][27] = "  jjj     jjj   "
  f32[106][28] = "  jjj     jjj   "
  f32[106][29] = "   jjjjjjjjj    "
  f32[106][30] = "    jjjjjjj     "

  f32[107][ 6] = "  kkk           "
  f32[107][ 7] = "  kkk           "
  f32[107][ 8] = "  kkk           "
  f32[107][ 9] = "  kkk           "
  f32[107][10] = "  kkk           "
  f32[107][11] = "  kkk           "
  f32[107][12] = "  kkk      kkk  "
  f32[107][13] = "  kkk     kkk   "
  f32[107][14] = "  kkk    kkk    "
  f32[107][15] = "  kkk   kkk     "
  f32[107][16] = "  kkk  kkk      "
  f32[107][17] = "  kkk kkk       "
  f32[107][18] = "  kkkkkk        "
  f32[107][19] = "  kkkkkk        "
  f32[107][20] = "  kkk kkk       "
  f32[107][21] = "  kkk  kkk      "
  f32[107][22] = "  kkk   kkk     "
  f32[107][23] = "  kkk    kkk    "
  f32[107][24] = "  kkk     kkk   "
  f32[107][25] = "  kkk      kkk  "

  f32[108][ 6] = "    lllll       "
  f32[108][ 7] = "    lllll       "
  f32[108][ 8] = "      lll       "
  f32[108][ 9] = "      lll       "
  f32[108][10] = "      lll       "
  f32[108][11] = "      lll       "
  f32[108][12] = "      lll       "
  f32[108][13] = "      lll       "
  f32[108][14] = "      lll       "
  f32[108][15] = "      lll       "
  f32[108][16] = "      lll       "
  f32[108][17] = "      lll       "
  f32[108][18] = "      lll       "
  f32[108][19] = "      lll       "
  f32[108][20] = "      lll       "
  f32[108][21] = "      lll       "
  f32[108][22] = "      lll       "
  f32[108][23] = "      lll       "
  f32[108][24] = "    lllllll     "
  f32[108][25] = "    lllllll     "

  f32[109][12] = " mmmmmmmmmmm    "
  f32[109][13] = " mmmmmmmmmmmm   "
  f32[109][14] = " mmm  mmm  mmm  "
  f32[109][15] = " mmm  mmm  mmm  "
  f32[109][16] = " mmm  mmm  mmm  "
  f32[109][17] = " mmm  mmm  mmm  "
  f32[109][18] = " mmm  mmm  mmm  "
  f32[109][19] = " mmm  mmm  mmm  "
  f32[109][20] = " mmm  mmm  mmm  "
  f32[109][21] = " mmm  mmm  mmm  "
  f32[109][22] = " mmm  mmm  mmm  "
  f32[109][23] = " mmm  mmm  mmm  "
  f32[109][24] = " mmm  mmm  mmm  "
  f32[109][25] = " mmm  mmm  mmm  "

  f32[110][12] = " nnnnnnnnnnn    "
  f32[110][13] = " nnnnnnnnnnnn   "
  f32[110][14] = " nnn       nnn  "
  f32[110][15] = " nnn       nnn  "
  f32[110][16] = " nnn       nnn  "
  f32[110][17] = " nnn       nnn  "
  f32[110][18] = " nnn       nnn  "
  f32[110][19] = " nnn       nnn  "
  f32[110][20] = " nnn       nnn  "
  f32[110][21] = " nnn       nnn  "
  f32[110][22] = " nnn       nnn  "
  f32[110][23] = " nnn       nnn  "
  f32[110][24] = " nnn       nnn  "
  f32[110][25] = " nnn       nnn  "

  f32[111][12] = "   ooooooooo    "
  f32[111][13] = "  ooooooooooo   "
  f32[111][14] = " ooo       ooo  "
  f32[111][15] = " ooo       ooo  "
  f32[111][16] = " ooo       ooo  "
  f32[111][17] = " ooo       ooo  "
  f32[111][18] = " ooo       ooo  "
  f32[111][19] = " ooo       ooo  "
  f32[111][20] = " ooo       ooo  "
  f32[111][21] = " ooo       ooo  "
  f32[111][22] = " ooo       ooo  "
  f32[111][23] = " ooo       ooo  "
  f32[111][24] = "  ooooooooooo   "
  f32[111][25] = "   ooooooooo    "

  f32[246][ 6] = "   ooo   ooo    "
  f32[246][ 7] = "   ooo   ooo    "
  f32[246][ 8] = "   ooo   ooo    "
  f32[246][ 9] = "   ooo   ooo    "
  f32[246][10] = "                "
  f32[246][11] = "                "
  f32[246][12] = "   ooooooooo    "
  f32[246][13] = "  ooooooooooo   "
  f32[246][14] = " ooo       ooo  "
  f32[246][15] = " ooo       ooo  "
  f32[246][16] = " ooo       ooo  "
  f32[246][17] = " ooo       ooo  "
  f32[246][18] = " ooo       ooo  "
  f32[246][19] = " ooo       ooo  "
  f32[246][20] = " ooo       ooo  "
  f32[246][21] = " ooo       ooo  "
  f32[246][22] = " ooo       ooo  "
  f32[246][23] = " ooo       ooo  "
  f32[246][24] = "  ooooooooooo   "
  f32[246][25] = "   ooooooooo    "

  f32[112][12] = " ppppppppppp    "
  f32[112][13] = " pppppppppppp   "
  f32[112][14] = " ppp       ppp  "
  f32[112][15] = " ppp       ppp  "
  f32[112][16] = " ppp       ppp  "
  f32[112][17] = " ppp       ppp  "
  f32[112][18] = " ppp       ppp  "
  f32[112][19] = " ppp       ppp  "
  f32[112][20] = " ppp       ppp  "
  f32[112][21] = " ppp       ppp  "
  f32[112][22] = " ppp       ppp  "
  f32[112][23] = " ppp       ppp  "
  f32[112][24] = " pppppppppppp   "
  f32[112][25] = " ppppppppppp    "
  f32[112][26] = " ppp            "
  f32[112][27] = " ppp            "
  f32[112][28] = " ppp            "
  f32[112][29] = " ppp            "
  f32[112][30] = " ppp            "

  f32[113][12] = "   qqqqqqqqqqq  "
  f32[113][13] = "  qqqqqqqqqqqq  "
  f32[113][14] = " qqq       qqq  "
  f32[113][15] = " qqq       qqq  "
  f32[113][16] = " qqq       qqq  "
  f32[113][17] = " qqq       qqq  "
  f32[113][18] = " qqq       qqq  "
  f32[113][19] = " qqq       qqq  "
  f32[113][20] = " qqq       qqq  "
  f32[113][21] = " qqq       qqq  "
  f32[113][22] = " qqq       qqq  "
  f32[113][23] = " qqq       qqq  "
  f32[113][24] = "  qqqqqqqqqqqq  "
  f32[113][25] = "   qqqqqqqqqqq  "
  f32[113][26] = "           qqq  "
  f32[113][27] = "           qqq  "
  f32[113][28] = "           qqq  "
  f32[113][29] = "           qqq  "
  f32[113][30] = "           qqq  "

  f32[114][12] = " rrr  rrrrrrrr  "
  f32[114][13] = " rrr rrrrrrrrr  "
  f32[114][14] = " rrrrrr         "
  f32[114][15] = " rrrrr          "
  f32[114][16] = " rrrr           "
  f32[114][17] = " rrr            "
  f32[114][18] = " rrr            "
  f32[114][19] = " rrr            "
  f32[114][20] = " rrr            "
  f32[114][21] = " rrr            "
  f32[114][22] = " rrr            "
  f32[114][23] = " rrr            "
  f32[114][24] = " rrr            "
  f32[114][25] = " rrr            "

  f32[115][12] = "   sssssssss    "
  f32[115][13] = "  sssssssssss   "
  f32[115][14] = " sss       sss  "
  f32[115][15] = " sss            "
  f32[115][16] = " sss            "
  f32[115][17] = " sss            "
  f32[115][18] = "  ssssssssss    "
  f32[115][19] = "   ssssssssss   "
  f32[115][20] = "           sss  "
  f32[115][21] = "           sss  "
  f32[115][22] = "           sss  "
  f32[115][23] = " sss       sss  "
  f32[115][24] = "  sssssssssss   "
  f32[115][25] = "   sssssssss    "

  f32[116][ 6] = "     ttt        "
  f32[116][ 7] = "     ttt        "
  f32[116][ 8] = "     ttt        "
  f32[116][ 9] = "     ttt        "
  f32[116][10] = "     ttt        "
  f32[116][11] = "     ttt        "
  f32[116][12] = " ttttttttttt    "
  f32[116][13] = " ttttttttttt    "
  f32[116][14] = "     ttt        "
  f32[116][15] = "     ttt        "
  f32[116][16] = "     ttt        "
  f32[116][17] = "     ttt        "
  f32[116][18] = "     ttt        "
  f32[116][19] = "     ttt        "
  f32[116][20] = "     ttt        "
  f32[116][21] = "     ttt        "
  f32[116][22] = "     ttt        "
  f32[116][23] = "     ttt        "
  f32[116][24] = "      tttttttt  "
  f32[116][25] = "       ttttttt  "

  f32[117][12] = " uuu       uuu  "
  f32[117][13] = " uuu       uuu  "
  f32[117][14] = " uuu       uuu  "
  f32[117][15] = " uuu       uuu  "
  f32[117][16] = " uuu       uuu  "
  f32[117][17] = " uuu       uuu  "
  f32[117][18] = " uuu       uuu  "
  f32[117][19] = " uuu       uuu  "
  f32[117][20] = " uuu       uuu  "
  f32[117][21] = " uuu       uuu  "
  f32[117][22] = " uuu       uuu  "
  f32[117][23] = " uuu       uuu  "
  f32[117][24] = "  uuuuuuuuuuuu  "
  f32[117][25] = "   uuuuuuuuuuu  "

  f32[252][ 6] = "   uuu   uuu    "
  f32[252][ 7] = "   uuu   uuu    "
  f32[252][ 8] = "   uuu   uuu    "
  f32[252][ 9] = "   uuu   uuu    "
  f32[252][10] = "                "
  f32[252][11] = "                "
  f32[252][12] = " uuu       uuu  "
  f32[252][13] = " uuu       uuu  "
  f32[252][14] = " uuu       uuu  "
  f32[252][15] = " uuu       uuu  "
  f32[252][16] = " uuu       uuu  "
  f32[252][17] = " uuu       uuu  "
  f32[252][18] = " uuu       uuu  "
  f32[252][19] = " uuu       uuu  "
  f32[252][20] = " uuu       uuu  "
  f32[252][21] = " uuu       uuu  "
  f32[252][22] = " uuu       uuu  "
  f32[252][23] = " uuu       uuu  "
  f32[252][24] = "  uuuuuuuuuuuu  "
  f32[252][25] = "   uuuuuuuuuuu  "

  f32[118][12] = " vvv       vvv  "
  f32[118][13] = " vvv       vvv  "
  f32[118][14] = " vvv       vvv  "
  f32[118][15] = "  vvv     vvv   "
  f32[118][16] = "  vvv     vvv   "
  f32[118][17] = "  vvv     vvv   "
  f32[118][18] = "   vvv   vvv    "
  f32[118][19] = "   vvv   vvv    "
  f32[118][20] = "   vvv   vvv    "
  f32[118][21] = "    vvv vvv     "
  f32[118][22] = "    vvv vvv     "
  f32[118][23] = "    vvv vvv     "
  f32[118][24] = "     vvvvv      "
  f32[118][25] = "     vvvvv      "

  f32[119][12] = " www       www  "
  f32[119][13] = " www       www  "
  f32[119][14] = " www       www  "
  f32[119][15] = " www       www  "
  f32[119][16] = " www  www  www  "
  f32[119][17] = " www  www  www  "
  f32[119][18] = " www  www  www  "
  f32[119][19] = " www  www  www  "
  f32[119][20] = " www  www  www  "
  f32[119][21] = " www  www  www  "
  f32[119][22] = " www  www  www  "
  f32[119][23] = " www  www  www  "
  f32[119][24] = "  wwwwwwwwwww   "
  f32[119][25] = "   wwwwwwwww    "

  f32[120][12] = " xxx       xxx  "
  f32[120][13] = " xxx       xxx  "
  f32[120][14] = " xxx       xxx  "
  f32[120][15] = "  xxx     xxx   "
  f32[120][16] = "   xxx   xxx    "
  f32[120][17] = "    xxx xxx     "
  f32[120][18] = "     xxxxx      "
  f32[120][19] = "     xxxxx      "
  f32[120][20] = "    xxx xxx     "
  f32[120][21] = "   xxx   xxx    "
  f32[120][22] = "  xxx     xxx   "
  f32[120][23] = " xxx       xxx  "
  f32[120][24] = " xxx       xxx  "
  f32[120][25] = " xxx       xxx  "

  f32[121][12] = " yyy       yyy  "
  f32[121][13] = " yyy       yyy  "
  f32[121][14] = " yyy       yyy  "
  f32[121][15] = " yyy       yyy  "
  f32[121][16] = " yyy       yyy  "
  f32[121][17] = " yyy       yyy  "
  f32[121][18] = " yyy       yyy  "
  f32[121][19] = " yyy       yyy  "
  f32[121][20] = " yyy       yyy  "
  f32[121][21] = " yyy       yyy  "
  f32[121][22] = " yyy       yyy  "
  f32[121][23] = " yyy       yyy  "
  f32[121][24] = "  yyyyyyyyyyyy  "
  f32[121][25] = "   yyyyyyyyyyy  "
  f32[121][26] = "           yyy  "
  f32[121][27] = "           yyy  "
  f32[121][28] = "           yyy  "
  f32[121][29] = "  yyyyyyyyyyy   "
  f32[121][30] = "  yyyyyyyyyy    "

  f32[122][12] = " zzzzzzzzzzzzz  "
  f32[122][13] = " zzzzzzzzzzzzz  "
  f32[122][14] = "          zzz   "
  f32[122][15] = "         zzz    "
  f32[122][16] = "        zzz     "
  f32[122][17] = "       zzz      "
  f32[122][18] = "      zzz       "
  f32[122][19] = "     zzz        "
  f32[122][20] = "    zzz         "
  f32[122][21] = "   zzz          "
  f32[122][22] = "  zzz           "
  f32[122][23] = " zzz            "
  f32[122][24] = " zzzzzzzzzzzzz  "
  f32[122][25] = " zzzzzzzzzzzzz  "

  f32[123][ 6] = "       {{{{{    "
  f32[123][ 7] = "      {{{{{{    "
  f32[123][ 8] = "     {{{        "
  f32[123][ 9] = "     {{{        "
  f32[123][10] = "     {{{        "
  f32[123][11] = "     {{{        "
  f32[123][12] = "     {{{        "
  f32[123][13] = "     {{{        "
  f32[123][14] = "     {{{        "
  f32[123][15] = "   {{{{         "
  f32[123][16] = "   {{{{         "
  f32[123][17] = "     {{{        "
  f32[123][18] = "     {{{        "
  f32[123][19] = "     {{{        "
  f32[123][20] = "     {{{        "
  f32[123][21] = "     {{{        "
  f32[123][22] = "     {{{        "
  f32[123][23] = "     {{{        "
  f32[123][24] = "      {{{{{{    "
  f32[123][25] = "       {{{{{    "

  f32[124][ 6] = "      |||       "
  f32[124][ 7] = "      |||       "
  f32[124][ 8] = "      |||       "
  f32[124][ 9] = "      |||       "
  f32[124][10] = "      |||       "
  f32[124][11] = "      |||       "
  f32[124][12] = "      |||       "
  f32[124][13] = "      |||       "
  f32[124][14] = "      |||       "
  f32[124][15] = "      |||       "
  f32[124][16] = "      |||       "
  f32[124][17] = "      |||       "
  f32[124][18] = "      |||       "
  f32[124][19] = "      |||       "
  f32[124][20] = "      |||       "
  f32[124][21] = "      |||       "
  f32[124][22] = "      |||       "
  f32[124][23] = "      |||       "
  f32[124][24] = "      |||       "
  f32[124][25] = "      |||       "

  f32[125][ 6] = "   }}}}}        "
  f32[125][ 7] = "   }}}}}}       "
  f32[125][ 8] = "       }}}      "
  f32[125][ 9] = "       }}}      "
  f32[125][10] = "       }}}      "
  f32[125][11] = "       }}}      "
  f32[125][12] = "       }}}      "
  f32[125][13] = "       }}}      "
  f32[125][14] = "       }}}      "
  f32[125][15] = "        }}}}    "
  f32[125][16] = "        }}}}    "
  f32[125][17] = "       }}}      "
  f32[125][18] = "       }}}      "
  f32[125][19] = "       }}}      "
  f32[125][20] = "       }}}      "
  f32[125][21] = "       }}}      "
  f32[125][22] = "       }}}      "
  f32[125][23] = "       }}}      "
  f32[125][24] = "   }}}}}}       "
  f32[125][25] = "   }}}}}        "

  f32[126][ 3] = "   ~~~~    ~~~  "
  f32[126][ 4] = "  ~~~~~~   ~~~  "
  f32[126][ 5] = " ~~~  ~~~  ~~~  "
  f32[126][ 6] = " ~~~  ~~~  ~~~  "
  f32[126][ 7] = " ~~~   ~~~~~~   "
  f32[126][ 8] = " ~~~    ~~~~    "

  f32[164][ 6] = "      eeeeee    "
  f32[164][ 7] = "     eeeeeeee   "
  f32[164][ 8] = "    eee    eee  "
  f32[164][ 9] = "   eee      eee "
  f32[164][10] = "  eee           "
  f32[164][11] = "  eee           "
  f32[164][12] = "  eee           "
  f32[164][13] = "eeeeeeeeeee     "
  f32[164][14] = "eeeeeeeeeee     "
  f32[164][15] = "  eee           "
  f32[164][16] = "  eee           "
  f32[164][17] = "eeeeeeeeeee     "
  f32[164][18] = "eeeeeeeeeee     "
  f32[164][19] = "  eee           "
  f32[164][20] = "  eee           "
  f32[164][21] = "  eee           "
  f32[164][22] = "   eee      eee "
  f32[164][23] = "    eee    eee  "
  f32[164][24] = "     eeeeeeee   "
  f32[164][25] = "      eeeeee    "

  f32[167][ 4] = "    ppppppp     "
  f32[167][ 5] = "   ppppppppp    "
  f32[167][ 6] = "  ppp     ppp   "
  f32[167][ 7] = "  ppp     ppp   "
  f32[167][ 8] = "  ppp           "
  f32[167][ 9] = "  pppp          "
  f32[167][10] = "   ppppppp      "
  f32[167][11] = "   pppppppp     "
  f32[167][12] = "  ppp    ppp    "
  f32[167][13] = "  ppp     ppp   "
  f32[167][14] = "  ppp     ppp   "
  f32[167][15] = "  ppp     ppp   "
  f32[167][16] = "  ppp     ppp   "
  f32[167][17] = "  ppp     ppp   "
  f32[167][18] = "   ppp    ppp   "
  f32[167][19] = "    pppppppp    "
  f32[167][20] = "     ppppppp    "
  f32[167][21] = "         pppp   "
  f32[167][22] = "          ppp   "
  f32[167][23] = "  ppp     ppp   "
  f32[167][24] = "  ppp     ppp   "
  f32[167][25] = "   ppppppppp    "
  f32[167][26] = "    ppppppp     "

  f32[176][ 3] = "    ooooooo     "
  f32[176][ 4] = "   ooooooooo    "
  f32[176][ 5] = "   ooo   ooo    "
  f32[176][ 6] = "   ooo   ooo    "
  f32[176][ 7] = "   ooo   ooo    "
  f32[176][ 8] = "   ooo   ooo    "
  f32[176][ 9] = "   ooooooooo    "
  f32[176][10] = "    ooooooo     "

  f32[178][ 3] = "     222222     "
  f32[178][ 4] = "    22222222    "
  f32[178][ 5] = "    222  222    "
  f32[178][ 6] = "    222  222    "
  f32[178][ 7] = "        222     "
  f32[178][ 8] = "       222      "
  f32[178][ 9] = "      222       "
  f32[178][10] = "     2222222    "
  f32[178][11] = "    22222222    "

  f32[179][ 3] = "    3333333     "
  f32[179][ 4] = "   333333333    "
  f32[179][ 5] = "   333   333    "
  f32[179][ 6] = "      33333     "
  f32[179][ 7] = "      33333     "
  f32[179][ 8] = "         333    "
  f32[179][ 9] = "   333   333    "
  f32[179][10] = "   333333333    "
  f32[179][11] = "    3333333     "

  f32[181][12] = " mmm       mmm  "
  f32[181][13] = " mmm       mmm  "
  f32[181][14] = " mmm       mmm  "
  f32[181][15] = " mmm       mmm  "
  f32[181][16] = " mmm       mmm  "
  f32[181][17] = " mmm       mmm  "
  f32[181][18] = " mmm       mmm  "
  f32[181][19] = " mmm       mmm  "
  f32[181][20] = " mmm       mmm  "
  f32[181][21] = " mmm      mmmm  "
  f32[181][22] = " mmm     mmmmm  "
  f32[181][23] = " mmm    mmmmmm  "
  f32[181][24] = " mmmmmmmmm mmm  "
  f32[181][25] = " mmmmmmmm  mmm  "
  f32[181][26] = " mmm            "
  f32[181][27] = " mmm            "
  f32[181][28] = " mmm            "
  f32[181][29] = " mmm            "
  f32[181][30] = " mmm            "

  f32[223][ 6] = "   ssssssss     "
  f32[223][ 7] = "  ssssssssss    "
  f32[223][ 8] = " sss      sss   "
  f32[223][ 9] = " sss      sss   "
  f32[223][10] = " sss      sss   "
  f32[223][11] = " sss      sss   "
  f32[223][12] = " sss      sss   "
  f32[223][13] = " sss     sss    "
  f32[223][14] = " sss ssssss     "
  f32[223][15] = " sss sssssss    "
  f32[223][16] = " sss      sss   "
  f32[223][17] = " sss       sss  "
  f32[223][18] = " sss       sss  "
  f32[223][19] = " sss       sss  "
  f32[223][19] = " sss       sss  "
  f32[223][20] = " sss       sss  "
  f32[223][21] = " sss       sss  "
  f32[223][22] = " sss       sss  "
  f32[223][23] = " sss ss    sss  "
  f32[223][24] = " sss  sssssss   "
  f32[223][25] = " sss   sssss    "
  f32[223][26] = " sss            "
  f32[223][27] = " sss            "
  f32[223][28] = " sss            "
  f32[223][29] = " sss            "
  f32[223][30] = " sss            "
}

// font ////////////////////////////////////////////////////////////////

func (X *console) ActFontsize() font.Size {
  return X.fontsize
}

func (X *console) SetFontsize (f font.Size) {
  if X == nil { panic ("quadratische Kacke") }
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
  return x < int(X.wd) && y < int(X.ht) && x1 < int(X.wd) && y1 < int(X.ht) // shit
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
  if y <= y1 { // Steigung positiv
    dy = y1 - y
    if dy <= dx { // Steigung <= 45°
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
    } else { // Steigung > 45°
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
  } else { // Steigung negativ
    dy = y - y1
    if dy <= dx { // Steigung >= -45°
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
    } else { // Steigung < -45°
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
  if y <= y1 { // Steigung positiv
    dy = y1 - y
    if dy <= dx { // Steigung <= 45°
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
    } else { // Steigung > 45°
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
    dy = y - y1 // Steigung negativ
    if dy <= dx { // Steigung >= -45°
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
    } else { // Steigung < -45°
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
  X.segs (xs, ys, X.Point)
  n := len (xs)
  if n > 1 {
    X.line (xs[n-1], ys[n-1], xs[0], ys[0], X.Point)
  }
}

func (X *console) PolygonInv (xs, ys []int) {
  if ! ok2 (xs, ys) { return }
  X.segs (xs, ys, X.PointInv)
  n := len (xs)
  if n > 1 {
    X.line (xs[n-1], ys[n-1], xs[0], ys[0], X.PointInv)
    X.PointInv (xs[0], ys[0])
    X.PointInv (xs[n-1], ys[n-1])
  }
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
//  if ! between (x - int(r), x + int(r), a) { return false }
/*
  if r == 0 { return a == x && b == y }
  z = a * a + b * b
  if z > r * r { z = z - r * r } else { z = r * r - z }
*/
  X.xx_, X.yy_, X.tt_, X.incident = a, b, int(t * t), false
  X.circ (x, y, r, false, X.onPoint)
  return X.incident
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

func (X *console) MouseEx() bool {
  return mouse.Ex()
}

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
  if ! mouse.Ex() || ! X.mouseOn || ! visible { return }
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
  if ! mouse.Ex() { return false }
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
  if ! mouse.Ex() { return false }
  lm, cm := X.MousePos()
  return l <= lm && lm < l + h && c <= cm && cm < c + w
}

func (X *console) UnderMouseGr (x, y, x1, y1 int, t uint) bool {
  if ! mouse.Ex() { return false }
  intord (&x, &y, &x1, &y1)
  xm, ym := X.MousePosGr()
  return x <= int(xm) + int(t) && int(xm) <= x1 + int(t) &&
         y <= int(ym) + int(t) && int(ym) <= y1 + int(t)
}

func (X *console) UnderMouse1 (x, y int, d uint) bool {
  if ! mouse.Ex() { return false }
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
  w, h, _, _ := ppmHeaderData (s)
  return w, h
}

func (X *console) PPMEncode (x0, y0, w, h uint) obj.Stream {
  s := X.Encode (x0, y0, w, h)
  return append (obj.Stream(X.PPMHeader (w, h)), s[2*4:]...)
}

func ppmHeaderData (s obj.Stream) (uint, uint, uint, int) {
  p := string(s[:2]); if p != "P6" { panic ("wrong ppm-header: " + p) }
  i := 3
  w, dw := number (s[i:])
  i += dw + 1
  h, dh := number (s[i:])
  i += dh + 1
  m, dm := number (s[i:])
  i += dm
  return w, h, m, i + 1
}

func (X *console) PPMDecode (st obj.Stream, x0, y0 uint) {
  w, h, _, j := ppmHeaderData (st)
  if w == 0 || h == 0 || w > X.Wd() || h > X.Ht() { return }
  i := 4 * uint(2)
  l := i + 3 * w * h
  e := make(obj.Stream, l)
  copy (e[:i], obj.Encode4 (uint16(x0), uint16(y0), uint16(w), uint16(h)))
  if under_S {
    ker.Panic ("underC and under_S")
  } else if under_X {
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
