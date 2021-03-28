package internal

// (c) Christian Maurer   v. 210324 - license see µU.go

// #include <termios.h>
// #include <unistd.h>
// #include <sys/ioctl.h>
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

void fin_(void) {
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
var
  active bool

func new_() {
  C.new()
  active = true
}

func read() byte {
  if ! active {
    return 0
  }
  return byte(C.read_())
}

func fin() {
  if active {
    C.fin_()
    active = false
  }
}
