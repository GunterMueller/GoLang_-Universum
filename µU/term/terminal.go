package term

// (c) Christian Maurer   v. 180106 - license see µU.go

// #include <termios.h>
// #include <unistd.h>
// #include <asm-generic/ioctls.h>
// #include <linux/kd.h>
/*
struct termios t;

void new (void) {
  ioctl (0, TCGETS, &t);
  struct termios t1;
  ioctl (0, TCGETS, &t1);
  t1.c_iflag = 0;
  t1.c_lflag &= ~(ISIG | ICANON | ECHO );
  t1.c_cc[VTIME] = 1;
  t1.c_cc[VMIN] = 3;
  ioctl (0, TCSETSW, &t1);
  ioctl (0, KDSKBMODE, K_MEDIUMRAW);
}

void fin (void) {
  ioctl (0, KDSKBMODE, K_XLATE);
  ioctl (0, TCSETS, &t);
}

char read_(void) {
  char b;
  read (0, &b, 1);
  return b;
}
*/
import
  "C"
type
  terminal struct {
  bool "active"
  byte
}
var
  a *terminal

func new_() Terminal {
  x := new(terminal)
  C.new()
  x.bool = true
  a = x
  return x
}

func (x *terminal) Read() byte {
  x.byte = byte(C.read_())
  return x.byte
}

func (x *terminal) Fin() {
  if x.bool {
    C.fin()
    x.bool = false
  }
}
