package ker

// (c) murus.org  v. 130217 - license see murus.go

//#include <stdlib.h>
//#include <fcntl.h>
//#include <unistd.h>
//#include <sys/ioctl.h>
//#include <sys/mman.h>
//#include <linux/vt.h>
//#include <linux/fb.h>
/*
void * framebuffer (int *x, int *y, int *b, int *a)
{
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
	"reflect"; "unsafe"
)

func Framebuffer() (x, y, b uint, fb []byte) {
  var xc, yc, bc, ac C.int
  f:= C.framebuffer (&xc, &yc, &bc, &ac)
  x, y, b = uint(xc), uint(yc), uint(bc)
  h:= (*reflect.SliceHeader)((unsafe.Pointer(&fb)))
  m:= int(x * y * ((b + 4) / 8)) // possible bitsizes 4, 15 !
  h.Cap, h.Len, h.Data = m, m, uintptr(f)
  if fb == nil {}
  return
}
