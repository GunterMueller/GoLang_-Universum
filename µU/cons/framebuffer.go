package cons

// (c) Christian Maurer   v. 210106 - license see µU.go

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
  "syscall"
  "strconv"
  . "µU/obj"
  . "µU/shape"
  . "µU/mode"
  "µU/ker"
  "µU/time"
)
const (
  esc1 = "\x1b["
  ClearScreen = esc1 + "H" + esc1 + "J"
  home = esc1 + "?25h" + esc1 + "?0c"
)
var (
  fbmemsize uint
  fbmem, fbcop,
  emptyBackground Stream
  visible bool // only for console switching
)

func consoleOn() {
  ker.ActivateConsole()
  n := width * height * uint(colourdepth)
  copy (fbmem[:n], fbcop[:n])
  visible = true
  c := actual
  c.Warp (c.ht1 * c.blinkY, c.wd1 * c.blinkX, c.consoleShape)
}

func consoleOff() {
  visible = false
  c := actual
  c.consoleShape = c.blinkShape
  c.Warp (c.ht1 * c.blinkY, c.wd1 * c.blinkX, Off)
  ker.DeactivateConsole()
}

func consoleFin() {
// TODO wait (blink())
// TODO fin (blink())
  c := actual
  finished = true
  time.Msleep (250) // provisorial
  c.cursorShape = Off
  print (ClearScreen + home)
}

var
  initialized bool

func maxResConsole() (uint, uint) {
  if framebufferOk() {
    return width, height
  }
  return 0, 0
}

func framebuffer() (x, y, b uint, fb Stream) {
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
  fullScreen = ModeOf (width, height)
  if Wd(fullScreen) != width || Ht(fullScreen) != height { ker.Panic ("fullScreen bug") }
  colourdepth = colbits / 8
  fbmemsize = width * height * colourdepth
  if uint(len (fbmem)) != fbmemsize {
    ker.Panic ("len (fbmem) == " + strconv.Itoa(len(fbmem)) +
               " != fbmemsize == " + strconv.Itoa(int(fbmemsize)))
  }
  fbcop = make(Stream, fbmemsize)
  emptyBackground = make(Stream, fbmemsize)
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
