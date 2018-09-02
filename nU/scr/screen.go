package scr

// (c) Christian Maurer   v. 180812 - license see nU.go

// #include <stdlib.h>
// #include <string.h>
// #include <unistd.h>
// #include <fcntl.h>
// #include <sys/ioctl.h>
// #include <linux/fb.h>
// #include <linux/vt.h>
/*
void wh (int *w, int *h) {
  struct fb_var_screeninfo v;
  int fd;
  *w = 0;
  *h = 0;
  if ((fd = open ("/dev/fb0", O_RDWR)) == -1) { return; }
  if (ioctl (fd, FBIOGET_VSCREENINFO, &v) == -1) { close (fd); return; }
  *w = v.xres;
  *h = v.yres;
}

void write_(const void *b) {
  write (1, b, strlen(b));
}
*/
import "C"
import ("unsafe";
        // "os/exec";
        "strconv"; "sync"; . "nU/obj"; "nU/col")

const (null = byte('0'); esc = byte(27))

type screen struct {
  w, h, nL, nC, lmax uint
  mutex sync.Mutex
}

var a *screen

func new_() Screen {
  x := new(screen)
  var cw, ch C.int
  C.wh (&cw, &ch)
  x.h, x.w = uint(ch), uint(cw)
  x.nL, x.nC = x.h / Ht1, x.w / Wd1
  x.Cls()
  a = x
  return x
}

func (x *screen) Wd() uint {
  return x.w
}

func (x *screen) Ht() uint {
  return x.h
}

func (x *screen) NLines() uint {
  return x.nL
}

func (x *screen) NColumns() uint {
  return x.nC
}

func (x *screen) Cls() {
  s := Stream(" [H [J")
  s[0] = esc
  s[3] = esc
  cs := C.CString(string(s)); defer C.free(unsafe.Pointer(cs))
  C.write_(unsafe.Pointer(cs))
}

func (x *screen) set (f col.Colour, b bool) {
  t := "4"; if b { t = "3" }
  s := []byte(" [" + t + "8;2;" + strconv.Itoa(int(f.R())) + ";" +
                                  strconv.Itoa(int(f.G())) + ";" +
                                  strconv.Itoa(int(f.B())) + "m")
  s[0] = 27
  cs := C.CString(string(s))
  defer C.free(unsafe.Pointer(cs))
  C.write_(unsafe.Pointer(cs))
}

func (x *screen) ColourF (f col.Colour) {
  x.set (f, true)
}

func (x *screen) ColourB (b col.Colour) {
  x.set (b, false)
}

func (x *screen) Colours (f, b col.Colour) {
  x.set (f, true)
  x.set (b, false)
}

func (x *screen) Switch (on bool) {
  s := Stream(" [25l")
  s[0] = esc
  if on {
    s = Stream(" [?25h") // [?0c")
    s[0] = esc
//    s[6] = esc
  }
  cs := C.CString(string(s))
  defer C.free(unsafe.Pointer(cs))
/*
  o := "off"
  if on { o = "on" }
  exec.Command ("setterm", "cursor", o).Run()
*/
}

func (x *screen) Warp (l, c uint) {
  l++
  c++
  s := Stream(" [  ;   H")
  s[0] = esc
  s[2] = byte(l) / 10 + null
  s[3] = byte(l) % 10 + null
  s[5] = byte(c) / 100 + null
  c %= 100
  s[6] = byte(c) / 10 + null
  s[7] = byte(c) % 10 + null
  cs := C.CString(string(s))
  defer C.free(unsafe.Pointer(cs))
  C.write_(unsafe.Pointer(cs))
}

func (x *screen) Fin() {
  x.set (col.White(), true)
  x.set (col.Black(), false)
  x.Warp (x.lmax + 1, 0)
  x.Switch (true)
}

func (x *screen) Write1 (b byte, l, c uint) {
  if l > x.lmax { x.lmax = l }
  x.Write (string(b), l, c)
}

func (x *screen) Write (s string, l, c uint) {
  if l > x.lmax { x.lmax = l }
  x.Warp (l, c)
  x.Switch (false)
  cs := C.CString(s)
  defer C.free(unsafe.Pointer(cs))
  C.write_(unsafe.Pointer(cs))
  x.Switch (false)
}

func (x *screen) WriteNat (n, l, c uint) {
  if l > x.lmax { x.lmax = l }
  s := strconv.Itoa(int(n))
  x.Write (s, l, c)
}

func (x *screen) Line (l, c, l1, c1 uint) {
  if l > x.lmax { x.lmax = l }
  if l1 > x.lmax { x.lmax = l1 }
  if l >= x.nL || l1 >= x.nL { return }
  if c >= x.nC || c1 >= x.nC { return }
  if l == l1 && c == c1 {
    x.Write1 ('*', l, c)
  }
  if l > l1 {
    c, c1 = c1, c
    l, l1 = l1, l
  }
  la, lb := l, l1
  if l1 < l {
    la, lb = lb, la
  }
  ca, cb := c, c1
  if c1 < c {
    ca, cb = cb, ca
  }
  if c == c1 {
    for la <= lb {
      x.Write1 (byte('|'), la, c)
      la++
    }
    return
  } // c != c1
  if l == l1 {
    for ca <= cb {
      x.Write1 (byte('-'), l, ca)
      ca++
    }
    return
  }
  b := byte('\\')
  if l < l1 && c > c1 || l > l1 && c < c1 {
    b = '/'
  }
  dl, dc := int(l1 - l), 0
  f := 0
  if c <= c1 {
    dc = int(c1 - c)
    if dc <= dl {
      for {
        x.Write1 (b, l, c)
        if l == l1 { break }
        l++
        f += 2 * dc
        if f > dl {
          c++
          f -= 2 * dl
        }
      }
    } else {
      for {
        x.Write1 (b, l, c)
        if c == c1 { break }
        c++
        f += 2 * dl
        if f > dc {
          l++
          f -= 2 * dc
        }
      }
    }
  } else {
    dc = int(c - c1)
    if dc <= dl {
      for {
        x.Write1 (b, l, c)
        if l == l1 { break }
        l++
        f += 2 * dc
        if f > dl {
          c--
          f -= 2 * dl
        }
      }
    } else {
      for {
        x.Write1 (b, l, c)
        if c == c1 { break }
        c--
        f += 2 * dl
        if f > dc {
          l++
          f -= 2 * dc
        }
      }
    }
  }
}

func (x *screen) Circle (c, l, r uint) {
  if l < r || l + r >= x.Ht() || c < r || c + r >= x.Wd() { return }
  if r == 0 {
    x.Write ("*", l, c)
    return
  }
  c1, l1 := uint(0), r
  f := 3
  f -= 2 * int(r)
  b := byte('*')
  for c1 <= l1 {
    x.Write1 (b, c - l1, 2 * (l - c1))
    x.Write1 (b, c + l1, 2 * (l - c1))
    x.Write1 (b, c - l1, 2 * (l + c1))
    x.Write1 (b, c + l1, 2 * (l + c1))
    x.Write1 (b, c - c1, 2 * (l - l1))
    x.Write1 (b, c + c1, 2 * (l - l1))
    x.Write1 (b, c - c1, 2 * (l + l1))
    x.Write1 (b, c + c1, 2 * (l + l1))
    c1++
    if f >= 0 {
      l1--
      f -= 4 * int(l1)
    }
    f += 4 * int(c1) + 2
  }
}

func (x *screen) Lock() {
  x.mutex.Lock()
}

func (x *screen) Unlock() {
  x.mutex.Unlock()
}
